// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraGatewaysEnableTerminationCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysEnableTerminationCmdGatewayId string

func InitLoraGatewaysEnableTerminationCmd() {
	LoraGatewaysEnableTerminationCmd.Flags().StringVar(&LoraGatewaysEnableTerminationCmdGatewayId, "gateway-id", "", TRAPI("ID of the target LoRaWAN gateway."))

	LoraGatewaysEnableTerminationCmd.RunE = LoraGatewaysEnableTerminationCmdRunE

	LoraGatewaysCmd.AddCommand(LoraGatewaysEnableTerminationCmd)
}

// LoraGatewaysEnableTerminationCmd defines 'enable-termination' subcommand
var LoraGatewaysEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/lora_gateways/{gateway_id}/enable_termination:post:summary"),
	Long:  TRAPI(`/lora_gateways/{gateway_id}/enable_termination:post:description`) + "\n\n" + createLinkToAPIReference("LoraGateway", "enableTerminationOnLoraGateway"),
}

func LoraGatewaysEnableTerminationCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectLoraGatewaysEnableTerminationCmdParams(ac)
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

func collectLoraGatewaysEnableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("gateway_id", "gateway-id", "path", parsedBody, LoraGatewaysEnableTerminationCmdGatewayId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraGatewaysEnableTerminationCmd("/lora_gateways/{gateway_id}/enable_termination"),
		query:  buildQueryForLoraGatewaysEnableTerminationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraGatewaysEnableTerminationCmd(path string) string {

	escapedGatewayId := url.PathEscape(LoraGatewaysEnableTerminationCmdGatewayId)

	path = strReplace(path, "{"+"gateway_id"+"}", escapedGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysEnableTerminationCmd() url.Values {
	result := url.Values{}

	return result
}
