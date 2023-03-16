/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package probes

import (
	"fmt"
	"local/cli/cmd/initialize/cmd"

	"github.com/spf13/cobra"
)

// ProbesCmd represents the probes command
var ProbesCmd = &cobra.Command{
	Use:   "probes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("probes called")
	},
}

func init() {
	cmd.InitializeRootCmd.AddCommand(ProbesCmd)
}
