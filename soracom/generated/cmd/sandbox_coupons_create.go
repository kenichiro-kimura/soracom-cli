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

// SandboxCouponsCreateCmdApplicableBillItemName holds value of 'applicableBillItemName' option
var SandboxCouponsCreateCmdApplicableBillItemName string

// SandboxCouponsCreateCmdExpiryYearMonth holds value of 'expiryYearMonth' option
var SandboxCouponsCreateCmdExpiryYearMonth string

// SandboxCouponsCreateCmdAmount holds value of 'amount' option
var SandboxCouponsCreateCmdAmount int64

// SandboxCouponsCreateCmdBody holds contents of request body to be sent
var SandboxCouponsCreateCmdBody string

func init() {
	SandboxCouponsCreateCmd.Flags().StringVar(&SandboxCouponsCreateCmdApplicableBillItemName, "applicable-bill-item-name", "", TRAPI(""))

	SandboxCouponsCreateCmd.Flags().StringVar(&SandboxCouponsCreateCmdExpiryYearMonth, "expiry-year-month", "", TRAPI(""))

	SandboxCouponsCreateCmd.Flags().Int64Var(&SandboxCouponsCreateCmdAmount, "amount", 0, TRAPI(""))

	SandboxCouponsCreateCmd.Flags().StringVar(&SandboxCouponsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SandboxCouponsCmd.AddCommand(SandboxCouponsCreateCmd)
}

// SandboxCouponsCreateCmd defines 'create' subcommand
var SandboxCouponsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/sandbox/coupons/create:post:summary"),
	Long:  TRAPI(`/sandbox/coupons/create:post:description`),
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

		param, err := collectSandboxCouponsCreateCmdParams(ac)
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

func collectSandboxCouponsCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSandboxCouponsCreateCmd()
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

	return &apiParams{
		method:      "POST",
		path:        buildPathForSandboxCouponsCreateCmd("/sandbox/coupons/create"),
		query:       buildQueryForSandboxCouponsCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSandboxCouponsCreateCmd(path string) string {

	return path
}

func buildQueryForSandboxCouponsCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSandboxCouponsCreateCmd() (string, error) {
	var result map[string]interface{}

	if SandboxCouponsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxCouponsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxCouponsCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SandboxCouponsCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SandboxCouponsCreateCmdBody)
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

	if SandboxCouponsCreateCmdApplicableBillItemName != "" {
		result["applicableBillItemName"] = SandboxCouponsCreateCmdApplicableBillItemName
	}

	if SandboxCouponsCreateCmdExpiryYearMonth != "" {
		result["expiryYearMonth"] = SandboxCouponsCreateCmdExpiryYearMonth
	}

	if SandboxCouponsCreateCmdAmount != 0 {
		result["amount"] = SandboxCouponsCreateCmdAmount
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
