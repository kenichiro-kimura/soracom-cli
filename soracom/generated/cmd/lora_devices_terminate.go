package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesTerminateCmdDeviceId holds value of 'device_id' option
var LoraDevicesTerminateCmdDeviceId string

func init() {
	LoraDevicesTerminateCmd.Flags().StringVar(&LoraDevicesTerminateCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))

	LoraDevicesCmd.AddCommand(LoraDevicesTerminateCmd)
}

// LoraDevicesTerminateCmd defines 'terminate' subcommand
var LoraDevicesTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/lora_devices/{device_id}/terminate:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/terminate:post:description`),
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

		param, err := collectLoraDevicesTerminateCmdParams(ac)
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

func collectLoraDevicesTerminateCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraDevicesTerminateCmd("/lora_devices/{device_id}/terminate"),
		query:  buildQueryForLoraDevicesTerminateCmd(),
	}, nil
}

func buildPathForLoraDevicesTerminateCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesTerminateCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesTerminateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
