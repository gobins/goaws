package apihandlers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getawssession() (awssession *session.Session) {
	sess := session.New(&aws.Config{Region: aws.String("ap-southeast-2")})
	return sess
}

func getec2client() (ec2client *ec2.EC2) {
	sess := getawssession()
	ec2client = ec2.New(sess)
	return ec2client
}

func gets3client() (s3client *s3.S3) {
	sess := getawssession()
	s3client = s3.New(sess)
	return s3client
}

func getcloudtrailclient() (cloudtrailclient *cloudtrail.CloudTrail) {
	sess := getawssession()
	cloudtrailclient = cloudtrail.New(sess)
	return cloudtrailclient
}
