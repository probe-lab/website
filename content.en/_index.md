---
title: Home
type: docs
---

# Introduction

The Protocol Benchmarking & Optimization Team (ProbeLab) is on a mission to measure the performance of Web3.0 network protocols, benchmark protocols against target performance milestones and propose improvements to their core design principles.

We focus on understanding the mechanics of internal network protocols, as well as how they interact with other parts of the system. Our expertise lies in network-layer protocols, and we are particularly active in the [IPFS](https://ipfs.io) and [Filecoin](https://filecoin.io) space, though our work is not limited to that. We dive deep into the protocol as an independent entity and investigate the exogenous factors that influence its performance.

Our team specializes in cross-protocol interoperation and network architecture, works to identify potential bottlenecks and inefficiencies in the system and provide solutions, accordingly. 

## Blog

[Check out our Blog ⇢](/blog)

## Recent Projects

Some of our recent major projects include:

{{< details title="Performance Benefit of Hydra nodes in the IPFS DHT" >}}
**TL;DR:** [Hydra Boosters](https://github.com/libp2p/hydra-booster) are a special type of DHT server node designed to accelerate content routing performance in the IPFS network. They were introduced in 2019 and were intended as an interim solution while exploring other DHT scalability techniques. The IPFS DHT network and its supporting protocols have advanced significantly since then, and the (not insignificant) cost of operating Hydras was put in question by our team. We have found that Hydras improve content routing performance by not more than 10-15% on average, which was considered minor, compared to its operational cost. The team carried out a progressive dial down of Hydras after communicating our intentions with the community (see [details](https://discuss.ipfs.tech/t/dht-hydra-peers-dialling-down-non-bridging-functionality-on-2022-12-01/15567)) and confirmed our performance estimates of a Hydra-less network.

[Read the full report ⇢](https://github.com/protocol/network-measurements/blob/master/results/rfm21-hydras-performance-contribution.md)

{{< /details >}}


{{< details title="libp2p NAT Hole Punching Success Rate" >}}
**TL;DR:** In December 2022, the ProbeLab team conducted a [NAT hole punching measurement campaign](https://discuss.libp2p.io/t/call-for-participation-nat-hole-punching-measurement-campaign/1690/8) to evaluate the NAT hole punching success rate of the [DCUtR protocol](https://github.com/libp2p/specs/blob/master/relay/DCUtR.md). The goal of this campaign was to not only measure the performance of the protocol, but also to identify its limitations and potential areas for optimization. The measurement was designed to provide insights into when and why the DCUtR protocol fails in NAT hole punching and to provide recommendations for improvement. In total, we tracked `6.25M` hole punches from `212` clients (API keys). The clients were deployed in `39` different countries and hole punched remote peers in `167` different countries. Our top findings were that:

- libp2p’s hole punching success rate is around 70%.
- Hole punch success doesn’t depend on the round trip time between the two peers.
- QUIC/UDP is not likelier to succeed than TCP.

[Read the full report ⇢](https://github.com/protocol/network-measurements/blob/master/results/rfm15-nat-hole-punching.md)
{{< /details >}}

{{< details title="The IPFS DHT Routing Table Health" >}}

**TL;DR:** Measuring the performance of decentralized peer-to-peer systems is often not as straightforward as for centrally operated systems. The goal of this study was to dive into the Kademlia DHT Routing Table Health in the IPFS network to understand better the state of the routing table in practice, as well as provide hints on libp2p and IPFS routing improvements.

Our study focused particularly on:

- the percentage of stale entries in the routing tables,
- the peer distribution in Kademila k-buckets,
- the number of missing peers in the k-buckets,
- whether IPFS nodes have their 20 closest peers in their routing tables.

Our measurements show that the Kademlia DHT Routing Table appears to be healthy on all measured aspects in the IPFS network. We showed that `95.21%` of the peers in the IPFS network know at least `18` of their `20` closest peers, which is surprisingly good given the high churn rate observed in the IPFS network. We found that on average 0.12 peers are missing per full k-bucket, and 19.76% per non-full k-bucket, which indicates great performance in terms of keeping routing tables up to date, given that the Kademlia DHT may miss some peers in the k-buckets by design.****

[Read the full report ⇢](https://github.com/protocol/network-measurements/blob/master/results/rfm19-dht-routing-table-health.md)
{{< /details >}}


{{< details title="The IPFS DHT Provider Record Liveness" >}}

**TL;DR:** Provider Records are a key component of the IPFS DHT. They link peers requesting content (CIDs) with peers storing and providing the content. If there are no provider records available in the network, the content can simply not be found (through the DHT). Provider records are stored on (`k=20`) specific peers in the network (called Provider Record Holders, for the purposes of our study), as indicated by the Kademlia routing algorithm. One reason a provider record might disappear from the network is if those peers that are tasked with storing the provider record leave the network. To avoid this situation, content publishers, i.e., those that store and serve content *republish* provider records every `republish.Interval`. 

The purpose of this study has been to identify if at least `1 out of k=20` peers is staying in the network for more than `republish.Interval`, which at the time of the study was 12hrs.

We found that:

1. `75%` of the initial Provider Record Holders stay online for more than `48` hours, showing a healthy current provider record's liveness.
2. `70%` of the PR Holders stay inside the `K=20` closest peers for more than `48` hours.
3. The non-hydra part of the network performs well, showing a healthy non-hydra dependency.
4. From the tested *`K`* values (*`K`*`= 15, 20, 25, 40`):
    - reducing `K` reduces the provider record traffic-related overheard by `25%`, decreases the provider record publication time by `2` seconds, maintains the same `70%` of active Provider Record Holders as `K=20`, and increases hydra-dependency.
    - increasing `K` increases the record publication time, adds extra traffic overhead, but also increases resilience to network fragmentation.

As a result of the study, we suggested increasing the provider record `republish.Interval` to 24 hours as the first action to reduce the overhead, but maintain the same performance.

[Read the full report ⇢](https://github.com/protocol/network-measurements/blob/master/results/rfm17-provider-record-liveness.md)
{{< /details >}}
