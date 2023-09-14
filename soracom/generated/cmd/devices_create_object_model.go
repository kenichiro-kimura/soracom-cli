// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// DevicesCreateObjectModelCmdCreatedTime holds value of 'createdTime' option
var DevicesCreateObjectModelCmdCreatedTime string

// DevicesCreateObjectModelCmdFormat holds value of 'format' option
var DevicesCreateObjectModelCmdFormat string

// DevicesCreateObjectModelCmdLastModifiedTime holds value of 'lastModifiedTime' option
var DevicesCreateObjectModelCmdLastModifiedTime string

// DevicesCreateObjectModelCmdObjectId holds value of 'objectId' option
var DevicesCreateObjectModelCmdObjectId string

// DevicesCreateObjectModelCmdObjectName holds value of 'objectName' option
var DevicesCreateObjectModelCmdObjectName string

// DevicesCreateObjectModelCmdOperatorId holds value of 'operatorId' option
var DevicesCreateObjectModelCmdOperatorId string

// DevicesCreateObjectModelCmdScope holds value of 'scope' option
var DevicesCreateObjectModelCmdScope string

// DevicesCreateObjectModelCmdBody holds contents of request body to be sent
var DevicesCreateObjectModelCmdBody string

func InitDevicesCreateObjectModelCmd() {
	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdCreatedTime, "created-time", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdFormat, "format", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdLastModifiedTime, "last-modified-time", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdObjectId, "object-id", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdObjectName, "object-name", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdOperatorId, "operator-id", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdScope, "scope", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	DevicesCreateObjectModelCmd.RunE = DevicesCreateObjectModelCmdRunE

	DevicesCmd.AddCommand(DevicesCreateObjectModelCmd)
}

// DevicesCreateObjectModelCmd defines 'create-object-model' subcommand
var DevicesCreateObjectModelCmd = &cobra.Command{
	Use:   "create-object-model",
	Short: TRAPI("/device_object_models:post:summary"),
	Long:  TRAPI(`/device_object_models:post:description`) + "\n\n" + createLinkToAPIReference("DeviceObjectModel", "createDeviceObjectModel"),
}

func DevicesCreateObjectModelCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectDevicesCreateObjectModelCmdParams(ac)
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
}

func collectDevicesCreateObjectModelCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForDevicesCreateObjectModelCmd()
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

	return &apiParams{
		method:      "POST",
		path:        buildPathForDevicesCreateObjectModelCmd("/device_object_models"),
		query:       buildQueryForDevicesCreateObjectModelCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesCreateObjectModelCmd(path string) string {

	return path
}

func buildQueryForDevicesCreateObjectModelCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForDevicesCreateObjectModelCmd() (string, error) {
	var result map[string]interface{}

	if DevicesCreateObjectModelCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(DevicesCreateObjectModelCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesCreateObjectModelCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if DevicesCreateObjectModelCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(DevicesCreateObjectModelCmdBody)
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

	if DevicesCreateObjectModelCmdCreatedTime != "" {
		result["createdTime"] = DevicesCreateObjectModelCmdCreatedTime
	}

	if DevicesCreateObjectModelCmdFormat != "" {
		result["format"] = DevicesCreateObjectModelCmdFormat
	}

	if DevicesCreateObjectModelCmdLastModifiedTime != "" {
		result["lastModifiedTime"] = DevicesCreateObjectModelCmdLastModifiedTime
	}

	if DevicesCreateObjectModelCmdObjectId != "" {
		result["objectId"] = DevicesCreateObjectModelCmdObjectId
	}

	if DevicesCreateObjectModelCmdObjectName != "" {
		result["objectName"] = DevicesCreateObjectModelCmdObjectName
	}

	if DevicesCreateObjectModelCmdOperatorId != "" {
		result["operatorId"] = DevicesCreateObjectModelCmdOperatorId
	}

	if DevicesCreateObjectModelCmdScope != "" {
		result["scope"] = DevicesCreateObjectModelCmdScope
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
