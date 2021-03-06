package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersDefaultPermissionsUpdateCmdOperatorId holds value of 'operator_id' option
var UsersDefaultPermissionsUpdateCmdOperatorId string

// UsersDefaultPermissionsUpdateCmdPermissions holds value of 'permissions' option
var UsersDefaultPermissionsUpdateCmdPermissions string

// UsersDefaultPermissionsUpdateCmdBody holds contents of request body to be sent
var UsersDefaultPermissionsUpdateCmdBody string

func init() {
	UsersDefaultPermissionsUpdateCmd.Flags().StringVar(&UsersDefaultPermissionsUpdateCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersDefaultPermissionsUpdateCmd.Flags().StringVar(&UsersDefaultPermissionsUpdateCmdPermissions, "permissions", "", TRAPI(""))

	UsersDefaultPermissionsUpdateCmd.Flags().StringVar(&UsersDefaultPermissionsUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	UsersDefaultPermissionsCmd.AddCommand(UsersDefaultPermissionsUpdateCmd)
}

// UsersDefaultPermissionsUpdateCmd defines 'update' subcommand
var UsersDefaultPermissionsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/operators/{operator_id}/users/default_permissions:put:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/default_permissions:put:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectUsersDefaultPermissionsUpdateCmdParams(ac)
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

func collectUsersDefaultPermissionsUpdateCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersDefaultPermissionsUpdateCmdOperatorId == "" {
		UsersDefaultPermissionsUpdateCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForUsersDefaultPermissionsUpdateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForUsersDefaultPermissionsUpdateCmd("/operators/{operator_id}/users/default_permissions"),
		query:       buildQueryForUsersDefaultPermissionsUpdateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForUsersDefaultPermissionsUpdateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersDefaultPermissionsUpdateCmdOperatorId, -1)

	return path
}

func buildQueryForUsersDefaultPermissionsUpdateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForUsersDefaultPermissionsUpdateCmd() (string, error) {
	var result map[string]interface{}

	if UsersDefaultPermissionsUpdateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(UsersDefaultPermissionsUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(UsersDefaultPermissionsUpdateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if UsersDefaultPermissionsUpdateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(UsersDefaultPermissionsUpdateCmdBody)
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

	if UsersDefaultPermissionsUpdateCmdPermissions != "" {
		result["permissions"] = UsersDefaultPermissionsUpdateCmdPermissions
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
