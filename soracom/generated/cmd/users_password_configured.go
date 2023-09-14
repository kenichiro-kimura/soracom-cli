// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersPasswordConfiguredCmdOperatorId holds value of 'operator_id' option
var UsersPasswordConfiguredCmdOperatorId string

// UsersPasswordConfiguredCmdUserName holds value of 'user_name' option
var UsersPasswordConfiguredCmdUserName string

func InitUsersPasswordConfiguredCmd() {
	UsersPasswordConfiguredCmd.Flags().StringVar(&UsersPasswordConfiguredCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersPasswordConfiguredCmd.Flags().StringVar(&UsersPasswordConfiguredCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersPasswordConfiguredCmd.RunE = UsersPasswordConfiguredCmdRunE

	UsersPasswordCmd.AddCommand(UsersPasswordConfiguredCmd)
}

// UsersPasswordConfiguredCmd defines 'configured' subcommand
var UsersPasswordConfiguredCmd = &cobra.Command{
	Use:   "configured",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/password:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/password:get:description`) + "\n\n" + createLinkToAPIReference("User", "hasUserPassword"),
}

func UsersPasswordConfiguredCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectUsersPasswordConfiguredCmdParams(ac)
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

func collectUsersPasswordConfiguredCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if UsersPasswordConfiguredCmdOperatorId == "" {
		UsersPasswordConfiguredCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersPasswordConfiguredCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersPasswordConfiguredCmd("/operators/{operator_id}/users/{user_name}/password"),
		query:  buildQueryForUsersPasswordConfiguredCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersPasswordConfiguredCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersPasswordConfiguredCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersPasswordConfiguredCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersPasswordConfiguredCmd() url.Values {
	result := url.Values{}

	return result
}
