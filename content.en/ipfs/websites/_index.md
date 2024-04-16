---
title: Website Monitoring
bookCollapseSection: true
plotly: true
weight: 40
aliases:
    - /websites/
---

# Websites

## Overall Time To First Byte {#websites-ttfb-latencies}

{{< plotly json="../../plots/latest/websites-ttfb-latencies.json" height="300px" >}}

The graph presents the Time to First Byte (TTFB) metric, which measures the
duration between the initiation of a request and the receipt of the first byte.
To construct this graph, TTFB percentiles were calculated on a daily basis
incorporating data from multiple regions. To smooth out the daily fluctuations
and identify overarching trends, a three-day rolling average was employed. This process
involved calculating the average TTFB values for each day, taking into account
the percentile values from the preceding two days as well. By incorporating this rolling
average, the graph provides a more stable representation of TTFB trends, aiding in the
identification of long-term patterns.

It is important to note that the resulting metric is an artificial composite
and does not reflect the specific performance of any individual region. Rather,
it allows to discern general tendencies and fluctuations in TTFB across the
combined dataset. By focusing on these overall trends, we can gain valuable
insights into the performance of Kubo TTFB performance as a whole, making it easier to
identify any notable deviations or improvements in that metric over time.

## Helia vs Kubo vs HTTP (90th percentile) {#websites-http-comparison-ttfb-p90}

{{< plotly json="../../plots/latest/websites-http-comparison-ttfb-p90.json" height="350px" >}}

The graph shows the 90th percentile of the Time to First Byte (TTFB) metric for
retrievals via HTTP, through Helia, and through Kubo across seven distinct
regions. The horizontal axis represents the seven regions, while the vertical
axis indicates TTFB in seconds.

By displaying separate bars for HTTP, Helia, and Kubo in each region, the graph allows
for a direct comparison of their TTFB performance. This comparison reveals
performance differences and provides insights into the efficiency of each
protocol/implementation in different geographic areas.

For Helia in the above graph, the different variations work out as follows:


`Helia (TG)` - Helia in NodeJS that is delegating [content-routing](https://docs.ipfs.tech/concepts/how-ipfs-works/#how-content-routing-works-in-ipfs), [peer-routing](https://docs.ipfs.tech/concepts/glossary/#peer-routing), and block retrieval entirely to [trustless gateways](https://specs.ipfs.tech/http-gateways/trustless-gateway/). js-libp2p is not used. This is the simplest case scenario for the Helia node since an external server is doing all of the heavy lifting.

`Helia (DR)` - Helia in NodeJS that [delegates content- and peer-routing to an external server using HTTP](https://docs.ipfs.tech/concepts/how-ipfs-works/#how-content-routing-works-in-ipfs) but then uses js-libp2p for direct peer retrieval. HTTP Trustless Gateways are not used.

`Helia` - Effectively the combination of "Helia (TG)" + "Helia (DR)".  Helia in NodeJS in parallel is:
 - delegating content- and peer-routing, as well as block retrieval entirely to [trustless gateways](https://specs.ipfs.tech/http-gateways/trustless-gateway/) AND
 - delegating content- and peer-routing to an external server using HTTP, but then uses js-libp2p for direct block retrieval from other peers. 
It's expected that this would yield more reliable results than "Helia (TG)" or "Helia (DR)" alone.

Important scenarios that aren't measured yet:
1. Helia in NodeJS where content, peer, and block retrieval are done entirely in a p2p fashion.  
2. The scenarios above but instead from within Chromium rather than NodeJS.

## Time To First Byte using Kubo

### (50th percentile) {#websites-web-vitals-heatmap-IPFS-KUBO-ttfb-p50}

{{< plotly json="../../plots/latest/websites-web-vitals-heatmap-IPFS-KUBO-ttfb-p50.json" height="700px" >}}

### (90th percentile) {#websites-web-vitals-heatmap-IPFS-KUBO-ttfb-p90}

{{< plotly json="../../plots/latest/websites-web-vitals-heatmap-IPFS-KUBO-ttfb-p90.json" height="700px" >}}

## First Contentful Paint using Kubo

### (50th percentile) {#websites-web-vitals-heatmap-IPFS-KUBO-fcp-p50}

{{< plotly json="../../plots/latest/websites-web-vitals-heatmap-IPFS-KUBO-fcp-p50.json" height="700px" >}}

### (90th percentile) {#websites-web-vitals-heatmap-IPFS-KUBO-fcp-p90}

{{< plotly json="../../plots/latest/websites-web-vitals-heatmap-IPFS-KUBO-fcp-p90.json" height="700px" >}}

## Largest Contentful Paint using Kubo

### (50th percentile) {#websites-web-vitals-heatmap-IPFS-KUBO-lcp-p50}

{{< plotly json="../../plots/latest/websites-web-vitals-heatmap-IPFS-KUBO-lcp-p50.json" height="700px" >}}

### (90th percentile) {#websites-web-vitals-heatmap-IPFS-KUBO-lcp-p90}

{{< plotly json="../../plots/latest/websites-web-vitals-heatmap-IPFS-KUBO-lcp-p90.json" height="700px" >}}
