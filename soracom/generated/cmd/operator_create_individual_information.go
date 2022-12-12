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

// OperatorCreateIndividualInformationCmdAddressLine1 holds value of 'addressLine1' option
var OperatorCreateIndividualInformationCmdAddressLine1 string

// OperatorCreateIndividualInformationCmdAddressLine2 holds value of 'addressLine2' option
var OperatorCreateIndividualInformationCmdAddressLine2 string

// OperatorCreateIndividualInformationCmdBuilding holds value of 'building' option
var OperatorCreateIndividualInformationCmdBuilding string

// OperatorCreateIndividualInformationCmdCity holds value of 'city' option
var OperatorCreateIndividualInformationCmdCity string

// OperatorCreateIndividualInformationCmdCountryCode holds value of 'countryCode' option
var OperatorCreateIndividualInformationCmdCountryCode string

// OperatorCreateIndividualInformationCmdFullName holds value of 'fullName' option
var OperatorCreateIndividualInformationCmdFullName string

// OperatorCreateIndividualInformationCmdOperatorId holds value of 'operator_id' option
var OperatorCreateIndividualInformationCmdOperatorId string

// OperatorCreateIndividualInformationCmdPhoneNumber holds value of 'phoneNumber' option
var OperatorCreateIndividualInformationCmdPhoneNumber string

// OperatorCreateIndividualInformationCmdState holds value of 'state' option
var OperatorCreateIndividualInformationCmdState string

// OperatorCreateIndividualInformationCmdZipCode holds value of 'zipCode' option
var OperatorCreateIndividualInformationCmdZipCode string

// OperatorCreateIndividualInformationCmdBody holds contents of request body to be sent
var OperatorCreateIndividualInformationCmdBody string

func init() {
	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdAddressLine1, "address-line1", "", TRAPI(""))

	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdAddressLine2, "address-line2", "", TRAPI(""))

	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdBuilding, "building", "", TRAPI(""))

	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdCity, "city", "", TRAPI(""))

	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdCountryCode, "country-code", "", TRAPI(""))

	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdFullName, "full-name", "", TRAPI(""))

	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdPhoneNumber, "phone-number", "", TRAPI(""))

	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdState, "state", "", TRAPI(""))

	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdZipCode, "zip-code", "", TRAPI(""))

	OperatorCreateIndividualInformationCmd.Flags().StringVar(&OperatorCreateIndividualInformationCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	OperatorCmd.AddCommand(OperatorCreateIndividualInformationCmd)
}

// OperatorCreateIndividualInformationCmd defines 'create-individual-information' subcommand
var OperatorCreateIndividualInformationCmd = &cobra.Command{
	Use:   "create-individual-information",
	Short: TRAPI("/operators/{operator_id}/individual_information:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/individual_information:post:description`) + "\n\n" + createLinkToAPIReference("Operator", "createIndividualInformation"),
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

		param, err := collectOperatorCreateIndividualInformationCmdParams(ac)
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

func collectOperatorCreateIndividualInformationCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if OperatorCreateIndividualInformationCmdOperatorId == "" {
		OperatorCreateIndividualInformationCmdOperatorId = ac.OperatorID
	}

	body, err = buildBodyForOperatorCreateIndividualInformationCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("countryCode", "country-code", "body", parsedBody, OperatorCreateIndividualInformationCmdCountryCode)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("fullName", "full-name", "body", parsedBody, OperatorCreateIndividualInformationCmdFullName)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("phoneNumber", "phone-number", "body", parsedBody, OperatorCreateIndividualInformationCmdPhoneNumber)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("zipCode", "zip-code", "body", parsedBody, OperatorCreateIndividualInformationCmdZipCode)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorCreateIndividualInformationCmd("/operators/{operator_id}/individual_information"),
		query:       buildQueryForOperatorCreateIndividualInformationCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorCreateIndividualInformationCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorCreateIndividualInformationCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorCreateIndividualInformationCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForOperatorCreateIndividualInformationCmd() (string, error) {
	var result map[string]interface{}

	if OperatorCreateIndividualInformationCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorCreateIndividualInformationCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorCreateIndividualInformationCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OperatorCreateIndividualInformationCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorCreateIndividualInformationCmdBody)
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

	if OperatorCreateIndividualInformationCmdAddressLine1 != "" {
		result["addressLine1"] = OperatorCreateIndividualInformationCmdAddressLine1
	}

	if OperatorCreateIndividualInformationCmdAddressLine2 != "" {
		result["addressLine2"] = OperatorCreateIndividualInformationCmdAddressLine2
	}

	if OperatorCreateIndividualInformationCmdBuilding != "" {
		result["building"] = OperatorCreateIndividualInformationCmdBuilding
	}

	if OperatorCreateIndividualInformationCmdCity != "" {
		result["city"] = OperatorCreateIndividualInformationCmdCity
	}

	if OperatorCreateIndividualInformationCmdCountryCode != "" {
		result["countryCode"] = OperatorCreateIndividualInformationCmdCountryCode
	}

	if OperatorCreateIndividualInformationCmdFullName != "" {
		result["fullName"] = OperatorCreateIndividualInformationCmdFullName
	}

	if OperatorCreateIndividualInformationCmdPhoneNumber != "" {
		result["phoneNumber"] = OperatorCreateIndividualInformationCmdPhoneNumber
	}

	if OperatorCreateIndividualInformationCmdState != "" {
		result["state"] = OperatorCreateIndividualInformationCmdState
	}

	if OperatorCreateIndividualInformationCmdZipCode != "" {
		result["zipCode"] = OperatorCreateIndividualInformationCmdZipCode
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
