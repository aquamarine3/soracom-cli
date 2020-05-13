// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersAuthKeysGenerateCmdOperatorId holds value of 'operator_id' option
var UsersAuthKeysGenerateCmdOperatorId string

// UsersAuthKeysGenerateCmdUserName holds value of 'user_name' option
var UsersAuthKeysGenerateCmdUserName string

func init() {
	UsersAuthKeysGenerateCmd.Flags().StringVar(&UsersAuthKeysGenerateCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersAuthKeysGenerateCmd.Flags().StringVar(&UsersAuthKeysGenerateCmdUserName, "user-name", "", TRAPI("user_name"))
	UsersAuthKeysCmd.AddCommand(UsersAuthKeysGenerateCmd)
}

// UsersAuthKeysGenerateCmd defines 'generate' subcommand
var UsersAuthKeysGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/auth_keys:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/auth_keys:post:description`),
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

		param, err := collectUsersAuthKeysGenerateCmdParams(ac)
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
		return prettyPrintStringAsJSON(body)

	},
}

func collectUsersAuthKeysGenerateCmdParams(ac *apiClient) (*apiParams, error) {
	if UsersAuthKeysGenerateCmdOperatorId == "" {
		UsersAuthKeysGenerateCmdOperatorId = ac.OperatorID
	}

	if UsersAuthKeysGenerateCmdUserName == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "user-name")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForUsersAuthKeysGenerateCmd("/operators/{operator_id}/users/{user_name}/auth_keys"),
		query:  buildQueryForUsersAuthKeysGenerateCmd(),
	}, nil
}

func buildPathForUsersAuthKeysGenerateCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersAuthKeysGenerateCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersAuthKeysGenerateCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersAuthKeysGenerateCmd() url.Values {
	result := url.Values{}

	return result
}
