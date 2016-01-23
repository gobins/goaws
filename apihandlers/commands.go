package apihandlers

import (
	"fmt"

	"github.com/apcera/termtables"
)

func getSubnetsFormatter() {
	table := termtables.CreateTable()
	table.AddHeaders("Name", "CIDR Block", "WRK", "Subnet Id")

	subnets := getSubnets("ap-southeast-2")
	data := parseSubnetsData(subnets)

	if data != nil {
		for _, row := range data {
			table.AddRow(row.subnetName, row.cidrBlock, row.subnetWrk, row.subnetID)
		}
	}

	fmt.Println(table.Render())
}
