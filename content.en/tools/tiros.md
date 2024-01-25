# Tiros

Tiros is a performance measurement tool for IPFS-hosted websites. It orchestrates the interplay between
a Kubo node and a headless Chrome instance. The goal is to measure website
metrics like [TTFB](https://web.dev/ttfb) or [FCP](https://web.dev/fcp) when loaded over IPFS and compare them with their HTTPS performance equivalents. For that, tiros instructs Chrome to request a website through the gateway of a local Kubo node or regularly over HTTPS.

{{< button href="https://github.com/dennis-tra/tiros" >}}GitHub{{< /button >}}

## Concepts

Our Tiros measurement consists of three components

1. `scheduler` - [tiros repository](https://github.com/probe-lab/tiros)
2. `chrome` - [`browserless/chrome`](https://github.com/browserless/chrome)
3. `kubo` - [ipfs/kubo](https://hub.docker.com/r/ipfs/kubo/)

The `scheduler` gets configured with a list of websites that will then be probed. Probing a website means: 1) loading it in headless chrome, and 2) taking [web-vitals](https://web.dev/vitals/) + TTI measurements. The scheduler probes each website by instructing headless chrome to load the following two URLs:

1. `http://localhost:8080/ipns/<website>` - this request hits the local Kubo node's gateway endpoint and loads the website over IPFS

2. `https://<website>` - this request uses regular HTTP to load the website's contents.

This allows us to compare the performance of IPFS against HTTP. The web-vitals measurements consist of the following metrics where the thresholds for `good`, `needs-improvement`, and `poor` come from [web.dev](https://web.dev).

| Metric   | Description                                                                                 |    Good | Needs Improvement |     Poor |
| -------- | ------------------------------------------------------------------------------------------- |--------:|------------------:|---------:|
| **CLS**  | [Cumulative Layout Shift](https://web.dev/cls/)                                             |  < 0.10 |            < 0.25 |  >= 0.25 |
| **FCP**  | [First Contentful Paint](https://web.dev/fcp/)                                              | < 1.80s |           < 3.00s | >= 3.00s |
| **LCP**  | [Largest Contentful Paint](https://web.dev/lcp/)                                            | < 2.50s |           < 4.00s | >= 4.00s |
| **TTFB** | [Time To First Byte](https://web.dev/ttfb/)                                                 | < 0.80s |           < 1.80s | >= 1.80s |
| **TTI**  | [Time To Interactive](https://developer.chrome.com/docs/lighthouse/performance/interactive) | < 3.80s |           < 7.30s | >= 7.30s |

## Deployment

We are running Tiros on ECS as scheduled tasks every six hours in the following
seven AWS regions:

```
eu-central-1
ap-south-1
af-southeast-2
sa-east-1
us-east-2
us-west-1
af-south-1
```

The list of websites that we are currently monitoring is configured [here](https://github.com/probe-lab/probelab-infra/blob/main/aws/tf/modules/tiros/_variables.tf#L49). At the time of writing, the list is:

```
blog.ipfs.tech
blog.libp2p.io
consensuslab.world
docs.ipfs.tech
docs.libp2p.io
drand.love
en.wikipedia-on-ipfs.org/wiki
filecoin.io
green.filecoin.io
ipfs.tech
ipld.io
libp2p.io
probelab.io
protocol.ai
research.protocol.ai
specs.ipfs.tech
strn.network
tldr.filecoin.io
web3.storage
```

Each task run consists of a series of measurements. In each run all the above websites are requested 5 times via Kubo and over HTTP. Then, the task holds for `{settle-time}` minutes (currently 10 minutes [[source](https://github.com/protocol/probelab-infra/blob/1d1867e41cc0a58d641f6edb28ccdf9660f5bdca/aws/tf/tiros.tf#L4)]) and does the same requests again to also take measurements with a "warmed up" kubo node.

In total, each task run measures `settle-times * retries * len([http, kubo]) * len(websites)` website requests. In our case it's `2 * 5 * 2 * 14 = 280` requests. As this task runs in all configured regions, we gather `280 * 7 = 1,960` data points every six hours.

On top of that, the above deployment and configuration is replicated for different Kubo versions. The idea is to measure the performance of the latest Kubo version and the most prevalent one in the network. The Kubo versions that we are measuring are configured [here](https://github.com/protocol/probelab-infra/blob/1d1867e41cc0a58d641f6edb28ccdf9660f5bdca/aws/tf/tiros.tf#L3).

## Contributing

Feel free to head over to the GitHub repository and dive in! [Open an issue](https://github.com/probe-lab/tiros/issues/new) or submit PRs.

{{< button href="https://github.com/probe-lab/tiros" >}}GitHub{{< /button >}}
