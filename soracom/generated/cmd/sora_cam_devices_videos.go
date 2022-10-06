// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	SoraCamDevicesCmd.AddCommand(SoraCamDevicesVideosCmd)
}

// SoraCamDevicesVideosCmd defines 'videos' subcommand
var SoraCamDevicesVideosCmd = &cobra.Command{
	Use:   "videos",
	Short: TRCLI("cli.sora-cam.devices.videos.summary"),
	Long:  TRCLI(`cli.sora-cam.devices.videos.description`),
}
