package template

import (
	"bytes"
	"fmt"
	"github.com/small-ek/ant-cli/utils"
)

type TableStructure struct {
	COLUMN_NAME    string `json:"COLUMN_NAME"`
	DATA_TYPE      string `json:"DATA_TYPE"`
	COLUMN_COMMENT string `json:"COLUMN_COMMENT"`
	COLUMN_TYPE    string `json:"COLUMN_TYPE"`
	COLUMN_KEY     string `json:"COLUMN_KEY"`
}

// database 数据库名称
// table 表名称
// tableStructure 表结构
// GenGormModel 生成gorm模型
func GenGormModel(database, table string, tableStructure []TableStructure) string {
	var buffer bytes.Buffer
	packageStr := fmt.Sprintf("package model \n\n")

	isImportDate := false
	isImportJson := false
	isImportDeletedAt := false
	for _, col := range tableStructure {
		if col.COLUMN_NAME == "deleted_at" || col.COLUMN_NAME == "delete_time" {
			isImportDeletedAt = true
		}

		switch col.DATA_TYPE {
		case "date", "datetime", "timestamp":
			isImportDate = true
		case "json":
			isImportJson = true
		}
	}
	buffer.WriteString(packageStr)
	importStr := fmt.Sprintf("import (\n")
	buffer.WriteString(importStr)

	//是否导入JSON
	if isImportJson == true {
		importSqlStr := fmt.Sprintf(`"github.com/small-ek/antgo/db/adb/sql"`)
		buffer.WriteString(importSqlStr)
	}
	//是否导软删除
	if isImportDeletedAt == true {
		importTimeStr := fmt.Sprintf(`
"gorm.io/gorm"`)
		buffer.WriteString(importTimeStr)
	}
	//是否导入时间
	if isImportDate == true {
		importTimeStr := fmt.Sprintf(`
"time"`)
		buffer.WriteString(importTimeStr)
	}

	buffer.WriteString("\n)\n\n")
	structStr := fmt.Sprintf("type %s struct { \n", utils.ToCamelCase(table))
	buffer.WriteString(structStr)
	for _, col := range tableStructure {
		if col.COLUMN_NAME == "id" {
			buffer.WriteString(fmt.Sprintf("    %s %s `gorm:\"column:%s;primaryKey;autoIncrement;\" json:\"%s\" form:\"%s\" comment:\"%s\"`\n", utils.ToCamelCase(col.COLUMN_NAME), sqlToGoType(col.DATA_TYPE, col.COLUMN_NAME), col.COLUMN_NAME, col.COLUMN_NAME, col.COLUMN_NAME, col.COLUMN_COMMENT))
		} else if col.COLUMN_NAME == "deleted_at" || col.COLUMN_NAME == "delete_time" {
			buffer.WriteString(fmt.Sprintf("    %s %s `gorm:\"column:%s\" json:\"%s,omitempty\" form:\"%s\" comment:\"%s\"`\n", utils.ToCamelCase(col.COLUMN_NAME), sqlToGoType(col.DATA_TYPE, col.COLUMN_NAME), col.COLUMN_NAME, col.COLUMN_NAME, col.COLUMN_NAME, col.COLUMN_COMMENT))
		} else {
			buffer.WriteString(fmt.Sprintf("    %s %s `gorm:\"column:%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`\n", utils.ToCamelCase(col.COLUMN_NAME), sqlToGoType(col.DATA_TYPE, col.COLUMN_NAME), col.COLUMN_NAME, col.COLUMN_NAME, col.COLUMN_NAME, col.COLUMN_COMMENT))
		}

	}
	buffer.WriteString("}\n\n")
	tableNameStr := fmt.Sprintf("func (%s) TableName() string {\n\treturn \"%s.%s\"\n}", utils.ToCamelCase(table), database, table)
	buffer.WriteString(tableNameStr)
	result := buffer.String()
	return result
}

// sqlToGoType 数据库类型
func sqlToGoType(sqlType, columnName string) string {
	if columnName == "deleted_at" || columnName == "delete_time" {
		return "gorm.DeletedAt"
	}
	switch sqlType {
	case "int", "tinyint", "smallint", "mediumint", "bigint":
		return "int"
	case "bit":
		return "uint8"
	case "varchar", "char", "text", "mediumtext", "longtext", "enum", "set":
		return "string"
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		return "[]byte"
	case "date", "datetime", "timestamp":
		return "time.Time"
	case "decimal", "float", "double":
		return "float64"
	case "json":
		return "sql.Json"
	default:
		return "interface{}"
	}
}
