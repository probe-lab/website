---
title: Week 2025-43 day 02
plotly: true
slug: 2025-43-02
weight: 1027489
---

# Filecoin report for GossipSub's Control Messages 2025-43 day 02

The following results show measurement data that were collected in **calendar week** 43  day 02 **of** 
2025 (`2025-10-21`).

This report provides charts for several metrics for all of GossipSub's topics on the Filecoin network.
The report focuses primarily at the Control Messages that GossipSub exchanges with other peers in the network.
We've produced these reports in collaboration with the Filecoin Foundation using our tool [Hermes](/tools/hermes).

## Sent and Received Control Messages

The following chart shows the number of Control Messages received by GossipSub over  intervals. The plot breaks down the number of `IHAVE`, `IWANT`, and `IDONTWANT` messages included in RPCs.

Some implementations, such as [`go-libp2p-pubsub`](https://github.com/libp2p/go-libp2p-pubsub), aggregate multiple messages within a single RPC to optimize communication. This chart counts the unique messages within RPCs rather than individual RPC transmissions.

> **Note:** `IHAVE` messages still include multiple topics within a single Control Message.

{{< plotly json="https://cdn.probelab.io/control_rpcs_reports/filecoin_mainnet/2025/10/21/control_rpc_ratio_ts_plot.json" height="600px" >}}

---

## Sent and Received Message IDs in Control RPCs

The following chart visualizes the number of **Message IDs** sent and received within each type of Control Message, aggregated over  intervals.

{{< plotly json="https://cdn.probelab.io/control_rpcs_reports/filecoin_mainnet/2025/10/21/control_rpc_effectiveness_msg_id_ts_plot.json" height="600px" >}}

---

## Per-Topic Sent and Received Message IDs in Control RPCs

This chart extends the previous visualization by breaking down sent and received **Message IDs** per topic, aggregated over  intervals.

{{< plotly json="https://cdn.probelab.io/control_rpcs_reports/filecoin_mainnet/2025/10/21/control_rpc_effectiveness_msg_id_per_topic_ts_plot.json" height="600px" >}}