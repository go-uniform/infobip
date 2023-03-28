package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
	"service/service/models"
	"time"
)

func init() {
	_base.Subscribe(_base.TargetAction("queue", "sms.send"), queuePushSmsSend)
}

func queuePushSmsSend(r uniform.IRequest, p diary.IPage) {
	var model models.SmsSend
	r.Read(&model)

	p.Notice("queue.sms.send", diary.M{
		"model": model,
	})

	if err := r.Conn().Request(p, _base.TargetAction("queue", "push"), time.Second, uniform.Request{
		Parameters: map[string]string{
			"queueName": "sms.send",
		},
		Model: model,
	}, func(sub uniform.IRequest, _ diary.IPage) {
		if sub.CanReply() {
			if err := sub.Reply(uniform.Request{}); err != nil {
				p.Error("reply", err.Error(), diary.M{
					"err": err,
				})
			}
		}
	}); err != nil {
		panic(err)
	}

	if r.CanReply() {
		if err := r.Reply(uniform.Request{}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
