package commands

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
	"strings"
)

func init() {
	_base.Subscribe(_base.TargetCommand("email.send"), emailSend)
}

func emailSend(r uniform.IRequest, p diary.IPage) {
	var params uniform.P
	r.Read(&params)

	if err := r.Conn().Request(p, _base.TargetAction("email", "send"), r.Remainder(), uniform.Request{
		Model: struct {
			From        string        `bson:"from"`
			To          []string      `bson:"to"`
			Cc          []string      `bson:"cc"`
			Bcc         []string      `bson:"bcc"`
			Subject     string        `bson:"subject"`
			Body        string        `bson:"body"`
			Attachments []interface{} `bson:"attachments"`
		}{
			From:    params["from"],
			To:      strings.Split(params["to"], ","),
			Subject: params["subject"],
			Body:    params["body"],
		},
	}, func(sub uniform.IRequest, _ diary.IPage) {
		if sub.HasError() {
			panic(sub.Error())
		}
	}); err != nil {
		panic(err)
	}

	if r.CanReply() {
		if err := r.Reply(uniform.Request{
			Model: "Email sent!",
		}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
