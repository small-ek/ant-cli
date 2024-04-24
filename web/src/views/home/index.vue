<script setup>
import {reactive, ref} from "vue";
import {useI18n} from "vue-i18n"
import {getDatabase, getTable, getTableList} from "@/api/db/index.js"

const {t} = useI18n()
const dbnameList = ref([]);
const tableList = ref([]);
const tableData = ref([]);

const form = reactive({
  dbname: '',
  table: ''
});


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
  },{
    title: '是否必填',
    dataIndex: 'required',
    slotName: 'required'
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
  console.log('values:', values, '\nerrors:', errors)
  if (errors !== undefined) {
    return
  }
  getTable({db: form.dbname, table: form.table}).then(res => {
    tableData.value = res.data
  })
}

</script>

<template>
  <div class="container">
    <a-card :style="{ width: '860px',marginTop:'50px' }" :title="$t('code.generation')" hoverable>
      <a-form ref="formRef" :rules="rules" :model="form" :style="{width:'600px'}" @submit="handleSubmit">
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

      <a-table :columns="columns" :data="tableData" :virtual-list-props="{height:500}" :pagination="false">
        <template #required="{ rowIndex }">
          <a-select :style="{width:'150px'}" v-model="tableData[rowIndex].required" placeholder="请选择">
            <a-option>是</a-option>
            <a-option>否</a-option>
          </a-select>
        </template>
      </a-table>
    </a-card>
  </div>
</template>
<style scoped>
.container {
  width: 800px;
  margin: 0 auto;
}
</style>