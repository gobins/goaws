package apihandlers

import (
	"testing"
)

func TestGetAWSSession(t *testing.T) {
	awssession := getawssession("ap-southeast-2")
	if awssession == nil {
		t.Error("Error getting ec2 session")
	}
}

func TestGetEC2Client(t *testing.T) {
	ec2client := getec2client("ap-southeast-2")
	if ec2client == nil {
		t.Error("Error getting an ec2 client")
	}
}

func TestGetS3Client(t *testing.T) {
	s3client := gets3client("ap-southeast-2")
	if s3client == nil {
		t.Error("Error getting a s3 client")
	}
}

func TestGetCloudtrailClient(t *testing.T) {
	cloudtrailclient := getcloudtrailclient("ap-southeast-2")
	if cloudtrailclient == nil {
		t.Error("Error getting cloudtrail client")
	}
}
