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

// OperatorAddContractCmdContractName holds value of 'contractName' option
var OperatorAddContractCmdContractName string

// OperatorAddContractCmdOperatorId holds value of 'operator_id' option
var OperatorAddContractCmdOperatorId string

// OperatorAddContractCmdBody holds contents of request body to be sent
var OperatorAddContractCmdBody string

func init() {
	OperatorAddContractCmd.Flags().StringVar(&OperatorAddContractCmdContractName, "contract-name", "", TRAPI(""))

	OperatorAddContractCmd.Flags().StringVar(&OperatorAddContractCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	OperatorAddContractCmd.Flags().StringVar(&OperatorAddContractCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	OperatorCmd.AddCommand(OperatorAddContractCmd)
}

// OperatorAddContractCmd defines 'add-contract' subcommand
var OperatorAddContractCmd = &cobra.Command{
	Use:   "add-contract",
	Short: TRAPI("/operators/{operator_id}/contracts:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/contracts:post:description`) + "\n\n" + createLinkToAPIReference("Operator", "addOperatorContract"),
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

		param, err := collectOperatorAddContractCmdParams(ac)
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

func collectOperatorAddContractCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if OperatorAddContractCmdOperatorId == "" {
		OperatorAddContractCmdOperatorId = ac.OperatorID
	}

	body, err = buildBodyForOperatorAddContractCmd()
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
		path:        buildPathForOperatorAddContractCmd("/operators/{operator_id}/contracts"),
		query:       buildQueryForOperatorAddContractCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorAddContractCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorAddContractCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorAddContractCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForOperatorAddContractCmd() (string, error) {
	var result map[string]interface{}

	if OperatorAddContractCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorAddContractCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorAddContractCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OperatorAddContractCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorAddContractCmdBody)
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

	if OperatorAddContractCmdContractName != "" {
		result["contractName"] = OperatorAddContractCmdContractName
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
