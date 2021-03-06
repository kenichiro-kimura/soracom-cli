package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SandboxOperatorsDeleteCmdOperatorId holds value of 'operator_id' option
var SandboxOperatorsDeleteCmdOperatorId string

func init() {
	SandboxOperatorsDeleteCmd.Flags().StringVar(&SandboxOperatorsDeleteCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	SandboxOperatorsCmd.AddCommand(SandboxOperatorsDeleteCmd)
}

// SandboxOperatorsDeleteCmd defines 'delete' subcommand
var SandboxOperatorsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/sandbox/operators/{operator_id}:delete:summary"),
	Long:  TRAPI(`/sandbox/operators/{operator_id}:delete:description`),
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

		param, err := collectSandboxOperatorsDeleteCmdParams(ac)
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

func collectSandboxOperatorsDeleteCmdParams(ac *apiClient) (*apiParams, error) {

	if SandboxOperatorsDeleteCmdOperatorId == "" {
		SandboxOperatorsDeleteCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSandboxOperatorsDeleteCmd("/sandbox/operators/{operator_id}"),
		query:  buildQueryForSandboxOperatorsDeleteCmd(),
	}, nil
}

func buildPathForSandboxOperatorsDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", SandboxOperatorsDeleteCmdOperatorId, -1)

	return path
}

func buildQueryForSandboxOperatorsDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
