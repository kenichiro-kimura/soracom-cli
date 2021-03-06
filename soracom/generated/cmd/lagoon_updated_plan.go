package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonUpdatedPlanCmdPlan holds value of 'plan' option
var LagoonUpdatedPlanCmdPlan string

// LagoonUpdatedPlanCmdBody holds contents of request body to be sent
var LagoonUpdatedPlanCmdBody string

func init() {
	LagoonUpdatedPlanCmd.Flags().StringVar(&LagoonUpdatedPlanCmdPlan, "plan", "", TRAPI(""))

	LagoonUpdatedPlanCmd.Flags().StringVar(&LagoonUpdatedPlanCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LagoonCmd.AddCommand(LagoonUpdatedPlanCmd)
}

// LagoonUpdatedPlanCmd defines 'updated-plan' subcommand
var LagoonUpdatedPlanCmd = &cobra.Command{
	Use:   "updated-plan",
	Short: TRAPI("/lagoon/plan:put:summary"),
	Long:  TRAPI(`/lagoon/plan:put:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectLagoonUpdatedPlanCmdParams(ac)
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

func collectLagoonUpdatedPlanCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonUpdatedPlanCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUpdatedPlanCmd("/lagoon/plan"),
		query:       buildQueryForLagoonUpdatedPlanCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonUpdatedPlanCmd(path string) string {

	return path
}

func buildQueryForLagoonUpdatedPlanCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLagoonUpdatedPlanCmd() (string, error) {
	var result map[string]interface{}

	if LagoonUpdatedPlanCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonUpdatedPlanCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUpdatedPlanCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LagoonUpdatedPlanCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonUpdatedPlanCmdBody)
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

	if LagoonUpdatedPlanCmdPlan != "" {
		result["plan"] = LagoonUpdatedPlanCmdPlan
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
