package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"html"
	"service/cmd/_base"
	"service/service"
	"strings"
	"time"
)

func init() {
	var from string
	var to string
	var subject string
	var body string

	cmd := _base.Command("queue.email.send", func(cmd *cobra.Command, args []string) {
		service.Command("queue.email.send", time.Minute, _base.NatsUri, _base.CompileNatsOptions(), map[string]interface{}{
			"from":    from,
			"to":      strings.Split(to, ","),
			"subject": subject,
			"body":    html.EscapeString(body), // escape string for security reasons
		}, func(bytes []byte) {
			fmt.Println(string(bytes))
		})
	}, "Queue an email message via CLI")

	cmd.Flags().StringVarP(&from, "from", "f", "noreply@uniform.co.za", "The from email address to be used when sending the email")
	cmd.Flags().StringVarP(&to, "to", "t", "", "The destination email address(es) to send the email to.")
	cmd.Flags().StringVarP(&subject, "subject", "s", "Test", "The message subject for the email to be sent.")
	cmd.Flags().StringVarP(&body, "body", "b", "Toto, I've got a feeling we're not in Kansas anymore.", "The message body for the email to be sent.")

	if err := cmd.MarkFlagRequired("to"); err != nil {
		panic(err)
	}

	_base.RootCmd.AddCommand(cmd)
}
