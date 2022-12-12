// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OperatorAuthKeysListCmdOperatorId holds value of 'operator_id' option
var OperatorAuthKeysListCmdOperatorId string

// OperatorAuthKeysListCmdOutputJSONL indicates to output with jsonl format
var OperatorAuthKeysListCmdOutputJSONL bool

func init() {
	OperatorAuthKeysListCmd.Flags().StringVar(&OperatorAuthKeysListCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	OperatorAuthKeysListCmd.Flags().BoolVar(&OperatorAuthKeysListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	OperatorAuthKeysCmd.AddCommand(OperatorAuthKeysListCmd)
}

// OperatorAuthKeysListCmd defines 'list' subcommand
var OperatorAuthKeysListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/auth_keys:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/auth_keys:get:description`) + "\n\n" + createLinkToAPIReference("Operator", "listOperatorAuthKeys"),
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

		param, err := collectOperatorAuthKeysListCmdParams(ac)
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
			if OperatorAuthKeysListCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectOperatorAuthKeysListCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorAuthKeysListCmdOperatorId == "" {
		OperatorAuthKeysListCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForOperatorAuthKeysListCmd("/operators/{operator_id}/auth_keys"),
		query:  buildQueryForOperatorAuthKeysListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorAuthKeysListCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorAuthKeysListCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorAuthKeysListCmd() url.Values {
	result := url.Values{}

	return result
}
