// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamLicensePacksListCmdOutputJSONL indicates to output with jsonl format
var SoraCamLicensePacksListCmdOutputJSONL bool

func InitSoraCamLicensePacksListCmd() {
	SoraCamLicensePacksListCmd.Flags().BoolVar(&SoraCamLicensePacksListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	SoraCamLicensePacksListCmd.RunE = SoraCamLicensePacksListCmdRunE

	SoraCamLicensePacksCmd.AddCommand(SoraCamLicensePacksListCmd)
}

// SoraCamLicensePacksListCmd defines 'list' subcommand
var SoraCamLicensePacksListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/sora_cam/license_packs:get:summary"),
	Long:  TRAPI(`/sora_cam/license_packs:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "listSoraCamLicensePacks"),
}

func SoraCamLicensePacksListCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamLicensePacksListCmdParams(ac)
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
		if SoraCamLicensePacksListCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectSoraCamLicensePacksListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamLicensePacksListCmd("/sora_cam/license_packs"),
		query:  buildQueryForSoraCamLicensePacksListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamLicensePacksListCmd(path string) string {

	return path
}

func buildQueryForSoraCamLicensePacksListCmd() url.Values {
	result := url.Values{}

	return result
}
