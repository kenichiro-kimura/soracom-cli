package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorUpdatePasswordCmdCurrentPassword holds value of 'currentPassword' option
var OperatorUpdatePasswordCmdCurrentPassword string

// OperatorUpdatePasswordCmdNewPassword holds value of 'newPassword' option
var OperatorUpdatePasswordCmdNewPassword string

// OperatorUpdatePasswordCmdOperatorId holds value of 'operator_id' option
var OperatorUpdatePasswordCmdOperatorId string

// OperatorUpdatePasswordCmdBody holds contents of request body to be sent
var OperatorUpdatePasswordCmdBody string

func init() {
	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdCurrentPassword, "current-password", "", TRAPI(""))

	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdNewPassword, "new-password", "", TRAPI(""))

	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorCmd.AddCommand(OperatorUpdatePasswordCmd)
}

// OperatorUpdatePasswordCmd defines 'update-password' subcommand
var OperatorUpdatePasswordCmd = &cobra.Command{
	Use:   "update-password",
	Short: TRAPI("/operators/{operator_id}/password:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/password:post:description`),
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

		param, err := collectOperatorUpdatePasswordCmdParams(ac)
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

func collectOperatorUpdatePasswordCmdParams(ac *apiClient) (*apiParams, error) {

	if OperatorUpdatePasswordCmdOperatorId == "" {
		OperatorUpdatePasswordCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForOperatorUpdatePasswordCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorUpdatePasswordCmd("/operators/{operator_id}/password"),
		query:       buildQueryForOperatorUpdatePasswordCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOperatorUpdatePasswordCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorUpdatePasswordCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorUpdatePasswordCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForOperatorUpdatePasswordCmd() (string, error) {
	var result map[string]interface{}

	if OperatorUpdatePasswordCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorUpdatePasswordCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorUpdatePasswordCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OperatorUpdatePasswordCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorUpdatePasswordCmdBody)
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

	if OperatorUpdatePasswordCmdCurrentPassword != "" {
		result["currentPassword"] = OperatorUpdatePasswordCmdCurrentPassword
	}

	if OperatorUpdatePasswordCmdNewPassword != "" {
		result["newPassword"] = OperatorUpdatePasswordCmdNewPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
