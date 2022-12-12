// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OrdersListSubscribersCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var OrdersListSubscribersCmdLastEvaluatedKey string

// OrdersListSubscribersCmdOrderId holds value of 'order_id' option
var OrdersListSubscribersCmdOrderId string

// OrdersListSubscribersCmdLimit holds value of 'limit' option
var OrdersListSubscribersCmdLimit int64

// OrdersListSubscribersCmdPaginate indicates to do pagination or not
var OrdersListSubscribersCmdPaginate bool

func init() {
	OrdersListSubscribersCmd.Flags().StringVar(&OrdersListSubscribersCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("Serial number of the last subscriber in the previous page that is set to response header with X-Soracom-Next-Key."))

	OrdersListSubscribersCmd.Flags().StringVar(&OrdersListSubscribersCmdOrderId, "order-id", "", TRAPI("Order ID. You can get it by calling [`Order:listOrders API`](#/Order/listOrders)."))

	OrdersListSubscribersCmd.Flags().Int64Var(&OrdersListSubscribersCmdLimit, "limit", 0, TRAPI("Max number of subscribers in a response."))

	OrdersListSubscribersCmd.Flags().BoolVar(&OrdersListSubscribersCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	OrdersCmd.AddCommand(OrdersListSubscribersCmd)
}

// OrdersListSubscribersCmd defines 'list-subscribers' subcommand
var OrdersListSubscribersCmd = &cobra.Command{
	Use:   "list-subscribers",
	Short: TRAPI("/orders/{order_id}/subscribers:get:summary"),
	Long:  TRAPI(`/orders/{order_id}/subscribers:get:description`) + "\n\n" + createLinkToAPIReference("Order", "listOrderedSubscribers"),
	RunE: func(cmd *cobra.Command, args []string) error {

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

		param, err := collectOrdersListSubscribersCmdParams(ac)
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
	},
}

func collectOrdersListSubscribersCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error

	err = checkIfRequiredStringParameterIsSupplied("order_id", "order-id", "path", parsedBody, OrdersListSubscribersCmdOrderId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForOrdersListSubscribersCmd("/orders/{order_id}/subscribers"),
		query:  buildQueryForOrdersListSubscribersCmd(),

		doPagination:                      OrdersListSubscribersCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOrdersListSubscribersCmd(path string) string {

	escapedOrderId := url.PathEscape(OrdersListSubscribersCmdOrderId)

	path = strReplace(path, "{"+"order_id"+"}", escapedOrderId, -1)

	return path
}

func buildQueryForOrdersListSubscribersCmd() url.Values {
	result := url.Values{}

	if OrdersListSubscribersCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", OrdersListSubscribersCmdLastEvaluatedKey)
	}

	if OrdersListSubscribersCmdLimit != 0 {
		result.Add("limit", sprintf("%d", OrdersListSubscribersCmdLimit))
	}

	return result
}
