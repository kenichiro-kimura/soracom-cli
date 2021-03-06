package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OrdersRegisterSubscribersCmdOrderId holds value of 'order_id' option
var OrdersRegisterSubscribersCmdOrderId string

func init() {
	OrdersRegisterSubscribersCmd.Flags().StringVar(&OrdersRegisterSubscribersCmdOrderId, "order-id", "", TRAPI("order_id"))

	OrdersCmd.AddCommand(OrdersRegisterSubscribersCmd)
}

// OrdersRegisterSubscribersCmd defines 'register-subscribers' subcommand
var OrdersRegisterSubscribersCmd = &cobra.Command{
	Use:   "register-subscribers",
	Short: TRAPI("/orders/{order_id}/subscribers/register:post:summary"),
	Long:  TRAPI(`/orders/{order_id}/subscribers/register:post:description`),
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

		param, err := collectOrdersRegisterSubscribersCmdParams(ac)
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

func collectOrdersRegisterSubscribersCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForOrdersRegisterSubscribersCmd("/orders/{order_id}/subscribers/register"),
		query:  buildQueryForOrdersRegisterSubscribersCmd(),
	}, nil
}

func buildPathForOrdersRegisterSubscribersCmd(path string) string {

	path = strings.Replace(path, "{"+"order_id"+"}", OrdersRegisterSubscribersCmdOrderId, -1)

	return path
}

func buildQueryForOrdersRegisterSubscribersCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
