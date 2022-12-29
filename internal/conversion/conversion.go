package conversion

import (
	"encoding/json"
	"fmt"
	"nmigo/internal/fs_ops"
	"os"
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
	ExtraConfig                 *ExtraConfig // Note, ExtraConfig is type.
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

func InitializeConversion() *Conversion {
	baseDir := getBaseDirectory()
	fileName := filepath.Join(baseDir, "config", "config.json")
	contents := fs_ops.ReadWholeFile(fileName)
	conversion := parseConfig(contents)

	if conversion.EnableExtraConfig {
		fileName = filepath.Join(baseDir, "config", "extra_config.json")
		contents = fs_ops.ReadWholeFile(fileName)
		conversion.ExtraConfig = parseExtraConfig(contents)
	}

	return &conversion
}

func getBaseDirectory() string {
	baseDir, ok := os.LookupEnv("aux_dir")

	if !ok {
		pwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		baseDir = filepath.Join(pwd, "..", "..")
	}

	return baseDir
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

func jsonMarshalSliceOfEmptyInterfaces(data []interface{}) (result []byte) {
	for _, dataItem := range data {
		result = append(result, jsonMarshalInterface(dataItem)...)
	}

	return result
}

func jsonMarshalInterface(data interface{}) []byte {
	bytes, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("Error %v", err)
		panic(err)
	}

	return bytes
}

func jsonMarshal(data map[string]interface{}) []byte {
	bytes, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("Error %v", err)
		panic(err)
	}

	return bytes
}

func jsonMarshalMapStringString(data map[string]string) []byte {
	bytes, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("Error %v", err)
		panic(err)
	}

	return bytes
}

func jsonUnmarshal(data []byte, valueContainer interface{}) {
	err := json.Unmarshal(data, &valueContainer)

	if err != nil {
		fmt.Printf("Error %v", err)
		panic(err)
	}
}

func convertSliceOfInterfacesToSliceOfStrings(configValue interface{}) []string {
	var excludeTablesSlice []string

	for _, tableToExclude := range configValue.([]interface{}) {
		excludeTablesSlice = append(excludeTablesSlice, fmt.Sprint(tableToExclude))
	}

	return excludeTablesSlice
}
