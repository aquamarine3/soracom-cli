package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// CredentialsDeleteCmdCredentialsId holds value of 'credentials_id' option
var CredentialsDeleteCmdCredentialsId string

func init() {
	CredentialsDeleteCmd.Flags().StringVar(&CredentialsDeleteCmdCredentialsId, "credentials-id", "", TR("Credentials ID"))

	CredentialsCmd.AddCommand(CredentialsDeleteCmd)
}

// CredentialsDeleteCmd defines 'delete' subcommand
var CredentialsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TR("credentials.delete_credential.delete.summary"),
	Long:  TR(`credentials.delete_credential.delete.description`),
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

		param, err := collectCredentialsDeleteCmdParams()
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

func collectCredentialsDeleteCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForCredentialsDeleteCmd("/credentials/{credentials_id}"),
		query:  buildQueryForCredentialsDeleteCmd(),
	}, nil
}

func buildPathForCredentialsDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"credentials_id"+"}", CredentialsDeleteCmdCredentialsId, -1)

	return path
}

func buildQueryForCredentialsDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
