---
title: Methodology
plotly: true
weight: 10
slug: methodology
aliases:
    - /discv5/methodology/
---

# discv5 reports methodology

## Report Title Format
Reports are titled after the year and calendar week (YYYY-WW) which the data in the corresponding report represent. For example, "Week 2024-09" represents calendar week 09 of 2024.

## DHT Crawling

We crawl the Ethereum `discv5` DHT network using the Nebula Crawler [[1](#references)].

The crawler gets a dump of the routing table of a remote peer by crafting `discv5` requests, matching every Kademlia [[2](#references)] bucket of the remote peer's routing table. Upon receiving a request for a node ID, a peer will compute in which bucket the target node ID falls in its own routing table, and return the content of this bucket. The buckets contain [Ethereum Node Records](https://ethereum.org/en/developers/docs/networking-layer/network-addresses/#enr) (ENR), with information of how to reach the corresponding nodes (fork_digest, attnet, tcp and udp addresses).

The crawler starts by connecting to [hardcoded bootstrap nodes](https://github.com/dennis-tra/nebula/blob/038116e979ade06ecc4518f287faace47c6580bf/config/bootstrap.go#L176-L199) extracted from [`eth2-networks`](https://github.com/eth-clients/eth2-networks/blob/master/shared/mainnet/bootstrap_nodes.txt). From there, it learns the routing table of the bootstrap nodes, and can continue crawling discovered peers until all of them have been queried. For future crawls, Nebula doesn't only depend on the hardcoded bootstrap nodes, but will also try to connect to all nodes discovered in previous crawls.

The `discv5` traffic described above is UDP only. In addition, the crawler tries to open a `libp2p` connection to all discovered nodes. Upon a successful `libp2p` connection, the crawler can learn additional information, such as user agent, supported protocols, quic addresses etc.

Discover more about Nebula's operations [here](https://probelab.io/tools/nebula/).

## Data filtering

The `discv5` DHT network is used by other projects besides Ethereum mainnet. Our reports mainly focus on how the `discv5` DHT network serves Ethereum mainnet. We only display data for nodes using the latest Ethereum mainnet fork digest `0x6a95a1a9` ([dencun](https://blog.ethereum.org/2024/02/27/dencun-mainnet-announcement)).

Sometimes, nodes refuse to open a libp2p connection with our crawler, for different reasons. Our plots only display _online_ nodes, that are answering the crawler's queries.

## Node classification

### User agents

User agent information is sourced from the libp2p identify protocol, which necessitates a successful libp2p connection. Given that libp2p connections can sometimes fail, despite a node's active participation in the DHT, we rely on user agent data acquired from either past or future successful libp2p connections associated with the same node ID. This method ensures that we can accurately capture and utilize user agent information, maintaining consistency in node identification and characterization over time.

### Geographical distribution

We employ MaxMind's GeoIP services to correlate IP addresses with their corresponding countries, utilizing the detailed resources provided by the [MaxMind GeoIP databases](https://www.maxmind.com/en/geoip-databases). This tool allows us to effectively determine the geographic locations of IP addresses, facilitating accurate country-level mapping and analysis.

### Cloud providers

We utilize Udger's datacenter identification services to map IP addresses to their respective commercial cloud providers, leveraging the comprehensive [Udger datacenter list](https://udger.com/resources/datacenter-list). This resource enables us to accurately identify and categorize IP addresses based on their association with known data centers.

## Differences to alternative tools/websites

Several teams have developed tools to report Ethereum's client diversity in websites such as [clientdiversity.org](https://clientdiversity.org/). Three tools we are aware of are: i) Sigma Prime's [blockprint](https://github.com/sigp/blockprint), ii) Migalab's [monitoreth.io](https://monitoreth.io/), and iii) [Rated.network](https://rated.network/). Given that there are variations in the diversity reported by these tools, here we attempt to clarify the differences in the methodologies used in an effort to explain these differences.

- Sigma Prime is counting the number of validators, while MigaLabs and ProbeLab are counting the number of clients (there can be multiple validators per client). Sigma Prime's blookprint tool calculates client diversity based on certain parameters of the proposed blocks. In particular, it reads how the blocks were built, and guesses which client was the builder.
- The differences between Migalabs and ProbeLab reports on client diversity are smaller and are primarily due to the different approaches to counting a node as online. Not all nodes accept incoming connections, and therefore, the counts are inherently divergent from each other.
- Rated.network has not revealed their methodology.


## References

[1] [Nebula Crawler](https://github.com/dennis-tra/nebula)

[2] [Kademlia paper](https://www.scs.stanford.edu/~dm/home/papers/kpos.pdf)

[3] [Ethereum Node Discovery Protocol v5](https://github.com/ethereum/devp2p/blob/master/discv5/discv5.md)
