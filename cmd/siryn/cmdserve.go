package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ehazlett/siryn/client"
	"github.com/ehazlett/siryn/server"
	"github.com/ehazlett/siryn/version"
)

func waitForInterrupt() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	for _ = range sigChan {
		return
	}
}

var cmdServe = cli.Command{
	Name:   "serve",
	Usage:  "start siryn server",
	Action: serve,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "listen, l",
			Usage: "Listen Address",
			Value: ":8080",
		},
		cli.StringFlag{
			Name:   "docker, d",
			Usage:  "Docker URL",
			Value:  "unix:///var/run/docker.sock",
			EnvVar: "DOCKER_HOST",
		},
		cli.StringFlag{
			Name:  "tls-ca-cert",
			Usage: "TLS CA Certificate",
			Value: "",
		},
		cli.StringFlag{
			Name:  "tls-cert",
			Usage: "TLS Certificate",
			Value: "",
		},
		cli.StringFlag{
			Name:  "tls-key",
			Usage: "TLS Key",
			Value: "",
		},
		cli.BoolFlag{
			Name:  "allow-insecure",
			Usage: "Allow insecure communication to daemon",
		},
		cli.StringFlag{
			Name:  "prometheus-path, p",
			Usage: "Prometheus binary path",
		},
		cli.StringFlag{
			Name:  "prometheus-config, c",
			Usage: "Prometheus config path",
			Value: "prometheus.yml",
		},
		cli.StringFlag{
			Name:  "pushgateway-path, g",
			Usage: "PushGateway binary path",
		},
		cli.IntFlag{
			Name:  "stat-interval, s",
			Usage: "Container stat reporting interval (in seconds < 60)",
			Value: 5,
		},
	},
}

func serve(c *cli.Context) {
	listenAddr := c.String("listen")
	dockerUrl := c.String("docker")
	tlsCACert := c.String("tls-ca-cert")
	tlsCert := c.String("tls-cert")
	tlsKey := c.String("tls-key")
	allowInsecure := c.Bool("allow-insecure")
	prometheusBinaryPath := c.String("prometheus-path")
	prometheusConfigPath := c.String("prometheus-config")
	pushGatewayBinaryPath := c.String("pushgateway-path")
	statInterval := c.Int("stat-interval")

	if prometheusBinaryPath == "" {
		p, err := exec.LookPath("prometheus")
		if err != nil {
			log.Fatal("unable to locate prometheus")
		}
		prometheusBinaryPath = p
	}

	if pushGatewayBinaryPath == "" {
		p, err := exec.LookPath("pushgateway")
		if err != nil {
			log.Fatal("unable to locate pushgateway")
		}
		pushGatewayBinaryPath = p
	}

	log.Infof("starting siryn %s: addr=%s", version.FullVersion, listenAddr)

	d, err := client.GetDockerClient(
		dockerUrl,
		tlsCACert,
		tlsCert,
		tlsKey,
		allowInsecure,
	)

	if err != nil {
		log.Fatal(err)
	}

	// start server
	cfg := &server.SirynConfig{
		DockerClient:          d,
		PrometheusBinaryPath:  prometheusBinaryPath,
		PrometheusConfigPath:  prometheusConfigPath,
		PushGatewayBinaryPath: pushGatewayBinaryPath,
		StatInterval:          statInterval,
		ListenAddr:            listenAddr,
	}

	srv, err := server.NewSirynServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
