# Hermes
Hermes is a light-node for [libp2p](https://libp2p.io/)-based networks that monitors the effectiveness and performance of the [GossipSub](https://docs.libp2p.io/concepts/pubsub/overview/) message propagation protocol. The tool behaves like a light-node in the network, connecting to other participants in the network, relying on a trusted local node that ensures we can reply any incoming request to maintain stable connections. The tool currently supports any [Ethereum](https://ethereum.org/en/) network (the `Consensus Layer`), although there might me more networks coming in the future.

{{< button href="https://github.com/probe-lab/hermes" >}}GitHub{{< /button >}}

## Components
Hermes a single component, the **light-node** for each network. However, it has two major requirements to work: 1) a **trusted local node** that will provide the right view of the chain's state and 2) an **event consumer** that will receive all the libp2p traces and will store them. The light-node is responsible for discovering peers, maintain the connection with them, and to monitor the internal events at the libp2p host.

## How does it work?

### Light-node
Hermes operates like a standard Libp2p node, discovering and connecting to the network while supporting all Libp2p protocols to ensure reliable and secure communication. It connects to a trusted local or remote node for certain RPCs that require chain state information, leveraging these trusted sources to enhance its capabilities.

It subscribes to all available GossipSub topics, allowing it to comprehensively receive, unveil and trace all the interactions with the network. This enables it to keep track of network activity efficiently.

Additionally, Hermes can submit these traces to any defined consumer, with current support for [AWS Kinesis](https://aws.amazon.com/kinesis/) and [Xatu](https://github.com/ethpandaops/xatu) from the Ethereum Foundation (EF). This feature facilitates the integration of network data with analytics and monitoring tools, aiding in the overall security and functionality of the network.

## What data does Hermes gather?
As `Hermes` maintains stable connections with the rest of the network nodes, it can emit debugging traces of how Libp2p and GossipSub work in the "wild". 

To summarize, `Hermes` gathers the following information from the Libp2p host and the GossipSub protocol:

- Connections and Disconnections from the libp2p host

- Control GossipSub messages:

    - Subscriptions from other nodes to topics

    - Graft and Prunes

    - Gossips peerscores 

    - IHaves and IWant messages

- Sent and received messages

- Protoco RPCs

## Deployment
`Hermes` has been developed with the intion of leving it running for an indefinite time duration. If you are interested in the data from other networks, reach out to us through any of the channels listed at the [Contact](/about/#contact) page.

## Contributing
Feel free to head over to the GitHub repository and dive in! [Open an issue](https://github.com/probe-lab/hermes) or submit PRs.

{{< button href="https://github.com/probe-lab/hermes" >}}GitHub{{< /button >}}
