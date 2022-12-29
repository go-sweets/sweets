package commands

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/manifoldco/promptui"
	"github.com/mix-go/xcli/flag"
	"github.com/mix-plus/cli/logic"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var (
	API   = "api"
	gRPC  = "grpc"
	Queue = "queue"
)

type NewCommand struct {
}

func (t *NewCommand) Main() {
	name := flag.Arguments().First().String("hello")
	name = strings.ReplaceAll(name, " ", "")

	promp := func(label string, items []string) string {
		prompt := promptui.Select{
			Label: label,
			Items: items,
		}
		prompt.HideSelected = true
		_, result, err := prompt.Run()
		if err != nil {
			return ""
		}
		return result
	}

	var selectType string
	switch promp("Select project type", []string{"API", "gRPC", "Queue"}) {
	case "API":
		selectType = API
		break
	case "gRPC":
		selectType = gRPC
		break
	case "Queue":
		selectType = Queue
		break
	default:
		return
	}

	t.NewProject(name, selectType)
}

func (t *NewCommand) NewProject(name, selectType string) {
	ver := ""
	switch selectType {
	case API, gRPC, Queue:
		ver = fmt.Sprintf("v%s", SkeletonVersion)
		break
	default:
		fmt.Println("Type error, only be console, api, web, grpc")
		return
	}

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

	srcDir := fmt.Sprintf("%s/pkg/mod/github.com/mix-plus/%s-skeleton@%s", goPath, selectType, ver)
	if _, err := os.Stat(srcDir); err != nil {
		fmt.Printf("No skeleton found 'github.com/mix-plus/%s-skeleton@%s'\n", selectType, ver)

		installFunc := func(installCmd string) error {
			cmd := exec.Command("go", installCmd, fmt.Sprintf("github.com/mix-plus/%s-skeleton@%s", selectType, ver))
			fmt.Printf("Install skeleton 'go %s github.com/mix-plus/%s-skeleton@%s'\n", installCmd, selectType, ver)
			total := 0
			fmt.Println("SelectType ", selectType)
			switch selectType {
			case API:
				total = 13834
			case gRPC:
				total = 15659
			}
			current := int64(0)
			bar := pb.StartNew(total)
			go func() {
				path := fmt.Sprintf("%s/pkg/mod/cache/download/github.com/mix-plus/%s-skeleton/@v/%s.zip", goPath, selectType, ver)
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
		_ = os.Remove(fmt.Sprintf("%s/bin/%s-skeleton", goPath, selectType))
		fmt.Println(fmt.Sprintf("Skeleton 'github.com/mix-plus/%s-skeleton@%s' installed successfully", selectType, ver))
	} else {
		fmt.Println(fmt.Sprintf("Local skeleton found 'github.com/mix-plus/%s-skeleton@%s'", selectType, ver))
	}

	fmt.Print(" - Generate code")
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dest := fmt.Sprintf("%s/%s", pwd, name)
	if !logic.CopyPath(srcDir, dest) {
		panic(errors.New("copy dir failed"))
	}
	fmt.Println(" > ok")

	fmt.Print(" - Processing package name")
	if err := logic.ReplaceAll(dest, fmt.Sprintf("github.com/mix-plus/%s-skeleton", selectType), name); err != nil {
		panic(errors.New("replace failed"))
	}
	if err := logic.ReplaceMod(dest); err != nil {
		panic(errors.New("replace go.mod failed"))
	}
	fmt.Println(" > ok")

	fmt.Println(fmt.Sprintf("Project '%s' is generated", name))
}
