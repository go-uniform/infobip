package cmd

import (
	"github.com/go-uniform/uniform"
	"github.com/spf13/cobra"
	"service/cmd/_base"
	"service/service"
	"service/service/info"
)

func init() {
	var level string
	var rate int
	var limit int
	var test bool
	var uri string
	var apiKey string

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run " + info.AppName + " service",
		Long:  "Run " + info.AppName + " service",
		Run: func(cmd *cobra.Command, args []string) {
			service.Execute(level, rate, limit, test, _base.NatsUri, _base.CompileNatsOptions(), uniform.M{
				"nats":       _base.NatsUri,
				"natsCert":   _base.NatsCert,
				"natsKey":    _base.NatsKey,
				"disableTls": _base.DisableTls,
				"lvl":        level,
				"rate":       rate,
				"limit":      limit,
				"test":       test,

				"uri":    uri,
				"apiKey": apiKey,
			})
		},
	}

	// set the service's environment configurations via many command-line-interface (CLI) arguments
	runCmd.Flags().StringVarP(&level, "lvl", "l", "notice", "The logging level ['trace', 'debug', 'info', 'notice', 'warning', 'error', 'fatal'] that service is running in")
	runCmd.Flags().IntVarP(&rate, "rate", "r", 1000, "The sample rate of the trace logs used for performance auditing [set to -1 to log every trace]")
	runCmd.Flags().IntVarP(&limit, "limit", "x", 1000, "The messages per second that each topic worker will be limited to [set to 0 or less for maximum throughput]")
	runCmd.Flags().BoolVar(&test, "test", false, "A flag indicating if service should enter into test mode")
	runCmd.Flags().StringVar(&uri, "uri", "", "The Infobip API BaseURI")
	runCmd.Flags().StringVar(&apiKey, "apiKey", "", "The Infobip API Key")

	if err := runCmd.MarkFlagRequired("uri"); err != nil {
		panic(err)
	}
	if err := runCmd.MarkFlagRequired("apiKey"); err != nil {
		panic(err)
	}

	_base.RootCmd.AddCommand(runCmd)
}
