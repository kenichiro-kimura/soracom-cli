package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraGatewaysTerminateCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysTerminateCmdGatewayId string

func init() {
	LoraGatewaysTerminateCmd.Flags().StringVar(&LoraGatewaysTerminateCmdGatewayId, "gateway-id", "", TRAPI("Device ID of the target LoRa gateway."))

	LoraGatewaysCmd.AddCommand(LoraGatewaysTerminateCmd)
}

// LoraGatewaysTerminateCmd defines 'terminate' subcommand
var LoraGatewaysTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/lora_gateways/{gateway_id}/terminate:post:summary"),
	Long:  TRAPI(`/lora_gateways/{gateway_id}/terminate:post:description`),
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

		param, err := collectLoraGatewaysTerminateCmdParams(ac)
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

func collectLoraGatewaysTerminateCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraGatewaysTerminateCmd("/lora_gateways/{gateway_id}/terminate"),
		query:  buildQueryForLoraGatewaysTerminateCmd(),
	}, nil
}

func buildPathForLoraGatewaysTerminateCmd(path string) string {

	path = strings.Replace(path, "{"+"gateway_id"+"}", LoraGatewaysTerminateCmdGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysTerminateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
