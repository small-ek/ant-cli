<script setup>
import logo from '@/assets/logo.png';
import {ref} from "vue";
import {IconClose, IconLanguage,} from '@arco-design/web-vue/es/icon';

const popupVisible1 = ref(false)
const setLang = (index) => {
  if (index === 1) {
    localStorage.setItem('lang', 'zh-CN')
  } else {
    localStorage.setItem('lang', 'en-US')
  }
  location.reload()
}
</script>

<template>
  <div class="menu-list">
    <a-menu mode="horizontal" theme="dark" :default-selected-keys="['1']">
      <a-menu-item key="0" :style="{ padding: 0, marginRight: '38px' }" disabled>
        <div class="logo-container">
          <img :src="logo"/>
        </div>
      </a-menu-item>
      <a-menu-item key="1">{{ $t("code.generation") }}</a-menu-item>

    </a-menu>
  </div>
  <div class="flex-menu">
    <a-trigger
        :trigger="['click', 'hover']"
        clickToClose
        position="top"
        v-model:popupVisible="popupVisible1"
    >
      <div :class="`button-trigger ${popupVisible1 ? 'button-trigger-active' : ''}`">
        <IconClose v-if="popupVisible1"/>
        <IconLanguage v-else style="font-size: 20px"/>
      </div>
      <template #content>
        <a-menu
            :style="{ marginBottom: '-4px' }"
            mode="popButton"
            :tooltipProps="{ position: 'left' }"
            showCollapseButton
        >
          <a-menu-item key="1" @click="setLang(1)">
            <template #icon>中</template>
            中文
          </a-menu-item>
          <a-menu-item key="2" @click="setLang(2)">
            <template #icon>En</template>
            English
          </a-menu-item>
        </a-menu>
      </template>
    </a-trigger>
  </div>
</template>

<style scoped lang="less">
.button-trigger {
  position: absolute;
  bottom: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  color: var(--color-white);
  font-size: 14px;
  border-radius: 50%;
  cursor: pointer;
  transition: all 0.1s;
}

/* button right */
.button-trigger:nth-child(1) {
  right: 30px;
  background-color: rgb(var(--arcoblue-6));
}

.button-trigger:nth-child(1).button-trigger-active {
  background-color: var(--color-primary-light-4);
}

.menu-list {
  box-sizing: border-box;
  width: 100%;
  padding: 0px;
  background-color: var(--color-neutral-2);
}

// logo
.logo-container {
  width: 65px;
  height: 65px;
  border-radius: 15px;
  background-image: linear-gradient(-225deg, #69EACB 0%, #EACCF8 48%, #6654F1 100%);
  cursor: pointer;

  img {
    width: 65px;
    height: 65px;
  }
}
</style>