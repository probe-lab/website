---
title: Week 2025-25 day 04
plotly: true
slug: 2025-25-04
weight: 1045580
---

# Filecoin report for GossipSub Bandwidth usage 2025-25 day 04

The following results show measurement data that were collected in **calendar week** 25  day 04 **of** 
2025 (`2025-06-19`).

This report provides charts for several metrics for all of GossipSub's topics on the Filecoin network.
The report focuses primarily on the overall bandwidth used by the gossipsub protocol when participating in the Filecoin network.
We've produced these reports in collaboration with the Filecoin Foundation using our tool [Hermes](/tools/hermes/).

## Incoming vs. Outgoing Bandwidth Usage
The following time series graph displays the **incoming (`in`)** and **outgoing (`out`)** bandwidth usage of the GossipSub protocol throughout the day. The data is aggregated into time windows of ****.

{{< plotly json="https://cdn.probelab.io/bandwidth_reports/filecoin_mainnet/2025/6/19/bandwidth_direction_aggregation_ts_plot.json" height="600px" >}}

## Bandwidth Usage by Message Type
Extending the previous visualization, this time series graph further breaks down the **incoming** and **outgoing** bandwidth usage by message type, providing deeper insight into the distribution of network traffic.

{{< plotly json="https://cdn.probelab.io/bandwidth_reports/filecoin_mainnet/2025/6/19/bandwidth_per_msg_type_ts_plot.json" height="600px" >}}

## Bandwidth Distribution by Message Type
The final visualization illustrates the percentage of total bandwidth attributed to each message type. This helps identify which types of messages contribute the most to network traffic.

{{< plotly json="https://cdn.probelab.io/bandwidth_reports/filecoin_mainnet/2025/6/19/bandwidth_total_per_type_plot.json" height="600px" >}}
