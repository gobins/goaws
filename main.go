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
	app.Commands = []cli.Command{
		{
			Name:  "get-subnets",
			Usage: "List all subnets",

			Action: func(c *cli.Context) {
				log.Debug("Calling apihandlers.GetSubnetsFormatted")
				apihandlers.GetSubnetsFormatted()
			},
		},
	}
	app.Run(os.Args)

}
