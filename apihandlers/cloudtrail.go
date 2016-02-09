package apihandlers

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

func lookupInstanceTrail() {
	log.Debug("Gathering Instance events for last 1 hour")
	cloudtrailclient := getcloudtrailclient()
	params := &cloudtrail.LookupEventsInput{
		EndTime: aws.Time(time.Now()),
		LookupAttributes: []*cloudtrail.LookupAttribute{
			{ // Required
				AttributeKey:   aws.String("EventName"),    // Required
				AttributeValue: aws.String("RunInstances"), // Required
			},
			// More values...
		},
		StartTime: aws.Time(time.Now().Add(-12 * time.Hour)),
	}
	resp, err := cloudtrailclient.LookupEvents(params)
	if err != nil {
		log.Error("Error Retrieving Instance Events")
		log.Error(err)
		return
	}
	log.Info(resp.Events)

}
