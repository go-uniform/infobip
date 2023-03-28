package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"service/cmd/_base"
	"service/service"
	"strings"
	"time"
)

func init() {
	var from string
	var to string
	var text string

	cmd := _base.Command("queue.sms.send", func(cmd *cobra.Command, args []string) {
		service.Command("queue.sms.send", time.Minute, _base.NatsUri, _base.CompileNatsOptions(), map[string]interface{}{
			"from": from,
			"to":   strings.Split(to, ","),
			"text": text,
		}, func(bytes []byte) {
			fmt.Println(string(bytes))
		})
	}, "Queue an SMS message via CLI")

	cmd.Flags().StringVarP(&from, "from", "f", "InfoSMS", "The from number, set this to InfoSMS to send from a long number.")
	cmd.Flags().StringVarP(&to, "to", "t", "", "The destination mobile number(s) to send the SMS to.")
	cmd.Flags().StringVarP(&text, "text", "b", "Toto, I've got a feeling we're not in Kansas anymore.", "The message text for the SMS to be sent.")

	if err := cmd.MarkFlagRequired("to"); err != nil {
		panic(err)
	}

	_base.RootCmd.AddCommand(cmd)
}
