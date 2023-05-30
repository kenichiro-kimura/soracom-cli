// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersPermissionsGetCmdOperatorId holds value of 'operator_id' option
var UsersPermissionsGetCmdOperatorId string

// UsersPermissionsGetCmdUserName holds value of 'user_name' option
var UsersPermissionsGetCmdUserName string

func InitUsersPermissionsGetCmd() {
	UsersPermissionsGetCmd.Flags().StringVar(&UsersPermissionsGetCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersPermissionsGetCmd.Flags().StringVar(&UsersPermissionsGetCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersPermissionsGetCmd.RunE = UsersPermissionsGetCmdRunE

	UsersPermissionsCmd.AddCommand(UsersPermissionsGetCmd)
}

// UsersPermissionsGetCmd defines 'get' subcommand
var UsersPermissionsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/permission:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/permission:get:description`) + "\n\n" + createLinkToAPIReference("User", "getUserPermission"),
}

func UsersPermissionsGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectUsersPermissionsGetCmdParams(ac)
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

func collectUsersPermissionsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if UsersPermissionsGetCmdOperatorId == "" {
		UsersPermissionsGetCmdOperatorId = ac.OperatorID
	}

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersPermissionsGetCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersPermissionsGetCmd("/operators/{operator_id}/users/{user_name}/permission"),
		query:  buildQueryForUsersPermissionsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersPermissionsGetCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersPermissionsGetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersPermissionsGetCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersPermissionsGetCmd() url.Values {
	result := url.Values{}

	return result
}
