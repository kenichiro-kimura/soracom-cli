// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SigfoxDevicesListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SigfoxDevicesListCmdLastEvaluatedKey string

// SigfoxDevicesListCmdTagName holds value of 'tag_name' option
var SigfoxDevicesListCmdTagName string

// SigfoxDevicesListCmdTagValue holds value of 'tag_value' option
var SigfoxDevicesListCmdTagValue string

// SigfoxDevicesListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var SigfoxDevicesListCmdTagValueMatchMode string

// SigfoxDevicesListCmdLimit holds value of 'limit' option
var SigfoxDevicesListCmdLimit int64

// SigfoxDevicesListCmdPaginate indicates to do pagination or not
var SigfoxDevicesListCmdPaginate bool

func init() {
	SigfoxDevicesListCmd.Flags().StringVar(&SigfoxDevicesListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The device ID of the last device retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next device onward."))

	SigfoxDevicesListCmd.Flags().StringVar(&SigfoxDevicesListCmdTagName, "tag-name", "", TRAPI("Tag name for filtering the search (exact match)."))

	SigfoxDevicesListCmd.Flags().StringVar(&SigfoxDevicesListCmdTagValue, "tag-value", "", TRAPI("Tag search string for filtering the search. Required when `tag_name` has been specified."))

	SigfoxDevicesListCmd.Flags().StringVar(&SigfoxDevicesListCmdTagValueMatchMode, "tag-value-match-mode", "exact", TRAPI("Tag match mode."))

	SigfoxDevicesListCmd.Flags().Int64Var(&SigfoxDevicesListCmdLimit, "limit", 0, TRAPI("Maximum number of Sigfox devices to retrieve."))

	SigfoxDevicesListCmd.Flags().BoolVar(&SigfoxDevicesListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	SigfoxDevicesCmd.AddCommand(SigfoxDevicesListCmd)
}

// SigfoxDevicesListCmd defines 'list' subcommand
var SigfoxDevicesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/sigfox_devices:get:summary"),
	Long:  TRAPI(`/sigfox_devices:get:description`),
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

		param, err := collectSigfoxDevicesListCmdParams(ac)
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

func collectSigfoxDevicesListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSigfoxDevicesListCmd("/sigfox_devices"),
		query:  buildQueryForSigfoxDevicesListCmd(),

		doPagination:                      SigfoxDevicesListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSigfoxDevicesListCmd(path string) string {

	return path
}

func buildQueryForSigfoxDevicesListCmd() url.Values {
	result := url.Values{}

	if SigfoxDevicesListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SigfoxDevicesListCmdLastEvaluatedKey)
	}

	if SigfoxDevicesListCmdTagName != "" {
		result.Add("tag_name", SigfoxDevicesListCmdTagName)
	}

	if SigfoxDevicesListCmdTagValue != "" {
		result.Add("tag_value", SigfoxDevicesListCmdTagValue)
	}

	if SigfoxDevicesListCmdTagValueMatchMode != "exact" {
		result.Add("tag_value_match_mode", SigfoxDevicesListCmdTagValueMatchMode)
	}

	if SigfoxDevicesListCmdLimit != 0 {
		result.Add("limit", sprintf("%d", SigfoxDevicesListCmdLimit))
	}

	return result
}
