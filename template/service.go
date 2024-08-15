package template

import (
	"fmt"
	"github.com/small-ek/ant-cli/utils"
)

// GenService 生成服务端
func GenService(table string, tableStructure []TableStructure) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	requestStr := ""
	for _, col := range tableStructure {
		if col.Required == 1 {
			requestStr += fmt.Sprintf("req.%s.%s = req.%s\n",
				utils.ToCamelCase(table),
				utils.ToCamelCase(col.FieldName),
				utils.ToCamelCase(col.FieldName),
			)
		}

	}
	return `package service

import (
	"` + getFileName + `/app/dao"
	"` + getFileName + `/app/models"
	"` + getFileName + `/app/entity/request"
)

type ` + humpTable + ` struct {
	req request.` + humpTable + `Request
	reqForm request.` + humpTable + `RequestForm
}

func New` + humpTable + `Service() *` + humpTable + ` {
	return &` + humpTable + `{}
}

//SetReq 设置参数
func (svc *` + humpTable + `) SetReq(req request.` + humpTable + `Request) *` + humpTable + ` {
	svc.req = req
	return svc
}

// SetReqForm 设置参数
func (svc *` + humpTable + `) SetReqForm(req request.` + humpTable + `RequestForm) *` + humpTable + ` {
	` + requestStr + `
	svc.reqForm = req
	return svc
}

// Index 分页
func (svc *` + humpTable + `) Index() ([]models.` + humpTable + `, int64, error) {
	return dao.New` + humpTable + `Dao().GetPage(svc.req.PageParam, svc.req.` + humpTable + `)
}

// Show 查询单个
func (svc *` + humpTable + `) Show() models.` + humpTable + ` {
	return dao.New` + humpTable + `Dao().GetById(svc.req.` + humpTable + `.Id)
}

// Store 添加
func (svc *` + humpTable + `) Store() error {
	return dao.New` + humpTable + `Dao().Create(&svc.reqForm.` + humpTable + `)
}

// Update 修改
func (svc *` + humpTable + `) Update() error {
	return dao.New` + humpTable + `Dao().Update(svc.reqForm.` + humpTable + `)
}

// Delete 删除
func (svc *` + humpTable + `) Delete() error {
	return dao.New` + humpTable + `Dao().DeleteById(svc.req.` + humpTable + `.Id)
}

`
}
