package router

import (
	"embed"
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

// Template 定义了代码生成模板的结构体
type Template struct {
	TableName    string                    `json:"table_name"`    // 表名称
	TableComment string                    `json:"table_comment"` // 表注释
	Fields       []template.TableStructure `json:"fields"`        // 表字段
	Package      string                    `json:"package"`       // 包名
	IsCreate     bool                      `json:"is_create"`     // 是否创建
	DataBase     string                    `json:"data_base"`     // 数据库
	ModuleType   string                    `json:"module_type"`   // 模块类型
	IsWeb        bool                      `json:"is_web"`        // 是否生成Web代码
	WebPackage   string                    `json:"web_package"`   // Web包名
}

// Router 初始化并返回一个Gin引擎实例
// Initialize and return a Gin engine instance
func Router() *gin.Engine {
	// 开发者模式
	// Developer mode
	if !config.GetBool("system.debug") {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	app := gin.New()
	app.Use(gin_middleware.Recovery())

	// 跨域处理
	// CORS handling
	if config.GetBool("system.cors") {
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{"*"}
		corsConfig.AllowHeaders = []string{"*"}
		app.Use(cors.New(corsConfig))
	}

	return app
}

// Load 加载路由并返回Gin引擎实例
// Load routes and return a Gin engine instance
func Load(embeddedFiles embed.FS) *gin.Engine {
	app := Router()

	// 设置静态资源路径
	// Set static resource path
	staticFiles, _ := fs.Sub(embeddedFiles, "web/dist")
	app.StaticFS("/web", http.FS(staticFiles))

	// 处理未找到的路由
	// Handle not found routes
	app.NoRoute(func(c *gin.Context) {
		data, err := embeddedFiles.ReadFile("web/dist/index.html")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		rawURL := c.Request.URL.String()
		ext := filepath.Ext(rawURL)
		switch ext {
		case ".js":
			c.Header("Content-Type", "application/javascript")
		case ".wasm":
			c.Header("Content-Type", "application/wasm")
		case ".css":
			c.Header("Content-Type", "text/css")
		case ".svg":
			c.Header("Content-Type", "image/svg+xml")
		case ".html", "/":
			c.Header("Content-Type", "text/html; charset=utf-8")
		case ".lani", ".lmat", ".lm", ".lh", ".ls":
			c.Header("Content-Type", "application/octet-stream")
		}
		c.String(http.StatusOK, string(data))
	})

	// 获取数据库列表
	// Get database list
	app.GET("api/database", func(c *gin.Context) {
		var result []string
		dbType := ant.GetConfig("connections.0.type")
		switch dbType {
		case "pgsql":
			ant.Db().Raw("SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT LIKE 'pg_%' AND schema_name <> 'information_schema';").Find(&result)
		case "mysql":
			ant.Db().Raw("SHOW DATABASES;").Find(&result)
		}
		c.JSON(http.StatusOK, result)
	})

	// 获取表列表
	// Get table list
	app.GET("api/table-list", func(c *gin.Context) {
		schema := c.Query("table")
		var list []map[string]interface{}
		dbType := ant.GetConfig("connections.0.type")
		switch dbType {
		case "pgsql":
			ant.Db().Raw("SELECT table_name, table_schema, obj_description((table_schema || '.' || table_name)::regclass, 'pg_class') AS table_comment FROM information_schema.tables WHERE table_schema = ? AND table_type = 'BASE TABLE';", schema).Find(&list)
		case "mysql":
			ant.Db().Raw("SELECT TABLE_NAME AS table_name, TABLE_ROWS AS table_rows, TABLE_COLLATION AS table_collation, TABLE_COMMENT AS table_comment FROM INFORMATION_SCHEMA.Tables WHERE table_schema = ?", schema).Find(&list)
		}
		c.JSON(http.StatusOK, list)
	})

	// 获取表字段
	// Get table fields
	app.GET("api/table-field", func(c *gin.Context) {
		db := c.Query("db")
		table := c.Query("table")
		var list []map[string]interface{}
		var sql string
		dbType := ant.GetConfig("connections.0.type")
		switch dbType {
		case "pgsql":
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
					cols.table_schema = ?  
					AND cols.table_name = ? 
					ORDER BY 
					cols.ordinal_position;`
		case "mysql":
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
		c.JSON(http.StatusOK, list)
	})

	// 预览代码
	// Preview code
	app.POST("api/code", func(c *gin.Context) {
		var code Template
		if err := c.BindJSON(&code); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
			return
		}

		modelStr := template.GenGormModel(code.DataBase, code.TableName, code.Fields)
		daoStr := template.GenDao(code.TableName, code.Fields)
		serviceStr := template.GenService(code.TableName, code.Fields)
		controllerStr := template.GenController(code.TableName, code.TableComment, code.Package)
		routeStr := template.GenRoute(code.TableName)
		requestStr := template.GenRequest(code.TableName, code.Fields)

		if code.IsCreate {
			if utils.Exists("./app/dao/"+code.TableName+".go") || utils.Exists("./app/service/"+code.TableName+".go") || utils.Exists("./app/entity/request/"+code.TableName+".go") || utils.Exists("./app/http/"+code.Package+"/"+code.TableName+".go") || utils.Exists("./routes/"+code.TableName+".go") {
				c.JSON(http.StatusConflict, gin.H{"message": "The file already exists"})
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
		if code.IsWeb {
			webStr, err := web.Views(code.TableName)
			if err != nil {
				c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
				return
			}
			result = append(result, map[string]interface{}{"name": "Web", "code": webStr})

			apiStr, err := web.Api(code.TableName)
			if err != nil {
				c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
				return
			}
			result = append(result, map[string]interface{}{"name": "Api", "code": apiStr})
		}
		c.JSON(http.StatusOK, result)
	})

	// 生成单个代码模块
	// Generate a single code module
	app.POST("api/generate_code", func(c *gin.Context) {
		var code Template
		if err := c.BindJSON(&code); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
			return
		}

		moduleConfig := map[string]struct {
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid module type"})
			return
		}

		// 检查文件是否已存在
		// Check if the file already exists
		if utils.Exists(config.Path) {
			c.JSON(http.StatusConflict, gin.H{"message": "The file already exists"})
			return
		}

		// 生成文件
		// Generate file
		fileContent := config.Generate()
		if err := utils.WriteFile(config.Path, fileContent); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	return app
}
