---
title: Bootstrappers
plotly: true
slug: bootstrappers
weight: 1
---

# Bootstrappers

Here's the place to discover the availability and uptime of Filecoin's bootstrapper nodes.

ProbeLab's monitoring infrastructure is probing the set of Filecoin bootstrappers indicated in this repository: https://github.com/filecoin-project/lotus/blob/master/build/bootstrap/mainnet.pi. Our probes are checking up on nodes' connectivity _every 5 mins_, utilising all available transports (i.e., TCP, QUIC, Websocket, WebTransport).

## How to read the dashboards

- Upon _three_ failed attempts to connect to a bootstrap node with any of the transport options (i.e., if none of the attempts with any of the transports is successful), the node is considered offline and an alert is triggered.
- If a node is offline, the corresponding tile at the top section of the dashboard turns red.
- Upon each failed attempt to connect to a node with a particular transport, the disruption is indicated as a red stripe in the timeseries dashboard of the corresponding transport.
- Unless all attempts to connect to a node with any of the transports fail, the top section tile remains green.

Below is the list of PeerIDs that are currently being probed in the order in which they appear under each transport protocol. For the full `multiaddress` of each node, please check the last section in the dashboard "Probed Multiaddresses".

| Operator| PeerID |
|---|---|
|FIL DevTTY| `12D3KooWAke3M2ji7tGNKx3BQkTHCyxVhtV1CN68z6Fkrpmfr37F`|
|Glif.io| `12D3KooWBF8cpp65hp2u9LK5mh19x67ftAam84z9LsfaquTDSBpt`|
|ChainSafe 1| `12D3KooWGnkd9GQKo3apkShQDaq1d6cKJJmsVe6KiQkacUk1T8oZ`|
|ChainSafe 2| `12D3KooWHQRSDFv4FvAjtU32shQ7znz7oRbLBryXzZ9NMK2feyyH`|
|ChainSafe 0| `12D3KooWKKkCZbcigsWTEu1cgNetNbZJqeNtysRtFpq7DTqw3eqH`|
|filincubator.com| `QmQu8C6deXwKvJP2D8B6QGyhngc3ZiDnFzEHBDx8yeBXST`|
	
{{< iframe src="https://grafana-public.probelab.io/public-dashboards/e65e04dbd0624714a81d114b221f1f91" title="Filecoin Bootstrappers Grafana Dashboard" css-class="filecoin-bootstrapper-grafana-iframe" width="100%" height="100vh">}}
