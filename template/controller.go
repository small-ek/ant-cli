package template

import (
	"github.com/small-ek/ant-cli/utils"
)

// GenController 生成控制器
func GenController(table, comment, packages string) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	return `package ` + packages + `

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/utils/page"
	"` + getFileName + `/app/http"
	"` + getFileName + `/app/request"
	"` + getFileName + `/app/service"
)

type ` + humpTable + `Controller struct {
	http.Base
	` + humpTable + `Service *service.` + humpTable + `
}

func New` + humpTable + `Controller() *` + humpTable + `Controller {
	return &` + humpTable + `Controller{
		` + humpTable + `Service: service.New` + humpTable + `Service(),
	}
}

//	@Summary		获取` + comment + `分页数据
//	@Description	获取` + comment + `分页
//	@Tags			` + humpTable + `
//	@Accept			json
//	@Produce		json
//	@Param		    data query request.` + humpTable + `Request true "分页获取列表"
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/` + table + ` [get] 路由
//
// Index 分页列表
func (ctrl *` + humpTable + `Controller) Index(c *gin.Context) {
	req := request.` + humpTable + `Request{
		PageParam: page.New(),
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
		return
	}

	list, total, err := ctrl.` + humpTable + `Service.SetReq(req).Index()
	if err != nil {
		ctrl.Fail(c, "异常错误", err.Error())
		return
	}
	ctrl.Success(c, "success", ctrl.Page(total, list))
}

//	@Summary		获取` + comment + `单条数据
//	@Description	获取` + comment + `详情
//	@Tags			` + humpTable + `
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/` + table + `/:id [get]
//
// Show 详情
func (ctrl *` + humpTable + `Controller) Show(c *gin.Context) {
	var req request.` + humpTable + `Request
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
		return
	}

	row := ctrl.` + humpTable + `Service.SetReq(req).Show()
	ctrl.Success(c, "success", row)
}

//	@Summary		创建` + comment + `数据
//	@Description	创建` + comment + `数据
//	@Tags			` + humpTable + `
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.` + humpTable + `RequestForm true "创建"
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/` + table + ` [post]
//
// Create 创建
func (ctrl *` + humpTable + `Controller) Create(c *gin.Context) {
	var req request.` + humpTable + `RequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
		return
	}

	if err := ctrl.` + humpTable + `Service.SetReqForm(req).Store(); err != nil {
		ctrl.Fail(c, "创建失败", err.Error())
		return
	}
	ctrl.Success(c, "success")
}

//	@Summary		修改` + comment + `数据
//	@Description	修改` + comment + `数据
//	@Tags			` + humpTable + `
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.` + humpTable + `RequestForm true "更新"
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/` + table + `/:id [put]
//
// Update 修改
func (ctrl *` + humpTable + `Controller) Update(c *gin.Context) {
	var req request.` + humpTable + `RequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
		return
	}

	if err := ctrl.` + humpTable + `Service.SetReqForm(req).Update();err!=nil{
		ctrl.Fail(c, "参数错误", err.Error())
	}
	ctrl.Success(c, "success")
}

//	@Summary		删除` + comment + `数据
//	@Description	删除` + comment + `数据
//	@Tags			` + humpTable + `
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/` + table + `/:id [delete]
//
// Delete 删除
func (ctrl *` + humpTable + `Controller) Delete(c *gin.Context) {
	var req request.` + humpTable + `Request
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
		return
	}

	if err := ctrl.` + humpTable + `Service.SetReq(req).Delete(); err != nil {
		ctrl.Fail(c, "删除失败", err.Error())
		return
	}

	ctrl.Success(c, "success")
}

`
}
