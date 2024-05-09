package template

import (
	"bytes"
	"fmt"
	"github.com/small-ek/ant-cli/utils"
)

type TableStructure struct {
	Comment   string `json:"comment"`    // 字段注释
	FieldName string `json:"field_name"` // 字段名称
	FieldType string `json:"field_type"` // 字段类型
	Required  int    `json:"required"`   // 是否必填
	IsSearch  int    `json:"is_search"`  // 是否搜索
	Indexes   string `json:"indexes"`    // 索引类型
	JoinTable string `json:"join_table"` // 关联表
	JoinField string `json:"join_field"` // 关联字段
	JoinType  string `json:"join_type"`  // 关联类型
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
		if col.FieldName == "deleted_at" || col.FieldName == "delete_time" {
			isImportDeletedAt = true
		}
		if col.FieldName != "deleted_at" && col.FieldName != "delete_time" {
			switch col.FieldType {
			case "date", "datetime", "timestamp":
				isImportDate = true
			case "json":
				isImportJson = true
			}
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
		if col.FieldName == "id" {
			buffer.WriteString(fmt.Sprintf("    %s %s `gorm:\"column:%s;primaryKey;autoIncrement;\" uri:\"%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`\n", utils.ToCamelCase(col.FieldName), sqlToGoType(col.FieldType, col.FieldName), col.FieldName, col.FieldName, col.FieldName, col.FieldName, col.Comment))
		} else if col.FieldType == "join" { //关联表
			//一对一
			if col.JoinType == "oneToOne" {
				buffer.WriteString(fmt.Sprintf("    %s %s `gorm:\"column:%s;foreignKey:%s;references:%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`\n", utils.ToCamelCase(col.FieldName), utils.ToCamelCase(col.JoinTable), col.FieldName, col.JoinField, col.FieldName, col.FieldName, col.FieldName, col.Comment))
			}
			//一对多
			if col.JoinType == "oneToMany" {
				buffer.WriteString(fmt.Sprintf("    %s []%s `gorm:\"column:%sforeignKey:%s;references:%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`\n", utils.ToCamelCase(col.FieldName), utils.ToCamelCase(col.JoinTable), col.FieldName, col.JoinField, col.FieldName, col.FieldName, col.FieldName, col.Comment))
			}
			//多对多
			if col.JoinType == "manyToMany" {
				buffer.WriteString(fmt.Sprintf("    %s []%s `gorm:\"column:%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`\n", utils.ToCamelCase(col.FieldName), utils.ToCamelCase(col.JoinTable), col.FieldName, col.FieldName, col.FieldName, col.Comment))
			}
		} else {
			buffer.WriteString(fmt.Sprintf("    %s %s `gorm:\"column:%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`\n", utils.ToCamelCase(col.FieldName), sqlToGoType(col.FieldType, col.FieldName), col.FieldName, col.FieldName, col.FieldName, col.Comment))
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
