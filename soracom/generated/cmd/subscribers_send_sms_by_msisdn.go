package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersSendSmsByMsisdnCmdMsisdn holds value of 'msisdn' option
var SubscribersSendSmsByMsisdnCmdMsisdn string

// SubscribersSendSmsByMsisdnCmdPayload holds value of 'payload' option
var SubscribersSendSmsByMsisdnCmdPayload string

// SubscribersSendSmsByMsisdnCmdEncodingType holds value of 'encodingType' option
var SubscribersSendSmsByMsisdnCmdEncodingType int64

// SubscribersSendSmsByMsisdnCmdBody holds contents of request body to be sent
var SubscribersSendSmsByMsisdnCmdBody string

func init() {
	SubscribersSendSmsByMsisdnCmd.Flags().StringVar(&SubscribersSendSmsByMsisdnCmdMsisdn, "msisdn", "", TRAPI("MSISDN of the target subscriber."))

	SubscribersSendSmsByMsisdnCmd.Flags().StringVar(&SubscribersSendSmsByMsisdnCmdPayload, "payload", "", TRAPI(""))

	SubscribersSendSmsByMsisdnCmd.Flags().Int64Var(&SubscribersSendSmsByMsisdnCmdEncodingType, "encoding-type", 0, TRAPI(""))

	SubscribersSendSmsByMsisdnCmd.Flags().StringVar(&SubscribersSendSmsByMsisdnCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SubscribersCmd.AddCommand(SubscribersSendSmsByMsisdnCmd)
}

// SubscribersSendSmsByMsisdnCmd defines 'send-sms-by-msisdn' subcommand
var SubscribersSendSmsByMsisdnCmd = &cobra.Command{
	Use:   "send-sms-by-msisdn",
	Short: TRAPI("/subscribers/msisdn/{msisdn}/send_sms:post:summary"),
	Long:  TRAPI(`/subscribers/msisdn/{msisdn}/send_sms:post:description`),
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

		param, err := collectSubscribersSendSmsByMsisdnCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectSubscribersSendSmsByMsisdnCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSubscribersSendSmsByMsisdnCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersSendSmsByMsisdnCmd("/subscribers/msisdn/{msisdn}/send_sms"),
		query:       buildQueryForSubscribersSendSmsByMsisdnCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSubscribersSendSmsByMsisdnCmd(path string) string {

	path = strings.Replace(path, "{"+"msisdn"+"}", SubscribersSendSmsByMsisdnCmdMsisdn, -1)

	return path
}

func buildQueryForSubscribersSendSmsByMsisdnCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSubscribersSendSmsByMsisdnCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersSendSmsByMsisdnCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersSendSmsByMsisdnCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersSendSmsByMsisdnCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SubscribersSendSmsByMsisdnCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersSendSmsByMsisdnCmdBody)
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

	if SubscribersSendSmsByMsisdnCmdPayload != "" {
		result["payload"] = SubscribersSendSmsByMsisdnCmdPayload
	}

	if SubscribersSendSmsByMsisdnCmdEncodingType != 0 {
		result["encodingType"] = SubscribersSendSmsByMsisdnCmdEncodingType
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
