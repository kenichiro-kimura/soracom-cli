// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// UsersUpdateCmdDescription holds value of 'description' option
var UsersUpdateCmdDescription string

// UsersUpdateCmdOperatorId holds value of 'operator_id' option
var UsersUpdateCmdOperatorId string

// UsersUpdateCmdUserName holds value of 'user_name' option
var UsersUpdateCmdUserName string

// UsersUpdateCmdBody holds contents of request body to be sent
var UsersUpdateCmdBody string

func init() {
	UsersUpdateCmd.Flags().StringVar(&UsersUpdateCmdDescription, "description", "", TRAPI(""))

	UsersUpdateCmd.Flags().StringVar(&UsersUpdateCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersUpdateCmd.Flags().StringVar(&UsersUpdateCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersUpdateCmd.Flags().StringVar(&UsersUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	UsersCmd.AddCommand(UsersUpdateCmd)
}

// UsersUpdateCmd defines 'update' subcommand
var UsersUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}:put:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}:put:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

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

		param, err := collectUsersUpdateCmdParams(ac)
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
	},
}

func collectUsersUpdateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if UsersUpdateCmdOperatorId == "" {
		UsersUpdateCmdOperatorId = ac.OperatorID
	}

	body, err = buildBodyForUsersUpdateCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersUpdateCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForUsersUpdateCmd("/operators/{operator_id}/users/{user_name}"),
		query:       buildQueryForUsersUpdateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersUpdateCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersUpdateCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersUpdateCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersUpdateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForUsersUpdateCmd() (string, error) {
	var result map[string]interface{}

	if UsersUpdateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(UsersUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(UsersUpdateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if UsersUpdateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(UsersUpdateCmdBody)
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

	if UsersUpdateCmdDescription != "" {
		result["description"] = UsersUpdateCmdDescription
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
