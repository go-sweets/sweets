package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os/exec"
	"regexp"
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
	resp, err := http.Get("https://raw.githubusercontent.com/mix-plus/go-mixplus/master/tools/mpctl/cmd/version.go")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Upgrade go-mixplus cli failed.")
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	re := regexp.MustCompile(`CLIVersion\s+=\s+"(\d+\.\d+\.\d+)"`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) != 2 {
		fmt.Println("Could not find version number")
		return
	}

	version := matches[1]
	isUpgrade := VersionCompare(version, CLIVersion)
	switch isUpgrade {
	case 0:
		fmt.Println("The go-mixplus cli is the latest version.")
	case 1:
		fmt.Println("Upgrade mpctl to version:", version)
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
	case 2:
		fmt.Println("The go-mixplus cli is a higher version than the latest version.")
	}
}
