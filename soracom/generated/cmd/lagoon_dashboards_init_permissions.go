package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonDashboardsInitPermissionsCmdDashboardId holds value of 'dashboard_id' option
var LagoonDashboardsInitPermissionsCmdDashboardId int64

func init() {
	LagoonDashboardsInitPermissionsCmd.Flags().Int64Var(&LagoonDashboardsInitPermissionsCmdDashboardId, "dashboard-id", 0, TRAPI("dashboard_id"))

	LagoonDashboardsCmd.AddCommand(LagoonDashboardsInitPermissionsCmd)
}

// LagoonDashboardsInitPermissionsCmd defines 'init-permissions' subcommand
var LagoonDashboardsInitPermissionsCmd = &cobra.Command{
	Use:   "init-permissions",
	Short: TRAPI("/lagoon/dashboards/{dashboard_id}/permissions/init:post:summary"),
	Long:  TRAPI(`/lagoon/dashboards/{dashboard_id}/permissions/init:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectLagoonDashboardsInitPermissionsCmdParams(ac)
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

func collectLagoonDashboardsInitPermissionsCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLagoonDashboardsInitPermissionsCmd("/lagoon/dashboards/{dashboard_id}/permissions/init"),
		query:  buildQueryForLagoonDashboardsInitPermissionsCmd(),
	}, nil
}

func buildPathForLagoonDashboardsInitPermissionsCmd(path string) string {

	path = strings.Replace(path, "{"+"dashboard_id"+"}", sprintf("%d", LagoonDashboardsInitPermissionsCmdDashboardId), -1)

	return path
}

func buildQueryForLagoonDashboardsInitPermissionsCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
