// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DevicesGetInstanceCmdDeviceId holds value of 'device_id' option
var DevicesGetInstanceCmdDeviceId string

// DevicesGetInstanceCmdInstance holds value of 'instance' option
var DevicesGetInstanceCmdInstance string

// DevicesGetInstanceCmdObject holds value of 'object' option
var DevicesGetInstanceCmdObject string

// DevicesGetInstanceCmdModel holds value of 'model' option
var DevicesGetInstanceCmdModel bool

func init() {
	DevicesGetInstanceCmd.Flags().StringVar(&DevicesGetInstanceCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesGetInstanceCmd.Flags().StringVar(&DevicesGetInstanceCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesGetInstanceCmd.Flags().StringVar(&DevicesGetInstanceCmdObject, "object", "", TRAPI("Object ID"))

	DevicesGetInstanceCmd.Flags().BoolVar(&DevicesGetInstanceCmdModel, "model", false, TRAPI("Whether or not to add model information"))
	DevicesCmd.AddCommand(DevicesGetInstanceCmd)
}

// DevicesGetInstanceCmd defines 'get-instance' subcommand
var DevicesGetInstanceCmd = &cobra.Command{
	Use:   "get-instance",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}:get:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}:get:description`),
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

		param, err := collectDevicesGetInstanceCmdParams(ac)
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

func collectDevicesGetInstanceCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, DevicesGetInstanceCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("instance", "instance", "path", parsedBody, DevicesGetInstanceCmdInstance)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("object", "object", "path", parsedBody, DevicesGetInstanceCmdObject)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesGetInstanceCmd("/devices/{device_id}/{object}/{instance}"),
		query:  buildQueryForDevicesGetInstanceCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesGetInstanceCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesGetInstanceCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	escapedInstance := url.PathEscape(DevicesGetInstanceCmdInstance)

	path = strReplace(path, "{"+"instance"+"}", escapedInstance, -1)

	escapedObject := url.PathEscape(DevicesGetInstanceCmdObject)

	path = strReplace(path, "{"+"object"+"}", escapedObject, -1)

	return path
}

func buildQueryForDevicesGetInstanceCmd() url.Values {
	result := url.Values{}

	if DevicesGetInstanceCmdModel != false {
		result.Add("model", sprintf("%t", DevicesGetInstanceCmdModel))
	}

	return result
}
