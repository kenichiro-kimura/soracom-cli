// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraGatewaysListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var LoraGatewaysListCmdLastEvaluatedKey string

// LoraGatewaysListCmdTagName holds value of 'tag_name' option
var LoraGatewaysListCmdTagName string

// LoraGatewaysListCmdTagValue holds value of 'tag_value' option
var LoraGatewaysListCmdTagValue string

// LoraGatewaysListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var LoraGatewaysListCmdTagValueMatchMode string

// LoraGatewaysListCmdLimit holds value of 'limit' option
var LoraGatewaysListCmdLimit int64

// LoraGatewaysListCmdPaginate indicates to do pagination or not
var LoraGatewaysListCmdPaginate bool

func init() {
	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The device ID of the last device retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next device onward."))

	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdTagName, "tag-name", "", TRAPI("Tag name for filtering the search (exact match)."))

	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdTagValue, "tag-value", "", TRAPI("Tag search string for filtering the search. Required when `tag_name` has been specified."))

	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdTagValueMatchMode, "tag-value-match-mode", "exact", TRAPI("Tag match mode."))

	LoraGatewaysListCmd.Flags().Int64Var(&LoraGatewaysListCmdLimit, "limit", 0, TRAPI("Maximum number of LoRa devices to retrieve."))

	LoraGatewaysListCmd.Flags().BoolVar(&LoraGatewaysListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	LoraGatewaysCmd.AddCommand(LoraGatewaysListCmd)
}

// LoraGatewaysListCmd defines 'list' subcommand
var LoraGatewaysListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/lora_gateways:get:summary"),
	Long:  TRAPI(`/lora_gateways:get:description`),
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

		param, err := collectLoraGatewaysListCmdParams(ac)
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

func collectLoraGatewaysListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraGatewaysListCmd("/lora_gateways"),
		query:  buildQueryForLoraGatewaysListCmd(),

		doPagination:                      LoraGatewaysListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraGatewaysListCmd(path string) string {

	return path
}

func buildQueryForLoraGatewaysListCmd() url.Values {
	result := url.Values{}

	if LoraGatewaysListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", LoraGatewaysListCmdLastEvaluatedKey)
	}

	if LoraGatewaysListCmdTagName != "" {
		result.Add("tag_name", LoraGatewaysListCmdTagName)
	}

	if LoraGatewaysListCmdTagValue != "" {
		result.Add("tag_value", LoraGatewaysListCmdTagValue)
	}

	if LoraGatewaysListCmdTagValueMatchMode != "exact" {
		result.Add("tag_value_match_mode", LoraGatewaysListCmdTagValueMatchMode)
	}

	if LoraGatewaysListCmdLimit != 0 {
		result.Add("limit", sprintf("%d", LoraGatewaysListCmdLimit))
	}

	return result
}
