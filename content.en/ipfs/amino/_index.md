---
title: Amino DHT
weight: 20
plotly: true
aliases:
    - /ipfsdht/
---

# IPFS DHT

One of the core components of IPFS is the [Distributed Hash Table (DHT)](https://docs.ipfs.tech/concepts/dht/#distributed-hash-tables-dhts), which enables nodes to locate and retrieve content from other nodes on the network. The IPFS DHT is a distributed key-value store that maps content addresses (_aka_ CIDs) to the nodes that are currently storing that content. It works by dividing the content address space into small, hash-based ["buckets"](https://docs.ipfs.tech/concepts/dht/#peer-buckets) that are distributed across the network. When a node wants to find a piece of content, it queries the DHT with the content's [CID](https://docs.ipfs.tech/concepts/content-addressing/#what-is-a-cid), and the DHT returns the nodes that are currently storing that content, in the form of [Provider Records](https://docs.ipfs.tech/concepts/dht/#provider-records). This allows content to be located and retrieved in a decentralized, efficient, and fault-tolerant manner. The IPFS DHT is a key component of the IPFS network, and is used by many IPFS implementations, tools, as well as a variety of decentralized applications and systems built on top of IPFS.

## Health

As a distributed system, IPFS relies on the coordinated participation of multiple nodes and servers to function correctly. Monitoring the
availability and long term stability of DHT servers over time can give insight into the health of the network. High churn in the network
can make content harder to locate and lead to longer [retrieval](https://docs.ipfs.tech/concepts/lifecycle/#_3-retrieval) times. Measuring [DHT server](https://docs.ipfs.tech/concepts/dht/#qualification) availability and expected lifetimes can help
assess the health and overall efficiency of the network.

### Availability

The [Nebula crawler](/tools/nebula) attempts to connect to [DHT Server](https://docs.ipfs.tech/concepts/dht/#qualification) peers in the IPFS DHT periodically. When a new DHT Server peer is discovered, the crawler records the start of a session of availability and extends the session length with every successful connection attempt. However, a failed connection terminates the session, and a later successful attempt starts a new session. Peers can have multiple sessions of availability during each measurement period.

In the following, a peer is classified as "online" if it was available for at least 80% of the measurement period. If a peer was available between 40% and 80% of the period, it is considered "mostly online," while "mostly offline" indicates availability between 10% and 40% of the time. Any peer that was available for less than 10% of the period is classified as "offline."

#### DHT Server Availability

{{< plotly json="../../plots/latest/dht-availability-classified-current.json" height="250px" >}}

#### DHT Server Availability, classified over time

{{< plotly json="../../plots/latest/dht-availability-classified-overall.json" height="600px" >}}

#### DHT Server Availability, classified by region

{{< plotly json="../../plots/latest/dht-availability-classified-region.json" height="600px" >}}

### Churn

The following displays the count of unique [Peer IDs](https://docs.ipfs.tech/concepts/glossary/#peer-id) that joined and left the network during the measurement period. The term "entered" refers to a peer that was offline at the start of the measurement period but appeared within it and remained online throughout. The term "left" refers to a DHT Server peer that was online at the start of the measurement period but went offline and did not come back online before the end of the measurement period.

{{< plotly json="../../plots/latest/dht-peers-entering-leaving.json" height="600px" >}}

The cumulative distribution of session lengths for peers found in the network is shown below. It plots the cumulative fraction of sessions longer than a given time, shown along the x-axis.

{{< plotly json="../../plots/latest/dht-peers-churn-cdf-overall.json" height="600px" >}}

## Capabilities

### Transports {#dht-transport-distribution}

{{< plotly json="../../plots/latest/dht-transport-distribution.json" height="400px" >}}

The above graph shows the distribution of transports that peers in the IPFS DHT are listening on. `WS` = WebSocket, `WSS` = Websocket Secure, `WT` = WebTransport. The data comes from a single crawl at the beginning of the indicated date in the lower right corner.

## Performance

Measuring the time it takes to [publish](https://docs.ipfs.tech/concepts/lifecycle/#_2-pinning) and [retrieve](https://docs.ipfs.tech/concepts/lifecycle/#_3-retrieval) information from the IPFS DHT is crucial for understanding the network's performance and identifying areas for improvement. It allows us to assess the network's efficiency in different regions, which is essential for global-scale applications that rely on the IPFS DHT. Measuring performance across different regions helps identify potential bottlenecks and optimize content delivery.

We have instrumented the following experiment to capture the [DHT Lookup](https://docs.ipfs.tech/concepts/dht/#lookup-algorithm) performance over time and from several different geographic locations.
We have set up IPFS DHT Server nodes in seven (7) different geographic locations. Each node periodically publishes a unique CID and makes it known to the rest of the nodes, who subsequently request it through the DHT (acting as clients). This way we capture the entire performance spectrum of the DHT, i.e., both the DHT Publish and the Lookup performance from each location.
In this section we present the average performance over all regions, as well as per region for both the DHT Lookup Performance and the DHT Publish Performance.

### Lookup Performance

{{< plotly json="../../plots/latest/dht-lookup-performance-quartiles.json" height="250px" >}}

The following plots show the distribution of timings for looking up [provider records](https://docs.ipfs.tech/concepts/dht/#provider-records) from various points across the world.

#### Median DHT Lookup Performance over time

{{< plotly json="../../plots/latest/dht-lookup-performance-overall.json" height="600px" >}}

#### DHT Lookup Performance Distribution

{{< plotly json="../../plots/latest/dht-lookup-performance-cdf.json" height="600px" >}}

#### DHT Lookup Performance Distribution, by region

{{< plotly json="../../plots/latest/dht-lookup-performance-cdf-region.json" height="600px" >}}

### Publish Performance

{{< plotly json="../../plots/latest/dht-publish-performance-quartiles.json" height="250px" >}}

The following plots show the distribution of timings for publishing provider records from various points across the world.

#### Median DHT Publish Performance over time

{{< plotly json="../../plots/latest/dht-publish-performance-overall.json" height="600px" >}}

#### DHT Publish Performance Distribution

{{< plotly json="../../plots/latest/dht-publish-performance-cdf.json" height="600px" >}}

#### DHT Publish Performance Distribution, by region

{{< plotly json="../../plots/latest/dht-publish-performance-cdf-region.json" height="600px" >}}

## Participation in the DHT

Measuring participation in the IPFS DHT is crucial to understanding the health and effectiveness of the network. A diverse and wide participation of software agents and peers helps ensure a robust and resilient network. Such diversity helps prevent centralization, provides greater redundancy, and increases the chances of content availability. Moreover, a wide participation allows for a more efficient distribution of content, improves load balancing, and can lead to faster content retrieval.

### Client vs Server Node Estimate

The plot presented below illustrates an estimate of the number of peers that exclusively function as clients. This estimate is derived by deducting the total number of unique peer IDs observed by the [bootstrap nodes](https://docs.ipfs.tech/concepts/glossary/#bootstrap-node), operated by Protocol Labs, from the number of unique [Peer IDs](https://docs.ipfs.tech/concepts/glossary/#peer-id) visited by the [Nebula crawler](/tools/nebula/) during the same period. Additionally, the plot also shows the number of unique IP addresses observed by the Nebula crawler.

{{< plotly json="../../plots/latest/ipfs-servers-vs-clients.json" height="400px" >}}

### DHT Server Software

The Nebula crawler records the software agents announced by peers registered in the IPFS DHT.
These peers act as DHT servers and record provider records pointing to content available from other peers in the network.

#### Most Frequent DHT Server Agents

{{< plotly json="../../plots/latest/top-dhtserver-agents.json" height="800px" >}}

Note that the x-axis in the above plot is represented using a log scale, which emphasizes the relatively smaller  populations of software agents compared to the much larger use of [Kubo](https://github.com/ipfs/kubo) (previously known as go-ipfs) within the DHT.

### Active Kubo Versions

[Kubo](https://github.com/ipfs/kubo) is the most prevelant software used for peers participating in the DHT. It adheres to a regular release cycle to introduce new features and improvements in performance, stability, and security. Measuring the distribution of Kubo versions provides insights into the adoption rate of new features and improvements and potential issues related to backward compatibility during protocol upgrades.

#### Kubo Version Distribution

{{< plotly json="../../plots/latest/kubo-version-distribution.json" height="800px" >}}

#### Recent Kubo Versions Over Time

In the following we show the change in distribution of the nine most recent releases of Kubo each week, grouping all prior releases into the "all others" category. When a new version appears, the oldest of the nine is moved to the "other" category.

{{< plotly json="../../plots/latest/recent-kubo-versions-over-time.json" height="600px" >}}

## DHT Key Density Monitoring

In Kademlia, every object indexed by the DHT requires a binary identifier. In the libp2p DHT implementation, peers are identified by the digest of `sha256(peer_id)` and CIDs are identified by the digest of `sha256(cid)`. This Kademlia identifier determines the location of an object within the Kademlia XOR keyspace.

The following plots examine the peer distribution within the keyspace, aiding in the identification of potential [Sybil](https://en.wikipedia.org/wiki/Sybil_attack) and eclipse attacks.

### Keyspace population distribution

**Description**: The plot illustrates the number of peers whose Kademlia identifier matches each prefix for all prefixes of a given size, for a given network crawl. Since the density of keyspace regions follows a [Poisson](https://en.wikipedia.org/wiki/Poisson_distribution) distribution, it is expected to observe some regions that are significantly more populated than others.

**How to read the plot:** The selected `depth` indicates the prefix size. There are `2^i` distinct prefixes at depth `i`. The slider at the bottom of the plot enables visualization of the population evolution over time across multiple crawls.

**What to look out for:** The red dashed line represents the expected density per region, corresponding to the number of peers matching a prefix. A bar significantly exceeding the expected density suggests that a region of the keyspace might be under an eclipse attack, especially if the value is above the risk threshold.

{{< plotly json="../../plots/latest/ipfs-regions-population.json" height="600px" >}}

### Keyspace density distribution

**Description:** As previously mentioned, the keyspace population follows a [Poisson](https://en.wikipedia.org/wiki/Poisson_distribution) distribution, which can make identifying outliers challenging. The plot below counts the number of regions for each population size and facilitates the identification and analysis of outliers. While it is normal for some regions to have populations above the average, the plot enables us to quantify these deviations.

**How to read the plot:** The red dashed line represents the expected number of regions for each population size. Note that the Poisson distribution is more evident at greater depths (longer prefix size), while analyzing data at lower depths provides limited insights. It is recommended to read the plot for depths between 9 and 13.

**What to look out for:** If an isolated bar appears on the far right beyond the risk threshold line, it may indicate a potential eclipse attack, warranting further investigation.

{{< plotly json="../../plots/latest/ipfs-density-distributions.json" height="600px" >}}
