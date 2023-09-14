// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// CredentialsListCmdOutputJSONL indicates to output with jsonl format
var CredentialsListCmdOutputJSONL bool

func InitCredentialsListCmd() {
	CredentialsListCmd.Flags().BoolVar(&CredentialsListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	CredentialsListCmd.RunE = CredentialsListCmdRunE

	CredentialsCmd.AddCommand(CredentialsListCmd)
}

// CredentialsListCmd defines 'list' subcommand
var CredentialsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/credentials:get:summary"),
	Long:  TRAPI(`/credentials:get:description`) + "\n\n" + createLinkToAPIReference("Credential", "listCredentials"),
}

func CredentialsListCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectCredentialsListCmdParams(ac)
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
		if CredentialsListCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectCredentialsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForCredentialsListCmd("/credentials"),
		query:  buildQueryForCredentialsListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForCredentialsListCmd(path string) string {

	return path
}

func buildQueryForCredentialsListCmd() url.Values {
	result := url.Values{}

	return result
}
