---
title: IPFS DHT
weight: 20
plotly: true
---

# IPFS DHT

{{< hint warning >}}
Please note that this website is a preview and is subject to change at any time. 
We strive to provide accurate and up-to-date information, but we cannot guarantee 
the completeness or reliability of the information presented during the preview. 
{{< /hint >}}

One of the core components of IPFS is the Distributed Hash Table (DHT), which enables nodes to locate and retrieve content from other nodes on the network. The IPFS DHT is a distributed key-value store that maps content addresses to the nodes that are currently storing that content. It works by dividing the content address space into small, hash-based "buckets" that are distributed across the network. When a node wants to find a piece of content, it queries the DHT with the content's hash, and the DHT returns the nodes that are currently storing that content. This allows content to be located and retrieved in a decentralized, efficient, and fault-tolerant manner. The IPFS DHT is a key component of the IPFS network, and is used by many IPFS applications and tools, including IPFS itself, as well as a variety of decentralized applications and systems built on top of IPFS.

## Health

As a distributed system, IPFS relies on the coordinated participation of multiple nodes and servers to function correctly. Monitoring the
availability and long term stability of DHT servers over time can give insight into the health of the network. High churn in the network
can make content harder to locate and lead to longer retrieval times. Measuring DHT server availability and expected lifetimes can help 
assess the health and overall efficiency of the network.

### Availability

The Nebula crawler attempts to connect to peers in the IPFS DHT periodically. When a new peer is discovered, the crawler records the start of a session of availability and extends the session length with every successful connection attempt. However, a failed connection terminates the session, and a later successful attempt starts a new session. Peers can have multiple sessions of availability during each measurement period. 

{{< plotly json="../../plots/latest/dht-availability-classified-current.json" height="250px" >}}


In the following plots, a peer is classified as "online" if it was available for at least 80% of the measurement period. If a peer was available between 40% and 80% of the period, it is considered "mostly online," while "mostly offline" indicates availability between 10% and 40% of the time. Any peer that was available for less than 10% of the period is classified as "offline."

{{< plotly json="../../plots/latest/dht-availability-classified-overall.json" height="600px" >}}

{{< plotly json="../../plots/latest/dht-availability-classified-region.json" height="600px" >}}

### Churn

{{< plotly json="../../plots/latest/dht-peers-entering-leaving.json" height="600px" >}}

This plot displays the count of unique peer IDs that joined and left the network during the measurement period. The term "entered" refers to a peer that was offline at the start of the period but appeared within it and remained online throughout. "left" refers to a peer that was online at the start of the period but went offline and did not come back online before the end of the measurement period.

## Performance

Measuring the time it takes to publish and retrieve information from the IPFS DHT is crucial for understanding the network's performance and identifying areas for improvement. It allows us to assess the network's efficiency in different regions, which is essential for global-scale applications that rely on the IPFS DHT. Measuring performance across different regions helps identify potential bottlenecks and optimize content delivery. 

(TBD: link to explanation of methodology)


### Lookup Performance

{{< plotly json="../../plots/latest/dht-lookup-performance-quartiles.json" height="250px" >}}

The following plots show the distribution of timings for looking up provider records from various points across the world. 

{{< plotly json="../../plots/latest/dht-lookup-performance-overall.json" height="600px" >}}

{{< plotly json="../../plots/latest/dht-lookup-performance-cdf.json" height="600px" >}}

{{< plotly json="../../plots/latest/dht-lookup-performance-cdf-region.json" height="600px" >}}

### Publish Performance

{{< plotly json="../../plots/latest/dht-publish-performance-quartiles.json" height="250px" >}}

The following plots show the distribution of timings for publishing provider records from various points across the world.

{{< plotly json="../../plots/latest/dht-publish-performance-overall.json" height="600px" >}}

{{< plotly json="../../plots/latest/dht-publish-performance-cdf.json" height="600px" >}}

{{< plotly json="../../plots/latest/dht-publish-performance-cdf-region.json" height="600px" >}}

## Participation in the DHT

Measuring participation in the IPFS DHT is crucial to understanding the health and effectiveness of the network. A diverse and wide participation of software agents and peers helps ensure a robust and resilient network. Such diversity helps prevent centralization, provides greater redundancy, and increases the chances of content availability. Moreover, a wide participation allows for a more efficient distribution of content, improves load balancing, and can lead to faster content retrieval. 

### Client vs Server Node Estimate

The plot presented below illustrates an estimate of the number of peers that exclusively function as clients. This estimate is derived by deducting the total number of unique peer IDs observed by the bootstrap nodes, operated by Protocol Labs, from the number of unique peer IDs visited by the Nebula crawler during the same period. Additionally, the plot also shows the number of unique IP addresses observed by the Nebula crawler.

{{< plotly json="../../plots/latest/ipfs-servers-vs-clients.json" height="400px" >}}

### DHT Server Software

The Nebula crawler records the software agents announced by peers registered in the IPFS DHT. 
These peers act as DHT servers and record provider records pointing to content available from other peers in the network.

{{< plotly json="../../plots/latest/top-dhtserver-agents.json" height="800px" >}}

Note that the x-axis in the above plot is represented using a log scale, which emphasizes the relatively smaller  populations of software agents compared to the much larger use of Kubo (previously known as go-ipfs) within the DHT. 



### Kubo Breakdown

Kubo is the most prevelant software used for peers participating in the DHT. It adheres to a regular release cycle to introduces new features and improvements in performance, stability, and security. Measuring the distribution of Kubo versions provides insights into the adoption rate of new features and improvements and potential issues related to backward compatibility during protocol upgrades. 

{{< plotly json="../../plots/latest/kubo-version-distribution.json" height="800px" >}}

{{< plotly json="../../plots/latest/recent-kubo-versions-over-time.json" height="600px" >}}



