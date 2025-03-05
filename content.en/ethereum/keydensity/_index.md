---
title: discv5 Key Density
plotly: true
weight: 90
---

# Ethereum Mainnet discv5 Key Density Monitoring

Every object indexed by discv5 requires a binary identifier. In the discv5 implementation, peers are identified by the digest of `secp256k1(peer_id)` and CIDs are identified by the digest of `keccak256(cid)`. This identifier determines the location of an object within the XOR keyspace.

The following plots examine the peer distribution within the keyspace, aiding in the identification of potential [Sybil](https://en.wikipedia.org/wiki/Sybil_attack) and eclipse attacks.

## Keyspace population distribution

**Description**: The plot illustrates the number of peers whose identifier matches each prefix for all prefixes of a given size, for a given network crawl. Since the density of keyspace regions follows a [Poisson](https://en.wikipedia.org/wiki/Poisson_distribution) distribution, it is expected to observe some regions that are significantly more populated than others.

**How to read the plot:** The selected `depth` indicates the prefix size. There are `2^i` distinct prefixes at depth `i`. The slider at the bottom of the plot enables visualization of the population evolution over time across multiple crawls.

**What to look out for:** The red dashed line represents the expected density per region, corresponding to the number of peers matching a prefix. A bar significantly exceeding the expected density suggests that a region of the keyspace might be under an eclipse attack, especially if the value is above the risk threshold.

{{< plotly json="../../../plots/latest/filecoin-regions-population.json" height="600px" >}}

## Keyspace density distribution

**Description:** As previously mentioned, the keyspace population follows a [Poisson](https://en.wikipedia.org/wiki/Poisson_distribution) distribution, which can make identifying outliers challenging. The plot below counts the number of regions for each population size and facilitates the identification and analysis of outliers. While it is normal for some regions to have populations above the average, the plot enables us to quantify these deviations.

**How to read the plot:** The red dashed line represents the expected number of regions for each population size. Note that the Poisson distribution is more evident at greater depths (longer prefix size), while analyzing data at lower depths provides limited insights. It is recommended to read the plot for depths between 6 and 8.

**What to look out for:** If an isolated bar appears on the far right beyond the risk threshold line, it may indicate a potential eclipse attack, warranting further investigation.

{{< plotly json="../../../plots/latest/filecoin-density-distributions.json" height="600px" >}}
