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

// SimsRegisterCmdGroupId holds value of 'groupId' option
var SimsRegisterCmdGroupId string

// SimsRegisterCmdRegistrationSecret holds value of 'registrationSecret' option
var SimsRegisterCmdRegistrationSecret string

// SimsRegisterCmdSimId holds value of 'sim_id' option
var SimsRegisterCmdSimId string

// SimsRegisterCmdBody holds contents of request body to be sent
var SimsRegisterCmdBody string

func init() {
	SimsRegisterCmd.Flags().StringVar(&SimsRegisterCmdGroupId, "group-id", "", TRAPI(""))

	SimsRegisterCmd.Flags().StringVar(&SimsRegisterCmdRegistrationSecret, "registration-secret", "", TRAPI(""))

	SimsRegisterCmd.Flags().StringVar(&SimsRegisterCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsRegisterCmd.Flags().StringVar(&SimsRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SimsCmd.AddCommand(SimsRegisterCmd)
}

// SimsRegisterCmd defines 'register' subcommand
var SimsRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/sims/{sim_id}/register:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/register:post:description`),
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

		param, err := collectSimsRegisterCmdParams(ac)
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

func collectSimsRegisterCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSimsRegisterCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("registrationSecret", "registration-secret", "body", parsedBody, SimsRegisterCmdRegistrationSecret)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsRegisterCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimsRegisterCmd("/sims/{sim_id}/register"),
		query:       buildQueryForSimsRegisterCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsRegisterCmd(path string) string {

	escapedSimId := url.PathEscape(SimsRegisterCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsRegisterCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimsRegisterCmd() (string, error) {
	var result map[string]interface{}

	if SimsRegisterCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimsRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(SimsRegisterCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SimsRegisterCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SimsRegisterCmdBody)
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

	if SimsRegisterCmdGroupId != "" {
		result["groupId"] = SimsRegisterCmdGroupId
	}

	if SimsRegisterCmdRegistrationSecret != "" {
		result["registrationSecret"] = SimsRegisterCmdRegistrationSecret
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
