// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// ShippingAddressesGetCmdOperatorId holds value of 'operator_id' option
var ShippingAddressesGetCmdOperatorId string

// ShippingAddressesGetCmdShippingAddressId holds value of 'shipping_address_id' option
var ShippingAddressesGetCmdShippingAddressId string

func init() {
	ShippingAddressesGetCmd.Flags().StringVar(&ShippingAddressesGetCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	ShippingAddressesGetCmd.Flags().StringVar(&ShippingAddressesGetCmdShippingAddressId, "shipping-address-id", "", TRAPI("shipping_address_id"))
	ShippingAddressesCmd.AddCommand(ShippingAddressesGetCmd)
}

// ShippingAddressesGetCmd defines 'get' subcommand
var ShippingAddressesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/shipping_addresses/{shipping_address_id}:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/shipping_addresses/{shipping_address_id}:get:description`) + "\n\n" + createLinkToAPIReference("ShippingAddress", "getShippingAddress"),
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) > 0 {
			return fmt.Errorf("unexpected arguments passed => %v", args)
		}

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

		param, err := collectShippingAddressesGetCmdParams(ac)
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

func collectShippingAddressesGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if ShippingAddressesGetCmdOperatorId == "" {
		ShippingAddressesGetCmdOperatorId = ac.OperatorID
	}

	err = checkIfRequiredStringParameterIsSupplied("shipping_address_id", "shipping-address-id", "path", parsedBody, ShippingAddressesGetCmdShippingAddressId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForShippingAddressesGetCmd("/operators/{operator_id}/shipping_addresses/{shipping_address_id}"),
		query:  buildQueryForShippingAddressesGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForShippingAddressesGetCmd(path string) string {

	escapedOperatorId := url.PathEscape(ShippingAddressesGetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedShippingAddressId := url.PathEscape(ShippingAddressesGetCmdShippingAddressId)

	path = strReplace(path, "{"+"shipping_address_id"+"}", escapedShippingAddressId, -1)

	return path
}

func buildQueryForShippingAddressesGetCmd() url.Values {
	result := url.Values{}

	return result
}
