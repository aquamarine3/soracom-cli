package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersSetGroupCmdGroupId holds value of 'groupId' option
var SubscribersSetGroupCmdGroupId string

// SubscribersSetGroupCmdImsi holds value of 'imsi' option
var SubscribersSetGroupCmdImsi string

// SubscribersSetGroupCmdOperatorId holds value of 'operatorId' option
var SubscribersSetGroupCmdOperatorId string

// SubscribersSetGroupCmdCreatedTime holds value of 'createdTime' option
var SubscribersSetGroupCmdCreatedTime int64

// SubscribersSetGroupCmdLastModifiedTime holds value of 'lastModifiedTime' option
var SubscribersSetGroupCmdLastModifiedTime int64

// SubscribersSetGroupCmdBody holds contents of request body to be sent
var SubscribersSetGroupCmdBody string

func init() {
	SubscribersSetGroupCmd.Flags().StringVar(&SubscribersSetGroupCmdGroupId, "group-id", "", TR(""))

	SubscribersSetGroupCmd.Flags().StringVar(&SubscribersSetGroupCmdImsi, "imsi", "", TR("subscribers.set_group.post.parameters.imsi.description"))

	SubscribersSetGroupCmd.Flags().StringVar(&SubscribersSetGroupCmdOperatorId, "operator-id", "", TR(""))

	SubscribersSetGroupCmd.Flags().Int64Var(&SubscribersSetGroupCmdCreatedTime, "created-time", 0, TR(""))

	SubscribersSetGroupCmd.Flags().Int64Var(&SubscribersSetGroupCmdLastModifiedTime, "last-modified-time", 0, TR(""))

	SubscribersSetGroupCmd.Flags().StringVar(&SubscribersSetGroupCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	SubscribersCmd.AddCommand(SubscribersSetGroupCmd)
}

// SubscribersSetGroupCmd defines 'set-group' subcommand
var SubscribersSetGroupCmd = &cobra.Command{
	Use:   "set-group",
	Short: TR("subscribers.set_group.post.summary"),
	Long:  TR(`subscribers.set_group.post.description`),
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

		param, err := collectSubscribersSetGroupCmdParams()
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

func collectSubscribersSetGroupCmdParams() (*apiParams, error) {

	body, err := buildBodyForSubscribersSetGroupCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersSetGroupCmd("/subscribers/{imsi}/set_group"),
		query:       buildQueryForSubscribersSetGroupCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSubscribersSetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersSetGroupCmdImsi, -1)

	return path
}

func buildQueryForSubscribersSetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSubscribersSetGroupCmd() (string, error) {
	if SubscribersSetGroupCmdBody != "" {
		if strings.HasPrefix(SubscribersSetGroupCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersSetGroupCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if SubscribersSetGroupCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return SubscribersSetGroupCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if SubscribersSetGroupCmdGroupId != "" {
		result["groupId"] = SubscribersSetGroupCmdGroupId
	}

	if SubscribersSetGroupCmdOperatorId != "" {
		result["operatorId"] = SubscribersSetGroupCmdOperatorId
	}

	if SubscribersSetGroupCmdCreatedTime != 0 {
		result["createdTime"] = SubscribersSetGroupCmdCreatedTime
	}

	if SubscribersSetGroupCmdLastModifiedTime != 0 {
		result["lastModifiedTime"] = SubscribersSetGroupCmdLastModifiedTime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
