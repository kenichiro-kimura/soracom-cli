// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsActivateCmdSimId holds value of 'sim_id' option
var SimsActivateCmdSimId string

func InitSimsActivateCmd() {
	SimsActivateCmd.Flags().StringVar(&SimsActivateCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsActivateCmd.RunE = SimsActivateCmdRunE

	SimsCmd.AddCommand(SimsActivateCmd)
}

// SimsActivateCmd defines 'activate' subcommand
var SimsActivateCmd = &cobra.Command{
	Use:   "activate",
	Short: TRAPI("/sims/{sim_id}/activate:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/activate:post:description`) + "\n\n" + createLinkToAPIReference("Sim", "activateSim"),
}

func SimsActivateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSimsActivateCmdParams(ac)
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

func collectSimsActivateCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsActivateCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsActivateCmd("/sims/{sim_id}/activate"),
		query:  buildQueryForSimsActivateCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsActivateCmd(path string) string {

	escapedSimId := url.PathEscape(SimsActivateCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsActivateCmd() url.Values {
	result := url.Values{}

	return result
}
