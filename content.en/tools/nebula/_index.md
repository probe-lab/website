![Nebula Logo](./nebula.png)

# Nebula

Nebula is a [libp2p](https://github.com/libp2p/go-libp2p-kad-dht) DHT crawler and monitor that tracks the liveliness of peers. The crawler connects to [DHT](https://en.wikipedia.org/wiki/Distributed_hash_table) bootstrap peers and then recursively follows all entries in their [k-buckets](https://en.wikipedia.org/wiki/Kademlia) until all peers have been visited. The crawler supports the [IPFS](https://ipfs.network), [Filecoin](https://filecoin.io), [Polkadot](https://polkadot.network/), [Kusama](https://kusama.network/), [Rococo](https://substrate.io/developers/rococo-network/), [Westend](https://wiki.polkadot.network/docs/maintain-networks#westend-test-network) networks [and more](https://github.com/dennis-tra/nebula/blob/a33a5fd493caaeb07e92ecc73c32ee87ae9e374f/pkg/config/config.go#L11).

{{< button href="https://github.com/dennis-tra/nebula" >}}GitHub{{< /button >}}

## Components

Nebula is split into two components 1) a **crawler** and 2) a **monitor**. The crawler is responsible for discovering peers in the DHT and the monitor is responsible for tracking their uptime.

## How does it work?

### Crawler

The Kademlia DHT network is a distributed system where each peer maintains a routing table containing other peers in the network with specific [XOR distances](https://en.wikipedia.org/wiki/Kademlia#Accelerated_lookups) to itself. These peers are grouped into so-called _k_-buckets based on the number of leading zeroes of the XOR between other's PeerIDs and the own PeerID. Also called shared prefix bits or common prefix length. For example, bucket 0 includes nodes without any shared prefix bits, and bucket 3 contains peers where the first three bits of both PeerIDs match. Each bucket contains a maximum of 20 entries.

To begin the crawl process, Nebula connects to a configurable set of bootstrap peers and successively follows every peer in their routing tables until it doesn't find any new peer. Given that Nebula knows the other peers' PeerIDs, it also knows about the _k_-buckets they maintain. In order to gather information from the routing tables of other peers, Nebula generates random keys that have a certain number of leading zeroes. These keys fall into each of the buckets that the other peers maintain. Nebula sends these keys to the other peers and asks if they know any peers that are closer to those keys. In response, the other peers provide Nebula with the closest peers they know of, which are all the peers in their respective buckets. Nebula performs this request in parallel for buckets numbered 0 to 15 [[source](https://github.com/dennis-tra/nebula/blob/a20481d35509411b03b3fcfcf0f148b49a5eda3f/pkg/crawl/crawl_p2p.go#L131)].

When a new peer is discovered, the crawler records the start of a session of availability and extends the session length with every subsequent successful connection attempt. In the opposite case, a failed connection terminates the session, and a later successful attempt starts a new session. Depending on the error that the connection attempt returns, Nebula will retry to connect or immediately mark the peer as offline. For example, if the other peer responds with an error that indicates that their resource limit is exceeded, then Nebula retries to connect another two times after five and ten seconds.

### Monitor

The monitoring process periodically queries the database for peers that Nebula considers to be online and tries to connect to them. That way we try to gather a precise measurement of the uptime. If the peer is dialable Nebula updates the session with the new uptime and if the peer was not dialable it "closes" the session. The peer is now considered offline. This allows for precise peer churn measurements.

The longer a peer is seen online, the lower the frequency of connection attempts that is made by Nebula to this particular peer. This is based on the assumption that if a peer remains online for a while, it is more likely to continue being online. Conversely, if Nebula discovers a new peer in the network, there is a high probability that it will go offline (churn) relatively quickly.

Nebula calculates when the next probe is due with the following formula

```
NOW + floor(max_interval, ceil(min_interval, 1.2 * (NOW - previous_successful_dial))
```

The maximum interval is set to 15m and the minimum interval is set to 1m.

The _monitoring_ process doesn't establish a proper libp2p connection (which involves the protocol negotiation) but only dials the peer on a transport level by calling [`DialPeer`](https://github.com/libp2p/go-libp2p/blob/8d771355b41297623e05b04a865d029a2522a074/p2p/net/swarm/swarm_dial.go#L229). It also liberally retries dialing peers if errors occur [[source](https://github.com/dennis-tra/nebula/blob/a20481d35509411b03b3fcfcf0f148b49a5eda3f/pkg/monitor/dialer.go#L117)].

## What data does Nebula gather?

The _crawler_ component establishes a proper libp2p connection to the remote peer. This means that Nebula and the remote peer exchange the list of supported protocols and user agent information. Furthermore, in order to connect to the remote peer, Nebula must have knowledge of the network addresses of the remote peer. It also measures the latency to dial, connect, and crawl. The dial latency includes only the establishment of a connection on the transport level and the connect latency also includes the protocol handshake. The crawl latency measures how long it took to extract information from other peers _k_-buckets.

To summarize, Nebula gathers the following information about all peers it was able to connect to:

- peer ID

- user agent

- supported protocols

- all advertised Multiaddresses

- connection latencies (dial, connect, crawl - see above)

- potential connection errors

The monitoring component periodically checks if peers are still online. This allows us to additionally measure sessions of uptime for each peer.

Because we crawl and probe the network periodically, and because Multiaddresses contain IP addresses we can also answer the following questions:

- When do peers update their IPFS node?

- How long has a peer been online?

- In which country/city is a peer located? (powered by [Maxmind](https://www.maxmind.com/en/home))

- Does the peer run in a data center? (powered by [Udger](https://udger.com/))

On top of the above, Nebula also tracks _neighbor_ information. We consider peers in _k_-buckets to be neighbors of the peer who maintains these _k_-buckets. This information spans a graph where each node is a peer and each edge corresponds to a _k_-bucket entry.

## Deployment

We are running the crawler to measure the following networks:

| Network  | Crawl Frequency  | Crawl duration |
| -------- | ---------------- | -------------- |
| IPFS     | every 30 minutes | ~5 minutes     |
| Filecoin | every 30 minutes | ~1 minute      |
| Polkadot | every 30 minutes | ~25 minutes    |
| Kusama   | every 30 minutes | ~9 minutes     |
| Westend  | every 30 minutes | ~2 minutes     |
| Rococo   | every 30 minutes | ~2 minutes     |

On this website we currently only report on the IPFS network. If you are interested in the data from the rest of the networks, reach out to us through any of the channels listed at the [Contact](https://probelab.io/contact/) page.

## Contributing

Feel free to head over to the GitHub repository and dive in! [Open an issue](https://github.com/dennis-tra/nebula/issues/new) or submit PRs.

{{< button href="https://github.com/dennis-tra/nebula" >}}GitHub{{< /button >}}
