---
title: Gateways
weight: 30
plotly: true
aliases:
    - /ipfsgateways/
---

# IPFS Gateways

An [IPFS gateway](https://docs.ipfs.tech/concepts/ipfs-gateway/) is a web-based interface that allows users to access IPFS content using a regular web browser. The gateway acts as a bridge between the IPFS network and the traditional web, allowing users to browse and retrieve IPFS content without having to install any specialized software or run a full IPFS node.

When a user requests content through an IPFS gateway, the gateway retrieves the content from the IPFS network (through [Bitswap](https://docs.ipfs.tech/concepts/bitswap/#bitswap) and/or [kubo](https://docs.ipfs.tech/install/command-line/#install-ipfs-kubo)) and delivers it to the user's browser using standard web protocols such as HTTP or HTTPS. This allows users to access IPFS content in the same way they would access any other web content, with a simple URL or link. IPFS gateways are a key part of the IPFS ecosystem, as they make IPFS content accessible to a wider audience and help to bridge the gap between the traditional web and the decentralized web.

## HTTP Requests to ipfs.io / dweb.link

IPFS gateways can be run by anyone who has access to an IPFS node, and there are many public gateways available on the internet. The following measurements are made using data derived from the nginx access logs for ipfs.io and dweb.link, the gateway services operated by Protocol Labs.

In the following plots, requests are _not_ deduplicated, i.e., if there are two requests for the same [CID](https://docs.ipfs.tech/concepts/content-addressing/#content-identifiers-cids), they count as two requests and not one. This is what makes sense for requests. Deduplicating requests would effectively count the number of CIDs requested and not the requests themselves.

#### Daily HTTP Requests to ipfs.io / dweb.link

{{< plotly json="../../plots/latest/gateway-requests-overall.json" height="450px" >}}

#### Daily HTTP Requests to ipfs.io / dweb.link, by region

{{< plotly json="../../plots/latest/gateway-requests-region.json" height="450px" >}}

#### Daily HTTP Requests to ipfs.io / dweb.link, by type

{{< plotly json="../../plots/latest/gateway-requests-type.json" height="450px" >}}

#### Daily Unique Clients accessing ipfs.io / dweb.link

In the following plot we aggregate the total number of unique client IP addresses that appear in the ipfs.io and dweb.link log files for each day. 
{{< plotly json="../../plots/latest/gateway-clients-overall.json" height="450px" >}}

#### Daily p95 TTFB accessing ipfs.io / dweb.link

The following graph shows the 95th percentile of the time to first byte (TTFB) for accessing the ipfs.io and dweb.link gateways.

{{< plotly json="../../plots/latest/gateway-ttfb-daily.json" height="450px" >}}

