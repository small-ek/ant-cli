package template

func Index(name string) string {
	return `package api

import (
	"` + name + `/app/http"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	http.Base
}

// Index
func (index *IndexController) Index(c *gin.Context) {
	c.String(200, "Hello AntGo")
}

`
}
