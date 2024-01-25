---
title: Tools & Data
bookCollapseSection: true
weight: 50
---

# Tools

ProbeLab is running multiple sets of experiments that aim to push the boundaries of what is possible in the fields of networking, distributed systems, and measurement of decentralized networks. On this page, we give an overview of the tools and infrastructure that we have developed to monitor live distributed networks. Currently, our efforts focus primarily on IPFS and its underlying, libp2p-based protocol stack.

## Nebula

[`Nebula`](https://github.com/dennis-tra/nebula) is a libp2p DHT crawler and monitor that is designed to track the liveliness and availability of peers. Nebula-based experiments are aimed at monitoring and improving the resilience and reliability of distributed systems by developing better tools for monitoring and managing decentralized peer-to-peer networks.

{{< button relref="/nebula" >}}Learn more{{< /button >}}

## Parsec

[`parsec`](https://github.com/probe-lab/parsec) is a DHT and IPNI performance measurement tool that is used to gather accurate data on the performance of DHT and IPNI lookups and publications. `parsec`-based experiments are aimed at improving the efficiency and speed of distributed systems by developing better algorithms for routing and data retrieval.

{{< button relref="/parsec" >}}Learn more{{< /button >}}

## Tiros

ProbeLab has built a website performance measurement tool, called [`tiros`](https://github.com/probe-lab/tiros) for websites hosted on IPFS. Tiros is designed to help developers monitor and optimize
the performance of their IPFS-hosted websites. `Tiros`-based experiments measure retrieval and rendering metrics of websites loaded over IPFS. It also measures and compares the IPFS metrics with their HTTPS counterparts.

{{< button relref="/tiros" >}}Learn more{{< /button >}}

## Data Provenance

The data produced by the tools above is used to generate the majority of plots on this website. Some external sources of data such as metrics produced by 
IPFS gateway infrastructure are also aggregated and used in their own plots or as a supplement to existing ones. Each plot is defined in an open format that
includes the underlying data sources and queries made to extract the view of the data.

{{< button relref="/data" >}}Learn more{{< /button >}}
