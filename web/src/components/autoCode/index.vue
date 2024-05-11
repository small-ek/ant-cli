<script setup>
import {defineEmits, toRefs} from 'vue'
import 'highlight.js/styles/atom-one-dark.css';

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  preCode: {
    type: Object,
    default: {}
  }
})

let {visible,preCode} = toRefs(props)
const emits = defineEmits(['update:visible'])

console.log(preCode)
const handleCancel = () => {
  emits('update:visible', false)
}

</script>

<template>
  <a-modal :width="800" :visible="visible" @cancel="handleCancel" :footer="false" draggable>
    <template #title>
      预览代码
    </template>
    <a-tabs default-active-key="2">
      <a-tab-pane key="1" title="Controller">
        {{preCode["controller"]}}
      </a-tab-pane>
      <a-tab-pane key="2" title="Tab 2">
        Content of Tab Panel 2
      </a-tab-pane>
      <a-tab-pane key="3">
        <template #title>Tab 3</template>
        Content of Tab Panel 3
      </a-tab-pane>
    </a-tabs>
  </a-modal>
</template>

<style scoped>

</style>