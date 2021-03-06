package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// QuerySubscribersCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var QuerySubscribersCmdLastEvaluatedKey string

// QuerySubscribersCmdSearchType holds value of 'search_type' option
var QuerySubscribersCmdSearchType string

// QuerySubscribersCmdGroup holds multiple values of 'group' option
var QuerySubscribersCmdGroup []string

// QuerySubscribersCmdIccid holds multiple values of 'iccid' option
var QuerySubscribersCmdIccid []string

// QuerySubscribersCmdImsi holds multiple values of 'imsi' option
var QuerySubscribersCmdImsi []string

// QuerySubscribersCmdMsisdn holds multiple values of 'msisdn' option
var QuerySubscribersCmdMsisdn []string

// QuerySubscribersCmdName holds multiple values of 'name' option
var QuerySubscribersCmdName []string

// QuerySubscribersCmdSerialNumber holds multiple values of 'serial_number' option
var QuerySubscribersCmdSerialNumber []string

// QuerySubscribersCmdTag holds multiple values of 'tag' option
var QuerySubscribersCmdTag []string

// QuerySubscribersCmdLimit holds value of 'limit' option
var QuerySubscribersCmdLimit int64

func init() {
	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The IMSI of the last subscriber retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next subscriber onward."))

	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdSearchType, "search-type", "", TRAPI("Type of the search ('AND searching' or 'OR searching')"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdGroup, "group", []string{}, TRAPI("Group name to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdIccid, "iccid", []string{}, TRAPI("ICCID to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdImsi, "imsi", []string{}, TRAPI("IMSI to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdMsisdn, "msisdn", []string{}, TRAPI("MSISDN to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdName, "name", []string{}, TRAPI("Name to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdSerialNumber, "serial-number", []string{}, TRAPI("Serial number to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdTag, "tag", []string{}, TRAPI("String of tag values to search"))

	QuerySubscribersCmd.Flags().Int64Var(&QuerySubscribersCmdLimit, "limit", 0, TRAPI("The maximum number of item to retrieve"))

	QueryCmd.AddCommand(QuerySubscribersCmd)
}

// QuerySubscribersCmd defines 'subscribers' subcommand
var QuerySubscribersCmd = &cobra.Command{
	Use:   "subscribers",
	Short: TRAPI("/query/subscribers:get:summary"),
	Long:  TRAPI(`/query/subscribers:get:description`),
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

		param, err := collectQuerySubscribersCmdParams(ac)
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

func collectQuerySubscribersCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForQuerySubscribersCmd("/query/subscribers"),
		query:  buildQueryForQuerySubscribersCmd(),
	}, nil
}

func buildPathForQuerySubscribersCmd(path string) string {

	return path
}

func buildQueryForQuerySubscribersCmd() string {
	result := []string{}

	if QuerySubscribersCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", QuerySubscribersCmdLastEvaluatedKey))
	}

	if QuerySubscribersCmdSearchType != "" {
		result = append(result, sprintf("%s=%s", "search_type", QuerySubscribersCmdSearchType))
	}

	for _, s := range QuerySubscribersCmdGroup {
		if s != "" {
			result = append(result, sprintf("%s=%s", "group", s))
		}
	}

	for _, s := range QuerySubscribersCmdIccid {
		if s != "" {
			result = append(result, sprintf("%s=%s", "iccid", s))
		}
	}

	for _, s := range QuerySubscribersCmdImsi {
		if s != "" {
			result = append(result, sprintf("%s=%s", "imsi", s))
		}
	}

	for _, s := range QuerySubscribersCmdMsisdn {
		if s != "" {
			result = append(result, sprintf("%s=%s", "msisdn", s))
		}
	}

	for _, s := range QuerySubscribersCmdName {
		if s != "" {
			result = append(result, sprintf("%s=%s", "name", s))
		}
	}

	for _, s := range QuerySubscribersCmdSerialNumber {
		if s != "" {
			result = append(result, sprintf("%s=%s", "serial_number", s))
		}
	}

	for _, s := range QuerySubscribersCmdTag {
		if s != "" {
			result = append(result, sprintf("%s=%s", "tag", s))
		}
	}

	if QuerySubscribersCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", QuerySubscribersCmdLimit))
	}

	return strings.Join(result, "&")
}
