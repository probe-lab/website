![Nebula Logo](./nebula.png)

# Nebula

Nebula is a [libp2p](https://github.com/libp2p/go-libp2p-kad-dht) DHT crawler
and monitor that tracks the liveliness of peers. The crawler connects
to [DHT](https://en.wikipedia.org/wiki/Distributed_hash_table) bootstrap peers
and then recursively follows all entries in
their [k-buckets](https://en.wikipedia.org/wiki/Kademlia) until all peers have
been visited. The crawler supports
the [IPFS](https://ipfs.network), [Filecoin](https://filecoin.io), [Polkadot](https://polkadot.network/), [Kusama](https://kusama.network/), [Rococo](https://substrate.io/developers/rococo-network/), [Westend](https://wiki.polkadot.network/docs/maintain-networks#westend-test-network)
networks [and more](https://github.com/dennis-tra/nebula/blob/a33a5fd493caaeb07e92ecc73c32ee87ae9e374f/pkg/config/config.go#L11).

{{< button href="https://github.com/dennis-tra/nebula" >}}GitHub{{< /button >}}

## Components

Nebula is split into two components 1) a **crawler** and 2) a **monitor**. The
crawler
is responsible for discovering peers in the DHT and the monitor periodically
connects to the peers to track their uptime.

## How does it work?

The Kademlia DHT network is a distributed system where each peer maintains a
routing table containing other peers in the network with specific XOR distances
to itself. Peers are grouped into "k-buckets" based on the number of shared
prefix bits between their
own [`PeerID`](https://docs.libp2p.io/concepts/peer-id/) and the remote PeerID.
For instance, bucket
0 includes nodes without any shared prefix bits, and each bucket contains a
maximum of 20 entries.

To begin the crawl process, Nebula connects to a set of bootstrap peers and
constructs a routing table for each peer it connects to. Next, it generates
random keys with common prefix lengths that fall into each peer's buckets and
asks remote peers if they know any peers that are closer (in XOR distance) to
the ones Nebula just constructed. This provides a list of all PeerIDs that a
peer has in its routing table. The process repeats for all found peers until no
new PeerIDs are found.

If configured to store results in a database, Nebula persists information for
every visited peer. This information includes latency measurements, current
multi-addresses, agent version, and supported protocols. If the peer is
dialable, Nebula also creates a session instance that contains information about
the node's uptime.

The monitoring process periodically queries the database for these sessions and
tries to dial the comprised peers. If the peer is dialable it updates the
session row with the new uptime and if the peer was not dialable it "closes" the
session. This allows for precise peer churn measurements.

## Installation

### From source

To compile it yourself run:

```shell
go install github.com/dennis-tra/nebula-crawler/cmd/nebula@latest # Go 1.19 or higher is required (may work with a lower version too)
```

Make sure the `$GOPATH/bin` is in your PATH variable to access the installed `nebula` executable.

## Usage

Nebula is a command line tool and provides the `crawl` sub-command. To simply crawl the IPFS network run:

```shell
nebula crawl --dry-run
```

The crawler can store its results as JSON documents or in a postgres database - the `--dry-run` flag prevents it from doing either. Nebula will print a summary of the crawl at the end instead. A crawl takes ~5-10 min depending on your internet connection. You can also specify the network you want to crawl by appending, e.g., `--network FILECOIN` and limit the number of peers to crawl by providing the `--limit` flag with the value of, e.g., `1000`. Example:

```shell
nebula crawl --dry-run --network FILECOIN --limit 1000
```

To store crawl results as JSON files provide the `--json-out` command line flag like so:

```shell
nebula crawl --json-out ./results/
```

After the crawl has finished, you will find the JSON files in the `./results/` subdirectory.

When providing only the `--json-out` command line flag you will see that the `*_neighbors.json` document is empty. This document would contain the full routing table information of each peer in the network which is quite a bit of data (~250MB as of April '23) and is therefore disabled by default. To populate the document you'll need to pass the `--neighbors` flag to the `crawl` subcommand.

```shell
nebula crawl --neighbors --json-out ./results/
```

The routing table information forms a graph and graph visualization tools often operate with [adjacency lists](https://en.wikipedia.org/wiki/Adjacency_list). To convert the `*_neighbors.json` document to an adjacency list, you can use [`jq`](https://stedolan.github.io/jq/) and the following command:

```shell
jq -r '.NeighborIDs[] as $neighbor | [.PeerID, $neighbor] | @csv' ./results/2023-04-16T14:32_neighbors.json > ./results/2023-04-16T14:32_neighbors.csv
```

There are a few more command line flags that are documented when you run `nebula crawl --help`.

When Nebula is configured to store its results in a postgres database (see below), then it also tracks session information of remote peers.


## Contributing

Feel free to head over to the GitHub repository and dive in! [Open an issue](https://github.com/dennis-tra/nebula/issues/new) or submit PRs.

{{< button href="https://github.com/dennis-tra/nebula" >}}GitHub{{< /button >}}
