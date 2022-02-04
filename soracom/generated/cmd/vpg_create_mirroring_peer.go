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

// VpgCreateMirroringPeerCmdDescription holds value of 'description' option
var VpgCreateMirroringPeerCmdDescription string

// VpgCreateMirroringPeerCmdIpAddress holds value of 'ipAddress' option
var VpgCreateMirroringPeerCmdIpAddress string

// VpgCreateMirroringPeerCmdProtocol holds value of 'protocol' option
var VpgCreateMirroringPeerCmdProtocol string

// VpgCreateMirroringPeerCmdVpgId holds value of 'vpg_id' option
var VpgCreateMirroringPeerCmdVpgId string

// VpgCreateMirroringPeerCmdEnabled holds value of 'enabled' option
var VpgCreateMirroringPeerCmdEnabled bool

// VpgCreateMirroringPeerCmdBody holds contents of request body to be sent
var VpgCreateMirroringPeerCmdBody string

func init() {
	VpgCreateMirroringPeerCmd.Flags().StringVar(&VpgCreateMirroringPeerCmdDescription, "description", "", TRAPI(""))

	VpgCreateMirroringPeerCmd.Flags().StringVar(&VpgCreateMirroringPeerCmdIpAddress, "ip-address", "", TRAPI(""))

	VpgCreateMirroringPeerCmd.Flags().StringVar(&VpgCreateMirroringPeerCmdProtocol, "protocol", "", TRAPI(""))

	VpgCreateMirroringPeerCmd.Flags().StringVar(&VpgCreateMirroringPeerCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))

	VpgCreateMirroringPeerCmd.Flags().BoolVar(&VpgCreateMirroringPeerCmdEnabled, "enabled", false, TRAPI(""))

	VpgCreateMirroringPeerCmd.Flags().StringVar(&VpgCreateMirroringPeerCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	VpgCmd.AddCommand(VpgCreateMirroringPeerCmd)
}

// VpgCreateMirroringPeerCmd defines 'create-mirroring-peer' subcommand
var VpgCreateMirroringPeerCmd = &cobra.Command{
	Use:   "create-mirroring-peer",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/junction/mirroring/peers:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/mirroring/peers:post:description`),
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

		param, err := collectVpgCreateMirroringPeerCmdParams(ac)
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

func collectVpgCreateMirroringPeerCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgCreateMirroringPeerCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgCreateMirroringPeerCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgCreateMirroringPeerCmd("/virtual_private_gateways/{vpg_id}/junction/mirroring/peers"),
		query:       buildQueryForVpgCreateMirroringPeerCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgCreateMirroringPeerCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgCreateMirroringPeerCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgCreateMirroringPeerCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgCreateMirroringPeerCmd() (string, error) {
	var result map[string]interface{}

	if VpgCreateMirroringPeerCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgCreateMirroringPeerCmdBody, "@") {
			fname := strings.TrimPrefix(VpgCreateMirroringPeerCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if VpgCreateMirroringPeerCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgCreateMirroringPeerCmdBody)
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

	if VpgCreateMirroringPeerCmdDescription != "" {
		result["description"] = VpgCreateMirroringPeerCmdDescription
	}

	if VpgCreateMirroringPeerCmdIpAddress != "" {
		result["ipAddress"] = VpgCreateMirroringPeerCmdIpAddress
	}

	if VpgCreateMirroringPeerCmdProtocol != "" {
		result["protocol"] = VpgCreateMirroringPeerCmdProtocol
	}

	if VpgCreateMirroringPeerCmdEnabled != false {
		result["enabled"] = VpgCreateMirroringPeerCmdEnabled
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
