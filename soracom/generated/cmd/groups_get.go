// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GroupsGetCmdGroupId holds value of 'group_id' option
var GroupsGetCmdGroupId string

func init() {
	GroupsGetCmd.Flags().StringVar(&GroupsGetCmdGroupId, "group-id", "", TRAPI("Target group ID."))
	GroupsCmd.AddCommand(GroupsGetCmd)
}

// GroupsGetCmd defines 'get' subcommand
var GroupsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/groups/{group_id}:get:summary"),
	Long:  TRAPI(`/groups/{group_id}:get:description`),
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

		param, err := collectGroupsGetCmdParams(ac)
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

func collectGroupsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("group_id", "group-id", "path", parsedBody, GroupsGetCmdGroupId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForGroupsGetCmd("/groups/{group_id}"),
		query:  buildQueryForGroupsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGroupsGetCmd(path string) string {

	escapedGroupId := url.PathEscape(GroupsGetCmdGroupId)

	path = strReplace(path, "{"+"group_id"+"}", escapedGroupId, -1)

	return path
}

func buildQueryForGroupsGetCmd() url.Values {
	result := url.Values{}

	return result
}
