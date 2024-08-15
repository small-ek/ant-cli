package template

import "github.com/small-ek/ant-cli/utils"

func GenRoute(table string) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	smallHumpTable := utils.ToCamelCaseLower(table)
	toKebabCase := utils.ToKebabCase(table)
	return `package routes

import (
	"github.com/gin-gonic/gin"
	"` + getFileName + `/app/http/api"
)

func ` + humpTable + `Route(route *gin.RouterGroup) {
	` + smallHumpTable + `Controller := api.New` + humpTable + `Controller()
	v1 := route.Group("` + toKebabCase + `")
	{
		v1.GET("", ` + smallHumpTable + `Controller.Index)
		v1.GET(":id", ` + smallHumpTable + `Controller.Show)
		v1.DELETE(":id", ` + smallHumpTable + `Controller.Delete)
		v1.POST("", ` + smallHumpTable + `Controller.Create)
		v1.PUT(":id", ` + smallHumpTable + `Controller.Update)
	}
}

`
}
