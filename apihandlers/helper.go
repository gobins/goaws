package apihandlers

import "github.com/aws/aws-sdk-go/service/ec2"

func getTagValue(tags []*ec2.Tag, key string) (value string) {
	var resp string
	for _, tag := range tags {
		if *tag.Key == key {
			resp = *tag.Value
		}
	}
	return resp
}
