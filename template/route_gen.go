package template

import "github.com/small-ek/ant-cli/utils"

func GenRoute(table string) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	return `package router

import (
	"github.com/gin-gonic/gin"
	"` + getFileName + `/app/http/api"
)

func ` + humpTable + `Route(Router *gin.RouterGroup) {
	` + humpTable + `Controller := api.New` + humpTable + `Controller()
	v1 := Router.Group("/v1/` + table + `")
	{
		v1.GET("", ` + humpTable + `Controller.Index)
		v1.GET("/:id", ` + humpTable + `Controller.Show)
		v1.DELETE("/:id", ` + humpTable + `Controller.Delete)
		v1.POST("", ` + humpTable + `Controller.Create)
		v1.PUT("/:id", ` + humpTable + `Controller.Update)
	}
}

`
}
