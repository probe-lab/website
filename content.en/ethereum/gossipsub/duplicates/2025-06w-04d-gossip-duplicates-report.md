---
title: Week 2025-06 day 04
plotly: true
slug: 2025-06-04
weight: 1045599
---

# Ethereum report for message duplicates 2025-06 day 04

The following results show measurement data that were collected in **calendar week** 06  day 04 **of** 
2025 in `2025-02-13`.

This report provides charts and metrics for the beacon block broadcasting latency on the Ethereum network.
We've produced these reports in collaboration with the Ethereum Foundation using out tool [Hermes](/tools/hermes/).

## Time distribution of duplicates

The following line plot displays the aggregated number of duplicates received over the entire hour every 5 minutes.
The graph includes most of the main Ethereum CL gossipsub topics. 

{{< plotly json="/plots/2025/02/06/04/ethereum_mainnet_ef-msg-duplicates-per-topic.json" height="600px" >}}

## Number of duplicates aggregated 

The next plot shows the Cumulative Distribution Function (CDF) of the number of duplicates for each received message.
Once again, the chart shows the CDF aggregated per topic.

{{< plotly json="/plots/2025/02/06/04/ethereum_mainnet_ef-cdf-duplicates-per-msg-topic.json" height="600px" >}}

## Size of messages per topic

The following chart shows the CDF distribution of the message size seen by our tooling. 

{{< plotly json="/plots/2025/02/06/04/ethereum_mainnet_ef-cdf-size-per-msg-topic.json" height="600px" >}}

## Time between the message arrival and the first duplicate

The next chart shows the CDF of the elapsed time between the arrival of a message and the arrival of its first duplicate.
Once again, the chart aggregates the distributions for each topic.

{{< plotly json="/plots/2025/02/06/04/ethereum_mainnet_ef-cdf-time-to-first-duplicate-per-msg-topic.json" height="600px" >}}

## Correlation between message size and number of duplicates

The following heatmap chart shows the correlation between the size of a message and the number of duplicate received.
The chart shows how many messages per topic were aggregated by defined message size and number of duplicate ranges.
We could expect larger messages to have larger number of duplicates, thus, this graph should help understanding
if there is any clear threshold.

{{< plotly json="/plots/2025/02/06/04/ethereum_mainnet_ef-correlation-size-duplicates-topic.json" height="600px" >}}