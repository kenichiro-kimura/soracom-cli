// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsGetCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsGetCmdNsId string

func InitLoraNetworkSetsGetCmd() {
	LoraNetworkSetsGetCmd.Flags().StringVar(&LoraNetworkSetsGetCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsGetCmd.RunE = LoraNetworkSetsGetCmdRunE

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsGetCmd)
}

// LoraNetworkSetsGetCmd defines 'get' subcommand
var LoraNetworkSetsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/lora_network_sets/{ns_id}:get:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}:get:description`) + "\n\n" + createLinkToAPIReference("LoraNetworkSet", "getLoraNetworkSet"),
}

func LoraNetworkSetsGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectLoraNetworkSetsGetCmdParams(ac)
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

func collectLoraNetworkSetsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("ns_id", "ns-id", "path", parsedBody, LoraNetworkSetsGetCmdNsId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraNetworkSetsGetCmd("/lora_network_sets/{ns_id}"),
		query:  buildQueryForLoraNetworkSetsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraNetworkSetsGetCmd(path string) string {

	escapedNsId := url.PathEscape(LoraNetworkSetsGetCmdNsId)

	path = strReplace(path, "{"+"ns_id"+"}", escapedNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsGetCmd() url.Values {
	result := url.Values{}

	return result
}
