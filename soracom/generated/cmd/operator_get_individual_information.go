// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OperatorGetIndividualInformationCmdOperatorId holds value of 'operator_id' option
var OperatorGetIndividualInformationCmdOperatorId string

func init() {
	OperatorGetIndividualInformationCmd.Flags().StringVar(&OperatorGetIndividualInformationCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	OperatorCmd.AddCommand(OperatorGetIndividualInformationCmd)
}

// OperatorGetIndividualInformationCmd defines 'get-individual-information' subcommand
var OperatorGetIndividualInformationCmd = &cobra.Command{
	Use:   "get-individual-information",
	Short: TRAPI("/operators/{operator_id}/individual_information:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/individual_information:get:description`),
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

		param, err := collectOperatorGetIndividualInformationCmdParams(ac)
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

func collectOperatorGetIndividualInformationCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorGetIndividualInformationCmdOperatorId == "" {
		OperatorGetIndividualInformationCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForOperatorGetIndividualInformationCmd("/operators/{operator_id}/individual_information"),
		query:  buildQueryForOperatorGetIndividualInformationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorGetIndividualInformationCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorGetIndividualInformationCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorGetIndividualInformationCmd() url.Values {
	result := url.Values{}

	return result
}