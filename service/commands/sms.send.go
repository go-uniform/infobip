package commands

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
)

func init() {
	_base.Subscribe(_base.TargetCommand("sms.send"), smsSend)
}

func smsSend(r uniform.IRequest, p diary.IPage) {
	var model struct {
		Queue bool     `bson:"queue"`
		From  string   `bson:"from"`
		To    []string `bson:"to"`
		Text  string   `bson:"text"`
	}
	r.Read(&model)

	if err := r.Conn().Request(p, _base.TargetAction("sms", "send"), r.Remainder(), uniform.Request{
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
			Model: "SMS sent!",
		}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
