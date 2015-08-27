package events

import (
	"fmt"
	"math"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/samalba/dockerclient"
)

type (
	EventHandler struct {
		client       *dockerclient.DockerClient
		statInterval int
	}

	eventArgs struct {
		Image string
	}
)

var (
	errorChan = make(chan error)
)

func NewEventHandler(client *dockerclient.DockerClient, statInterval int) *EventHandler {
	return &EventHandler{
		client:       client,
		statInterval: statInterval,
	}
}

// Handle starts container stats monitoring if label matches
func (h *EventHandler) Handle(e *dockerclient.Event, ec chan error, args ...interface{}) {
	log.Debug(fmt.Sprintf("event: date=%d type=%s image=%s container=%s", e.Time, e.Status, e.From, e.Id))
	// get container info for event
	c, err := h.client.InspectContainer(e.Id)
	if err != nil {
		ec <- err
		return
	}

	monitor := false

	for k, _ := range c.Config.Labels {
		if strings.Index(k, "siryn") == 0 {
			monitor = true
			break
		}
	}

	id := e.Id
	image := c.Config.Image

	if monitor {
		switch e.Status {
		case "start":
			log.Debugf("starting stats: id=%s image=%s", id, image)
			if err := h.startStats(id, image); err != nil {
				ec <- err
				return
			}
		case "kill", "die", "stop", "destroy":
			log.Debugf("resetting stats: id=%s image=%s", id, image)
			if err := h.resetStats(id, image); err != nil {
				ec <- err
				return
			}
		}
	}
}

func (h *EventHandler) startStats(id string, image string) error {
	log.Debugf("gathering stats: id=%s image=%s interval=%d", id, image, h.statInterval)
	args := eventArgs{
		Image: image,
	}
	go h.handleStats(id, h.sendContainerStats, errorChan, args)

	return nil
}

func (h *EventHandler) handleStats(id string, cb dockerclient.StatCallback, ec chan error, args ...interface{}) {
	go h.client.StartMonitorStats(id, cb, ec, args...)
}

func (h *EventHandler) sendContainerStats(id string, stats *dockerclient.Stats, ec chan error, args ...interface{}) {
	// report on interval
	timestamp := time.Now()

	rem := math.Mod(float64(timestamp.Second()), float64(h.statInterval))
	if rem != 0 {
		return
	}

	image := ""
	if len(args) > 0 {
		arg := args[0]
		evtArgs := arg.(eventArgs)
		image = evtArgs.Image
	}

	totalUsage := stats.CpuStats.CpuUsage.TotalUsage
	memPercent := float64(stats.MemoryStats.Usage) / float64(stats.MemoryStats.Limit) * 100.0

	counterCpuTotalUsage.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "cpu",
	}).Set(float64(totalUsage))

	counterMemoryUsage.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "memory",
	}).Set(float64(stats.MemoryStats.Usage))

	counterMemoryMaxUsage.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "memory",
	}).Set(float64(stats.MemoryStats.MaxUsage))

	counterMemoryPercent.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "memory",
	}).Set(float64(memPercent))

	counterNetworkRxBytes.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.RxBytes))

	counterNetworkRxPackets.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.RxPackets))

	counterNetworkRxErrors.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.RxErrors))

	counterNetworkRxDropped.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.RxDropped))

	counterNetworkTxBytes.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.TxBytes))

	counterNetworkTxPackets.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.TxPackets))

	counterNetworkTxErrors.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.TxErrors))

	counterNetworkTxDropped.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.TxDropped))
}

func (h *EventHandler) resetStats(id, image string) error {
	counterCpuTotalUsage.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "cpu",
	}).Set(float64(0.0))

	counterMemoryUsage.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "memory",
	}).Set(float64(0.0))

	counterMemoryMaxUsage.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "memory",
	}).Set(float64(0.0))

	counterMemoryPercent.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "memory",
	}).Set(float64(0.0))

	counterNetworkRxBytes.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(0.0))

	counterNetworkRxPackets.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(0.0))

	counterNetworkRxErrors.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(0.0))

	counterNetworkRxDropped.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(0.0))

	counterNetworkTxBytes.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(0.0))

	counterNetworkTxPackets.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(0.0))

	counterNetworkTxErrors.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(0.0))

	counterNetworkTxDropped.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(0.0))

	return nil
}
