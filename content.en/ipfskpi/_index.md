---
title: IPFS KPIs
plotly: true
---
# IPFS KPIs

IPFS relies on the coordinated participation of a swarm of independent peers to function correctly. Therefore, measuring the performance and health of the network is crucial for ensuring its reliability and efficiency. In this context, we summarise a number of network characteristics as key performance indicators. 

## Network Size & Stability

### Client vs Server Node Estimate

{{< plotly json="../plots/ipfs-servers-vs-clients-latest.json" height="400px" >}}

## Content Routing 

IPFS employs several content routing systems, with the Kademlia distributed hash table being the most established. Within the network, peers commonly use this system to locate other peers that hold the content being requested. We measure the availability of DHT server nodes and the latency of DHT lookups for random content.

### DHT server availability

{{< plotly json="../plots/dht-availability-classified-online-latest.json" height="250px" >}}

{{< plotly json="../../plots/dht-availability-classified-overall-latest.json" height="600px" >}}

### DHT server performance

{{< plotly json="../plots/dht-lookup-performance-quartiles-latest.json" height="250px" >}}

{{< plotly json="../../plots/dht-lookup-performance-overall-latest.json" height="600px" >}}

## Websites and Traffic

A common use-case for IPFS is hosting websites, addressed using IPNS or DNSLink. We monitor the time it takes to load sample websites through a browser and the number of requests to the public Protocol Labs operated IPFS gateways. 

{{< plotly json="../plots/websites-ttfb-quartiles-latest.json" height="250px" >}}

TBD: IPFS hosted websites, time to first byte, median, over time

TBD: HTTP Originated Requests (gateways), over time
