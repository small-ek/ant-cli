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
  package_name: 'api',
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
      message: t('verify.comment'),
    },
  ],
  join_type: [
    {
      required: true,
      message: t('verify.join_type'),
    },
  ],
  field_name: [
    {
      required: true,
      message: t('verify.field_name'),
    },
  ],
  join_table: [
    {
      required: true,
      message: t('verify.join_table'),
    },
  ],
  join_field: [
    {
      required: true,
      message: t('verify.join_field'),
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
  package_name: [
    {
      required: true,
      message: t('verify.packageName'),
    },
  ],
}

const columns = [
  {
    title: t('table_columns.comment'),
    dataIndex: 'comment',
    ellipsis: true,
    tooltip: true,
    width: 140
  },
  {
    title: t('table_columns.field_type'),
    dataIndex: 'field_type',
    ellipsis: true,
    tooltip: true,
    width: 100
  },
  {
    title: t('table_columns.field_name'),
    dataIndex: 'field_name',
    ellipsis: true,
    tooltip: true,
    width: 130
  },
  {
    title: t('table_columns.required'),
    dataIndex: 'required',
    slotName: 'required',
    width: 155
  },
  {
    title: t('table_columns.is_search'),
    dataIndex: 'is_search',
    slotName: 'is_search',
    width: 155
  },
  {
    title: t('table_columns.conditions'),
    dataIndex: 'conditions',
    slotName: 'conditions',
    width: 130
  },
  {
    title: t('table_columns.option'),
    dataIndex: 'option',
    slotName: 'option',
    width: 160
  }
];

//关联模型
const correlationModel = [{name: t('correlationModel.oneToOne'), value: "oneToOne"}, {name: t('correlationModel.oneToMany'), value: "oneToMany"}, {name: t('correlationModel.manyToMany'), value: "manyToMany"}]

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
  if (form.table === "" || tableData.value.length === 0 || form.package_name === "") {
    Message.error(t('tips.fail'))
    return
  }
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
      Message.info(t('tips.success'))
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
      <a-form :rules="rules" :model="form" :style="{width:'850px'}" @submit="handleSubmit">
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
          <a-input v-model="form.package_name" :placeholder="$t('code.input')"/>
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
        {{ $t("form.addField") }}
      </a-button>
      <!--表格-->
      <a-table style="margin-top: 10px" :stripe="true" :hoverable="true" :columns="columns" :data="tableData" :pagination="false">
        <template #required="{ rowIndex }">
          <a-radio-group v-if="tableData[rowIndex].field_type!='join'" v-model="tableData[rowIndex].required">
            <a-radio :value="1">{{ $t('tableSelect.ok') }}</a-radio>
            <a-radio :value="0">{{ $t('tableSelect.no') }}</a-radio>
          </a-radio-group>
        </template>
        <template #is_search="{ rowIndex }">
          <a-radio-group v-if="tableData[rowIndex].field_type!='join'" v-model="tableData[rowIndex].is_search">
            <a-radio :value="1">{{ $t('tableSelect.ok') }}</a-radio>
            <a-radio :value="0">{{ $t('tableSelect.no') }}</a-radio>
          </a-radio-group>
        </template>
        <template #conditions="{ rowIndex }">
          <a-select v-if="tableData[rowIndex].is_search==1" :style="{width:'100px'}" v-model="tableData[rowIndex].conditions" :placeholder="$t('code.select')">
            <a-option v-for="row in conditionsList" :value="row">{{ row }}</a-option>
          </a-select>
        </template>
        <template #option="{ rowIndex }">
          <a-space v-if="tableData[rowIndex].field_type=='join'">
            <a-button type="text" style="margin-left: 10px" shape="round" size="mini" @click="editForm(tableData[rowIndex],rowIndex)">
              <template #icon>
                <icon-eye/>
              </template>
              {{ $t('tableOption.edit') }}
            </a-button>
            <a-button type="text" status="danger" shape="round" size="mini" @click="delTable(rowIndex)">
              <template #icon>
                <icon-delete/>
              </template>
              {{ $t('tableOption.delete') }}
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
          <a-form-item field="comment" :label="$t('form2.chinese_name')">
            <a-input v-model="formField.comment"/>
          </a-form-item>
          <a-form-item field="join_type" :label="$t('form2.join_type')" :help="$t('form2.help')">
            <a-select :placeholder="$t('code.select')" v-model="formField.join_type" allow-clear allow-search>
              <a-option :value="row.value" v-for="row in correlationModel">{{ row["name"] }}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="field_name" :label="$t('form2.field_name')">
            <a-select :placeholder="$t('code.select')" v-model="formField.field_name" allow-clear allow-search>
              <a-option :value="row.field_name" v-for="row in tableData">{{ row["field_name"] }}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="join_table" :label="$t('form2.join_table')">
            <a-select :placeholder="$t('code.select')" v-model="formField.join_table" allow-clear allow-search @change="associationTable">
              <a-option :value="row.table_name" v-for="row in tableList">{{ row["table_name"] }}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="join_field" :label="$t('form2.join_field')">
            <a-select :placeholder="$t('code.select')" v-model="formField.join_field" allow-clear allow-search>
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
        {{ $t('tableOption.preview') }}
      </a-button>
      <a-popconfirm :content="t('tips.is_gen_code')" type="info" @ok="getPreviewCode(true)">
        <a-button type="primary" style="margin-left: 20px" shape="round">
          <template #icon>
            <icon-desktop/>
          </template>
          {{ $t("tableOption.gen_code") }}
        </a-button>
      </a-popconfirm>

    </div>
    <!--预览代码-->
    <template v-if="preCodeStatus">
      <autoCode v-model:visible="preCodeStatus" v-model:preCode="preCode"></autoCode>
    </template>
  </div>
</template>
<style scoped>
.container {
  width: 1100px;
  margin: 0 auto;
}
</style>