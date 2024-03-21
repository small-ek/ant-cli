<script setup>
import {reactive} from "vue";

const form = reactive({
  name: '',
  post: '',
  isRead: false,
});
const handleSubmit = (data) => {
  console.log(data);
};

</script>
<template>
  <div class="container">
    <a-card :style="{ width: '860px',marginTop:'50px' }" title="代码生成" hoverable>
      <a-form :model="form" :style="{ width: '600px' }" @submit="handleSubmit">
        <a-form-item field="name" tooltip="Please enter username" label="Username">
          <a-input
              v-model="form.name"
              placeholder="please enter your username..."
          />
        </a-form-item>
        <a-form-item field="post" label="Post">
          <a-input v-model="form.post" placeholder="please enter your post..." />
        </a-form-item>
        <a-form-item field="isRead">
          <a-checkbox v-model="form.isRead"> I have read the manual </a-checkbox>
        </a-form-item>
        <a-form-item>
          <a-button html-type="submit">Submit</a-button>
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