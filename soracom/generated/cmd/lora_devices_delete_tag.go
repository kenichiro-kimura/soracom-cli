// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraDevicesDeleteTagCmdDeviceId holds value of 'device_id' option
var LoraDevicesDeleteTagCmdDeviceId string

// LoraDevicesDeleteTagCmdTagName holds value of 'tag_name' option
var LoraDevicesDeleteTagCmdTagName string

func InitLoraDevicesDeleteTagCmd() {
	LoraDevicesDeleteTagCmd.Flags().StringVar(&LoraDevicesDeleteTagCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRaWAN device."))

	LoraDevicesDeleteTagCmd.Flags().StringVar(&LoraDevicesDeleteTagCmdTagName, "tag-name", "", TRAPI("Tag name to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))

	LoraDevicesDeleteTagCmd.RunE = LoraDevicesDeleteTagCmdRunE

	LoraDevicesCmd.AddCommand(LoraDevicesDeleteTagCmd)
}

// LoraDevicesDeleteTagCmd defines 'delete-tag' subcommand
var LoraDevicesDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TRAPI("/lora_devices/{device_id}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/tags/{tag_name}:delete:description`) + "\n\n" + createLinkToAPIReference("LoraDevice", "deleteLoraDeviceTag"),
}

func LoraDevicesDeleteTagCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectLoraDevicesDeleteTagCmdParams(ac)
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

func collectLoraDevicesDeleteTagCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, LoraDevicesDeleteTagCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("tag_name", "tag-name", "path", parsedBody, LoraDevicesDeleteTagCmdTagName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForLoraDevicesDeleteTagCmd("/lora_devices/{device_id}/tags/{tag_name}"),
		query:  buildQueryForLoraDevicesDeleteTagCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraDevicesDeleteTagCmd(path string) string {

	escapedDeviceId := url.PathEscape(LoraDevicesDeleteTagCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	escapedTagName := url.PathEscape(LoraDevicesDeleteTagCmdTagName)

	path = strReplace(path, "{"+"tag_name"+"}", escapedTagName, -1)

	return path
}

func buildQueryForLoraDevicesDeleteTagCmd() url.Values {
	result := url.Values{}

	return result
}
