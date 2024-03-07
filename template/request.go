package template

import (
	"github.com/small-ek/ant-cli/utils"
)

// GenService 生成服务端
func GenRequest(table string) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	return `package request

import (
	"github.com/small-ek/antgo/utils/page"
	"` + getFileName + `/app/model"
)

type ` + humpTable + `Request struct {
	model.` + humpTable + `
	page.PageParam
}

`
}
