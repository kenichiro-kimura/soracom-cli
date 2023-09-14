// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

func InitBillsGetLatestCmd() {

	BillsGetLatestCmd.RunE = BillsGetLatestCmdRunE

	BillsCmd.AddCommand(BillsGetLatestCmd)
}

// BillsGetLatestCmd defines 'get-latest' subcommand
var BillsGetLatestCmd = &cobra.Command{
	Use:   "get-latest",
	Short: TRAPI("/bills/latest:get:summary"),
	Long:  TRAPI(`/bills/latest:get:description`) + "\n\n" + createLinkToAPIReference("Billing", "getLatestBilling"),
}

func BillsGetLatestCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectBillsGetLatestCmdParams(ac)
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

func collectBillsGetLatestCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForBillsGetLatestCmd("/bills/latest"),
		query:  buildQueryForBillsGetLatestCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForBillsGetLatestCmd(path string) string {

	return path
}

func buildQueryForBillsGetLatestCmd() url.Values {
	result := url.Values{}

	return result
}
