package gorm

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mix-plus/go-mixplus/pkg/contains"
	"github.com/mix-plus/go-mixplus/pkg/plugins/gorm/filter"
	"github.com/mix-plus/go-mixplus/pkg/str"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var CmdGorm = &cobra.Command{
	Use:   "gorm",
	Short: "Generate gorm model from database. Example: mpctl gen gorm",
	Long:  "Generate gorm model from database. Example: mpctl gen gorm",
	Run:   run,
}

var (
	config             string
	dsn                string
	db                 string
	tables             string
	exclude            string
	association        string
	outPath            string
	outFile            string
	modelPkgName       string
	fieldWithStringTag string
	onlyModel          bool
	withUnitTest       bool
	fieldNullable      bool
	fieldWithIndexTag  bool
	fieldWithTypeTag   bool
	fieldSignable      bool
)

// argParse is parser for cmd
func init() {
	config = "etc/gen.yml"
	dsn = ""
	db = "mysql"
	tables = ""
	exclude = "schema_migrations"
	association = ""
	outPath = "internal/dao/query"
	outFile = ""
	modelPkgName = "internal/dao/model"
	fieldWithStringTag = ""
	onlyModel = false
	withUnitTest = false
	fieldNullable = false
	fieldWithIndexTag = false
	fieldWithTypeTag = false
	fieldSignable = true
	CmdGorm.PersistentFlags().StringVarP(&config, "config", "c", config, "is path for gen.yml")
	CmdGorm.PersistentFlags().StringVarP(&dsn, "dsn", "", dsn, "consult[https://gorm.io/docs/connecting_to_the_database.html]")
	CmdGorm.PersistentFlags().StringVarP(&db, "db", "", db, "input mysql|postgres|sqlite|sqlserver|clickhouse. consult[https://gorm.io/docs/connecting_to_the_database.html]")
	CmdGorm.PersistentFlags().StringVarP(&tables, "tables", "t", tables, "enter the required data table or leave it blank")
	CmdGorm.PersistentFlags().StringVarP(&exclude, "exclude", "e", exclude, "enter the exclude data table or leave it blank")
	CmdGorm.PersistentFlags().StringVarP(&association, "association", "a", association, "enter the association data table or leave it blank, index1: table name; index2: relation table name; index3: field name; index4: relation type(has_one/has_many/belongs_to/many_to_many); index5: gorm tag. Example: -a \"user|role|Role|has_one|foreignKey:RoleID\"")
	CmdGorm.PersistentFlags().StringVarP(&outPath, "outPath", "p", outPath, "specify a directory for output")
	CmdGorm.PersistentFlags().StringVarP(&outFile, "outFile", "", outFile, "query code file name, default: gen.go")
	CmdGorm.PersistentFlags().StringVarP(&modelPkgName, "modelPkgName", "m", modelPkgName, "generated model code's package name")
	CmdGorm.PersistentFlags().StringVarP(&fieldWithStringTag, "fieldWithStringTag", "s", fieldWithStringTag, "field need add ,string json tag, index1: table name; index i: field name Example: -s \"user|role_id,user_group|lock_expire|wrong\"")
	CmdGorm.PersistentFlags().BoolVarP(&onlyModel, "onlyModel", "o", onlyModel, "only generate models (without query file)")
	CmdGorm.PersistentFlags().BoolVarP(&withUnitTest, "withUnitTest", "", withUnitTest, "generate unit test for query code")
	CmdGorm.PersistentFlags().BoolVarP(&fieldNullable, "fieldNullable", "", fieldNullable, "generate with pointer when field is nullable")
	CmdGorm.PersistentFlags().BoolVarP(&fieldWithIndexTag, "fieldWithIndexTag", "", fieldWithIndexTag, "generate field with gorm index tag")
	CmdGorm.PersistentFlags().BoolVarP(&fieldWithTypeTag, "fieldWithTypeTag", "", fieldWithTypeTag, "generate field with gorm column type tag")
	CmdGorm.PersistentFlags().BoolVarP(&fieldSignable, "fieldSignable", "", fieldSignable, "detect integer field's unsigned type, adjust generated data type")
}

// DBType database type
type DBType string

const (
	// dbMySQL Gorm Drivers mysql || postgres || sqlite || sqlserver
	dbMySQL      DBType = "mysql"
	dbPostgres   DBType = "postgres"
	dbSQLite     DBType = "sqlite"
	dbSQLServer  DBType = "sqlserver"
	dbClickHouse DBType = "clickhouse"
)

