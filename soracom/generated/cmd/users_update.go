package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersUpdateCmdDescription holds value of 'description' option
var UsersUpdateCmdDescription string

// UsersUpdateCmdOperatorId holds value of 'operator_id' option
var UsersUpdateCmdOperatorId string

// UsersUpdateCmdUserName holds value of 'user_name' option
var UsersUpdateCmdUserName string

// UsersUpdateCmdBody holds contents of request body to be sent
var UsersUpdateCmdBody string

func init() {
	UsersUpdateCmd.Flags().StringVar(&UsersUpdateCmdDescription, "description", "", TR(""))

	UsersUpdateCmd.Flags().StringVar(&UsersUpdateCmdOperatorId, "operator-id", "", TR("operator_id"))

	UsersUpdateCmd.Flags().StringVar(&UsersUpdateCmdUserName, "user-name", "", TR("user_name"))

	UsersUpdateCmd.Flags().StringVar(&UsersUpdateCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	UsersCmd.AddCommand(UsersUpdateCmd)
}

// UsersUpdateCmd defines 'update' subcommand
var UsersUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TR("users.update_user.put.summary"),
	Long:  TR(`users.update_user.put.description`),
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

		param, err := collectUsersUpdateCmdParams()
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

func collectUsersUpdateCmdParams() (*apiParams, error) {

	body, err := buildBodyForUsersUpdateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForUsersUpdateCmd("/operators/{operator_id}/users/{user_name}"),
		query:       buildQueryForUsersUpdateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForUsersUpdateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersUpdateCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersUpdateCmdUserName, -1)

	return path
}

func buildQueryForUsersUpdateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForUsersUpdateCmd() (string, error) {
	if UsersUpdateCmdBody != "" {
		if strings.HasPrefix(UsersUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(UsersUpdateCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if UsersUpdateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return UsersUpdateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if UsersUpdateCmdDescription != "" {
		result["description"] = UsersUpdateCmdDescription
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
