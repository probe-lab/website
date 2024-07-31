---
title: Methodology
weight: 10
slug: methodology
---

# Methodology for block arrivals reports 

## Basics on Ethereum time deadlines
The Ethereum specification defines strict time deadlines on its consensus layer:
- The spec first defines epochs, which are 6.4 minutes long and are composed of 32 slots of 12 seconds each.
- Although not explicitly defined in the specification, each epoch is further split into three periods:
    - Block composition and its propagation (0 - 4th second, red boxes in the plot below).
    - Attestation composition and its propagation (4th - 8th second, blue boxes in the plot below).
    - Aggregation of attestations and their propagation (8th - 12th second, yellow boxes in the plot below).

![img](../slot_time_division.png)

In theory, nodes compose the block at the slot's starting time, so the propagation starts at`t0`of the slot, giving Gossipsub 4 seconds to propagate over the block over the entire network.
After those 4 seconds, validators must attest to their perception of the chain, including the last block they saw. Ensuring that blocks arrive before the 4-second deadline is crucial to reduce the chances of splitting the network (fork) between nodes that saw a message and nodes that didn't.


## Why it's important to monitor block propagation latency
The defined time windows are not hardcoded anywhere; they are just guidelines to ensure the best performance. However, it's not always in the best interest of the validators to follow them. In fact, it's often the case that validators benefit from delaying message propagation to get extra rewards from the network.

The most well-known case of validators benefiting from artificially delaying the propagation of messages is the "Miner Extractable Value" or [MEV](https://coinmarketcap.com/academy/glossary/miner-extractable-value-mev). 
The strict 12-second window that PoS defines enables validators to focus primarily on producing the most profitable blocks, by including the highest value payloads in the blocks they propose. 
This is done by ordering the available transactions in the most profitable way. In this context, the public bids submitted to [public relays](https://ethstaker.cc/mev-relay-list) can then be selected by validators that have to propose new blocks. 

## How does MEV affect the block proposal and its propagation?

The more time these block builders have, the more time they have to develop a more profitable execution payload. This ultimately incentive validators to "wait" a little bit more beyond the defined `t0` mark, as the MEV reward might be substantially higher due to those extra 1-2 seconds (check out the [live dashboard](https://payload.de/data/) of MEV bids and their reward). 

The most significant limitation in those cases is still the 4-second line, where validators must attest to the last block they saw. Thus, the block proposer can wait longer for a higher rewarding bid without being too greedy and exceed the point where the block propagation couldn't be done by 2/3s of the validators.

## Data source
All the data displayed in the Block Arrival Time reports was collected from a [public database](https://github.com/ethpandaops/xatu-data) maintained by the [EthPandaOps](https://ethpandaops.io/) team of the Ethereum Foundation.

The raw data includes data points from a set of "sentry nodes" located in the following locations:

| Continent | Country   |	Client      |
| ---       | ---       | ---           |
| EU        | FI        | lighthouse    |
|           |           | lodestar      |
|           |           | nimbus        |
|           |           | prysm         |
|           |           | teku          |
| NA        | US        | lighthouse    |
|           |           | lodestar      |
|           |           | nimbus        |
|           |           | prysm         |
|           |           | teku          |
| OC        | AU        | lighthouse    | 
|           |           | lodestar      |
|           |           | nimbus        |
|           |           | prysm         |
|           |           | teku          |  
