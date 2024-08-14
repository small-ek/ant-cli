package template

import (
	"fmt"
)

func Base(name string) string {
	return fmt.Sprintf("package vo\n\nimport (\n\t\"github.com/gin-gonic/gin\"\n\t\"github.com/small-ek/antgo/os/alog\"\n\t\"github.com/small-ek/antgo/utils/response\"\n\t\"go.uber.org/zap\"\n)\n\ntype Base struct {\n\tCode int `json:\"code\"` //Status Code\n}\n\n// Success 成功返回\nfunc (b *Base) Success(c *gin.Context, msg string, data ...interface{}) {\n\tif b.Code == 0 {\n\t\tb.Code = 200\n\t}\n\tc.SecureJSON(200, response.Success(msg, b.Code, data...))\n}\n\n// Fail 错误返回\nfunc (b *Base) Fail(c *gin.Context, msg string, err ...string) {\n\tif len(err) > 0 {\n\t\talog.Write.Debug(\"Return error\", zap.Any(\"error\", err))\n\t}\n\tif b.Code == 0 {\n\t\tb.Code = 422\n\t}\n\tc.SecureJSON(200, response.Fail(msg, b.Code, err...))\n\n}\n\n// SetStatus 修改状态\nfunc (b *Base) SetStatus(status int) *Base {\n\tb.Code = status\n\treturn b\n}\n\n// Page 分页数据\nfunc (b *Base) Page(total int64, list interface{}) response.Page {\n\treturn response.Page{\n\t\tTotal: total,\n\t\tItems: list,\n\t}\n}\n")
	//package vo
	//
	//import (
	//	"github.com/gin-gonic/gin"
	//"github.com/small-ek/antgo/os/alog"
	//"github.com/small-ek/antgo/utils/response"
	//"go.uber.org/zap"
	//)
	//
	//type Base struct {
	//	Code int `json:"code"` //Status Code
	//}
	//
	//// Success 成功返回
	//func (b *Base) Success(c *gin.Context, msg string, data ...interface{}) {
	//if b.Code == 0 {
	//b.Code = 200
	//}
	//c.SecureJSON(200, response.Success(msg, b.Code, data...))
	//}
	//
	//// Fail 错误返回
	//func (b *Base) Fail(c *gin.Context, msg string, err ...string) {
	//if len(err) > 0 {
	//alog.Write.Debug("Return error", zap.Any("error", err))
	//}
	//if b.Code == 0 {
	//b.Code = 422
	//}
	//c.SecureJSON(200, response.Fail(msg, b.Code, err...))
	//
	//}
	//
	//// SetStatus 修改状态
	//func (b *Base) SetStatus(status int) *Base {
	//b.Code = status
	//return b
	//}
	//
	//// Page 分页数据
	//func (b *Base) Page(total int64, list interface{}) response.Page {
	//return response.Page{
	//Total: total,
	//Items: list,
	//}
	//}

}
