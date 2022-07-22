// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// QuerySimsCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var QuerySimsCmdLastEvaluatedKey string

// QuerySimsCmdSearchType holds value of 'search_type' option
var QuerySimsCmdSearchType string

// QuerySimsCmdSessionStatus holds value of 'session_status' option
var QuerySimsCmdSessionStatus string

// QuerySimsCmdBundles holds multiple values of 'bundles' option
var QuerySimsCmdBundles []string

// QuerySimsCmdGroup holds multiple values of 'group' option
var QuerySimsCmdGroup []string

// QuerySimsCmdIccid holds multiple values of 'iccid' option
var QuerySimsCmdIccid []string

// QuerySimsCmdImsi holds multiple values of 'imsi' option
var QuerySimsCmdImsi []string

// QuerySimsCmdMsisdn holds multiple values of 'msisdn' option
var QuerySimsCmdMsisdn []string

// QuerySimsCmdName holds multiple values of 'name' option
var QuerySimsCmdName []string

// QuerySimsCmdSerialNumber holds multiple values of 'serial_number' option
var QuerySimsCmdSerialNumber []string

// QuerySimsCmdSimId holds multiple values of 'sim_id' option
var QuerySimsCmdSimId []string

// QuerySimsCmdTag holds multiple values of 'tag' option
var QuerySimsCmdTag []string

// QuerySimsCmdLimit holds value of 'limit' option
var QuerySimsCmdLimit int64

// QuerySimsCmdPaginate indicates to do pagination or not
var QuerySimsCmdPaginate bool

// QuerySimsCmdOutputJSONL indicates to output with jsonl format
var QuerySimsCmdOutputJSONL bool

func init() {
	QuerySimsCmd.Flags().StringVar(&QuerySimsCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The SIM ID of the last SIM retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next SIM onward."))

	QuerySimsCmd.Flags().StringVar(&QuerySimsCmdSearchType, "search-type", "and", TRAPI("Type of the search ('AND searching' or 'OR searching')"))

	QuerySimsCmd.Flags().StringVar(&QuerySimsCmdSessionStatus, "session-status", "NA", TRAPI("Status of the session to search (ONLINE or OFFLINE)"))

	QuerySimsCmd.Flags().StringSliceVar(&QuerySimsCmdBundles, "bundles", []string{}, TRAPI("Bundles type to search"))

	QuerySimsCmd.Flags().StringSliceVar(&QuerySimsCmdGroup, "group", []string{}, TRAPI("Group name to search"))

	QuerySimsCmd.Flags().StringSliceVar(&QuerySimsCmdIccid, "iccid", []string{}, TRAPI("ICCID to search"))

	QuerySimsCmd.Flags().StringSliceVar(&QuerySimsCmdImsi, "imsi", []string{}, TRAPI("IMSI to search"))

	QuerySimsCmd.Flags().StringSliceVar(&QuerySimsCmdMsisdn, "msisdn", []string{}, TRAPI("MSISDN to search"))

	QuerySimsCmd.Flags().StringSliceVar(&QuerySimsCmdName, "name", []string{}, TRAPI("Name to search"))

	QuerySimsCmd.Flags().StringSliceVar(&QuerySimsCmdSerialNumber, "serial-number", []string{}, TRAPI("Serial number to search"))

	QuerySimsCmd.Flags().StringSliceVar(&QuerySimsCmdSimId, "sim-id", []string{}, TRAPI("Identifier of the SIM to search"))

	QuerySimsCmd.Flags().StringSliceVar(&QuerySimsCmdTag, "tag", []string{}, TRAPI("String of tag values to search"))

	QuerySimsCmd.Flags().Int64Var(&QuerySimsCmdLimit, "limit", 10, TRAPI("The maximum number of item to retrieve"))

	QuerySimsCmd.Flags().BoolVar(&QuerySimsCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	QuerySimsCmd.Flags().BoolVar(&QuerySimsCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	QueryCmd.AddCommand(QuerySimsCmd)
}

// QuerySimsCmd defines 'sims' subcommand
var QuerySimsCmd = &cobra.Command{
	Use:   "sims",
	Short: TRAPI("/query/sims:get:summary"),
	Long:  TRAPI(`/query/sims:get:description`) + "\n\n" + createLinkToAPIReference("Query", "searchSims"),
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

		param, err := collectQuerySimsCmdParams(ac)
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
			if QuerySimsCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectQuerySimsCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForQuerySimsCmd("/query/sims"),
		query:  buildQueryForQuerySimsCmd(),

		doPagination:                      QuerySimsCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForQuerySimsCmd(path string) string {

	return path
}

func buildQueryForQuerySimsCmd() url.Values {
	result := url.Values{}

	if QuerySimsCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", QuerySimsCmdLastEvaluatedKey)
	}

	if QuerySimsCmdSearchType != "and" {
		result.Add("search_type", QuerySimsCmdSearchType)
	}

	if QuerySimsCmdSessionStatus != "NA" {
		result.Add("session_status", QuerySimsCmdSessionStatus)
	}

	for _, s := range QuerySimsCmdBundles {
		if s != "" {
			result.Add("bundles", s)
		}
	}

	for _, s := range QuerySimsCmdGroup {
		if s != "" {
			result.Add("group", s)
		}
	}

	for _, s := range QuerySimsCmdIccid {
		if s != "" {
			result.Add("iccid", s)
		}
	}

	for _, s := range QuerySimsCmdImsi {
		if s != "" {
			result.Add("imsi", s)
		}
	}

	for _, s := range QuerySimsCmdMsisdn {
		if s != "" {
			result.Add("msisdn", s)
		}
	}

	for _, s := range QuerySimsCmdName {
		if s != "" {
			result.Add("name", s)
		}
	}

	for _, s := range QuerySimsCmdSerialNumber {
		if s != "" {
			result.Add("serial_number", s)
		}
	}

	for _, s := range QuerySimsCmdSimId {
		if s != "" {
			result.Add("sim_id", s)
		}
	}

	for _, s := range QuerySimsCmdTag {
		if s != "" {
			result.Add("tag", s)
		}
	}

	if QuerySimsCmdLimit != 10 {
		result.Add("limit", sprintf("%d", QuerySimsCmdLimit))
	}

	return result
}
