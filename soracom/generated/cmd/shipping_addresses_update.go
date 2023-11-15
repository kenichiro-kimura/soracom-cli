// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// ShippingAddressesUpdateCmdAddressLine1 holds value of 'addressLine1' option
var ShippingAddressesUpdateCmdAddressLine1 string

// ShippingAddressesUpdateCmdAddressLine2 holds value of 'addressLine2' option
var ShippingAddressesUpdateCmdAddressLine2 string

// ShippingAddressesUpdateCmdBuilding holds value of 'building' option
var ShippingAddressesUpdateCmdBuilding string

// ShippingAddressesUpdateCmdCity holds value of 'city' option
var ShippingAddressesUpdateCmdCity string

// ShippingAddressesUpdateCmdCompanyName holds value of 'companyName' option
var ShippingAddressesUpdateCmdCompanyName string

// ShippingAddressesUpdateCmdCountryCode holds value of 'countryCode' option
var ShippingAddressesUpdateCmdCountryCode string

// ShippingAddressesUpdateCmdDepartment holds value of 'department' option
var ShippingAddressesUpdateCmdDepartment string

// ShippingAddressesUpdateCmdEmail holds value of 'email' option
var ShippingAddressesUpdateCmdEmail string

// ShippingAddressesUpdateCmdFullName holds value of 'fullName' option
var ShippingAddressesUpdateCmdFullName string

// ShippingAddressesUpdateCmdOperatorId holds value of 'operator_id' option
var ShippingAddressesUpdateCmdOperatorId string

// ShippingAddressesUpdateCmdPhoneNumber holds value of 'phoneNumber' option
var ShippingAddressesUpdateCmdPhoneNumber string

// ShippingAddressesUpdateCmdShippingAddressId holds value of 'shipping_address_id' option
var ShippingAddressesUpdateCmdShippingAddressId string

// ShippingAddressesUpdateCmdState holds value of 'state' option
var ShippingAddressesUpdateCmdState string

// ShippingAddressesUpdateCmdZipCode holds value of 'zipCode' option
var ShippingAddressesUpdateCmdZipCode string

// ShippingAddressesUpdateCmdBody holds contents of request body to be sent
var ShippingAddressesUpdateCmdBody string

func InitShippingAddressesUpdateCmd() {
	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdAddressLine1, "address-line1", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdAddressLine2, "address-line2", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdBuilding, "building", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdCity, "city", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdCompanyName, "company-name", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdCountryCode, "country-code", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdDepartment, "department", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdEmail, "email", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdFullName, "full-name", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdOperatorId, "operator-id", "", TRAPI("Operator ID."))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdPhoneNumber, "phone-number", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdShippingAddressId, "shipping-address-id", "", TRAPI("Shipping address ID."))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdState, "state", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdZipCode, "zip-code", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	ShippingAddressesUpdateCmd.RunE = ShippingAddressesUpdateCmdRunE

	ShippingAddressesCmd.AddCommand(ShippingAddressesUpdateCmd)
}

// ShippingAddressesUpdateCmd defines 'update' subcommand
var ShippingAddressesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/operators/{operator_id}/shipping_addresses/{shipping_address_id}:put:summary"),
	Long:  TRAPI(`/operators/{operator_id}/shipping_addresses/{shipping_address_id}:put:description`) + "\n\n" + createLinkToAPIReference("ShippingAddress", "updateShippingAddress"),
}

func ShippingAddressesUpdateCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectShippingAddressesUpdateCmdParams(ac)
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

func collectShippingAddressesUpdateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if ShippingAddressesUpdateCmdOperatorId == "" {
		ShippingAddressesUpdateCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	body, err = buildBodyForShippingAddressesUpdateCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("addressLine1", "address-line1", "body", parsedBody, ShippingAddressesUpdateCmdAddressLine1)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("city", "city", "body", parsedBody, ShippingAddressesUpdateCmdCity)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("shipping_address_id", "shipping-address-id", "path", parsedBody, ShippingAddressesUpdateCmdShippingAddressId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("state", "state", "body", parsedBody, ShippingAddressesUpdateCmdState)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("zipCode", "zip-code", "body", parsedBody, ShippingAddressesUpdateCmdZipCode)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForShippingAddressesUpdateCmd("/operators/{operator_id}/shipping_addresses/{shipping_address_id}"),
		query:       buildQueryForShippingAddressesUpdateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForShippingAddressesUpdateCmd(path string) string {

	escapedOperatorId := url.PathEscape(ShippingAddressesUpdateCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedShippingAddressId := url.PathEscape(ShippingAddressesUpdateCmdShippingAddressId)

	path = strReplace(path, "{"+"shipping_address_id"+"}", escapedShippingAddressId, -1)

	return path
}

func buildQueryForShippingAddressesUpdateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForShippingAddressesUpdateCmd() (string, error) {
	var result map[string]interface{}

	if ShippingAddressesUpdateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(ShippingAddressesUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(ShippingAddressesUpdateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if ShippingAddressesUpdateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(ShippingAddressesUpdateCmdBody)
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

	if ShippingAddressesUpdateCmdAddressLine1 != "" {
		result["addressLine1"] = ShippingAddressesUpdateCmdAddressLine1
	}

	if ShippingAddressesUpdateCmdAddressLine2 != "" {
		result["addressLine2"] = ShippingAddressesUpdateCmdAddressLine2
	}

	if ShippingAddressesUpdateCmdBuilding != "" {
		result["building"] = ShippingAddressesUpdateCmdBuilding
	}

	if ShippingAddressesUpdateCmdCity != "" {
		result["city"] = ShippingAddressesUpdateCmdCity
	}

	if ShippingAddressesUpdateCmdCompanyName != "" {
		result["companyName"] = ShippingAddressesUpdateCmdCompanyName
	}

	if ShippingAddressesUpdateCmdCountryCode != "" {
		result["countryCode"] = ShippingAddressesUpdateCmdCountryCode
	}

	if ShippingAddressesUpdateCmdDepartment != "" {
		result["department"] = ShippingAddressesUpdateCmdDepartment
	}

	if ShippingAddressesUpdateCmdEmail != "" {
		result["email"] = ShippingAddressesUpdateCmdEmail
	}

	if ShippingAddressesUpdateCmdFullName != "" {
		result["fullName"] = ShippingAddressesUpdateCmdFullName
	}

	if ShippingAddressesUpdateCmdPhoneNumber != "" {
		result["phoneNumber"] = ShippingAddressesUpdateCmdPhoneNumber
	}

	if ShippingAddressesUpdateCmdState != "" {
		result["state"] = ShippingAddressesUpdateCmdState
	}

	if ShippingAddressesUpdateCmdZipCode != "" {
		result["zipCode"] = ShippingAddressesUpdateCmdZipCode
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
