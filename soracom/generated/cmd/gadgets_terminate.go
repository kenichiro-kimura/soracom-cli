package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GadgetsTerminateCmdProductId holds value of 'product_id' option
var GadgetsTerminateCmdProductId string

// GadgetsTerminateCmdSerialNumber holds value of 'serial_number' option
var GadgetsTerminateCmdSerialNumber string

func init() {
	GadgetsTerminateCmd.Flags().StringVar(&GadgetsTerminateCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsTerminateCmd.Flags().StringVar(&GadgetsTerminateCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))

	GadgetsCmd.AddCommand(GadgetsTerminateCmd)
}

// GadgetsTerminateCmd defines 'terminate' subcommand
var GadgetsTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/terminate:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/terminate:post:description`),
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

		param, err := collectGadgetsTerminateCmdParams(ac)
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

func collectGadgetsTerminateCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForGadgetsTerminateCmd("/gadgets/{product_id}/{serial_number}/terminate"),
		query:  buildQueryForGadgetsTerminateCmd(),
	}, nil
}

func buildPathForGadgetsTerminateCmd(path string) string {

	path = strings.Replace(path, "{"+"product_id"+"}", GadgetsTerminateCmdProductId, -1)

	path = strings.Replace(path, "{"+"serial_number"+"}", GadgetsTerminateCmdSerialNumber, -1)

	return path
}

func buildQueryForGadgetsTerminateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
