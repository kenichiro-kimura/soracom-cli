package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GroupsListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var GroupsListCmdLastEvaluatedKey string

// GroupsListCmdTagName holds value of 'tag_name' option
var GroupsListCmdTagName string

// GroupsListCmdTagValue holds value of 'tag_value' option
var GroupsListCmdTagValue string

// GroupsListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var GroupsListCmdTagValueMatchMode string

// GroupsListCmdLimit holds value of 'limit' option
var GroupsListCmdLimit int64

func init() {
	GroupsListCmd.Flags().StringVar(&GroupsListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The last Group ID retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next group onward."))

	GroupsListCmd.Flags().StringVar(&GroupsListCmdTagName, "tag-name", "", TRAPI("Tag name of the group. Filters through all groups that exactly match the tag name. When tag_name is specified, tag_value is required."))

	GroupsListCmd.Flags().StringVar(&GroupsListCmdTagValue, "tag-value", "", TRAPI("Tag value of the groups."))

	GroupsListCmd.Flags().StringVar(&GroupsListCmdTagValueMatchMode, "tag-value-match-mode", "", TRAPI("Tag match mode."))

	GroupsListCmd.Flags().Int64Var(&GroupsListCmdLimit, "limit", 0, TRAPI("Maximum number of results per response page."))

	GroupsCmd.AddCommand(GroupsListCmd)
}

// GroupsListCmd defines 'list' subcommand
var GroupsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/groups:get:summary"),
	Long:  TRAPI(`/groups:get:description`),
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

		param, err := collectGroupsListCmdParams(ac)
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

func collectGroupsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForGroupsListCmd("/groups"),
		query:  buildQueryForGroupsListCmd(),
	}, nil
}

func buildPathForGroupsListCmd(path string) string {

	return path
}

func buildQueryForGroupsListCmd() string {
	result := []string{}

	if GroupsListCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", GroupsListCmdLastEvaluatedKey))
	}

	if GroupsListCmdTagName != "" {
		result = append(result, sprintf("%s=%s", "tag_name", GroupsListCmdTagName))
	}

	if GroupsListCmdTagValue != "" {
		result = append(result, sprintf("%s=%s", "tag_value", GroupsListCmdTagValue))
	}

	if GroupsListCmdTagValueMatchMode != "" {
		result = append(result, sprintf("%s=%s", "tag_value_match_mode", GroupsListCmdTagValueMatchMode))
	}

	if GroupsListCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", GroupsListCmdLimit))
	}

	return strings.Join(result, "&")
}
