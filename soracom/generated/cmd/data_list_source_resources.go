// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DataListSourceResourcesCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var DataListSourceResourcesCmdLastEvaluatedKey string

// DataListSourceResourcesCmdResourceType holds value of 'resource_type' option
var DataListSourceResourcesCmdResourceType string

// DataListSourceResourcesCmdLimit holds value of 'limit' option
var DataListSourceResourcesCmdLimit int64

// DataListSourceResourcesCmdPaginate indicates to do pagination or not
var DataListSourceResourcesCmdPaginate bool

// DataListSourceResourcesCmdOutputJSONL indicates to output with jsonl format
var DataListSourceResourcesCmdOutputJSONL bool

func init() {
	DataListSourceResourcesCmd.Flags().StringVar(&DataListSourceResourcesCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of 'resourceId' in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	DataListSourceResourcesCmd.Flags().StringVar(&DataListSourceResourcesCmdResourceType, "resource-type", "", TRAPI("Type of data source resource"))

	DataListSourceResourcesCmd.Flags().Int64Var(&DataListSourceResourcesCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	DataListSourceResourcesCmd.Flags().BoolVar(&DataListSourceResourcesCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	DataListSourceResourcesCmd.Flags().BoolVar(&DataListSourceResourcesCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	DataCmd.AddCommand(DataListSourceResourcesCmd)
}

// DataListSourceResourcesCmd defines 'list-source-resources' subcommand
var DataListSourceResourcesCmd = &cobra.Command{
	Use:   "list-source-resources",
	Short: TRAPI("/data/resources:get:summary"),
	Long:  TRAPI(`/data/resources:get:description`) + "\n\n" + createLinkToAPIReference("DataEntry", "listDataSourceResources"),
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

		param, err := collectDataListSourceResourcesCmdParams(ac)
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
			if DataListSourceResourcesCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectDataListSourceResourcesCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDataListSourceResourcesCmd("/data/resources"),
		query:  buildQueryForDataListSourceResourcesCmd(),

		doPagination:                      DataListSourceResourcesCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDataListSourceResourcesCmd(path string) string {

	return path
}

func buildQueryForDataListSourceResourcesCmd() url.Values {
	result := url.Values{}

	if DataListSourceResourcesCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", DataListSourceResourcesCmdLastEvaluatedKey)
	}

	if DataListSourceResourcesCmdResourceType != "" {
		result.Add("resource_type", DataListSourceResourcesCmdResourceType)
	}

	if DataListSourceResourcesCmdLimit != 0 {
		result.Add("limit", sprintf("%d", DataListSourceResourcesCmdLimit))
	}

	return result
}
