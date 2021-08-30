// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsEnableTerminationCmdSimId holds value of 'sim_id' option
var SimsEnableTerminationCmdSimId string

func init() {
	SimsEnableTerminationCmd.Flags().StringVar(&SimsEnableTerminationCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))
	SimsCmd.AddCommand(SimsEnableTerminationCmd)
}

// SimsEnableTerminationCmd defines 'enable-termination' subcommand
var SimsEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/sims/{sim_id}/enable_termination:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/enable_termination:post:description`),
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

		param, err := collectSimsEnableTerminationCmdParams(ac)
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

func collectSimsEnableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsEnableTerminationCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsEnableTerminationCmd("/sims/{sim_id}/enable_termination"),
		query:  buildQueryForSimsEnableTerminationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsEnableTerminationCmd(path string) string {

	escapedSimId := url.PathEscape(SimsEnableTerminationCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsEnableTerminationCmd() url.Values {
	result := url.Values{}

	return result
}
