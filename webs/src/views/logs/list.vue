<script setup>
import { ref } from 'vue';
import LoginLogs from './components/login-logs.vue';

const activeKey = ref(1);

</script>

<template>
  <a-tabs v-model:activeKey="activeKey">
    <a-tab-pane :key="1" tab="登录日志">
        <login-logs />
    </a-tab-pane>
    <!-- <a-tab-pane :key="2" tab="Tab 2" force-render>Content of Tab Pane 2</a-tab-pane> -->
    <!-- <a-tab-pane :key="3" tab="Tab 3">Content of Tab Pane 3</a-tab-pane> -->
  </a-tabs>
</template>

<style scoped>

</style>