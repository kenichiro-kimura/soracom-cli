package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsAddPermissionCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsAddPermissionCmdNsId string

// LoraNetworkSetsAddPermissionCmdOperatorId holds value of 'operatorId' option
var LoraNetworkSetsAddPermissionCmdOperatorId string

// LoraNetworkSetsAddPermissionCmdBody holds contents of request body to be sent
var LoraNetworkSetsAddPermissionCmdBody string

func init() {
	LoraNetworkSetsAddPermissionCmd.Flags().StringVar(&LoraNetworkSetsAddPermissionCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsAddPermissionCmd.Flags().StringVar(&LoraNetworkSetsAddPermissionCmdOperatorId, "operator-id", "", TRAPI(""))

	LoraNetworkSetsAddPermissionCmd.Flags().StringVar(&LoraNetworkSetsAddPermissionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsAddPermissionCmd)
}

// LoraNetworkSetsAddPermissionCmd defines 'add-permission' subcommand
var LoraNetworkSetsAddPermissionCmd = &cobra.Command{
	Use:   "add-permission",
	Short: TRAPI("/lora_network_sets/{ns_id}/add_permission:post:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}/add_permission:post:description`),
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

		param, err := collectLoraNetworkSetsAddPermissionCmdParams(ac)
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

func collectLoraNetworkSetsAddPermissionCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLoraNetworkSetsAddPermissionCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraNetworkSetsAddPermissionCmd("/lora_network_sets/{ns_id}/add_permission"),
		query:       buildQueryForLoraNetworkSetsAddPermissionCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraNetworkSetsAddPermissionCmd(path string) string {

	path = strings.Replace(path, "{"+"ns_id"+"}", LoraNetworkSetsAddPermissionCmdNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsAddPermissionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraNetworkSetsAddPermissionCmd() (string, error) {
	var result map[string]interface{}

	if LoraNetworkSetsAddPermissionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LoraNetworkSetsAddPermissionCmdBody, "@") {
			fname := strings.TrimPrefix(LoraNetworkSetsAddPermissionCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LoraNetworkSetsAddPermissionCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraNetworkSetsAddPermissionCmdBody)
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

	if LoraNetworkSetsAddPermissionCmdOperatorId != "" {
		result["operatorId"] = LoraNetworkSetsAddPermissionCmdOperatorId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
