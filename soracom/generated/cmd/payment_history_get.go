// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// PaymentHistoryGetCmdPaymentTransactionId holds value of 'payment_transaction_id' option
var PaymentHistoryGetCmdPaymentTransactionId string

func InitPaymentHistoryGetCmd() {
	PaymentHistoryGetCmd.Flags().StringVar(&PaymentHistoryGetCmdPaymentTransactionId, "payment-transaction-id", "", TRAPI("Payment transaction ID. It can be obtained via [Billing:getBillingHistory API](#/Billing/getBillingHistory) or [Billing:getBilling API](#/Billing/getBilling)."))

	PaymentHistoryGetCmd.RunE = PaymentHistoryGetCmdRunE

	PaymentHistoryCmd.AddCommand(PaymentHistoryGetCmd)
}

// PaymentHistoryGetCmd defines 'get' subcommand
var PaymentHistoryGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/payment_history/transactions/{payment_transaction_id}:get:summary"),
	Long:  TRAPI(`/payment_history/transactions/{payment_transaction_id}:get:description`) + "\n\n" + createLinkToAPIReference("Payment", "getPaymentTransaction"),
}

func PaymentHistoryGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectPaymentHistoryGetCmdParams(ac)
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

func collectPaymentHistoryGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("payment_transaction_id", "payment-transaction-id", "path", parsedBody, PaymentHistoryGetCmdPaymentTransactionId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForPaymentHistoryGetCmd("/payment_history/transactions/{payment_transaction_id}"),
		query:  buildQueryForPaymentHistoryGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForPaymentHistoryGetCmd(path string) string {

	escapedPaymentTransactionId := url.PathEscape(PaymentHistoryGetCmdPaymentTransactionId)

	path = strReplace(path, "{"+"payment_transaction_id"+"}", escapedPaymentTransactionId, -1)

	return path
}

func buildQueryForPaymentHistoryGetCmd() url.Values {
	result := url.Values{}

	return result
}
