// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SubscribersActivateCmdImsi holds value of 'imsi' option
var SubscribersActivateCmdImsi string

func init() {
	SubscribersActivateCmd.Flags().StringVar(&SubscribersActivateCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))
	SubscribersCmd.AddCommand(SubscribersActivateCmd)
}

// SubscribersActivateCmd defines 'activate' subcommand
var SubscribersActivateCmd = &cobra.Command{
	Use:   "activate",
	Short: TRAPI("/subscribers/{imsi}/activate:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/activate:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

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
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectSubscribersActivateCmdParams(ac)
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

func collectSubscribersActivateCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, SubscribersActivateCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersActivateCmd("/subscribers/{imsi}/activate"),
		query:  buildQueryForSubscribersActivateCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersActivateCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersActivateCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersActivateCmd() url.Values {
	result := url.Values{}

	return result
}
