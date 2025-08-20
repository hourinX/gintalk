<script setup>
import { ref, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import { CommentOutlined, TeamOutlined } from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()

const activeIcon = ref('');
console.log('Current route name:', route.name)
if (route.name) {
  activeIcon.value = route.name.toLowerCase() === 'chat' ? 'chats' : route.name.toLowerCase()
} else {
  activeIcon.value = 'chats'
}

const icons = [
  { key: 'chats', component: CommentOutlined, route: '/chat' },
  { key: 'relations', component: TeamOutlined, route: '/relations' },
]

const handleIconClick = (key,route) => {
  activeIcon.value = key
  router.push(route)
}

watch(
  () => route.name,
  (newName) => {
    if (newName) {
      activeIcon.value = newName.toLowerCase() === 'chat' ? 'chats' : newName.toLowerCase()
    }
  },
  { immediate: true }
)
</script>

<template>
  <a-layout>
    <a-layout-header class="header">
      <div class="logo no-select">
        <img src="@/assets/logo.svg" alt="Logo" class="logo-img" />
        <span class="logo-text">GinTalk</span>
      </div>
    </a-layout-header>
    <a-layout>
      <a-layout-sider width="50" class="sider" style="background: #fff">
        <div class="icon-menu">
          <component
            v-for="item in icons"
            :is="item.component"
            :key="item.key"
            class="menu-icon"
            :class="{ active: activeIcon === item.key }"
            @click="handleIconClick(item.key,item.route)"
          />
        </div>
      </a-layout-sider>
      <a-layout style="padding: 0 12px 12px">
        <a-layout-content
          :style="{ background: '#fff', padding: '24px', margin: 0, minHeight: '280px' }"
        >
          <router-view />
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-layout>
</template>

<style scoped>
#components-layout-demo-top-side-2 .logo {
  float: left;
  width: 120px;
  height: 32px;
  margin: 16px 24px 16px 0;
  background: rgba(255, 255, 255, 0.3);
  display: flex;
  flex-direction: row;
  align-items: center;
  align-content: center;
  gap: 8px;
}

.no-select {
  user-select: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
}

.ant-row-rtl #components-layout-demo-top-side-2 .logo {
  float: right;
  margin: 16px 0 16px 24px;
}

.site-layout-background {
  background: #fff;
}

.ant-layout-header {
  background: var(--color-primary) !important;
  color: #fff;
  display: flex;
  align-items: center;
  padding: 0 20px;
}

.logo-img {
  height: 36px;
  width: 36px;
  user-drag: none;
  -webkit-user-drag: none;
}

.logo-text {
  font-size: 20px;
  font-weight: 600;
  line-height: 1;
  color: #fff;
  margin-left: 5px;
}

.sider {
  background: #001529;
  display: flex;
  justify-content: center;
  padding: 16px 0;
}

.icon-menu {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 24px;
}

.menu-icon {
  font-size: 28px;
  color: #acacac !important;
  cursor: pointer;
  transition: color 0.3s;
}

.menu-icon:hover {
  color: #595959 !important;
}

.menu-icon.active {
  color: var(--color-primary) !important;
}
</style>