package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesSetObjectModelScopeCmdModelId holds value of 'model_id' option
var DevicesSetObjectModelScopeCmdModelId string

// DevicesSetObjectModelScopeCmdScope holds value of 'scope' option
var DevicesSetObjectModelScopeCmdScope string

// DevicesSetObjectModelScopeCmdBody holds contents of request body to be sent
var DevicesSetObjectModelScopeCmdBody string

func init() {
	DevicesSetObjectModelScopeCmd.Flags().StringVar(&DevicesSetObjectModelScopeCmdModelId, "model-id", "", TRAPI("Target device object model ID"))

	DevicesSetObjectModelScopeCmd.Flags().StringVar(&DevicesSetObjectModelScopeCmdScope, "scope", "", TRAPI(""))

	DevicesSetObjectModelScopeCmd.Flags().StringVar(&DevicesSetObjectModelScopeCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	DevicesCmd.AddCommand(DevicesSetObjectModelScopeCmd)
}

// DevicesSetObjectModelScopeCmd defines 'set-object-model-scope' subcommand
var DevicesSetObjectModelScopeCmd = &cobra.Command{
	Use:   "set-object-model-scope",
	Short: TRAPI("/device_object_models/{model_id}/set_scope:post:summary"),
	Long:  TRAPI(`/device_object_models/{model_id}/set_scope:post:description`),
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

		param, err := collectDevicesSetObjectModelScopeCmdParams(ac)
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

func collectDevicesSetObjectModelScopeCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForDevicesSetObjectModelScopeCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForDevicesSetObjectModelScopeCmd("/device_object_models/{model_id}/set_scope"),
		query:       buildQueryForDevicesSetObjectModelScopeCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForDevicesSetObjectModelScopeCmd(path string) string {

	path = strings.Replace(path, "{"+"model_id"+"}", DevicesSetObjectModelScopeCmdModelId, -1)

	return path
}

func buildQueryForDevicesSetObjectModelScopeCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForDevicesSetObjectModelScopeCmd() (string, error) {
	var result map[string]interface{}

	if DevicesSetObjectModelScopeCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(DevicesSetObjectModelScopeCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesSetObjectModelScopeCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if DevicesSetObjectModelScopeCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(DevicesSetObjectModelScopeCmdBody)
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

	if DevicesSetObjectModelScopeCmdScope != "" {
		result["scope"] = DevicesSetObjectModelScopeCmdScope
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
