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
	if e.Status == "start" {
		// TODO: filter for siryn label
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

		if monitor {
			if err := h.startStats(e.Id, c.Config.Image); err != nil {
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

	gaugeCpuTotalUsage.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "cpu",
	}).Set(float64(totalUsage))

	gaugeMemoryUsage.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "memory",
	}).Set(float64(stats.MemoryStats.Usage))

	gaugeMemoryMaxUsage.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "memory",
	}).Set(float64(stats.MemoryStats.MaxUsage))

	gaugeMemoryPercent.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "memory",
	}).Set(float64(memPercent))

	gaugeNetworkRxBytes.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.RxBytes))

	gaugeNetworkRxPackets.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.RxPackets))

	gaugeNetworkRxErrors.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.RxErrors))

	gaugeNetworkRxDropped.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.RxDropped))

	gaugeNetworkTxBytes.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.TxBytes))

	gaugeNetworkTxPackets.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.TxPackets))

	gaugeNetworkTxErrors.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.TxErrors))

	gaugeNetworkTxDropped.With(prometheus.Labels{
		"container": id,
		"image":     image,
		"type":      "network",
	}).Set(float64(stats.NetworkStats.TxDropped))
}
