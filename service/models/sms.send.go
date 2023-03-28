package models

type SmsSend struct {
	From string   `bson:"from"`
	To   []string `bson:"to"`
	Text string   `bson:"text"`
}
