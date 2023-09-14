// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsSessionEventsCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SimsSessionEventsCmdLastEvaluatedKey string

// SimsSessionEventsCmdSimId holds value of 'sim_id' option
var SimsSessionEventsCmdSimId string

// SimsSessionEventsCmdFrom holds value of 'from' option
var SimsSessionEventsCmdFrom int64

// SimsSessionEventsCmdLimit holds value of 'limit' option
var SimsSessionEventsCmdLimit int64

// SimsSessionEventsCmdTo holds value of 'to' option
var SimsSessionEventsCmdTo int64

// SimsSessionEventsCmdPaginate indicates to do pagination or not
var SimsSessionEventsCmdPaginate bool

// SimsSessionEventsCmdOutputJSONL indicates to output with jsonl format
var SimsSessionEventsCmdOutputJSONL bool

func InitSimsSessionEventsCmd() {
	SimsSessionEventsCmd.Flags().StringVar(&SimsSessionEventsCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The time stamp of the last event retrieved on the previous page. By specifying this parameter, you can continue to retrieve the list from the next event onward."))

	SimsSessionEventsCmd.Flags().StringVar(&SimsSessionEventsCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsSessionEventsCmd.Flags().Int64Var(&SimsSessionEventsCmdFrom, "from", 0, TRAPI("Start time for the events search range (UNIX time in milliseconds)."))

	SimsSessionEventsCmd.Flags().Int64Var(&SimsSessionEventsCmdLimit, "limit", 0, TRAPI("Maximum number of events to retrieve."))

	SimsSessionEventsCmd.Flags().Int64Var(&SimsSessionEventsCmdTo, "to", 0, TRAPI("End time for the events search range (UNIX time in milliseconds)."))

	SimsSessionEventsCmd.Flags().BoolVar(&SimsSessionEventsCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	SimsSessionEventsCmd.Flags().BoolVar(&SimsSessionEventsCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	SimsSessionEventsCmd.RunE = SimsSessionEventsCmdRunE

	SimsCmd.AddCommand(SimsSessionEventsCmd)
}

// SimsSessionEventsCmd defines 'session-events' subcommand
var SimsSessionEventsCmd = &cobra.Command{
	Use:   "session-events",
	Short: TRAPI("/sims/{sim_id}/events/sessions:get:summary"),
	Long:  TRAPI(`/sims/{sim_id}/events/sessions:get:description`) + "\n\n" + createLinkToAPIReference("Sim", "listSimSessionEvents"),
}

func SimsSessionEventsCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSimsSessionEventsCmdParams(ac)
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
		if SimsSessionEventsCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectSimsSessionEventsCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsSessionEventsCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSimsSessionEventsCmd("/sims/{sim_id}/events/sessions"),
		query:  buildQueryForSimsSessionEventsCmd(),

		doPagination:                      SimsSessionEventsCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsSessionEventsCmd(path string) string {

	escapedSimId := url.PathEscape(SimsSessionEventsCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsSessionEventsCmd() url.Values {
	result := url.Values{}

	if SimsSessionEventsCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SimsSessionEventsCmdLastEvaluatedKey)
	}

	if SimsSessionEventsCmdFrom != 0 {
		result.Add("from", sprintf("%d", SimsSessionEventsCmdFrom))
	}

	if SimsSessionEventsCmdLimit != 0 {
		result.Add("limit", sprintf("%d", SimsSessionEventsCmdLimit))
	}

	if SimsSessionEventsCmdTo != 0 {
		result.Add("to", sprintf("%d", SimsSessionEventsCmdTo))
	}

	return result
}
