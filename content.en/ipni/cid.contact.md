---
title: cid.contact
plotly: true
---

# `cid.contact`

This page shows performance metrics for the [cid.contact](https://cid.contact)
InterPlanetary Network Indexer. We gathered the data with our
[parsec](/tools/parsec) measurement tool. More information about the methodology
[below](#methodology).

## Trends

{{< plotly json="../../plots/latest/ipni-trend-latencies-cid.contact.json" height="300px" >}}

The graph shows different percentiles of the "Time To First Provider Record" (TTFPR)
metrics over time. To measure the TTFPR, we start the stopwatch right before we
do the `https://{ipni}/cid/{cid}` HTTP request to the indexer and stop it when we have received and parsed
the full record.

In the above graph we aggregate these latency measurements on a daily basis and
calculate the 50th and 90th percentiles incorporating data from multiple
regions. In our measurement setup, we do multiple requests for the same CID. To
differentiate between cached and uncached latencies, we show the latencies of the
first requests for a CID as "uncached" and subsequent request latencies as
"cached".

It is important to note that the resulting TTFPR percentiles are artificial composites
and do not reflect the specific performance of any individual region. Rather,
it allows to discern general tendencies and fluctuations in TTFPR across the
combined dataset. By focusing on these overall trends, we can gain valuable
insights into the performance of IPNI TTFPR performance as a whole, making it easier to
identify any notable deviations or improvements in that metric over time.

The sample size corresponds to the number of cached retrievals at a specific day
summed across all considered regions. The number of uncached retrievals is usually
exactly the same and therefore not shown separately.

## Lookup Latencies

The graphs presented below display the latencies for the "Time To First Provider
Record" (TTFPR) observed from various regions over the past few days. Each line
represents the Cumulative Distribution Function (CDF) of the lookup latency in a
specific region. Note that the x-axis is presented in a
logarithmic scale. We distinguish between two types of latencies: uncached,
which refers to the first request for a CID, and cached, which pertains to
subsequent requests for a CID.

### Uncached

{{< plotly json="../../plots/latest/ipni-snapshot-uncached-latencies-cdf-cid.contact.json" height="300px" >}}

### Cached

{{< plotly json="../../plots/latest/ipni-snapshot-cached-latencies-cdf-cid.contact.json" height="300px" >}}

### DHT Comparison

{{< plotly json="../../plots/latest/ipni-snapshot-dht-comparison-cid.contact-p90.json" height="300px" >}}

The bar chart above illustrates the 90th percentile of latency associated with
looking up a provider record using two different methods: the Distributed Hash
Table (DHT) and this network indexer.

The network indexer's lookup latencies are divided into two categories: uncached and cached latencies. 
When a CID is looked up for the first time, the response is cached at the edge. The "uncached" latency
bars indicate the latencies for initial CID lookups, while the "cached" latency bars represent the latencies for subsequent lookups.

### Methodology

#### Setup

We are using our [parsec](/tools/parsec) measurement tool to gather publication
and retrieval latency information from an IPNI. `parsec` is a Distributed Hash Table (DHT) and InterPlanetary Network Intexer (IPNI) lookup performance measurement tool and split into two components 1) a **scheduler** and 2) a **server**.

The server is a lean libp2p peer that supports and participates in IPNI gossipsub topics. It exposes a lean HTTP API that allows the scheduler to issue publication and retrieval operations. We are running seven servers in [different AWS regions](/tools/parsec/#deployment).

#### Operation

The _scheduler_ node in our privileged `us-east-1` region instructs one of our seven _server_ nodes to publish a random CID to an `cid.contact`. When the server receives such request, it goes ahead and _announces_ its possession of the CID/content to `cid.contact`.
Announcing the possession of a CID/content can be done via HTTP and/or GossipSub and involves the generation of an advertisement that has its own CID (AdCID). For now, we're only using the HTTP method and making a request to the `/announce` endpoint of `cid.contact`. In our case, the advertisement actually will include 300 CIDs 1) the original one that was received from the scheduler and 2) random "probe" CIDs, which will become relevant later.

At this point, the CID was not indexed yet, so the server waits until `cid.contact` reaches out to it to sync the advertisement chain. To detect that moment, we started watching GraphSync events for the AdCID right when we called the `/announce` endpoint. After `cid.c ontact` reached out to the server, and we received a GraphSync completion event for the AdCID, we can be relatively certain that the CID was indexed. However, there might still be some internal IPNI delay until it is fully indexed. Therefore, the server requests the "probe" CIDs to check if the advertisement was indeed indexed with a timeout of five minutes every second [1] [2]. If the request for that CID returns a record, we know the advertisement was indexed, and the server gives control back to the scheduler, which then starts instructing the remaining servers to request the original CID.

Now, the scheduler instructs the other servers to retrieve the CID from the indexer by calling the `/cid/{cid}` endpoint. Every time a retrieval request resolves, the scheduler instructs that server to retrieve the CID again (up to five times). Each server measures the latency of individual requests. The stopwatch starts right before the HTTP request starts and stops right after we received and parsed the response. Each of the above request latencies becomes an individual data point in our database.

#### Classification
##### Uncached

We are classifying _uncached_ as the "first request from **any** region." The assumption is that the request from one region doesn't affect the cache in another region. Because `cid.contact` uses [CloudFront with its >400 points of presence](https://docs.aws.amazon.com/whitepapers/latest/aws-fault-isolation-boundaries/points-of-presence.html) as it's caching layer it is likely that this assumption holds true.

##### Cached

We are classifying _cached_ requests as only "request number 2 from any region". This means we are currently discarding all data points >2.

#### References

Additional context in [this GitHub issue](https://github.com/probe-lab/network-measurements/issues/47).

---

**Footnotes**

[1] If the request for the probe CID returns a 404 (aka the advertisement was not indexed at this point), the negative response is cached for 5 minutes, so even if the indexing was done just a second later, the server would only find out after 5 minutes. Therefore, we are generating 300 CIDs and check them every second.

[2] The server cannot just check if the latest advertisement for the provider as returned by `/providers/{serverPeerID}` matches the AdCID because the `/providers/{serverPeerID}` response is cached for one hour. So a subsequent round of publications won't see the updated record.