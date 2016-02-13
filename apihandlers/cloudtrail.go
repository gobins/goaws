package apihandlers

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

func lookupInstanceTrail(key, value string) []*cloudtrail.Event {
	log.Debug("Gathering Instance events for last 1 hour")
	cloudtrailclient := getcloudtrailclient()
	params := &cloudtrail.LookupEventsInput{
		EndTime: aws.Time(time.Now()),
		LookupAttributes: []*cloudtrail.LookupAttribute{
			{ // Required
				AttributeKey:   aws.String(key),   // Required
				AttributeValue: aws.String(value), // Required
			},
			// More values...
		},
		StartTime: aws.Time(time.Now().Add(-1 * time.Hour)),
	}
	resp, err := cloudtrailclient.LookupEvents(params)
	if err != nil {
		log.Error("Error Retrieving Instance Events")
		log.Error(err)
	}
	return resp.Events
}

type eventData struct {
	eventID    string
	eventName  string
	username   string
	resourceID string
}

func parseCloudtrailEvents(events []*cloudtrail.Event) []eventData {
	log.Debug("Parsing Cloudtrail Events Data")
	resp := make([]eventData, 0, 20)
	for _, event := range events {
		parsedData := new(eventData)
		parsedData.eventID = *event.EventId
		parsedData.eventName = *event.EventName
		parsedData.username = *event.Username
		for _, resource := range event.Resources {
			parsedData.resourceID = *resource.ResourceName
			resp = append(resp, *parsedData)
		}
	}
	return resp
}
