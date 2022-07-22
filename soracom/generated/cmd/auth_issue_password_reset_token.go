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

// AuthIssuePasswordResetTokenCmdEmail holds value of 'email' option
var AuthIssuePasswordResetTokenCmdEmail string

// AuthIssuePasswordResetTokenCmdBody holds contents of request body to be sent
var AuthIssuePasswordResetTokenCmdBody string

func init() {
	AuthIssuePasswordResetTokenCmd.Flags().StringVar(&AuthIssuePasswordResetTokenCmdEmail, "email", "", TRAPI(""))

	AuthIssuePasswordResetTokenCmd.Flags().StringVar(&AuthIssuePasswordResetTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	AuthCmd.AddCommand(AuthIssuePasswordResetTokenCmd)
}

// AuthIssuePasswordResetTokenCmd defines 'issue-password-reset-token' subcommand
var AuthIssuePasswordResetTokenCmd = &cobra.Command{
	Use:   "issue-password-reset-token",
	Short: TRAPI("/auth/password_reset_token/issue:post:summary"),
	Long:  TRAPI(`/auth/password_reset_token/issue:post:description`) + "\n\n" + createLinkToAPIReference("Auth", "issuePasswordResetToken"),
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

		param, err := collectAuthIssuePasswordResetTokenCmdParams(ac)
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

func collectAuthIssuePasswordResetTokenCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForAuthIssuePasswordResetTokenCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("email", "email", "body", parsedBody, AuthIssuePasswordResetTokenCmdEmail)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForAuthIssuePasswordResetTokenCmd("/auth/password_reset_token/issue"),
		query:       buildQueryForAuthIssuePasswordResetTokenCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForAuthIssuePasswordResetTokenCmd(path string) string {

	return path
}

func buildQueryForAuthIssuePasswordResetTokenCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForAuthIssuePasswordResetTokenCmd() (string, error) {
	var result map[string]interface{}

	if AuthIssuePasswordResetTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(AuthIssuePasswordResetTokenCmdBody, "@") {
			fname := strings.TrimPrefix(AuthIssuePasswordResetTokenCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if AuthIssuePasswordResetTokenCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(AuthIssuePasswordResetTokenCmdBody)
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

	if AuthIssuePasswordResetTokenCmdEmail != "" {
		result["email"] = AuthIssuePasswordResetTokenCmdEmail
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
