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

## Network Size & Stability

### Client vs Server Node Estimate

The total number of peers in the network is estimated using the number of unique peer IDs seen by Protocol Labs' bootstrap nodes. The number of unique DHT server peer IDs visited by the Nebula crawler is subtracted estimate the number of peers that exclusively function as clients.

{{< plotly json="../plots/ipfs-servers-vs-clients-latest.json" height="400px" >}}

## Content Routing 

IPFS employs several content routing systems, with the Kademlia distributed hash table being the most established. Within the network, peers commonly use this system to locate other peers that hold the content being requested. We measure the availability of DHT server nodes and the latency of DHT lookups for random content.

### DHT server availability

{{< plotly json="../plots/dht-availability-classified-online-latest.json" height="250px" >}}

{{< plotly json="../../plots/dht-availability-classified-overall-latest.json" height="600px" >}}

### DHT server performance

{{< plotly json="../plots/dht-lookup-performance-quartiles-latest.json" height="250px" >}}

{{< plotly json="../../plots/dht-lookup-performance-overall-latest.json" height="600px" >}}

### IPNI utilization

IPNI is set of protocols that describe how data can be indexed across the IPFS and Filecoin networks. Network indexers complement the IPFS DHT to enable peers to locate content-addressed data. The data in the plot below shows the number of requests made per day to the network indexers operated by cid.contact.

{{< plotly json="../plots/ipni-requests-overall-latest.json" height="600px" >}}


## Websites and Traffic

A common use-case for IPFS is hosting websites, addressed using IPNS or DNSLink. We monitor the time it takes to load sample websites through a browser and the number of requests to the public Protocol Labs operated IPFS gateways. 

{{< plotly json="../plots/websites-ttfb-quartiles-latest.json" height="250px" >}}

The following plot shows the total number of requests made per day to the public IPFS gateways operated by Protocol Labs (ipfs.io and dweb.link).

{{< plotly json="../plots/gateway-requests-daily-latest.json" height="600px" >}}


