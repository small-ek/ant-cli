package template

func RouterIndex(name string) string {
	return `package router

import (
	"` + name + `/app/http/api"
	"github.com/gin-gonic/gin"
)

func IndexRoute(Router *gin.RouterGroup) {
	IndexController := new(api.IndexController)
	Router.GET("/", IndexController.Index)
}

`
}
