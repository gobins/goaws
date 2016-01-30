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
	var tagname string
	var tagvalue string
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
					Usage:       "Environment Name",
					Destination: &environment,
				},
			},
			Action: func(c *cli.Context) {
				log.Debug("Calling apihandlers.GetInstancesFormatted")
				apihandlers.GetInstancesFormatted(environment)
			},
		},
		{
			Name:  "update-tag",
			Usage: "Updates tag for all objects in subnet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "environment",
					Usage:       "Environment Name",
					Destination: &environment,
				},
				cli.StringFlag{
					Name:        "tagname",
					Usage:       "Key of the tag",
					Destination: &tagname,
				},
				cli.StringFlag{
					Name:        "tagvalue",
					Usage:       "Value of the tag",
					Destination: &tagvalue,
				},
			},
			Action: func(c *cli.Context) {
				log.Debug("Calling apihandlers.UpdateEnvTags")
				apihandlers.UpdateEnvTags(tagname, tagvalue, environment)
			},
		},
	}
	app.Run(os.Args)

}
