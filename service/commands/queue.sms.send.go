package commands

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
	"service/service/models"
)

func init() {
	_base.Subscribe(_base.TargetCommand("queue.sms.send"), queueSmsSend)
}

func queueSmsSend(r uniform.IRequest, p diary.IPage) {
	var model models.SmsSend
	r.Read(&model)

	if err := r.Conn().Request(p, _base.TargetAction("queue", "sms.send"), r.Remainder(), uniform.Request{
		Model: model,
	}, func(sub uniform.IRequest, _ diary.IPage) {
		if sub.HasError() {
			panic(sub.Error())
		}
	}); err != nil {
		panic(err)
	}

	if r.CanReply() {
		if err := r.Reply(uniform.Request{
			Model: "SMS queued!",
		}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
