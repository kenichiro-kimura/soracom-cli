package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgTerminateCmdVpgId holds value of 'vpg_id' option
var VpgTerminateCmdVpgId string

func init() {
	VpgTerminateCmd.Flags().StringVar(&VpgTerminateCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgCmd.AddCommand(VpgTerminateCmd)
}

// VpgTerminateCmd defines 'terminate' subcommand
var VpgTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/terminate:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/terminate:post:description`),
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

		param, err := collectVpgTerminateCmdParams(ac)
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

func collectVpgTerminateCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgTerminateCmd("/virtual_private_gateways/{vpg_id}/terminate"),
		query:  buildQueryForVpgTerminateCmd(),
	}, nil
}

func buildPathForVpgTerminateCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgTerminateCmdVpgId, -1)

	return path
}

func buildQueryForVpgTerminateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
