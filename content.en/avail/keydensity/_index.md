---
title: Avail DHT Key Density
bookHidden: true
bookCollapseSection: true
plotly: true
---

# Amino Key Density Monitoring

In Kademlia, every object indexed by the DHT requires a binary identifier. In the libp2p DHT implementation, peers are identified by the digest of `sha256(peer_id)` and CIDs are identified by the digest of `sha256(cid)`. This Kademlia identifier determines the location of an object within the Kademlia XOR keyspace.

The following plots examine the peer distribution within the keyspace, aiding in the identification of potential [Sybil](https://en.wikipedia.org/wiki/Sybil_attack) and eclipse attacks.

## Keyspace population distribution

**Description**: The plot illustrates the number of peers whose Kademlia identifier matches each prefix for all prefixes of a given size, for a given network crawl. Since the density of keyspace regions follows a [Poisson](https://en.wikipedia.org/wiki/Poisson_distribution) distribution, it is expected to observe some regions that are significantly more populated than others.

**How to read the plot:** The selected `depth` indicates the prefix size. There are `2^i` distinct prefixes at depth `i`. The slider at the bottom of the plot enables visualization of the population evolution over time across multiple crawls.

**What to look out for:** The red dashed line represents the expected density per region, corresponding to the number of peers matching a prefix. A bar exceeding the expected density by more than twice suggests that a region of the keyspace might be under an eclipse attack.

{{< plotly json="../../plots/latest/avail-regions-population.json" height="600px" >}}

## Keyspace density distribution

**Description:** As previously mentioned, the keyspace population follows a [Poisson](https://en.wikipedia.org/wiki/Poisson_distribution) distribution, which can make identifying outliers challenging. The plot below counts the number of regions for each population size and facilitates the identification and analysis of outliers. While it is normal for some regions to have populations above the average, the plot enables us to quantify these deviations.

**How to read the plot:** The red dashed line represents the expected number of regions for each population size. Note that the Poisson distribution is more evident at greater depths (longer prefix size), while analyzing data at lower depths provides limited insights. It is recommended to read the plot for depths between 9 and 13.

**What to look out for:** If a bar significantly exceeds its expected value on the right side of the plot, or if an isolated bar appears on the far right, it may indicate a potential eclipse attack, warranting further investigation.

{{< plotly json="../../plots/latest/avail-density-distributions.json" height="600px" >}}
