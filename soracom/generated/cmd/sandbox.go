package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(SandboxCmd)
}

// SandboxCmd defines 'sandbox' subcommand
var SandboxCmd = &cobra.Command{
	Use:   "sandbox",
	Short: TRCLI("cli.sandbox.summary"),
	Long:  TRCLI(`cli.sandbox.description`),
}
