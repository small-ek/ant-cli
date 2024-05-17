package router

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/ant-cli/template"
	"github.com/small-ek/ant-cli/utils"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/frame/middleware"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/gin_cors"
	"io/fs"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type Template struct {
	TableName    string                    `json:"table_name"`    // 表名称
	TableComment string                    `json:"table_comment"` // 表注释
	Fields       []template.TableStructure `json:"fields"`        // 表字段
	Package      string                    `json:"package"`       // 包名
	IsCreate     bool                      `json:"is_create"`     // 是否创建
	DataBase     string                    `json:"data_base"`     // 数据库
}

func Router() *gin.Engine {
	//开发者模式
	if config.GetBool("system.debug") == false {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	var app = gin.New()
	app.Use(gin.Logger()).Use(middleware.Recovery())
	//跨域处理
	app.Use(gin_cors.Cors)

	return app
}

// Load 加载路由
func Load(f embed.FS) *gin.Engine {
	app := Router()
	//添加路由组前缀
	//Group := app.Group("")
	//注册路由

	st, _ := fs.Sub(f, "web/dist")
	//设置资源路径
	app.StaticFS("/web", http.FS(st))

	//找不到默认跳转这里
	app.NoRoute(func(c *gin.Context) {
		data, err := f.ReadFile("web/dist/index.html")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		rawURL := c.Request.URL.String()
		fmt.Println(rawURL)
		ext := filepath.Ext(rawURL)
		if ext == ".js" {
			c.Header("Content-Type", "application/javascript")
		}
		if ext == ".wasm" {
			c.Header("Content-Type", "application/wasm")
		}
		if ext == ".css" {
			c.Header("Content-Type", "text/css")
		}
		if ext == ".svg" {
			c.Header("Content-Type", "image/svg+xml")
		}
		if ext == ".html" || rawURL == "/" {
			c.Header("Content-Type", "text/html; charset=utf-8")
		}
		if ext == ".lani" || ext == ".lmat" || ext == ".lm" || ext == ".lh" || ext == ".ls" {
			c.Header("Content-Type", "application/octet-stream")
		}
		c.String(http.StatusOK, string(data))
	})

	//获取数据库
	app.GET("api/database", func(c *gin.Context) {
		var result = []string{}
		ant.Db().Raw("SHOW DATABASES;").Find(&result)
		c.JSON(200, result)
	})

	//获取表列表
	app.GET("api/table-list", func(c *gin.Context) {
		var table = c.Query("table")
		var list = []map[string]interface{}{}
		ant.Db().Raw("SELECT TABLE_NAME AS table_name,TABLE_ROWS AS table_rows,TABLE_COLLATION AS table_collation,TABLE_COMMENT AS table_comment FROM INFORMATION_SCHEMA.Tables WHERE table_schema = ?", table).Find(&list)
		c.JSON(200, list)
	})

	//获取表字段
	app.GET("api/table-field", func(c *gin.Context) {
		var db = c.Query("db")
		var table = c.Query("table")
		var list = []map[string]interface{}{}
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
		ant.Db().Raw(sql, db, table).Find(&list)
		c.JSON(200, list)
	})

	//预览代码
	app.POST("api/code", func(c *gin.Context) {
		var code Template
		err := c.BindJSON(&code)
		if err != nil {
			return
		}

		modelStr := template.GenGormModel(code.DataBase, code.TableName, code.Fields)
		daoStr := template.GenDao(code.TableName, code.Fields)
		serviceStr := template.GenService(code.TableName, code.Fields)
		controllerStr := template.GenController(code.TableName, code.TableComment)
		routeStr := template.GenRoute(code.TableName)
		requestStr := template.GenRequest(code.TableName, code.Fields)

		if code.IsCreate == true {
			utils.WriteFile("./app/model/"+code.TableName+".go", modelStr)
			utils.WriteFile("./app/dao/"+code.TableName+".go", daoStr)
			utils.WriteFile("./app/service/"+code.TableName+".go", serviceStr)
			utils.WriteFile("./app/request/"+code.TableName+".go", requestStr)
			utils.WriteFile("./app/http/index/"+code.TableName+".go", controllerStr)
			utils.WriteFile("./router/"+code.TableName+".go", routeStr)
		}
		c.JSON(200, []map[string]interface{}{
			{"name": "Route", "code": routeStr},
			{"name": "Controller", "code": controllerStr},
			{"name": "Request", "code": requestStr},
			{"name": "Service", "code": serviceStr},
			{"name": "Dao", "code": daoStr},
			{"name": "Model", "code": modelStr},
		})
	})

	return app
}
