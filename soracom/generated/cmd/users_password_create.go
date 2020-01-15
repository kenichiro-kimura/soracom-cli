// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"io/ioutil"

	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// UsersPasswordCreateCmdOperatorId holds value of 'operator_id' option
var UsersPasswordCreateCmdOperatorId string

// UsersPasswordCreateCmdPassword holds value of 'password' option
var UsersPasswordCreateCmdPassword string

// UsersPasswordCreateCmdUserName holds value of 'user_name' option
var UsersPasswordCreateCmdUserName string

// UsersPasswordCreateCmdBody holds contents of request body to be sent
var UsersPasswordCreateCmdBody string

func init() {
	UsersPasswordCreateCmd.Flags().StringVar(&UsersPasswordCreateCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersPasswordCreateCmd.Flags().StringVar(&UsersPasswordCreateCmdPassword, "password", "", TRAPI(""))

	UsersPasswordCreateCmd.Flags().StringVar(&UsersPasswordCreateCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersPasswordCreateCmd.MarkFlagRequired("user-name")

	UsersPasswordCreateCmd.Flags().StringVar(&UsersPasswordCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	UsersPasswordCmd.AddCommand(UsersPasswordCreateCmd)
}

// UsersPasswordCreateCmd defines 'create' subcommand
var UsersPasswordCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/password:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/password:post:description`),
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

		param, err := collectUsersPasswordCreateCmdParams(ac)
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

		return prettyPrintStringAsJSON(body)

	},
}

func collectUsersPasswordCreateCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersPasswordCreateCmdOperatorId == "" {
		UsersPasswordCreateCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForUsersPasswordCreateCmd()
	if err != nil {
		return nil, err
	}

	contentType := "application/json"

	return &apiParams{
		method:      "POST",
		path:        buildPathForUsersPasswordCreateCmd("/operators/{operator_id}/users/{user_name}/password"),
		query:       buildQueryForUsersPasswordCreateCmd(),
		contentType: contentType,
		body:        body,
	}, nil
}

func buildPathForUsersPasswordCreateCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersPasswordCreateCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersPasswordCreateCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersPasswordCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForUsersPasswordCreateCmd() (string, error) {
	var result map[string]interface{}

	if UsersPasswordCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(UsersPasswordCreateCmdBody, "@") {
			fname := strings.TrimPrefix(UsersPasswordCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if UsersPasswordCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(UsersPasswordCreateCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if UsersPasswordCreateCmdPassword != "" {
		result["password"] = UsersPasswordCreateCmdPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
