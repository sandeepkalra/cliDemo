/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"local/cli/cmd"
	"local/cli/cmd/get/cmd/probes"
	"local/cli/cmd/get/cmd/stats"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	getCmd.AddCommand(probes.ProbesCmd)
	getCmd.AddCommand(stats.StatsCmd)
	cmd.RootCmd.AddCommand(getCmd)
}
