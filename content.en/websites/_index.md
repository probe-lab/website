---
title: IPFS Websites
weight: 30
bookCollapseSection: true
plotly: true
---

# Websites

## Overall Time To First Byte {#websites-overall-latency-kubo-ttfb}

{{< plotly json="../../plots/latest/websites-overall-latency-KUBO-ttfb.json" height="300px" >}}

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

## Kubo vs HTTP (90th percentile) {#websites-http-comparison-ttfb-p90}

{{< plotly json="../../plots/latest/websites-http-comparison-ttfb-p90.json" height="350px" >}}

The graph shows the 90th percentile of the Time to First Byte (TTFB) metric for
both HTTP and Kubo protocols across seven distinct regions. The horizontal axis
represents the seven regions, while the vertical axis indicates TTFB in seconds.

By displaying separate bars for HTTP and Kubo in each region, the graph allows
for a direct comparison of their TTFB performance. This comparison reveals
performance differences and provides insights into the efficiency of each
protocol in different geographic areas.

## Time To First Byte using Kubo

### (50th percentile) {#web-vitals-heatmap-KUBO-ttfb-p50}

{{< plotly json="../../plots/latest/web-vitals-heatmap-KUBO-ttfb-p50.json" height="700px" >}}

### (90th percentile) {#web-vitals-heatmap-KUBO-ttfb-p90}

{{< plotly json="../../plots/latest/web-vitals-heatmap-KUBO-ttfb-p90.json" height="700px" >}}

[//]: # ()
[//]: # (### &#40;99th percentile&#41; {#web-vitals-heatmap-KUBO-ttfb-p99})

[//]: # ()
[//]: # ({{< plotly json="../../plots/latest/web-vitals-heatmap-KUBO-ttfb-p99.json" height="700px" >}})

## First Contentful Paint using Kubo

### (50th percentile) {#web-vitals-heatmap-KUBO-fcp-p50}

{{< plotly json="../../plots/latest/web-vitals-heatmap-KUBO-fcp-p50.json" height="700px" >}}

### (90th percentile) {#web-vitals-heatmap-KUBO-fcp-p90}

{{< plotly json="../../plots/latest/web-vitals-heatmap-KUBO-fcp-p90.json" height="700px" >}}

[//]: # ()
[//]: # (### &#40;99th percentile&#41; {#web-vitals-heatmap-KUBO-fcp-p99})

[//]: # ()
[//]: # ({{< plotly json="../../plots/latest/web-vitals-heatmap-KUBO-fcp-p99.json" height="700px" >}})

## Largest Contentful Paint using Kubo

### (50th percentile) {#web-vitals-heatmap-KUBO-lcp-p50}

{{< plotly json="../../plots/latest/web-vitals-heatmap-KUBO-lcp-p50.json" height="700px" >}}

### (90th percentile) {#web-vitals-heatmap-KUBO-lcp-p90}

{{< plotly json="../../plots/latest/web-vitals-heatmap-KUBO-lcp-p90.json" height="700px" >}}

[//]: # ()
[//]: # (### &#40;99th percentile&#41; {#web-vitals-heatmap-KUBO-lcp-p99})

[//]: # ()
[//]: # ({{< plotly json="../../plots/latest/web-vitals-heatmap-KUBO-lcp-p99.json" height="700px" >}})