package template

func Router(name string) string {
	return `package router

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/frame/middleware"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/gin_cors"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"io/ioutil"
	_ "` + name + `/docs"
	"` + name + `/router"
)

func Router() *gin.Engine {
	var app = gin.New()
	//开发者模式
	if config.GetBool("system.debug") == false {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	} else {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

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
	Group := app.Group("")
	//注册路由
	router.IndexRoute(Group)

	return app
}
`
}
