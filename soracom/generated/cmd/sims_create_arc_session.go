// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsCreateArcSessionCmdSimId holds value of 'sim_id' option
var SimsCreateArcSessionCmdSimId string

func init() {
	SimsCreateArcSessionCmd.Flags().StringVar(&SimsCreateArcSessionCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))
	SimsCmd.AddCommand(SimsCreateArcSessionCmd)
}

// SimsCreateArcSessionCmd defines 'create-arc-session' subcommand
var SimsCreateArcSessionCmd = &cobra.Command{
	Use:   "create-arc-session",
	Short: TRAPI("/sims/{sim_id}/sessions/arc:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/sessions/arc:post:description`),
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

		param, err := collectSimsCreateArcSessionCmdParams(ac)
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

func collectSimsCreateArcSessionCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsCreateArcSessionCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsCreateArcSessionCmd("/sims/{sim_id}/sessions/arc"),
		query:  buildQueryForSimsCreateArcSessionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsCreateArcSessionCmd(path string) string {

	escapedSimId := url.PathEscape(SimsCreateArcSessionCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsCreateArcSessionCmd() url.Values {
	result := url.Values{}

	return result
}
