package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

var StatsBeamExportCmdOperatorId string

var StatsBeamExportCmdPeriod string

var StatsBeamExportCmdFrom int64

var StatsBeamExportCmdTo int64

var StatsBeamExportCmdBody string

func init() {
	StatsBeamExportCmd.Flags().StringVar(&StatsBeamExportCmdOperatorId, "operator-id", "", TR("stats.export_beam_stats.post.parameters.operator_id.description"))

	StatsBeamExportCmd.Flags().StringVar(&StatsBeamExportCmdPeriod, "period", "", TR(""))

	StatsBeamExportCmd.Flags().Int64Var(&StatsBeamExportCmdFrom, "from", 0, TR(""))

	StatsBeamExportCmd.Flags().Int64Var(&StatsBeamExportCmdTo, "to", 0, TR(""))

	StatsBeamExportCmd.Flags().StringVar(&StatsBeamExportCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	StatsBeamCmd.AddCommand(StatsBeamExportCmd)
}

var StatsBeamExportCmd = &cobra.Command{
	Use:   "export",
	Short: TR("stats.export_beam_stats.post.summary"),
	Long:  TR(`stats.export_beam_stats.post.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		param, err := collectStatsBeamExportCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectStatsBeamExportCmdParams() (*apiParams, error) {

	body, err := buildBodyForStatsBeamExportCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForStatsBeamExportCmd("/stats/beam/operators/{operator_id}/export"),
		query:       buildQueryForStatsBeamExportCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForStatsBeamExportCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", StatsBeamExportCmdOperatorId, -1)

	return path
}

func buildQueryForStatsBeamExportCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForStatsBeamExportCmd() (string, error) {
	if StatsBeamExportCmdBody != "" {
		if strings.HasPrefix(StatsBeamExportCmdBody, "@") {
			fname := strings.TrimPrefix(StatsBeamExportCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if StatsBeamExportCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return StatsBeamExportCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if StatsBeamExportCmdPeriod != "" {
		result["period"] = StatsBeamExportCmdPeriod
	}

	if StatsBeamExportCmdFrom != 0 {
		result["from"] = StatsBeamExportCmdFrom
	}

	if StatsBeamExportCmdTo != 0 {
		result["to"] = StatsBeamExportCmdTo
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
