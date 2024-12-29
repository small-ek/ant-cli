package router

import (
	"embed"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/ant-cli/template"
	"github.com/small-ek/ant-cli/template/web"
	"github.com/small-ek/ant-cli/utils"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/frame/gin_middleware"
	"github.com/small-ek/antgo/os/config"
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
	ModuleType   string                    `json:"module_type"`   // 模块类型
	IsWeb        bool                      `json:"is_web"`        // 是否web
	WebPackage   string                    `json:"web_package"`   // web包名
}

func Router() *gin.Engine {
	//开发者模式
	if config.GetBool("system.debug") == false {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	var app = gin.New()
	app.Use(gin_middleware.Recovery())
	//跨域处理
	if config.GetBool("system.cors") == true {
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{"*"}
		corsConfig.AllowHeaders = []string{"*"}
		app.Use(cors.New(corsConfig))
	}

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
		if ant.GetConfig("connections.0.type") == "pgsql" {
			ant.Db().Raw("SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT LIKE 'pg_%' AND schema_name <> 'information_schema';").Find(&result)
		}
		if ant.GetConfig("connections.0.type") == "mysql" {
			ant.Db().Raw("SHOW DATABASES;").Find(&result)
		}
		c.JSON(200, result)
	})

	//获取表列表
	app.GET("api/table-list", func(c *gin.Context) {
		var table = c.Query("table")
		var list = []map[string]interface{}{}
		if ant.GetConfig("connections.0.type") == "pgsql" {
			ant.Db().Raw("SELECT table_name,table_schema,obj_description((table_schema || '.' || table_name)::regclass, 'pg_class') AS table_comment FROM information_schema.tables WHERE table_schema = ? AND table_type = 'BASE TABLE';", table).Find(&list)
		}
		if ant.GetConfig("connections.0.type") == "mysql" {
			ant.Db().Raw("SELECT TABLE_NAME AS table_name,TABLE_ROWS AS table_rows,TABLE_COLLATION AS table_collation,TABLE_COMMENT AS table_comment FROM INFORMATION_SCHEMA.Tables WHERE table_schema = ?", table).Find(&list)
		}
		c.JSON(200, list)
	})

	//获取表字段
	app.GET("api/table-field", func(c *gin.Context) {
		var db = c.Query("db")
		var table = c.Query("table")
		var list = []map[string]interface{}{}
		var sql string
		if ant.GetConfig("connections.0.type") == "pgsql" {
			sql = `SELECT 
					cols.column_name AS field_name,
					cols.data_type AS field_type,
					pgd.description AS comment,
					CASE 
						WHEN EXISTS (
							SELECT 1 
							FROM pg_index i
							JOIN pg_attribute a ON a.attnum = ANY(i.indkey) 
							WHERE i.indrelid = (cols.table_schema || '.' || cols.table_name)::regclass
							  AND a.attname = cols.column_name
							  AND i.indisprimary
						) THEN 'PRIMARY KEY'
						WHEN EXISTS (
							SELECT 1 
							FROM pg_index i
							JOIN pg_attribute a ON a.attnum = ANY(i.indkey) 
							WHERE i.indrelid = (cols.table_schema || '.' || cols.table_name)::regclass
							  AND a.attname = cols.column_name
							  AND i.indisunique
						) THEN 'UNIQUE'
						ELSE NULL
					END AS indexes
					FROM 
					information_schema.columns cols
					LEFT JOIN 
					pg_description pgd 
					ON (pgd.objoid = (cols.table_schema || '.' || cols.table_name)::regclass 
						AND pgd.objsubid = cols.ordinal_position)
					WHERE 
					cols.table_schema = ?  -- 替换为你的 schema
					AND cols.table_name = ? -- 替换为你的表名
					ORDER BY 
					cols.ordinal_position;`
		}
		if ant.GetConfig("connections.0.type") == "mysql" {
			sql = `SELECT 
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
		}
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
		controllerStr := template.GenController(code.TableName, code.TableComment, code.Package)
		routeStr := template.GenRoute(code.TableName)
		requestStr := template.GenRequest(code.TableName, code.Fields)

		if code.IsCreate == true {
			if utils.Exists("./app/dao/"+code.TableName+".go") || utils.Exists("./app/service/"+code.TableName+".go") || utils.Exists("./app/entity/request/"+code.TableName+".go") || utils.Exists("./app/http/"+code.Package+"/"+code.TableName+".go") || utils.Exists("./routes/"+code.TableName+".go") {
				c.JSON(409, gin.H{"message": "The file already exists"})
				return
			}
			utils.WriteFile("./app/entity/models/"+code.TableName+".go", modelStr)
			utils.WriteFile("./app/dao/"+code.TableName+".go", daoStr)
			utils.WriteFile("./app/service/"+code.TableName+".go", serviceStr)
			utils.WriteFile("./app/entity/request/"+code.TableName+".go", requestStr)
			utils.WriteFile("./app/http/"+code.Package+"/"+code.TableName+".go", controllerStr)
			utils.WriteFile("./routes/"+code.TableName+".go", routeStr)
		}
		result := []map[string]interface{}{
			{"name": "Route", "code": routeStr},
			{"name": "Controller", "code": controllerStr},
			{"name": "Request", "code": requestStr},
			{"name": "Service", "code": serviceStr},
			{"name": "Dao", "code": daoStr},
			{"name": "Model", "code": modelStr},
		}
		if code.IsWeb == true {
			webStr, err := web.Views(code.TableName)
			if err != nil {
				c.JSON(409, gin.H{"message": err.Error()})
				return
			}
			result = append(result, map[string]interface{}{"name": "Web", "code": webStr})

			apiStr, err := web.Api(code.TableName)
			if err != nil {
				c.JSON(409, gin.H{"message": err.Error()})
				return
			}
			result = append(result, map[string]interface{}{"name": "Api", "code": apiStr})
		}
		c.JSON(200, result)
	})

	//生成单个代码
	app.POST("api/generate_code", func(c *gin.Context) {
		var code Template
		if err := c.BindJSON(&code); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload", "details": err.Error()})
			return
		}

		var moduleConfig = map[string]struct {
			Path     string
			Generate func() string
		}{
			"Dao": {
				Path:     "./app/dao/" + code.TableName + ".go",
				Generate: func() string { return template.GenDao(code.TableName, code.Fields) },
			},
			"Model": {
				Path:     "./app/entity/models/" + code.TableName + ".go",
				Generate: func() string { return template.GenGormModel(code.DataBase, code.TableName, code.Fields) },
			},
			"Service": {
				Path:     "./app/service/" + code.TableName + ".go",
				Generate: func() string { return template.GenService(code.TableName, code.Fields) },
			},
			"Request": {
				Path:     "./app/entity/request/" + code.TableName + ".go",
				Generate: func() string { return template.GenRequest(code.TableName, code.Fields) },
			},
			"Controller": {
				Path:     "./app/http/" + code.Package + "/" + code.TableName + ".go",
				Generate: func() string { return template.GenController(code.TableName, code.TableComment, code.Package) },
			},
			"Route": {
				Path:     "./routes/" + code.TableName + ".go",
				Generate: func() string { return template.GenRoute(code.TableName) },
			},
		}

		config, exists := moduleConfig[code.ModuleType]
		if !exists {
			c.JSON(400, gin.H{"error": "Invalid module type"})
			return
		}

		// 检查文件是否已存在
		if utils.Exists(config.Path) {
			c.JSON(409, gin.H{"message": "The file already exists"})
			return
		}

		// 生成文件
		fileContent := config.Generate()
		if err := utils.WriteFile(config.Path, fileContent); err != nil {
			c.JSON(500, gin.H{"error": "Failed to write file", "details": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "success"})
	})

	return app
}
