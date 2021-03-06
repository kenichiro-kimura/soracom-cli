package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorVerifyMfaRevokeTokenCmdBackupCode holds value of 'backupCode' option
var OperatorVerifyMfaRevokeTokenCmdBackupCode string

// OperatorVerifyMfaRevokeTokenCmdEmail holds value of 'email' option
var OperatorVerifyMfaRevokeTokenCmdEmail string

// OperatorVerifyMfaRevokeTokenCmdPassword holds value of 'password' option
var OperatorVerifyMfaRevokeTokenCmdPassword string

// OperatorVerifyMfaRevokeTokenCmdToken holds value of 'token' option
var OperatorVerifyMfaRevokeTokenCmdToken string

// OperatorVerifyMfaRevokeTokenCmdBody holds contents of request body to be sent
var OperatorVerifyMfaRevokeTokenCmdBody string

func init() {
	OperatorVerifyMfaRevokeTokenCmd.Flags().StringVar(&OperatorVerifyMfaRevokeTokenCmdBackupCode, "backup-code", "", TRAPI(""))

	OperatorVerifyMfaRevokeTokenCmd.Flags().StringVar(&OperatorVerifyMfaRevokeTokenCmdEmail, "email", "", TRAPI(""))

	OperatorVerifyMfaRevokeTokenCmd.Flags().StringVar(&OperatorVerifyMfaRevokeTokenCmdPassword, "password", "", TRAPI(""))

	OperatorVerifyMfaRevokeTokenCmd.Flags().StringVar(&OperatorVerifyMfaRevokeTokenCmdToken, "token", "", TRAPI(""))

	OperatorVerifyMfaRevokeTokenCmd.Flags().StringVar(&OperatorVerifyMfaRevokeTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorCmd.AddCommand(OperatorVerifyMfaRevokeTokenCmd)
}

// OperatorVerifyMfaRevokeTokenCmd defines 'verify-mfa-revoke-token' subcommand
var OperatorVerifyMfaRevokeTokenCmd = &cobra.Command{
	Use:   "verify-mfa-revoke-token",
	Short: TRAPI("/operators/mfa_revoke_token/verify:post:summary"),
	Long:  TRAPI(`/operators/mfa_revoke_token/verify:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectOperatorVerifyMfaRevokeTokenCmdParams(ac)
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

func collectOperatorVerifyMfaRevokeTokenCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForOperatorVerifyMfaRevokeTokenCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorVerifyMfaRevokeTokenCmd("/operators/mfa_revoke_token/verify"),
		query:       buildQueryForOperatorVerifyMfaRevokeTokenCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOperatorVerifyMfaRevokeTokenCmd(path string) string {

	return path
}

func buildQueryForOperatorVerifyMfaRevokeTokenCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForOperatorVerifyMfaRevokeTokenCmd() (string, error) {
	var result map[string]interface{}

	if OperatorVerifyMfaRevokeTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorVerifyMfaRevokeTokenCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorVerifyMfaRevokeTokenCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OperatorVerifyMfaRevokeTokenCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorVerifyMfaRevokeTokenCmdBody)
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

	if OperatorVerifyMfaRevokeTokenCmdBackupCode != "" {
		result["backupCode"] = OperatorVerifyMfaRevokeTokenCmdBackupCode
	}

	if OperatorVerifyMfaRevokeTokenCmdEmail != "" {
		result["email"] = OperatorVerifyMfaRevokeTokenCmdEmail
	}

	if OperatorVerifyMfaRevokeTokenCmdPassword != "" {
		result["password"] = OperatorVerifyMfaRevokeTokenCmdPassword
	}

	if OperatorVerifyMfaRevokeTokenCmdToken != "" {
		result["token"] = OperatorVerifyMfaRevokeTokenCmdToken
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
