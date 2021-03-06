package cmd

import (
	"os"
	"strings"

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

		_, body, err := ac.callAPI(param)
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

func collectGadgetsGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForGadgetsGetCmd("/gadgets/{product_id}/{serial_number}"),
		query:  buildQueryForGadgetsGetCmd(),
	}, nil
}

func buildPathForGadgetsGetCmd(path string) string {

	path = strings.Replace(path, "{"+"product_id"+"}", GadgetsGetCmdProductId, -1)

	path = strings.Replace(path, "{"+"serial_number"+"}", GadgetsGetCmdSerialNumber, -1)

	return path
}

func buildQueryForGadgetsGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
