---
title: docs.libp2p.io
plotly: true
---

# docs.libp2p.io

This page shows performance metrics for the website [docs.libp2p.io](https://docs.libp2p.io).

## Overall Performance {#website-performance-gauge-docslibp2pio-kubo}

{{< plotly json="../../plots/latest/website-performance-gauge-docs.libp2p.io-KUBO.json" height="300px" >}}

The graph presents a comparison of two crucial web performance metrics: Time to First Byte (TTFB) and First Contentful Paint (FCP).
The data displayed shows the 90th percentile of both metrics and was gathered during the previous week.
To aid in comparison, the percentage difference between the previous week and the week before is displayed below each gauge indicator. This enables easy identification of performance improvements or regressions over time.

By default, the graph focuses on the `eu-central-1` region, which can be customized as needed in the lower left corner.
The graph utilizes shaded areas in different colors to denote performance categories based on predefined thresholds set by web-vitals. Green represents "good" performance, yellow indicates "needs improvement," and red signifies "poor" performance (see [Metrics](#metrics)).

## Progression

### Web-Vitals Metrics (90th Percentile) {#website-metrics-docslibp2pio}

{{< plotly json="../../plots/latest/website-metrics-docs.libp2p.io.json" height="300px" >}}

The [Web-Vitals](https://web.dev/vitals/) Metrics graph provides high-level insights into the websites' performance, allowing you to monitor key metrics over the past 30 days.

To construct this graph, the 90th percentiles for each metric were calculated on a daily basis incorporating data from multiple regions.
To smooth out the daily fluctuations and identify overarching trends, a three-day rolling average was employed.
This process involved calculating the average percentile values for each day, taking into account the percentile values from the preceding two days as well.
By incorporating this rolling average, the graph provides a more stable representation of metric trends, aiding in the identification of long-term patterns.

The graph showcases five essential web vitals metrics, enabling you to assess the performance of your web pages thoroughly. These metrics include:

1. **TTFB** - _Time to First Byte_: This metric measures the time it takes for a user's browser to receive the first byte of data from a web server. It indicates the responsiveness of the server and the network.
2. **FCP** - _First Contentful Paint_: FCP measures the time it takes for the first piece of content to appear on a user's screen, such as text, images, or graphics. It provides insights into the perceived loading speed of your web page.
3. **LCP** - _Largest Contentful Paint_: LCP measures the time it takes for the largest content element, such as an image or a block of text, to become visible to the user. It reflects the time until the main content of the page is rendered.
4. **CLS** - _Cumulative Layout Shift_: CLS evaluates the visual stability of a web page by measuring the unexpected layout shifts that occur during the loading process. It helps determine whether content shifts on the screen, which can lead to a poor user experience.
5. **TTI** - _Time to Interactive_: TTI measures the time it takes for a web page to become fully interactive and responsive to user input. It indicates when users can effectively engage with the page, interact with elements, or start using web applications.

It is important to note that the resulting metrics are artificial composites and do not reflect the specific performance of any individual region. Rather, it allows discerning general tendencies and fluctuations in these metrics across the combined dataset.

### Unique Website Providers per Day {#website-providers-docslibp2pio}

{{< plotly json="../../plots/latest/website-providers-docs.libp2p.io.json" height="350px" >}}

This graph showcases the unique providing peers as identified by distinct [PeerIDs](https://docs.libp2p.io/concepts/fundamentals/peers/#peer-id) discovered throughout a specific day with our [website measurement infrastructure](/tools/tiros).
Each peer was attempted to be connected to, and based on the outcome, it was categorized accordingly.

If a peer remained unreachable during the entire day, it was labeled as "Unreachable".
On the other hand, if we were able to connect to a peer at least once during the day,
it was marked as either "Reachable Relayed" or "Reachable Non-Relayed."

"Reachable Relayed" indicates that the peer solely advertised relay multiaddresses, meaning we could only establish a connection to it through a proxying peer. These peers act as intermediaries, facilitating the connection between our system and the reachable peer.
It is important to note that this scenario has performance implications, as relaying introduces additional latency and potential bottlenecks.

"Reachable Non-Relayed" represents peers that were publicly reachable and usually offer the best performance.

In addition to the peer-related information, the graph also includes black markers that represent the number of website deployments per day.
Deployments are determined by monitoring the CIDs found within the websites' IPNS records. If the CID changes, we consider this a new deployment.

### Time To First Byte Rating {#website-providers-docslibp2pio}

{{< plotly json="../../plots/latest/website-metric-good-rating-docs.libp2p.io-KUBO-ttfb.json" height="300px" >}}

## Snapshot

### Website Probes

{{< plotly json="../../plots/latest/web-vitals-probes-count-docs.libp2p.io.json" height="150px" >}}

### Web-Vitals Metrics measured from Europe using Kubo {#web-vitals-barchart-docslibp2pio-kubo-eu-central-1}

<small>[What do `CLS`, `FCP`, `LCP`, etc. mean?](#metrics) | [What do `Fatal`, `Undefined`, `Poor` etc. mean?](#values)</small>

{{< plotly json="../../plots/latest/web-vitals-barchart-docs.libp2p.io-KUBO-eu-central-1.json" height="400px" >}}

### Website Probing Success rate from different Regions {#website-retrieval-errors-docslibp2pio}

{{< plotly json="../../plots/latest/website-retrieval-errors-docs.libp2p.io.json" height="350px" >}}

### Kubo Metrics by Region

#### Time To First Byte {#website-metric-cdf-docslibp2pio-kubo-ttfb}

{{< plotly json="../../plots/latest/website-metric-cdf-docs.libp2p.io-KUBO-ttfb.json" height="320px" >}}

#### First Contentful Paint {#website-metric-cdf-docslibp2pio-kubo-fcp}

{{< plotly json="../../plots/latest/website-metric-cdf-docs.libp2p.io-KUBO-fcp.json" height="320px" >}}

#### Largest Contentful Paint {#website-metric-cdf-docslibp2pio-kubo-lcp}

{{< plotly json="../../plots/latest/website-metric-cdf-docs.libp2p.io-KUBO-lcp.json" height="320px" >}}

### Time To First Byte Kubo/HTTP Latency Ratio {#website-http-ratio-docslibp2pio}

{{< plotly json="../../plots/latest/website-http-ratio-docs.libp2p.io.json" height="500px" >}}

We calculated different percentiles for the Time To First Byte metric in different regions for website requests that were done via Kubo and via plain HTTP.
Then we divided the values of Kubo by the ones from HTTP. A resulting number greater than `1` means that Kubo was slower than HTTP in that region for that percentile.
Conversely, a number less than `1` means that Kubo was faster.

## Glossary

### Metrics

#### Time to First Byte (TTFB)

This metric measures the time it takes for a user's browser to receive the first byte of data for a website. It indicates the responsiveness of the server and/or the network.

#### First Contentful Paint (FCP)

FCP measures the time it takes for the first piece of content to appear on a user's screen, such as text, images, or graphics. It provides insights into the perceived loading speed of your web page.

#### Largest Contentful Paint (LCP)

LCP measures the time it takes for the largest content element, such as an image or a block of text, to become visible to the user. It reflects the time until the main content of the page is rendered.

#### Cumulative Layout Shift (CLS)

CLS evaluates the visual stability of a web page by measuring the unexpected layout shifts that occur during the loading process. It helps determine whether content shifts on the screen, which can lead to a poor user experience.

#### Time to Interactive (TTI)

TTI measures the time it takes for a web page to become fully interactive and responsive to user input. It indicates when users can effectively engage with the page, interact with elements, or start using web applications.

#### Thresholds

| Metric   | Description                                                                                 |    Good | Needs Improvement |     Poor |
|----------|---------------------------------------------------------------------------------------------|--------:|------------------:|---------:|
| **TTFB** | [Time To First Byte](https://web.dev/ttfb/)                                                 | < 0.80s |           < 1.80s | >= 1.80s |
| **FCP**  | [First Contentful Paint](https://web.dev/fcp/)                                              | < 1.80s |           < 3.00s | >= 3.00s |
| **LCP**  | [Largest Contentful Paint](https://web.dev/lcp/)                                            | < 2.50s |           < 4.00s | >= 4.00s |
| **CLS**  | [Cumulative Layout Shift](https://web.dev/cls/)                                             | < 0.10s |           < 0.25s | >= 0.25s |
| **TTI**  | [Time To Interactive](https://developer.chrome.com/docs/lighthouse/performance/interactive) | < 3.80s |           < 7.30s | >= 7.30s |


### Values

| Value                 | Description                                                                                          |
|-----------------------|------------------------------------------------------------------------------------------------------|
| **Good**              | The value was smaller than the `good` threshold (see [Metrics](#metrics))                            |
| **Needs Improvement** | The value was larger than the `good` but smaller than the `poor` threshold (see [Metrics](#metrics)) |
| **Poor**              | The value was larger than the `poor` threshold (see [Metrics](#metrics))                             |
| **Undefined**         | We could not gather the metric because of internal measurement errors (our fault)                    |
| **Fatal**             | We could not gather the metric because we could not retrieve the website at all                      |
