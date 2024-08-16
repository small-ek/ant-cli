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

type GenApi struct {
}

// Action
func (b GenApi) Action(c *cli.Context) error {
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
	var getTable = map[string]interface{}{}
	ant.Db().Raw("SELECT TABLE_NAME AS table_name,TABLE_ROWS AS table_rows,TABLE_COLLATION AS table_collation,TABLE_COMMENT AS table_comment FROM INFORMATION_SCHEMA.Tables WHERE table_schema = ? AND table_name=?", tableStr[0], tableStr[1]).Find(&getTable)
	// 生成Model
	getModelStr := template.GenGormModel(tableStr[0], tableStr[1], tableStructure)
	utils.WriteFile("./app/entity/models/"+tableStr[1]+".go", getModelStr)
	// 生成Dao
	getDaoStr := template.GenDao(tableStr[1], tableStructure)
	utils.WriteFile("./app/dao/"+tableStr[1]+".go", getDaoStr)
	// 生成Service
	getServiceStr := template.GenService(tableStr[1], tableStructure)
	utils.WriteFile("./app/service/"+tableStr[1]+".go", getServiceStr)
	// 生成request
	getRequestStr := template.GenRequest(tableStr[1], tableStructure)
	utils.WriteFile("./app/entity/request/"+tableStr[1]+".go", getRequestStr)
	// 生成Controller
	getControllerStr := template.GenController(tableStr[1], getTable["table_comment"].(string), "api")
	utils.WriteFile("./app/http/api/"+tableStr[1]+".go", getControllerStr)
	// 生成Route
	getRouteStr := template.GenRoute(tableStr[1])
	utils.WriteFile("./routes/"+tableStr[1]+".go", getRouteStr)
	return nil
}
