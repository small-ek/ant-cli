package template

func RouterIndex(name string) string {
	return `package router

import (
	"` + name + `/app/http/index"
	"github.com/gin-gonic/gin"
)

func IndexRoute(Router *gin.RouterGroup) {
	Router.GET("/", new(index.IndexController).Index)
}

`
}
