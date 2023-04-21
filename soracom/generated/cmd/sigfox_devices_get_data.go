// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SigfoxDevicesGetDataCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesGetDataCmdDeviceId string

// SigfoxDevicesGetDataCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SigfoxDevicesGetDataCmdLastEvaluatedKey string

// SigfoxDevicesGetDataCmdSort holds value of 'sort' option
var SigfoxDevicesGetDataCmdSort string

// SigfoxDevicesGetDataCmdFrom holds value of 'from' option
var SigfoxDevicesGetDataCmdFrom int64

// SigfoxDevicesGetDataCmdLimit holds value of 'limit' option
var SigfoxDevicesGetDataCmdLimit int64

// SigfoxDevicesGetDataCmdTo holds value of 'to' option
var SigfoxDevicesGetDataCmdTo int64

// SigfoxDevicesGetDataCmdPaginate indicates to do pagination or not
var SigfoxDevicesGetDataCmdPaginate bool

// SigfoxDevicesGetDataCmdOutputJSONL indicates to output with jsonl format
var SigfoxDevicesGetDataCmdOutputJSONL bool

func init() {
	SigfoxDevicesGetDataCmd.Flags().StringVar(&SigfoxDevicesGetDataCmdDeviceId, "device-id", "", TRAPI("Device ID of the target subscriber that generated data entries."))

	SigfoxDevicesGetDataCmd.Flags().StringVar(&SigfoxDevicesGetDataCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of 'time' in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	SigfoxDevicesGetDataCmd.Flags().StringVar(&SigfoxDevicesGetDataCmdSort, "sort", "desc", TRAPI("Sort order of the data entries. Either descending (latest data entry first) or ascending (oldest data entry first)."))

	SigfoxDevicesGetDataCmd.Flags().Int64Var(&SigfoxDevicesGetDataCmdFrom, "from", 0, TRAPI("Start time for the data entries search range (unixtime in milliseconds)."))

	SigfoxDevicesGetDataCmd.Flags().Int64Var(&SigfoxDevicesGetDataCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	SigfoxDevicesGetDataCmd.Flags().Int64Var(&SigfoxDevicesGetDataCmdTo, "to", 0, TRAPI("End time for the data entries search range (unixtime in milliseconds)."))

	SigfoxDevicesGetDataCmd.Flags().BoolVar(&SigfoxDevicesGetDataCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	SigfoxDevicesGetDataCmd.Flags().BoolVar(&SigfoxDevicesGetDataCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	SigfoxDevicesCmd.AddCommand(SigfoxDevicesGetDataCmd)
}

// SigfoxDevicesGetDataCmd defines 'get-data' subcommand
var SigfoxDevicesGetDataCmd = &cobra.Command{
	Use:   "get-data",
	Short: TRAPI("/sigfox_devices/{device_id}/data:get:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/data:get:description`) + "\n\n" + createLinkToAPIReference("SigfoxDevice", "getDataFromSigfoxDevice"),
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

		param, err := collectSigfoxDevicesGetDataCmdParams(ac)
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
			if SigfoxDevicesGetDataCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSigfoxDevicesGetDataCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SigfoxDevicesGetDataCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSigfoxDevicesGetDataCmd("/sigfox_devices/{device_id}/data"),
		query:  buildQueryForSigfoxDevicesGetDataCmd(),

		doPagination:                      SigfoxDevicesGetDataCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSigfoxDevicesGetDataCmd(path string) string {

	escapedDeviceId := url.PathEscape(SigfoxDevicesGetDataCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesGetDataCmd() url.Values {
	result := url.Values{}

	if SigfoxDevicesGetDataCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SigfoxDevicesGetDataCmdLastEvaluatedKey)
	}

	if SigfoxDevicesGetDataCmdSort != "desc" {
		result.Add("sort", SigfoxDevicesGetDataCmdSort)
	}

	if SigfoxDevicesGetDataCmdFrom != 0 {
		result.Add("from", sprintf("%d", SigfoxDevicesGetDataCmdFrom))
	}

	if SigfoxDevicesGetDataCmdLimit != 0 {
		result.Add("limit", sprintf("%d", SigfoxDevicesGetDataCmdLimit))
	}

	if SigfoxDevicesGetDataCmdTo != 0 {
		result.Add("to", sprintf("%d", SigfoxDevicesGetDataCmdTo))
	}

	return result
}
