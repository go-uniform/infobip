package infobip

import "github.com/go-diary/diary"

/* Infobip API

Reference: https://www.infobip.com/docs/api
*/

type infobip struct {
	Page    diary.IPage
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

func NewInfobipConnector(page diary.IPage, baseUri, apiKey string) IInfobip {
	var instance IInfobip
	page.Scope("mongo", func(p diary.IPage) {
		instance = &infobip{
			BaseUri: baseUri,
			ApiKey:  apiKey,
		}
	})
	return instance
}
