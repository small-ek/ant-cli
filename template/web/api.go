package web

import (
	"bytes"
	"github.com/small-ek/ant-cli/utils"
	"text/template"
)

func Api(table string) (string, error) {
	const tpl = `import http from "@/utils/request.js";

export function {{.importsApi}}() {
    /**
     * 获取列表
     * @param current_page 当前页
     * @param page_size 每页条数
     * @param filter_map 过滤条件
     * @param order 排序字段
     * @param desc 是否降序
     * @returns {Promise<axios.AxiosResponse<any>>} 返回列表
     */
    const getList = async (current_page = 1, page_size = 2000, filter_map, order = [], desc = []) => {
        return http.get('{{.importsApiName}}', {
            params: {
                current_page: current_page,
                page_size: page_size,
                filter_map: JSON.stringify(filter_map),
                order: order.length > 0 ? order : ['id'],
                desc: desc.length > 0 ? desc : [true],
            }
        });
    }

    /**
     * getSysAdminUsers 获取详情
     * @param id 主键
     * @returns {Promise<axios.AxiosResponse<any>>}
     */
    const show = async (id) => {
        return http.get('{{.importsApiName}}/' + id);
    }

    /**
     * deleteSysAdminUsers 删除
     * @returns {Promise<axios.AxiosResponse<any>>}
     * @param ids 删除的主键
     */
    const deletes = async (ids = []) => {
        return http.delete('{{.importsApiName}}', {
            data: {
                ids: ids
            }
        });
    }

    /**
     * createSysAdminUsers 创建
     * @param data 创建数据
     * @returns {Promise<axios.AxiosResponse<any>>}
     */
    const creates = async (data) => {
        return http.post('{{.importsApiName}}', data);
    }

    /**
     * updateSysAdminUsers 更新
     * @param data 更新数据
     * @returns {Promise<axios.AxiosResponse<any>>}
     */
    const updates = async (data) => {
        return http.put('{{.importsApiName}}/' + data.id, data);
    }
    return {
        getList, show, deletes, creates, updates
    }
}`
	// 小写驼峰
	tableToCamelCaseLower := utils.ToCamelCaseLower(table)
	// 转换为短横线命名
	tableToKebabCase := utils.ToKebabCase(table)
	data := map[string]string{
		"importsApi":     tableToCamelCaseLower,
		"importsApiName": tableToKebabCase,
	}
	// 创建一个新的模板并解析
	tmpl, err := template.New("user").Parse(tpl)
	if err != nil {
		return "", err
	}

	// 使用 bytes.Buffer 来捕获模板输出
	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		return "", err
	}

	// 返回渲染后的字符串
	return result.String(), nil
}
