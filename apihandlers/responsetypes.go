package apihandlers

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Items struct {
	items []Item
}

type TagSet struct {
	Items `json:"items"`
}
type ResourceItems struct {
	ResourceId string `json:"resourceId"`
}

type UserIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principalId"`
	Arn         string `json:"arn"`
	AccountId   string `json:"accountId"`
	AccessKeyId string `json:"accessKeyId"`
	UserName    string `json:"userName"`
}

type ResponseElements struct {
	Return string `json:"_return"`
}

type ResourcesSet struct {
	ResourceItems `json:"items"`
}

type RequestParameters struct {
	ResourcesSet `json:"resourcesSet"`
	TagSet       `json:"tagSet"`
}

type EventResponse struct {
	EventVersion    string `json:"eventVersion"`
	UserIdentity    `json:"userIdentity"`
	EventTime       string `json:"eventTime"`
	EventSource     string `json:"eventSource"`
	EventName       string `json:"eventName"`
	AwsRegion       string `json:"awsRegion"`
	SourceIPAddress string `json:"sourceIPAddress"`
	UserAgent       string `json:"userAgent"`
	//RequestParameters  `json:"requestParameters"`
	//ResponseElements   `json:"responseElements"`
	RequestID          string `json:"requestID"`
	EventID            string `json:"eventID"`
	EventType          string `json:"eventType"`
	RecipientAccountId string `json:"recipientAccountId"`
}
