package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersDefaultPermissionsGetCmdOperatorId holds value of 'operator_id' option
var UsersDefaultPermissionsGetCmdOperatorId string

func init() {
	UsersDefaultPermissionsGetCmd.Flags().StringVar(&UsersDefaultPermissionsGetCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersDefaultPermissionsCmd.AddCommand(UsersDefaultPermissionsGetCmd)
}

// UsersDefaultPermissionsGetCmd defines 'get' subcommand
var UsersDefaultPermissionsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/users/default_permissions:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/default_permissions:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectUsersDefaultPermissionsGetCmdParams(ac)
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

func collectUsersDefaultPermissionsGetCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersDefaultPermissionsGetCmdOperatorId == "" {
		UsersDefaultPermissionsGetCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersDefaultPermissionsGetCmd("/operators/{operator_id}/users/default_permissions"),
		query:  buildQueryForUsersDefaultPermissionsGetCmd(),
	}, nil
}

func buildPathForUsersDefaultPermissionsGetCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersDefaultPermissionsGetCmdOperatorId, -1)

	return path
}

func buildQueryForUsersDefaultPermissionsGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
