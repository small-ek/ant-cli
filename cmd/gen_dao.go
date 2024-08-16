package cmd

import (
	"errors"
	"flag"
	"github.com/small-ek/ant-cli/template"
	"github.com/small-ek/ant-cli/utils"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/urfave/cli/v2"
	"strings"
)

type GenDao struct {
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

	var tableStructure []template.TableStructure
	var sql = `SELECT 
				COLUMN_NAME AS field_name,
				DATA_TYPE AS field_type,
				COLUMN_COMMENT AS comment,
				COLUMN_KEY AS indexes
				FROM 
				INFORMATION_SCHEMA.COLUMNS 
				WHERE 
				TABLE_SCHEMA = ?
				AND 
				TABLE_NAME = ?
                ORDER BY ORDINAL_POSITION;`

	ant.Db().Raw(sql, tableStr[0], tableStr[1]).Find(&tableStructure)
	if len(tableStructure) == 0 {
		return errors.New("Database or data table does not exist")
	}
	//生成Model
	getModelStr := template.GenGormModel(tableStr[0], tableStr[1], tableStructure)
	utils.WriteFile("./app/entity/models/"+tableStr[1]+".go", getModelStr)
	//生成Dao
	getDaoStr := template.GenDao(tableStr[1], tableStructure)
	utils.WriteFile("./app/dao/"+tableStr[1]+".go", getDaoStr)

	return nil
}
