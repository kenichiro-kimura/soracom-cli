// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgDeleteMirroringPeerCmdIpaddr holds value of 'ipaddr' option
var VpgDeleteMirroringPeerCmdIpaddr string

// VpgDeleteMirroringPeerCmdVpgId holds value of 'vpg_id' option
var VpgDeleteMirroringPeerCmdVpgId string

func init() {
	VpgDeleteMirroringPeerCmd.Flags().StringVar(&VpgDeleteMirroringPeerCmdIpaddr, "ipaddr", "", TRAPI("IP address of mirroring peer"))

	VpgDeleteMirroringPeerCmd.Flags().StringVar(&VpgDeleteMirroringPeerCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))
	VpgCmd.AddCommand(VpgDeleteMirroringPeerCmd)
}

// VpgDeleteMirroringPeerCmd defines 'delete-mirroring-peer' subcommand
var VpgDeleteMirroringPeerCmd = &cobra.Command{
	Use:   "delete-mirroring-peer",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}:delete:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}:delete:description`),
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

		param, err := collectVpgDeleteMirroringPeerCmdParams(ac)
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

func collectVpgDeleteMirroringPeerCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("ipaddr", "ipaddr", "path", parsedBody, VpgDeleteMirroringPeerCmdIpaddr)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgDeleteMirroringPeerCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgDeleteMirroringPeerCmd("/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}"),
		query:  buildQueryForVpgDeleteMirroringPeerCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgDeleteMirroringPeerCmd(path string) string {

	escapedIpaddr := url.PathEscape(VpgDeleteMirroringPeerCmdIpaddr)

	path = strReplace(path, "{"+"ipaddr"+"}", escapedIpaddr, -1)

	escapedVpgId := url.PathEscape(VpgDeleteMirroringPeerCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgDeleteMirroringPeerCmd() url.Values {
	result := url.Values{}

	return result
}
