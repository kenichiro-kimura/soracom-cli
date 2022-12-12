// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// RolesListUsersCmdOperatorId holds value of 'operator_id' option
var RolesListUsersCmdOperatorId string

// RolesListUsersCmdRoleId holds value of 'role_id' option
var RolesListUsersCmdRoleId string

// RolesListUsersCmdOutputJSONL indicates to output with jsonl format
var RolesListUsersCmdOutputJSONL bool

func init() {
	RolesListUsersCmd.Flags().StringVar(&RolesListUsersCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	RolesListUsersCmd.Flags().StringVar(&RolesListUsersCmdRoleId, "role-id", "", TRAPI("role_id"))

	RolesListUsersCmd.Flags().BoolVar(&RolesListUsersCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	RolesCmd.AddCommand(RolesListUsersCmd)
}

// RolesListUsersCmd defines 'list-users' subcommand
var RolesListUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: TRAPI("/operators/{operator_id}/roles/{role_id}/users:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles/{role_id}/users:get:description`) + "\n\n" + createLinkToAPIReference("Role", "listRoleAttachedUsers"),
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

		param, err := collectRolesListUsersCmdParams(ac)
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
			if RolesListUsersCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectRolesListUsersCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if RolesListUsersCmdOperatorId == "" {
		RolesListUsersCmdOperatorId = ac.OperatorID
	}

	err = checkIfRequiredStringParameterIsSupplied("role_id", "role-id", "path", parsedBody, RolesListUsersCmdRoleId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForRolesListUsersCmd("/operators/{operator_id}/roles/{role_id}/users"),
		query:  buildQueryForRolesListUsersCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForRolesListUsersCmd(path string) string {

	escapedOperatorId := url.PathEscape(RolesListUsersCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedRoleId := url.PathEscape(RolesListUsersCmdRoleId)

	path = strReplace(path, "{"+"role_id"+"}", escapedRoleId, -1)

	return path
}

func buildQueryForRolesListUsersCmd() url.Values {
	result := url.Values{}

	return result
}
