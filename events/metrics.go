package events

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	counterCpuTotalUsage = prometheus.NewGaugeVec(
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
	counterMemoryUsage = prometheus.NewGaugeVec(
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
	counterMemoryMaxUsage = prometheus.NewGaugeVec(
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
	counterMemoryPercent = prometheus.NewGaugeVec(
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
	counterNetworkRxBytes = prometheus.NewGaugeVec(
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
	counterNetworkRxPackets = prometheus.NewGaugeVec(
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
	counterNetworkRxErrors = prometheus.NewGaugeVec(
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
	counterNetworkRxDropped = prometheus.NewGaugeVec(
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
	counterNetworkTxBytes = prometheus.NewGaugeVec(
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
	counterNetworkTxPackets = prometheus.NewGaugeVec(
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
	counterNetworkTxErrors = prometheus.NewGaugeVec(
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
	counterNetworkTxDropped = prometheus.NewGaugeVec(
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
	// register the prometheus counters
	prometheus.MustRegister(counterCpuTotalUsage)
	prometheus.MustRegister(counterMemoryUsage)
	prometheus.MustRegister(counterMemoryMaxUsage)
	prometheus.MustRegister(counterMemoryPercent)
	prometheus.MustRegister(counterNetworkRxBytes)
	prometheus.MustRegister(counterNetworkRxPackets)
	prometheus.MustRegister(counterNetworkRxErrors)
	prometheus.MustRegister(counterNetworkRxDropped)
	prometheus.MustRegister(counterNetworkTxBytes)
	prometheus.MustRegister(counterNetworkTxPackets)
	prometheus.MustRegister(counterNetworkTxErrors)
	prometheus.MustRegister(counterNetworkTxDropped)
}
