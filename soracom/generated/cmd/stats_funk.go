// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	StatsCmd.AddCommand(StatsFunkCmd)
}

// StatsFunkCmd defines 'funk' subcommand
var StatsFunkCmd = &cobra.Command{
	Use:   "funk",
	Short: TRCLI("cli.stats.funk.summary"),
	Long:  TRCLI(`cli.stats.funk.description`),
}