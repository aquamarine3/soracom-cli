package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd defines 'soracom' command
var RootCmd = &cobra.Command{
	Use:   "soracom",
	Short: "soracom command",
	Long:  `A command line tool to invoke SORACOM API`,
}

var specifiedProfileName string
var specifiedCoverageType string

func init() {
	RootCmd.PersistentFlags().StringVar(&specifiedProfileName, "profile", "", "Specify profile name")
	RootCmd.PersistentFlags().StringVar(&specifiedCoverageType, "coverage-type", "", "Specify coverage type")
}
