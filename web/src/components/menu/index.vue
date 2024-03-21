<script setup>
import logo from '@/assets/logo.png';
</script>

<template>
  <div class="menu-demo">
    <a-menu mode="horizontal" theme="dark" :default-selected-keys="['1']">
      <a-menu-item key="0" :style="{ padding: 0, marginRight: '38px' }" disabled>
        <div
            :style="{
            width: '60px',
            height: '60px',
            borderRadius: '2px',
            background: 'var(--color-fill-3)',
            cursor: 'text',
          }"
        >
          <img
              width="60px"
              height="60px"
              :src="logo"
          />
        </div>
      </a-menu-item>
      <a-menu-item key="1">代码生成</a-menu-item>

    </a-menu>
  </div>
</template>

<style scoped>
.menu-demo {
  box-sizing: border-box;
  width: 100%;
  padding: 0px;
  background-color: var(--color-neutral-2);
}
</style>