---
title: Methodology
weight: 10
slug: methodology
---

# Methodology for block arrivals reports 

## Basics on Ethereum time deadlines
The Ethereum specification defines strict time deadlines on its consensus layer:
- It first defines epochs, which are the compilation of 32 slots sparsed 12 seconds away
- (Not that directly defined on the specifications) It then defines a second one within the slot itself, which is divided into three sub-time windows for:
    - Block composition and its propagation.
    - Attestation composition and its propagation.
    - Aggregation of attestations and their propagation.

![img](../slot_time_division.png)

In theory, nodes should compose the block at the slot's start starting time, so the propagation starts at`t0`of the slot, giving it 4 seconds to propagate over the network entirely.
After those 4 seconds, validators must attest to their perception of the chain, including the last block they saw. This way, ensuring that blocks arrive before the 4-second deadline is crucial to reduce the chances of splitting the network between nodes that saw the message and nodes that didn't.

However, with the addition of more components or protocol upgrades, the network starts to see side techniques that not only maximize the reward but also play with the protocol at some sort of time game.    

## How the message propagation can be delayed?
The defined time windows are not hardcoded anywhere; they are just guidelines to ensure the best performance. But it's only sometimes in the interest of the validators to follow them. 

Let's put an example with [MEV](https://coinmarketcap.com/academy/glossary/miner-extractable-value-mev). 
The strict 12-second window that PoS defines opens a new space for a new set of users that can focus primarily on producing the most profitable block's payload. 
They do it by ordering the available transactions in the most profitable order. In this context, the public bids submitted to [public relays](https://ethstaker.cc/mev-relay-list) can then be selected by validators that have to propose new blocks. 

How does it affect the block proposal and its propagation?

The more time these block builders have, the more time they have to develop a more profitable execution payload. This ultimately incentive validators to "wait" a little bit more beyond the defined `t0` mark, as the MEV reward might be substantially higher due to those extra 1-2 seconds (check out the [live dashboard](https://payload.de/data/) of MEV bids and their reward). 

The most significant limitation in those cases is still the 4-second line, where validators must attest to the last block they saw. Thus, the block proposer can wait longer for a higher rewarding bid without being too greedy to the point where the block propagation couldn't be done to 2/3s of the validators.

## Data source
All the data displayed in the dashboard was collected from a [public database](https://github.com/ethpandaops/xatu-data) maintained by the [EthPandaOps](https://ethpandaops.io/) team from the Ethereum Foundation.

The raw data includes samples or data points from a set of "sentry nodes" located in the following locations:

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
