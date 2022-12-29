package conversion

import (
	"encoding/json"
	"fmt"
	"nmigo/internal/fs_ops"
	"path/filepath"
)

type Conversion struct {
	Source                      DbConfig
	Target                      DbConfig
	MaxEachDbConnectionPoolSize int
	Encoding                    string
	Schema                      string
	ExcludeTables               []string
	IncludeTables               []string
	MigrateOnlyData             bool
	Delimiter                   string
	EnableExtraConfig           bool
	ExtraConfig                 *ExtraConfig // Note, ExtraConfig is nilable type.
	LogsDirPath                 string
	DataTypesMapAddr            string
	IndexTypesMapAddr           string
}

type ExtraConfig struct {
	Tables      []ExtraConfigTable
	ForeignKeys []ExtraConfigForeignKey
}

type ExtraConfigTable struct {
	Name    ExtraConfigTableName
	Columns []ExtraConfigTableColumn
}

type ExtraConfigTableName struct {
	Original string
	New      string
}

type ExtraConfigTableColumn struct {
	Original string
	New      string
}

type ExtraConfigForeignKey struct {
	ConstraintName       string
	TableName            string
	ColumnName           string
	ReferencedTableName  string
	ReferencedColumnName string
	UpdateRule           string
	DeleteRule           string
}

type DbConfig struct {
	Host     string
	Port     int
	Database string
	Charset  string
	User     string
	Password string
}

func InitializeConversion(baseDir string) *Conversion {
	fileName := filepath.Join(baseDir, "config", "config.json")
	contents := fs_ops.ReadWholeFile(fileName)
	conversion := parseConfig(contents)
	conversion.LogsDirPath = filepath.Join(baseDir, "logs_directory")
	conversion.DataTypesMapAddr = filepath.Join(baseDir, "config", "data_types_map.json")
	conversion.IndexTypesMapAddr = filepath.Join(baseDir, "config", "index_types_map.json")

	if conversion.EnableExtraConfig {
		fileName = filepath.Join(baseDir, "config", "extra_config.json")
		contents = fs_ops.ReadWholeFile(fileName)
		conversion.ExtraConfig = parseExtraConfig(contents)
	}

	return &conversion
}

func parseConfig(configData []byte) Conversion {
	var conversion Conversion
	jsonUnmarshal(configData, &conversion)
	return conversion
}

func parseExtraConfig(configData []byte) *ExtraConfig {
	var extraConfig ExtraConfig
	jsonUnmarshal(configData, &extraConfig)
	return &extraConfig
}

func jsonUnmarshal(data []byte, valueContainer interface{}) {
	err := json.Unmarshal(data, &valueContainer)

	if err != nil {
		fmt.Printf("Error %v", err)
		panic(err)
	}
}
