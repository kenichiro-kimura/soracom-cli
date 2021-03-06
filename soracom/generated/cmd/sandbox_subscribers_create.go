package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SandboxSubscribersCreateCmdSubscription holds value of 'subscription' option
var SandboxSubscribersCreateCmdSubscription string

// SandboxSubscribersCreateCmdBody holds contents of request body to be sent
var SandboxSubscribersCreateCmdBody string

func init() {
	SandboxSubscribersCreateCmd.Flags().StringVar(&SandboxSubscribersCreateCmdSubscription, "subscription", "", TRAPI(""))

	SandboxSubscribersCreateCmd.Flags().StringVar(&SandboxSubscribersCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SandboxSubscribersCmd.AddCommand(SandboxSubscribersCreateCmd)
}

// SandboxSubscribersCreateCmd defines 'create' subcommand
var SandboxSubscribersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/sandbox/subscribers/create:post:summary"),
	Long:  TRAPI(`/sandbox/subscribers/create:post:description`),
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

		param, err := collectSandboxSubscribersCreateCmdParams(ac)
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

func collectSandboxSubscribersCreateCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSandboxSubscribersCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSandboxSubscribersCreateCmd("/sandbox/subscribers/create"),
		query:       buildQueryForSandboxSubscribersCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSandboxSubscribersCreateCmd(path string) string {

	return path
}

func buildQueryForSandboxSubscribersCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSandboxSubscribersCreateCmd() (string, error) {
	var result map[string]interface{}

	if SandboxSubscribersCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxSubscribersCreateCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxSubscribersCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SandboxSubscribersCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SandboxSubscribersCreateCmdBody)
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

	if SandboxSubscribersCreateCmdSubscription != "" {
		result["subscription"] = SandboxSubscribersCreateCmdSubscription
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
