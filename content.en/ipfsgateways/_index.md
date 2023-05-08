---
title: IPFS Gateways
weight: 30
plotly: true
---

# IPFS Gateways

{{< hint warning >}}
Please note that this website is a preview and is subject to change at any time. 
We strive to provide accurate and up-to-date information, but we cannot guarantee 
the completeness or reliability of the information presented during the preview. 
{{< /hint >}}


An IPFS gateway is a web-based interface that allows users to access IPFS content using a regular web browser. The gateway acts as a bridge between the IPFS network and the traditional web, allowing users to browse and retrieve IPFS content without having to install any specialized software or run a full IPFS node.

When a user requests content through an IPFS gateway, the gateway retrieves the content from the IPFS network and delivers it to the user's browser using standard web protocols such as HTTP or HTTPS. This allows users to access IPFS content in the same way they would access any other web content, with a simple URL or link. IPFS gateways are a key part of the IPFS ecosystem, as they make IPFS content accessible to a wider audience and help to bridge the gap between the traditional web and the decentralized web.


IPFS gateways can be run by anyone who has access to an IPFS node, and there are many public gateways available on the internet. The following measurements are made using data from the gateway services operated by Protocol Labs.


{{< plotly json="../../plots/latest/gateway-requests-overall.json" height="600px" >}}

{{< plotly json="../../plots/latest/gateway-requests-region.json" height="600px" >}}

{{< plotly json="../../plots/latest/gateway-requests-type.json" height="600px" >}}
