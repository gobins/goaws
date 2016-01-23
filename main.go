package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "goaws"
	app.Usage = "Aws cli using Golang SDK"

	//Initialise EC2 client

	//Defining sub-commands
	app.Commands = []cli.Command{
		{
			Name:  "get-subnets",
			Usage: "List all subnets",
			Action: func(c *cli.Context) {
				fmt.Println("Testing", c.Args().First())
			},
		},
	}
	app.Run(os.Args)
}
