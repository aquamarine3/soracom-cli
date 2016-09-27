package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// EventHandlersDeleteCmdHandlerId holds value of 'handler_id' option
var EventHandlersDeleteCmdHandlerId string

func init() {
	EventHandlersDeleteCmd.Flags().StringVar(&EventHandlersDeleteCmdHandlerId, "handler-id", "", TR("event_handlers.delete_event_handler.delete.parameters.handler_id.description"))

	EventHandlersCmd.AddCommand(EventHandlersDeleteCmd)
}

// EventHandlersDeleteCmd defines 'delete' subcommand
var EventHandlersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TR("event_handlers.delete_event_handler.delete.summary"),
	Long:  TR(`event_handlers.delete_event_handler.delete.description`),
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

		param, err := collectEventHandlersDeleteCmdParams()
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

func collectEventHandlersDeleteCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForEventHandlersDeleteCmd("/event_handlers/{handler_id}"),
		query:  buildQueryForEventHandlersDeleteCmd(),
	}, nil
}

func buildPathForEventHandlersDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"handler_id"+"}", EventHandlersDeleteCmdHandlerId, -1)

	return path
}

func buildQueryForEventHandlersDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
