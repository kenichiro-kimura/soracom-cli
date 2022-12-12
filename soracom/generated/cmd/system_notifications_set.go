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

// SystemNotificationsSetCmdOperatorId holds value of 'operator_id' option
var SystemNotificationsSetCmdOperatorId string

// SystemNotificationsSetCmdPassword holds value of 'password' option
var SystemNotificationsSetCmdPassword string

// SystemNotificationsSetCmdType holds value of 'type' option
var SystemNotificationsSetCmdType string

// SystemNotificationsSetCmdEmailIdList holds multiple values of 'emailIdList' option
var SystemNotificationsSetCmdEmailIdList []string

// SystemNotificationsSetCmdBody holds contents of request body to be sent
var SystemNotificationsSetCmdBody string

func init() {
	SystemNotificationsSetCmd.Flags().StringVar(&SystemNotificationsSetCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	SystemNotificationsSetCmd.Flags().StringVar(&SystemNotificationsSetCmdPassword, "password", "", TRAPI("Password of the operator. This is necessary when type is primary."))

	SystemNotificationsSetCmd.Flags().StringVar(&SystemNotificationsSetCmdType, "type", "", TRAPI("system notification type"))

	SystemNotificationsSetCmd.Flags().StringSliceVar(&SystemNotificationsSetCmdEmailIdList, "email-id-list", []string{}, TRAPI(""))

	SystemNotificationsSetCmd.Flags().StringVar(&SystemNotificationsSetCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SystemNotificationsCmd.AddCommand(SystemNotificationsSetCmd)
}

// SystemNotificationsSetCmd defines 'set' subcommand
var SystemNotificationsSetCmd = &cobra.Command{
	Use:   "set",
	Short: TRAPI("/operators/{operator_id}/system_notifications/{type}:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/system_notifications/{type}:post:description`) + "\n\n" + createLinkToAPIReference("SystemNotification", "setSystemNotification"),
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

		param, err := collectSystemNotificationsSetCmdParams(ac)
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

func collectSystemNotificationsSetCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if SystemNotificationsSetCmdOperatorId == "" {
		SystemNotificationsSetCmdOperatorId = ac.OperatorID
	}

	body, err = buildBodyForSystemNotificationsSetCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("type", "type", "path", parsedBody, SystemNotificationsSetCmdType)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringSliceParameterIsSupplied("emailIdList", "email-id-list", "body", parsedBody, SystemNotificationsSetCmdEmailIdList)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSystemNotificationsSetCmd("/operators/{operator_id}/system_notifications/{type}"),
		query:       buildQueryForSystemNotificationsSetCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSystemNotificationsSetCmd(path string) string {

	escapedOperatorId := url.PathEscape(SystemNotificationsSetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedType := url.PathEscape(SystemNotificationsSetCmdType)

	path = strReplace(path, "{"+"type"+"}", escapedType, -1)

	return path
}

func buildQueryForSystemNotificationsSetCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSystemNotificationsSetCmd() (string, error) {
	var result map[string]interface{}

	if SystemNotificationsSetCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SystemNotificationsSetCmdBody, "@") {
			fname := strings.TrimPrefix(SystemNotificationsSetCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SystemNotificationsSetCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SystemNotificationsSetCmdBody)
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

	if SystemNotificationsSetCmdPassword != "" {
		result["password"] = SystemNotificationsSetCmdPassword
	}

	if len(SystemNotificationsSetCmdEmailIdList) != 0 {
		result["emailIdList"] = SystemNotificationsSetCmdEmailIdList
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
