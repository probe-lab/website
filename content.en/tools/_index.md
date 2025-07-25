---
title: Tools & Data
bookCollapseSection: true
weight: 60
---

# Tools

ProbeLab is running multiple sets of experiments that aim to push the boundaries of what is possible in the fields of networking, distributed systems, and measurement of decentralized networks. On this page, we give an overview of the tools and infrastructure that we have developed to monitor live distributed networks. Our efforts focus primarily on libp2p-based networks, but since 2024-Q2 we are expanding to other protocol stacks.

## Nebula

[`Nebula`](https://github.com/dennis-tra/nebula) is a DHT crawler and monitor that is designed to track the liveliness and availability of peers. Nebula-based experiments are aimed at monitoring and improving the resilience and reliability of distributed systems by developing better tools for monitoring and managing decentralized peer-to-peer networks. Nebula currently supports libp2p-, discv4- and discv5-based networks.

{{< button relref="/nebula" >}}Learn more{{< /button >}}

## Hermes

[Hermes](https://github.com/probe-lab/hermes) is a [GossipSub](https://docs.libp2p.io/concepts/pubsub/overview/) listener and tracer for [libp2p](https://libp2p.io/)-based networks. `Hermes`-based experiments aim to measure the efficiency and performance of the GossipSub message broadcasting protocol in any `libp2p`-based network. `Hermes` can help developers tune their network's protocols based on the message propagation latency and control message overhead. `Hermes` currently supports the Ethereum Consensus Layer network and Filecoin network.

{{< button relref="/hermes" >}}Learn more{{< /button >}}

## Ants

[ants-watch](https://github.com/probe-lab/ants-watch) is a DHT client monitoring tool. It is able to log the activity of all nodes in a DHT network by carefully placing ants in the DHT keyspace. For nodes to utilize the DHT they need to perform routing table maintenance tasks. These tasks consist of sending requests to several other nodes close to oneself in the DHT keyspace. ants watch ensures that at least one of these requests will always hit one of the deployed ants. When a request hits an ant, we record information about the requesting peer like agent version, supported protocols, IP addresses, and more.  

{{< button relref="/ants" >}}Learn more{{< /button >}}

## Akai 

[akai](https://github.com/probe-lab/akai) is a libp2p generic data availability (DA) sampler. It was originaly developed to perform Data Availability Sampling (DAS) of DA Cells in the Avail DA DHT. However, it has been extended to support more DHT networks. As of 2025-Q3, Akai supports sampling items in the Avail network (Mainnet, Turing, Hex), Celestia network (Mainnet, Mocha-4), and IPFS (Amino DHT).

{{< button ref="[/ants](https://github.com/probe-lab/akai)" >}}Learn more{{< /button >}}

## Parsec

[`parsec`](https://github.com/probe-lab/parsec) is a DHT and [IPNI](https://cid.contact/) performance measurement tool that is used to gather accurate data on the performance of lookups and publications. `parsec`-based experiments are aimed at improving the efficiency and speed of distributed systems by developing better algorithms for routing and data retrieval.

{{< button relref="/parsec" >}}Learn more{{< /button >}}

## Tiros

ProbeLab has built a website performance measurement tool, called [`tiros`](https://github.com/probe-lab/tiros) for websites hosted on IPFS. Tiros is designed to help developers monitor and optimize the performance of their IPFS-hosted websites. `Tiros`-based experiments measure retrieval and rendering metrics of websites loaded over IPFS. It also measures and compares the IPFS metrics with their HTTPS counterparts.

{{< button relref="/tiros" >}}Learn more{{< /button >}}

## Data Provenance

The data produced by the tools above is used to generate the majority of plots on this website. Some external sources of data such as metrics produced by 
IPFS gateway infrastructure are also aggregated and used in their own plots or as a supplement to existing ones. Each plot is defined in an open format that includes the underlying data sources and queries made to extract the view of the data.

{{< button relref="/data" >}}Learn more{{< /button >}}
