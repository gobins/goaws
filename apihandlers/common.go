package apihandlers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func updateTag(tagname, tagvalue string, resources []*string) bool {
	log.Debug("Updating tag" + tagname + ":" + tagvalue)
	log.Debug(resources)
	params := &ec2.CreateTagsInput{
		Resources: resources,
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(tagname),
				Value: aws.String(tagvalue),
			},
		},
	}
	ec2client := getec2client()
	resp, err := ec2client.CreateTags(params)
	log.Debug(resp)
	if err != nil {
		log.Error("Error updating tag")
		log.Error(err)
	}
	return true
}
