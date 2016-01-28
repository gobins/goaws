package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gobins/goaws/apihandlers"
)

func main() {
	app := cli.NewApp()
	app.Name = "goaws"
	app.Usage = "Aws cli using Golang SDK"

	//Defining sub-commands
	var environment string
	app.Commands = []cli.Command{
		{
			Name:  "get-subnets",
			Usage: "List all subnets",

			Action: func(c *cli.Context) {
				log.Debug("Calling apihandlers.GetSubnetsFormatted")
				apihandlers.GetSubnetsFormatted()
			},
		},
		{
			Name:  "get-instances",
			Usage: "List all subnets",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "environment",
					Usage:       "test",
					Destination: &environment,
				},
			},
			Action: func(c *cli.Context) {
				log.Debug("Calling apihandlers.GetInstancesFormatted")
				apihandlers.GetInstancesFormatted(environment)
			},
		},
	}
	app.Run(os.Args)

}
