package template

import (
	"github.com/small-ek/ant-cli/utils"
)

// GenService 生成服务端
func GenService(table string, _ []TableStructure) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	return `package service

import (
	"context"
	"` + getFileName + `/app/dao"
	"` + getFileName + `/app/entity/models"
	"` + getFileName + `/app/entity/request"
	"github.com/small-ek/antgo/utils/conv"
)

type ` + humpTable + ` struct {
	dao *dao.` + humpTable + `Dao
	ctx     context.Context
}

func New` + humpTable + `Service(ctx context.Context) *` + humpTable + ` {
	return &` + humpTable + `{
		dao: dao.New` + humpTable + `Dao(ctx, nil),
		ctx: ctx,
	}
}

// SetCtx 设置上下文
func (svc *` + humpTable + `) SetCtx(ctx context.Context) *` + humpTable + ` {
	svc.ctx = ctx
	svc.dao = dao.New` + humpTable + `Dao(ctx, nil)
	return svc
}

// bindForm 将请求参数绑定到模型
func (svc *` + humpTable + `) bindForm(req request.` + humpTable + `RequestForm) (models.` + humpTable + `, error) {
	var data models.` + humpTable + `
	if err := conv.ToStruct(req, &data); err != nil {
		return data, err
	}
	return data, nil
}

// Index 分页
func (svc *` + humpTable + `) Index(req request.` + humpTable + `Request) ([]models.` + humpTable + `, int64, error) {
	return svc.dao.GetPage(req.PageParam)
}

// Show 查询单个
func (svc *` + humpTable + `) Show(req request.` + humpTable + `Request) models.` + humpTable + ` {
	return svc.dao.GetById(req.Id)
}

// Store 添加
func (svc *` + humpTable + `) Store(req request.` + humpTable + `RequestForm) error {
	data, err := svc.bindForm(req)
	if err != nil {
		return err
	}
	_, err = svc.dao.Create(&data)
	return err
}

// Update 修改
func (svc *` + humpTable + `) Update(req request.` + humpTable + `RequestForm) error {
	data, err := svc.bindForm(req)
	if err != nil {
		return err
	}
	return svc.dao.Update(data)
}

// Delete 删除
func (svc *` + humpTable + `) Delete(req request.` + humpTable + `Request) error {
	return svc.dao.DeleteById(req.Id)
}

// Deletes 批量删除
func (svc *` + humpTable + `) Deletes(req request.IdsRequest) error {
	return svc.dao.DeleteByIds(req.Ids)
}
`
}
