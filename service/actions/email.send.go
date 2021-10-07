package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
	"service/service/info"
	"service/service/integrations/infobip"
)

func init() {
	_base.Subscribe(_base.TargetAction("email", "send"), emailSend)
}

func emailSend(r uniform.IRequest, p diary.IPage) {
	var model infobip.EmailSendRequest
	r.Read(model)

	p.Notice("email.send", nil)

	if info.TestMode {
		p.Notice("email.send.test-mode", diary.M{
			"from":        model.From,
			"to":          model.To,
			"cc":          model.Cc,
			"subject":     model.Subject,
			"body":        model.Text,
			"attachments": model.Attachments,
		})

		if r.CanReply() {
			if err := r.Reply(uniform.Request{
				Model: infobip.EmailSendResponse{},
			}); err != nil {
				p.Error("reply", err.Error(), diary.M{
					"err": err,
				})
			}
		}
		return
	}

	info.Infobip.SendEmail(model)

	if r.CanReply() {
		if err := r.Reply(uniform.Request{
			Model: infobip.EmailSendResponse{},
		}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}
