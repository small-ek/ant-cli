package template

import (
	"fmt"
)

func Base(name string) string {
	return fmt.Sprintf("package vo\n\nimport (\n\t\"github.com/gin-gonic/gin\"\n\t\"github.com/small-ek/antgo/os/alog\"\n\t\"github.com/small-ek/antgo/os/config\"\n\t\"github.com/small-ek/antgo/utils/response\"\n\t\"go.uber.org/zap\"\n)\n\ntype Base struct {\n\tCode int `json:\"code\"` //Status Code\n}\n\n// Success 成功返回\nfunc (b *Base) Success(c *gin.Context, code string, data ...interface{}) {\n\tmsg := GetMessage(code)\n\tc.SecureJSON(200, response.Success(code, msg, data...))\n}\n\n// Fail 错误返回\nfunc (b *Base) Fail(c *gin.Context, code string, err ...error) {\n\tif len(err) > 0 && err[0] != nil {\n\t\talog.Write.Error(\"Return error\", zap.Errors(\"Fail Error\", err))\n\t}\n\n\tmsg := GetMessage(code)\n\n\tif config.GetBool(\"system.debug\") == true && len(err) > 0 && err[0] != nil {\n\t\tc.SecureJSON(200, response.Fail(code, msg, err[0].Error()))\n\t\treturn\n\t}\n\tc.SecureJSON(200, response.Fail(code, msg))\n\n}\n\n// SetCode 修改状态\nfunc (b *Base) SetCode(code int) *Base {\n\tb.Code = code\n\treturn b\n}\n\n// Page 分页数据\nfunc (b *Base) Page(total int64, list interface{}) response.Page {\n\treturn response.Page{\n\t\tTotal: total,\n\t\tItems: list,\n\t}\n}\n\n// GetDeviceId 获取设备号\nfunc (b *Base) GetDeviceId(c *gin.Context) string {\n\treturn c.GetHeader(\"device-id\")\n}\n\n// GetIp 获取客户端真实IP\nfunc (b *Base) GetIp(c *gin.Context) string {\n\tip := c.GetHeader(\"X-Forwarded-For\")\n\tif ip == \"\" {\n\t\tip = c.ClientIP()\n\t}\n\treturn ip\n}\n")
	//package vo

	//import (
	//	"github.com/gin-gonic/gin"
	//"github.com/small-ek/antgo/os/alog"
	//"github.com/small-ek/antgo/os/config"
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
	//code := 0
	//if b.Code != 0 {
	//code = b.Code
	//}
	//
	//c.SecureJSON(200, response.Success(msg, code, data...))
	//}
	//
	//// Fail 错误返回
	//func (b *Base) Fail(c *gin.Context, msg string, err ...error) {
	//if len(err) > 0 {
	//alog.Write.Error("Return error", zap.Errors("Fail Error", err))
	//}
	//code := 1
	//if b.Code != 0 {
	//code = b.Code
	//}
	//if config.GetBool("system.debug") == true {
	//c.SecureJSON(200, response.Fail(msg, code, err[0].Error()))
	//return
	//}
	//c.SecureJSON(200, response.Fail(msg, code))
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
	//
	//
	//// GetDeviceId 获取设备号
	//func (b *Base) GetDeviceId(c *gin.Context) string {
	//return c.GetHeader("device-id")
	//}

}
