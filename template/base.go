package template

import (
	"fmt"
)

func Base(name string) string {
	return fmt.Sprintf("package vo\n\nimport (\n\t\"github.com/gin-gonic/gin\"\n\t\"github.com/small-ek/antgo/os/alog\"\n\t\"github.com/small-ek/antgo/utils/response\"\n\t\"go.uber.org/zap\"\n)\n\ntype Base struct {\n\tCode int `json:\"code\"` //Status Code\n}\n\n// Success 成功返回\nfunc (b *Base) Success(c *gin.Context, msg string, data ...interface{}) {\n\tcode := 200\n\tif b.Code != 0 {\n\t\tcode = b.Code\n\t}\n\n\tc.SecureJSON(200, response.Success(msg, code, data...))\n}\n\n// Fail 错误返回\nfunc (b *Base) Fail(c *gin.Context, msg string, err ...string) {\n\tif len(err) > 0 {\n\t\talog.Write.Debug(\"Return error\", zap.Any(\"error\", err))\n\t}\n\tcode := 422\n\tif b.Code != 0 {\n\t\tcode = b.Code\n\t}\n\n\tc.SecureJSON(200, response.Fail(msg, code, err...))\n\n}\n\n// SetCode 修改状态\nfunc (b *Base) SetCode(code int) *Base {\n\tb.Code = code\n\treturn b\n}\n\n// Page 分页数据\nfunc (b *Base) Page(total int64, list interface{}) response.Page {\n\treturn response.Page{\n\t\tTotal: total,\n\t\tItems: list,\n\t}\n}\n")
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
	//code := 200
	//if b.Code != 0 {
	//code = b.Code
	//}
	//
	//c.SecureJSON(200, response.Success(msg, code, data...))
	//}
	//
	//// Fail 错误返回
	//func (b *Base) Fail(c *gin.Context, msg string, err ...string) {
	//if len(err) > 0 {
	//alog.Write.Debug("Return error", zap.Any("error", err))
	//}
	//code := 422
	//if b.Code != 0 {
	//code = b.Code
	//}
	//
	//c.SecureJSON(200, response.Fail(msg, code, err...))
	//
	//}
	//
	//// SetCode 修改状态
	//func (b *Base) SetCode(code int) *Base {
	//b.Code = code
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
