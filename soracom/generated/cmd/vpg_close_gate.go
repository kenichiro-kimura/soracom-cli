// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgCloseGateCmdVpgId holds value of 'vpg_id' option
var VpgCloseGateCmdVpgId string

func InitVpgCloseGateCmd() {
	VpgCloseGateCmd.Flags().StringVar(&VpgCloseGateCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgCloseGateCmd.RunE = VpgCloseGateCmdRunE

	VpgCmd.AddCommand(VpgCloseGateCmd)
}

// VpgCloseGateCmd defines 'close-gate' subcommand
var VpgCloseGateCmd = &cobra.Command{
	Use:   "close-gate",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/close:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/close:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "closeGate"),
}

func VpgCloseGateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgCloseGateCmdParams(ac)
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

func collectVpgCloseGateCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgCloseGateCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgCloseGateCmd("/virtual_private_gateways/{vpg_id}/gate/close"),
		query:  buildQueryForVpgCloseGateCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgCloseGateCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgCloseGateCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgCloseGateCmd() url.Values {
	result := url.Values{}

	return result
}
