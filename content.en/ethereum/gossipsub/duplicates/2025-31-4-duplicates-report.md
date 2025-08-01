---
title: Week 2025-31 day 04
plotly: true
slug: 2025-31-04
weight: 1027571
---

# Ethereum report for message duplicates 2025-31 day 04

The following results show measurement data that were collected in **calendar week** 31  day 04 **of** 
2025 (`2025-07-31`).

This report provides charts for several metrics for all of GossipSub's topics on the Ethereum network.
The report focuses primarily on the broadcasting latency and number of duplicates.
We've produced these reports in collaboration with the Ethereum Foundation using out tool [Hermes](/tools/hermes/).

## Time distribution of duplicates

The following line plot displays the number of duplicates received per hour, aggregated in  intervals.
The graph includes most of the main Ethereum CL gossipsub topics. 

{{< plotly json="https://cdn.probelab.io/duplicates_reports/ethereum_mainnet/2025/7/31/duplicates_per_topic_ts_plot.json" height="600px" >}}

## Number of duplicates aggregated 

The next plot shows the Cumulative Distribution Function (CDF) of the number of duplicates for each received message.
Once again, the chart shows the CDF aggregated per topic.

{{< plotly json="https://cdn.probelab.io/duplicates_reports/ethereum_mainnet/2025/7/31/duplicates_cdf_per_msg_id_plot.json" height="600px" >}}

## Size of messages per topic

The following chart shows the CDF distribution of the message size seen by our tooling. 

{{< plotly json="https://cdn.probelab.io/duplicates_reports/ethereum_mainnet/2025/7/31/duplicates_cdf_of_msg_size_per_topic_plot.json" height="600px" >}}

## Time between the message arrival and the first duplicate

The next chart shows the CDF of the elapsed time between the arrival of a message and the arrival of its first duplicate.
Once again, the chart aggregates the distributions for each topic.

{{< plotly json="https://cdn.probelab.io/duplicates_reports/ethereum_mainnet/2025/7/31/duplicates_cdf_of_time_to_first_duplicate_plot.json" height="600px" >}}

## Correlation between message size and number of duplicates
Under the same network bandwidth, larger messages take longer to be fully transmitted between peers. Since a peer can only detect a message after downloading it completely, the likelihood of receiving duplicates increases with message size.

The heatmap below shows the correlation between message size and duplicate count, showing the number of messages per topic aggregated by "message size" and "number of duplicates" ranges.

{{< plotly json="https://cdn.probelab.io/duplicates_reports/ethereum_mainnet/2025/7/31/duplicates_correlation_between_size_and_duplicates_plot.json" height="600px" >}}