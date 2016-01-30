package apihandlers

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/apcera/termtables"
	"github.com/aws/aws-sdk-go/aws"
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
	if subnetID != "" {
		instances := getAllInstancesInSubnet(subnetID)
		data := parseInstancesData(instances)
		fmt.Println(data)
		if data != nil {
			for _, row := range data {
				table.AddRow(row.name, row.instanceState, row.instanceWrk, row.launchedBy, row.instanceType, row.instancesID)
			}
		}
	}
	fmt.Println(table.Render())
}

//UpdateEnvTags updates tag with a value for all objects in subnet
func UpdateEnvTags(envname string, wrk string) {
	log.Debug("Updating tags in all objects in subnet")
	subnetID := getSubnetIDByTag("Name", envname)
	instances := getAllInstancesInSubnet(subnetID)
	parsedData := parseInstancesData(instances)
	resources := make([]*string, 5, 20)
	for _, data := range parsedData {
		resources = append(resources, aws.String(data.instancesID))
	}
	updateTag("Name", wrk, resources)
	GetInstancesFormatted(envname)
}
