// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersAuthKeysListCmdOperatorId holds value of 'operator_id' option
var UsersAuthKeysListCmdOperatorId string

// UsersAuthKeysListCmdUserName holds value of 'user_name' option
var UsersAuthKeysListCmdUserName string

// UsersAuthKeysListCmdOutputJSONL indicates to output with jsonl format
var UsersAuthKeysListCmdOutputJSONL bool

func InitUsersAuthKeysListCmd() {
	UsersAuthKeysListCmd.Flags().StringVar(&UsersAuthKeysListCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersAuthKeysListCmd.Flags().StringVar(&UsersAuthKeysListCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersAuthKeysListCmd.Flags().BoolVar(&UsersAuthKeysListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	UsersAuthKeysListCmd.RunE = UsersAuthKeysListCmdRunE

	UsersAuthKeysCmd.AddCommand(UsersAuthKeysListCmd)
}

// UsersAuthKeysListCmd defines 'list' subcommand
var UsersAuthKeysListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/auth_keys:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/auth_keys:get:description`) + "\n\n" + createLinkToAPIReference("User", "listUserAuthKeys"),
}

func UsersAuthKeysListCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectUsersAuthKeysListCmdParams(ac)
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
		if UsersAuthKeysListCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectUsersAuthKeysListCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if UsersAuthKeysListCmdOperatorId == "" {
		UsersAuthKeysListCmdOperatorId = ac.OperatorID
	}

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersAuthKeysListCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersAuthKeysListCmd("/operators/{operator_id}/users/{user_name}/auth_keys"),
		query:  buildQueryForUsersAuthKeysListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersAuthKeysListCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersAuthKeysListCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersAuthKeysListCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersAuthKeysListCmd() url.Values {
	result := url.Values{}

	return result
}
