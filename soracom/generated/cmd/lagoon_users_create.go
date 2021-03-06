package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonUsersCreateCmdRole holds value of 'role' option
var LagoonUsersCreateCmdRole string

// LagoonUsersCreateCmdUserEmail holds value of 'userEmail' option
var LagoonUsersCreateCmdUserEmail string

// LagoonUsersCreateCmdUserPassword holds value of 'userPassword' option
var LagoonUsersCreateCmdUserPassword string

// LagoonUsersCreateCmdBody holds contents of request body to be sent
var LagoonUsersCreateCmdBody string

func init() {
	LagoonUsersCreateCmd.Flags().StringVar(&LagoonUsersCreateCmdRole, "role", "", TRAPI(""))

	LagoonUsersCreateCmd.Flags().StringVar(&LagoonUsersCreateCmdUserEmail, "user-email", "", TRAPI(""))

	LagoonUsersCreateCmd.Flags().StringVar(&LagoonUsersCreateCmdUserPassword, "user-password", "", TRAPI(""))

	LagoonUsersCreateCmd.Flags().StringVar(&LagoonUsersCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LagoonUsersCmd.AddCommand(LagoonUsersCreateCmd)
}

// LagoonUsersCreateCmd defines 'create' subcommand
var LagoonUsersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/lagoon/users:post:summary"),
	Long:  TRAPI(`/lagoon/users:post:description`),
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

		param, err := collectLagoonUsersCreateCmdParams(ac)
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

func collectLagoonUsersCreateCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonUsersCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLagoonUsersCreateCmd("/lagoon/users"),
		query:       buildQueryForLagoonUsersCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonUsersCreateCmd(path string) string {

	return path
}

func buildQueryForLagoonUsersCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLagoonUsersCreateCmd() (string, error) {
	var result map[string]interface{}

	if LagoonUsersCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonUsersCreateCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUsersCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LagoonUsersCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonUsersCreateCmdBody)
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

	if LagoonUsersCreateCmdRole != "" {
		result["role"] = LagoonUsersCreateCmdRole
	}

	if LagoonUsersCreateCmdUserEmail != "" {
		result["userEmail"] = LagoonUsersCreateCmdUserEmail
	}

	if LagoonUsersCreateCmdUserPassword != "" {
		result["userPassword"] = LagoonUsersCreateCmdUserPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
