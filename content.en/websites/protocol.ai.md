---
title: protocol.ai
plotly: true
---

# protocol.ai

This page shows performance metrics for the
website [protocol.ai](https://protocol.ai).

## Web-Vitals Metrics (Kubo) {#web-vitals-barchart-protocolai-kubo-eu-central-1}

{{< plotly json="../../plots/latest/web-vitals-barchart-protocol.ai-KUBO-eu-central-1.json" height="400px" >}}





## Metrics

| Metric   | Description                                                                                 |    Good | Needs Improvement |    Poor |
|----------|---------------------------------------------------------------------------------------------|--------:|------------------:|--------:|
| **CLS**  | [Cumulative Layout Shift](https://web.dev/cls/)                                             | < 0.10s |           < 0.25s | > 0.25s |
| **FCP**  | [First Contentful Paint](https://web.dev/fcp/)                                              | < 1.80s |           < 3.00s | > 3.00s |
| **LCP**  | [Largest Contentful Paint](https://web.dev/lcp/)                                            | < 2.50s |           < 4.00s | > 4.00s |
| **TTFB** | [Time To First Byte](https://web.dev/ttfb/)                                                 | < 0.80s |           < 1.80s | > 1.80s |
| **TTI**  | [Time To Interactive](https://developer.chrome.com/docs/lighthouse/performance/interactive) | < 3.80s |           < 7.30s | > 7.30s |

## Values

| Value                 | Description                                                                                          |
|-----------------------|------------------------------------------------------------------------------------------------------|
| **Good**              | The value was smaller than the `good` threshold (see [Metrics](#metrics))                            |
| **Needs Improvement** | The value was larger than the `good` but smaller than the `poor` threshold (see [Metrics](#metrics)) |
| **Poor**              | The value was larger than the `poor` threshold (see [Metrics](#metrics))                             |
| **Undefined**         | We could not gather the metric because of internal measurement errors (our fault)                    |
| **Fatal**             | We could not gather the metric because we could not retrieve the website at all                      |
