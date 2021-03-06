package cmd

import (
	"encoding/json"

	"io/ioutil"

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

func init() {
	VpgOpenGateCmd.Flags().StringVar(&VpgOpenGateCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgOpenGateCmd.Flags().Int64Var(&VpgOpenGateCmdVxlanId, "vxlan-id", 0, TRAPI(""))

	VpgOpenGateCmd.Flags().BoolVar(&VpgOpenGateCmdPrivacySeparatorEnabled, "privacy-separator-enabled", false, TRAPI(""))

	VpgOpenGateCmd.Flags().StringVar(&VpgOpenGateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgOpenGateCmd)
}

// VpgOpenGateCmd defines 'open-gate' subcommand
var VpgOpenGateCmd = &cobra.Command{
	Use:   "open-gate",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/open:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/open:post:description`),
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

		param, err := collectVpgOpenGateCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectVpgOpenGateCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForVpgOpenGateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgOpenGateCmd("/virtual_private_gateways/{vpg_id}/gate/open"),
		query:       buildQueryForVpgOpenGateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgOpenGateCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgOpenGateCmdVpgId, -1)

	return path
}

func buildQueryForVpgOpenGateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgOpenGateCmd() (string, error) {
	var result map[string]interface{}

	if VpgOpenGateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgOpenGateCmdBody, "@") {
			fname := strings.TrimPrefix(VpgOpenGateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if VpgOpenGateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
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

	if VpgOpenGateCmdVxlanId != 0 {
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
