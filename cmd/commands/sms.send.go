package commands

import (
	"github.com/spf13/cobra"
	"service/cmd/_base"
	"service/service"
	"time"
)

func init() {
	var from string
	var to string
	var body string

	cmd := _base.Command("sms.send", func(cmd *cobra.Command, args []string) {
		service.Command("sms.send", time.Second, _base.NatsUri, _base.CompileNatsOptions(), map[string]string{
			"from": from,
			"to": to,
			"body": body,
		}, nil)
	}, "Send an SMS message to a target mobile number via CLI")

	cmd.Flags().StringVarP(&from, "from", "f", "InfoSMS", "The from number, set this to InfoSMS to send from a long number.")
	cmd.Flags().StringVarP(&to, "to", "t", "", "The destination mobile number(s) to send the SMS to.")
	cmd.Flags().StringVarP(&body, "body", "b", "Toto, I've got a feeling we're not in Kansas anymore.", "The message body for the SMS to be sent.")

	if err := cmd.MarkFlagRequired("to"); err != nil {
		panic(err)
	}

	_base.RootCmd.AddCommand(cmd)
}