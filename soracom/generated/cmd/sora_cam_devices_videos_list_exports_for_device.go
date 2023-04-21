// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesVideosListExportsForDeviceCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesVideosListExportsForDeviceCmdDeviceId string

// SoraCamDevicesVideosListExportsForDeviceCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SoraCamDevicesVideosListExportsForDeviceCmdLastEvaluatedKey string

// SoraCamDevicesVideosListExportsForDeviceCmdSort holds value of 'sort' option
var SoraCamDevicesVideosListExportsForDeviceCmdSort string

// SoraCamDevicesVideosListExportsForDeviceCmdLimit holds value of 'limit' option
var SoraCamDevicesVideosListExportsForDeviceCmdLimit int64

// SoraCamDevicesVideosListExportsForDeviceCmdPaginate indicates to do pagination or not
var SoraCamDevicesVideosListExportsForDeviceCmdPaginate bool

// SoraCamDevicesVideosListExportsForDeviceCmdOutputJSONL indicates to output with jsonl format
var SoraCamDevicesVideosListExportsForDeviceCmdOutputJSONL bool

func init() {
	SoraCamDevicesVideosListExportsForDeviceCmd.Flags().StringVar(&SoraCamDevicesVideosListExportsForDeviceCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesVideosListExportsForDeviceCmd.Flags().StringVar(&SoraCamDevicesVideosListExportsForDeviceCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("Value of the x-soracom-next-key header in the response to the last listSoraCamDeviceVideoExportsForDevice request. By specifying this parameter, you can continue to retrieve the list from the last request."))

	SoraCamDevicesVideosListExportsForDeviceCmd.Flags().StringVar(&SoraCamDevicesVideosListExportsForDeviceCmdSort, "sort", "desc", TRAPI("Sort order. The list in the response is sorted in ascending ('asc') or descending ('desc') order of 'requestedTime'. The default is 'desc' i.e. newer items are sorted first."))

	SoraCamDevicesVideosListExportsForDeviceCmd.Flags().Int64Var(&SoraCamDevicesVideosListExportsForDeviceCmdLimit, "limit", 10, TRAPI("Maximum number of items to retrieve in one request. Note that the response may contain fewer items than the specified limit."))

	SoraCamDevicesVideosListExportsForDeviceCmd.Flags().BoolVar(&SoraCamDevicesVideosListExportsForDeviceCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	SoraCamDevicesVideosListExportsForDeviceCmd.Flags().BoolVar(&SoraCamDevicesVideosListExportsForDeviceCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	SoraCamDevicesVideosCmd.AddCommand(SoraCamDevicesVideosListExportsForDeviceCmd)
}

// SoraCamDevicesVideosListExportsForDeviceCmd defines 'list-exports-for-device' subcommand
var SoraCamDevicesVideosListExportsForDeviceCmd = &cobra.Command{
	Use:   "list-exports-for-device",
	Short: TRAPI("/sora_cam/devices/{device_id}/videos/exports:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/videos/exports:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "listSoraCamDeviceVideoExportsForDevice"),
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

		param, err := collectSoraCamDevicesVideosListExportsForDeviceCmdParams(ac)
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
			if SoraCamDevicesVideosListExportsForDeviceCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSoraCamDevicesVideosListExportsForDeviceCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesVideosListExportsForDeviceCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesVideosListExportsForDeviceCmd("/sora_cam/devices/{device_id}/videos/exports"),
		query:  buildQueryForSoraCamDevicesVideosListExportsForDeviceCmd(),

		doPagination:                      SoraCamDevicesVideosListExportsForDeviceCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesVideosListExportsForDeviceCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesVideosListExportsForDeviceCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesVideosListExportsForDeviceCmd() url.Values {
	result := url.Values{}

	if SoraCamDevicesVideosListExportsForDeviceCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SoraCamDevicesVideosListExportsForDeviceCmdLastEvaluatedKey)
	}

	if SoraCamDevicesVideosListExportsForDeviceCmdSort != "desc" {
		result.Add("sort", SoraCamDevicesVideosListExportsForDeviceCmdSort)
	}

	if SoraCamDevicesVideosListExportsForDeviceCmdLimit != 10 {
		result.Add("limit", sprintf("%d", SoraCamDevicesVideosListExportsForDeviceCmdLimit))
	}

	return result
}
