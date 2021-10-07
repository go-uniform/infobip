package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
	"service/service/info"
	"service/service/integrations/infobip"
)

func init() {
	_base.Subscribe(_base.TargetAction("sms", "send"), smsSend)
}

func smsSend(r uniform.IRequest, p diary.IPage) {
	var model infobip.SmsTextAdvanceRequest
	r.Read(model)

	p.Notice("email.send", nil)

	if info.TestMode {
		p.Notice("email.send.test-mode", diary.M{
			"messages": model.Messages,
		})

		if r.CanReply() {
			if err := r.Reply(uniform.Request{
				Model: infobip.SmsTextAdvanceResponse{},
			}); err != nil {
				p.Error("reply", err.Error(), diary.M{
					"err": err,
				})
			}
		}
		return
	}

	info.Infobip.SmsTextAdvanced(model)

	if r.CanReply() {
		if err := r.Reply(uniform.Request{
			Model: infobip.SmsTextAdvanceResponse{},
		}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
