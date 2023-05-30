// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// PaymentStatementsExportCmdExportMode holds value of 'export_mode' option
var PaymentStatementsExportCmdExportMode string

// PaymentStatementsExportCmdPaymentStatementId holds value of 'payment_statement_id' option
var PaymentStatementsExportCmdPaymentStatementId string

func InitPaymentStatementsExportCmd() {
	PaymentStatementsExportCmd.Flags().StringVar(&PaymentStatementsExportCmdExportMode, "export-mode", "", TRAPI("Export mode (async, sync)"))

	PaymentStatementsExportCmd.Flags().StringVar(&PaymentStatementsExportCmdPaymentStatementId, "payment-statement-id", "", TRAPI("Payment statement ID"))

	PaymentStatementsExportCmd.RunE = PaymentStatementsExportCmdRunE

	PaymentStatementsCmd.AddCommand(PaymentStatementsExportCmd)
}

// PaymentStatementsExportCmd defines 'export' subcommand
var PaymentStatementsExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/payment_statements/{payment_statement_id}/export:post:summary"),
	Long:  TRAPI(`/payment_statements/{payment_statement_id}/export:post:description`) + "\n\n" + createLinkToAPIReference("Payment", "exportPaymentStatement"),
}

func PaymentStatementsExportCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectPaymentStatementsExportCmdParams(ac)
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
	rawOutput = true

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectPaymentStatementsExportCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error

	err = checkIfRequiredStringParameterIsSupplied("payment_statement_id", "payment-statement-id", "path", parsedBody, PaymentStatementsExportCmdPaymentStatementId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForPaymentStatementsExportCmd("/payment_statements/{payment_statement_id}/export"),
		query:  buildQueryForPaymentStatementsExportCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForPaymentStatementsExportCmd(path string) string {

	escapedPaymentStatementId := url.PathEscape(PaymentStatementsExportCmdPaymentStatementId)

	path = strReplace(path, "{"+"payment_statement_id"+"}", escapedPaymentStatementId, -1)

	return path
}

func buildQueryForPaymentStatementsExportCmd() url.Values {
	result := url.Values{}

	if PaymentStatementsExportCmdExportMode != "" {
		result.Add("export_mode", PaymentStatementsExportCmdExportMode)
	}

	return result
}
