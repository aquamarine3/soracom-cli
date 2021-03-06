// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsSetToStandbyCmdSimId holds value of 'sim_id' option
var SimsSetToStandbyCmdSimId string

func init() {
	SimsSetToStandbyCmd.Flags().StringVar(&SimsSetToStandbyCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))
	SimsCmd.AddCommand(SimsSetToStandbyCmd)
}

// SimsSetToStandbyCmd defines 'set-to-standby' subcommand
var SimsSetToStandbyCmd = &cobra.Command{
	Use:   "set-to-standby",
	Short: TRAPI("/sims/{sim_id}/set_to_standby:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/set_to_standby:post:description`),
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

		param, err := collectSimsSetToStandbyCmdParams(ac)
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

func collectSimsSetToStandbyCmdParams(ac *apiClient) (*apiParams, error) {
	if SimsSetToStandbyCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsSetToStandbyCmd("/sims/{sim_id}/set_to_standby"),
		query:  buildQueryForSimsSetToStandbyCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsSetToStandbyCmd(path string) string {

	escapedSimId := url.PathEscape(SimsSetToStandbyCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsSetToStandbyCmd() url.Values {
	result := url.Values{}

	return result
}
