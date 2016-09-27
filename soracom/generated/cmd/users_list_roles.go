package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersListRolesCmdOperatorId holds value of 'operator_id' option
var UsersListRolesCmdOperatorId string

// UsersListRolesCmdUserName holds value of 'user_name' option
var UsersListRolesCmdUserName string

func init() {
	UsersListRolesCmd.Flags().StringVar(&UsersListRolesCmdOperatorId, "operator-id", "", TR("operator_id"))

	UsersListRolesCmd.Flags().StringVar(&UsersListRolesCmdUserName, "user-name", "", TR("user_name"))

	UsersCmd.AddCommand(UsersListRolesCmd)
}

// UsersListRolesCmd defines 'list-roles' subcommand
var UsersListRolesCmd = &cobra.Command{
	Use:   "list-roles",
	Short: TR("roles.list_user_roles.get.summary"),
	Long:  TR(`roles.list_user_roles.get.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		param, err := collectUsersListRolesCmdParams()
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

func collectUsersListRolesCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersListRolesCmd("/operators/{operator_id}/users/{user_name}/roles"),
		query:  buildQueryForUsersListRolesCmd(),
	}, nil
}

func buildPathForUsersListRolesCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersListRolesCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersListRolesCmdUserName, -1)

	return path
}

func buildQueryForUsersListRolesCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
