package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsCreateCmdCreatedTime holds value of 'createdTime' option
var LoraNetworkSetsCreateCmdCreatedTime string

// LoraNetworkSetsCreateCmdLastModifiedTime holds value of 'lastModifiedTime' option
var LoraNetworkSetsCreateCmdLastModifiedTime string

// LoraNetworkSetsCreateCmdNetworkSetId holds value of 'networkSetId' option
var LoraNetworkSetsCreateCmdNetworkSetId string

// LoraNetworkSetsCreateCmdOperatorId holds value of 'operatorId' option
var LoraNetworkSetsCreateCmdOperatorId string

// LoraNetworkSetsCreateCmdBody holds contents of request body to be sent
var LoraNetworkSetsCreateCmdBody string

func init() {
	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdCreatedTime, "created-time", "", TRAPI(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdLastModifiedTime, "last-modified-time", "", TRAPI(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdNetworkSetId, "network-set-id", "", TRAPI(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdOperatorId, "operator-id", "", TRAPI(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsCreateCmd)
}

// LoraNetworkSetsCreateCmd defines 'create' subcommand
var LoraNetworkSetsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/lora_network_sets:post:summary"),
	Long:  TRAPI(`/lora_network_sets:post:description`),
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

		param, err := collectLoraNetworkSetsCreateCmdParams(ac)
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

func collectLoraNetworkSetsCreateCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLoraNetworkSetsCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraNetworkSetsCreateCmd("/lora_network_sets"),
		query:       buildQueryForLoraNetworkSetsCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraNetworkSetsCreateCmd(path string) string {

	return path
}

func buildQueryForLoraNetworkSetsCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraNetworkSetsCreateCmd() (string, error) {
	var result map[string]interface{}

	if LoraNetworkSetsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LoraNetworkSetsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(LoraNetworkSetsCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LoraNetworkSetsCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraNetworkSetsCreateCmdBody)
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

	if LoraNetworkSetsCreateCmdCreatedTime != "" {
		result["createdTime"] = LoraNetworkSetsCreateCmdCreatedTime
	}

	if LoraNetworkSetsCreateCmdLastModifiedTime != "" {
		result["lastModifiedTime"] = LoraNetworkSetsCreateCmdLastModifiedTime
	}

	if LoraNetworkSetsCreateCmdNetworkSetId != "" {
		result["networkSetId"] = LoraNetworkSetsCreateCmdNetworkSetId
	}

	if LoraNetworkSetsCreateCmdOperatorId != "" {
		result["operatorId"] = LoraNetworkSetsCreateCmdOperatorId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
