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
	ExtraConfig                 ExtraConfig
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
	fileName := filepath.Join(baseDir, "configs", "config.json")
	contents := fs_ops.ReadWholeFile(fileName)
	conversion := parseConfig(contents)

	if conversion.EnableExtraConfig {
		fileName = filepath.Join(baseDir, "configs", "extra_config.json")
		contents = fs_ops.ReadWholeFile(fileName)
		conversion.ExtraConfig = parseExtraConfig(contents)
	}

	return &conversion
}

func getBaseDirectory() string {
	baseDir, auxDirIsSet := os.LookupEnv("aux_dir")

	if !auxDirIsSet {
		pwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		baseDir = pwd
	}

	return baseDir
}

func parseConfig(configData []byte) Conversion {
	var config map[string]interface{}
	jsonUnmarshal(configData, &config)

	var source DbConfig
	bytesSliceSource := jsonMarshal(config["source"].(map[string]interface{}))
	jsonUnmarshal(bytesSliceSource, &source)

	var target DbConfig
	bytesSliceTarget := jsonMarshal(config["target"].(map[string]interface{}))
	jsonUnmarshal(bytesSliceTarget, &target)

	maxEachDbConnectionPoolSize := int(config["max_each_db_connection_pool_size"].(float64))
	encoding := config["encoding"].(string)
	schema := config["schema"].(string)
	excludeTables := convertSliceOfInterfacesToSliceOfStrings(config["exclude_tables"])
	includeTables := convertSliceOfInterfacesToSliceOfStrings(config["include_tables"])
	migrateOnlyData := config["migrate_only_data"].(bool)
	delimiter := config["delimiter"].(string)
	enableExtraConfig := config["enable_extra_config"].(bool)

	return Conversion{
		Source:                      source,
		Target:                      target,
		MaxEachDbConnectionPoolSize: maxEachDbConnectionPoolSize,
		Encoding:                    encoding,
		Schema:                      schema,
		ExcludeTables:               excludeTables,
		IncludeTables:               includeTables,
		MigrateOnlyData:             migrateOnlyData,
		Delimiter:                   delimiter,
		EnableExtraConfig:           enableExtraConfig,
	}
}

func parseExtraConfig(configData []byte) ExtraConfig {
	var config map[string]interface{}
	var extraConfig ExtraConfig
	//var extraConfigTables []ExtraConfigTable
	var extraConfigForeignKeys []ExtraConfigForeignKey
	//var extraConfigTableName ExtraConfigTableName
	//var extraConfigTableColumns []ExtraConfigTableColumn

	jsonUnmarshal(configData, &config)

	// var xx []map[string]string
	x := config["foreign_keys"].([]interface{})
	fmt.Println(x)
	bytesForeignKeys := jsonMarshalSliceOfEmptyInterfaces(x)
	jsonUnmarshal(bytesForeignKeys, &extraConfigForeignKeys)

	return extraConfig
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
