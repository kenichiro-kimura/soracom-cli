package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	CredentialsCmd.AddCommand(CredentialsListCmd)
}

// CredentialsListCmd defines 'list' subcommand
var CredentialsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/credentials:get:summary"),
	Long:  TRAPI(`/credentials:get:description`),
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

		param, err := collectCredentialsListCmdParams(ac)
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

func collectCredentialsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForCredentialsListCmd("/credentials"),
		query:  buildQueryForCredentialsListCmd(),
	}, nil
}

func buildPathForCredentialsListCmd(path string) string {

	return path
}

func buildQueryForCredentialsListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
