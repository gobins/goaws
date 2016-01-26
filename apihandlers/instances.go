package apihandlers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type instanceData struct {
	name          string
	instancesID   string
	instanceState string
	instanceWrk   string
	instanceType  string
	launchedBy    string
}

func getAllInstancesInSubnet(subnetID string) (instances []*ec2.Instance) {
	ec2client := getec2client()
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("subnet-id"),
				Values: []*string{
					aws.String(subnetID),
				},
			},
		},
	}
	resp, err := ec2client.DescribeInstances(params)
	if err != nil {
		log.Error("Error retrieving all instances in subnet")
	}
	reservation := resp.Reservations[0]
	instances = reservation.Instances
	return instances
}

func parseInstancesData(instances []*ec2.Instance) (response []instanceData) {
	resp := make([]instanceData, 0, 20)
	for _, instance := range instances {
		parsedData := new(instanceData)
		parsedData.instancesID = *instance.InstanceId
		parsedData.instanceState = *instance.State.Name
		parsedData.instanceType = *instance.InstanceType
		tags := instance.Tags
		parsedData.name = getTagValue(tags, "Name")
		parsedData.launchedBy = getTagValue(tags, "Launched_by")
		parsedData.instanceWrk = getTagValue(tags, "WRK")
		resp = append(resp, *parsedData)
	}
	return resp
}
