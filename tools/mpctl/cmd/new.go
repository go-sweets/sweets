package cmd

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/mix-go/xcli/flag"
	"github.com/mix-plus/cli/logic"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var NewCmd = &cobra.Command{
	Use:     "new",
	Short:   "create a service template",
	Long:    "Create a service project using the repository template. Example: mpctl new helloworld",
	Run:     newRun,
	Version: SkeletonVersion,
}

var (
	repoURL string
	name    string
)

func init() {
	repoURL = "github.com/mix-plus/mixplus-skeleton"

	name = flag.Arguments().First().String("hello")
	name = strings.ReplaceAll(name, " ", "")
}

func newRun(_ *cobra.Command, args []string) {
	envCmd := "go env GOPATH"
	cmd := exec.Command("go", "env", "GOPATH")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Exec error: %v\n", err)
		return
	}
	goPath := string(out[:len(out)-1])
	if goPath == "" {
		fmt.Printf("$GOPATH is not configured, see '%s'\n", envCmd)
		return
	}

	dr := ":"
	if runtime.GOOS == "windows" {
		dr = ";"
	}
	if strings.Contains(goPath, dr) {
		fmt.Printf("$GOPATH cannot have multiple directories, see '%s'\n", envCmd)
		return
	}

	srcDir := fmt.Sprintf("%s/pkg/mod/%s", goPath, repoURL+"@v"+SkeletonVersion)
	if _, err := os.Stat(srcDir); err != nil {
		fmt.Printf("No skeleton found '%s'\n", repoURL)

		installFunc := func(installCmd string) error {
			cmd := exec.Command("go", installCmd, fmt.Sprintf("%s", repoURL))
			fmt.Printf("Install skeleton 'go %s %s'\n", installCmd, repoURL)
			total := 0
			current := int64(0)
			bar := pb.StartNew(total)
			go func() {
				path := fmt.Sprintf("%s/pkg/mod/cache/download/%s/@v/%s.zip", goPath, repoURL, SkeletonVersion)
				fmt.Println("path ", path)
				for {
					f, err := os.Open(path)
					if err != nil {
						continue
					}
					fi, err := f.Stat()
					if err != nil {
						_ = f.Close()
						continue
					}
					current = fi.Size()
					bar.SetCurrent(current)
					_ = f.Close()
					time.Sleep(time.Millisecond * 100)
				}
			}()
			err = cmd.Run()
			if err == nil {
				bar.SetTotal(current)
				bar.SetCurrent(current)
			} else {
				bar.SetTotal(0)
				bar.SetCurrent(0)
			}
			bar.Finish()
			return err
		}

		if err1 := installFunc("get"); err1 != nil {
			if err2 := installFunc("install"); err2 != nil {
				fmt.Println(fmt.Sprintf("Install failed: %s: %s", err1, err2))
				return
			}
		}

		time.Sleep(2 * time.Second) // 等待一会，让 gomod 完成解压
		_ = os.Remove(fmt.Sprintf("%s/bin/mixplus-skeleton", goPath))
		fmt.Println(fmt.Sprintf("Skeleton '%s' installed successfully", repoURL))
	} else {
		fmt.Println(fmt.Sprintf("Local skeleton found '%s'", repoURL))
	}

	fmt.Print(" - Generate code")
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dest := fmt.Sprintf("%s/%s", pwd, name)
	if !logic.CopyPath(srcDir, dest) {
		panic(errors.New(fmt.Sprintf("copy dir failed srcdir %s to %s", srcDir, dest)))
	}
	fmt.Println(" > ok")

	fmt.Print(" - Processing package name")
	if err := logic.ReplaceAll(dest, fmt.Sprintf("%s", repoURL), name); err != nil {
		panic(errors.New("replace failed"))
	}
	if err := logic.ReplaceMod(dest); err != nil {
		panic(errors.New("replace go.mod failed"))
	}
	fmt.Println(" > ok")

	fmt.Println(fmt.Sprintf("Project '%s' is generated", name))
}
