package web

import (
	"bytes"
	"github.com/small-ek/ant-cli/utils"
	"text/template"
)

func Views(table string) (string, error) {
	const tpl = `<script setup>
import FilterBar from "@/components/filter-bar/index.vue";
import Table from "@/components/table/index.vue";
import {onMounted} from "vue";
import {Message} from "@arco-design/web-vue";
import {formatTime} from "@/utils/time.js";
import EditForm from "@/components/edit-form/index.vue";
{{.importsApi}}

import {
  columns,
  formData,
  formList,
  formListIndexMap,
  formRef,
  formRules,
  formTitle,
  ids,
  list,
  onSelect,
  page,
  searchList,
  tableRef
} from "./index.js";

import {useNavigation} from "@/utils/base.js";
import {useRoute, useRouter} from "vue-router";
const router = useRouter();
const {jump} = useNavigation(router);
const route = useRoute()
const fetchPageList = async (params) => {
  const res = await {{.importsApiName}}().getList(params.currentPage || 1, params.pageSize || page.value.pageSize, params.filter_map || page.value.searchForm, params.order || page.value.order, params.desc || page.value.desc);
  list.value = res.data.items;
  page.value.current = params.currentPage;
  page.value.total = res.data.total;
};

const search = (row) => {
  page.value.searchForm = row;
  page.value.current = 1;
  fetchPageList({
    currentPage: 1
  });
};

onMounted(() => {
  fetchPageList({
    currentPage: page.value.current,
    pageSize: page.value.pageSize,
    order: page.value.order,
    desc: page.value.desc
  });
});

const changePage = (current) => {
  fetchPageList({
    currentPage: current
  });
};

const pageSizeChange = (size) => {
  page.value.pageSize = size;
  fetchPageList();
};

const sorterChange = (field, sort) => {
  if (sort && field) {
    page.value.order = [field];
    page.value.desc = [sort === 'descend'];
    page.value.current = 1
    reload()
  }
};

const updatedStatus = (status, row) => {
  row.status = status;
  updates(row).then((res) => {
    if (res.code === 0) {
      Message.success('操作成功');
    } else {
      Message.error('操作失败');
      row.status = status === 2 ? 1 : 2;
    }
  });
};


const showEdit = (row) => {
  // 如果是编辑，密码默认隐藏
  formTitle.value = row ? '编辑' : '添加';
  if (row) {
    {{.importsApiName}}().show(row.id).then((res) => {
      if (res.code === 0) {
        formData.value = res.data;
        formRef.value.setVisible(true);
      }
    });
  } else {
    formData.value = {};
    formRef.value.setVisible(true);
  }

};

const submit = (row) => {
  const action = row.id ? {{.importsApiName}}().updates : {{.importsApiName}}().creates;
  action(row).then((res) => {
    if (res.code === 0) {
     Message.success('操作成功');
        formRef.value ? formRef.value.setVisible(false) : null;
        reload();
    }
  });
};

const deletesItem = (id) => {
  if (id > 0) {
    ids.value.push(id);
  }
  if (ids.value.length === 0) {
    Message.warning('请选择要删除的项目');
    return;
  }
  {{.importsApiName}}().deletes(ids.value).then((res) => {
    if (res.code === 0) {
      Message.success('操作成功');
      ids.value = [];
      tableRef.value.clearSelected();
      reload();
    }
  });
};
// 重新加载
const reload = () => {
  fetchPageList({
    currentPage: page.value.current,
    pageSize: page.value.pageSize,
    filter_map: page.value.searchForm
  });
};
</script>

<template>
  <div class="container">
    <!--过滤栏-->
    <FilterBar v-if="searchList.length>0" :model="searchList" @search="search">
    </FilterBar>

    <!-- 表格-->
    <div class="ant-card ant-table-card">
      <Table ref="tableRef" :columns="columns" :data="list" :total="page.total" :current="page.current"
             :pageSize="page.pageSize" @changePage="changePage"
             @pageSizeChange="pageSizeChange"
             @sorterChange="sorterChange" @select="onSelect" @deletes="deletesItem" @showEdit="showEdit">
        <template #created_at="{ record }">
			{{.timeFormat}}
        </template>
        <template #optional="{ record }">
          <a-button size="mini" @click="showEdit(record)">编辑</a-button>

          <a-popconfirm content="确认要删除当前项目吗?" okText="确认删除" cancelText="取消"
                        @ok="deletesItem(record.id)">
            <a-button size="mini" type="outline" class="ml-10" status="danger">
              删除
            </a-button>
          </a-popconfirm>
        </template>
      </Table>
    </div>
    <!-- 编辑表单-->
    <EditForm :title="formTitle" ref="formRef" :model="formList" :form="formData" :rules="formRules" @submit="submit">
    </EditForm>
  </div>

</template>

<style scoped>

</style>`
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
