/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"local/cli/cmd"
	"local/cli/cmd/initialize/cmd/fsm"
	"local/cli/cmd/initialize/cmd/memory"
	"local/cli/cmd/initialize/cmd/probes"
	"local/cli/cmd/initialize/cmd/stats"

	"github.com/spf13/cobra"
)

// initializeCmd represents the initialize command
var initializeCmd = &cobra.Command{
	Use:   "initialize",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	initializeCmd.AddCommand(probes.ProbesCmd)
	initializeCmd.AddCommand(stats.StatsCmd)
	initializeCmd.AddCommand(memory.MemoryCmd)
	initializeCmd.AddCommand(fsm.FSMCmd)
	cmd.RootCmd.AddCommand(initializeCmd)
}
