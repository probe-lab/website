---
title: Block Arrival Times
plotly: True
---

# Ethereum block broadcasting latency  

The following report represents a compilation of charts and metrics introducing the measured beacon block arrival times on the Ethereum network. The methodology will be available [here](../methodology) for further details.    

{{< hint info >}}
💡 A more detailed report is available on [this ethresear.ch](https://ethresear.ch/t/gossipsub-message-propagation-latency/19982) post.
{{< /hint >}}

_NOTE: dates of the data 14th to 16th of June (both included)_

## Block arrival time within the slot 
The line plot displays the Cumulative Distribution Function (CDF) of the block arrival times within the slot at the different clients (listed [here](../methodology#data-source)), showing each percentile per client type aggregating all the locations.

{{< plotly json="../../../plots/latest/ethereum-block-broadcasting-latency/gossipsub_arrival_times_within_slot_by_agent_on_mainnet_beacon_block.json" height="600px" >}}

The following line plot shows the CDF of the block arrival time within the slot, aggregating all the data points by the continent that the clients were in.

{{< plotly json="../../../plots/latest/ethereum-block-broadcasting-latency/gossipsub_arrival_times_within_slot_on_by_continent_mainnet_beacon_block.json" height="600px" >}}

## Block arrival time to block size correlation
The following heatmap shows the concentration block arrival times, aggregating the block size by a 50KB range. The displayed numbers correspond to the represented percentage of blocks falling into the block size and arrival time out of the total of blocks.  

{{< plotly json="../../../plots/latest/ethereum-block-broadcasting-latency/message_arrival_time_and_size_correlation_heatmap_on_mainnet_beacon_block.json" height="600px" >}}

## Block arrival time over epochs
The line plot shows the maximum, minimum, mean and median values of the block arrival times aggregated over ranges of 4 epochs.   

_NOTE: The time aggregations are done over windows of 4 epochs (or 1536 seconds)_

{{< plotly json="../../../plots/latest/ethereum-block-broadcasting-latency/message_arrivals_max_min_on_1536s_window_on_topic_mainnet_beacon_block.json" height="600px" >}}

The following line plot shows a time distribution of the mean block arrival time aggregated by client type and over time windows of 4 epochs.

{{< plotly json="../../../plots/latest/ethereum-block-broadcasting-latency/message_arrival_time_on_1536s_window_by_agents_on_topic_mainnet_beacon_block.json" height="600px" >}}

Lastly, the line plot displays the time distribution of the mean block arrival times aggregated by continent and over time windows of 4 epochs.

{{< plotly json="../../../plots/latest/ethereum-block-broadcasting-latency/message_arrival_time_on_1536s_window_by_continent_on_topic_mainnet_beacon_block.json" height="600px" >}}





