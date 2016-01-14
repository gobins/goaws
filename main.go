package main

import (
  "os"
  "github.com/codegangsta/cli"
  "github.com/gobins/goaws/apihandlers"
)

func func main() {
  app :=cli.NewApp().Run(os.Args)
  app.Name = "goaws"
  app.Usage = "Aws cli using Golang SDK"

  //Initialise EC2 client

  //Defining sub-commands
  app.Commands = []cli.Command{
    {
      Name: "get-subnets",
      Usage: "List all subnets"
      Action: func(c *cli.Context){
          lister.List(getsubnets())
      }
  }
  }
}
