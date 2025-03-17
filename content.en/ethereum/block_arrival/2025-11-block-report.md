---
title: Week 2025-11
plotly: true
slug: 2025-11
weight: 1045594
---

# Ethereum block broadcasting latency report 2025-11

The following results show measurement data that were collected in **calendar week** 11 **of** 2025 from `2025-03-10` to `2025-03-17`.

This report provides charts and metrics for the beacon block broadcasting latency on the Ethereum network. The methodology we used is available [here](../methodology).

## Block arrival time within the slot
The line plot displays the Cumulative Distribution Function (CDF) of the block arrival times within the slot for the different Ethereum clients (listed [here](../methodology#data-source)). The line plot aggregates all the locations we collect data from for each client.

{{< plotly json="https://cdn.probelab.io/block_arrival_reports/ethereum_mainnet/2025/3/10/block_arrival_times_cdf_by_client_plot.json" height="600px" >}}

The following line plot shows the CDF of the block arrival time within the slot, aggregating all the data points by the continent that the clients were in.

{{< plotly json="https://cdn.probelab.io/block_arrival_reports/ethereum_mainnet/2025/3/10/block_arrival_times_cdf_by_continent_plot.json" height="600px" >}}

## Block arrival time to block size correlation
The following heatmap shows the concentration of block arrival times, aggregating the block size in 50KB range increments. The number displayed within each square represents the corresponding percentage of blocks falling into: i) the block size, and ii) arrival time, out of the total number of blocks seen.

{{< plotly json="https://cdn.probelab.io/block_arrival_reports/ethereum_mainnet/2025/3/10/block_arrival_correlation_with_size_plot.json" height="600px" >}}

## Block arrival time over epochs
The line plot shows the maximum, minimum, mean and median values of the block arrival times aggregated over a range of 4 epochs.

_NOTE: The time aggregation is done over windows of 4 epochs (or 1536 seconds)._

{{< plotly json="https://cdn.probelab.io/block_arrival_reports/ethereum_mainnet/2025/3/10/block_arrival_time_series_percentiles_plot.json" height="600px" >}}

The following line plot shows the time distribution of the mean block arrival time aggregated by client type and over time windows of 4 epochs.

{{< plotly json="https://cdn.probelab.io/block_arrival_reports/ethereum_mainnet/2025/3/10/block_arrival_time_series_clients_plot.json" height="600px" >}}

Lastly, the following line plot displays the time distribution of the mean block arrival times, aggregated by continent and over time windows of 4 epochs.

{{< plotly json="https://cdn.probelab.io/block_arrival_reports/ethereum_mainnet/2025/3/10/block_arrival_time_series_continent_plot.json" height="600px" >}}
