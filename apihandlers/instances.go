package apihandlers

import (
	"sort"

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
	volumes       []string
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

	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, instance)
		}
	}
	return instances
}

func parseInstancesData(instances []*ec2.Instance) (response []instanceData) {
	resp := make([]instanceData, 0, 20)
	for _, instance := range instances {
		volumes := make([]string, 0, 5)
		parsedData := new(instanceData)
		parsedData.instancesID = *instance.InstanceId
		parsedData.instanceState = *instance.State.Name
		parsedData.instanceType = *instance.InstanceType
		tags := instance.Tags
		parsedData.name = getTagValue(tags, "Name")
		parsedData.launchedBy = getTagValue(tags, "Launched_by")
		parsedData.instanceWrk = getTagValue(tags, "WRK")
		//log.Info(instance.BlockDeviceMappings)
		for _, blockdevices := range instance.BlockDeviceMappings {
			volumes = append(volumes, *blockdevices.Ebs.VolumeId)
		}
		parsedData.volumes = volumes
		resp = append(resp, *parsedData)
	}
	instanceDataSorter(resp)
	return resp
}

//NameSorter is an implementation of Sort
type NameSorter []instanceData

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].name < a[j].name }

func instanceDataSorter(data []instanceData) {
	sort.Sort(NameSorter(data))
}
