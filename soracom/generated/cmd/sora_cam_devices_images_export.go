// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// SoraCamDevicesImagesExportCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesImagesExportCmdDeviceId string

// SoraCamDevicesImagesExportCmdImageFilters holds multiple values of 'imageFilters' option
var SoraCamDevicesImagesExportCmdImageFilters []string

// SoraCamDevicesImagesExportCmdTime holds value of 'time' option
var SoraCamDevicesImagesExportCmdTime int64

// SoraCamDevicesImagesExportCmdBody holds contents of request body to be sent
var SoraCamDevicesImagesExportCmdBody string

func init() {
	SoraCamDevicesImagesExportCmd.Flags().StringVar(&SoraCamDevicesImagesExportCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesImagesExportCmd.Flags().StringSliceVar(&SoraCamDevicesImagesExportCmdImageFilters, "image-filters", []string{}, TRAPI("Image filters to be applied to the exported image."))

	SoraCamDevicesImagesExportCmd.Flags().Int64Var(&SoraCamDevicesImagesExportCmdTime, "time", 0, TRAPI("Time to be exported (unix time in milliseconds). A still image is exported from the recorded video taken at the specified time.'"))

	SoraCamDevicesImagesExportCmd.Flags().StringVar(&SoraCamDevicesImagesExportCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SoraCamDevicesImagesCmd.AddCommand(SoraCamDevicesImagesExportCmd)
}

// SoraCamDevicesImagesExportCmd defines 'export' subcommand
var SoraCamDevicesImagesExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/sora_cam/devices/{device_id}/images/exports:post:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/images/exports:post:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "exportSoraCamDeviceRecordedImage"),
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

		param, err := collectSoraCamDevicesImagesExportCmdParams(ac)
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

func collectSoraCamDevicesImagesExportCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraCamDevicesImagesExportCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesImagesExportCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("time", "time", "body", parsedBody, SoraCamDevicesImagesExportCmdTime)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSoraCamDevicesImagesExportCmd("/sora_cam/devices/{device_id}/images/exports"),
		query:       buildQueryForSoraCamDevicesImagesExportCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesImagesExportCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesImagesExportCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesImagesExportCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraCamDevicesImagesExportCmd() (string, error) {
	var result map[string]interface{}

	if SoraCamDevicesImagesExportCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraCamDevicesImagesExportCmdBody, "@") {
			fname := strings.TrimPrefix(SoraCamDevicesImagesExportCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SoraCamDevicesImagesExportCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraCamDevicesImagesExportCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if len(SoraCamDevicesImagesExportCmdImageFilters) != 0 {
		result["imageFilters"] = SoraCamDevicesImagesExportCmdImageFilters
	}

	if SoraCamDevicesImagesExportCmdTime != 0 {
		result["time"] = SoraCamDevicesImagesExportCmdTime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
