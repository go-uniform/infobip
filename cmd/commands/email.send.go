package commands

import (
	"github.com/spf13/cobra"
	"html"
	"service/cmd/_base"
	"service/service"
	"time"
)

func init() {
	var fromName string
	var from string
	var to string
	var subject string
	var body string

	cmd := _base.Command("email.send", func(cmd *cobra.Command, args []string) {
		service.Command("email.send", time.Second, _base.NatsUri, _base.CompileNatsOptions(), map[string]string{
			"fromName": fromName,
			"from": from,
			"to": to,
			"subject": subject,
			"body": html.EscapeString(body), // escape string for security reasons
		}, nil)
	}, "Send an email message to a target email address via CLI")

	cmd.Flags().StringVarP(&fromName, "fromName", "", "Uprate", "The from email address name to be used when sending the email")
	cmd.Flags().StringVarP(&from, "from", "f", "noreply@uprate.co.za", "The from email address to be used when sending the email")
	cmd.Flags().StringVarP(&to, "to", "t", "", "The destination email address(es) to send the email to.")
	cmd.Flags().StringVarP(&subject, "subject", "s", "Test", "The message subject for the email to be sent.")
	cmd.Flags().StringVarP(&body, "body", "b", "Toto, I've got a feeling we're not in Kansas anymore.", "The message body for the email to be sent.")

	if err := cmd.MarkFlagRequired("to"); err != nil {
		panic(err)
	}

	_base.RootCmd.AddCommand(cmd)
}