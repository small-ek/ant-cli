<script setup>
import {reactive, ref} from "vue";
import {useI18n} from "vue-i18n"
import {getDatabase, getTable, getTableList, previewCode} from "../../api/db/index.js";
import autoCode from "../../components/autoCode/index.vue";
import {Message} from "@arco-design/web-vue";

const {t} = useI18n()
const dbnameList = ref([]);
const tableList = ref([]);
const tableData = ref([]);
const relevanceFieldList = ref([]);
const visible = ref(false);
const preCodeStatus = ref(false)
const preCode = ref({})

/**
 * 追加字段
 * @param values
 * @param errors
 */
const newField = ({values, errors}) => {
  if (errors !== undefined) {
    return
  }
  if (formField.value.update_index == 0) {
    tableData.value.push({
      comment: formField.value.comment,
      field_name: formField.value.field_name,
      field_type: "join",
      join_type: formField.value.join_type,
      join_table: formField.value.join_table,
      join_field: formField.value.join_field,
      required: 0,
      is_search: 0,
      conditions: ""
    })
  } else {
    tableData.value[formField.value.update_index] = {
      comment: formField.value.comment,
      field_name: formField.value.field_name,
      field_type: "join",
      join_type: formField.value.join_type,
      join_table: formField.value.join_table,
      join_field: formField.value.join_field,
      required: 0,
      is_search: 0,
      conditions: ""
    }
  }

  visible.value = false
}

const form = reactive({
  dbname: '',
  table: '',
  package_name: 'index',
  table_comment: ''
});

const formField = ref({
  comment: "",
  field_name: "",
  field_type: "",
  join_table: "",
  join_field: "",
  join_type: "",
  update_index: 0
});

const rulesForm = {
  comment: [
    {
      required: true,
      message: t('verify.databaseName'),
    },
  ],
  join_type: [
    {
      required: true,
      message: "请选择关联类型",
    },
  ],
  field_name: [
    {
      required: true,
      message: "请选择当前表字段",
    },
  ],
  join_table: [
    {
      required: true,
      message: "请选择关联表",
    },
  ],
  join_field: [
    {
      required: true,
      message: "请选择关联表字段",
    },
  ]
}

const rules = {
  dbname: [
    {
      required: true,
      message: t('verify.databaseName'),
    },
  ],
  table: [
    {
      required: true,
      message: t('verify.tableName'),
    },
  ],
}

const columns = [
  {
    title: '中文名称',
    dataIndex: 'comment',
    ellipsis: true,
    tooltip: true,
    width: 150
  },
  {
    title: '数据库类型',
    dataIndex: 'field_type',
    ellipsis: true,
    tooltip: true,
    width: 120
  },
  {
    title: '数据库字段名称',
    dataIndex: 'field_name',
    ellipsis: true,
    tooltip: true,
    width: 180
  },
  {
    title: '是否必填',
    dataIndex: 'required',
    slotName: 'required',
    width: 150
  },
  {
    title: '是否搜索',
    dataIndex: 'is_search',
    slotName: 'is_search',
    width: 150
  },
  {
    title: '查询条件',
    dataIndex: 'conditions',
    slotName: 'conditions',
    width: 150
  },
  {
    title: '操作',
    dataIndex: 'option',
    slotName: 'option',
    width: 180
  }
];

//关联模型
const correlationModel = [{name: "一对一", value: "oneToOne"}, {name: "一对多", value: "oneToMany"}, {name: "多对多", value: "manyToMany"}]

//获取数据库
getDatabase().then(res => {
  dbnameList.value = res.data
})

//设置表注释
const setTableComment = (value) => {
  tableList.value.forEach(item => {
    if (item.table_name == value) {
      form.table_comment = item.table_comment
    }
  })
}

//获取数据库表
const onchangeDbName = () => {
  getTableList({table: form.dbname}).then(res => {
    tableList.value = res.data
  })
}
const handleSubmit = ({values, errors}) => {
  if (errors !== undefined) {
    return
  }

  getTable({db: form.dbname, table: form.table}).then(res => {
    const data = []
    for (let i = 0; i < res.data.length; i++) {
      data.push({
        comment: res.data[i].comment,
        field_name: res.data[i].field_name,
        field_type: res.data[i].field_type,
        indexes: res.data[i].indexes,
        required: 0,
        is_search: 0
      })
    }
    tableData.value = data
  })
}
const associationTable = () => {
  getTable({db: form.dbname, table: formField.value.join_table}).then(res => {
    relevanceFieldList.value = res.data
  })
}

// getPreviewCode 获取预览代码
const getPreviewCode = (is_create) => {
  previewCode({
    table_name: form.table,
    fields: tableData.value,
    package: form.package_name,
    is_create: is_create,
    data_base: form.dbname,
    table_comment: form.table_comment
  }).then(res => {
    if (is_create == false) {
      preCodeStatus.value = true
    } else {
      Message.info("生成成功")
    }
    preCode.value = res.data
  })
}

// editForm 编辑表格
const editForm = (row, index) => {
  formField.value = {
    comment: row.comment,
    field_name: row.field_name,
    field_type: row.field_type,
    join_table: row.join_table,
    join_field: row.join_field,
    join_type: row.join_type,
    update_index: index
  }
  visible.value = true
}
const conditionsList = ["=", "!=", "<", ">", "<=", ">=", "LIKE", "IN", "NOT IN", "BETWEEN", "NOT BETWEEN"];
// delTable 删除表格
const delTable = (index) => {
  tableData.value.splice(index, 1)
}
</script>

