package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ehazlett/siryn/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "siryn"
	app.Version = version.FullVersion
	app.Usage = "docker container monitoring"
	app.Author = "@ehazlett"
	app.Email = ""
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}
	app.Commands = []cli.Command{
		cmdServe,
		cmdInfo,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
