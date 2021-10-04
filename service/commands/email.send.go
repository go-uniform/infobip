package commands

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
)

func init() {
	_base.Subscribe(_base.TargetCommand("email.send"), emailSend)
}

func emailSend(r uniform.IRequest, p diary.IPage) {
	// todo: send email

	if r.CanReply() {
		if err := r.Reply(uniform.Request{
			Model: "pong",
		}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
