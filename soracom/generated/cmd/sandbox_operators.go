package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	SandboxCmd.AddCommand(SandboxOperatorsCmd)
}

// SandboxOperatorsCmd defines 'operators' subcommand
var SandboxOperatorsCmd = &cobra.Command{
	Use:   "operators",
	Short: TRCLI("cli.sandbox.operators.summary"),
	Long:  TRCLI(`cli.sandbox.operators.description`),
}
