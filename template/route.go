package template

func Router(name string) string {
	return `package router

import (
	"` + name + `/router"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/gin_cors"
)

func Router() *gin.Engine {
	var app = gin.New()
	app.Use(gin.Logger())
	//跨域处理
	if config.GetBool("system.cors") == true {
		app.Use(gin_cors.Cors)
	}

	//开发者模式
	if config.GetBool("system.debug") == false {
		gin.SetMode(gin.ReleaseMode)
	}

	return app
}

// Load 加载路由
func Load() *gin.Engine {
	app := Router()
	//添加路由组前缀
	Group := app.Group("")
	//注册路由
	router.IndexRoute(Group)

	return app
}
`
}
