package template

import (
	"github.com/small-ek/ant-cli/utils"
)

// GenController 生成控制器
func GenController(table, comment, packages string) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	toKebabCase := utils.ToKebabCase(table)
	return `package ` + packages + `

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/utils/page"
	"` + getFileName + `/app/entity/vo"
	"` + getFileName + `/app/entity/request"
	"` + getFileName + `/app/service"
)

type ` + humpTable + `Controller struct {
	vo.Base
	` + humpTable + `Service *service.` + humpTable + `
}

func New` + humpTable + `Controller() *` + humpTable + `Controller {
	return &` + humpTable + `Controller{
		` + humpTable + `Service: service.New` + humpTable + `Service(),
	}
}

//	@Tags			` + comment + `
//	@Summary		获取` + comment + `分页数据
//	@Accept			json
//	@Produce		json
//	@Param		    data query page.PageParam true "分页参数"
//	@Param		    filter_map		query	string		false	"JSON格式的过滤字段" example("{\"key1\":\"value1\",\"key2\": 1}")
//	@Param		    extra			query	string		false	"其他参数" example("{\"key1\":\"value1\",\"key2\": 1}")
//	@Success		0	{object}	response.Write{data=response.Page{items=[]models.` + humpTable + `}}
//	@Failure		1	{object}	response.Write
//	@Router			/` + toKebabCase + ` [get] 路由
//
// Index 分页列表
func (ctrl *` + humpTable + `Controller) Index(c *gin.Context) {
	req := request.` + humpTable + `Request{
		PageParam: page.New(),
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}

	list, total, err := ctrl.` + humpTable + `Service.SetReq(c,req).Index()
	if err != nil {
		ctrl.Fail(c, vo.FAILED, err)
		return
	}
	ctrl.Success(c, vo.SUCCESS, ctrl.Page(total, list))
}

//	@Tags			` + comment + `
//	@Summary		获取` + comment + `详情数据
//	@Accept			json
//	@Produce		json
//	@Success		0	{object}	response.Write{data=models.` + humpTable + `}
//	@Failure		1	{object}	response.Write
//	@Router			/` + toKebabCase + `/:id [get]
//
// Show 详情
func (ctrl *` + humpTable + `Controller) Show(c *gin.Context) {
	var req request.` + humpTable + `Request
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}

	result := ctrl.` + humpTable + `Service.SetReq(c,req).Show()
	ctrl.Success(c, vo.SUCCESS, result)
}

//	@Tags			` + comment + `
//	@Summary		创建` + comment + `数据
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.` + humpTable + `RequestForm true "创建参数"
//	@Success		0	{object}	response.Write
//	@Failure		1	{object}	response.Write
//	@Router			/` + toKebabCase + ` [post]
//
// Create 创建
func (ctrl *` + humpTable + `Controller) Create(c *gin.Context) {
	var req request.` + humpTable + `RequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}

	if err := ctrl.` + humpTable + `Service.SetReq(c,req).Store(); err != nil {
		ctrl.Fail(c, vo.CREATION_FAILED, err)
		return
	}
	ctrl.Success(c, vo.CREATION_SUCCESS)
}

//	@Tags			` + comment + `
//	@Summary		修改` + comment + `数据
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.` + humpTable + `RequestForm true "更新参数"
//	@Success		0	{object}	response.Write
//	@Failure		1	{object}	response.Write
//	@Router			/` + toKebabCase + `/:id [put]
//
// Update 修改
func (ctrl *` + humpTable + `Controller) Update(c *gin.Context) {
	var req request.` + humpTable + `RequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}

	if err := ctrl.` + humpTable + `Service.SetReq(c,req).Update();err!=nil{
		ctrl.Fail(c, vo.UPDATE_FAILED, err)
		return
	}
	ctrl.Success(c, vo.UPDATE_SUCCESS)
}

//	@Tags			` + comment + `
//	@Summary		删除` + comment + `数据
//	@Accept			json
//	@Produce		json
//	@Success		0	{object}	response.Write
//	@Failure		1	{object}	response.Write
//	@Router			/` + toKebabCase + `/:id [delete]
//
// Delete 删除
func (ctrl *` + humpTable + `Controller) Delete(c *gin.Context) {
	var req request.` + humpTable + `Request
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}

	if err := ctrl.` + humpTable + `Service.SetReq(c,req).Delete(); err != nil {
		ctrl.Fail(c, vo.DELETE_FAILED, err)
		return
	}

	ctrl.Success(c, vo.DELETE_SUCCESS)
}

`
}
