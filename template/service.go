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
			requestStr += fmt.Sprintf("req.%s.%s = req.%s",
				utils.ToCamelCase(table),
				utils.ToCamelCase(col.FieldName),
				utils.ToCamelCase(col.FieldName),
			)
		}

	}
	return `package service

import (
	"` + getFileName + `/app/dao"
	"` + getFileName + `/app/model"
	"` + getFileName + `/app/request"
)

type ` + humpTable + ` struct {
	req request.` + humpTable + `Request
	reqForm request.` + humpTable + `RequestForm
}

func New` + humpTable + `Service() *` + humpTable + ` {
	return &` + humpTable + `{}
}

//SetReq 设置参数
func (s *` + humpTable + `) SetReq(req request.` + humpTable + `Request) *` + humpTable + ` {
	s.req = req
	return s
}

// SetReqForm 设置参数
func (s *` + humpTable + `) SetReqForm(req request.` + humpTable + `RequestForm) *` + humpTable + ` {
	` + requestStr + `
	s.reqForm = req
	return s
}

// Index 分页
func (s *` + humpTable + `) Index() ([]model.` + humpTable + `, int64, error) {
	return dao.New` + humpTable + `Dao().GetPage(s.req.PageParam, s.req.` + humpTable + `)
}

// Show 查询单个
func (s *` + humpTable + `) Show() model.` + humpTable + ` {
	return dao.New` + humpTable + `Dao().GetById(s.req.` + humpTable + `.Id)
}

// Store 添加
func (s *` + humpTable + `) Store() error {
	return dao.New` + humpTable + `Dao().Create(&s.reqForm.` + humpTable + `)
}

// Update 修改
func (s *` + humpTable + `) Update() error {
	return dao.New` + humpTable + `Dao().Update(s.reqForm.` + humpTable + `)
}

// Delete 删除
func (s *` + humpTable + `) Delete() error {
	return dao.New` + humpTable + `Dao().DeleteById(s.req.` + humpTable + `.Id)
}

`
}
