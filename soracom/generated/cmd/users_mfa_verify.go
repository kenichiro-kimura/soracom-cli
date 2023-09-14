// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// UsersMfaVerifyCmdMfaOTPCode holds value of 'mfaOTPCode' option
var UsersMfaVerifyCmdMfaOTPCode string

// UsersMfaVerifyCmdOperatorId holds value of 'operator_id' option
var UsersMfaVerifyCmdOperatorId string

// UsersMfaVerifyCmdUserName holds value of 'user_name' option
var UsersMfaVerifyCmdUserName string

// UsersMfaVerifyCmdBody holds contents of request body to be sent
var UsersMfaVerifyCmdBody string

func InitUsersMfaVerifyCmd() {
	UsersMfaVerifyCmd.Flags().StringVar(&UsersMfaVerifyCmdMfaOTPCode, "mfa-otpcode", "", TRAPI(""))

	UsersMfaVerifyCmd.Flags().StringVar(&UsersMfaVerifyCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersMfaVerifyCmd.Flags().StringVar(&UsersMfaVerifyCmdUserName, "user-name", "", TRAPI("SAM user name"))

	UsersMfaVerifyCmd.Flags().StringVar(&UsersMfaVerifyCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	UsersMfaVerifyCmd.RunE = UsersMfaVerifyCmdRunE

	UsersMfaCmd.AddCommand(UsersMfaVerifyCmd)
}

// UsersMfaVerifyCmd defines 'verify' subcommand
var UsersMfaVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/mfa/verify:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/mfa/verify:post:description`) + "\n\n" + createLinkToAPIReference("User", "verifyUserMFA"),
}

func UsersMfaVerifyCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectUsersMfaVerifyCmdParams(ac)
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
}

func collectUsersMfaVerifyCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if UsersMfaVerifyCmdOperatorId == "" {
		UsersMfaVerifyCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	body, err = buildBodyForUsersMfaVerifyCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersMfaVerifyCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForUsersMfaVerifyCmd("/operators/{operator_id}/users/{user_name}/mfa/verify"),
		query:       buildQueryForUsersMfaVerifyCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersMfaVerifyCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersMfaVerifyCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersMfaVerifyCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersMfaVerifyCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForUsersMfaVerifyCmd() (string, error) {
	var result map[string]interface{}

	if UsersMfaVerifyCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(UsersMfaVerifyCmdBody, "@") {
			fname := strings.TrimPrefix(UsersMfaVerifyCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if UsersMfaVerifyCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(UsersMfaVerifyCmdBody)
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

	if UsersMfaVerifyCmdMfaOTPCode != "" {
		result["mfaOTPCode"] = UsersMfaVerifyCmdMfaOTPCode
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
