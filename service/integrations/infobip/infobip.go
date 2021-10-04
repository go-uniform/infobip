package infobip

/* Infobip API

Reference: https://www.infobip.com/docs/api
*/

type infobip struct {
	BaseUri string
	ApiKey  string
}

type Infobip interface {
	SmsTextAdvanced()
}

func NewInfobipConnector() {

}
