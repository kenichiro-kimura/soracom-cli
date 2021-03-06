package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersMfaRevokeCmdOperatorId holds value of 'operator_id' option
var UsersMfaRevokeCmdOperatorId string

// UsersMfaRevokeCmdUserName holds value of 'user_name' option
var UsersMfaRevokeCmdUserName string

func init() {
	UsersMfaRevokeCmd.Flags().StringVar(&UsersMfaRevokeCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersMfaRevokeCmd.Flags().StringVar(&UsersMfaRevokeCmdUserName, "user-name", "", TRAPI("SAM user name"))

	UsersMfaCmd.AddCommand(UsersMfaRevokeCmd)
}

// UsersMfaRevokeCmd defines 'revoke' subcommand
var UsersMfaRevokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/mfa:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/mfa:delete:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectUsersMfaRevokeCmdParams(ac)
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

func collectUsersMfaRevokeCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersMfaRevokeCmdOperatorId == "" {
		UsersMfaRevokeCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForUsersMfaRevokeCmd("/operators/{operator_id}/users/{user_name}/mfa"),
		query:  buildQueryForUsersMfaRevokeCmd(),
	}, nil
}

func buildPathForUsersMfaRevokeCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersMfaRevokeCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersMfaRevokeCmdUserName, -1)

	return path
}

func buildQueryForUsersMfaRevokeCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
