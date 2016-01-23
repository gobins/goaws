package main

import (
	"os"

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
				//api := apihandlers.New()
				apihandlers.GetSubnetsFormatted()
			},
		},
	}
	app.Run(os.Args)

}
