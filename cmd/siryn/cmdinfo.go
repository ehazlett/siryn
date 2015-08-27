package main

import (
	"fmt"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ehazlett/siryn/version"
)

var cmdInfo = cli.Command{
	Name:   "info",
	Usage:  "show info",
	Action: info,
}

func info(c *cli.Context) {
	sirynVersion := fmt.Sprintf("siryn %s", version.FullVersion)
	promVersion := "unknown"

	promPath, err := exec.LookPath("prometheus")
	if err != nil {
		log.Fatal("unable to find prometheus binary")
	}

	promCmd := exec.Command(promPath, "-version")

	promOut, err := promCmd.Output()
	if err != nil {
		log.Fatalf("unable to get prometheus version: %s", err)
	}

	promVersion = string(promOut)

	fmt.Printf("%s\n%s", sirynVersion, promVersion)
}
