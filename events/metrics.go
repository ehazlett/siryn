package events

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	gaugeCpuTotalUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "cpu_total_time_nanoseconds",
			Help:      "Total CPU time used in nanoseconds",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeMemoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "memory_usage_bytes",
			Help:      "Memory used in bytes",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeMemoryMaxUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "memory_max_usage_bytes",
			Help:      "Memory max used in bytes",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeMemoryPercent = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "memory_usage_percent",
			Help:      "Percentage of memory used",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeNetworkRxBytes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "network_rx_bytes",
			Help:      "Network (rx) in bytes",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeNetworkRxPackets = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "network_rx_packets_total",
			Help:      "Network (rx) packet total",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeNetworkRxErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "network_rx_errors_total",
			Help:      "Network (rx) error total",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeNetworkRxDropped = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "network_rx_dropped_total",
			Help:      "Network (rx) dropped total",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeNetworkTxBytes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "network_tx_bytes",
			Help:      "Network (tx) in bytes",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeNetworkTxPackets = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "network_tx_packets_total",
			Help:      "Network (tx) packet total",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeNetworkTxErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "network_tx_errors_total",
			Help:      "Network (tx) error total",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
	gaugeNetworkTxDropped = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "siryn",
			Subsystem: "docker",
			Name:      "network_tx_dropped_total",
			Help:      "Network (tx) dropped total",
		},
		[]string{
			"container",
			"image",
			"type",
		},
	)
)

func init() {
	// register the prometheus gauges
	prometheus.MustRegister(gaugeCpuTotalUsage)
	prometheus.MustRegister(gaugeMemoryUsage)
	prometheus.MustRegister(gaugeMemoryMaxUsage)
	prometheus.MustRegister(gaugeMemoryPercent)
	prometheus.MustRegister(gaugeNetworkRxBytes)
	prometheus.MustRegister(gaugeNetworkRxPackets)
	prometheus.MustRegister(gaugeNetworkRxErrors)
	prometheus.MustRegister(gaugeNetworkRxDropped)
	prometheus.MustRegister(gaugeNetworkTxBytes)
	prometheus.MustRegister(gaugeNetworkTxPackets)
	prometheus.MustRegister(gaugeNetworkTxErrors)
	prometheus.MustRegister(gaugeNetworkTxDropped)
}
