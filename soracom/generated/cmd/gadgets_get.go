// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GadgetsGetCmdProductId holds value of 'product_id' option
var GadgetsGetCmdProductId string

// GadgetsGetCmdSerialNumber holds value of 'serial_number' option
var GadgetsGetCmdSerialNumber string

func init() {
	GadgetsGetCmd.Flags().StringVar(&GadgetsGetCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsGetCmd.Flags().StringVar(&GadgetsGetCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))
	GadgetsCmd.AddCommand(GadgetsGetCmd)
}

// GadgetsGetCmd defines 'get' subcommand
var GadgetsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}:get:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}:get:description`),
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

		param, err := collectGadgetsGetCmdParams(ac)
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

func collectGadgetsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("product_id", "product-id", "path", parsedBody, GadgetsGetCmdProductId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("serial_number", "serial-number", "path", parsedBody, GadgetsGetCmdSerialNumber)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForGadgetsGetCmd("/gadgets/{product_id}/{serial_number}"),
		query:  buildQueryForGadgetsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGadgetsGetCmd(path string) string {

	escapedProductId := url.PathEscape(GadgetsGetCmdProductId)

	path = strReplace(path, "{"+"product_id"+"}", escapedProductId, -1)

	escapedSerialNumber := url.PathEscape(GadgetsGetCmdSerialNumber)

	path = strReplace(path, "{"+"serial_number"+"}", escapedSerialNumber, -1)

	return path
}

func buildQueryForGadgetsGetCmd() url.Values {
	result := url.Values{}

	return result
}
