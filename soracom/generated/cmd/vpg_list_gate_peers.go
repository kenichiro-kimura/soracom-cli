// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgListGatePeersCmdVpgId holds value of 'vpg_id' option
var VpgListGatePeersCmdVpgId string

// VpgListGatePeersCmdOutputJSONL indicates to output with jsonl format
var VpgListGatePeersCmdOutputJSONL bool

func InitVpgListGatePeersCmd() {
	VpgListGatePeersCmd.Flags().StringVar(&VpgListGatePeersCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgListGatePeersCmd.Flags().BoolVar(&VpgListGatePeersCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	VpgListGatePeersCmd.RunE = VpgListGatePeersCmdRunE

	VpgCmd.AddCommand(VpgListGatePeersCmd)
}

// VpgListGatePeersCmd defines 'list-gate-peers' subcommand
var VpgListGatePeersCmd = &cobra.Command{
	Use:   "list-gate-peers",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/peers:get:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/peers:get:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "listGatePeers"),
}

func VpgListGatePeersCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgListGatePeersCmdParams(ac)
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
		if VpgListGatePeersCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectVpgListGatePeersCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgListGatePeersCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForVpgListGatePeersCmd("/virtual_private_gateways/{vpg_id}/gate/peers"),
		query:  buildQueryForVpgListGatePeersCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgListGatePeersCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgListGatePeersCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgListGatePeersCmd() url.Values {
	result := url.Values{}

	return result
}
