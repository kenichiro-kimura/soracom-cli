// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// EmailsListCmdOperatorId holds value of 'operator_id' option
var EmailsListCmdOperatorId string

// EmailsListCmdOutputJSONL indicates to output with jsonl format
var EmailsListCmdOutputJSONL bool

func InitEmailsListCmd() {
	EmailsListCmd.Flags().StringVar(&EmailsListCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	EmailsListCmd.Flags().BoolVar(&EmailsListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	EmailsListCmd.RunE = EmailsListCmdRunE

	EmailsCmd.AddCommand(EmailsListCmd)
}

// EmailsListCmd defines 'list' subcommand
var EmailsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/emails:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/emails:get:description`) + "\n\n" + createLinkToAPIReference("Email", "listEmails"),
}

func EmailsListCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectEmailsListCmdParams(ac)
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
		if EmailsListCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectEmailsListCmdParams(ac *apiClient) (*apiParams, error) {
	if EmailsListCmdOperatorId == "" {
		EmailsListCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForEmailsListCmd("/operators/{operator_id}/emails"),
		query:  buildQueryForEmailsListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForEmailsListCmd(path string) string {

	escapedOperatorId := url.PathEscape(EmailsListCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForEmailsListCmd() url.Values {
	result := url.Values{}

	return result
}
