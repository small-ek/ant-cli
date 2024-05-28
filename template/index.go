package template

func Index(name string) string {
	return `package api

import (
	"` + name + `/app/entity/vo"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	vo.Base
}

// Index
func (index *IndexController) Index(c *gin.Context) {
	c.String(200, "Hello AntGo")
}

`
}
