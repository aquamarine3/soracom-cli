package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersSuspendCmdImsi holds value of 'imsi' option
var SubscribersSuspendCmdImsi string

func init() {
	SubscribersSuspendCmd.Flags().StringVar(&SubscribersSuspendCmdImsi, "imsi", "", TR("subscribers.suspend_subscriber.post.parameters.imsi.description"))

	SubscribersCmd.AddCommand(SubscribersSuspendCmd)
}

// SubscribersSuspendCmd defines 'suspend' subcommand
var SubscribersSuspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: TR("subscribers.suspend_subscriber.post.summary"),
	Long:  TR(`subscribers.suspend_subscriber.post.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectSubscribersSuspendCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectSubscribersSuspendCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersSuspendCmd("/subscribers/{imsi}/suspend"),
		query:  buildQueryForSubscribersSuspendCmd(),
	}, nil
}

func buildPathForSubscribersSuspendCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersSuspendCmdImsi, -1)

	return path
}

func buildQueryForSubscribersSuspendCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
