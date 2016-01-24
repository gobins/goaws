package apihandlers

import (
	"testing"
)

func TestGetAWSSession(t *testing.T) {
	awssession := getawssession()
	if awssession == nil {
		t.Error("Error getting ec2 session")
	}
}

func TestGetEC2Client(t *testing.T) {
	ec2client := getec2client()
	if ec2client == nil {
		t.Error("Error getting an ec2 client")
	}
}

func TestGetS3Client(t *testing.T) {
	s3client := gets3client()
	if s3client == nil {
		t.Error("Error getting a s3 client")
	}
}

func TestGetCloudtrailClient(t *testing.T) {
	cloudtrailclient := getcloudtrailclient()
	if cloudtrailclient == nil {
		t.Error("Error getting cloudtrail client")
	}
}
