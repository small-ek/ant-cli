package web

import (
	"bytes"
	"github.com/small-ek/ant-cli/utils"
	"text/template"
)

func ViewsJs(table string) (string, error) {
	const tpl = `import {reactive, ref} from "vue";

//分页参数
const page = ref({
    current: 1,
    pageSize: 10,
    total: 0,
    searchForm: {},
    order: ['id'],
    desc: [true]
})
const tableRef = ref(null)
const list = ref([]);
const formRef = ref(null)
const formData = ref({})
const ids = ref([])
const formTitle = ref("")
//搜索列表
const searchList = ref([
	{{. searchList}}    
]);

//表格列
const columns = [
    {
        title: '标识',
        dataIndex: 'id',
        width: 120,
        sortable: {
            sortDirections: ['ascend', 'descend']
        },
        visible: true
    },
    {
        title: '产品名称',
        dataIndex: 'product_name',
        tooltip: true,
        ellipsis: true,
        visible: true
    },
    {
        title: '贷款期限',
        dataIndex: 'loan_term',
        tooltip: true,
        ellipsis: true,
        visible: true
    },{
        title: '资金方',
        dataIndex: 'fund_provider',
        tooltip: true,
        ellipsis: true,
        visible: true
    },{
        title: '单价',
        dataIndex: 'unit_price',
        tooltip: true,
        ellipsis: true,
        visible: true
    },{
        title: '操作',
        slotName: 'optional',
        width: 220,
        visible: true
    }
];

//表单验证
const formRules = reactive({
    name: [
        {required: true, message: '请输入角色名称', trigger: 'blur'},
    ]
})

//表单列表
const formList = ref([
    {
        label: '角色名称',
        key: 'name',
        value: "",
        type: 'input',
        placeholder: '请输入角色名称'
    },
    {
        label: '角色菜单',
        key: 'menu_ids',
        value: [],
        type: 'tree',
        options: [],
        expandedKeys: [],
        checkedStrategy: 'all',
        checkStrictly: true
    },
    {
        label: '角色API',
        key: 'api_ids',
        value: [],
        type: 'tree',
        options: [],
        expandedKeys: [],
    }
]);

/**
 * 表单列表索引映射
 * @type {{}}
 */
const formListIndexMap = formList.value.reduce((map, item, index) => {
    map[item.key] = index;
    return map;
}, {});

/**
 * 搜索框列表索引映射
 * @type {{}}
 */
const searchListIndexMap = searchList.value.reduce((map, item, index) => {
    map[item.key] = index;
    return map;
}, {});

const onSelect = (selectedRowKeys) => {
    ids.value = selectedRowKeys;
};

export {
    onSelect,
    searchList,
    columns,
    page,
    formList,
    list,
    formRef,
    formData,
    formRules,
    ids,
    tableRef,
    formTitle,
    formListIndexMap,
    searchListIndexMap
};`
	// 小写驼峰
	tableToCamelCaseLower := utils.ToCamelCaseLower(table)
	// 转换为短横线命名
	tableToKebabCase := utils.ToKebabCase(table)
	data := map[string]string{
		"importsApi":     `import {` + tableToCamelCaseLower + `} from "@/api/` + tableToKebabCase + `.js";`,
		"importsApiName": tableToCamelCaseLower,
		"timeFormat":     `{{formatTime(record.created_at)}}`,
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
