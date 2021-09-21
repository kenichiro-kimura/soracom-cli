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

// GroupsPutConfigCmdGroupId holds value of 'group_id' option
var GroupsPutConfigCmdGroupId string

// GroupsPutConfigCmdNamespace holds value of 'namespace' option
var GroupsPutConfigCmdNamespace string

// GroupsPutConfigCmdBody holds contents of request body to be sent
var GroupsPutConfigCmdBody string

func init() {
	GroupsPutConfigCmd.Flags().StringVar(&GroupsPutConfigCmdGroupId, "group-id", "", TRAPI("Target group."))

	GroupsPutConfigCmd.Flags().StringVar(&GroupsPutConfigCmdNamespace, "namespace", "", TRAPI("Target configuration."))

	GroupsPutConfigCmd.Flags().StringVar(&GroupsPutConfigCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	GroupsCmd.AddCommand(GroupsPutConfigCmd)
}

// GroupsPutConfigCmd defines 'put-config' subcommand
var GroupsPutConfigCmd = &cobra.Command{
	Use:   "put-config",
	Short: TRAPI("/groups/{group_id}/configuration/{namespace}:put:summary"),
	Long:  TRAPI(`/groups/{group_id}/configuration/{namespace}:put:description`),
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

		param, err := collectGroupsPutConfigCmdParams(ac)
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

func collectGroupsPutConfigCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForGroupsPutConfigCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("group_id", "group-id", "path", parsedBody, GroupsPutConfigCmdGroupId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("namespace", "namespace", "path", parsedBody, GroupsPutConfigCmdNamespace)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForGroupsPutConfigCmd("/groups/{group_id}/configuration/{namespace}"),
		query:       buildQueryForGroupsPutConfigCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGroupsPutConfigCmd(path string) string {

	escapedGroupId := url.PathEscape(GroupsPutConfigCmdGroupId)

	path = strReplace(path, "{"+"group_id"+"}", escapedGroupId, -1)

	escapedNamespace := url.PathEscape(GroupsPutConfigCmdNamespace)

	path = strReplace(path, "{"+"namespace"+"}", escapedNamespace, -1)

	return path
}

func buildQueryForGroupsPutConfigCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForGroupsPutConfigCmd() (string, error) {
	var b []byte
	var err error

	if GroupsPutConfigCmdBody != "" {
		if strings.HasPrefix(GroupsPutConfigCmdBody, "@") {
			fname := strings.TrimPrefix(GroupsPutConfigCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if GroupsPutConfigCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(GroupsPutConfigCmdBody)
		}

		if err != nil {
			return "", err
		}
	}

	if b == nil {
		b = []byte{}
	}

	return string(b), nil
}
