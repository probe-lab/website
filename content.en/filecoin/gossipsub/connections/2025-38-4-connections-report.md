---
title: Week 2025-38 day 04
plotly: true
slug: 2025-38-04
weight: 1027522
---

# Filecoin report for GossipSub Connections and Disconnections 2025-38 day 04

The following results show measurement data that were collected in **calendar week** 38  day 04 **of** 
2025 (`2025-09-18`).

This report provides a set of charts summarizing the most relevant metrics about the connection of a Libp2p node in the Filecoin network.
We've produced these reports in collaboration with the Filecoin Foundation, using our tool [Hermes](/tools/hermes/).

## Duration of the Connections
The following graph shows the Cumulative Distribution Function (CDF) of the connection duration that our control Hermes instance established with nodes from the Filecoin network.

{{< plotly json="https://cdn.probelab.io/connections_reports/filecoin_mainnet/2025/9/18/connection_duration_cdf_plot.json" height="600px" >}}

## Connections and Disconnections over time
The following graph shows all connection and disconnection events that the Hermes node recorded over the day based on the direction (inbound or outbound).
{{< plotly json="https://cdn.probelab.io/connections_reports/filecoin_mainnet/2025/9/18/connect_disconnect_ts_plot.json" height="600px" >}}