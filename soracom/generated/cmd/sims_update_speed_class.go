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

// SimsUpdateSpeedClassCmdSimId holds value of 'sim_id' option
var SimsUpdateSpeedClassCmdSimId string

// SimsUpdateSpeedClassCmdSpeedClass holds value of 'speedClass' option
var SimsUpdateSpeedClassCmdSpeedClass string

// SimsUpdateSpeedClassCmdBody holds contents of request body to be sent
var SimsUpdateSpeedClassCmdBody string

func init() {
	SimsUpdateSpeedClassCmd.Flags().StringVar(&SimsUpdateSpeedClassCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsUpdateSpeedClassCmd.Flags().StringVar(&SimsUpdateSpeedClassCmdSpeedClass, "speed-class", "", TRAPI(""))

	SimsUpdateSpeedClassCmd.Flags().StringVar(&SimsUpdateSpeedClassCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SimsCmd.AddCommand(SimsUpdateSpeedClassCmd)
}

// SimsUpdateSpeedClassCmd defines 'update-speed-class' subcommand
var SimsUpdateSpeedClassCmd = &cobra.Command{
	Use:   "update-speed-class",
	Short: TRAPI("/sims/{sim_id}/update_speed_class:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/update_speed_class:post:description`),
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

		param, err := collectSimsUpdateSpeedClassCmdParams(ac)
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

func collectSimsUpdateSpeedClassCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSimsUpdateSpeedClassCmd()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &parsedBody)
	if err != nil {
		return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
	}
	contentType := "application/json"

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsUpdateSpeedClassCmdSimId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("speedClass", "speed-class", "body", parsedBody, SimsUpdateSpeedClassCmdSpeedClass)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimsUpdateSpeedClassCmd("/sims/{sim_id}/update_speed_class"),
		query:       buildQueryForSimsUpdateSpeedClassCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsUpdateSpeedClassCmd(path string) string {

	escapedSimId := url.PathEscape(SimsUpdateSpeedClassCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsUpdateSpeedClassCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimsUpdateSpeedClassCmd() (string, error) {
	var result map[string]interface{}

	if SimsUpdateSpeedClassCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimsUpdateSpeedClassCmdBody, "@") {
			fname := strings.TrimPrefix(SimsUpdateSpeedClassCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SimsUpdateSpeedClassCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SimsUpdateSpeedClassCmdBody)
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

	if SimsUpdateSpeedClassCmdSpeedClass != "" {
		result["speedClass"] = SimsUpdateSpeedClassCmdSpeedClass
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
