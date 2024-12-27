package template

import (
	"bytes"
	"fmt"
	"github.com/small-ek/ant-cli/utils"
)

type TableStructure struct {
	Comment    string `json:"comment"`    // 字段注释
	FieldName  string `json:"field_name"` // 字段名称
	FieldType  string `json:"field_type"` // 字段类型
	Required   int    `json:"required"`   // 是否必填
	IsSearch   int    `json:"is_search"`  // 是否搜索
	Conditions string `json:"conditions"` // 查询条件
	Indexes    string `json:"indexes"`    // 索引类型
	JoinTable  string `json:"join_table"` // 关联表
	JoinField  string `json:"join_field"` // 关联字段
	JoinType   string `json:"join_type"`  // 关联类型
}

// database 数据库名称
// table 表名称
// tableStructure 表结构
// GenGormModel 生成gorm模型
func GenGormModel(database, table string, tableStructure []TableStructure) string {
	var buffer bytes.Buffer
	packageStr := fmt.Sprintf("package models \n\n")

	isImportDate := false
	isImportJson := false
	isImportDeletedAt := false
	for _, col := range tableStructure {
		if col.FieldName == "deleted_at" || col.FieldName == "delete_time" {
			isImportDeletedAt = true
		}

		if col.FieldName != "deleted_at" && col.FieldName != "delete_time" {
			switch col.FieldType {
			case "date", "datetime", "timestamp", "time":
				isImportDate = true
			case "json":
				isImportJson = true
			}
		}

	}
	buffer.WriteString(packageStr)
	importStr := fmt.Sprintf("import (")
	buffer.WriteString(importStr)

	//是否导入JSON
	if isImportJson == true {
		importSqlStr := fmt.Sprintf(`"github.com/small-ek/antgo/db/adb/asql"`)
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
		if col.FieldName == "created_at" || col.FieldName == "updated_at" || col.FieldName == "create_time" || col.FieldName == "update_time" {
			buffer.WriteString(fmt.Sprintf("    %s %s `gorm:\"column:%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`%s\n",
				utils.ToCamelCase(col.FieldName),
				sqlToGoType(col.FieldType, col.FieldName),
				col.FieldName,
				col.FieldName,
				col.FieldName,
				utils.RemoveNewlines(col.Comment),
				utils.GetComment(col.Comment)))
		} else if col.FieldName == "deleted_at" || col.FieldName == "delete_time" {
			buffer.WriteString(fmt.Sprintf("    %s gorm.DeletedAt `gorm:\"column:%s\" json:\"-\" form:\"%s\" comment:\"%s\"`%s\n",
				utils.ToCamelCase(col.FieldName),
				col.FieldName,
				col.FieldName,
				utils.RemoveNewlines(col.Comment),
				utils.GetComment(col.Comment)))
		} else if col.FieldType == "join" { //关联表
			switch col.JoinType {
			case "oneToOne":
				buffer.WriteString(fmt.Sprintf("    %s %s `gorm:\"foreignKey:%s;references:%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`%s\n",
					utils.ToCamelCase(col.JoinTable),
					utils.ToCamelCase(col.JoinTable),
					utils.ToCamelCase(col.JoinField),
					utils.ToCamelCase(col.FieldName),
					col.JoinTable,
					col.JoinTable,
					utils.RemoveNewlines(col.Comment),
					utils.GetComment(col.Comment)))
			case "oneToMany":
				buffer.WriteString(fmt.Sprintf("    %s []%s `gorm:\"foreignKey:%s;references:%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`%s\n",
					utils.ToCamelCase(col.JoinTable),
					utils.ToCamelCase(col.JoinTable),
					utils.ToCamelCase(col.JoinField),
					utils.ToCamelCase(col.FieldName),
					col.JoinTable,
					col.JoinTable,
					utils.RemoveNewlines(col.Comment),
					utils.GetComment(col.Comment)))
			case "manyToMany":
				buffer.WriteString(fmt.Sprintf("    %s []%s `gorm:\"many2many:%s;foreignKey:%s;References:%s;joinForeignKey:%s;joinReferences:%s\" json:\"%s\" form:\"%s\" comment:\"%s\" `%s\n",
					utils.ToCamelCase(col.JoinTable),
					utils.ToCamelCase(col.JoinTable),
					utils.Many2Many(table, col.JoinTable),
					utils.ToCamelCase(col.FieldName),
					utils.ToCamelCase(col.JoinField),
					table+"_id",
					col.JoinTable+"_id",
					col.JoinTable,
					col.JoinTable,
					utils.RemoveNewlines(col.Comment),
					utils.GetComment(col.Comment)))
			}

		} else if col.FieldName == "id" {
			buffer.WriteString(fmt.Sprintf("    %s %s `gorm:\"column:%s;primaryKey;autoIncrement;\" uri:\"%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`%s\n",
				utils.ToCamelCase(col.FieldName),
				noNullSqlToGoType(col.FieldType, col.FieldName),
				col.FieldName,
				col.FieldName,
				col.FieldName,
				col.FieldName,
				utils.RemoveNewlines(col.Comment),
				utils.GetComment(col.Comment)))
		} else {
			buffer.WriteString(fmt.Sprintf("    %s %s `gorm:\"column:%s%s\" json:\"%s\" form:\"%s\" comment:\"%s\"`%s\n",
				utils.ToCamelCase(col.FieldName),
				sqlToGoType(col.FieldType, col.FieldName),
				col.FieldName,
				sqlDefault(col.FieldType),
				col.FieldName,
				col.FieldName,
				utils.RemoveNewlines(col.Comment),
				utils.GetComment(col.Comment)))
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
	case "int", "tinyint", "smallint", "mediumint", "bigint", "serial", "bigserial", "int2", "int4", "int8", "integer", "smallserial", "serial2", "serial4", "serial8", "bigserial2", "bigserial4", "bigserial8", "oid":
		return "*int"
	case "bit", "varbit":
		return "uint8"
	case "varchar", "char", "text", "mediumtext", "longtext", "set", "character varying", "character", "uuid", "enum":
		return "string"
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob", "bytea":
		return "[]byte"
	case "date", "datetime", "timestamp", "timestamptz", "time", "timetz":
		return "time.Time"
	case "decimal", "float", "double", "numeric", "real", "double precision":
		return "float64"
	case "json", "jsonb":
		return "asql.Json"
	case "boolean":
		return "*bool"
	default:
		return "interface{}"
	}
}

// noNullSqlToGoType 非指针数据库类型
func noNullSqlToGoType(sqlType, columnName string) string {
	if columnName == "deleted_at" || columnName == "delete_time" {
		return "gorm.DeletedAt"
	}
	switch sqlType {
	case "int", "tinyint", "smallint", "mediumint", "bigint", "serial", "bigserial", "int2", "int4", "int8", "integer", "smallserial", "serial2", "serial4", "serial8", "bigserial2", "bigserial4", "bigserial8", "oid":
		return "int"
	case "bit", "varbit":
		return "uint8"
	case "varchar", "char", "text", "mediumtext", "longtext", "set", "character varying", "character", "uuid", "enum":
		return "string"
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob", "bytea":
		return "[]byte"
	case "date", "datetime", "timestamp", "timestamptz", "time", "timetz":
		return "time.Time"
	case "decimal", "float", "double", "numeric", "real", "double precision":
		return "float64"
	case "json", "jsonb":
		return "asql.Json"
	case "boolean":
		return "bool"
	default:
		return "interface{}"
	}
}

// sqlDefault sql默认值
func sqlDefault(sqlType string) string {
	switch sqlType {
	case "enum":
		return ";default:null"
	default:
		return ""
	}
}
