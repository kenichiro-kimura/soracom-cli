package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DataDeleteEntryCmdResourceId holds value of 'resource_id' option
var DataDeleteEntryCmdResourceId string

// DataDeleteEntryCmdResourceType holds value of 'resource_type' option
var DataDeleteEntryCmdResourceType string

// DataDeleteEntryCmdTime holds value of 'time' option
var DataDeleteEntryCmdTime int64

func init() {
	DataDeleteEntryCmd.Flags().StringVar(&DataDeleteEntryCmdResourceId, "resource-id", "", TRAPI("ID of data source resource"))

	DataDeleteEntryCmd.Flags().StringVar(&DataDeleteEntryCmdResourceType, "resource-type", "", TRAPI("Type of data source resource"))

	DataDeleteEntryCmd.Flags().Int64Var(&DataDeleteEntryCmdTime, "time", 0, TRAPI("Timestamp of the target data entry to delete (unixtime in milliseconds)."))

	DataCmd.AddCommand(DataDeleteEntryCmd)
}

// DataDeleteEntryCmd defines 'delete-entry' subcommand
var DataDeleteEntryCmd = &cobra.Command{
	Use:   "delete-entry",
	Short: TRAPI("/data/{resource_type}/{resource_id}/{time}:delete:summary"),
	Long:  TRAPI(`/data/{resource_type}/{resource_id}/{time}:delete:description`),
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

		param, err := collectDataDeleteEntryCmdParams(ac)
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

func collectDataDeleteEntryCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForDataDeleteEntryCmd("/data/{resource_type}/{resource_id}/{time}"),
		query:  buildQueryForDataDeleteEntryCmd(),
	}, nil
}

func buildPathForDataDeleteEntryCmd(path string) string {

	path = strings.Replace(path, "{"+"resource_id"+"}", DataDeleteEntryCmdResourceId, -1)

	path = strings.Replace(path, "{"+"resource_type"+"}", DataDeleteEntryCmdResourceType, -1)

	return path
}

func buildQueryForDataDeleteEntryCmd() string {
	result := []string{}

	if DataDeleteEntryCmdTime != 0 {
		result = append(result, sprintf("%s=%d", "time", DataDeleteEntryCmdTime))
	}

	return strings.Join(result, "&")
}
