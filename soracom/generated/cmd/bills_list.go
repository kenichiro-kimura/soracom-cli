package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	BillsCmd.AddCommand(BillsListCmd)
}

// BillsListCmd defines 'list' subcommand
var BillsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/bills:get:summary"),
	Long:  TRAPI(`/bills:get:description`),
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

		param, err := collectBillsListCmdParams(ac)
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

func collectBillsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForBillsListCmd("/bills"),
		query:  buildQueryForBillsListCmd(),
	}, nil
}

func buildPathForBillsListCmd(path string) string {

	return path
}

func buildQueryForBillsListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
