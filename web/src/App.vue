<script setup>
import Menu from "../src/components/menu/index.vue"
</script>
<template>
  <div class="bodys">
    <Menu></Menu>
    <router-view></router-view>
  </div>
</template>


<style scoped>
.bodys {
  min-width: 1200px;
  background-color: var(--color-neutral-1);
  height: 100vh;
}
</style>
