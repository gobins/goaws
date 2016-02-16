package apihandlers

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/apcera/termtables"
	"github.com/aws/aws-sdk-go/aws"
)

// produce json output
func js(what string, data interface{}) {
	fmt.Println("{", what, ": ")
	js, _ := json.MarshalIndent(data, "  ", "  ")
	fmt.Println(string(js))
	fmt.Println("}")
}

//GetSubnetsFormatted retrieve all subnet in a region and format it
func GetSubnetsFormatted(format string) {
	log.Debug("Creating Output Table for Subnets Data")

	subnets := getAllSubnets()
	data := parseSubnetsData(subnets)
	if format == "table" {
		table := termtables.CreateTable()
		table.AddHeaders("Name", "CIDR Block", "WRK", "Subnet Id")
		if data != nil {
			for _, row := range data {
				table.AddRow(row.SubnetName, row.CidrBlock, row.SubnetWrk, row.SubnetID)
			}
		}

		fmt.Println(table.Render())
	} else {
		js("subnets", data)
	}

}

//GetInstancesFormatted retrieve all instances in the subnet
func GetInstancesFormatted(envname, format string) {
	log.Debug("Creating Output Table for all instances data in the subnet")
	subnetID := getSubnetIDByTag("Name", envname)
	if subnetID != "" {
		instances := getAllInstancesInSubnet(subnetID)
		data := parseInstancesData(instances)
		if format == "table" {
			table := termtables.CreateTable()
			table.AddHeaders("Name", "State", "WRK", "Launched By", "Instance Type", "Instance ID")
			if data != nil {
				for _, row := range data {
					table.AddRow(row.Name, row.InstanceState, row.InstanceWrk, row.LaunchedBy, row.InstanceType, row.InstancesID)
				}
			}
			fmt.Println(table.Render())
		} else {
			js("instances", data)
		}
	}
}

//UpdateEnvTags updates tag with a value for all objects in subnet
func UpdateEnvTags(tagname, tagvalue, envname string) {
	log.Debug("Updating tags in all objects in subnet")
	subnetID := getSubnetIDByTag("Name", envname)
	instances := getAllInstancesInSubnet(subnetID)
	parsedData := parseInstancesData(instances)
	resources := make([]*string, 0, 20)
	resources = append(resources, &subnetID)
	for _, data := range parsedData {
		resources = append(resources, aws.String(data.InstancesID))
		for _, volume := range data.Volumes {
			resources = append(resources, aws.String(volume))
		}
	}
	table := termtables.CreateTable()
	table.AddHeaders("Updating Resources")
	for _, resource := range resources {
		table.AddRow(*resource)
	}
	fmt.Println(table.Render())
	updateTag(tagname, tagvalue, resources)
	GetInstancesFormatted(envname, "table")
}

//GetTrail returns events captured in cloudtrail
func GetTrail(key, value, format string) {
	resp := lookupInstanceTrail(key, value)
	events := parseCloudtrailEvents(resp)
	if format == "table" {
		table := termtables.CreateTable()
		table.AddHeaders("EventID", "ResourceID", "Username", "EventName")
		for _, event := range events {
			table.AddRow(event.EventID, event.ResourceID, event.Username, event.EventName)
		}
		fmt.Println(table.Render())
	} else {
		js("events", events)
	}
}
