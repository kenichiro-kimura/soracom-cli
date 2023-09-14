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

// VpgCreatePacketCaptureSessionCmdPrefix holds value of 'prefix' option
var VpgCreatePacketCaptureSessionCmdPrefix string

// VpgCreatePacketCaptureSessionCmdVpgId holds value of 'vpg_id' option
var VpgCreatePacketCaptureSessionCmdVpgId string

// VpgCreatePacketCaptureSessionCmdDuration holds value of 'duration' option
var VpgCreatePacketCaptureSessionCmdDuration int64

// VpgCreatePacketCaptureSessionCmdBody holds contents of request body to be sent
var VpgCreatePacketCaptureSessionCmdBody string

func InitVpgCreatePacketCaptureSessionCmd() {
	VpgCreatePacketCaptureSessionCmd.Flags().StringVar(&VpgCreatePacketCaptureSessionCmdPrefix, "prefix", "", TRAPI(""))

	VpgCreatePacketCaptureSessionCmd.Flags().StringVar(&VpgCreatePacketCaptureSessionCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))

	VpgCreatePacketCaptureSessionCmd.Flags().Int64Var(&VpgCreatePacketCaptureSessionCmdDuration, "duration", 0, TRAPI(""))

	VpgCreatePacketCaptureSessionCmd.Flags().StringVar(&VpgCreatePacketCaptureSessionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCreatePacketCaptureSessionCmd.RunE = VpgCreatePacketCaptureSessionCmdRunE

	VpgCmd.AddCommand(VpgCreatePacketCaptureSessionCmd)
}

// VpgCreatePacketCaptureSessionCmd defines 'create-packet-capture-session' subcommand
var VpgCreatePacketCaptureSessionCmd = &cobra.Command{
	Use:   "create-packet-capture-session",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/packet_capture_sessions:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/packet_capture_sessions:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "createPacketCaptureSession"),
}

func VpgCreatePacketCaptureSessionCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgCreatePacketCaptureSessionCmdParams(ac)
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

func collectVpgCreatePacketCaptureSessionCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgCreatePacketCaptureSessionCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgCreatePacketCaptureSessionCmdVpgId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("duration", "duration", "body", parsedBody, VpgCreatePacketCaptureSessionCmdDuration)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgCreatePacketCaptureSessionCmd("/virtual_private_gateways/{vpg_id}/packet_capture_sessions"),
		query:       buildQueryForVpgCreatePacketCaptureSessionCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgCreatePacketCaptureSessionCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgCreatePacketCaptureSessionCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgCreatePacketCaptureSessionCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgCreatePacketCaptureSessionCmd() (string, error) {
	var result map[string]interface{}

	if VpgCreatePacketCaptureSessionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgCreatePacketCaptureSessionCmdBody, "@") {
			fname := strings.TrimPrefix(VpgCreatePacketCaptureSessionCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if VpgCreatePacketCaptureSessionCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgCreatePacketCaptureSessionCmdBody)
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

	if VpgCreatePacketCaptureSessionCmdPrefix != "" {
		result["prefix"] = VpgCreatePacketCaptureSessionCmdPrefix
	}

	if VpgCreatePacketCaptureSessionCmd.Flags().Lookup("duration").Changed {
		result["duration"] = VpgCreatePacketCaptureSessionCmdDuration
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
