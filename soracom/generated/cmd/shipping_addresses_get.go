package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ShippingAddressesGetCmdOperatorId holds value of 'operator_id' option
var ShippingAddressesGetCmdOperatorId string

// ShippingAddressesGetCmdShippingAddressId holds value of 'shipping_address_id' option
var ShippingAddressesGetCmdShippingAddressId string

func init() {
	ShippingAddressesGetCmd.Flags().StringVar(&ShippingAddressesGetCmdOperatorId, "operator-id", "", TR("operator_id"))

	ShippingAddressesGetCmd.Flags().StringVar(&ShippingAddressesGetCmdShippingAddressId, "shipping-address-id", "", TR("shipping_address_id"))

	ShippingAddressesCmd.AddCommand(ShippingAddressesGetCmd)
}

// ShippingAddressesGetCmd defines 'get' subcommand
var ShippingAddressesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TR("shipping_addresses.get_shipping_address.get.summary"),
	Long:  TR(`shipping_addresses.get_shipping_address.get.description`),
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

		param, err := collectShippingAddressesGetCmdParams()
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

func collectShippingAddressesGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForShippingAddressesGetCmd("/operators/{operator_id}/shipping_addresses/{shipping_address_id}"),
		query:  buildQueryForShippingAddressesGetCmd(),
	}, nil
}

func buildPathForShippingAddressesGetCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", ShippingAddressesGetCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"shipping_address_id"+"}", ShippingAddressesGetCmdShippingAddressId, -1)

	return path
}

func buildQueryForShippingAddressesGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
