module github.com/mix-plus/go-mixplus/tools/mpctl

go 1.18

require (
	github.com/cheggaaa/pb/v3 v3.1.0
	github.com/fatih/color v1.13.0
	github.com/go-gorp/gorp/v3 v3.1.0
	github.com/go-sql-driver/mysql v1.7.0
	github.com/lib/pq v1.10.7
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/mix-go/xcli v1.1.20
	github.com/mix-plus/go-mixplus/pkg/contains v1.1.0
	github.com/mix-plus/go-mixplus/pkg/plugins/gorm/filter v1.1.0
	github.com/mix-plus/go-mixplus/pkg/str v1.1.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pkg/errors v0.9.1
	github.com/rubenv/sql-migrate v1.5.1
	github.com/spf13/cobra v1.6.1
	github.com/spf13/viper v1.14.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/clickhouse v0.5.1
	gorm.io/driver/mysql v1.5.1
	gorm.io/driver/postgres v1.5.2
	gorm.io/driver/sqlite v1.5.2
	gorm.io/driver/sqlserver v1.5.1
	gorm.io/gen v0.3.23
	gorm.io/gorm v1.25.2
)

require (
	github.com/ClickHouse/ch-go v0.53.0 // indirect
	github.com/ClickHouse/clickhouse-go/v2 v2.8.3 // indirect
	github.com/VividCortex/ewma v1.1.1 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-faster/city v1.0.1 // indirect
	github.com/go-faster/errors v0.6.1 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/microsoft/go-mssqldb v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/paulmach/orb v0.9.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	go.opentelemetry.io/otel v1.14.0 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/datatypes v1.1.1-0.20230130040222-c43177d3cf8c // indirect
	gorm.io/hints v1.1.0 // indirect
	gorm.io/plugin/dbresolver v1.3.0 // indirect
)

replace (
	github.com/mix-plus/go-mixplus/pkg/contains => ../../pkg/contains
	github.com/mix-plus/go-mixplus/pkg/plugins/gorm/filter => ../../pkg/plugins/gorm/filter
	github.com/mix-plus/go-mixplus/pkg/str => ../../pkg/str
)
