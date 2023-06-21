---
title: docs.ipfs.tech
plotly: true
---

# `docs.ipfs.tech`

This page shows performance metrics for the website [docs.ipfs.tech](https://docs.ipfs.tech). We gathered the data with our [Tiros](/tools/tiros) measurement tool.

We initially present an Overview of the performance, followed by Trends, i.e., the performance over time, for the last 30 days. We finally take a snapshot of the performance seen during the most recent week and present performance details over this period. The combination of long-term (Trends) and short-term (Snapshot) results provide a comprehensive view of the site's performance, as well as potential areas of improvement.

## Overview

### Performance over Kubo {#website-snapshot-performance-gauge-docsipfstech-kubo}

{{< plotly json="../../plots/latest/website-snapshot-performance-gauge-docs.ipfs.tech-KUBO.json" height="300px" >}}

The graph presents a comparison of two crucial Kubo web performance metrics (90th percentile): Time to First Byte (TTFB) and First Contentful Paint (FCP).
The data displayed shows the 90th percentile of both metrics and was gathered during the previous week.
To aid in comparison, the percentage difference between the previous week and the week before is displayed below each gauge indicator. This enables easy identification of performance improvements or regressions over time. The plots on this page are updated weekly on Monday.

By default, the graph focuses on the `eu-central-1` region, which can be customized as needed in the lower left corner through the drop-down menu.
The graph utilizes shaded areas in different colors to denote performance categories based on predefined thresholds set by web-vitals. Green represents "good" performance, yellow indicates "needs improvement," and red signifies "poor" performance (see [Metrics](#metrics)).

## Trends

The website measurement data can be viewed from two distinct perspectives: as a snapshot of the most recent data points (see [Snapshot](#snapshot)) and as a progression over time. This section provides insights into the overall trends by presenting metrics over time. Analyzing data solely as a snapshot can offer a momentary glimpse into the current state, but it may lack context and fail to capture the bigger picture. However, by examining data over time, patterns and trends emerge, allowing for a deeper understanding of the data's trajectory and potential future outcomes. This section's focus on metrics over time enhances the ability to identify and interpret these trends, enabling informed decision-making and strategic planning.

### Web-Vitals Metrics (90th Percentile) {#website-trend-metrics-docsipfstech}

{{< plotly json="../../plots/latest/website-trend-metrics-docs.ipfs.tech.json" height="300px" >}}

The [Web-Vitals](https://web.dev/vitals/) Metrics graph provides high-level insights into the websites' performance, allowing you to monitor key metrics over the past 30 days.

To construct this graph, the 90th percentiles for each metric were calculated on a daily basis incorporating data from multiple regions.
To smooth out the daily fluctuations and identify overarching trends, a three-day rolling average was employed.
This process involved calculating the average percentile values for each day, taking into account the percentile values from the preceding two days as well.
By incorporating this rolling average, the graph provides a more stable representation of metric trends, aiding in the identification of long-term patterns.

The graph showcases five essential web vitals metrics, enabling you to assess the performance of your web pages thoroughly. These metrics include Time to First Byte (TTFB), First Contentful Paint (FCP), Largest Contentful Paint (LCP), Cumulative Layout Shift (CLS), and Time to Interactive (TTI) (see [Metrics](#metrics) for more information).

It is important to note that the resulting metrics are artificial composites and do not reflect the specific performance of any individual region. Rather, it allows discerning general tendencies and fluctuations in these metrics across the combined dataset.

### Unique IPFS Website Providers per Day {#website-trend-providers-docsipfstech}

{{< plotly json="../../plots/latest/website-trend-providers-docs.ipfs.tech.json" height="350px" >}}

One of the primary reasons why a website (CID) might not be available over Kubo is that there are no providers for the content. This graph showcases the unique providing peers as identified by distinct [PeerIDs](https://docs.libp2p.io/concepts/fundamentals/peers/#peer-id) discovered throughout a specific day in the IPFS DHT. We look up the IPNS record for the website which gives us the current CID of the website's content. Then, we look up all provider records in the IPFS DHT and record the distinct PeerIDs of the providing peers. Finally, we try to connect to all the discovered peers, and based on the outcome, classify them as:

- `Unreachable`: The peer remained unreachable during the entire day. We weren't able to connect to it a single time, which might happen if providers advertise a [provider record](https://docs.ipfs.tech/concepts/dht/#distributed-hash-tables-dhts) and then leave the network (go offline), leaving a stale record behind.
- `Reachable Relayed`: The providing peer solely advertised [relay multiaddresses](https://docs.libp2p.io/concepts/nat/circuit-relay/), meaning we could only establish a connection to it through a proxying peer. Proxying peers act as intermediaries, facilitating the connection between our monitoring system and the providing peer. It is important to note that this scenario has performance implications, as relaying introduces additional latency and potential bottlenecks _at the connection setup level_. Once the connection is set up, the actual data transfer takes place directly between the two peers, so there is no performance implication during that stage. It is also worth noting that where a "Reachable Relayed" peer is counted, the original providing peer has been reached and it was not only the proxying peer that our probes reached.
- `Reachable Non-Relayed` represents peers that were publicly reachable and usually offer the best performance.

In order for a website (or CID more in general) to be available and accessible in the network, there needs to be at least one “Reachable Provider”. Obviously, the more the “Reachable Providers”, the “healthier” the content will be. This is because the content is more resilient to peer churn (i.e., some of the providers leaving the network) and is also available from more locations, which likely means faster retrieval.

In addition to the peer-related information, the graph also includes black markers that represent the number of website deployments per day (count shown on the right handside y-axis).
Deployments are determined by monitoring the CIDs found within the websites' IPNS records. If the CID changes, we consider this a new deployment.

#### Known Stable Providers {#website-trend-hosters-docsipfstech}

{{< plotly json="../../plots/latest/website-trend-hosters-docs.ipfs.tech.json" height="250px" >}}

For the above graph, we obtained the PeerIDs from two hosting providers/pinning services: [Protocol Labs' IPFS Websites Collab-Cluster](https://collab.ipfscluster.io/) and [Fleek](https://fleek.co). We monitor how many of their PeerIDs appear in the list of peers providing the website to the DHT on a daily basis. We gather the list of providing peers every six hours from seven different vantage points with two different Kubo versions (see [Tiros](/tools/tiros)), and aggregate the distinct peers we have found. Then we count the number of peerIDs that belong to either Fleek or PL's Collab-Cluster. We monitor six PeerIDs for Fleek and seven PeerIDs for PL's Collab-Cluster ([source](https://github.com/plprobelab/website/blob/main/config/plotdefs-website/website-trend-hosters.yaml#L8)). The sum of both bars should always be less than or equal to the number of `Reachable Non-Relayed` peers in the previous graph.

More importantly, if both bars (for Fleek and the PL Cluster) are at zero, then this very likely means that the website has no stable providers on the IPFS DHT.

### IPFS Retrieval Errors {#website-trend-retrieval-errors-docsipfstech-kubo}

{{< plotly json="../../plots/latest/website-trend-retrieval-errors-docs.ipfs.tech-KUBO.json" height="350px" >}}

This graph shows error rates of website requests via Kubo over the past 30 days. It combines measurements from all our measurement regions. The x-axis represents days, while the y-axis displays error rates. Additionally, black markers indicate the actual number of requests (from our infrastructure) per day with the corresponding count shown on the right handside y-axis. This graph offers a concise overview of error rates and probing volume, aiding users in assessing website availability in IPFS.

## Snapshot

This section presents a snapshot of the most recent week's data, offering a concise overview of the current state. By focusing on this specific timeframe, readers gain immediate insights into the prevailing metrics and performance indicators. While a single snapshot may lack the context of historical data, it serves as a valuable tool for assessing the present situation. Analyzing the data in this way allows for quick identification of key trends, patterns, and potential areas of concern or success. This section's emphasis on the snapshot of data enables decision-makers to make informed, real-time assessments and take immediate actions based on the current status.

### Website Probes {#website-snapshot-probes-count-docsipfstech}

{{< plotly json="../../plots/latest/website-snapshot-probes-count-docs.ipfs.tech.json" height="150px" >}}

We perform on average 500 requests per week from each of the seven AWS regions
where our infrastructure is deployed using [Kubo](https://github.com/ipfs/kubo)
and HTTP. Above is the number of requests for each of the request methods. The number may
vary depending on errors during the fetching process, which we look into more
detail further down.

### Web-Vitals Metrics measured from Europe using Kubo {#website-snapshot-web-vitals-barchart-docsipfstech-kubo-eu-central-1}

{{< plotly json="../../plots/latest/website-snapshot-web-vitals-barchart-docs.ipfs.tech-KUBO-eu-central-1.json" height="400px" >}}

<small>[What do `Fatal`, `Undefined`, `Poor` etc. mean?](#values)</small>

During the designated time period (indicated in the bottom right corner), we conducted multiple measurements for the five metrics shown along the x-axis. The y-axis represents the proportion of measurement outcomes from the total number of measurements specifically taken in the eu-central-1 region. This visual representation allows us to analyze the distribution of the metric ratings within the specified time frame for that particular region.

### Website Probing Success rate from different Regions {#website-snapshot-retrieval-errors-docsipfstech}

{{< plotly json="../../plots/latest/website-snapshot-retrieval-errors-docs.ipfs.tech.json" height="350px" >}}

While the graph on [IPFS Retrieval Errors](#website-trend-retrieval-errors-docsipfstech-kubo) further up shows the Kubo retrieval errors over time, this graph shows the retrieval errors as seen from different regions in the specified time interval (bottom right corner). Alongside the Kubo retrieval outcomes it also shows the HTTP results. The black markers again show the number of probes performed in each region with each request method (count shown on the right handside y-axis).

### Kubo Metrics by Region

This series of graphs presents a comprehensive analysis of latency performance across different regions using [Cumulative Distribution Functions (CDFs)](https://en.wikipedia.org/wiki/Cumulative_distribution_function). The primary focus is on three crucial metrics: Time to First Byte (TTFB), First Contentful Paint (FCP), and Largest Contentful Paint (LCP). Each graph in the series shows the CDFs of a specific metric from all our measured regions. CDFs allow for a holistic view of the latency distribution, showcasing how often specific latency values occur within each region.

To provide context and aid interpretation, the graphs incorporate shaded background areas. These areas are color-coded, with green representing good performance, yellow indicating areas that require improvement, and red denoting poor performance. The thresholds are defined by [web-vitals](https://web.dev/vitals) (more info below in [Metrics](#metrics)). By analyzing the position of the CDF lines in relation to the shaded regions, one can quickly identify regions with superior, average, or subpar latency performance for TTFB, FCP, and LCP.

#### Time To First Byte {#website-snapshot-metric-cdf-docsipfstech-kubo-ttfb}

{{< plotly json="../../plots/latest/website-snapshot-metric-cdf-docs.ipfs.tech-KUBO-ttfb.json" height="320px" >}}

#### First Contentful Paint {#website-snapshot-metric-cdf-docsipfstech-kubo-fcp}

{{< plotly json="../../plots/latest/website-snapshot-metric-cdf-docs.ipfs.tech-KUBO-fcp.json" height="320px" >}}

#### Largest Contentful Paint {#website-snapshot-metric-cdf-docsipfstech-kubo-lcp}

{{< plotly json="../../plots/latest/website-snapshot-metric-cdf-docs.ipfs.tech-KUBO-lcp.json" height="320px" >}}

### Kubo vs HTTP Latency Comparison (TTFB) {#website-snapshot-http-ratio-docsipfstech-ttfb}

{{< plotly json="../../plots/latest/website-snapshot-http-ratio-docs.ipfs.tech-ttfb.json" height="500px" >}}

We calculated different percentiles for the Time To First Byte (TTFB) metric in different regions for website requests that were done via Kubo and via plain HTTP. Then we divided the values of Kubo by the ones from HTTP. A resulting number greater than `1` means that Kubo was slower than HTTP in that region for that percentile. Conversely, a number less than `1` means that Kubo was faster.

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
| **CLS**  | [Cumulative Layout Shift](https://web.dev/cls/)                                             | < 0.10  |           < 0.25  | >= 0.25  |
| **TTI**  | [Time To Interactive](https://developer.chrome.com/docs/lighthouse/performance/interactive) | < 3.80s |           < 7.30s | >= 7.30s |


### Values

| Value                 | Description                                                                                          |
|-----------------------|------------------------------------------------------------------------------------------------------|
| **Good**              | The value was smaller than the `good` threshold (see [Metrics](#metrics))                            |
| **Needs Improvement** | The value was larger than the `good` but smaller than the `poor` threshold (see [Metrics](#metrics)) |
| **Poor**              | The value was larger than the `poor` threshold (see [Metrics](#metrics))                             |
| **Undefined**         | We could not gather the metric because of internal measurement errors (our fault)                    |
| **Fatal**             | We could not gather the metric because we could not retrieve the website at all                      |