package apihandlers

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/apcera/termtables"
)

//GetSubnetsFormatted retrieve all subnet in a region and format it
func GetSubnetsFormatted() {
	log.Debug("Creating Output Table for Subnets Data")
	table := termtables.CreateTable()
	table.AddHeaders("Name", "CIDR Block", "WRK", "Subnet Id")

	subnets := getAllSubnets()
	data := parseSubnetsData(subnets)

	if data != nil {
		for _, row := range data {
			table.AddRow(row.subnetName, row.cidrBlock, row.subnetWrk, row.subnetID)
		}
	}

	fmt.Println(table.Render())
}

//GetInstancesFormatted retrieve all insntances in the subnet
func GetInstancesFormatted(envname string) {
	log.Debug("Creating Output Table for all instances data in the subnet")
	table := termtables.CreateTable()
	table.AddHeaders("Name", "State", "WRK", "Launched By", "Instance Type", "Instance ID")
	subnetID := getSubnetIDByTag("Name", envname)
	instances := getAllInstancesInSubnet(subnetID)
	data := parseInstancesData(instances)
	if data != nil {
		for _, row := range data {
			table.AddRow(row.name, row.instanceState, row.instanceWrk, row.launchedBy, row.instanceType, row.instancesID)
		}
	}
	fmt.Println(table.Render())
}
