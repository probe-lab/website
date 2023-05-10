# Parsec

`parsec` is a DHT lookup performance measurement tool. It specifically measures the `PUT` and `GET` performance of the
IPFS public DHT but could also be configured to measure
other [libp2p-kad-dht](https://github.com/libp2p/specs/blob/master/kad-dht/README.md) networks.
The setup is split into two components: a scheduler and a server.

The server is just a normal libp2p peer that supports and participates in the public IPFS DHT and exposes a [lean HTTP
API](https://github.com/plprobelab/parsec/blob/main/server.yaml) that allows the scheduler to issue publication and retrieval operations. Currently, the scheduler goes around all
seven server nodes, instructs one to publish provider records for a random data blob and asks the other six to look them up.
All seven servers take timing measurements about the publication or retrieval latencies and
report back the results to the scheduler. The scheduler then tracks this information in a database for later analysis.

## Concepts

Next to the concept of servers and schedulers there's the concept of a `fleet`. A fleet is a set of server nodes that
have a common configuration. For example, we are running three different fleets with seven nodes each (in different regions): 1) `default` 2) `optprov` 3) `fullrt`.
Each of these three fleets are configured differently. The `default` fleet uses the default configuration in the `go-libp2p-kad-dht` repository, the `optprov` fleet uses the optimistic provide configuration to publish data into the DHT, and the `fullrt` fleet uses the accelerated DHT client.

Schedulers are then configured to interface with any combination of fleets. Right now, we have one scheduler for each fleet. As said above, it asks one node to publish content, then instructs the others to find the provider records, and then repeats the process with the next peer. However,
we could configure a scheduler that does the same thing but with nodes from multiple fleets e.g., `default`+`fullrt` to check if content that's published with one implementation is reachable with another one.

## Deployment

We are running `parsec` servers in the following AWS regions:

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
the schedulers run. We currently have three fleets of seven nodes each:

1. `default` - uses the stock `go-libp2p-kad-dht` configuration
2. `optprov` - uses the Optimistic Provide approach to publish content
3. `fullrt` - uses the accelerated DHT client (full routing table)

Each fleet has its own scheduler and they don't interfere with each other.


## Contributing

Feel free to head over to the GitHub repository and dive in! [Open an issue](https://github.com/plprobelab/parsec/issues/new) or submit PRs.

{{< button href="https://github.com/plprobelab/parsec" >}}GitHub{{< /button >}}
