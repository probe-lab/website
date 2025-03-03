---
title: Ants Watch
weight: 50
---

# Ants Watch

`ants watch` is a DHT client monitoring tool. It is able to log the activity of all nodes in a DHT network by carefully placing ants in the DHT keyspace. For nodes to utilize the DHT they need to perform routing table maintenance tasks. These tasks consist of sending requests to several other nodes close to oneself in the DHT keyspace. `ants watch` ensures that at least one of these requests will always hit one of the deployed ants. When a request hits an ant, we record information about the requesting peer like agent version, supported protocols, IP addresses, and more.

{{< button href="https://github.com/probe-lab/ants-watch" >}}GitHub{{< /button >}}

## How does it work?

* An `ant` is a lightweight [libp2p DHT node](https://github.com/libp2p/go-libp2p-kad-dht), participating in the DHT network, and logging incoming requests.
* `ants` participate in the DHT network as DHT server nodes. `ants` need to be dialable by other nodes in the network. Hence, `ants-watch` must run on a public IP address either with port forwarding properly configured (including local and gateway firewalls) or UPnP enabled.
* The tool releases `ants` (i.e., spawns new `ant` nodes) at targeted locations in the keyspace in order to _occupy_ and _watch_ the full keyspace.
* The tool's logic is based on the fact that peer routing requests are distributed to `k` closest nodes in the keyspace and routing table updates by DHT client (and server) nodes need to find the `k` closest DHT server peers to themselves. Therefore, placing approximately 1 `ant` node every `k` DHT server nodes can capture all DHT client nodes over time.
* The routing table update process varies across implementations, but is by default set to 10 mins in the go-libp2p implementation. This means that `ants` will record the existence of DHT client nodes approximately every 10 mins (or whatever the routing table update interval is).
* Depending on the network size, the number of `ants` as well as their location in the keyspace is adjusted automatically.
* Network size and peers distribution is obtained by querying an external [Nebula database](https://github.com/dennis-tra/nebula).
* All `ants` run from within the same process, sharing the same DHT records.
* The `ant queen` is responsible for spawning, adjusting the number and monitoring the ants as well as gathering their logs and persisting them to a central database.
* `ants-watch` does not operate like a crawler, where after one run the number of DHT client nodes is captured.
* `ants-watch` logs all received DHT requests and therefore, it must run continuously to provide the number of DHT client nodes over time.

## What data does `ants-watch` gather?
`ants watch` can collect information about the requesting peer. Information gathered includes: 
- agent version,
- supported protocols,
- IP addresses (and therefore, geolocation and cloud infrastructure deployment),
- node uptime.

## Setup & Deployment
For detailed information on how to setup and deploy `ants`, please refer to the tool's Github repository: https://github.com/probe-lab/ants-watch/blob/dev/README.md#setup.

## Contributing
Feel free to head over to the GitHub repository and dive in! [Open an issue](https://github.com/probe-lab/ants-watch) or submit PRs.

{{< button href="https://github.com/probe-lab/ants-watch" >}}GitHub{{< /button >}}
