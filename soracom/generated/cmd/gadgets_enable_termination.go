package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GadgetsEnableTerminationCmdProductId holds value of 'product_id' option
var GadgetsEnableTerminationCmdProductId string

// GadgetsEnableTerminationCmdSerialNumber holds value of 'serial_number' option
var GadgetsEnableTerminationCmdSerialNumber string

func init() {
	GadgetsEnableTerminationCmd.Flags().StringVar(&GadgetsEnableTerminationCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsEnableTerminationCmd.Flags().StringVar(&GadgetsEnableTerminationCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))

	GadgetsCmd.AddCommand(GadgetsEnableTerminationCmd)
}

// GadgetsEnableTerminationCmd defines 'enable-termination' subcommand
var GadgetsEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/enable_termination:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/enable_termination:post:description`),
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

		param, err := collectGadgetsEnableTerminationCmdParams(ac)
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

func collectGadgetsEnableTerminationCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForGadgetsEnableTerminationCmd("/gadgets/{product_id}/{serial_number}/enable_termination"),
		query:  buildQueryForGadgetsEnableTerminationCmd(),
	}, nil
}

func buildPathForGadgetsEnableTerminationCmd(path string) string {

	path = strings.Replace(path, "{"+"product_id"+"}", GadgetsEnableTerminationCmdProductId, -1)

	path = strings.Replace(path, "{"+"serial_number"+"}", GadgetsEnableTerminationCmdSerialNumber, -1)

	return path
}

func buildQueryForGadgetsEnableTerminationCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
