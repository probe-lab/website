---
title: IPFS KPIs
weight: 10
plotly: true
aliases:
    - /ipfskpi/
---
# IPFS KPIs

IPFS relies on the coordinated participation of a swarm of independent peers to function correctly. Therefore, measuring the performance and health of the network is crucial for ensuring its reliability and efficiency. In this context, we summarise a number of network characteristics as key performance indicators. 
Our KPIs are currently focusing primarily on the public [IPFS DHT](https://docs.ipfs.tech/concepts/dht/) (although we do report on other parts of the architecture too). Understanding that IPFS is an ecosystem of [content routing subsystems](https://docs.ipfs.tech/concepts/how-ipfs-works/#subsystems-overview), over time, we plan to expand to other content routing subsystems too.

## Network Size & Stability

### Client vs Server Node Estimate

The total number of peers in the network is estimated using the number of unique [Peer IDs](https://docs.ipfs.tech/concepts/glossary/#peer-id) seen by Protocol Labs' [bootstrap nodes](https://docs.ipfs.tech/concepts/glossary/#bootstrap-node). The number of unique [DHT Server](https://docs.ipfs.tech/concepts/dht/#routing-tables) peer IDs identified by the [Nebula crawler](/tools/nebula/) is then subtracted from the total number of peers (seen by the bootstrappers) to estimate the number of peers that exclusively function as DHT Clients.

{{< plotly json="../../plots/latest/ipfs-servers-vs-clients.json" height="400px" >}}

### Unique Software Agents

The total number of unique software agents operating in the network is estimated from those seen by Protocol Labs' [bootstrap nodes](https://docs.ipfs.tech/concepts/glossary/#bootstrap-node) when a peer connects. The number of unique agents seen by the [Nebula crawler](/tools/nebula/) when crawling the IPFS DHT is included for comparison. The software agent strings have not been refined or processed, resulting in the count treating major and minor versions of each software agent as distinct entries.


{{< plotly json="../../plots/latest/ipfs-unique-agents.json" height="400px" >}}


## Content Routing 

IPFS employs several content routing subsystems, with the [Kademlia](https://docs.ipfs.tech/concepts/dht/#kademlia) Distributed Hash Table (DHT) being the most established. Within the network, peers commonly use this system to locate other peers that hold the content being requested. We measure the availability of DHT server nodes and the latency of [DHT lookups](https://docs.ipfs.tech/concepts/dht/#lookup-algorithm) for random content.

### DHT server availability

We categorize [DHT Server](https://docs.ipfs.tech/concepts/dht/#routing-tables) nodes with regard to their "availability" as follows:
- "Online": the node has been found online for 80% of time or more.
- "Mostly Online": the node has been found online between 40%-80% of time.
- "Mostly Offline": the node has been found online between 10%-40% of time.
- "Offline": the node has been found online less than 10% of time.  

Data is collected using the [Nebula crawler](/tools/nebula/).

#### DHT Server Availability

{{< plotly json="../../plots/latest/dht-availability-classified-current.json" height="250px" >}}


#### DHT Server Availability, classified over time

{{< plotly json="../../plots/latest/dht-availability-classified-overall.json" height="600px" >}}

### DHT Lookup performance

We use [Parsec](/tools/parsec) to capture the [DHT Lookup](https://docs.ipfs.tech/concepts/dht/#lookup-algorithm) performance over time and from several different geographic locations. In this section we present the average performance over all regions.

#### DHT Lookup Performance

{{< plotly json="../../plots/latest/dht-lookup-performance-quartiles.json" height="250px" >}}

#### Historic DHT Lookup Performance

The historic trend over time is currently provided by [Parsec](/tools/parsec). Before 21 April 2023 an older script was used that queried IPFS preload servers to measure DHT lookup performance.

{{< plotly json="../../plots/latest/dht-lookup-performance-long.json" height="600px" >}}

### IPNI utilization

[IPNI](https://github.com/ipni) is a set of protocols that describe how data can be indexed across the IPFS and Filecoin networks. Network indexers complement the [IPFS DHT](https://docs.ipfs.tech/concepts/dht/) to enable peers to locate content-addressed data. The data in the plot below shows the number of requests made per day to the network indexers operated by [cid.contact](https://cid.contact/).

{{< plotly json="../../plots/latest/ipni-requests-overall.json" height="600px" >}}


## Websites

A common use-case for IPFS is hosting websites, addressed using [IPNS](https://docs.ipfs.tech/concepts/dht/) or [DNSLink](https://docs.ipfs.tech/concepts/dnslink/). We monitor the time it takes to load sample websites using the [Tiros](/tools/tiros) monitoring tool.

#### Time to First Byte for IPFS Hosted Websites

{{< plotly json="../../plots/latest/websites-ttfb-quartiles.json" height="250px" >}}


## HTTP Gateway Usage

#### Gateway Requests

The following plot shows the total number of requests made per day to the ipfs.io [public IPFS gateways](https://docs.ipfs.tech/concepts/ipfs-gateway/#gateway-providers) operated by [Interplanetary Shipyard](https://ipshipyard.com/). Data is collated from nginx access logs that front the gateway infrastructure.

{{< plotly json="../../plots/latest/gateway-requests-overall.json" height="600px" >}}


