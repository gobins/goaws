package apihandlers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getSubnets(awsregion string) []*ec2.Subnet {
	ec2client := getec2client(awsregion)
	params := &ec2.DescribeSubnetsInput{
		DryRun: aws.Bool(false),
	}
	resp, err := ec2client.DescribeSubnets(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	return resp.Subnets
}

func parseSubnetsData(subnets []ec2.Subnet) {
	if subnets != nil {
		type subnetData struct {
			subnetID   string
			cidrBlock  string
			subnetWrk  string
			subnetName string
		}

		for _, subnet := range subnets {
			parsedData := new(subnetData)
			parsedData.subnetID = *subnet.SubnetId
			parsedData.cidrBlock = *subnet.CidrBlock
			tags := subnet.Tags
			if tags != nil {
				//	test := getTagValue(tags, "WRK")
			}
		}
	}
}

func getTagValue(tags []*Tag, key string) (value string) {
	var resp string
	for _, tag := range tags {
		if tag.Key == key {
			value := tag.Value
		}
	}

	return resp
}
