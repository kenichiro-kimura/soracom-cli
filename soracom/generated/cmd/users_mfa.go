package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	UsersCmd.AddCommand(UsersMfaCmd)
}

// UsersMfaCmd defines 'mfa' subcommand
var UsersMfaCmd = &cobra.Command{
	Use:   "mfa",
	Short: TRCLI("cli.users.mfa.summary"),
	Long:  TRCLI(`cli.users.mfa.description`),
}
