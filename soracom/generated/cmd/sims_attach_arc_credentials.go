// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/soracom/soracom-cli/generators/lib"

	"github.com/spf13/cobra"
)

// SimsAttachArcCredentialsCmdArcClientPeerPublicKey holds value of 'arcClientPeerPublicKey' option
var SimsAttachArcCredentialsCmdArcClientPeerPublicKey string

// SimsAttachArcCredentialsCmdSimId holds value of 'sim_id' option
var SimsAttachArcCredentialsCmdSimId string

// SimsAttachArcCredentialsCmdBody holds contents of request body to be sent
var SimsAttachArcCredentialsCmdBody string

func InitSimsAttachArcCredentialsCmd() {
	SimsAttachArcCredentialsCmd.Flags().StringVar(&SimsAttachArcCredentialsCmdArcClientPeerPublicKey, "arc-client-peer-public-key", "", TRAPI("if this parameter is missing, the sever generates keypair"))

	SimsAttachArcCredentialsCmd.Flags().StringVar(&SimsAttachArcCredentialsCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM"))

	SimsAttachArcCredentialsCmd.Flags().StringVar(&SimsAttachArcCredentialsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SimsAttachArcCredentialsCmd.RunE = SimsAttachArcCredentialsCmdRunE

	SimsCmd.AddCommand(SimsAttachArcCredentialsCmd)
}

// SimsAttachArcCredentialsCmd defines 'attach-arc-credentials' subcommand
var SimsAttachArcCredentialsCmd = &cobra.Command{
	Use:   "attach-arc-credentials",
	Short: TRAPI("/sims/{sim_id}/credentials/arc:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/credentials/arc:post:description`) + "\n\n" + createLinkToAPIReference("Sim", "attachArcSimCredentials"),
}

func SimsAttachArcCredentialsCmdRunE(cmd *cobra.Command, args []string) error {
	lib.WarnfStderr(TRCLI("cli.deprecated-api") + "\n")

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

	param, err := collectSimsAttachArcCredentialsCmdParams(ac)
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

func collectSimsAttachArcCredentialsCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSimsAttachArcCredentialsCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsAttachArcCredentialsCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimsAttachArcCredentialsCmd("/sims/{sim_id}/credentials/arc"),
		query:       buildQueryForSimsAttachArcCredentialsCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsAttachArcCredentialsCmd(path string) string {

	escapedSimId := url.PathEscape(SimsAttachArcCredentialsCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsAttachArcCredentialsCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimsAttachArcCredentialsCmd() (string, error) {
	var result map[string]interface{}

	if SimsAttachArcCredentialsCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimsAttachArcCredentialsCmdBody, "@") {
			fname := strings.TrimPrefix(SimsAttachArcCredentialsCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SimsAttachArcCredentialsCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SimsAttachArcCredentialsCmdBody)
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

	if SimsAttachArcCredentialsCmdArcClientPeerPublicKey != "" {
		result["arcClientPeerPublicKey"] = SimsAttachArcCredentialsCmdArcClientPeerPublicKey
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