type CmdGenParams struct {
	DSN                *string   `yaml:"dsn"`
	DB                 *string   `yaml:"db"`
	Tables             *[]string `yaml:"tables"`
	Exclude            *[]string `yaml:"exclude"`
	Association        *[]string `yaml:"association"`
	OutPath            *string   `yaml:"outPath" mapstructure:"out-path"`
	OutFile            *string   `yaml:"outFile" mapstructure:"out-file"`
	ModelPkgName       *string   `yaml:"modelPkgName" mapstructure:"model-pkg-name"`
	FieldWithStringTag *[]string `yaml:"fieldWithStringTag" mapstructure:"field-with-string-tag"`
	OnlyModel          *bool     `yaml:"onlyModel" mapstructure:"only-model"`
	WithUnitTest       *bool     `yaml:"withUnitTest" mapstructure:"with-unit-test"`
	FieldNullable      *bool     `yaml:"fieldNullable" mapstructure:"field-nullable"`
	FieldWithIndexTag  *bool     `yaml:"fieldWithIndexTag" mapstructure:"field-with-index-tag"`
	FieldWithTypeTag   *bool     `yaml:"fieldWithTypeTag" mapstructure:"field-with-type-tag"`
	FieldSignable      *bool     `yaml:"fieldSignable" mapstructure:"field-signable"`
}

type CmdParams struct {
	Gen *CmdGenParams `yaml:"gen"`
}

// connectDB choose db type for connection to database
func connectDB(t DBType, dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("dsn cannot be empty")
	}

	switch t {
	case dbMySQL:
		return gorm.Open(mysql.Open(dsn))
	case dbPostgres:
		return gorm.Open(postgres.Open(dsn))
	case dbSQLite:
		return gorm.Open(sqlite.Open(dsn))
	case dbSQLServer:
		return gorm.Open(sqlserver.Open(dsn))
	case dbClickHouse:
		return gorm.Open(clickhouse.Open(dsn))
	default:
		return nil, fmt.Errorf("unknow db %q (support mysql || postgres || sqlite || sqlserver for now)", t)
	}
}

// genModels is gorm/gen generated models
func genModels(cfg *CmdGenParams) (err error) {
	targetTables := *cfg.Tables
	if len(targetTables) == 0 {
		targetTables, err = newDB(cfg).Migrator().GetTables()
		if err != nil {
			return fmt.Errorf("GORM migrator get all tables fail: %w", err)
		}
	}

	excludeTables := *cfg.Exclude
	// get association tables by associationTables
	for _, item := range *cfg.Association {
		arr := strings.Split(item, "|")
		if !contains.Contains(excludeTables, arr[0]) {
			excludeTables = append(excludeTables, arr[0])
		}
		if !contains.Contains(excludeTables, arr[1]) {
			excludeTables = append(excludeTables, arr[1])
		}
	}

	// remove excludeTables
	simpleTables := make([]string, 0, len(targetTables))
	for _, item := range targetTables {
		if !contains.Contains(excludeTables, item) {
			simpleTables = append(simpleTables, item)
		}
	}

	var models []interface{}
	g := newGenerator(cfg)
	relations := make([]string, 0, len(*cfg.Association))
	sources := make([]string, 0, len(*cfg.Association))
	for _, item := range *cfg.Association {
		arr := strings.Split(item, "|")
		arr2 := strings.Split(arr[4], ":")
		source := arr[0]
		relation := arr[1]
		fieldName := arr[2]
		// get gorm tag
		tag := field.GormTag{}
		tag.Set(arr2[0], arr2[1])
		// save source and relation
		if !contains.Contains(relations, relation) {
			relations = append(relations, relation)
		}
		if !contains.Contains(sources, source) {
			sources = append(sources, source)
		}

		relationNs := needAddStringTag(cfg, relation)
		sourceNs := needAddStringTag(cfg, source)
		// generate model with opt
		m := g.GenerateModel(
			source,
			gen.FieldRelate(
				field.RelationshipType(arr[3]),
				fieldName,
				newGenerator(cfg).GenerateModel(relation, gen.FieldJSONTagWithNS(relationNs)),
				&field.RelateConfig{GORMTag: tag},
			),
			gen.FieldJSONTagWithNS(sourceNs),
		)
		models = append(models, m)
	}

	// relation in sources means generate model with opt, not in is simple
	for _, item := range relations {
		if !contains.Contains(sources, item) {
			simpleTables = append(simpleTables, item)
		}
	}

	for _, item := range simpleTables {
		ns := needAddStringTag(cfg, item)
		m := g.GenerateModel(item, gen.FieldJSONTagWithNS(ns))
		models = append(models, m)
	}

	if !*cfg.OnlyModel {
		g.ApplyInterface(func(filter.Filter) {}, models...)
	}

	g.Execute()
	return
}

