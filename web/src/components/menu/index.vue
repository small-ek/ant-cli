<script setup>
import logo from '@/assets/logo.png';
</script>

<template>
  <div class="menu-list">
    <a-menu mode="horizontal" theme="dark" :default-selected-keys="['1']">
      <a-menu-item key="0" :style="{ padding: 0, marginRight: '38px' }" disabled>
        <div class="logo-container">
          <img :src="logo"/>
        </div>
      </a-menu-item>
      <a-menu-item key="1">代码生成</a-menu-item>

    </a-menu>
  </div>
</template>

<style scoped lang="less">
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