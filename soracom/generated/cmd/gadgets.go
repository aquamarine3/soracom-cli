// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(GadgetsCmd)
}

// GadgetsCmd defines 'gadgets' subcommand
var GadgetsCmd = &cobra.Command{
	Use:   "gadgets",
	Short: TRCLI("cli.gadgets.summary"),
	Long:  TRCLI(`cli.gadgets.description`),
}
