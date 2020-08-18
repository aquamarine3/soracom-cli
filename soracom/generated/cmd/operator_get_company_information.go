// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OperatorGetCompanyInformationCmdOperatorId holds value of 'operator_id' option
var OperatorGetCompanyInformationCmdOperatorId string

func init() {
	OperatorGetCompanyInformationCmd.Flags().StringVar(&OperatorGetCompanyInformationCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	OperatorCmd.AddCommand(OperatorGetCompanyInformationCmd)
}

// OperatorGetCompanyInformationCmd defines 'get-company-information' subcommand
var OperatorGetCompanyInformationCmd = &cobra.Command{
	Use:   "get-company-information",
	Short: TRAPI("/operators/{operator_id}/company_information:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/company_information:get:description`),
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

		param, err := collectOperatorGetCompanyInformationCmdParams(ac)
		if err != nil {
			return err
		}

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectOperatorGetCompanyInformationCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorGetCompanyInformationCmdOperatorId == "" {
		OperatorGetCompanyInformationCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForOperatorGetCompanyInformationCmd("/operators/{operator_id}/company_information"),
		query:  buildQueryForOperatorGetCompanyInformationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorGetCompanyInformationCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorGetCompanyInformationCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorGetCompanyInformationCmd() url.Values {
	result := url.Values{}

	return result
}