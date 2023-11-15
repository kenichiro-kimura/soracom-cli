// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SandboxOperatorsDeleteCmdOperatorId holds value of 'operator_id' option
var SandboxOperatorsDeleteCmdOperatorId string

func InitSandboxOperatorsDeleteCmd() {
	SandboxOperatorsDeleteCmd.Flags().StringVar(&SandboxOperatorsDeleteCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	SandboxOperatorsDeleteCmd.RunE = SandboxOperatorsDeleteCmdRunE

	SandboxOperatorsCmd.AddCommand(SandboxOperatorsDeleteCmd)
}

// SandboxOperatorsDeleteCmd defines 'delete' subcommand
var SandboxOperatorsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/sandbox/operators/{operator_id}:delete:summary"),
	Long:  TRAPI(`/sandbox/operators/{operator_id}:delete:description`) + "\n\n" + createLinkToAPIReference("API Sandbox: Operator", "sandboxDeleteOperator"),
}

func SandboxOperatorsDeleteCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSandboxOperatorsDeleteCmdParams(ac)
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

func collectSandboxOperatorsDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	if SandboxOperatorsDeleteCmdOperatorId == "" {
		SandboxOperatorsDeleteCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSandboxOperatorsDeleteCmd("/sandbox/operators/{operator_id}"),
		query:  buildQueryForSandboxOperatorsDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSandboxOperatorsDeleteCmd(path string) string {

	escapedOperatorId := url.PathEscape(SandboxOperatorsDeleteCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForSandboxOperatorsDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
