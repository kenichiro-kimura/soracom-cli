package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonLicensePacksUpdateCmdBody holds contents of request body to be sent
var LagoonLicensePacksUpdateCmdBody string

func init() {

	LagoonLicensePacksUpdateCmd.Flags().StringVar(&LagoonLicensePacksUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LagoonLicensePacksCmd.AddCommand(LagoonLicensePacksUpdateCmd)
}

// LagoonLicensePacksUpdateCmd defines 'update' subcommand
var LagoonLicensePacksUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/lagoon/license_packs:put:summary"),
	Long:  TRAPI(`/lagoon/license_packs:put:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectLagoonLicensePacksUpdateCmdParams(ac)
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

func collectLagoonLicensePacksUpdateCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonLicensePacksUpdateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonLicensePacksUpdateCmd("/lagoon/license_packs"),
		query:       buildQueryForLagoonLicensePacksUpdateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonLicensePacksUpdateCmd(path string) string {

	return path
}

func buildQueryForLagoonLicensePacksUpdateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLagoonLicensePacksUpdateCmd() (string, error) {
	var result map[string]interface{}

	if LagoonLicensePacksUpdateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonLicensePacksUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonLicensePacksUpdateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LagoonLicensePacksUpdateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonLicensePacksUpdateCmdBody)
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

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
