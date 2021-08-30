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

// DevicesPutResourceCmdDeviceId holds value of 'device_id' option
var DevicesPutResourceCmdDeviceId string

// DevicesPutResourceCmdInstance holds value of 'instance' option
var DevicesPutResourceCmdInstance string

// DevicesPutResourceCmdObject holds value of 'object' option
var DevicesPutResourceCmdObject string

// DevicesPutResourceCmdResource holds value of 'resource' option
var DevicesPutResourceCmdResource string

// DevicesPutResourceCmdBody holds contents of request body to be sent
var DevicesPutResourceCmdBody string

func init() {
	DevicesPutResourceCmd.Flags().StringVar(&DevicesPutResourceCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesPutResourceCmd.Flags().StringVar(&DevicesPutResourceCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesPutResourceCmd.Flags().StringVar(&DevicesPutResourceCmdObject, "object", "", TRAPI("Object ID"))

	DevicesPutResourceCmd.Flags().StringVar(&DevicesPutResourceCmdResource, "resource", "", TRAPI("Resource ID"))

	DevicesPutResourceCmd.Flags().StringVar(&DevicesPutResourceCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	DevicesCmd.AddCommand(DevicesPutResourceCmd)
}

// DevicesPutResourceCmd defines 'put-resource' subcommand
var DevicesPutResourceCmd = &cobra.Command{
	Use:   "put-resource",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}/{resource}:put:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}/{resource}:put:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
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

		param, err := collectDevicesPutResourceCmdParams(ac)
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

func collectDevicesPutResourceCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForDevicesPutResourceCmd()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &parsedBody)
	if err != nil {
		return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
	}
	contentType := "application/json"

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, DevicesPutResourceCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("instance", "instance", "path", parsedBody, DevicesPutResourceCmdInstance)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("object", "object", "path", parsedBody, DevicesPutResourceCmdObject)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("resource", "resource", "path", parsedBody, DevicesPutResourceCmdResource)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForDevicesPutResourceCmd("/devices/{device_id}/{object}/{instance}/{resource}"),
		query:       buildQueryForDevicesPutResourceCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesPutResourceCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesPutResourceCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	escapedInstance := url.PathEscape(DevicesPutResourceCmdInstance)

	path = strReplace(path, "{"+"instance"+"}", escapedInstance, -1)

	escapedObject := url.PathEscape(DevicesPutResourceCmdObject)

	path = strReplace(path, "{"+"object"+"}", escapedObject, -1)

	escapedResource := url.PathEscape(DevicesPutResourceCmdResource)

	path = strReplace(path, "{"+"resource"+"}", escapedResource, -1)

	return path
}

func buildQueryForDevicesPutResourceCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForDevicesPutResourceCmd() (string, error) {
	var result map[string]interface{}

	if DevicesPutResourceCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(DevicesPutResourceCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesPutResourceCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if DevicesPutResourceCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(DevicesPutResourceCmdBody)
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
