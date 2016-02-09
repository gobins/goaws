package main

import (
	"fmt"
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
			Name:  "get-trail",
			Usage: "Return Cloudtrail events",

			Action: func(c *cli.Context) {
				apihandlers.GetTrail()
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
				var errorFlag bool
				if environment == "" {
					fmt.Println("environment is required for update-tag subcommand")
					errorFlag = true
				}
				if errorFlag == true {
					os.Exit(1)
				}
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
				var errorFlag bool
				if tagname == "" {
					fmt.Println("tagname is required for update-tag subcommand")
					errorFlag = true
				}
				if tagvalue == "" {
					fmt.Println("tagvalue is required for update-tag subcommand")
					errorFlag = true
				}
				if environment == "" {
					fmt.Println("environment is required for update-tag subcommand")
					errorFlag = true
				}
				if errorFlag == true {
					os.Exit(1)
				}
				apihandlers.UpdateEnvTags(tagname, tagvalue, environment)
			},
		},
	}
	app.Run(os.Args)

}
