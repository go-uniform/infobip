package service

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
	"service/service/info"
	"service/service/integrations/infobip"
	"sync"
	"time"
)

func RunBefore(shutdown chan bool, group *sync.WaitGroup, p diary.IPage) {
	uri, ok := info.Args["uri"].(string)
	if !ok {
		panic("uri must be a string")
	}
	apiKey, ok := info.Args["apiKey"].(string)
	if !ok {
		panic("apiKey must be a string")
	}

	info.Uri = uri
	info.ApiKey = apiKey

	info.Infobip = infobip.NewInfobipConnector(p, uri, apiKey, info.Virtual)

	var queues = []struct {
		QueueName string `bson:"queueName"`
	}{
		{QueueName: "email.send"},
		{QueueName: "sms.send"},
	}
	for _, queue := range queues {
		if err := info.Conn.Request(p, _base.TargetAction("queue", "create"), time.Second*10, uniform.Request{
			Model: queue,
		}, func(r uniform.IRequest, p diary.IPage) {
			if r.HasError() {
				panic(r.Error())
			}
			if r.CanReply() {
				if err := r.Reply(uniform.Request{}); err != nil {
					p.Error("reply", err.Error(), diary.M{
						"err": err,
					})
				}
			}
		}); err != nil {
			panic(err)
		}
	}
}
