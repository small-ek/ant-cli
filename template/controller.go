package template

import (
	"bytes"
	"context"
	"text/template"

	"github.com/small-ek/ant-cli/utils"
	"github.com/small-ek/antgo/os/alog"
	"go.uber.org/zap"
)

// ControllerData 定义生成控制器需要的数据结构
type ControllerData struct {
	Package    string // 包名
	TableName  string // 原始表名
	TableCamel string // 驼峰表名
	TableKebab string // kebab-case 表名
	Comment    string // 控制器注释
	FileName   string // 项目文件名
	HasCreate  bool   // 是否生成 Create 方法
	HasUpdate  bool   // 是否生成 Update 方法
	HasDelete  bool   // 是否生成 Delete 方法
	HasIndex   bool   // 是否生成 Index 方法
	HasShow    bool   // 是否生成 Show 方法
}

// GenController 根据模板生成控制器代码
func GenController(data ControllerData) string {
	// 内部自动处理的参数
	humpTable := utils.ToCamelCase(data.TableName)
	kebabTable := utils.ToKebabCase(data.TableName)
	fileName := utils.GetFileName()
	data.TableCamel = humpTable
	data.TableKebab = kebabTable
	data.FileName = fileName

	tpl := `package {{.Package}}

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/utils/page"
	"{{.FileName}}/app/entity/vo"
	"{{.FileName}}/app/entity/request"
	"{{.FileName}}/app/service"
)

type {{.TableCamel}}Controller struct {
	vo.Base
}

func New{{.TableCamel}}Controller() *{{.TableCamel}}Controller {
	return &{{.TableCamel}}Controller{}
}

{{- if .HasIndex }}
// @Tags {{.Comment}}
// @Summary 获取{{.Comment}}分页数据
// @Accept json
// @Produce json
// @Param data query page.PageParam true "分页参数"
// @Param filter_map query string false "JSON格式的过滤字段" example("{\"key1\":\"value1\"}")
// @Param extra query string false "其他参数" example("{\"key1\":\"value1\"}")
// @Success 0 {object} response.Write{data=response.Page{items=[]models.{{.TableCamel}}}}
// @Failure 1 {object} response.Write
// @Router /{{.TableKebab}} [get]
func (ctrl *{{.TableCamel}}Controller) Index(c *gin.Context) {
	req := request.{{.TableCamel}}Request{
		PageParam: page.New(),
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}

	list, total, err := service.New{{.TableCamel}}Service(c).Index(req)
	if err != nil {
		ctrl.Fail(c, vo.FAILED, err)
		return
	}
	ctrl.Success(c, vo.SUCCESS, ctrl.Page(total, list))
}
{{- end }}

{{- if .HasShow }}
// @Tags {{.Comment}}
// @Summary 获取{{.Comment}}详情数据
// @Accept json
// @Produce json
// @Success 0 {object} response.Write{data=models.{{.TableCamel}}}
// @Failure 1 {object} response.Write
// @Router /{{.TableKebab}}/:id [get]
func (ctrl *{{.TableCamel}}Controller) Show(c *gin.Context) {
	var req request.{{.TableCamel}}Request
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}

	result := service.New{{.TableCamel}}Service(c).Show(req)
	ctrl.Success(c, vo.SUCCESS, result)
}
{{- end }}

{{- if .HasCreate }}
// @Tags {{.Comment}}
// @Summary 创建{{.Comment}}数据
// @Accept json
// @Produce json
// @Param data body request.{{.TableCamel}}RequestForm true "创建参数"
// @Success 0 {object} response.Write
// @Failure 1 {object} response.Write
// @Router /{{.TableKebab}} [post]
func (ctrl *{{.TableCamel}}Controller) Create(c *gin.Context) {
	var req request.{{.TableCamel}}RequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}

	if err := service.New{{.TableCamel}}Service(c).Store(req); err != nil {
		ctrl.Fail(c, vo.CREATION_FAILED, err)
		return
	}
	ctrl.Success(c, vo.CREATION_SUCCESS)
}
{{- end }}

{{- if .HasUpdate }}
// @Tags {{.Comment}}
// @Summary 修改{{.Comment}}数据
// @Accept json
// @Produce json
// @Param data body request.{{.TableCamel}}RequestForm true "更新参数"
// @Success 0 {object} response.Write
// @Failure 1 {object} response.Write
// @Router /{{.TableKebab}}/:id [put]
func (ctrl *{{.TableCamel}}Controller) Update(c *gin.Context) {
	var req request.{{.TableCamel}}RequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}

	if err := service.New{{.TableCamel}}Service(c).Update(req); err != nil {
		ctrl.Fail(c, vo.UPDATE_FAILED, err)
		return
	}
	ctrl.Success(c, vo.UPDATE_SUCCESS)
}
{{- end }}

{{- if .HasDelete }}
// @Tags {{.Comment}}
// @Summary 删除{{.Comment}}数据
// @Accept json
// @Produce json
// @Param data body request.IdsRequest true "删除参数"
// @Success 0 {object} response.Write
// @Failure 1 {object} response.Write
// @Router /{{.TableKebab}} [delete]
func (ctrl *{{.TableCamel}}Controller) Delete(c *gin.Context) {
	var req request.IdsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}

	if err := service.New{{.TableCamel}}Service(c).Deletes(req); err != nil {
		ctrl.Fail(c, vo.DELETE_FAILED, err)
		return
	}

	ctrl.Success(c, vo.DELETE_SUCCESS)
}
{{- end }}
`

	var buf bytes.Buffer
	t := template.Must(template.New("controller").Parse(tpl))
	if err := t.Execute(&buf, data); err != nil {
		alog.Error(context.Background(), "GenController", zap.Error(err))
		return ""
	}
	return buf.String()
}
