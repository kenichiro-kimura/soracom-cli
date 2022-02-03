// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OperatorEnableMfaCmdOperatorId holds value of 'operator_id' option
var OperatorEnableMfaCmdOperatorId string

func init() {
	OperatorEnableMfaCmd.Flags().StringVar(&OperatorEnableMfaCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	OperatorCmd.AddCommand(OperatorEnableMfaCmd)
}

// OperatorEnableMfaCmd defines 'enable-mfa' subcommand
var OperatorEnableMfaCmd = &cobra.Command{
	Use:   "enable-mfa",
	Short: TRAPI("/operators/{operator_id}/mfa:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/mfa:post:description`),
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

		param, err := collectOperatorEnableMfaCmdParams(ac)
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
	},
}

func collectOperatorEnableMfaCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorEnableMfaCmdOperatorId == "" {
		OperatorEnableMfaCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForOperatorEnableMfaCmd("/operators/{operator_id}/mfa"),
		query:  buildQueryForOperatorEnableMfaCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorEnableMfaCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorEnableMfaCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorEnableMfaCmd() url.Values {
	result := url.Values{}

	return result
}
