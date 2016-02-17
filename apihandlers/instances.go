package apihandlers

import (
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type InstanceData struct {
	Name          string `json:"name"`
	InstancesID   string `json:"instance_id"`
	InstanceState string `json:"state"`
	InstanceWrk   string `json:"wrk"`
	InstanceType  string `json:"type"`
	LaunchedBy    string `json:"launched_by"`
	Volumes       []string `json:"volumes"`
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

func parseInstancesData(instances []*ec2.Instance) (response []InstanceData) {
	resp := make([]InstanceData, 0, 20)
	for _, instance := range instances {
		volumes := make([]string, 0, 5)
		parsedData := new(InstanceData)
		parsedData.InstancesID = *instance.InstanceId
		parsedData.InstanceState = *instance.State.Name
		parsedData.InstanceType = *instance.InstanceType
		tags := instance.Tags
		parsedData.Name = getTagValue(tags, "Name")
		parsedData.LaunchedBy = getTagValue(tags, "Launched_by")
		parsedData.InstanceWrk = getTagValue(tags, "WRK")
		//log.Info(instance.BlockDeviceMappings)
		for _, blockdevices := range instance.BlockDeviceMappings {
			volumes = append(volumes, *blockdevices.Ebs.VolumeId)
		}
		parsedData.Volumes = volumes
		resp = append(resp, *parsedData)
	}
	instanceDataSorter(resp)
	return resp
}

//NameSorter is an implementation of Sort
type NameSorter []InstanceData

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

func instanceDataSorter(data []InstanceData) {
	sort.Sort(NameSorter(data))
}
