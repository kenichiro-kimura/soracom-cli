// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsStopPacketCaptureSessionCmdSessionId holds value of 'session_id' option
var SimsStopPacketCaptureSessionCmdSessionId string

// SimsStopPacketCaptureSessionCmdSimId holds value of 'sim_id' option
var SimsStopPacketCaptureSessionCmdSimId string

func init() {
	SimsStopPacketCaptureSessionCmd.Flags().StringVar(&SimsStopPacketCaptureSessionCmdSessionId, "session-id", "", TRAPI("Packet capture session ID"))

	SimsStopPacketCaptureSessionCmd.Flags().StringVar(&SimsStopPacketCaptureSessionCmdSimId, "sim-id", "", TRAPI("SIM ID"))
	SimsCmd.AddCommand(SimsStopPacketCaptureSessionCmd)
}

// SimsStopPacketCaptureSessionCmd defines 'stop-packet-capture-session' subcommand
var SimsStopPacketCaptureSessionCmd = &cobra.Command{
	Use:   "stop-packet-capture-session",
	Short: TRAPI("/sims/{sim_id}/packet_capture_sessions/{session_id}/stop:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/packet_capture_sessions/{session_id}/stop:post:description`),
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

		param, err := collectSimsStopPacketCaptureSessionCmdParams(ac)
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

func collectSimsStopPacketCaptureSessionCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("session_id", "session-id", "path", parsedBody, SimsStopPacketCaptureSessionCmdSessionId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsStopPacketCaptureSessionCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsStopPacketCaptureSessionCmd("/sims/{sim_id}/packet_capture_sessions/{session_id}/stop"),
		query:  buildQueryForSimsStopPacketCaptureSessionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsStopPacketCaptureSessionCmd(path string) string {

	escapedSessionId := url.PathEscape(SimsStopPacketCaptureSessionCmdSessionId)

	path = strReplace(path, "{"+"session_id"+"}", escapedSessionId, -1)

	escapedSimId := url.PathEscape(SimsStopPacketCaptureSessionCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsStopPacketCaptureSessionCmd() url.Values {
	result := url.Values{}

	return result
}
