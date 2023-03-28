package models

type EmailSend struct {
	From        string        `bson:"from"`
	To          []string      `bson:"to"`
	Cc          []string      `bson:"cc"`
	Bcc         []string      `bson:"bcc"`
	Subject     string        `bson:"subject"`
	Body        string        `bson:"body"`
	Attachments []interface{} `bson:"attachments"`
}
