package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesUnsetGroupCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesUnsetGroupCmdDeviceId string

func init() {
	SigfoxDevicesUnsetGroupCmd.Flags().StringVar(&SigfoxDevicesUnsetGroupCmdDeviceId, "device-id", "", TRAPI("Device ID of the target Sigfox device."))

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesUnsetGroupCmd)
}

// SigfoxDevicesUnsetGroupCmd defines 'unset-group' subcommand
var SigfoxDevicesUnsetGroupCmd = &cobra.Command{
	Use:   "unset-group",
	Short: TRAPI("/sigfox_devices/{device_id}/unset_group:post:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/unset_group:post:description`),
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

		param, err := collectSigfoxDevicesUnsetGroupCmdParams(ac)
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

func collectSigfoxDevicesUnsetGroupCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSigfoxDevicesUnsetGroupCmd("/sigfox_devices/{device_id}/unset_group"),
		query:  buildQueryForSigfoxDevicesUnsetGroupCmd(),
	}, nil
}

func buildPathForSigfoxDevicesUnsetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesUnsetGroupCmdDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesUnsetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