func run(cmd *cobra.Command, args []string) {
	cfg, err := parseConfig(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: parse config fail: %s\033[m\n", err.Error())
		return
	}

	err = genModels(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: Generate gorm failed: %s\033[m\n", err.Error())
		return
	}

	fmt.Println("\n🍺 Generate gorm succeeded")
	fmt.Printf("model path %s\n", color.GreenString(*cfg.OutPath))
	fmt.Printf("query path %s\n", color.GreenString(*cfg.ModelPkgName))
}

func newDB(cfg *CmdGenParams) *gorm.DB {
	gormDB, err := connectDB(DBType(*cfg.DB), *cfg.DSN)
	if err != nil {
		log.Fatalln("connect db server fail:", err)
	}
	return gormDB
}

func needAddStringTag(cfg *CmdGenParams, tableName string) func(columnName string) string {
	for _, item := range *cfg.FieldWithStringTag {
		arr := strings.Split(item, "|")
		if arr[0] == tableName {
			newArr := make([]string, len(arr)-1, len(arr)-1)
			copy(newArr, arr[1:])
			return func(columnName string) string {
				if columnName == "id" {
					return "id,string"
				}
				tag := str.CamelCaseLowerFirst(columnName)
				for _, v := range newArr {
					if v == columnName {
						tag = tag + ",string"
						break
					}
				}
				return tag
			}
		}
	}
	return nil
}

func newGenerator(cfg *CmdGenParams) *gen.Generator {
	g := gen.NewGenerator(gen.Config{
		OutPath:           *cfg.OutPath,
		OutFile:           *cfg.OutFile,
		ModelPkgPath:      *cfg.ModelPkgName,
		WithUnitTest:      *cfg.WithUnitTest,
		FieldNullable:     *cfg.FieldNullable,
		FieldWithIndexTag: *cfg.FieldWithIndexTag,
		FieldWithTypeTag:  *cfg.FieldWithTypeTag,
		FieldSignable:     *cfg.FieldSignable,
	})
	g.UseDB(newDB(cfg))
	var dataMap = map[string]func(gorm.ColumnType) (dataType string){
		"decimal": func(columnType gorm.ColumnType) (dataType string) {
			return "decimal.Decimal"
		},
		"datetime": func(columnType gorm.ColumnType) (dataType string) {
			return "carbon.DateTime"
		},
	}

	g.WithDataTypeMap(dataMap)
	g.WithImportPkgPath("github.com/golang-module/carbon/v2")
	g.WithJSONTagNameStrategy(func(columnName string) string {
		if columnName == "id" {
			return "id,string"
		}
		return str.CamelCaseLowerFirst(columnName)
	})
	return g
}

func parseConfig(cmd *cobra.Command) (*CmdGenParams, error) {
	viper.SetDefault("gen.dsn", dsn)
	viper.SetDefault("gen.db", db)
	viper.SetDefault("gen.tables", tables)
	viper.SetDefault("gen.exclude", exclude)
	viper.SetDefault("gen.association", association)
	viper.SetDefault("gen.out-path", outPath)
	viper.SetDefault("gen.out-file", outFile)
	viper.SetDefault("gen.model-pkg-name", modelPkgName)
	viper.SetDefault("gen.field-with-string-tag", fieldWithStringTag)
	viper.SetDefault("gen.only-model", onlyModel)
	viper.SetDefault("gen.with-unit-test", withUnitTest)
	viper.SetDefault("gen.field-nullable", fieldNullable)
	viper.SetDefault("gen.field-with-index-tag", fieldWithIndexTag)
	viper.SetDefault("gen.field-with-type-tag", fieldWithTypeTag)
	viper.SetDefault("gen.field-signable", fieldSignable)

	configPath, _ := cmd.Flags().GetString("config")
	if configPath != "" {
		viper.SetConfigFile(configPath)
		// if yml value not exist use default value
		viper.MergeInConfig()
	}
	var cfg CmdParams
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("viper failed to parse config: %w", err)
	}
	for _, item := range *cfg.Gen.Association {
		arr := strings.Split(item, "|")
		if len(arr) != 5 {
			return nil, errors.Errorf("invalid association tables: %s", item)
		}
		arr2 := strings.Split(arr[4], ":")
		if len(arr2) != 2 {
			return nil, errors.Errorf("invalid association tables gorm tag: %s", arr[4])
		}
	}
	for _, item := range *cfg.Gen.FieldWithStringTag {
		arr := strings.Split(item, "|")
		if len(arr) <= 1 {
			return nil, errors.Errorf("invalid field with string tag: %s", item)
		}
	}
	return cfg.Gen, nil
}
