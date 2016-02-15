package apihandlers

import (
	log "github.com/Sirupsen/logrus"
	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Volume struct {
	Name     string `json:"name"`
	VolumeId string `json:"volume_id"`
	State    string `json:"state"`
	Size     int64  `json:"size"`
	Wrk      string `json:"wrk"`
	Attached []string `json:"attached"`
}

func getAllVolumes() (volumes[]*ec2.Volume) {
	params := &ec2.DescribeVolumesInput{}
	
	ec2client := getec2client()
	resp, err := ec2client.DescribeVolumes(params)
	if err != nil {
		log.Error("Error retrieving volumes")
	}

	for _, volume := range resp.Volumes {
		volumes = append(volumes, volume)
	}
	return volumes
}

func parseVolumeData(volumes []*ec2.Volume, detached bool) (response []Volume) {
	resp := make([]Volume, 0, 20)
	for _, vol := range volumes {
		if detached && len(vol.Attachments) > 0 {
			continue
		}
		vms := make([]string, 0, 5)
		tags := vol.Tags
		name := getTagValue(tags, "Name")
		wrk  := getTagValue(tags, "WRK")
		for _, attmnt := range vol.Attachments {
			vms = append(vms, *attmnt.InstanceId)
		}
		parsedData := Volume{
			Name:     name,
			VolumeId: *vol.VolumeId,
			State:    *vol.State,
			Size:     *vol.Size,
			Wrk:      wrk,
			Attached: vms,
		}
		resp = append(resp, parsedData)
	}
	return resp
}
