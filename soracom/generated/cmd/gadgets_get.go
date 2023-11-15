// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GadgetsGetCmdProductId holds value of 'product_id' option
var GadgetsGetCmdProductId string

// GadgetsGetCmdSerialNumber holds value of 'serial_number' option
var GadgetsGetCmdSerialNumber string

func InitGadgetsGetCmd() {
	GadgetsGetCmd.Flags().StringVar(&GadgetsGetCmdProductId, "product-id", "", TRAPI("Product ID of the target Gadget API compatible device.- 'button': Soracom LTE-M Button powered by AWS.- 'wimax': Soracom Cloud Camera Services Cellular Pack."))

	GadgetsGetCmd.Flags().StringVar(&GadgetsGetCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target Gadget API compatible device."))

	GadgetsGetCmd.RunE = GadgetsGetCmdRunE

	GadgetsCmd.AddCommand(GadgetsGetCmd)
}

// GadgetsGetCmd defines 'get' subcommand
var GadgetsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}:get:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}:get:description`) + "\n\n" + createLinkToAPIReference("Gadget", "getGadget"),
}

func GadgetsGetCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
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
