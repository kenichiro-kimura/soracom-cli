package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GadgetsListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var GadgetsListCmdLastEvaluatedKey string

// GadgetsListCmdProductId holds value of 'product_id' option
var GadgetsListCmdProductId string

// GadgetsListCmdTagName holds value of 'tag_name' option
var GadgetsListCmdTagName string

// GadgetsListCmdTagValue holds value of 'tag_value' option
var GadgetsListCmdTagValue string

// GadgetsListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var GadgetsListCmdTagValueMatchMode string

// GadgetsListCmdLimit holds value of 'limit' option
var GadgetsListCmdLimit int64

func init() {
	GadgetsListCmd.Flags().StringVar(&GadgetsListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The ID ({product_id}/{serial_number}) of the last gadget retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next device onward."))

	GadgetsListCmd.Flags().StringVar(&GadgetsListCmdProductId, "product-id", "", TRAPI("Product ID for filtering the search."))

	GadgetsListCmd.Flags().StringVar(&GadgetsListCmdTagName, "tag-name", "", TRAPI("Tag name for filtering the search (exact match)."))

	GadgetsListCmd.Flags().StringVar(&GadgetsListCmdTagValue, "tag-value", "", TRAPI("Tag search string for filtering the search. Required when `tag_name` has been specified."))

	GadgetsListCmd.Flags().StringVar(&GadgetsListCmdTagValueMatchMode, "tag-value-match-mode", "", TRAPI("Tag match mode."))

	GadgetsListCmd.Flags().Int64Var(&GadgetsListCmdLimit, "limit", 0, TRAPI("Maximum number of gadgets to retrieve."))

	GadgetsCmd.AddCommand(GadgetsListCmd)
}

// GadgetsListCmd defines 'list' subcommand
var GadgetsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/gadgets:get:summary"),
	Long:  TRAPI(`/gadgets:get:description`),
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

		param, err := collectGadgetsListCmdParams(ac)
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

func collectGadgetsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForGadgetsListCmd("/gadgets"),
		query:  buildQueryForGadgetsListCmd(),
	}, nil
}

func buildPathForGadgetsListCmd(path string) string {

	return path
}

func buildQueryForGadgetsListCmd() string {
	result := []string{}

	if GadgetsListCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", GadgetsListCmdLastEvaluatedKey))
	}

	if GadgetsListCmdProductId != "" {
		result = append(result, sprintf("%s=%s", "product_id", GadgetsListCmdProductId))
	}

	if GadgetsListCmdTagName != "" {
		result = append(result, sprintf("%s=%s", "tag_name", GadgetsListCmdTagName))
	}

	if GadgetsListCmdTagValue != "" {
		result = append(result, sprintf("%s=%s", "tag_value", GadgetsListCmdTagValue))
	}

	if GadgetsListCmdTagValueMatchMode != "" {
		result = append(result, sprintf("%s=%s", "tag_value_match_mode", GadgetsListCmdTagValueMatchMode))
	}

	if GadgetsListCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", GadgetsListCmdLimit))
	}

	return strings.Join(result, "&")
}
