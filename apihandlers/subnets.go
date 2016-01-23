package apihandlers

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type subnetData struct {
	subnetID   string
	cidrBlock  string
	subnetWrk  string
	subnetName string
}

func getSubnets(awsregion string) []*ec2.Subnet {
	ec2client := getec2client(awsregion)
	params := &ec2.DescribeSubnetsInput{
		DryRun: aws.Bool(false),
	}
	log.Debug("Calling DescribeSubnets")
	resp, err := ec2client.DescribeSubnets(params)
	if err != nil {
		fmt.Println(err.Error())
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