<template>
  <div class="container">
    <a-card :style="{ width: '1000px',marginTop:'50px' }" :title="$t('code.generation')" hoverable>
      <a-form :rules="rules" :model="form" :style="{width:'600px'}" @submit="handleSubmit">
        <a-form-item field="dbname" :label="$t('code.databaseName')" validate-trigger="blur">
          <a-select v-model="form.dbname" :placeholder="$t('code.select')" allow-clear allow-search @change="onchangeDbName">
            <a-option :value="row" v-for="row in dbnameList">{{ row }}</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="table" :label="$t('code.tableName')" validate-trigger="blur">
          <a-select v-model="form.table" :placeholder="$t('code.select')" allow-clear allow-search @change="setTableComment">
            <a-option :value="row.table_name" v-for="row in tableList">{{ row.table_name }}</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="package_name" :label="$t('code.packageName')" validate-trigger="blur">
          <a-input v-model="form.package_name" placeholder="请输入包名"/>
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button html-type="submit" type="primary" shape="round">{{ $t("code.verify") }}</a-button>
          </a-space>
        </a-form-item>
      </a-form>
      <a-button type="primary" shape="round" @click="visible=true">
        <template #icon>
          <icon-plus/>
        </template>
        新建关联表
      </a-button>
      <!--表格-->
      <a-table style="margin-top: 10px" :columns="columns" :data="tableData" :pagination="false">
        <template #required="{ rowIndex }">
          <a-select v-if="tableData[rowIndex].field_type!='join'" :style="{width:'100px'}" v-model="tableData[rowIndex].required" placeholder="请选择">
            <a-option :value="1">是</a-option>
            <a-option :value="0">否</a-option>
          </a-select>
        </template>
        <template #is_search="{ rowIndex }">
          <a-select v-if="tableData[rowIndex].field_type!='join'" :style="{width:'100px'}" v-model="tableData[rowIndex].is_search" placeholder="请选择">
            <a-option :value="1">是</a-option>
            <a-option :value="0">否</a-option>
          </a-select>
        </template>
        <template #conditions="{ rowIndex }">
          <a-select v-if="tableData[rowIndex].is_search==1" :style="{width:'100px'}" v-model="tableData[rowIndex].conditions" placeholder="请选择">
            <a-option v-for="row in conditionsList" :value="row">{{ row }}</a-option>
          </a-select>
        </template>
        <template #option="{ rowIndex }">
          <a-space v-if="tableData[rowIndex].field_type=='join'">
            <a-button type="text" style="margin-left: 10px" shape="round" size="mini" @click="editForm(tableData[rowIndex],rowIndex)">
              <template #icon>
                <icon-eye/>
              </template>
              编辑
            </a-button>
            <a-button type="text" status="danger" style="margin-left: 10px" shape="round" size="mini" @click="delTable(rowIndex)">
              <template #icon>
                <icon-delete/>
              </template>
              删除
            </a-button>
          </a-space>
        </template>
      </a-table>
    </a-card>
    <!--表单-->
    <a-modal v-model:visible="visible" :footer="false" draggable>
      <template #title>
        编辑字段
      </template>
      <div>
        <a-form :model="formField" :rules="rulesForm" @submit="newField">
          <a-form-item field="comment" label="中文名称">
            <a-input v-model="formField.comment"/>
          </a-form-item>
          <a-form-item field="join_type" label="关联模型" help="多对多模型,中间表要求表名+字段名称拼接，必须驼峰命名,例如(user_refer_id)，user_refer是表名称">
            <a-select placeholder="请选择" v-model="formField.join_type" allow-clear allow-search>
              <a-option :value="row.value" v-for="row in correlationModel">{{ row["name"] }}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="field_name" label="当前表字段">
            <a-select placeholder="请选择" v-model="formField.field_name" allow-clear allow-search>
              <a-option :value="row.field_name" v-for="row in tableData">{{ row["field_name"] }}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="join_table" label="关联表">
            <a-select placeholder="请选择" v-model="formField.join_table" allow-clear allow-search @change="associationTable">
              <a-option :value="row.table_name" v-for="row in tableList">{{ row["table_name"] }}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="join_field" label="关联表字段">
            <a-select placeholder="请选择" v-model="formField.join_field" allow-clear allow-search>
              <a-option :value="row.field_name" v-for="row in relevanceFieldList">{{ row["field_name"] }}</a-option>
            </a-select>
          </a-form-item>

          <a-form-item>
            <a-button html-type="submit" type="primary" shape="round">{{ $t("code.verify") }}</a-button>
          </a-form-item>
        </a-form>
      </div>
    </a-modal>
    <div style="margin-top: 15px">
      <a-button type="primary" shape="round" @click="getPreviewCode(false)">
        <template #icon>
          <icon-eye/>
        </template>
        预览代码
      </a-button>
      <a-button type="primary" style="margin-left: 20px" shape="round" @click="getPreviewCode(true)">
        <template #icon>
          <icon-desktop/>
        </template>
        生成代码
      </a-button>
    </div>
    <!--预览代码-->
    <template v-if="preCodeStatus">
      <autoCode v-model:visible="preCodeStatus" v-model:preCode="preCode"></autoCode>
    </template>
  </div>
</template>
<style scoped>
.container {
  width: 800px;
  margin: 0 auto;
}
</style>