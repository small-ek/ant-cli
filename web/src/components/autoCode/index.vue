<script setup>
import {defineEmits, nextTick, onMounted, toRefs} from 'vue'
import hljs from 'highlight.js'
import "highlight.js/styles/monokai-sublime.min.css";
import go from 'highlight.js/lib/languages/go.js';
import javascript from 'highlight.js/lib/languages/javascript.js';

onMounted(() => {
  nextTick(() => {
    document.querySelectorAll("pre").forEach((block) => {
      hljs.highlightElement(block, {language: go, theme: 'monokai-sublime'});
      hljs.highlightElement(block, {language: javascript, theme: 'monokai-sublime'});
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


const emits = defineEmits(['update:visible'])

const handleCancel = () => {
  emits('update:visible', false)
}

const copyCode = () => {
  
}
</script>

<template>

  <a-modal :width="800" :visible="visible" @cancel="handleCancel" :footer="false" draggable>
    <template #title>
      预览代码
    </template>
    <a-tabs default-active-key="2">
      <a-tab-pane key="1" title="Route">
        <pre><code class="language-code">{{ preCode["route"] }}</code></pre>
      </a-tab-pane>
      <a-tab-pane key="2" title="Controller">
        <pre><code class="language-code">{{ preCode["controller"] }}</code></pre>
      </a-tab-pane>
      <a-tab-pane key="3" title="Service">
        <pre><code class="language-code">{{ preCode["service"] }}</code></pre>
      </a-tab-pane>
      <a-tab-pane key="4" title="Dao">
        <pre><code class="language-code">{{ preCode["dao"] }}</code></pre>
      </a-tab-pane>
      <a-tab-pane key="5" title="Model">
        <pre><code class="language-code">{{ preCode["model"] }}</code></pre>
      </a-tab-pane>
    </a-tabs>
    <div>
      <a-button type="primary" @click="copyCode">
        <template #icon>
          <icon-copy />
        </template>
        复制</a-button>
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