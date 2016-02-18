package apihandlers

import (
	"encoding/json"
	"strings"
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
	EventID            string `json:"event_id"`
	EventName          string `json:"event_name"`
	Username           string `json:"username"`
	ResourceID         string `json:"resource_id"`
	EventTime          string `json:"event_time"`
	UserIdentity       string `json:"user_identity"`
	EventSource        string `json:"event_source"`
	AwsRegion          string `json:"aws_region"`
	SourceIPAddress    string `json:"source_IPAddress"`
	UserAgent          string `json:"user_agent"`
	RequestID          string `json:"request_id"`
	EventType          string `json:"event_type"`
	RecipientAccountId string `json:"recipient_accountId"`
	AccountType        string `json:"account_type"`
	PrincipalId        string `json:"principal_id"`
	Arn                string `json:"arn"`
	AccountId          string `json:"account_id"`
	AccessKeyId        string `json:"access_key_Id"`
}

func parseCloudtrailEvents(events []*cloudtrail.Event) []EventData {
	log.Debug("Parsing Cloudtrail Events Data")
	resp := make([]EventData, 0, 20)
	for _, event := range events {
		parsedData := new(EventData)
		parsedData.EventID = *event.EventId
		parsedData.EventName = *event.EventName
		parsedData.Username = *event.Username

		response := &EventResponse{}
		reader := strings.NewReader(*event.CloudTrailEvent)
		err := json.NewDecoder(reader).Decode(&response)
		if err != nil {
			log.Fatal(err)
		}

		parsedData.EventTime = response.EventTime
		parsedData.EventSource = response.EventSource
		parsedData.AwsRegion = response.AwsRegion
		parsedData.SourceIPAddress = response.SourceIPAddress
		parsedData.UserAgent = response.UserAgent
		parsedData.RequestID = response.RequestID
		parsedData.EventType = response.EventType
		parsedData.RecipientAccountId = response.RecipientAccountId
		parsedData.AccountType = response.UserIdentity.Type
		parsedData.PrincipalId = response.UserIdentity.PrincipalId
		parsedData.Arn = response.UserIdentity.Arn
		parsedData.AccountId = response.UserIdentity.AccountId
		parsedData.AccessKeyId = response.UserIdentity.AccessKeyId
		if len(event.Resources) == 0 {
			resp = append(resp, *parsedData)
		} else {
			for _, resource := range event.Resources {
				parsedData.ResourceID = *resource.ResourceName
				resp = append(resp, *parsedData)
			}
		}
	}
	return resp
}
