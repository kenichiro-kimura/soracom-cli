package cmd

import (
	"encoding/json"

	"io/ioutil"

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

	RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdPermission, "permission", "", TRAPI(""))

	RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdRoleId, "role-id", "", TRAPI("role_id"))

	RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	RolesCmd.AddCommand(RolesUpdateCmd)
}

// RolesUpdateCmd defines 'update' subcommand
var RolesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/operators/{operator_id}/roles/{role_id}:put:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles/{role_id}:put:description`),
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

		param, err := collectRolesUpdateCmdParams(ac)
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

func collectRolesUpdateCmdParams(ac *apiClient) (*apiParams, error) {

	if RolesUpdateCmdOperatorId == "" {
		RolesUpdateCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForRolesUpdateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForRolesUpdateCmd("/operators/{operator_id}/roles/{role_id}"),
		query:       buildQueryForRolesUpdateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForRolesUpdateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", RolesUpdateCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"role_id"+"}", RolesUpdateCmdRoleId, -1)

	return path
}

func buildQueryForRolesUpdateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
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
