package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var UpgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "upgrade the go-mixplus cli",
	Long:    "upgrade the go-mixplus cli",
	Run:     upgradeRun,
	Version: CLIVersion,
}

var cliRepoUrl string

func init() {
	cliRepoUrl = "github.com/mix-plus/go-mixplus/tools/mpctl"
}

func upgradeRun(_ *cobra.Command, args []string) {
	upgradeFunc := func(upgradeCmd string) (err error) {
		cmd := exec.Command("go", upgradeCmd, cliRepoUrl+"@v"+CLIVersion)
		fmt.Printf("Upgrade go-mixplus cli: %s\n", cmd.String())
		err = cmd.Run()
		if err != nil {
			fmt.Println("Upgrade go-mixplus cli failed.", err.Error())
		}
		return err
	}
	if err1 := upgradeFunc("get"); err1 != nil {
		if err2 := upgradeFunc("install"); err2 != nil {
			fmt.Println("Upgrade go-mixplus cli failed.")
		}
	}

	fmt.Println(" > ok.")
}
