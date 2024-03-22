<script setup>
import {reactive} from "vue";

const handleSubmit = ({values, errors}) => {
  console.log('values:', values, '\nerrors:', errors)
}

const form = reactive({
  name: '',
  password: '',
  password2: '',
  email: '',
  ip: '192.168.2.1',
  url: '',
  match: ''
});

const rules = {
  name: [
    {
      required: true,
      message:'name is required',
    },
  ],
  password: [
    {
      required: true,
      message:'password is required',
    },
  ],
  password2: [
    {
      required: true,
      message:'password is required',
    },
    {
      validator: (value, cb) => {
        if (value !== form.password) {
          cb('two passwords do not match')
        } else {
          cb()
        }
      }
    }
  ],
  email: [
    {
      type: 'email',
      required: true,
    }
  ],
  ip: [
    {
      type: 'ip',
      required: true,
    }
  ],
  url: [
    {
      type: 'url',
      required: true,
    }
  ],
  match: [
    {
      required: true,
      validator: (value, cb) => {
        return new Promise((resolve) => {
          if (!value) {
            cb('Please enter match')
          }
          if (value !== 'match') {
            cb('match must be match!')
          }
          resolve()
        })
      }
    }
  ],
}

</script>
<template>
  <div class="container">
    <a-card :style="{ width: '860px',marginTop:'50px' }" title="代码生成" hoverable>
      <a-form ref="formRef" :rules="rules" :model="form" :style="{width:'600px'}" @submit="handleSubmit">
        <a-form-item field="dbname" label="数据库名" validate-trigger="blur">
          <a-select v-model="form.db" placeholder="请选择" allow-clear>
            <a-option value="section one">Section One</a-option>
            <a-option value="section two">Section Two</a-option>
            <a-option value="section three">Section Three</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="table" label="表名" validate-trigger="blur">
          <a-select v-model="form.table" placeholder="请选择" allow-clear>
            <a-option value="section one">Section One</a-option>
            <a-option value="section two">Section Two</a-option>
            <a-option value="section three">Section Three</a-option>
          </a-select>
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button html-type="submit">确认</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>
<style scoped>
.container {
  width: 800px;
  margin: 0 auto;
}
</style>