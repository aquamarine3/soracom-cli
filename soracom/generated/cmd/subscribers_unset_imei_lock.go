package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersUnsetImeiLockCmdImsi holds value of 'imsi' option
var SubscribersUnsetImeiLockCmdImsi string

func init() {
	SubscribersUnsetImeiLockCmd.Flags().StringVar(&SubscribersUnsetImeiLockCmdImsi, "imsi", "", TR("subscribers.unset_imei_lock.post.parameters.imsi.description"))

	SubscribersCmd.AddCommand(SubscribersUnsetImeiLockCmd)
}

// SubscribersUnsetImeiLockCmd defines 'unset-imei-lock' subcommand
var SubscribersUnsetImeiLockCmd = &cobra.Command{
	Use:   "unset-imei-lock",
	Short: TR("subscribers.unset_imei_lock.post.summary"),
	Long:  TR(`subscribers.unset_imei_lock.post.description`),
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

		param, err := collectSubscribersUnsetImeiLockCmdParams()
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

func collectSubscribersUnsetImeiLockCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersUnsetImeiLockCmd("/subscribers/{imsi}/unset_imei_lock"),
		query:  buildQueryForSubscribersUnsetImeiLockCmd(),
	}, nil
}

func buildPathForSubscribersUnsetImeiLockCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersUnsetImeiLockCmdImsi, -1)

	return path
}

func buildQueryForSubscribersUnsetImeiLockCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
