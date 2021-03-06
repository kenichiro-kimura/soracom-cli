package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GadgetsSetGroupCmdGroupId holds value of 'groupId' option
var GadgetsSetGroupCmdGroupId string

// GadgetsSetGroupCmdOperatorId holds value of 'operatorId' option
var GadgetsSetGroupCmdOperatorId string

// GadgetsSetGroupCmdProductId holds value of 'product_id' option
var GadgetsSetGroupCmdProductId string

// GadgetsSetGroupCmdSerialNumber holds value of 'serial_number' option
var GadgetsSetGroupCmdSerialNumber string

// GadgetsSetGroupCmdCreatedTime holds value of 'createdTime' option
var GadgetsSetGroupCmdCreatedTime int64

// GadgetsSetGroupCmdLastModifiedTime holds value of 'lastModifiedTime' option
var GadgetsSetGroupCmdLastModifiedTime int64

// GadgetsSetGroupCmdBody holds contents of request body to be sent
var GadgetsSetGroupCmdBody string

func init() {
	GadgetsSetGroupCmd.Flags().StringVar(&GadgetsSetGroupCmdGroupId, "group-id", "", TRAPI(""))

	GadgetsSetGroupCmd.Flags().StringVar(&GadgetsSetGroupCmdOperatorId, "operator-id", "", TRAPI(""))

	GadgetsSetGroupCmd.Flags().StringVar(&GadgetsSetGroupCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsSetGroupCmd.Flags().StringVar(&GadgetsSetGroupCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))

	GadgetsSetGroupCmd.Flags().Int64Var(&GadgetsSetGroupCmdCreatedTime, "created-time", 0, TRAPI(""))

	GadgetsSetGroupCmd.Flags().Int64Var(&GadgetsSetGroupCmdLastModifiedTime, "last-modified-time", 0, TRAPI(""))

	GadgetsSetGroupCmd.Flags().StringVar(&GadgetsSetGroupCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	GadgetsCmd.AddCommand(GadgetsSetGroupCmd)
}

// GadgetsSetGroupCmd defines 'set-group' subcommand
var GadgetsSetGroupCmd = &cobra.Command{
	Use:   "set-group",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/set_group:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/set_group:post:description`),
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

		param, err := collectGadgetsSetGroupCmdParams(ac)
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

func collectGadgetsSetGroupCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForGadgetsSetGroupCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForGadgetsSetGroupCmd("/gadgets/{product_id}/{serial_number}/set_group"),
		query:       buildQueryForGadgetsSetGroupCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForGadgetsSetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"product_id"+"}", GadgetsSetGroupCmdProductId, -1)

	path = strings.Replace(path, "{"+"serial_number"+"}", GadgetsSetGroupCmdSerialNumber, -1)

	return path
}

func buildQueryForGadgetsSetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForGadgetsSetGroupCmd() (string, error) {
	var result map[string]interface{}

	if GadgetsSetGroupCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(GadgetsSetGroupCmdBody, "@") {
			fname := strings.TrimPrefix(GadgetsSetGroupCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if GadgetsSetGroupCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(GadgetsSetGroupCmdBody)
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

	if GadgetsSetGroupCmdGroupId != "" {
		result["groupId"] = GadgetsSetGroupCmdGroupId
	}

	if GadgetsSetGroupCmdOperatorId != "" {
		result["operatorId"] = GadgetsSetGroupCmdOperatorId
	}

	if GadgetsSetGroupCmdCreatedTime != 0 {
		result["createdTime"] = GadgetsSetGroupCmdCreatedTime
	}

	if GadgetsSetGroupCmdLastModifiedTime != 0 {
		result["lastModifiedTime"] = GadgetsSetGroupCmdLastModifiedTime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
