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
			requestStr += fmt.Sprintf("svc.req.%s.%s = value.%s\n",
				utils.ToCamelCase(table),
				utils.ToCamelCase(col.FieldName),
				utils.ToCamelCase(col.FieldName),
			)
		}

	}
	return `package service

import (
	"github.com/small-ek/antgo/os/alog"
	"go.uber.org/zap"
	"reflect"
	"` + getFileName + `/app/dao"
	"` + getFileName + `/app/entity/models"
	"` + getFileName + `/app/entity/request"
	"github.com/small-ek/antgo/utils/conv"
)

type ` + humpTable + ` struct {
	req request.` + humpTable + `Request
	reqForm request.` + humpTable + `RequestForm
	reqIds  request.IdsRequest
}

func New` + humpTable + `Service() *` + humpTable + ` {
	return &` + humpTable + `{}
}

//SetReq 设置参数
func (svc *` + humpTable + `) SetReq(ctx context.Context,req interface{}) *` + humpTable + ` {
	svc.ctx = ctx
	switch value := req.(type) {
		case request.IdsRequest:
			svc.reqIds = value
		case request.` + humpTable + `Request:
			svc.req = value
		case request.` + humpTable + `RequestForm:
			svc.reqForm.` + humpTable + ` = models.` + humpTable + `{}
			conv.ToStruct(value, &svc.reqForm.` + humpTable + `)
		default:
			alog.Write.Error("SetReq", zap.Any("Unsupported request type", reflect.TypeOf(value)))
	}
	return svc
}

// SetCtx 设置上下文
func (svc *TLearningResourceTag) SetCtx(ctx context.Context) *` + humpTable + ` {
	svc.ctx = ctx
	return svc
}

// Index 分页
func (svc *` + humpTable + `) Index() ([]models.` + humpTable + `, int64, error) {
	return dao.New` + humpTable + `Dao(svc.ctx,nil).GetPage(svc.req.PageParam)
}

// Show 查询单个
func (svc *` + humpTable + `) Show() models.` + humpTable + ` {
	return dao.New` + humpTable + `Dao(svc.ctx,nil).GetById(svc.req.` + humpTable + `.Id)
}

// Store 添加
func (svc *` + humpTable + `) Store() error {
	return dao.New` + humpTable + `Dao(svc.ctx,nil).Create(svc.reqForm.` + humpTable + `)
}

// Update 修改
func (svc *` + humpTable + `) Update() error {
	return dao.New` + humpTable + `Dao(svc.ctx,nil).Update(svc.reqForm.` + humpTable + `)
}

// Delete 删除
func (svc *` + humpTable + `) Delete() error {
	return dao.New` + humpTable + `Dao(svc.ctx,nil).DeleteById(svc.req.` + humpTable + `.Id)
}

// Deletes 批量删除
func (svc *` + humpTable + `) Deletes() error {
	return dao.New` + humpTable + `Dao(svc.ctx,nil).DeleteByIds(svc.reqIds.Ids)
}
`
}
