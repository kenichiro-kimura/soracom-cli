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

// OperatorVerifyCmdToken holds value of 'token' option
var OperatorVerifyCmdToken string

// OperatorVerifyCmdBody holds contents of request body to be sent
var OperatorVerifyCmdBody string

func InitOperatorVerifyCmd() {
	OperatorVerifyCmd.Flags().StringVar(&OperatorVerifyCmdToken, "token", "", TRAPI(""))

	OperatorVerifyCmd.Flags().StringVar(&OperatorVerifyCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorVerifyCmd.RunE = OperatorVerifyCmdRunE

	OperatorCmd.AddCommand(OperatorVerifyCmd)
}

// OperatorVerifyCmd defines 'verify' subcommand
var OperatorVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: TRAPI("/operators/verify:post:summary"),
	Long:  TRAPI(`/operators/verify:post:description`) + "\n\n" + createLinkToAPIReference("Operator", "verifyOperator"),
}

func OperatorVerifyCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectOperatorVerifyCmdParams(ac)
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

func collectOperatorVerifyCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForOperatorVerifyCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("token", "token", "body", parsedBody, OperatorVerifyCmdToken)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorVerifyCmd("/operators/verify"),
		query:       buildQueryForOperatorVerifyCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorVerifyCmd(path string) string {

	return path
}

func buildQueryForOperatorVerifyCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForOperatorVerifyCmd() (string, error) {
	var result map[string]interface{}

	if OperatorVerifyCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorVerifyCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorVerifyCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if OperatorVerifyCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorVerifyCmdBody)
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

	if OperatorVerifyCmdToken != "" {
		result["token"] = OperatorVerifyCmdToken
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
