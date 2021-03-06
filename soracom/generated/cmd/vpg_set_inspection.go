package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgSetInspectionCmdVpgId holds value of 'vpg_id' option
var VpgSetInspectionCmdVpgId string

// VpgSetInspectionCmdEnabled holds value of 'enabled' option
var VpgSetInspectionCmdEnabled bool

// VpgSetInspectionCmdBody holds contents of request body to be sent
var VpgSetInspectionCmdBody string

func init() {
	VpgSetInspectionCmd.Flags().StringVar(&VpgSetInspectionCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))

	VpgSetInspectionCmd.Flags().BoolVar(&VpgSetInspectionCmdEnabled, "enabled", false, TRAPI(""))

	VpgSetInspectionCmd.Flags().StringVar(&VpgSetInspectionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgSetInspectionCmd)
}

// VpgSetInspectionCmd defines 'set-inspection' subcommand
var VpgSetInspectionCmd = &cobra.Command{
	Use:   "set-inspection",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/junction/set_inspection:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/set_inspection:post:description`),
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

		param, err := collectVpgSetInspectionCmdParams(ac)
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

func collectVpgSetInspectionCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForVpgSetInspectionCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgSetInspectionCmd("/virtual_private_gateways/{vpg_id}/junction/set_inspection"),
		query:       buildQueryForVpgSetInspectionCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgSetInspectionCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgSetInspectionCmdVpgId, -1)

	return path
}

func buildQueryForVpgSetInspectionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgSetInspectionCmd() (string, error) {
	var result map[string]interface{}

	if VpgSetInspectionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgSetInspectionCmdBody, "@") {
			fname := strings.TrimPrefix(VpgSetInspectionCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if VpgSetInspectionCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgSetInspectionCmdBody)
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

	if VpgSetInspectionCmdEnabled != false {
		result["enabled"] = VpgSetInspectionCmdEnabled
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
