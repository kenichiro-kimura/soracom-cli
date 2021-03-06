package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersUnsetGroupCmdImsi holds value of 'imsi' option
var SubscribersUnsetGroupCmdImsi string

func init() {
	SubscribersUnsetGroupCmd.Flags().StringVar(&SubscribersUnsetGroupCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersCmd.AddCommand(SubscribersUnsetGroupCmd)
}

// SubscribersUnsetGroupCmd defines 'unset-group' subcommand
var SubscribersUnsetGroupCmd = &cobra.Command{
	Use:   "unset-group",
	Short: TRAPI("/subscribers/{imsi}/unset_group:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/unset_group:post:description`),
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

		param, err := collectSubscribersUnsetGroupCmdParams(ac)
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

func collectSubscribersUnsetGroupCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersUnsetGroupCmd("/subscribers/{imsi}/unset_group"),
		query:  buildQueryForSubscribersUnsetGroupCmd(),
	}, nil
}

func buildPathForSubscribersUnsetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersUnsetGroupCmdImsi, -1)

	return path
}

func buildQueryForSubscribersUnsetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
