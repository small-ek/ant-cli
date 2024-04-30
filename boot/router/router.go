package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/frame/middleware"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/gin_cors"
	"io/fs"
	"io/ioutil"
	"net/http"
)

type Template struct {
	TableName string                   `json:"table_name"`
	Fields    []map[string]interface{} `json:"fields"`
	Package   string                   `json:"package"`
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
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	//获取数据库
	app.GET("api/get_database", func(c *gin.Context) {
		var result = []string{}
		ant.Db().Raw("SHOW DATABASES;").Find(&result)
		c.JSON(200, result)
	})

	//获取表列表
	app.GET("api/get_table_list", func(c *gin.Context) {
		var table = c.Query("table")
		var list = []map[string]interface{}{}
		ant.Db().Raw("SELECT TABLE_NAME AS table_name,TABLE_ROWS AS table_rows,TABLE_COLLATION AS table_collation,TABLE_COMMENT AS table_comment FROM INFORMATION_SCHEMA.Tables WHERE table_schema = ?", table).Find(&list)
		c.JSON(200, list)
	})

	//获取表字段
	app.GET("api/get_table", func(c *gin.Context) {
		var db = c.Query("db")
		var table = c.Query("table")
		var list = []map[string]interface{}{}
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
				TABLE_NAME = ?
                ORDER BY ORDINAL_POSITION;`
		ant.Db().Raw(sql, db, table).Find(&list)
		c.JSON(200, list)
	})

	//预览代码
	app.GET("api/previewCode", func(c *gin.Context) {

		c.JSON(200, gin.H{})
	})

	return app
}
