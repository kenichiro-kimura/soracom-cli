package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersAttachRoleCmdOperatorId holds value of 'operator_id' option
var UsersAttachRoleCmdOperatorId string

// UsersAttachRoleCmdRoleId holds value of 'roleId' option
var UsersAttachRoleCmdRoleId string

// UsersAttachRoleCmdUserName holds value of 'user_name' option
var UsersAttachRoleCmdUserName string

// UsersAttachRoleCmdBody holds contents of request body to be sent
var UsersAttachRoleCmdBody string

func init() {
	UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdRoleId, "role-id", "", TRAPI(""))

	UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	UsersCmd.AddCommand(UsersAttachRoleCmd)
}

// UsersAttachRoleCmd defines 'attach-role' subcommand
var UsersAttachRoleCmd = &cobra.Command{
	Use:   "attach-role",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/roles:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/roles:post:description`),
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

		param, err := collectUsersAttachRoleCmdParams(ac)
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

func collectUsersAttachRoleCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersAttachRoleCmdOperatorId == "" {
		UsersAttachRoleCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForUsersAttachRoleCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForUsersAttachRoleCmd("/operators/{operator_id}/users/{user_name}/roles"),
		query:       buildQueryForUsersAttachRoleCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForUsersAttachRoleCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersAttachRoleCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersAttachRoleCmdUserName, -1)

	return path
}

func buildQueryForUsersAttachRoleCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForUsersAttachRoleCmd() (string, error) {
	var result map[string]interface{}

	if UsersAttachRoleCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(UsersAttachRoleCmdBody, "@") {
			fname := strings.TrimPrefix(UsersAttachRoleCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if UsersAttachRoleCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(UsersAttachRoleCmdBody)
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

	if UsersAttachRoleCmdRoleId != "" {
		result["roleId"] = UsersAttachRoleCmdRoleId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
