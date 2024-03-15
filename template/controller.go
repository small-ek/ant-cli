package template

import (
	"github.com/small-ek/ant-cli/utils"
)

// GenController 生成控制器
func GenController(table string) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	return `package index

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/utils/page"
	"github.com/small-ek/antgo/utils/response"
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

// Index
func (ctrl *` + humpTable + `Controller) Index(c *gin.Context) {
	req := request.` + humpTable + `Request{
		PageParam: page.New(),
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.SecureJSON(200, response.Fail("参数错误", err.Error()))
		return
	}

	list, total, err := ctrl.` + humpTable + `Service.SetReq(req).Index()
	if err != nil {
		c.SecureJSON(200, response.Fail("参数错误", err.Error()))
		return
	}

	c.SecureJSON(200, response.Success("success", response.Page{
		Total: total,
		List:  list,
	}))
}

// Show
func (ctrl *` + humpTable + `Controller) Show(c *gin.Context) {
	req := request.` + humpTable + `Request{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.SecureJSON(200, response.Fail("参数错误", err.Error()))
		return
	}

	row := ctrl.` + humpTable + `Service.SetReq(req).Show()

	c.SecureJSON(200, response.Success("success", row))
}

// Create
func (ctrl *` + humpTable + `Controller) Create(c *gin.Context) {
	req := request.` + humpTable + `Request{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.SecureJSON(200, response.Fail("参数错误", err.Error()))
		return
	}

	if err := ctrl.` + humpTable + `Service.SetReq(req).Store(); err != nil {
		c.SecureJSON(200, response.Fail("error", err.Error()))
		return
	}

	c.SecureJSON(200, response.Success("success"))
}

// Update
func (ctrl *` + humpTable + `Controller) Update(c *gin.Context) {
	req := request.` + humpTable + `Request{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.SecureJSON(200, response.Fail("参数错误", err.Error()))
		return
	}

	row := ctrl.` + humpTable + `Service.SetReq(req).Update()

	c.SecureJSON(200, response.Success("success", row))
}

// Delete
func (ctrl *` + humpTable + `Controller) Delete(c *gin.Context) {
	req := request.` + humpTable + `Request{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.SecureJSON(200, response.Fail("参数错误", err.Error()))
		return
	}

	if err := ctrl.` + humpTable + `Service.SetReq(req).Delete(); err != nil {
		c.SecureJSON(200, response.Fail("参数错误", err.Error()))
		return
	}

	c.SecureJSON(200, response.Success("success"))
}

`
}
