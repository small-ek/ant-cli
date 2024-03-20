package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/frame/middleware"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/gin_cors"
	"io/fs"
	"io/ioutil"
	"net/http"
)

func Router() *gin.Engine {
	//开发者模式
	if config.GetBool("system.debug") == false {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	var app = gin.New()
	app.Use(gin.Logger()).Use(middleware.Recovery())
	//跨域处理
	if config.GetBool("system.cors") == true {
		app.Use(gin_cors.Cors)
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
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	app.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})

	return app
}
