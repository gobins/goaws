package apihandlers

import (
	"github.com/apcera/termtables"
)

func getSubnetsFormatter() {
	table := termtables.CreateTable()
	table.AddHeaders("Name", "CIDR Block", "WRK", "Subnet Id")

	subnets := getSubnets("ap-southeast-2")

}
