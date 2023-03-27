package main

import (
	"github.com/mix-plus/go-mixplus/tools/mpctl/cmd"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "go-mixplus",
	Short:   "go-mixplus: An elegant toolkit for Go microservices.",
	Long:    `go-mixplus: An elegant toolkit for Go microservices.`,
	Version: cmd.CLIVersion,
}

func init() {
	rootCmd.AddCommand(cmd.NewCmd)
	rootCmd.AddCommand(cmd.UpgradeCmd)
}
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
