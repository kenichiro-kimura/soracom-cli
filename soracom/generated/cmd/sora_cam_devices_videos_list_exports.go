// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesVideosListExportsCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesVideosListExportsCmdDeviceId string

// SoraCamDevicesVideosListExportsCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SoraCamDevicesVideosListExportsCmdLastEvaluatedKey string

// SoraCamDevicesVideosListExportsCmdSort holds value of 'sort' option
var SoraCamDevicesVideosListExportsCmdSort string

// SoraCamDevicesVideosListExportsCmdLimit holds value of 'limit' option
var SoraCamDevicesVideosListExportsCmdLimit int64

// SoraCamDevicesVideosListExportsCmdPaginate indicates to do pagination or not
var SoraCamDevicesVideosListExportsCmdPaginate bool

// SoraCamDevicesVideosListExportsCmdOutputJSONL indicates to output with jsonl format
var SoraCamDevicesVideosListExportsCmdOutputJSONL bool

func init() {
	SoraCamDevicesVideosListExportsCmd.Flags().StringVar(&SoraCamDevicesVideosListExportsCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device. If this ID is not specified, all compatible camera devices owned by the operator will be returned."))

	SoraCamDevicesVideosListExportsCmd.Flags().StringVar(&SoraCamDevicesVideosListExportsCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("Value of the x-soracom-next-key header in the response to the last export request. By specifying this parameter, you can continue to retrieve the list from the last export request."))

	SoraCamDevicesVideosListExportsCmd.Flags().StringVar(&SoraCamDevicesVideosListExportsCmdSort, "sort", "desc", TRAPI("Sort order. The list in the response is sorted in ascending (`asc`) or descending (`desc`) order of `requestedTime`. The default is `desc` i.e. newer items are sorted first."))

	SoraCamDevicesVideosListExportsCmd.Flags().Int64Var(&SoraCamDevicesVideosListExportsCmdLimit, "limit", 10, TRAPI("Maximum number of data related to export processing to be acquired in one request. Note that the number of data that can be acquired may be less than the specified number."))

	SoraCamDevicesVideosListExportsCmd.Flags().BoolVar(&SoraCamDevicesVideosListExportsCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	SoraCamDevicesVideosListExportsCmd.Flags().BoolVar(&SoraCamDevicesVideosListExportsCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	SoraCamDevicesVideosCmd.AddCommand(SoraCamDevicesVideosListExportsCmd)
}

// SoraCamDevicesVideosListExportsCmd defines 'list-exports' subcommand
var SoraCamDevicesVideosListExportsCmd = &cobra.Command{
	Use:   "list-exports",
	Short: TRAPI("/sora_cam/devices/videos/exports:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/videos/exports:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "listSoraCamDeviceVideoExports"),
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

		param, err := collectSoraCamDevicesVideosListExportsCmdParams(ac)
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
			if SoraCamDevicesVideosListExportsCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSoraCamDevicesVideosListExportsCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesVideosListExportsCmd("/sora_cam/devices/videos/exports"),
		query:  buildQueryForSoraCamDevicesVideosListExportsCmd(),

		doPagination:                      SoraCamDevicesVideosListExportsCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesVideosListExportsCmd(path string) string {

	return path
}

func buildQueryForSoraCamDevicesVideosListExportsCmd() url.Values {
	result := url.Values{}

	if SoraCamDevicesVideosListExportsCmdDeviceId != "" {
		result.Add("device_id", SoraCamDevicesVideosListExportsCmdDeviceId)
	}

	if SoraCamDevicesVideosListExportsCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SoraCamDevicesVideosListExportsCmdLastEvaluatedKey)
	}

	if SoraCamDevicesVideosListExportsCmdSort != "desc" {
		result.Add("sort", SoraCamDevicesVideosListExportsCmdSort)
	}

	if SoraCamDevicesVideosListExportsCmdLimit != 10 {
		result.Add("limit", sprintf("%d", SoraCamDevicesVideosListExportsCmdLimit))
	}

	return result
}
