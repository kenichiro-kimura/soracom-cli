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

// SimsPutCountryMappingEntriesCmdIccid holds value of 'iccid' option
var SimsPutCountryMappingEntriesCmdIccid string

// SimsPutCountryMappingEntriesCmdSimId holds value of 'sim_id' option
var SimsPutCountryMappingEntriesCmdSimId string

// SimsPutCountryMappingEntriesCmdBody holds contents of request body to be sent
var SimsPutCountryMappingEntriesCmdBody string

func init() {
	SimsPutCountryMappingEntriesCmd.Flags().StringVar(&SimsPutCountryMappingEntriesCmdIccid, "iccid", "", TRAPI("Iccid of the target profile"))

	SimsPutCountryMappingEntriesCmd.Flags().StringVar(&SimsPutCountryMappingEntriesCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsPutCountryMappingEntriesCmd.Flags().StringVar(&SimsPutCountryMappingEntriesCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SimsCmd.AddCommand(SimsPutCountryMappingEntriesCmd)
}

// SimsPutCountryMappingEntriesCmd defines 'put-country-mapping-entries' subcommand
var SimsPutCountryMappingEntriesCmd = &cobra.Command{
	Use:   "put-country-mapping-entries",
	Short: TRAPI("/sims/{sim_id}/profiles/{iccid}/subscription_containers/country_mapping:put:summary"),
	Long:  TRAPI(`/sims/{sim_id}/profiles/{iccid}/subscription_containers/country_mapping:put:description`) + "\n\n" + createLinkToAPIReference("Sim", "putSubscriptionContainerCountryMappingEntries"),
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

		param, err := collectSimsPutCountryMappingEntriesCmdParams(ac)
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

func collectSimsPutCountryMappingEntriesCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSimsPutCountryMappingEntriesCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("iccid", "iccid", "path", parsedBody, SimsPutCountryMappingEntriesCmdIccid)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsPutCountryMappingEntriesCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForSimsPutCountryMappingEntriesCmd("/sims/{sim_id}/profiles/{iccid}/subscription_containers/country_mapping"),
		query:       buildQueryForSimsPutCountryMappingEntriesCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsPutCountryMappingEntriesCmd(path string) string {

	escapedIccid := url.PathEscape(SimsPutCountryMappingEntriesCmdIccid)

	path = strReplace(path, "{"+"iccid"+"}", escapedIccid, -1)

	escapedSimId := url.PathEscape(SimsPutCountryMappingEntriesCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsPutCountryMappingEntriesCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimsPutCountryMappingEntriesCmd() (string, error) {
	var result map[string]interface{}

	if SimsPutCountryMappingEntriesCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimsPutCountryMappingEntriesCmdBody, "@") {
			fname := strings.TrimPrefix(SimsPutCountryMappingEntriesCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SimsPutCountryMappingEntriesCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SimsPutCountryMappingEntriesCmdBody)
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
