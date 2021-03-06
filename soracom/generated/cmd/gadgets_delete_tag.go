package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GadgetsDeleteTagCmdProductId holds value of 'product_id' option
var GadgetsDeleteTagCmdProductId string

// GadgetsDeleteTagCmdSerialNumber holds value of 'serial_number' option
var GadgetsDeleteTagCmdSerialNumber string

// GadgetsDeleteTagCmdTagName holds value of 'tag_name' option
var GadgetsDeleteTagCmdTagName string

func init() {
	GadgetsDeleteTagCmd.Flags().StringVar(&GadgetsDeleteTagCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsDeleteTagCmd.Flags().StringVar(&GadgetsDeleteTagCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))

	GadgetsDeleteTagCmd.Flags().StringVar(&GadgetsDeleteTagCmdTagName, "tag-name", "", TRAPI("Tag name to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))

	GadgetsCmd.AddCommand(GadgetsDeleteTagCmd)
}

// GadgetsDeleteTagCmd defines 'delete-tag' subcommand
var GadgetsDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/tags/{tag_name}:delete:description`),
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

		param, err := collectGadgetsDeleteTagCmdParams(ac)
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

func collectGadgetsDeleteTagCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForGadgetsDeleteTagCmd("/gadgets/{product_id}/{serial_number}/tags/{tag_name}"),
		query:  buildQueryForGadgetsDeleteTagCmd(),
	}, nil
}

func buildPathForGadgetsDeleteTagCmd(path string) string {

	path = strings.Replace(path, "{"+"product_id"+"}", GadgetsDeleteTagCmdProductId, -1)

	path = strings.Replace(path, "{"+"serial_number"+"}", GadgetsDeleteTagCmdSerialNumber, -1)

	path = strings.Replace(path, "{"+"tag_name"+"}", GadgetsDeleteTagCmdTagName, -1)

	return path
}

func buildQueryForGadgetsDeleteTagCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
