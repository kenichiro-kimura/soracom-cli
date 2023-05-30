// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// BillsGetCmdYyyyMM holds value of 'yyyyMM' option
var BillsGetCmdYyyyMM string

func InitBillsGetCmd() {
	BillsGetCmd.Flags().StringVar(&BillsGetCmdYyyyMM, "yyyy-mm", "", TRAPI("Target year and month"))

	BillsGetCmd.RunE = BillsGetCmdRunE

	BillsCmd.AddCommand(BillsGetCmd)
}

// BillsGetCmd defines 'get' subcommand
var BillsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/bills/{yyyyMM}:get:summary"),
	Long:  TRAPI(`/bills/{yyyyMM}:get:description`) + "\n\n" + createLinkToAPIReference("Billing", "getBilling"),
}

func BillsGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectBillsGetCmdParams(ac)
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

func collectBillsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("yyyyMM", "yyyy-mm", "path", parsedBody, BillsGetCmdYyyyMM)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForBillsGetCmd("/bills/{yyyyMM}"),
		query:  buildQueryForBillsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForBillsGetCmd(path string) string {

	escapedYyyyMM := url.PathEscape(BillsGetCmdYyyyMM)

	path = strReplace(path, "{"+"yyyyMM"+"}", escapedYyyyMM, -1)

	return path
}

func buildQueryForBillsGetCmd() url.Values {
	result := url.Values{}

	return result
}
