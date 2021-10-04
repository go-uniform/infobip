package infobip

import "time"

/* Send SMS message
99% of all use cases can be achieved by using this API method. Everything from sending a simple single message to a
single destination, up to batch sending of personalized messages to the thousands of recipients with a single API
request. Language, transliteration, scheduling and every advanced feature you can think of is supported.

POST: /sms/2/text/advanced

Reference: https://www.infobip.com/docs/api#channels/sms/send-sms-message
*/

/*
SmsTextAdvanceRequest
*/

type SmsTextAdvanceRequestMessageDeliveryWindowTime struct {
	Hour int32
	Minute int32
}

type SmsTextAdvanceRequestMessageDeliveryWindow struct {
	Days []string
	From SmsTextAdvanceRequestMessageDeliveryWindowTime
	To SmsTextAdvanceRequestMessageDeliveryWindowTime
}

type SmsTextAdvanceRequestMessageDestination struct {
	MessageId string `json:",omitempty"`
	To string
}

type SmsTextAdvanceRequestMessageLanguage struct {
	LanguageCode string
}

type SmsTextAdvanceRequestMessageRegionalIndiaDlt struct {
	ContentTemplateId string `json:",omitempty"`
	PrincipalEntityId string
}

type SmsTextAdvanceRequestMessageRegional struct {
	IndiaDlt *SmsTextAdvanceRequestMessageRegionalIndiaDlt `json:",omitempty"`
}

type SmsTextAdvanceRequestMessage struct {
	CallbackData string `json:",omitempty"`
	DeliveryTimeWindow *SmsTextAdvanceRequestMessageDeliveryWindow `json:",omitempty"`
	Destinations []SmsTextAdvanceRequestMessageDestination
	Flash bool `json:",omitempty"`
	From string `json:",omitempty"`
	IntermediateReport bool `json:",omitempty"`
	Language *SmsTextAdvanceRequestMessageLanguage `json:",omitempty"`
	NotifyContentType string `json:",omitempty"`
	NotifyUrl string `json:",omitempty"`
	Regional *SmsTextAdvanceRequestMessageRegional `json:",omitempty"`
	SendAt *time.Time `json:",omitempty"`
	Text string `json:",omitempty"`
	Transliteration string `json:",omitempty"`
	ValidityPeriod int64 `json:",omitempty"`
}

type SmsTextAdvanceRequestSendingSpeedLimit struct {
	Amount int32
	TimeUnit string `json:",omitempty"`
}

type SmsTextAdvanceRequestTracking struct {
	BaseUrl string `json:",omitempty"`
	ProcessKey string `json:",omitempty"`
	Track string `json:",omitempty"`
	Type string `json:",omitempty"`
}

type SmsTextAdvanceRequest struct {
	BulkId string `json:",omitempty"`
	Messages []SmsTextAdvanceRequestMessage `json:",omitempty"`
	SendingSpeedLimit *SmsTextAdvanceRequestSendingSpeedLimit `json:",omitempty"`
	Tracking *SmsTextAdvanceRequestTracking `json:",omitempty"`
}

/*
SmsTextAdvanceResponse
*/

type SmsTextAdvanceResponseMessageStatus struct {
	Action string
	Description string
	GroupId int32
	GroupName string
	Id int32
	Name string
}

type SmsTextAdvanceResponseMessage struct {
	MessageId string
	Status SmsTextAdvanceResponseMessageStatus
	To string
}

type SmsTextAdvanceResponse struct {
	BulkId string
	Messages []SmsTextAdvanceResponseMessage
}

/*
SmsTextAdvanceError
*/

type SmsTextAdvanceErrorRequestErrorServiceException struct {
	MessageId string
	Text string // details error description
}

type SmsTextAdvanceErrorRequestError struct {
	ServiceException SmsTextAdvanceErrorRequestErrorServiceException
}

type SmsTextAdvanceError struct {
	RequestError SmsTextAdvanceErrorRequestError
}

/*
Request
*/

func (i *infobip) SmsTextAdvanced(request SmsTextAdvanceRequest) SmsTextAdvanceResponse {
	// todo: implement http request
	panic("not yet implemented")
}
