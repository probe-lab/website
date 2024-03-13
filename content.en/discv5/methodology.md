---
title: "discv5 reports methodology"
plotly: true
slug: methodology
---

# discv5 reports methodology

## DHT Crawling

We crawl the `discv5` network using the Nebula Crawler [[1](#references)].

The crawler can get a dump of the routing table of a remote peer by crafting `discv5` requests nodes, matching every Kademlia [[2](#references)] bucket of the remote peer's routing table. Upon receiving a request for a node ID, a peer will compute in which bucket the target node ID falls in its own routing table, and return the content of this bucket. The buckets contain [Ethereum Node Records](https://ethereum.org/en/developers/docs/networking-layer/network-addresses/#enr) (ENR), containing information of how to reach the corresponding nodes (fork_digest, attnet, tcp and udp addresses).

The crawler starts by connecting to [hardcoded bootstrap nodes](https://github.com/dennis-tra/nebula/blob/038116e979ade06ecc4518f287faace47c6580bf/config/bootstrap.go#L176-L199) extracted from [`eth2-networks`](https://github.com/eth-clients/eth2-networks/blob/master/shared/mainnet/bootstrap_nodes.txt). From there, it learns the routing table of the bootstrap nodes, and can continue crawling discovered peers until all of them have been queried. For future crawls, Nebula doesn't only depend on the hardcoded bootstrap nodes, but will also try to connect to all nodes discovered in previous crawls.

The `discv5` traffic described above is UDP only. In addition, the crawler tries to open a `libp2p` connection to all discovered nodes. Upon a successful `libp2p` connection, the crawler can learn additional information, such as user agent, supported protocol, quic addresses etc.

## Data filtering

The `discv5` DHT network is used by other projects besides Ethereum mainnet. Our reports mainly focus on how the `discv5` DHT network serves Ethereum mainnet. We only display data for nodes using the latest Ethereum mainnet fork digest `0x6a95a1a9` ([dencun](https://blog.ethereum.org/2024/02/27/dencun-mainnet-announcement)).

Sometimes, some nodes refuse to open a libp2p connection with our crawler, for different reasons. Our plots only display _online_ nodes, that are answering the crawler's queries.

### Peer ID mismatch

We've noted that a single Ethereum node may generate multiple Ethereum Node Records (ENRs) that specify the same IP address and port but different libp2p peer IDs. This occurs when the node updates its libp2p peer ID, leading to an outdated ENR alongside a current one. Using any ENR to make a `discv5` request to the remote node should work. It's crucial to use the most recent ENR when initiating discv5 requests or establishing libp2p connections to ensure accurate node representation and avoid duplicate counting of the same node under different peer IDs.

We classify nodes as online if they meet either of the following criteria: they successfully establish a libp2p connection, or they respond positively to discv5 requests but encounter a libp2p connection error that is not a `peer_id_mismatch`.

## Node classification

### User agents

User agent information is sourced from the libp2p identify protocol, which necessitates a successful libp2p connection. Given that libp2p connections can sometimes fail, despite a node's active participation in the DHT, we rely on user agent data acquired from either past or future successful libp2p connections associated with the same node ID. This method ensures that we can accurately capture and utilize user agent information, maintaining consistency in node identification and characterization over time.

### Geographical distribution

We employ MaxMind's GeoIP services to correlate IP addresses with their corresponding countries, utilizing the detailed resources provided by the [MaxMind GeoIP databases](https://www.maxmind.com/en/geoip-databases). This tool allows us to effectively determine the geographic locations of IP addresses, facilitating accurate country-level mapping and analysis.

### Cloud providers

We utilize Udger's datacenter identification services to map IP addresses to their respective commercial cloud providers, leveraging the comprehensive [Udger datacenter list](https://udger.com/resources/datacenter-list). This resource enables us to accurately identify and categorize IP addresses based on their association with known data centers.

## References

[1] [Nebula Crawler](https://github.com/dennis-tra/nebula)

[2] [Kademlia paper](https://www.scs.stanford.edu/~dm/home/papers/kpos.pdf)

[3] [Ethereum Node Discovery Protocol v5](https://github.com/ethereum/devp2p/blob/master/discv5/discv5.md)
