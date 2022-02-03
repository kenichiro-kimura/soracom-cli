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

// LoraNetworkSetsRevokePermissionCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsRevokePermissionCmdNsId string

// LoraNetworkSetsRevokePermissionCmdOperatorId holds value of 'operatorId' option
var LoraNetworkSetsRevokePermissionCmdOperatorId string

// LoraNetworkSetsRevokePermissionCmdBody holds contents of request body to be sent
var LoraNetworkSetsRevokePermissionCmdBody string

func init() {
	LoraNetworkSetsRevokePermissionCmd.Flags().StringVar(&LoraNetworkSetsRevokePermissionCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsRevokePermissionCmd.Flags().StringVar(&LoraNetworkSetsRevokePermissionCmdOperatorId, "operator-id", "", TRAPI(""))

	LoraNetworkSetsRevokePermissionCmd.Flags().StringVar(&LoraNetworkSetsRevokePermissionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsRevokePermissionCmd)
}

// LoraNetworkSetsRevokePermissionCmd defines 'revoke-permission' subcommand
var LoraNetworkSetsRevokePermissionCmd = &cobra.Command{
	Use:   "revoke-permission",
	Short: TRAPI("/lora_network_sets/{ns_id}/revoke_permission:post:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}/revoke_permission:post:description`),
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

		param, err := collectLoraNetworkSetsRevokePermissionCmdParams(ac)
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

func collectLoraNetworkSetsRevokePermissionCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLoraNetworkSetsRevokePermissionCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("ns_id", "ns-id", "path", parsedBody, LoraNetworkSetsRevokePermissionCmdNsId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraNetworkSetsRevokePermissionCmd("/lora_network_sets/{ns_id}/revoke_permission"),
		query:       buildQueryForLoraNetworkSetsRevokePermissionCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraNetworkSetsRevokePermissionCmd(path string) string {

	escapedNsId := url.PathEscape(LoraNetworkSetsRevokePermissionCmdNsId)

	path = strReplace(path, "{"+"ns_id"+"}", escapedNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsRevokePermissionCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLoraNetworkSetsRevokePermissionCmd() (string, error) {
	var result map[string]interface{}

	if LoraNetworkSetsRevokePermissionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LoraNetworkSetsRevokePermissionCmdBody, "@") {
			fname := strings.TrimPrefix(LoraNetworkSetsRevokePermissionCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LoraNetworkSetsRevokePermissionCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraNetworkSetsRevokePermissionCmdBody)
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

	if LoraNetworkSetsRevokePermissionCmdOperatorId != "" {
		result["operatorId"] = LoraNetworkSetsRevokePermissionCmdOperatorId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
