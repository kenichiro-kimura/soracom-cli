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

// UsersCreateCmdDescription holds value of 'description' option
var UsersCreateCmdDescription string

// UsersCreateCmdOperatorId holds value of 'operator_id' option
var UsersCreateCmdOperatorId string

// UsersCreateCmdUserName holds value of 'user_name' option
var UsersCreateCmdUserName string

// UsersCreateCmdBody holds contents of request body to be sent
var UsersCreateCmdBody string

func init() {
	UsersCreateCmd.Flags().StringVar(&UsersCreateCmdDescription, "description", "", TRAPI(""))

	UsersCreateCmd.Flags().StringVar(&UsersCreateCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersCreateCmd.Flags().StringVar(&UsersCreateCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersCreateCmd.Flags().StringVar(&UsersCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	UsersCmd.AddCommand(UsersCreateCmd)
}

// UsersCreateCmd defines 'create' subcommand
var UsersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}:post:description`) + "\n\n" + createLinkToAPIReference("User", "createUser"),
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

		param, err := collectUsersCreateCmdParams(ac)
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

func collectUsersCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if UsersCreateCmdOperatorId == "" {
		UsersCreateCmdOperatorId = ac.OperatorID
	}

	body, err = buildBodyForUsersCreateCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersCreateCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForUsersCreateCmd("/operators/{operator_id}/users/{user_name}"),
		query:       buildQueryForUsersCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersCreateCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersCreateCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersCreateCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForUsersCreateCmd() (string, error) {
	var result map[string]interface{}

	if UsersCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(UsersCreateCmdBody, "@") {
			fname := strings.TrimPrefix(UsersCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if UsersCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(UsersCreateCmdBody)
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

	if UsersCreateCmdDescription != "" {
		result["description"] = UsersCreateCmdDescription
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
