<script setup>
import {reactive, ref} from "vue";
import {useI18n} from "vue-i18n"
import {getDatabase, getTable, getTableList} from "@/api/db/index.js"

const {t} = useI18n()
const dbnameList = ref([]);
const tableList = ref([]);
const tableData = ref([]);
const relevanceFieldList = ref([]);
const visible = ref(false);

/**
 * 追加字段
 * @param values
 * @param errors
 */
const newField = ({values, errors}) => {
  if (errors !== undefined) {
    return
  }

  tableData.value.push({
    COLUMN_COMMENT: formField.value.name,
    COLUMN_NAME: formField.value.filed,
    COLUMN_TYPE: "varchar",
    required: 0,
    isSearch: 0
  })
  visible.value = false
}

const form = reactive({
  dbname: '',
  table: ''
});

const formField = ref({
  name: "",
  filed: "",
  is_join_table: false,
  column_name: "",
  join_table: "",
  join_column_filed: ""
});

const rulesForm = {
  name: [
    {
      required: true,
      message: t('verify.databaseName'),
    },
  ],
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
    dataIndex: 'COLUMN_COMMENT',
    ellipsis: true,
    tooltip: true,
  },
  {
    title: '数据库类型',
    dataIndex: 'COLUMN_TYPE',
    ellipsis: true,
    tooltip: true,
  },
  {
    title: '数据库字段名称',
    dataIndex: 'COLUMN_NAME',
    ellipsis: true,
    tooltip: true,
  },
  {
    title: '是否必填',
    dataIndex: 'required',
    slotName: 'required'
  },
  {
    title: '是否搜索',
    dataIndex: 'isSearch',
    slotName: 'isSearch'
  }
];


getDatabase().then(res => {
  dbnameList.value = res.data
})

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
    for (let i = 0; i < res.data.length; i++) {
      res.data[i].required = 0
      res.data[i].isSearch = 0
    }
    tableData.value = res.data
  })
}

const associationTable = () => {
  getTable({db: form.dbname, table: formField.value.join_table}).then(res => {
    relevanceFieldList.value = res.data
  })
}

</script>

<template>
  <div class="container">
    <a-card :style="{ width: '860px',marginTop:'50px' }" :title="$t('code.generation')" hoverable>
      <a-form :rules="rules" :model="form" :style="{width:'600px'}" @submit="handleSubmit">
        <a-form-item field="dbname" :label="$t('code.databaseName')" validate-trigger="blur">
          <a-select v-model="form.dbname" :placeholder="$t('code.select')" allow-clear allow-search @change="onchangeDbName">
            <a-option :value="row" v-for="row in dbnameList">{{ row }}</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="table" :label="$t('code.tableName')" validate-trigger="blur">
          <a-select v-model="form.table" :placeholder="$t('code.select')" allow-clear allow-search>
            <a-option :value="row.table_name" v-for="row in tableList">{{ row.table_name }}</a-option>
          </a-select>
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button html-type="submit" type="primary" shape="round">{{ $t("code.verify") }}</a-button>
          </a-space>
        </a-form-item>
      </a-form>
      <a-form-item>
        <a-button type="primary" shape="round" @click="visible=true">
          <template #icon>
            <icon-plus/>
          </template>
          新建关联表
        </a-button>
      </a-form-item>
      <a-table :columns="columns" :data="tableData" :virtual-list-props="{height:500}" :pagination="false">
        <template #required="{ rowIndex }">
          <a-select :style="{width:'150px'}" v-model="tableData[rowIndex].required" placeholder="请选择">
            <a-option :value="1">是</a-option>
            <a-option :value="0">否</a-option>
          </a-select>
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
          <a-form-item field="name" label="中文名称">
            <a-input v-model="formField.name"/>
          </a-form-item>
          <a-form-item field="filed" label="结构体变量">
            <a-input v-model="formField.filed"/>
          </a-form-item>
          <a-form-item field="is_join_table" label="是否关联表">
            <a-switch v-model="formField.is_join_table"/>
          </a-form-item>
          <template v-if="formField.is_join_table">
            <a-form-item field="column_name" label="当前关联字段">
              <a-select placeholder="请选择" v-model="formField.column_name" allow-clear allow-search>
                <a-option :value="row.COLUMN_NAME" v-for="row in tableData">{{ row["COLUMN_NAME"] }}</a-option>
              </a-select>
            </a-form-item>
            <a-form-item field="join_table" label="关联表">
              <a-select placeholder="请选择" v-model="formField.join_table" allow-clear allow-search @change="associationTable">
                <a-option :value="row.table_name" v-for="row in tableList">{{ row["table_name"] }}</a-option>
              </a-select>
            </a-form-item>
            <a-form-item field="join_column_filed" label="关联表字段">
              <a-select placeholder="请选择" v-model="formField.join_column_filed" allow-clear allow-search>
                <a-option :value="row.COLUMN_NAME" v-for="row in relevanceFieldList">{{ row["COLUMN_NAME"] }}</a-option>
              </a-select>
            </a-form-item>
          </template>
          <a-form-item>
            <a-button html-type="submit" type="primary" shape="round">{{ $t("code.verify") }}</a-button>
          </a-form-item>
        </a-form>
      </div>
    </a-modal>
  </div>
</template>
<style scoped>
.container {
  width: 800px;
  margin: 0 auto;
}
</style>