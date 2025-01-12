package template

func Code(name string) string {
	return `
package vo

// Code 业务状态码
const (
	INVALID_REQUEST_PARAMETERS = "INVALID_REQUEST_PARAMETERS" // INVALID_REQUEST_PARAMETERS 无效的请求参数

	SUCCESS = "SUCCESS" // SUCCESS 成功

	FAILED = "FAILED" // FAILED 失败

	CREATION_FAILED = "CREATION_FAILED" // CREATION_FAILED 创建失败

	CREATION_SUCCESS = "CREATION_SUCCESS" // CREATION_SUCCESS 创建成功

	UPDATE_FAILED = "UPDATE_FAILED" // UPDATE_FAILED 更新失败

	UPDATE_SUCCESS = "UPDATE_SUCCESS" // UPDATE_SUCCESS 更新成功

	DELETE_FAILED = "DELETE_FAILED" // DELETE_FAILED 删除失败

	DELETE_SUCCESS = "DELETE_SUCCESS" // DELETE_SUCCESS 删除成功
)

// 全局映射表
var Messages = map[string]string{
	"INVALID_REQUEST_PARAMETERS": "请求参数无效",
	"SUCCESS":                    "操作成功",
	"FAILED":                     "操作失败",

	"CREATION_FAILED":  "创建操作失败",
	"CREATION_SUCCESS": "创建操作成功",

	"UPDATE_FAILED":  "更新操作失败",
	"UPDATE_SUCCESS": "更新操作成功",

	"DELETE_FAILED":  "删除操作失败",
	"DELETE_SUCCESS": "删除操作成功",
}

func GetMessage(code string) string {
	if msg, exists := Messages[code]; exists {
		return msg
	}
	return "未知错误，请联系客服解决。"
}


`

}
