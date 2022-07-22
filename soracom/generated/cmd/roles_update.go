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

// RolesUpdateCmdDescription holds value of 'description' option
var RolesUpdateCmdDescription string

// RolesUpdateCmdOperatorId holds value of 'operator_id' option
var RolesUpdateCmdOperatorId string

// RolesUpdateCmdPermission holds value of 'permission' option
var RolesUpdateCmdPermission string

// RolesUpdateCmdRoleId holds value of 'role_id' option
var RolesUpdateCmdRoleId string

// RolesUpdateCmdBody holds contents of request body to be sent
var RolesUpdateCmdBody string

func init() {
	RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdDescription, "description", "", TRAPI(""))

	RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdPermission, "permission", "", TRAPI("Permission as JSON"))

	RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdRoleId, "role-id", "", TRAPI("role_id"))

	RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	RolesCmd.AddCommand(RolesUpdateCmd)
}

// RolesUpdateCmd defines 'update' subcommand
var RolesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/operators/{operator_id}/roles/{role_id}:put:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles/{role_id}:put:description`) + "\n\n" + createLinkToAPIReference("Role", "updateRole"),
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

		param, err := collectRolesUpdateCmdParams(ac)
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

func collectRolesUpdateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if RolesUpdateCmdOperatorId == "" {
		RolesUpdateCmdOperatorId = ac.OperatorID
	}

	body, err = buildBodyForRolesUpdateCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("permission", "permission", "body", parsedBody, RolesUpdateCmdPermission)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("role_id", "role-id", "path", parsedBody, RolesUpdateCmdRoleId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForRolesUpdateCmd("/operators/{operator_id}/roles/{role_id}"),
		query:       buildQueryForRolesUpdateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForRolesUpdateCmd(path string) string {

	escapedOperatorId := url.PathEscape(RolesUpdateCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedRoleId := url.PathEscape(RolesUpdateCmdRoleId)

	path = strReplace(path, "{"+"role_id"+"}", escapedRoleId, -1)

	return path
}

func buildQueryForRolesUpdateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForRolesUpdateCmd() (string, error) {
	var result map[string]interface{}

	if RolesUpdateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(RolesUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(RolesUpdateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if RolesUpdateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(RolesUpdateCmdBody)
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

	if RolesUpdateCmdDescription != "" {
		result["description"] = RolesUpdateCmdDescription
	}

	if RolesUpdateCmdPermission != "" {
		result["permission"] = RolesUpdateCmdPermission
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
