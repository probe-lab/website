---
title: IPFS DHT
plotly: true
---

# IPFS DHT

{{< hint warning >}}
Please note that this website is a preview and is subject to change at any time. 
We strive to provide accurate and up-to-date information, but we cannot guarantee 
the completeness or reliability of the information presented during the preview. 
{{< /hint >}}

One of the core components of IPFS is the Distributed Hash Table (DHT), which enables nodes to locate and retrieve content from other nodes on the network. The IPFS DHT is a distributed key-value store that maps content addresses to the nodes that are currently storing that content. It works by dividing the content address space into small, hash-based "buckets" that are distributed across the network. When a node wants to find a piece of content, it queries the DHT with the content's hash, and the DHT returns the nodes that are currently storing that content. This allows content to be located and retrieved in a decentralized, efficient, and fault-tolerant manner. The IPFS DHT is a key component of the IPFS network, and is used by many IPFS applications and tools, including IPFS itself, as well as a variety of decentralized applications and systems built on top of IPFS.



## Performance

### Lookup Performance

{{< plotly json="/website/plots/dht-lookup-performance-overall-latest.json" height="600px" >}}

{{< plotly json="/website/plots/dht-lookup-performance-cdf-latest.json" height="600px" >}}

{{< plotly json="/website/plots/dht-lookup-performance-cdf-region-latest.json" height="600px" >}}

### Publish Performance

{{< plotly json="/website/plots/dht-publish-performance-overall-latest.json" height="600px" >}}

{{< plotly json="/website/plots/dht-publish-performance-cdf-latest.json" height="600px" >}}

{{< plotly json="/website/plots/dht-publish-performance-cdf-region-latest.json" height="600px" >}}

## Participation in the DHT


### DHT Server Software

The Nebula crawler records the software agents announced by peers registered in the IPFS DHT. 
These peers act as DHT servers and record provider records pointing to content available from other peers in the
network.

Note that he x-axis in the accompanying plot is represented using a log scale, which emphasizes the relatively smaller 
populations of software agents compared to the much larger use of Kubo (previously known as go-ipfs) within the DHT. 

{{< plotly json="/website/plots/top-dhtserver-agents-latest.json" height="800px" >}}


### Kubo Breakdown

{{< plotly json="/website/plots/kubo-version-distribution-latest.json" height="800px" >}}

{{< plotly json="/website/plots/recent-kubo-versions-over-time-latest.json" height="600px" >}}



### Client vs Server Node Estimate



## Health
