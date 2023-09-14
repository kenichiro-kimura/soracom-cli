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

// SoraletsUploadCmdContentType holds value of 'content-type' option
var SoraletsUploadCmdContentType string

// SoraletsUploadCmdSoraletId holds value of 'soralet_id' option
var SoraletsUploadCmdSoraletId string

// SoraletsUploadCmdBody holds contents of request body to be sent
var SoraletsUploadCmdBody string

func InitSoraletsUploadCmd() {
	SoraletsUploadCmd.Flags().StringVar(&SoraletsUploadCmdContentType, "content-type", "", TRAPI("Content type of the file to upload"))

	SoraletsUploadCmd.Flags().StringVar(&SoraletsUploadCmdSoraletId, "soralet-id", "", TRAPI("The identifier of Soralet."))

	SoraletsUploadCmd.Flags().StringVar(&SoraletsUploadCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SoraletsUploadCmd.RunE = SoraletsUploadCmdRunE

	SoraletsCmd.AddCommand(SoraletsUploadCmd)
}

// SoraletsUploadCmd defines 'upload' subcommand
var SoraletsUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: TRAPI("/soralets/{soralet_id}/versions:post:summary"),
	Long:  TRAPI(`/soralets/{soralet_id}/versions:post:description`) + "\n\n" + createLinkToAPIReference("Soralet", "uploadSoraletCode"),
}

func SoraletsUploadCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraletsUploadCmdParams(ac)
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

func collectSoraletsUploadCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraletsUploadCmd()
	if err != nil {
		return nil, err
	}
	contentType := SoraletsUploadCmdContentType

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("soralet_id", "soralet-id", "path", parsedBody, SoraletsUploadCmdSoraletId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSoraletsUploadCmd("/soralets/{soralet_id}/versions"),
		query:       buildQueryForSoraletsUploadCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraletsUploadCmd(path string) string {

	escapedSoraletId := url.PathEscape(SoraletsUploadCmdSoraletId)

	path = strReplace(path, "{"+"soralet_id"+"}", escapedSoraletId, -1)

	return path
}

func buildQueryForSoraletsUploadCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraletsUploadCmd() (string, error) {
	var b []byte
	var err error

	if SoraletsUploadCmdBody != "" {
		if strings.HasPrefix(SoraletsUploadCmdBody, "@") {
			fname := strings.TrimPrefix(SoraletsUploadCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SoraletsUploadCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraletsUploadCmdBody)
		}

		if err != nil {
			return "", err
		}
	}

	if b == nil {
		b = []byte{}
	}

	return string(b), nil
}
