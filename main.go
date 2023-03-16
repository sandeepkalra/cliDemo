/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"local/cli/cmd"
	_ "local/cli/cmd/get"
	_ "local/cli/cmd/initialize"
	_ "local/cli/cmd/set"
)

func main() {
	cmd.Execute()
}
