package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonCreateUserCmdRole holds value of 'role' option
var LagoonCreateUserCmdRole string

// LagoonCreateUserCmdUserEmail holds value of 'userEmail' option
var LagoonCreateUserCmdUserEmail string

// LagoonCreateUserCmdUserPassword holds value of 'userPassword' option
var LagoonCreateUserCmdUserPassword string

// LagoonCreateUserCmdBody holds contents of request body to be sent
var LagoonCreateUserCmdBody string

func init() {
	LagoonCreateUserCmd.Flags().StringVar(&LagoonCreateUserCmdRole, "role", "", TRAPI(""))

	LagoonCreateUserCmd.Flags().StringVar(&LagoonCreateUserCmdUserEmail, "user-email", "", TRAPI(""))

	LagoonCreateUserCmd.Flags().StringVar(&LagoonCreateUserCmdUserPassword, "user-password", "", TRAPI(""))

	LagoonCreateUserCmd.Flags().StringVar(&LagoonCreateUserCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LagoonCmd.AddCommand(LagoonCreateUserCmd)
}

// LagoonCreateUserCmd defines 'create-user' subcommand
var LagoonCreateUserCmd = &cobra.Command{
	Use:   "create-user",
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

		param, err := collectLagoonCreateUserCmdParams(ac)
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

func collectLagoonCreateUserCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonCreateUserCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLagoonCreateUserCmd("/lagoon/users"),
		query:       buildQueryForLagoonCreateUserCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonCreateUserCmd(path string) string {

	return path
}

func buildQueryForLagoonCreateUserCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLagoonCreateUserCmd() (string, error) {
	var result map[string]interface{}

	if LagoonCreateUserCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonCreateUserCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonCreateUserCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LagoonCreateUserCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonCreateUserCmdBody)
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

	if LagoonCreateUserCmdRole != "" {
		result["role"] = LagoonCreateUserCmdRole
	}

	if LagoonCreateUserCmdUserEmail != "" {
		result["userEmail"] = LagoonCreateUserCmdUserEmail
	}

	if LagoonCreateUserCmdUserPassword != "" {
		result["userPassword"] = LagoonCreateUserCmdUserPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
