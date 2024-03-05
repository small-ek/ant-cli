package cmd

import (
	"errors"
	"flag"
	"fmt"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/urfave/cli/v2"
	"strings"
)

type GenDao struct {
}

type TableStructure struct {
	COLUMN_NAME    string `json:"COLUMN_NAME"`
	DATA_TYPE      string `json:"DATA_TYPE"`
	COLUMN_COMMENT string `json:"COLUMN_COMMENT"`
	COLUMN_TYPE    string `json:"COLUMN_TYPE"`
	COLUMN_KEY     string `json:"COLUMN_KEY"`
}

// Action
func (b GenDao) Action(c *cli.Context) error {
	tableName := c.Args().First()
	if len(tableName) == 0 {
		return errors.New("Please enter the database alias and table name")
	}
	configPath := flag.String("config", "./config/config.toml", "Configuration file path")
	ant.New(*configPath)
	tableStr := strings.Split(tableName, ".")

	if len(tableStr) < 2 {
		return errors.New("Please enter the database alias and table name")
	}

	var tableStructure []TableStructure
	var sql = `SELECT 
				COLUMN_NAME,
				DATA_TYPE,
				COLUMN_COMMENT,
					COLUMN_TYPE,
					COLUMN_KEY
				FROM 
				INFORMATION_SCHEMA.COLUMNS 
				WHERE 
				TABLE_SCHEMA = ?
				AND 
				TABLE_NAME = ?;`

	ant.Db().Raw(sql, tableStr[0], tableStr[1]).Find(&tableStructure)

	fmt.Println("type Admin struct {")
	for _, col := range tableStructure {
		fmt.Printf("    %s %s `json:\"%s\" form:\"%s\" comment:\"%s\"`\n", toCamelCase(col.COLUMN_NAME), sqlToGoType(col.DATA_TYPE), col.COLUMN_NAME, col.COLUMN_NAME, col.COLUMN_COMMENT)
	}
	fmt.Println("}")
	return nil
}

// toCamelCase 驼峰
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

// sqlToGoType 数据库类型
func sqlToGoType(sqlType string) string {
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
