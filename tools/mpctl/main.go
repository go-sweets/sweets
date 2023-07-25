package main

import (
	"github.com/mix-plus/go-mixplus/tools/mpctl/internal"
	"github.com/mix-plus/go-mixplus/tools/mpctl/internal/gen"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "go-mixplus",
	Short:   "go-mixplus: An elegant toolkit for Go microservices.",
	Long:    `go-mixplus: An elegant toolkit for Go microservices.`,
	Version: internal.CLIVersion,
}

func init() {
	rootCmd.AddCommand(internal.NewCmd)
	rootCmd.AddCommand(internal.UpgradeCmd)
	rootCmd.AddCommand(gen.CmdGen)
}
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
