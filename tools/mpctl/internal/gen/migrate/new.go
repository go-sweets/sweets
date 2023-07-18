package migrate

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

var migrateNew = &cobra.Command{
	Use:   "new",
	Short: "Create a new a database migration.",
	Long:  "Create a new a database migration.",
	Run:   MigrateNewRun,
}

var templateContent = `
-- +migrate Up

-- +migrate Down
`
var tpl *template.Template

func init() {
	tpl = template.Must(template.New("new_migration").Parse(templateContent))
}

func MigrateNewRun(cmd *cobra.Command, args []string) {
	name := cmd.Flags().Arg(0)
	if name == "" {
		panic("Please provide a name for the migration.")
	}
	ConfigFlags(cmd)
	env, err := GetEnvironment()
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%s-%s.sql", time.Now().Format("20060102150405"), strings.TrimSpace(name))
	pathName := path.Join(env.Dir, fileName)
	f, err := os.Create(pathName)

	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	if err := tpl.Execute(f, nil); err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Created migration %s", pathName))
}
