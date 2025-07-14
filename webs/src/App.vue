<script setup>
import zhCN from 'ant-design-vue/es/locale/zh_CN'
const theme = {
  token: {
    'colorPrimary': "#ff0000"
  }
}


</script>

<template>
  <a-config-provider :theme="theme" :locale="zhCN">
    <router-view />
  </a-config-provider>
  
</template>

<style scoped>

</style>
