// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// CellLocationsGetCmdCid holds value of 'cid' option
var CellLocationsGetCmdCid string

// CellLocationsGetCmdEci holds value of 'eci' option
var CellLocationsGetCmdEci string

// CellLocationsGetCmdEcid holds value of 'ecid' option
var CellLocationsGetCmdEcid string

// CellLocationsGetCmdLac holds value of 'lac' option
var CellLocationsGetCmdLac string

// CellLocationsGetCmdMcc holds value of 'mcc' option
var CellLocationsGetCmdMcc string

// CellLocationsGetCmdMnc holds value of 'mnc' option
var CellLocationsGetCmdMnc string

// CellLocationsGetCmdTac holds value of 'tac' option
var CellLocationsGetCmdTac string

func InitCellLocationsGetCmd() {
	CellLocationsGetCmd.Flags().StringVar(&CellLocationsGetCmdCid, "cid", "", TRAPI("CID - Cell ID (for 3G)"))

	CellLocationsGetCmd.Flags().StringVar(&CellLocationsGetCmdEci, "eci", "", TRAPI("ECID - Enhanced Cell ID (for 4G) - specify either 'ecid' or 'eci'"))

	CellLocationsGetCmd.Flags().StringVar(&CellLocationsGetCmdEcid, "ecid", "", TRAPI("ECID - Enhanced Cell ID (for 4G) - specify either 'ecid' or 'eci'"))

	CellLocationsGetCmd.Flags().StringVar(&CellLocationsGetCmdLac, "lac", "", TRAPI("LAC - Location Area Code (for 3G)"))

	CellLocationsGetCmd.Flags().StringVar(&CellLocationsGetCmdMcc, "mcc", "", TRAPI("MCC - Mobile Country Code"))

	CellLocationsGetCmd.Flags().StringVar(&CellLocationsGetCmdMnc, "mnc", "", TRAPI("MNC - Mobile Network Code"))

	CellLocationsGetCmd.Flags().StringVar(&CellLocationsGetCmdTac, "tac", "", TRAPI("TAC - Tracking Area Code (for 4G)"))

	CellLocationsGetCmd.RunE = CellLocationsGetCmdRunE

	CellLocationsCmd.AddCommand(CellLocationsGetCmd)
}

// CellLocationsGetCmd defines 'get' subcommand
var CellLocationsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/cell_locations:get:summary"),
	Long:  TRAPI(`/cell_locations:get:description`) + "\n\n" + createLinkToAPIReference("CellLocation", "getCellLocation"),
}

func CellLocationsGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectCellLocationsGetCmdParams(ac)
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
}

func collectCellLocationsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error

	err = checkIfRequiredStringParameterIsSupplied("mcc", "mcc", "query", parsedBody, CellLocationsGetCmdMcc)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("mnc", "mnc", "query", parsedBody, CellLocationsGetCmdMnc)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForCellLocationsGetCmd("/cell_locations"),
		query:  buildQueryForCellLocationsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForCellLocationsGetCmd(path string) string {

	return path
}

func buildQueryForCellLocationsGetCmd() url.Values {
	result := url.Values{}

	if CellLocationsGetCmdCid != "" {
		result.Add("cid", CellLocationsGetCmdCid)
	}

	if CellLocationsGetCmdEci != "" {
		result.Add("eci", CellLocationsGetCmdEci)
	}

	if CellLocationsGetCmdEcid != "" {
		result.Add("ecid", CellLocationsGetCmdEcid)
	}

	if CellLocationsGetCmdLac != "" {
		result.Add("lac", CellLocationsGetCmdLac)
	}

	if CellLocationsGetCmdMcc != "" {
		result.Add("mcc", CellLocationsGetCmdMcc)
	}

	if CellLocationsGetCmdMnc != "" {
		result.Add("mnc", CellLocationsGetCmdMnc)
	}

	if CellLocationsGetCmdTac != "" {
		result.Add("tac", CellLocationsGetCmdTac)
	}

	return result
}
