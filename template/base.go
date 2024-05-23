package template

func Base(name string) string {
	return `package http

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/alog"
	"github.com/small-ek/antgo/utils/response"
	"go.uber.org/zap"
)

type Base struct {
}

// Success 成功返回
func (b *Base) Success(c *gin.Context, msg string, data ...interface{}) {
	c.SecureJSON(200, response.Success(msg, data...))
}

// Fail 错误返回
func (b *Base) Fail(c *gin.Context, msg string, err ...string) {
	if len(err) > 0 {
		alog.Write.Debug("Return error", zap.Any("error", err))
	}

	c.SecureJSON(200, response.Fail(msg, err...))
}

// Page 分页数据
func (b *Base) Page(total int64, list interface{}) response.Page {
	return response.Page{
		Total: total,
		List:  list,
	}
}
`
}
