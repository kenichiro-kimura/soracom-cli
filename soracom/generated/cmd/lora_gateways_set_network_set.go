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

// LoraGatewaysSetNetworkSetCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysSetNetworkSetCmdGatewayId string

// LoraGatewaysSetNetworkSetCmdNetworkSetId holds value of 'networkSetId' option
var LoraGatewaysSetNetworkSetCmdNetworkSetId string

// LoraGatewaysSetNetworkSetCmdBody holds contents of request body to be sent
var LoraGatewaysSetNetworkSetCmdBody string

func InitLoraGatewaysSetNetworkSetCmd() {
	LoraGatewaysSetNetworkSetCmd.Flags().StringVar(&LoraGatewaysSetNetworkSetCmdGatewayId, "gateway-id", "", TRAPI("ID of the target LoRaWAN gateway."))

	LoraGatewaysSetNetworkSetCmd.Flags().StringVar(&LoraGatewaysSetNetworkSetCmdNetworkSetId, "network-set-id", "", TRAPI(""))

	LoraGatewaysSetNetworkSetCmd.Flags().StringVar(&LoraGatewaysSetNetworkSetCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraGatewaysSetNetworkSetCmd.RunE = LoraGatewaysSetNetworkSetCmdRunE

	LoraGatewaysCmd.AddCommand(LoraGatewaysSetNetworkSetCmd)
}

// LoraGatewaysSetNetworkSetCmd defines 'set-network-set' subcommand
var LoraGatewaysSetNetworkSetCmd = &cobra.Command{
	Use:   "set-network-set",
	Short: TRAPI("/lora_gateways/{gateway_id}/set_network_set:post:summary"),
	Long:  TRAPI(`/lora_gateways/{gateway_id}/set_network_set:post:description`) + "\n\n" + createLinkToAPIReference("LoraGateway", "setLoraNetworkSet"),
}

func LoraGatewaysSetNetworkSetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectLoraGatewaysSetNetworkSetCmdParams(ac)
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

func collectLoraGatewaysSetNetworkSetCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLoraGatewaysSetNetworkSetCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("gateway_id", "gateway-id", "path", parsedBody, LoraGatewaysSetNetworkSetCmdGatewayId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraGatewaysSetNetworkSetCmd("/lora_gateways/{gateway_id}/set_network_set"),
		query:       buildQueryForLoraGatewaysSetNetworkSetCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraGatewaysSetNetworkSetCmd(path string) string {

	escapedGatewayId := url.PathEscape(LoraGatewaysSetNetworkSetCmdGatewayId)

	path = strReplace(path, "{"+"gateway_id"+"}", escapedGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysSetNetworkSetCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLoraGatewaysSetNetworkSetCmd() (string, error) {
	var result map[string]interface{}

	if LoraGatewaysSetNetworkSetCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LoraGatewaysSetNetworkSetCmdBody, "@") {
			fname := strings.TrimPrefix(LoraGatewaysSetNetworkSetCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LoraGatewaysSetNetworkSetCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraGatewaysSetNetworkSetCmdBody)
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

	if LoraGatewaysSetNetworkSetCmdNetworkSetId != "" {
		result["networkSetId"] = LoraGatewaysSetNetworkSetCmdNetworkSetId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
