// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SubscribersSetToStandbyCmdImsi holds value of 'imsi' option
var SubscribersSetToStandbyCmdImsi string

func InitSubscribersSetToStandbyCmd() {
	SubscribersSetToStandbyCmd.Flags().StringVar(&SubscribersSetToStandbyCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersSetToStandbyCmd.RunE = SubscribersSetToStandbyCmdRunE

	SubscribersCmd.AddCommand(SubscribersSetToStandbyCmd)
}

// SubscribersSetToStandbyCmd defines 'set-to-standby' subcommand
var SubscribersSetToStandbyCmd = &cobra.Command{
	Use:   "set-to-standby",
	Short: TRAPI("/subscribers/{imsi}/set_to_standby:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/set_to_standby:post:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "setSubscriberToStandby"),
}

func SubscribersSetToStandbyCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSubscribersSetToStandbyCmdParams(ac)
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
}

func collectSubscribersSetToStandbyCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, SubscribersSetToStandbyCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersSetToStandbyCmd("/subscribers/{imsi}/set_to_standby"),
		query:  buildQueryForSubscribersSetToStandbyCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersSetToStandbyCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersSetToStandbyCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersSetToStandbyCmd() url.Values {
	result := url.Values{}

	return result
}
