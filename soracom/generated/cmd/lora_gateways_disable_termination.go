package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraGatewaysDisableTerminationCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysDisableTerminationCmdGatewayId string

func init() {
	LoraGatewaysDisableTerminationCmd.Flags().StringVar(&LoraGatewaysDisableTerminationCmdGatewayId, "gateway-id", "", TR("lora_gateways.disable_termination.parameters.gateway_id.description"))

	LoraGatewaysCmd.AddCommand(LoraGatewaysDisableTerminationCmd)
}

// LoraGatewaysDisableTerminationCmd defines 'disable-termination' subcommand
var LoraGatewaysDisableTerminationCmd = &cobra.Command{
	Use:   "disable-termination",
	Short: TR("lora_gateways.disable_termination.summary"),
	Long:  TR(`lora_gateways.disable_termination.description`),
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

		param, err := collectLoraGatewaysDisableTerminationCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectLoraGatewaysDisableTerminationCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraGatewaysDisableTerminationCmd("/lora_gateways/{gateway_id}/disable_termination"),
		query:  buildQueryForLoraGatewaysDisableTerminationCmd(),
	}, nil
}

func buildPathForLoraGatewaysDisableTerminationCmd(path string) string {

	path = strings.Replace(path, "{"+"gateway_id"+"}", LoraGatewaysDisableTerminationCmdGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysDisableTerminationCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}