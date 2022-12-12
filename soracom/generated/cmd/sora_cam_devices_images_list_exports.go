// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesImagesListExportsCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesImagesListExportsCmdDeviceId string

// SoraCamDevicesImagesListExportsCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SoraCamDevicesImagesListExportsCmdLastEvaluatedKey string

// SoraCamDevicesImagesListExportsCmdSort holds value of 'sort' option
var SoraCamDevicesImagesListExportsCmdSort string

// SoraCamDevicesImagesListExportsCmdLimit holds value of 'limit' option
var SoraCamDevicesImagesListExportsCmdLimit int64

// SoraCamDevicesImagesListExportsCmdPaginate indicates to do pagination or not
var SoraCamDevicesImagesListExportsCmdPaginate bool

// SoraCamDevicesImagesListExportsCmdOutputJSONL indicates to output with jsonl format
var SoraCamDevicesImagesListExportsCmdOutputJSONL bool

func init() {
	SoraCamDevicesImagesListExportsCmd.Flags().StringVar(&SoraCamDevicesImagesListExportsCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device. If this ID is not specified, all compatible camera devices owned by the operator will be returned."))

	SoraCamDevicesImagesListExportsCmd.Flags().StringVar(&SoraCamDevicesImagesListExportsCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("Value of the x-soracom-next-key header in the response to the last export request. By specifying this parameter, you can continue to retrieve the list from the last export request."))

	SoraCamDevicesImagesListExportsCmd.Flags().StringVar(&SoraCamDevicesImagesListExportsCmdSort, "sort", "desc", TRAPI("Sort order. The list in the response is sorted in ascending (`asc`) or descending (`desc`) order of `requestedTime`. The default is `desc` i.e. newer items are sorted first."))

	SoraCamDevicesImagesListExportsCmd.Flags().Int64Var(&SoraCamDevicesImagesListExportsCmdLimit, "limit", 10, TRAPI("Maximum number of data related to export processing to be acquired in one request. Note that the number of data that can be acquired may be less than the specified number."))

	SoraCamDevicesImagesListExportsCmd.Flags().BoolVar(&SoraCamDevicesImagesListExportsCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	SoraCamDevicesImagesListExportsCmd.Flags().BoolVar(&SoraCamDevicesImagesListExportsCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	SoraCamDevicesImagesCmd.AddCommand(SoraCamDevicesImagesListExportsCmd)
}

// SoraCamDevicesImagesListExportsCmd defines 'list-exports' subcommand
var SoraCamDevicesImagesListExportsCmd = &cobra.Command{
	Use:   "list-exports",
	Short: TRAPI("/sora_cam/devices/images/exports:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/images/exports:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "listSoraCamDeviceImageExports"),
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

		param, err := collectSoraCamDevicesImagesListExportsCmdParams(ac)
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
			if SoraCamDevicesImagesListExportsCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSoraCamDevicesImagesListExportsCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesImagesListExportsCmd("/sora_cam/devices/images/exports"),
		query:  buildQueryForSoraCamDevicesImagesListExportsCmd(),

		doPagination:                      SoraCamDevicesImagesListExportsCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesImagesListExportsCmd(path string) string {

	return path
}

func buildQueryForSoraCamDevicesImagesListExportsCmd() url.Values {
	result := url.Values{}

	if SoraCamDevicesImagesListExportsCmdDeviceId != "" {
		result.Add("device_id", SoraCamDevicesImagesListExportsCmdDeviceId)
	}

	if SoraCamDevicesImagesListExportsCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SoraCamDevicesImagesListExportsCmdLastEvaluatedKey)
	}

	if SoraCamDevicesImagesListExportsCmdSort != "desc" {
		result.Add("sort", SoraCamDevicesImagesListExportsCmdSort)
	}

	if SoraCamDevicesImagesListExportsCmdLimit != 10 {
		result.Add("limit", sprintf("%d", SoraCamDevicesImagesListExportsCmdLimit))
	}

	return result
}
