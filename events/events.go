package events

import (
	"fmt"
	"math"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/samalba/dockerclient"
)

type (
	EventHandler struct {
		client       *dockerclient.DockerClient
		statInterval int
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
			if err := h.startStats(e.Id); err != nil {
				ec <- err
				return
			}
		}
	}
}

func (h *EventHandler) startStats(id string) error {
	log.Debugf("gathering stats: id=%s interval=%d", id, h.statInterval)
	go h.handleStats(id, h.sendContainerStats, errorChan, nil)

	return nil
}

func (h *EventHandler) handleStats(id string, cb dockerclient.StatCallback, ec chan error, args ...interface{}) {
	go h.client.StartMonitorStats(id, cb, ec, args)
}

func (h *EventHandler) sendContainerStats(id string, stats *dockerclient.Stats, ec chan error, args ...interface{}) {
	// report on interval
	timestamp := time.Now()

	rem := math.Mod(float64(timestamp.Second()), float64(h.statInterval))
	if rem != 0 {
		return
	}

	totalUsage := stats.CpuStats.CpuUsage.TotalUsage

	log.Debugf("stats: id=%s cpuUsage.total=%d", id, totalUsage)
}
