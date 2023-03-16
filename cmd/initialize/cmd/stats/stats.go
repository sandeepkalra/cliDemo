/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package stats

import (
	"fmt"
	"local/cli/cmd/initialize/cmd"

	"github.com/spf13/cobra"
)

// StatsCmd represents the stats command
var StatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stats called")
	},
}

func init() {
	cmd.InitializeRootCmd.AddCommand(StatsCmd)
}
