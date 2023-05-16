---
title: IPFS KPIs
weight: 10
plotly: true
---
# IPFS KPIs

{{< hint warning >}}
Please note that this website is a preview and is subject to change at any time. 
We strive to provide accurate and up-to-date information, but we cannot guarantee 
the completeness or reliability of the information presented during the preview. 
{{< /hint >}}


IPFS relies on the coordinated participation of a swarm of independent peers to function correctly. Therefore, measuring the performance and health of the network is crucial for ensuring its reliability and efficiency. In this context, we summarise a number of network characteristics as key performance indicators. 
Our KPIs are currently focusing primarily on the public [IPFS DHT](https://docs.ipfs.tech/concepts/dht/) (although we do report on other parts of the architecture too). Understanding that IPFS is an ecosystem of content routing subsystems, over time, we plan to expand to other content routing subsystems too.

## Network Size & Stability

### Client vs Server Node Estimate

The total number of peers in the network is estimated using the number of unique [Peer IDs](https://docs.ipfs.tech/concepts/glossary/#peer-id) seen by Protocol Labs' [bootstrap nodes](https://docs.ipfs.tech/concepts/glossary/#bootstrap-node). The number of unique [DHT Server](https://docs.ipfs.tech/concepts/dht/#routing-tables) peer IDs identified by the [Nebula crawler](https://github.com/dennis-tra/nebula) is then subtracted from the total number of peers (seen by the bootstrappers) to estimate the number of peers that exclusively function as DHT Clients.

{{< plotly json="../../plots/latest/ipfs-servers-vs-clients.json" height="400px" >}}

## Content Routing 

IPFS employs several content routing subsystems, with the [Kademlia](https://docs.ipfs.tech/concepts/dht/#kademlia) Distributed Hash Table (DHT) being the most established. Within the network, peers commonly use this system to locate other peers that hold the content being requested. We measure the availability of DHT server nodes and the latency of [DHT lookups](https://docs.ipfs.tech/concepts/dht/#lookup-algorithm) for random content.

### DHT server availability

We categorize [DHT Server](https://docs.ipfs.tech/concepts/dht/#routing-tables) nodes with regard to their "availability" as follows:
- "Online": the node has been found online for 80% of time or more.
- "Mostly Online": the node has been found online between 40%-80% of time.
- "Mostly Offline": the node has been found online between 10%-40% of time.
- "Offline": the node has been found online less than 10% of time.  

{{< plotly json="../../plots/latest/dht-availability-classified-online.json" height="250px" >}}

{{< plotly json="../../plots/latest/dht-availability-classified-overall.json" height="600px" >}}

### DHT Lookup performance

We have instrumented the following experiment to capture the [DHT Lookup](https://docs.ipfs.tech/concepts/dht/#lookup-algorithm) performance over time and from several different geographic locations.
We have set up IPFS DHT Server nodes in seven (7) different geographic locations. Each node periodically publishes a unique [Content Identifier (CID)](https://docs.ipfs.tech/concepts/content-addressing/#content-identifiers-cids) and makes it known to the rest of the nodes, who subsequently request it through the DHT (acting as clients). This way we capture the DHT Lookup performance from each location.
In this section we present the average performance over all regions.

{{< plotly json="../../plots/latest/dht-lookup-performance-quartiles.json" height="250px" >}}

{{< plotly json="../../plots/latest/dht-lookup-performance-long.json" height="600px" >}}

### IPNI utilization

[IPNI](https://github.com/ipni) is a set of protocols that describe how data can be indexed across the IPFS and Filecoin networks. Network indexers complement the [IPFS DHT](https://docs.ipfs.tech/concepts/dht/) to enable peers to locate content-addressed data. The data in the plot below shows the number of requests made per day to the network indexers operated by [cid.contact](https://cid.contact/).

{{< plotly json="../../plots/latest/ipni-requests-overall.json" height="600px" >}}


## Websites and Traffic

A common use-case for IPFS is hosting websites, addressed using [IPNS](https://docs.ipfs.tech/concepts/dht/) or [DNSLink](https://docs.ipfs.tech/concepts/dnslink/). We monitor the time it takes to load sample websites through a browser and the number of requests to the public Protocol Labs operated [IPFS gateways](https://docs.ipfs.tech/concepts/ipfs-gateway/). 

{{< plotly json="../../plots/latest/websites-ttfb-quartiles.json" height="250px" >}}

The following plot shows the total number of requests made per day to the [public IPFS gateways](https://docs.ipfs.tech/concepts/ipfs-gateway/#gateway-providers) operated by Protocol Labs (ipfs.io and dweb.link).

{{< plotly json="../../plots/latest/gateway-requests-overall.json" height="600px" >}}


