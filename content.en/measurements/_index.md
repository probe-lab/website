---
title: Measurements
plotly: true
---

# Measurements

{{< hint warning >}}
Please note that this website is a preview and is subject to change at any time. 
We strive to provide accurate and up-to-date information, but we cannot guarantee 
the completeness or reliability of the information presented during the preview. 
{{< /hint >}}

An overview of a selection of current measurements are included here to give an 


## DHT Server Availability

This measurement shows the number of servers participating in the IPFS Distributed Hash Table (DHT) that were online for at least 80% of the measurement period. IPFS is a dynamic network and peers may enter and leave at any time so the number online at any one time may vary. Read more about how we measure this in the [IPFS DHT](./ipfsdht#availability) section.

{{< plotly json="../plots/dht-availability-classified-online-latest.json" height="250px" >}}

## Website Performance

ProbeLab operates a service that measures the time taken to retrieve pages from a selection of websites served via IPFS. The time to first byte is the time it took to receive the first byte of the first response from each website.

{{< plotly json="../plots/websites-ttfb-quartiles-latest.json" height="250px" >}}

## IPFS DHT

These measurements show the currently observed performance of the IPFS DHT when looking up and publishing provider records. See more in the [IPFS DHT](./ipfsdht#performance) section.

{{< plotly json="../plots/dht-lookup-performance-quartiles-latest.json" height="250px" >}}
{{< plotly json="../plots/dht-publish-performance-quartiles-latest.json" height="250px" >}}


