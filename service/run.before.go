package service

import (
	"github.com/go-diary/diary"
	"service/service/info"
	"service/service/integrations/infobip"
	"sync"
)

func RunBefore(shutdown chan bool, group *sync.WaitGroup, p diary.IPage) {
	uri, ok := info.Args["uri"].(string)
	if !ok {
		panic("mongo uri must be a string")
	}
	apiKey, ok := info.Args["apiKey"].(string)
	if !ok {
		panic("mongo authSource must be a string")
	}

	info.Uri = uri
	info.ApiKey = apiKey

	info.Infobip = infobip.NewInfobipConnector(p, uri, apiKey)
}
