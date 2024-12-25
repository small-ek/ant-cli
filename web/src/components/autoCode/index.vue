<script setup>
import {defineEmits, nextTick, onMounted, ref, toRefs} from 'vue'
import hljs from 'highlight.js'
import "highlight.js/styles/monokai-sublime.min.css";
import go from 'highlight.js/lib/languages/go.js';
import clipboard from 'clipboard';
import {Message} from "@arco-design/web-vue";
import {useI18n} from "vue-i18n";
const {t} = useI18n()
const activeKey = ref(1)
onMounted(() => {
  nextTick(() => {
    document.querySelectorAll("pre").forEach((block) => {
      hljs.highlightElement(block, {language: go, theme: 'monokai-sublime'});

    });
  })

})

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

let {visible, preCode} = toRefs(props)
console.log(preCode)

const emits = defineEmits(['update:visible'])

const handleCancel = () => {
  emits('update:visible', false)
}

const copyCode = () => {
  clipboard.copy(preCode.value[activeKey.value - 1]['code'])
  Message.info({content: "复制成功", duration: 2000, id: 'copy2'})
}

const generateCode = () => {
  emits('generateCode', preCode.value[activeKey.value - 1]["name"])
}

const onChangeTab = (index) => {
  clipboard.copy(preCode.value[index - 1]['code'])
  Message.info({content: "已复制", duration: 2000, id: 'copy'})
}
</script>

<template>

  <a-modal :width="800" :visible="visible" @cancel="handleCancel" :footer="false" draggable>
    <template #title>
      预览代码
    </template>
    <a-tabs v-model:active-key="activeKey" @change="onChangeTab">
      <a-tab-pane :key="index+1" :title="row['name']" v-for="(row,index) in preCode">
        <pre><code class="language-code" v-html="row['code']"></code></pre>
      </a-tab-pane>

    </a-tabs>
    <div>
      <a-popconfirm :content="t('tips.is_gen_code')" type="info" @ok="generateCode()">
        <a-button type="primary" style="margin-right: 20px" >
          <template #icon>
            <icon-code />
          </template>
          生成当前代码
        </a-button>
      </a-popconfirm>

      <a-button type="primary" style="float: right" @click="copyCode">
        <template #icon>
          <icon-copy/>
        </template>
        复制
      </a-button>
    </div>
  </a-modal>
</template>

<style scoped>
.hljs {
  padding: 10px;
  border-radius: 15px;
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>