// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// VpgOpenGateCmdVpgId holds value of 'vpg_id' option
var VpgOpenGateCmdVpgId string

// VpgOpenGateCmdVxlanId holds value of 'vxlanId' option
var VpgOpenGateCmdVxlanId int64

// VpgOpenGateCmdPrivacySeparatorEnabled holds value of 'privacySeparatorEnabled' option
var VpgOpenGateCmdPrivacySeparatorEnabled bool

// VpgOpenGateCmdBody holds contents of request body to be sent
var VpgOpenGateCmdBody string

func InitVpgOpenGateCmd() {
	VpgOpenGateCmd.Flags().StringVar(&VpgOpenGateCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgOpenGateCmd.Flags().Int64Var(&VpgOpenGateCmdVxlanId, "vxlan-id", 10, TRAPI(""))

	VpgOpenGateCmd.Flags().BoolVar(&VpgOpenGateCmdPrivacySeparatorEnabled, "privacy-separator-enabled", false, TRAPI(""))

	VpgOpenGateCmd.Flags().StringVar(&VpgOpenGateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgOpenGateCmd.RunE = VpgOpenGateCmdRunE

	VpgCmd.AddCommand(VpgOpenGateCmd)
}

// VpgOpenGateCmd defines 'open-gate' subcommand
var VpgOpenGateCmd = &cobra.Command{
	Use:   "open-gate",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/open:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/open:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "openGate"),
}

func VpgOpenGateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgOpenGateCmdParams(ac)
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

func collectVpgOpenGateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgOpenGateCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgOpenGateCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgOpenGateCmd("/virtual_private_gateways/{vpg_id}/gate/open"),
		query:       buildQueryForVpgOpenGateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgOpenGateCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgOpenGateCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgOpenGateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgOpenGateCmd() (string, error) {
	var result map[string]interface{}

	if VpgOpenGateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgOpenGateCmdBody, "@") {
			fname := strings.TrimPrefix(VpgOpenGateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if VpgOpenGateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgOpenGateCmdBody)
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

	if VpgOpenGateCmd.Flags().Lookup("vxlan-id").Changed {
		result["vxlanId"] = VpgOpenGateCmdVxlanId
	}

	if VpgOpenGateCmdPrivacySeparatorEnabled != false {
		result["privacySeparatorEnabled"] = VpgOpenGateCmdPrivacySeparatorEnabled
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
