<script setup>
import Menu from "@/components/menu/index.vue"
</script>
<template>
  <div class="bodys">
    <Menu></Menu>
    <router-view></router-view>
  </div>
</template>



<style scoped>
.bodys{
  min-width: 1200px;
}
</style>
