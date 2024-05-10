<script setup>
import {ref} from "vue";
const props = defineProps({
  data:{
    type: Object,
    default() {
      return {}
    }
  }
})
console.log(props.data)
const visible = ref(true)
console.log(visible)
</script>

<template>
  <div>212121</div>
  <a-modal v-model:visible="visible" :footer="false" draggable>
    <template #title>
      编辑字段
    </template>
    <div>
      11221
    </div>
  </a-modal>
</template>

<style scoped>

</style>