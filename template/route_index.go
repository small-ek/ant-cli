package template

func RouterIndex(name string) string {
	return `package routes

import (
	"` + name + `/app/http/api"
	"github.com/gin-gonic/gin"
)

func IndexRoute(route *gin.RouterGroup) {
	IndexController := new(api.IndexController)
	route.GET("/", IndexController.Index)
}

`
}
