// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DevicesGetResourceCmdDeviceId holds value of 'device_id' option
var DevicesGetResourceCmdDeviceId string

// DevicesGetResourceCmdInstance holds value of 'instance' option
var DevicesGetResourceCmdInstance string

// DevicesGetResourceCmdObject holds value of 'object' option
var DevicesGetResourceCmdObject string

// DevicesGetResourceCmdResource holds value of 'resource' option
var DevicesGetResourceCmdResource string

// DevicesGetResourceCmdModel holds value of 'model' option
var DevicesGetResourceCmdModel bool

func InitDevicesGetResourceCmd() {
	DevicesGetResourceCmd.Flags().StringVar(&DevicesGetResourceCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesGetResourceCmd.Flags().StringVar(&DevicesGetResourceCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesGetResourceCmd.Flags().StringVar(&DevicesGetResourceCmdObject, "object", "", TRAPI("Object ID"))

	DevicesGetResourceCmd.Flags().StringVar(&DevicesGetResourceCmdResource, "resource", "", TRAPI("Resource ID"))

	DevicesGetResourceCmd.Flags().BoolVar(&DevicesGetResourceCmdModel, "model", false, TRAPI("Whether or not to add model information"))

	DevicesGetResourceCmd.RunE = DevicesGetResourceCmdRunE

	DevicesCmd.AddCommand(DevicesGetResourceCmd)
}

// DevicesGetResourceCmd defines 'get-resource' subcommand
var DevicesGetResourceCmd = &cobra.Command{
	Use:   "get-resource",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}/{resource}:get:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}/{resource}:get:description`) + "\n\n" + createLinkToAPIReference("Device", "readDeviceResource"),
}

func DevicesGetResourceCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectDevicesGetResourceCmdParams(ac)
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
}

func collectDevicesGetResourceCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, DevicesGetResourceCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("instance", "instance", "path", parsedBody, DevicesGetResourceCmdInstance)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("object", "object", "path", parsedBody, DevicesGetResourceCmdObject)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("resource", "resource", "path", parsedBody, DevicesGetResourceCmdResource)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesGetResourceCmd("/devices/{device_id}/{object}/{instance}/{resource}"),
		query:  buildQueryForDevicesGetResourceCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesGetResourceCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesGetResourceCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	escapedInstance := url.PathEscape(DevicesGetResourceCmdInstance)

	path = strReplace(path, "{"+"instance"+"}", escapedInstance, -1)

	escapedObject := url.PathEscape(DevicesGetResourceCmdObject)

	path = strReplace(path, "{"+"object"+"}", escapedObject, -1)

	escapedResource := url.PathEscape(DevicesGetResourceCmdResource)

	path = strReplace(path, "{"+"resource"+"}", escapedResource, -1)

	return path
}

func buildQueryForDevicesGetResourceCmd() url.Values {
	result := url.Values{}

	if DevicesGetResourceCmdModel != false {
		result.Add("model", sprintf("%t", DevicesGetResourceCmdModel))
	}

	return result
}
