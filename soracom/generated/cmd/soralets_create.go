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

// SoraletsCreateCmdDescription holds value of 'description' option
var SoraletsCreateCmdDescription string

// SoraletsCreateCmdSoraletId holds value of 'soraletId' option
var SoraletsCreateCmdSoraletId string

// SoraletsCreateCmdBody holds contents of request body to be sent
var SoraletsCreateCmdBody string

func init() {
	SoraletsCreateCmd.Flags().StringVar(&SoraletsCreateCmdDescription, "description", "", TRAPI(""))

	SoraletsCreateCmd.Flags().StringVar(&SoraletsCreateCmdSoraletId, "soralet-id", "", TRAPI(""))

	SoraletsCreateCmd.Flags().StringVar(&SoraletsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SoraletsCmd.AddCommand(SoraletsCreateCmd)
}

// SoraletsCreateCmd defines 'create' subcommand
var SoraletsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/soralets:post:summary"),
	Long:  TRAPI(`/soralets:post:description`),
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

		param, err := collectSoraletsCreateCmdParams(ac)
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

func collectSoraletsCreateCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForSoraletsCreateCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if SoraletsCreateCmdSoraletId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "soralet-id")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSoraletsCreateCmd("/soralets"),
		query:       buildQueryForSoraletsCreateCmd(),
		contentType: contentType,
		body:        body,
	}, nil
}

func buildPathForSoraletsCreateCmd(path string) string {

	return path
}

func buildQueryForSoraletsCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraletsCreateCmd() (string, error) {
	var result map[string]interface{}

	if SoraletsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraletsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(SoraletsCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SoraletsCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraletsCreateCmdBody)
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

	if SoraletsCreateCmdDescription != "" {
		result["description"] = SoraletsCreateCmdDescription
	}

	if SoraletsCreateCmdSoraletId != "" {
		result["soraletId"] = SoraletsCreateCmdSoraletId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}