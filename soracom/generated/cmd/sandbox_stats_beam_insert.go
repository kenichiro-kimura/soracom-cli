package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SandboxStatsBeamInsertCmdImsi holds value of 'imsi' option
var SandboxStatsBeamInsertCmdImsi string

// SandboxStatsBeamInsertCmdUnixtime holds value of 'unixtime' option
var SandboxStatsBeamInsertCmdUnixtime int64

// SandboxStatsBeamInsertCmdBody holds contents of request body to be sent
var SandboxStatsBeamInsertCmdBody string

func init() {
	SandboxStatsBeamInsertCmd.Flags().StringVar(&SandboxStatsBeamInsertCmdImsi, "imsi", "", TRAPI("IMSI"))

	SandboxStatsBeamInsertCmd.Flags().Int64Var(&SandboxStatsBeamInsertCmdUnixtime, "unixtime", 0, TRAPI(""))

	SandboxStatsBeamInsertCmd.Flags().StringVar(&SandboxStatsBeamInsertCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SandboxStatsBeamCmd.AddCommand(SandboxStatsBeamInsertCmd)
}

// SandboxStatsBeamInsertCmd defines 'insert' subcommand
var SandboxStatsBeamInsertCmd = &cobra.Command{
	Use:   "insert",
	Short: TRAPI("/sandbox/stats/beam/subscribers/{imsi}:post:summary"),
	Long:  TRAPI(`/sandbox/stats/beam/subscribers/{imsi}:post:description`),
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

		param, err := collectSandboxStatsBeamInsertCmdParams(ac)
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

func collectSandboxStatsBeamInsertCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSandboxStatsBeamInsertCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSandboxStatsBeamInsertCmd("/sandbox/stats/beam/subscribers/{imsi}"),
		query:       buildQueryForSandboxStatsBeamInsertCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSandboxStatsBeamInsertCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SandboxStatsBeamInsertCmdImsi, -1)

	return path
}

func buildQueryForSandboxStatsBeamInsertCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSandboxStatsBeamInsertCmd() (string, error) {
	var result map[string]interface{}

	if SandboxStatsBeamInsertCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxStatsBeamInsertCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxStatsBeamInsertCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SandboxStatsBeamInsertCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SandboxStatsBeamInsertCmdBody)
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

	if SandboxStatsBeamInsertCmdUnixtime != 0 {
		result["unixtime"] = SandboxStatsBeamInsertCmdUnixtime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
