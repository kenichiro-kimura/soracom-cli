// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var VpgListCmdLastEvaluatedKey string

// VpgListCmdTagName holds value of 'tag_name' option
var VpgListCmdTagName string

// VpgListCmdTagValue holds value of 'tag_value' option
var VpgListCmdTagValue string

// VpgListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var VpgListCmdTagValueMatchMode string

// VpgListCmdLimit holds value of 'limit' option
var VpgListCmdLimit int64

// VpgListCmdPaginate indicates to do pagination or not
var VpgListCmdPaginate bool

func init() {
	VpgListCmd.Flags().StringVar(&VpgListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The last group ID retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next VPG onward."))

	VpgListCmd.Flags().StringVar(&VpgListCmdTagName, "tag-name", "", TRAPI("Tag name of the VPG. Filters through all VPGs that exactly match the tag name. When tag_name is specified, tag_value is required."))

	VpgListCmd.Flags().StringVar(&VpgListCmdTagValue, "tag-value", "", TRAPI("Tag value of the VPGs."))

	VpgListCmd.Flags().StringVar(&VpgListCmdTagValueMatchMode, "tag-value-match-mode", "exact", TRAPI("Tag match mode."))

	VpgListCmd.Flags().Int64Var(&VpgListCmdLimit, "limit", 0, TRAPI("Maximum number of results per response page."))

	VpgListCmd.Flags().BoolVar(&VpgListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	VpgCmd.AddCommand(VpgListCmd)
}

// VpgListCmd defines 'list' subcommand
var VpgListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/virtual_private_gateways:get:summary"),
	Long:  TRAPI(`/virtual_private_gateways:get:description`),
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

		param, err := collectVpgListCmdParams(ac)
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

func collectVpgListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForVpgListCmd("/virtual_private_gateways"),
		query:  buildQueryForVpgListCmd(),

		doPagination:                      VpgListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgListCmd(path string) string {

	return path
}

func buildQueryForVpgListCmd() url.Values {
	result := url.Values{}

	if VpgListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", VpgListCmdLastEvaluatedKey)
	}

	if VpgListCmdTagName != "" {
		result.Add("tag_name", VpgListCmdTagName)
	}

	if VpgListCmdTagValue != "" {
		result.Add("tag_value", VpgListCmdTagValue)
	}

	if VpgListCmdTagValueMatchMode != "exact" {
		result.Add("tag_value_match_mode", VpgListCmdTagValueMatchMode)
	}

	if VpgListCmdLimit != 0 {
		result.Add("limit", sprintf("%d", VpgListCmdLimit))
	}

	return result
}
