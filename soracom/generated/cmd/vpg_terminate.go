// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgTerminateCmdVpgId holds value of 'vpg_id' option
var VpgTerminateCmdVpgId string

func InitVpgTerminateCmd() {
	VpgTerminateCmd.Flags().StringVar(&VpgTerminateCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgTerminateCmd.RunE = VpgTerminateCmdRunE

	VpgCmd.AddCommand(VpgTerminateCmd)
}

// VpgTerminateCmd defines 'terminate' subcommand
var VpgTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/terminate:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/terminate:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "terminateVirtualPrivateGateway"),
}

func VpgTerminateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgTerminateCmdParams(ac)
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

func collectVpgTerminateCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgTerminateCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgTerminateCmd("/virtual_private_gateways/{vpg_id}/terminate"),
		query:  buildQueryForVpgTerminateCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgTerminateCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgTerminateCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgTerminateCmd() url.Values {
	result := url.Values{}

	return result
}
