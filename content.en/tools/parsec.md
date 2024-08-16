---
title: Parsec
weight: 20
---

# Parsec

`parsec` is a Distributed Hash Table (DHT) and InterPlanetary Network Intexer ([IPNI](https://github.com/ipni)) lookup performance measurement tool. It specifically measures the `PUT` and `GET` performance of the **IPFS** public DHT but could also be configured to measure other [libp2p-kad-dht](https://github.com/libp2p/specs/blob/master/kad-dht/README.md) networks. The setup is split into two components 1) a **scheduler** and 2) a **server**.

The server is just a normal libp2p peer that supports and participates in the public IPFS DHT and IPNI gossipsub topics, and exposes a [lean HTTP API](https://github.com/probe-lab/parsec/blob/main/server.yaml) that allows the scheduler to issue publication and retrieval operations to either the DHT or an IPNI. It is important to note that we are _not_ running Kubo. Currently, the scheduler goes around all our deployed server nodes (see [below](#deployment)), instructs one to publish provider records for a random data blob and asks the other six to look it up. All servers take timing measurements about the publication and retrieval latencies and/or report back the results to the scheduler. The scheduler then tracks this information in a database for later analysis.

{{< button href="https://github.com/probe-lab/parsec" >}}GitHub{{< /button >}}

## Concepts

Next to the concept of servers and schedulers there's the concept of a `fleet`. A fleet is a set of server nodes that have a common configuration. For example, we are running four different fleets with seven nodes each (in different regions): 1) `default` 2) `optprov` 3) `fullrt` 4) `cid.contact`. Each of these four fleets is configured differently. The `default` fleet uses the default configuration in the `go-libp2p-kad-dht` repository, the `optprov` fleet uses the optimistic provide configuration to publish data into the DHT, the `fullrt` fleet uses the accelerated DHT client, and the `cid.contact` fleet publishes to and retrieves from the cid.contact network indexer.

Schedulers are then configured to interface with any combination of fleets. Right now, we have one scheduler for each fleet. As said above, it asks one node to publish content, then instructs the others to find the provider records, and then repeats the process with the next peer. However, it is possible to configure a scheduler that does the same thing but with nodes from multiple fleets e.g., `default`+`fullrt` to check if content that's published with one implementation is reachable from another one.

## Deployment

As of 2023-05, we are running `parsec` servers in the following AWS regions:

```
eu-central-1
us-east-2
us-west-1
ap-south-1
ap-southeast-2
af-south-1
sa-east-1
```

Our privileged region is `us-east-1`. That's where our main database resides and
the schedulers run. We currently have four fleets of seven nodes each:

1. `default` - uses the stock `go-libp2p-kad-dht` configuration
2. `optprov` - uses the [Optimistic Provide](https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#optimistic-provide) approach to publish content
3. `fullrt` - uses the [accelerated DHT client](https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#accelerated-dht-client) (full routing table)
4. `cid.contact` - uses the InterPlanetary network indexer [cid.contact](https://cid.contact)

Each fleet has its own scheduler and they don't interfere with each other.

## Contributing

Feel free to head over to the GitHub repository and dive in! [Open an issue](https://github.com/probe-lab/parsec/issues/new) or submit PRs.

{{< button href="https://github.com/probe-lab/parsec" >}}GitHub{{< /button >}}
