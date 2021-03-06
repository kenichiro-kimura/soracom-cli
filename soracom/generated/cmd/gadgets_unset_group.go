package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GadgetsUnsetGroupCmdProductId holds value of 'product_id' option
var GadgetsUnsetGroupCmdProductId string

// GadgetsUnsetGroupCmdSerialNumber holds value of 'serial_number' option
var GadgetsUnsetGroupCmdSerialNumber string

func init() {
	GadgetsUnsetGroupCmd.Flags().StringVar(&GadgetsUnsetGroupCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsUnsetGroupCmd.Flags().StringVar(&GadgetsUnsetGroupCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))

	GadgetsCmd.AddCommand(GadgetsUnsetGroupCmd)
}

// GadgetsUnsetGroupCmd defines 'unset-group' subcommand
var GadgetsUnsetGroupCmd = &cobra.Command{
	Use:   "unset-group",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/unset_group:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/unset_group:post:description`),
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

		param, err := collectGadgetsUnsetGroupCmdParams(ac)
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

func collectGadgetsUnsetGroupCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForGadgetsUnsetGroupCmd("/gadgets/{product_id}/{serial_number}/unset_group"),
		query:  buildQueryForGadgetsUnsetGroupCmd(),
	}, nil
}

func buildPathForGadgetsUnsetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"product_id"+"}", GadgetsUnsetGroupCmdProductId, -1)

	path = strings.Replace(path, "{"+"serial_number"+"}", GadgetsUnsetGroupCmdSerialNumber, -1)

	return path
}

func buildQueryForGadgetsUnsetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
