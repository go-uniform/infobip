package infobip

import (
	"github.com/go-diary/diary"
)

/* Infobip API

Reference: https://www.infobip.com/docs/api
*/

type infobip struct {
	Page    diary.IPage
	BaseUri string
	ApiKey  string
	Virtual bool
}

type IInfobip interface {
	SmsTextAdvanced(request SmsTextAdvanceRequest) SmsTextAdvanceResponse
	PersonCreate(request PersonCreateRequest) PersonCreateResponse
	PersonRemove(request PersonRemoveQueryRequest) PersonRemoveResponse
	PersonUpdate(request PersonUpdateRequest, queryRequest PersonUpdateQueryRequest) PersonUpdateResponse
	SendEmail(request EmailSendRequest) EmailSendResponse
}

func NewInfobipConnector(page diary.IPage, baseUri, apiKey string, virtual bool) IInfobip {
	var instance IInfobip
	page.Scope("infobip", func(p diary.IPage) {
		instance = &infobip{
			Page:    page,
			BaseUri: baseUri,
			ApiKey:  apiKey,
			Virtual: virtual,
		}
	})
	return instance
}
