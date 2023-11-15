// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesAtomcamSettingsGetMotionTaggingCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomcamSettingsGetMotionTaggingCmdDeviceId string

func InitSoraCamDevicesAtomcamSettingsGetMotionTaggingCmd() {
	SoraCamDevicesAtomcamSettingsGetMotionTaggingCmd.Flags().StringVar(&SoraCamDevicesAtomcamSettingsGetMotionTaggingCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesAtomcamSettingsGetMotionTaggingCmd.RunE = SoraCamDevicesAtomcamSettingsGetMotionTaggingCmdRunE

	SoraCamDevicesAtomcamSettingsCmd.AddCommand(SoraCamDevicesAtomcamSettingsGetMotionTaggingCmd)
}

// SoraCamDevicesAtomcamSettingsGetMotionTaggingCmd defines 'get-motion-tagging' subcommand
var SoraCamDevicesAtomcamSettingsGetMotionTaggingCmd = &cobra.Command{
	Use:   "get-motion-tagging",
	Short: TRAPI("/sora_cam/devices/{device_id}/atomcam/settings/motion_tagging:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atomcam/settings/motion_tagging:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "getSoraCamDeviceAtomCamSettingsMotionTagging"),
}

func SoraCamDevicesAtomcamSettingsGetMotionTaggingCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectSoraCamDevicesAtomcamSettingsGetMotionTaggingCmdParams(ac)
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

func collectSoraCamDevicesAtomcamSettingsGetMotionTaggingCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomcamSettingsGetMotionTaggingCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesAtomcamSettingsGetMotionTaggingCmd("/sora_cam/devices/{device_id}/atomcam/settings/motion_tagging"),
		query:  buildQueryForSoraCamDevicesAtomcamSettingsGetMotionTaggingCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomcamSettingsGetMotionTaggingCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomcamSettingsGetMotionTaggingCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomcamSettingsGetMotionTaggingCmd() url.Values {
	result := url.Values{}

	return result
}
