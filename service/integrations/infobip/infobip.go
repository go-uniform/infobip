package infobip

/* Infobip API

Reference: https://www.infobip.com/docs/api
*/

type infobip struct {
	BaseUri string
	ApiKey  string
}

type IInfobip interface {
	SmsTextAdvanced(request SmsTextAdvanceRequest) SmsTextAdvanceResponse
	PersonCreate(request PersonCreateRequest) PersonCreateResponse
	PersonRemove(request PersonRemoveQueryRequest) PersonRemoveResponse
	PersonUpdate(request PersonUpdateRequest, queryRequest PersonUpdateQueryRequest) PersonUpdateResponse
	SendEmail(request EmailSendRequest) EmailSendResponse
}

func NewInfobipConnector(baseUri, apiKey string) IInfobip {
	return &infobip{
		BaseUri: baseUri,
		ApiKey:  apiKey,
	}
}
