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

// SigfoxDevicesSetGroupCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesSetGroupCmdDeviceId string

// SigfoxDevicesSetGroupCmdGroupId holds value of 'groupId' option
var SigfoxDevicesSetGroupCmdGroupId string

// SigfoxDevicesSetGroupCmdOperatorId holds value of 'operatorId' option
var SigfoxDevicesSetGroupCmdOperatorId string

// SigfoxDevicesSetGroupCmdCreatedTime holds value of 'createdTime' option
var SigfoxDevicesSetGroupCmdCreatedTime int64

// SigfoxDevicesSetGroupCmdLastModifiedTime holds value of 'lastModifiedTime' option
var SigfoxDevicesSetGroupCmdLastModifiedTime int64

// SigfoxDevicesSetGroupCmdBody holds contents of request body to be sent
var SigfoxDevicesSetGroupCmdBody string

func init() {
	SigfoxDevicesSetGroupCmd.Flags().StringVar(&SigfoxDevicesSetGroupCmdDeviceId, "device-id", "", TRAPI("Device ID of the target Sigfox device."))

	SigfoxDevicesSetGroupCmd.Flags().StringVar(&SigfoxDevicesSetGroupCmdGroupId, "group-id", "", TRAPI(""))

	SigfoxDevicesSetGroupCmd.Flags().StringVar(&SigfoxDevicesSetGroupCmdOperatorId, "operator-id", "", TRAPI(""))

	SigfoxDevicesSetGroupCmd.Flags().Int64Var(&SigfoxDevicesSetGroupCmdCreatedTime, "created-time", 0, TRAPI(""))

	SigfoxDevicesSetGroupCmd.Flags().Int64Var(&SigfoxDevicesSetGroupCmdLastModifiedTime, "last-modified-time", 0, TRAPI(""))

	SigfoxDevicesSetGroupCmd.Flags().StringVar(&SigfoxDevicesSetGroupCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SigfoxDevicesCmd.AddCommand(SigfoxDevicesSetGroupCmd)
}

// SigfoxDevicesSetGroupCmd defines 'set-group' subcommand
var SigfoxDevicesSetGroupCmd = &cobra.Command{
	Use:   "set-group",
	Short: TRAPI("/sigfox_devices/{device_id}/set_group:post:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/set_group:post:description`),
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

		param, err := collectSigfoxDevicesSetGroupCmdParams(ac)
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

func collectSigfoxDevicesSetGroupCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSigfoxDevicesSetGroupCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SigfoxDevicesSetGroupCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSigfoxDevicesSetGroupCmd("/sigfox_devices/{device_id}/set_group"),
		query:       buildQueryForSigfoxDevicesSetGroupCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSigfoxDevicesSetGroupCmd(path string) string {

	escapedDeviceId := url.PathEscape(SigfoxDevicesSetGroupCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesSetGroupCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSigfoxDevicesSetGroupCmd() (string, error) {
	var result map[string]interface{}

	if SigfoxDevicesSetGroupCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SigfoxDevicesSetGroupCmdBody, "@") {
			fname := strings.TrimPrefix(SigfoxDevicesSetGroupCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SigfoxDevicesSetGroupCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SigfoxDevicesSetGroupCmdBody)
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

	if SigfoxDevicesSetGroupCmdGroupId != "" {
		result["groupId"] = SigfoxDevicesSetGroupCmdGroupId
	}

	if SigfoxDevicesSetGroupCmdOperatorId != "" {
		result["operatorId"] = SigfoxDevicesSetGroupCmdOperatorId
	}

	if SigfoxDevicesSetGroupCmdCreatedTime != 0 {
		result["createdTime"] = SigfoxDevicesSetGroupCmdCreatedTime
	}

	if SigfoxDevicesSetGroupCmdLastModifiedTime != 0 {
		result["lastModifiedTime"] = SigfoxDevicesSetGroupCmdLastModifiedTime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
