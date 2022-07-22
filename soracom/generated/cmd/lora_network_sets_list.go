// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var LoraNetworkSetsListCmdLastEvaluatedKey string

// LoraNetworkSetsListCmdTagName holds value of 'tag_name' option
var LoraNetworkSetsListCmdTagName string

// LoraNetworkSetsListCmdTagValue holds value of 'tag_value' option
var LoraNetworkSetsListCmdTagValue string

// LoraNetworkSetsListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var LoraNetworkSetsListCmdTagValueMatchMode string

// LoraNetworkSetsListCmdLimit holds value of 'limit' option
var LoraNetworkSetsListCmdLimit int64

// LoraNetworkSetsListCmdPaginate indicates to do pagination or not
var LoraNetworkSetsListCmdPaginate bool

// LoraNetworkSetsListCmdOutputJSONL indicates to output with jsonl format
var LoraNetworkSetsListCmdOutputJSONL bool

func init() {
	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The ID of the last network set retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next device onward."))

	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdTagName, "tag-name", "", TRAPI("Tag name for filtering the search (exact match)."))

	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdTagValue, "tag-value", "", TRAPI("Tag search string for filtering the search. Required when `tag_name` has been specified."))

	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdTagValueMatchMode, "tag-value-match-mode", "exact", TRAPI("Tag match mode."))

	LoraNetworkSetsListCmd.Flags().Int64Var(&LoraNetworkSetsListCmdLimit, "limit", 0, TRAPI("Maximum number of LoRa devices to retrieve."))

	LoraNetworkSetsListCmd.Flags().BoolVar(&LoraNetworkSetsListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	LoraNetworkSetsListCmd.Flags().BoolVar(&LoraNetworkSetsListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsListCmd)
}

// LoraNetworkSetsListCmd defines 'list' subcommand
var LoraNetworkSetsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/lora_network_sets:get:summary"),
	Long:  TRAPI(`/lora_network_sets:get:description`) + "\n\n" + createLinkToAPIReference("LoraNetworkSet", "listLoraNetworkSets"),
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

		param, err := collectLoraNetworkSetsListCmdParams(ac)
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
			if LoraNetworkSetsListCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectLoraNetworkSetsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraNetworkSetsListCmd("/lora_network_sets"),
		query:  buildQueryForLoraNetworkSetsListCmd(),

		doPagination:                      LoraNetworkSetsListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraNetworkSetsListCmd(path string) string {

	return path
}

func buildQueryForLoraNetworkSetsListCmd() url.Values {
	result := url.Values{}

	if LoraNetworkSetsListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", LoraNetworkSetsListCmdLastEvaluatedKey)
	}

	if LoraNetworkSetsListCmdTagName != "" {
		result.Add("tag_name", LoraNetworkSetsListCmdTagName)
	}

	if LoraNetworkSetsListCmdTagValue != "" {
		result.Add("tag_value", LoraNetworkSetsListCmdTagValue)
	}

	if LoraNetworkSetsListCmdTagValueMatchMode != "exact" {
		result.Add("tag_value_match_mode", LoraNetworkSetsListCmdTagValueMatchMode)
	}

	if LoraNetworkSetsListCmdLimit != 0 {
		result.Add("limit", sprintf("%d", LoraNetworkSetsListCmdLimit))
	}

	return result
}
