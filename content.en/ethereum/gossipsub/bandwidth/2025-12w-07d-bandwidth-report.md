---
title: Week 2025-12 day 07
plotly: true
slug: 2025-12-07
weight: 1027701
---

# Ethereum report for GossipSub Bandwidth usage 2025-12 day 07

The following results show measurement data that were collected in **calendar week** 12  day 07 **of** 
2025 (`2025-03-23`).

This report provides charts for several metrics for all of Gossipsub's topics on the Ethereum network.
The report focuses primarily on the overall bandwidth used by the gossipsub protocol when participating in the Ethereum network.
We've produced these reports in collaboration with the Ethereum Foundation using out tool [Hermes](/tools/hermes/).

## Incoming vs. Outgoing Bandwidth Usage
The following time series graph displays the **incoming (`in`)** and **outgoing (`out`)** bandwidth usage of the GossipSub protocol throughout the day. The data is aggregated into time windows of **5 minutes**.

{{< plotly json="https://cdn.probelab.io/bandwidth_reports/ethereum_mainnet/2025/3/23/bandwidth_direction_aggregation_ts_plot.json" height="600px" >}}

## Bandwidth Usage by Message Type
Extending the previous visualization, this time series graph further breaks down the **incoming** and **outgoing** bandwidth usage by message type, providing deeper insight into the distribution of network traffic.

{{< plotly json="https://cdn.probelab.io/bandwidth_reports/ethereum_mainnet/2025/3/23/bandwidth_per_msg_type_ts_plot.json" height="600px" >}}

## Bandwidth Distribution by Message Type
The final visualization illustrates the percentage of total bandwidth attributed to each message type. This helps identify which types of messages contribute the most to network traffic.

{{< plotly json="https://cdn.probelab.io/bandwidth_reports/ethereum_mainnet/2025/3/23/bandwidth_total_per_type_plot.json" height="600px" >}}
