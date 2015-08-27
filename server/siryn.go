package server

import (
	"fmt"
	"net/http"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/ehazlett/siryn/events"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/samalba/dockerclient"
)

var (
	prometheusArgs  = []string{}
	pushGatewayArgs = []string{}
)

type SirynConfig struct {
	PrometheusBinaryPath  string
	PrometheusConfigPath  string
	PushGatewayBinaryPath string
	StatInterval          int
	DockerClient          *dockerclient.DockerClient
	ListenAddr            string
}

type SirynServer struct {
	config         *SirynConfig
	prometheusCmd  *exec.Cmd
	pushGatewayCmd *exec.Cmd
}

func NewSirynServer(cfg *SirynConfig) (*SirynServer, error) {
	return &SirynServer{
		config: cfg,
	}, nil
}

func (s *SirynServer) Run() error {
	if err := s.StartPrometheus(); err != nil {
		return err
	}

	if err := s.StartPushGateway(); err != nil {
		return err
	}

	if err := s.StartEventHandler(); err != nil {
		return err
	}

	// start prometheus listener
	http.Handle("/metrics", prometheus.Handler())

	if err := http.ListenAndServe(s.config.ListenAddr, nil); err != nil {
		return err
	}

	return nil
}

func (s *SirynServer) StartPrometheus() error {
	log.Info("starting prometheus")
	args := append(prometheusArgs, fmt.Sprintf("-config.file=%s",
		s.config.PrometheusConfigPath))
	cmd, err := startPrometheus(s.config.PrometheusBinaryPath, args)
	if err != nil {
		return err
	}

	s.prometheusCmd = cmd

	return nil
}

func (s *SirynServer) StartPushGateway() error {
	log.Info("starting pushgateway")
	cmd, err := startPushGateway(s.config.PushGatewayBinaryPath, pushGatewayArgs)
	if err != nil {
		return err
	}

	s.pushGatewayCmd = cmd

	return nil
}

func (s *SirynServer) StartEventHandler() error {
	errChan := make(chan (error))
	go func() {
		err := <-errChan
		log.Error(err)
	}()

	// event handler
	h := events.NewEventHandler(s.config.DockerClient, s.config.StatInterval)
	s.config.DockerClient.StartMonitorEvents(h.Handle, errChan)

	return nil
}
