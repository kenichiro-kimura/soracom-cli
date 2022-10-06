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

// SimsDownlinkPingCmdSimId holds value of 'sim_id' option
var SimsDownlinkPingCmdSimId string

// SimsDownlinkPingCmdNumberOfPingRequests holds value of 'numberOfPingRequests' option
var SimsDownlinkPingCmdNumberOfPingRequests int64

// SimsDownlinkPingCmdTimeoutSeconds holds value of 'timeoutSeconds' option
var SimsDownlinkPingCmdTimeoutSeconds int64

// SimsDownlinkPingCmdBody holds contents of request body to be sent
var SimsDownlinkPingCmdBody string

func init() {
	SimsDownlinkPingCmd.Flags().StringVar(&SimsDownlinkPingCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsDownlinkPingCmd.Flags().Int64Var(&SimsDownlinkPingCmdNumberOfPingRequests, "number-of-ping-requests", 1, TRAPI("the number of attempt ping"))

	SimsDownlinkPingCmd.Flags().Int64Var(&SimsDownlinkPingCmdTimeoutSeconds, "timeout-seconds", 1, TRAPI("the timeout duration of each ping"))

	SimsDownlinkPingCmd.Flags().StringVar(&SimsDownlinkPingCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SimsCmd.AddCommand(SimsDownlinkPingCmd)
}

// SimsDownlinkPingCmd defines 'downlink-ping' subcommand
var SimsDownlinkPingCmd = &cobra.Command{
	Use:   "downlink-ping",
	Short: TRAPI("/sims/{sim_id}/downlink/ping:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/downlink/ping:post:description`) + "\n\n" + createLinkToAPIReference("Sim", "sendDownlinkPing"),
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

		param, err := collectSimsDownlinkPingCmdParams(ac)
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

func collectSimsDownlinkPingCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSimsDownlinkPingCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsDownlinkPingCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimsDownlinkPingCmd("/sims/{sim_id}/downlink/ping"),
		query:       buildQueryForSimsDownlinkPingCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsDownlinkPingCmd(path string) string {

	escapedSimId := url.PathEscape(SimsDownlinkPingCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsDownlinkPingCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimsDownlinkPingCmd() (string, error) {
	var result map[string]interface{}

	if SimsDownlinkPingCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimsDownlinkPingCmdBody, "@") {
			fname := strings.TrimPrefix(SimsDownlinkPingCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SimsDownlinkPingCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SimsDownlinkPingCmdBody)
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

	if SimsDownlinkPingCmdNumberOfPingRequests != 1 {
		result["numberOfPingRequests"] = SimsDownlinkPingCmdNumberOfPingRequests
	}

	if SimsDownlinkPingCmdTimeoutSeconds != 1 {
		result["timeoutSeconds"] = SimsDownlinkPingCmdTimeoutSeconds
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
