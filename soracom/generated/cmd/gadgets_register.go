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

// GadgetsRegisterCmdProductId holds value of 'product_id' option
var GadgetsRegisterCmdProductId string

// GadgetsRegisterCmdSerialNumber holds value of 'serial_number' option
var GadgetsRegisterCmdSerialNumber string

// GadgetsRegisterCmdBody holds contents of request body to be sent
var GadgetsRegisterCmdBody string

func InitGadgetsRegisterCmd() {
	GadgetsRegisterCmd.Flags().StringVar(&GadgetsRegisterCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsRegisterCmd.Flags().StringVar(&GadgetsRegisterCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))

	GadgetsRegisterCmd.Flags().StringVar(&GadgetsRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	GadgetsRegisterCmd.RunE = GadgetsRegisterCmdRunE

	GadgetsCmd.AddCommand(GadgetsRegisterCmd)
}

// GadgetsRegisterCmd defines 'register' subcommand
var GadgetsRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/register:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/register:post:description`) + "\n\n" + createLinkToAPIReference("Gadget", "registerGadget"),
}

func GadgetsRegisterCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectGadgetsRegisterCmdParams(ac)
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

func collectGadgetsRegisterCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForGadgetsRegisterCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("product_id", "product-id", "path", parsedBody, GadgetsRegisterCmdProductId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("serial_number", "serial-number", "path", parsedBody, GadgetsRegisterCmdSerialNumber)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForGadgetsRegisterCmd("/gadgets/{product_id}/{serial_number}/register"),
		query:       buildQueryForGadgetsRegisterCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGadgetsRegisterCmd(path string) string {

	escapedProductId := url.PathEscape(GadgetsRegisterCmdProductId)

	path = strReplace(path, "{"+"product_id"+"}", escapedProductId, -1)

	escapedSerialNumber := url.PathEscape(GadgetsRegisterCmdSerialNumber)

	path = strReplace(path, "{"+"serial_number"+"}", escapedSerialNumber, -1)

	return path
}

func buildQueryForGadgetsRegisterCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForGadgetsRegisterCmd() (string, error) {
	var result map[string]interface{}

	if GadgetsRegisterCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(GadgetsRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(GadgetsRegisterCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if GadgetsRegisterCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(GadgetsRegisterCmdBody)
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
