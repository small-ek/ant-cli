package router

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/frame/middleware"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/gin_cors"
	"io/ioutil"
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
func Load() *gin.Engine {
	app := Router()
	//添加路由组前缀
	//Group := app.Group("")
	////注册路由
	app.LoadHTMLGlob("public/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	// 设置一个路由，当访问/时，渲染名为"index"的模板
	app.GET("/index.html", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Hello, Gin!",
		})
	})

	return app
}
