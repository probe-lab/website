---
title: Week 2025-34 day 07
plotly: true
slug: 2025-34-07
weight: 1027547
---

# Filecoin report for GossipSub PeerExchange 2025-34 day 07

The following results show measurement data that were collected in **calendar week** 34  day 07 **of** 
2025 (`2025-08-24`).

This report provides a set of charts summarizing the peer exchange related metrics for a Libp2p node in the Filecoin network.
We've produced these reports in collaboration with the Filecoin Foundation, using our tool [Hermes](/tools/hermes/).

## Shared Peer Records over Prune Messages
The next graph shows, as a histogram, the distribution of how many **Prune** messages (Y axis) included how many peers (X axis), 
illustrating the effectiveness of the GossipSub PeerExchange in the Filecoin network.

{{< plotly json="https://cdn.probelab.io/peer_exchange_reports/filecoin_mainnet/2025/8/24/peer_exchange_total_prunes_with_peers_plot.json" height="600px" >}}

## Number of Unique Peers Shared over Time
The following graph shows the number of peers that were exchanged with the Hermes instance over time and topic.

{{< plotly json="https://cdn.probelab.io/peer_exchange_reports/filecoin_mainnet/2025/8/24/peer_exchange_total_peers_ts_plot.json" height="600px" >}}