---
title: protocol.ai
plotly: true
---

# protocol.ai

This page shows performance metrics for the website [protocol.ai](https://protocol.ai).

## Overall Performance {#website-performance-gauge-protocolai-kubo}

{{< plotly json="../../plots/latest/website-performance-gauge-protocol.ai-KUBO.json" height="300px" >}}

The graph presents a comparison of two crucial web performance metrics: Time to First Byte (TTFB) and First Contentful Paint (FCP).
The data displayed shows the 90th percentile of both metrics and was gathered during the previous week.
To aid in comparison, the percentage difference between the previous week and the week before is displayed below each gauge indicator. This enables easy identification of performance improvements or regressions over time.

By default, the graph focuses on the `eu-central-1` region, which can be customized as needed in the lower left corner.
The graph utilizes shaded areas in different colors to denote performance categories based on predefined thresholds set by web-vitals. Green represents "good" performance, yellow indicates "needs improvement," and red signifies "poor" performance (see [Metrics](#metrics)).

## Progression

### Unique Website Providers per Day {#website-providers-protocolai}

{{< plotly json="../../plots/latest/website-providers-protocol.ai.json" height="350px" >}}

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



### Time To First Byte Rating {#website-providers-protocolai}

{{< plotly json="../../plots/latest/website-metric-good-rating-protocol.ai-KUBO-ttfb.json" height="300px" >}}

## Snapshot

### Website Probes

{{< plotly json="../../plots/latest/web-vitals-probes-count-protocol.ai.json" height="150px" >}}

### Web-Vitals Metrics measured from Europe using Kubo {#web-vitals-barchart-protocolai-kubo-eu-central-1}

<small>[What do `CLS`, `FCP`, `LCP`, etc. mean?](#metrics) | [What do `Fatal`, `Undefined`, `Poor` etc. mean?](#values)</small>

{{< plotly json="../../plots/latest/web-vitals-barchart-protocol.ai-KUBO-eu-central-1.json" height="400px" >}}

### Website Probing Success rate from different Regions {#website-retrieval-errors-protocolai}

{{< plotly json="../../plots/latest/website-retrieval-errors-protocol.ai.json" height="350px" >}}

### Kubo Metrics by Region

#### Time To First Byte {#website-metric-cdf-protocolai-kubo-ttfb}

{{< plotly json="../../plots/latest/website-metric-cdf-protocol.ai-KUBO-ttfb.json" height="320px" >}}

#### First Contentful Paint {#website-metric-cdf-protocolai-kubo-fcp}

{{< plotly json="../../plots/latest/website-metric-cdf-protocol.ai-KUBO-fcp.json" height="320px" >}}

#### Largest Contentful Paint {#website-metric-cdf-protocolai-kubo-lcp}

{{< plotly json="../../plots/latest/website-metric-cdf-protocol.ai-KUBO-lcp.json" height="320px" >}}

### Time To First Byte Kubo/HTTP Latency Ratio {#website-http-ratio-protocolai}

{{< plotly json="../../plots/latest/website-http-ratio-protocol.ai.json" height="500px" >}}

We calculated different percentiles for the Time To First Byte metric in different regions for website requests that were done via Kubo and via plain HTTP.
Then we divided the values of Kubo by the ones from HTTP. A resulting number greater than `1` means that Kubo was slower than HTTP in that region for that percentile.
Conversely, a number less than `1` means that Kubo was faster.

## Glossary

### Metrics

| Metric   | Description                                                                                 |    Good | Needs Improvement |     Poor |
|----------|---------------------------------------------------------------------------------------------|--------:|------------------:|---------:|
| **CLS**  | [Cumulative Layout Shift](https://web.dev/cls/)                                             | < 0.10s |           < 0.25s | >= 0.25s |
| **FCP**  | [First Contentful Paint](https://web.dev/fcp/)                                              | < 1.80s |           < 3.00s | >= 3.00s |
| **LCP**  | [Largest Contentful Paint](https://web.dev/lcp/)                                            | < 2.50s |           < 4.00s | >= 4.00s |
| **TTFB** | [Time To First Byte](https://web.dev/ttfb/)                                                 | < 0.80s |           < 1.80s | >= 1.80s |
| **TTI**  | [Time To Interactive](https://developer.chrome.com/docs/lighthouse/performance/interactive) | < 3.80s |           < 7.30s | >= 7.30s |

### Values

| Value                 | Description                                                                                          |
|-----------------------|------------------------------------------------------------------------------------------------------|
| **Good**              | The value was smaller than the `good` threshold (see [Metrics](#metrics))                            |
| **Needs Improvement** | The value was larger than the `good` but smaller than the `poor` threshold (see [Metrics](#metrics)) |
| **Poor**              | The value was larger than the `poor` threshold (see [Metrics](#metrics))                             |
| **Undefined**         | We could not gather the metric because of internal measurement errors (our fault)                    |
| **Fatal**             | We could not gather the metric because we could not retrieve the website at all                      |
