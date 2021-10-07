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
	var model struct {
		From string   `bson:"from"`
		To   []string `bson:"to"`
		Text string   `bson:"text"`
	}
	r.Read(&model)

	p.Notice("sms.send", nil)

	if info.TestMode {
		p.Notice("sms.send.test-mode", diary.M{
			"model": model,
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

	destinations := make([]infobip.SmsTextAdvanceRequestMessageDestination, 0)
	for _, to := range model.To {
		destinations = append(destinations, infobip.SmsTextAdvanceRequestMessageDestination{
			To: to,
		})
	}

	info.Infobip.SmsTextAdvanced(infobip.SmsTextAdvanceRequest{
		Messages: []infobip.SmsTextAdvanceRequestMessage{
			{
				From:         model.From,
				Destinations: destinations,
				Text:         model.Text,
			},
		},
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
}
