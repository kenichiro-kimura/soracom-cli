// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

func InitBillsListCmd() {

	BillsListCmd.RunE = BillsListCmdRunE

	BillsCmd.AddCommand(BillsListCmd)
}

// BillsListCmd defines 'list' subcommand
var BillsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/bills:get:summary"),
	Long:  TRAPI(`/bills:get:description`) + "\n\n" + createLinkToAPIReference("Billing", "getBillingHistory"),
}

func BillsListCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectBillsListCmdParams(ac)
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

func collectBillsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForBillsListCmd("/bills"),
		query:  buildQueryForBillsListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForBillsListCmd(path string) string {

	return path
}

func buildQueryForBillsListCmd() url.Values {
	result := url.Values{}

	return result
}
