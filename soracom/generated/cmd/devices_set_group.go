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

// DevicesSetGroupCmdDeviceId holds value of 'device_id' option
var DevicesSetGroupCmdDeviceId string

// DevicesSetGroupCmdBody holds contents of request body to be sent
var DevicesSetGroupCmdBody string

func init() {
	DevicesSetGroupCmd.Flags().StringVar(&DevicesSetGroupCmdDeviceId, "device-id", "", TRAPI("Device to update"))

	DevicesSetGroupCmd.Flags().StringVar(&DevicesSetGroupCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	DevicesCmd.AddCommand(DevicesSetGroupCmd)
}

// DevicesSetGroupCmd defines 'set-group' subcommand
var DevicesSetGroupCmd = &cobra.Command{
	Use:   "set-group",
	Short: TRAPI("/devices/{device_id}/set_group:post:summary"),
	Long:  TRAPI(`/devices/{device_id}/set_group:post:description`),
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

		param, err := collectDevicesSetGroupCmdParams(ac)
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

func collectDevicesSetGroupCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForDevicesSetGroupCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, DevicesSetGroupCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForDevicesSetGroupCmd("/devices/{device_id}/set_group"),
		query:       buildQueryForDevicesSetGroupCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesSetGroupCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesSetGroupCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForDevicesSetGroupCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForDevicesSetGroupCmd() (string, error) {
	var result map[string]interface{}

	if DevicesSetGroupCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(DevicesSetGroupCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesSetGroupCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if DevicesSetGroupCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(DevicesSetGroupCmdBody)
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

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
