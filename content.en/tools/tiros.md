# Tiros

Tiros is an IPFS website measurement tool. It orchestrates the interplay between
a Kubo node and a headless Chrome instance. The goal is to measure website
metrics like [TTFB](https://web.dev/ttfb) or [FCP](https://web.dev/fcp) when
loaded over IPFS and compare them with their HTTPS performance equivalents.
For that, tiros instructs Chrome to request a website through the gateway
of a local Kubo node or regularly over HTTPS.

{{< button href="https://github.com/dennis-tra/nebula" >}}GitHub{{< /button >}}

## Deployment

We are running Tiros on ECS as a scheduled task every six hours in the following
seven AWS regions:

```
eu-central-1
ap-south-1
af-southeast-2
sa-east-1
us-east-2
us-west-1
af-south-1
```

The list of websites that we are currently monitoring is configured [here](https://github.com/protocol/probelab-infra/blob/aabe20d28e833364e0ed17e651d5f810e524cbb9/aws/tf/modules/tiros/_variables.tf#L49).
At the time of writing the list is:

```
filecoin.io
libp2p.io
ipld.io
protocol.ai
docs.ipfs.tech
tldr.filecoin.io
green.filecoin.io
drand.love
web3.storage
consensuslab.world
strn.network
research.protocol.ai
ipfs.tech
blog.ipfs.tech
docs.libp2p.io
blog.libp2p.io
en.wikipedia-on-ipfs.org/wiki
```

Each ECS task consists of three containers:

1. `scheduler` - [tiros repository](https://github.com/plprobelab/tiros)
2. `chrome` - via [`browserless/chrome`](https://github.com/browserless/chrome)
3. `kubo` - via [ipfs/kubo](https://hub.docker.com/r/ipfs/kubo/)

`kubo` is running with `LIBP2P_RCMGR=0` which disables the [libp2p Network Resource Manager](https://github.com/libp2p/go-libp2p-resource-manager#readme).

The `scheduler` gets configured with a list of websites that will then be probed. A typical website config looks like this `ipfs.io,docs.libp2p.io,ipld.io`. The scheduler probes each website via `kubo` by requesting `http://localhost:8080/ipns/<website>` and via HTTP  by requesting`https://<website>`. Port `8080` is the default `kubo` HTTP-Gateway port. The `scheduler` uses [`go-rod`](https://github.com/go-rod/rod) to communicate with the `browserless/chrome` instance. The following excerpt is a gist of what's happening when requesting a website:

```go
browser := rod.New().Context(ctx).ControlURL("ws://localhost:3000")) // default CDP chrome port

browser.Connect()
defer browser.Close()

var metricsStr string
rod.Try(func() {
    browser = browser.Context(c.Context).MustIncognito()
    browser.MustSetCookies() // clears cookies if parameter is empty
    
    page := browser.MustPage() // Get a handle of a new page in our incognito browser
    
    page.MustEvalOnNewDocument(jsOnNewDocument) // clears the cache by running `localStorage.clear()`
    
    // disable caching in general
    proto.NetworkSetCacheDisabled{CacheDisabled: true}.Call(page) // make sure we're not hitting the cache


    // finally navigate to url and fail out of rod.Try by panicking
    page.Timeout(websiteRequestTimeout).Navigate(url)
    page.Timeout(websiteRequestTimeout).WaitLoad() // waits for onload event
    page.Timeout(websiteRequestTimeout).WaitIdle(time.Minute)

    page.MustEval(wrapInFn(jsTTIPolyfill)) // add TTI polyfill
    page.MustEval(wrapInFn(jsWebVitalsIIFE)) // add web-vitals

    // finally actually measure the stuff
    metricsStr = page.MustEval(jsMeasurement).Str()
    
    page.MustClose()
})
// parse metricsStr
```

`jsOnNewDocument` contains javascript that gets executed on a new page before anything happens. We're subscribing to performance events which is necessary for TTI polyfill and we're clearing the local storage. This is the code ([link to source](https://github.com/dennis-tra/tiros/blob/main/js/onNewDocument.js)):

```javascript
// From https://github.com/GoogleChromeLabs/tti-polyfill#usage
!function(){if('PerformanceLongTaskTiming' in window){var g=window.__tti={e:[]};
    g.o=new PerformanceObserver(function(l){g.e=g.e.concat(l.getEntries())});
    g.o.observe({entryTypes:['longtask']})}}();

localStorage.clear();
```

Then, after the website has loaded we are adding a [TTI polyfill](https://github.com/dennis-tra/tiros/blob/main/js/tti-polyfill.js) and [web-vitals](https://github.com/dennis-tra/tiros/blob/main/js/web-vitals.iife.js) to the page.

We got the tti-polyfill from [GoogleChromeLabs/tti-polyfill](https://github.com/GoogleChromeLabs/tti-polyfill/blob/master/tti-polyfill.js) (archived in favor of the [First Input Delay](https://web.dev/fid/) metric).
We got the web-vitals javascript from [GoogleChrome/web-vitals](https://github.com/GoogleChrome/web-vitals) by building it ourselves with `npm run build` and then copying the `web-vitals.iife.js` (`iife` = immediately invoked function execution)

Then we execute the following javascript on that page ([link to source](https://github.com/dennis-tra/tiros/blob/main/js/measurement.js)):

```javascript
async () => {

    const onTTI = async (callback) => {
        const tti = await window.ttiPolyfill.getFirstConsistentlyInteractive({})

        // https://developer.chrome.com/docs/lighthouse/performance/interactive/#how-lighthouse-determines-your-tti-score
        let rating = "good";
        if (tti > 7300) {
            rating = "poor";
        } else if (tti > 3800) {
            rating = "needs-improvement";
        }

        callback({
            name: "TTI",
            value: tti,
            rating: rating,
            delta: tti,
            entries: [],
        });
    };

    const {onCLS, onFCP, onLCP, onTTFB} = window.webVitals;

    const wrapMetric = (metricFn) =>
        new Promise((resolve, reject) => {
            const timeout = setTimeout(() => resolve(null), 10000);
            metricFn(
                (metric) => {
                    clearTimeout(timeout);
                    resolve(metric);
                },
                {reportAllChanges: true}
            );
        });

    const data = await Promise.all([
        wrapMetric(onCLS),
        wrapMetric(onFCP),
        wrapMetric(onLCP),
        wrapMetric(onTTFB),
        wrapMetric(onTTI),
    ]);

    return JSON.stringify(data);
}
```

This function will return a JSON array of the following format:

```json
[
  {
    "name": "CLS",
    "value": 1.3750143983783765e-05,
    "rating": "good",
    ...
  },
  {
    "name": "FCP",
    "value": 872,
    "rating": "good",
    ...
  },
  {
    "name": "LCP",
    "value": 872,
    "rating": "good",
    ...
  },
  {
    "name": "TTFB",
    "value": 717,
    "rating": "good",
    ...
  },
  {
    "name": "TTI",
    "value": 999,
    "rating": "good",
    ...
  }
]
```

If the website request went through the `kubo` gateway we're running one round of garbage collection by calling the `/api/v0/repo/gc` endpoint. With this we make sure that the next request to that website won't come from the local kubo node cache.

To also measure a "warmed up" kubo node, we also configured a "settle time". This is just the time to wait before the first website requests are made. After the scheduler has looped through all websites we configured another settle time of 10min before all websites are requested again. Each run in between settles also has a "times" counter which is set to `5` right now in our deployment. This means that we request a single website 5 times in between each settle times. The loop looks like this:

```go
for _, settle := range c.IntSlice("settle-times") {
    time.Sleep(time.Duration(settle) * time.Second)
    for i := 0; i < c.Int("times"); i++ {
        for _, mType := range []string{models.MeasurementTypeKUBO, models.MeasurementTypeHTTP} {
            for _, website := range websites {

                pr, _ := t.Probe(c, websiteURL(c, website, mType))
                
                t.Save(c, pr, website, mType, i)

                if mType == models.MeasurementTypeKUBO {
                    t.KuboGC(c.Context)
                }
            }
        }
    }
}
```

So in total, each run measures `settle-times * times * len([http, kubo]) * len(websites)` website requests. In our case it's `2 * 5 * 2 * 14 = 280` requests. This takes around `1h` because some websites time out and the second settle time is configured to be `10m`.


## Contributing

Feel free to head over to the GitHub repository and dive in! [Open an issue](https://github.com/plprobelab/tiros/issues/new) or submit PRs.

{{< button href="https://github.com/plprobelab/tiros" >}}GitHub{{< /button >}}
