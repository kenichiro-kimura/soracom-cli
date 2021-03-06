package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgSetRedirectionCmdDescription holds value of 'description' option
var VpgSetRedirectionCmdDescription string

// VpgSetRedirectionCmdGateway holds value of 'gateway' option
var VpgSetRedirectionCmdGateway string

// VpgSetRedirectionCmdVpgId holds value of 'vpg_id' option
var VpgSetRedirectionCmdVpgId string

// VpgSetRedirectionCmdEnabled holds value of 'enabled' option
var VpgSetRedirectionCmdEnabled bool

// VpgSetRedirectionCmdBody holds contents of request body to be sent
var VpgSetRedirectionCmdBody string

func init() {
	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdDescription, "description", "", TRAPI(""))

	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdGateway, "gateway", "", TRAPI(""))

	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))

	VpgSetRedirectionCmd.Flags().BoolVar(&VpgSetRedirectionCmdEnabled, "enabled", false, TRAPI(""))

	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgSetRedirectionCmd)
}

// VpgSetRedirectionCmd defines 'set-redirection' subcommand
var VpgSetRedirectionCmd = &cobra.Command{
	Use:   "set-redirection",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/junction/set_redirection:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/set_redirection:post:description`),
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

		param, err := collectVpgSetRedirectionCmdParams(ac)
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

func collectVpgSetRedirectionCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForVpgSetRedirectionCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgSetRedirectionCmd("/virtual_private_gateways/{vpg_id}/junction/set_redirection"),
		query:       buildQueryForVpgSetRedirectionCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgSetRedirectionCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgSetRedirectionCmdVpgId, -1)

	return path
}

func buildQueryForVpgSetRedirectionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgSetRedirectionCmd() (string, error) {
	var result map[string]interface{}

	if VpgSetRedirectionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgSetRedirectionCmdBody, "@") {
			fname := strings.TrimPrefix(VpgSetRedirectionCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if VpgSetRedirectionCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgSetRedirectionCmdBody)
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

	if VpgSetRedirectionCmdDescription != "" {
		result["description"] = VpgSetRedirectionCmdDescription
	}

	if VpgSetRedirectionCmdGateway != "" {
		result["gateway"] = VpgSetRedirectionCmdGateway
	}

	if VpgSetRedirectionCmdEnabled != false {
		result["enabled"] = VpgSetRedirectionCmdEnabled
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
