// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersDetachRoleCmdOperatorId holds value of 'operator_id' option
var UsersDetachRoleCmdOperatorId string

// UsersDetachRoleCmdRoleId holds value of 'role_id' option
var UsersDetachRoleCmdRoleId string

// UsersDetachRoleCmdUserName holds value of 'user_name' option
var UsersDetachRoleCmdUserName string

func InitUsersDetachRoleCmd() {
	UsersDetachRoleCmd.Flags().StringVar(&UsersDetachRoleCmdOperatorId, "operator-id", "", TRAPI("Operator ID."))

	UsersDetachRoleCmd.Flags().StringVar(&UsersDetachRoleCmdRoleId, "role-id", "", TRAPI("Role ID."))

	UsersDetachRoleCmd.Flags().StringVar(&UsersDetachRoleCmdUserName, "user-name", "", TRAPI("SAM user name."))

	UsersDetachRoleCmd.RunE = UsersDetachRoleCmdRunE

	UsersCmd.AddCommand(UsersDetachRoleCmd)
}

// UsersDetachRoleCmd defines 'detach-role' subcommand
var UsersDetachRoleCmd = &cobra.Command{
	Use:   "detach-role",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/roles/{role_id}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/roles/{role_id}:delete:description`) + "\n\n" + createLinkToAPIReference("Role", "detachRole"),
}

func UsersDetachRoleCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectUsersDetachRoleCmdParams(ac)
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

func collectUsersDetachRoleCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if UsersDetachRoleCmdOperatorId == "" {
		UsersDetachRoleCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	err = checkIfRequiredStringParameterIsSupplied("role_id", "role-id", "path", parsedBody, UsersDetachRoleCmdRoleId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersDetachRoleCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForUsersDetachRoleCmd("/operators/{operator_id}/users/{user_name}/roles/{role_id}"),
		query:  buildQueryForUsersDetachRoleCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersDetachRoleCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersDetachRoleCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedRoleId := url.PathEscape(UsersDetachRoleCmdRoleId)

	path = strReplace(path, "{"+"role_id"+"}", escapedRoleId, -1)

	escapedUserName := url.PathEscape(UsersDetachRoleCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersDetachRoleCmd() url.Values {
	result := url.Values{}

	return result
}
