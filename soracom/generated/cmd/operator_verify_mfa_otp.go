package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorVerifyMfaOtpCmdMfaOTPCode holds value of 'mfaOTPCode' option
var OperatorVerifyMfaOtpCmdMfaOTPCode string

// OperatorVerifyMfaOtpCmdOperatorId holds value of 'operator_id' option
var OperatorVerifyMfaOtpCmdOperatorId string

// OperatorVerifyMfaOtpCmdBody holds contents of request body to be sent
var OperatorVerifyMfaOtpCmdBody string

func init() {
	OperatorVerifyMfaOtpCmd.Flags().StringVar(&OperatorVerifyMfaOtpCmdMfaOTPCode, "mfa-otpcode", "", TRAPI(""))

	OperatorVerifyMfaOtpCmd.Flags().StringVar(&OperatorVerifyMfaOtpCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorVerifyMfaOtpCmd.Flags().StringVar(&OperatorVerifyMfaOtpCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorCmd.AddCommand(OperatorVerifyMfaOtpCmd)
}

// OperatorVerifyMfaOtpCmd defines 'verify-mfa-otp' subcommand
var OperatorVerifyMfaOtpCmd = &cobra.Command{
	Use:   "verify-mfa-otp",
	Short: TRAPI("/operators/{operator_id}/mfa/verify:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/mfa/verify:post:description`),
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

		param, err := collectOperatorVerifyMfaOtpCmdParams(ac)
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

func collectOperatorVerifyMfaOtpCmdParams(ac *apiClient) (*apiParams, error) {

	if OperatorVerifyMfaOtpCmdOperatorId == "" {
		OperatorVerifyMfaOtpCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForOperatorVerifyMfaOtpCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorVerifyMfaOtpCmd("/operators/{operator_id}/mfa/verify"),
		query:       buildQueryForOperatorVerifyMfaOtpCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOperatorVerifyMfaOtpCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorVerifyMfaOtpCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorVerifyMfaOtpCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForOperatorVerifyMfaOtpCmd() (string, error) {
	var result map[string]interface{}

	if OperatorVerifyMfaOtpCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorVerifyMfaOtpCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorVerifyMfaOtpCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OperatorVerifyMfaOtpCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorVerifyMfaOtpCmdBody)
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

	if OperatorVerifyMfaOtpCmdMfaOTPCode != "" {
		result["mfaOTPCode"] = OperatorVerifyMfaOtpCmdMfaOTPCode
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
