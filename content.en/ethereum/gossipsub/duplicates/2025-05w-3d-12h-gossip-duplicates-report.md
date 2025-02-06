---
title: Week 2025-05-3 12h
plotly: true
slug: 2025-05
weight: 1045600
---

# Ethereum report for message duplicates 2025-05-3 hour 12

The following results show measurement data that were collected in **calendar week** 05 **of** 
2025 ir (`2025-01-29` at hour `12`.

This report provides charts and metrics for the beacon block broadcasting latency on the Ethereum network.
The methodology we used is available [here](../methodology).

## Time distribution of duplicates

The following line plot displays the aggregated number of duplicates received over the entire hour every 1 minute.
The graph includes most of the main Ethereum CL gossipsub topics. 

{{< plotly json="/plots/2025/02/05/ethereum_mainnet-msg-duplicates-per-topic.json" height="600px" >}}

## Number of duplicates aggregated 

The next plot shows the Cumulative Distribution Function (CDF) of the number of duplicates for each received message.
Once again, the chart shows the CDF aggregated per topic.

{{< plotly json="/plots/2025/02/05/ethereum_mainnet-cdf-duplicates-per-msg-topic.json" height="600px" >}}

## Size of messages per topic

The following chart shows the CDF distribution of the message size seen by our tooling. The chart displays the distribution 
aggregated by topic. 

{{< plotly json="/plots/2025/02/05/ethereum_mainnet-cdf-size-per-msg-topic.json" height="600px" >}}

## Time between the message arrival and the first duplicate

The next chart shows the CDF of the elapsed time between the arrival of the message and the timestamp of its first duplicate.
Once again, the chart aggregates the distributions for each topic.

{{< plotly json="/plots/2025/02/05/ethereum_mainnet-cdf-time-to-first-duplicate-per-msg-topic.json" height="600px" >}}

## Correlation between message size and number of duplicates

The following chart shows the correlation between the size of a message and the number of duplicate received aggregated 
by topic. We could expect larger messages to have larger number of duplicates, thus, this graph should help understanding
if there is any clear threshold.

{{< plotly json="/plots/2025/02/05/ethereum_mainnet-correlation-size-duplicates-topic.json" height="600px" >}}