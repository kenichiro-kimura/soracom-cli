// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersListRolesCmdOperatorId holds value of 'operator_id' option
var UsersListRolesCmdOperatorId string

// UsersListRolesCmdUserName holds value of 'user_name' option
var UsersListRolesCmdUserName string

// UsersListRolesCmdOutputJSONL indicates to output with jsonl format
var UsersListRolesCmdOutputJSONL bool

func InitUsersListRolesCmd() {
	UsersListRolesCmd.Flags().StringVar(&UsersListRolesCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersListRolesCmd.Flags().StringVar(&UsersListRolesCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersListRolesCmd.Flags().BoolVar(&UsersListRolesCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	UsersListRolesCmd.RunE = UsersListRolesCmdRunE

	UsersCmd.AddCommand(UsersListRolesCmd)
}

// UsersListRolesCmd defines 'list-roles' subcommand
var UsersListRolesCmd = &cobra.Command{
	Use:   "list-roles",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/roles:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/roles:get:description`) + "\n\n" + createLinkToAPIReference("Role", "listUserRoles"),
}

func UsersListRolesCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectUsersListRolesCmdParams(ac)
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
		if UsersListRolesCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectUsersListRolesCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if UsersListRolesCmdOperatorId == "" {
		UsersListRolesCmdOperatorId = ac.OperatorID
	}

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersListRolesCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersListRolesCmd("/operators/{operator_id}/users/{user_name}/roles"),
		query:  buildQueryForUsersListRolesCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersListRolesCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersListRolesCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersListRolesCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersListRolesCmd() url.Values {
	result := url.Values{}

	return result
}
