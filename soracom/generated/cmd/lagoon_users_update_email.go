package cmd

import (
	"encoding/json"

	"io/ioutil"

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

func collectLagoonUsersUpdateEmailCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonUsersUpdateEmailCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUsersUpdateEmailCmd("/lagoon/users/{lagoon_user_id}/email"),
		query:       buildQueryForLagoonUsersUpdateEmailCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonUsersUpdateEmailCmd(path string) string {

	path = strings.Replace(path, "{"+"lagoon_user_id"+"}", sprintf("%d", LagoonUsersUpdateEmailCmdLagoonUserId), -1)

	return path
}

func buildQueryForLagoonUsersUpdateEmailCmd() string {
	result := []string{}

	return strings.Join(result, "&")
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
