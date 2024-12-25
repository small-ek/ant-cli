package template

import (
	"fmt"
	"github.com/small-ek/ant-cli/utils"
)

// GenService 生成服务端
func GenRequest(table string, tableStructure []TableStructure) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	requestStr := ""
	for _, col := range tableStructure {
		if col.Required == 1 {

			requestStr += fmt.Sprintf("    %s %s `json:\"%s\" form:\"%s\" %s comment:\"%s\"`%s\n",
				utils.ToCamelCase(col.FieldName),
				sqlToGoType(col.FieldType, col.FieldName),
				col.FieldName,
				col.FieldName,
				utils.GetTag(col.Required),
				col.Comment,
				utils.GetComment(col.Comment))
		}

	}
	return `package request

import (
	"github.com/small-ek/antgo/utils/page"
	"` + getFileName + `/app/entity/models"
)

type ` + humpTable + `Request struct {
	models.` + humpTable + `
	page.PageParam
}

type ` + humpTable + `RequestForm struct {
	models.` + humpTable + `
    ` + requestStr + `
}
`
}
