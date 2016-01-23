package apihandlers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getawssession(awsregion string) (awssession *session.Session) {
	sess := session.New(&aws.Config{Region: aws.String(awsregion)})
	return sess
}

func getec2client(awsregion string) (ec2client *ec2.EC2) {
	sess := getawssession(awsregion)
	ec2client = ec2.New(sess)
	return ec2client
}

func gets3client(awsregion string) (s3client *s3.S3) {
	sess := getawssession(awsregion)
	s3client = s3.New(sess)
	return s3client
}

func getcloudtrailclient(awsregion string) (cloudtrailclient *cloudtrail.CloudTrail) {
	sess := getawssession(awsregion)
	cloudtrailclient = cloudtrail.New(sess)
	return cloudtrailclient
}
