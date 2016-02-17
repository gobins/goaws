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

//EventData type for handling cloudtrail events data
type EventData struct {
	EventID    string `json:"event_id"`
	EventName  string `json:"name"`
	Username   string `json:"username"`
	ResourceID string `json:"resource_id"`
	EventTime  string `json:"event_time"`
}

func parseCloudtrailEvents(events []*cloudtrail.Event) []EventData {
	log.Debug("Parsing Cloudtrail Events Data")
	resp := make([]EventData, 0, 20)
	for _, event := range events {
		parsedData := new(EventData)
		parsedData.EventID = *event.EventId
		parsedData.EventName = *event.EventName
		parsedData.Username = *event.Username
		for _, resource := range event.Resources {
			parsedData.ResourceID = *resource.ResourceName
			resp = append(resp, *parsedData)
		}
	}
	return resp
}
