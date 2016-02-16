package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/godlikeachilles/goaws/apihandlers"
)

func main() {
	app := cli.NewApp()
	app.Name = "goaws"
	app.Usage = "Aws cli using Golang SDK"

	//Defining sub-commands
	var environment string
	var tagname string
	var tagvalue string
	var attkey string
	var attvalue string
	var format string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "format",
			Usage: "json or table. defaults to table.",
			Destination: &format,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "get-subnets",
			Usage: "List all subnets",

			Action: func(c *cli.Context) {
				if format == "" || format == "table" {
					format = "table"
				} else if format != "json" {
					fmt.Println("Unsupported format", format)
					os.Exit(1)
				}
				log.Debug("Calling apihandlers.GetSubnetsFormatted")
				apihandlers.GetSubnetsFormatted(format)
			},
		},
		{
			Name:  "get-trail",
			Usage: "Return Cloudtrail events",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "attkey",
					Usage:       "EventId | EventName | Username | ResourceType | ResourceName",
					Destination: &attkey,
				},
				cli.StringFlag{
					Name:        "attvalue",
					Usage:       "Attribute Value for the LookupAttribute",
					Destination: &attvalue,
				},
			},
			Action: func(c *cli.Context) {
				var errorFlag bool
				if attkey == "" {
					fmt.Println("attkey is required for update-tag subcommand")
					errorFlag = true
				}
				if attvalue == "" {
					fmt.Println("attvalue is required for update-tag subcommand")
					errorFlag = true
				}
				if errorFlag == true {
					os.Exit(1)
				}
				apihandlers.GetTrail(attkey, attvalue, format)
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
				apihandlers.GetInstancesFormatted(environment, format)
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
