// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock holds value of 'destinationCidrBlock' option
var VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock string

// VpgCreateVpcPeeringConnectionCmdPeerOwnerId holds value of 'peerOwnerId' option
var VpgCreateVpcPeeringConnectionCmdPeerOwnerId string

// VpgCreateVpcPeeringConnectionCmdPeerRegion holds value of 'peerRegion' option
var VpgCreateVpcPeeringConnectionCmdPeerRegion string

// VpgCreateVpcPeeringConnectionCmdPeerVpcId holds value of 'peerVpcId' option
var VpgCreateVpcPeeringConnectionCmdPeerVpcId string

// VpgCreateVpcPeeringConnectionCmdVpgId holds value of 'vpg_id' option
var VpgCreateVpcPeeringConnectionCmdVpgId string

// VpgCreateVpcPeeringConnectionCmdBody holds contents of request body to be sent
var VpgCreateVpcPeeringConnectionCmdBody string

func init() {
	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock, "destination-cidr-block", "", TRAPI(""))

	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdPeerOwnerId, "peer-owner-id", "", TRAPI(""))

	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdPeerRegion, "peer-region", "", TRAPI(""))

	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdPeerVpcId, "peer-vpc-id", "", TRAPI(""))

	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	VpgCmd.AddCommand(VpgCreateVpcPeeringConnectionCmd)
}

// VpgCreateVpcPeeringConnectionCmd defines 'create-vpc-peering-connection' subcommand
var VpgCreateVpcPeeringConnectionCmd = &cobra.Command{
	Use:   "create-vpc-peering-connection",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/vpc_peering_connections:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/vpc_peering_connections:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "createVpcPeeringConnection"),
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

		param, err := collectVpgCreateVpcPeeringConnectionCmdParams(ac)
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

func collectVpgCreateVpcPeeringConnectionCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgCreateVpcPeeringConnectionCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgCreateVpcPeeringConnectionCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgCreateVpcPeeringConnectionCmd("/virtual_private_gateways/{vpg_id}/vpc_peering_connections"),
		query:       buildQueryForVpgCreateVpcPeeringConnectionCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgCreateVpcPeeringConnectionCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgCreateVpcPeeringConnectionCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgCreateVpcPeeringConnectionCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgCreateVpcPeeringConnectionCmd() (string, error) {
	var result map[string]interface{}

	if VpgCreateVpcPeeringConnectionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgCreateVpcPeeringConnectionCmdBody, "@") {
			fname := strings.TrimPrefix(VpgCreateVpcPeeringConnectionCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if VpgCreateVpcPeeringConnectionCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgCreateVpcPeeringConnectionCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock != "" {
		result["destinationCidrBlock"] = VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock
	}

	if VpgCreateVpcPeeringConnectionCmdPeerOwnerId != "" {
		result["peerOwnerId"] = VpgCreateVpcPeeringConnectionCmdPeerOwnerId
	}

	if VpgCreateVpcPeeringConnectionCmdPeerRegion != "" {
		result["peerRegion"] = VpgCreateVpcPeeringConnectionCmdPeerRegion
	}

	if VpgCreateVpcPeeringConnectionCmdPeerVpcId != "" {
		result["peerVpcId"] = VpgCreateVpcPeeringConnectionCmdPeerVpcId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
