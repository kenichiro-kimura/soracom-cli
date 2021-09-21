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

// SubscribersSetImeiLockCmdImei holds value of 'imei' option
var SubscribersSetImeiLockCmdImei string

// SubscribersSetImeiLockCmdImsi holds value of 'imsi' option
var SubscribersSetImeiLockCmdImsi string

// SubscribersSetImeiLockCmdBody holds contents of request body to be sent
var SubscribersSetImeiLockCmdBody string

func init() {
	SubscribersSetImeiLockCmd.Flags().StringVar(&SubscribersSetImeiLockCmdImei, "imei", "", TRAPI(""))

	SubscribersSetImeiLockCmd.Flags().StringVar(&SubscribersSetImeiLockCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersSetImeiLockCmd.Flags().StringVar(&SubscribersSetImeiLockCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SubscribersCmd.AddCommand(SubscribersSetImeiLockCmd)
}

// SubscribersSetImeiLockCmd defines 'set-imei-lock' subcommand
var SubscribersSetImeiLockCmd = &cobra.Command{
	Use:   "set-imei-lock",
	Short: TRAPI("/subscribers/{imsi}/set_imei_lock:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/set_imei_lock:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
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

		param, err := collectSubscribersSetImeiLockCmdParams(ac)
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

func collectSubscribersSetImeiLockCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSubscribersSetImeiLockCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, SubscribersSetImeiLockCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersSetImeiLockCmd("/subscribers/{imsi}/set_imei_lock"),
		query:       buildQueryForSubscribersSetImeiLockCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersSetImeiLockCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersSetImeiLockCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersSetImeiLockCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSubscribersSetImeiLockCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersSetImeiLockCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersSetImeiLockCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersSetImeiLockCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SubscribersSetImeiLockCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersSetImeiLockCmdBody)
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

	if SubscribersSetImeiLockCmdImei != "" {
		result["imei"] = SubscribersSetImeiLockCmdImei
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
