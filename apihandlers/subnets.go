package apihandlers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/apcera/termtables"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type subnetData struct {
	subnetID   string
	cidrBlock  string
	subnetWrk  string
	subnetName string
}

func getAllSubnets() []*ec2.Subnet {
	ec2client := getec2client()
	params := &ec2.DescribeSubnetsInput{
		DryRun: aws.Bool(false),
	}
	log.Debug("Calling DescribeSubnets")
	resp, err := ec2client.DescribeSubnets(params)
	if err != nil {
		log.Error("Error calling DescribeSubnets")
	}
	return resp.Subnets
}

func parseSubnetsData(subnets []*ec2.Subnet) (response []subnetData) {

	resp := make([]subnetData, 0, 4)
	log.Debug("Parsing Subnets Data")
	if subnets != nil {
		for _, subnet := range subnets {
			parsedData := new(subnetData)
			parsedData.subnetID = *subnet.SubnetId
			parsedData.cidrBlock = *subnet.CidrBlock
			tags := subnet.Tags
			if tags != nil {
				parsedData.subnetWrk = getTagValue(tags, "WRK")
				parsedData.subnetName = getTagValue(tags, "Name")
			}
			resp = append(resp, *parsedData)
		}
	}
	return resp
}

func getSubnetIDByTag(tagname, tagvalue string) (subnetID string) {
	ec2client := getec2client()
	params := &ec2.DescribeSubnetsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:" + tagname),
				Values: []*string{
					aws.String(tagvalue),
				},
			},
		},
	}

	resp, err := ec2client.DescribeSubnets(params)

	if err != nil {
		log.Error("Error calling DescribeSubnets")
		return subnetID
	}
	subnets := resp.Subnets
	if subnets == nil {
		log.Error("Cannot find subnet with the tag name/value")
		return subnetID
	} else if len(subnets) > 1 {
		log.Error("More than one subnet with the tag exists")
		table := termtables.CreateTable()
		table.AddHeaders("Subnet Id", "CIDR Block", tagname)

		for _, subnet := range subnets {
			table.AddRow(*subnet.SubnetId, *subnet.CidrBlock, tagvalue)
		}
		log.Error(table.Render())
		log.Error("Returning the first subnet in the list")
	} else {
		subnet := resp.Subnets[0]
		subnetID = *subnet.SubnetId
	}
	return subnetID
}
