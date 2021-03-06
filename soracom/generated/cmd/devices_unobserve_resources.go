package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesUnobserveResourcesCmdDeviceId holds value of 'device_id' option
var DevicesUnobserveResourcesCmdDeviceId string

// DevicesUnobserveResourcesCmdInstance holds value of 'instance' option
var DevicesUnobserveResourcesCmdInstance string

// DevicesUnobserveResourcesCmdObject holds value of 'object' option
var DevicesUnobserveResourcesCmdObject string

func init() {
	DevicesUnobserveResourcesCmd.Flags().StringVar(&DevicesUnobserveResourcesCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesUnobserveResourcesCmd.Flags().StringVar(&DevicesUnobserveResourcesCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesUnobserveResourcesCmd.Flags().StringVar(&DevicesUnobserveResourcesCmdObject, "object", "", TRAPI("Object ID"))

	DevicesCmd.AddCommand(DevicesUnobserveResourcesCmd)
}

// DevicesUnobserveResourcesCmd defines 'unobserve-resources' subcommand
var DevicesUnobserveResourcesCmd = &cobra.Command{
	Use:   "unobserve-resources",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}/unobserve:post:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}/unobserve:post:description`),
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

		param, err := collectDevicesUnobserveResourcesCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectDevicesUnobserveResourcesCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForDevicesUnobserveResourcesCmd("/devices/{device_id}/{object}/{instance}/unobserve"),
		query:  buildQueryForDevicesUnobserveResourcesCmd(),
	}, nil
}

func buildPathForDevicesUnobserveResourcesCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", DevicesUnobserveResourcesCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"instance"+"}", DevicesUnobserveResourcesCmdInstance, -1)

	path = strings.Replace(path, "{"+"object"+"}", DevicesUnobserveResourcesCmdObject, -1)

	return path
}

func buildQueryForDevicesUnobserveResourcesCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
