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

// LagoonUsersUpdateEmailCmdUserEmail holds value of 'userEmail' option
var LagoonUsersUpdateEmailCmdUserEmail string

// LagoonUsersUpdateEmailCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUsersUpdateEmailCmdLagoonUserId int64

// LagoonUsersUpdateEmailCmdBody holds contents of request body to be sent
var LagoonUsersUpdateEmailCmdBody string

func init() {
	LagoonUsersUpdateEmailCmd.Flags().StringVar(&LagoonUsersUpdateEmailCmdUserEmail, "user-email", "", TRAPI(""))

	LagoonUsersUpdateEmailCmd.Flags().Int64Var(&LagoonUsersUpdateEmailCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUsersUpdateEmailCmd.Flags().StringVar(&LagoonUsersUpdateEmailCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonUsersCmd.AddCommand(LagoonUsersUpdateEmailCmd)
}

// LagoonUsersUpdateEmailCmd defines 'update-email' subcommand
var LagoonUsersUpdateEmailCmd = &cobra.Command{
	Use:   "update-email",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}/email:put:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}/email:put:description`),
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

		param, err := collectLagoonUsersUpdateEmailCmdParams(ac)
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

func collectLagoonUsersUpdateEmailCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLagoonUsersUpdateEmailCmd()
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

	err = checkIfRequiredIntegerParameterIsSupplied("lagoon_user_id", "lagoon-user-id", "path", parsedBody, LagoonUsersUpdateEmailCmdLagoonUserId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUsersUpdateEmailCmd("/lagoon/users/{lagoon_user_id}/email"),
		query:       buildQueryForLagoonUsersUpdateEmailCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonUsersUpdateEmailCmd(path string) string {

	path = strReplace(path, "{"+"lagoon_user_id"+"}", url.PathEscape(sprintf("%d", LagoonUsersUpdateEmailCmdLagoonUserId)), -1)

	return path
}

func buildQueryForLagoonUsersUpdateEmailCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonUsersUpdateEmailCmd() (string, error) {
	var result map[string]interface{}

	if LagoonUsersUpdateEmailCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonUsersUpdateEmailCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUsersUpdateEmailCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LagoonUsersUpdateEmailCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonUsersUpdateEmailCmdBody)
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

	if LagoonUsersUpdateEmailCmdUserEmail != "" {
		result["userEmail"] = LagoonUsersUpdateEmailCmdUserEmail
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
