package template

import (
	"github.com/small-ek/ant-cli/utils"
)

// GenService 生成服务端
func GenService(table string) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	return `package service

import (
	"` + getFileName + `/app/dao"
	"` + getFileName + `/app/model"
	"` + getFileName + `/app/request"
)

type ` + humpTable + ` struct {
	req request.` + humpTable + `Request
}

func New` + humpTable + `Service() *` + humpTable + ` {
	return &` + humpTable + `{}
}

//SetReq 设置参数
func (s *` + humpTable + `) SetReq(req request.` + humpTable + `Request) *` + humpTable + ` {
	s.req = req
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
	return dao.New` + humpTable + `Dao().Create(&s.req.` + humpTable + `)
}

// Update 修改
func (s *` + humpTable + `) Update() error {
	return dao.New` + humpTable + `Dao().Update(s.req.` + humpTable + `)
}

// Delete 删除
func (s *` + humpTable + `) Delete() error {
	return dao.New` + humpTable + `Dao().DeleteById(s.req.` + humpTable + `.Id)
}

`
}
