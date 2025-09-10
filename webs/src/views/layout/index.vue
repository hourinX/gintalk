<script setup>
import { ref, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import { message,Modal } from 'ant-design-vue'
import { CommentOutlined, TeamOutlined, FundProjectionScreenOutlined, LogoutOutlined } from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()

const activeIcon = ref('');
if (route.name) {
  activeIcon.value = route.name.toLowerCase() === 'chat' ? 'chats' : route.name.toLowerCase()
} else {
  activeIcon.value = 'chats'
}

const icons = [
  { key: 'chats', component: CommentOutlined, route: '/chat', title: '聊天' },
  { key: 'relations', component: TeamOutlined, route: '/relations', title: '联系人' },
  { key: 'logs', component: FundProjectionScreenOutlined, route: '/logs', title: '日志' },
  { key: 'logout', component: LogoutOutlined, route: '/logout', title: '退出登录' },
]

const handleIconClick = (key,route) => {
  if (key === 'logout') {
    Modal.confirm({
        title: '退出提醒',
        content: `是否确定登出当前系统？`,
        okText: '登出',
        cancelText: '取消',
        onOk() {
          localStorage.removeItem('access_token')
          localStorage.removeItem('refresh_token')
          localStorage.removeItem('expire_time')
          router.push('/login')
          message.info('have logged out successfully');
        },
        onCancel() {},
    })
    return
  }
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
  <a-layout class="main-layout">
    <a-layout-header class="header">
      <div class="logo no-select">
        <img src="@/assets/logo.svg" alt="Logo" class="logo-img" />
        <span class="logo-text">GinTalk</span>
      </div>
    </a-layout-header>

    <a-layout class="body-layout">
      <a-layout-sider width="50" class="sider" style="background: #fff">
        <div class="icon-menu">
          <a-tooltip
            v-for="item in icons"
            :key="item.key"
            :title="item.title"
            placement="rightTop"
          >
            <component
              :is="item.component"
              :key="item.key"
              class="menu-icon"
              :class="{ active: activeIcon === item.key }"
              @click="handleIconClick(item.key,item.route)"
            />
          </a-tooltip>
        </div>
      </a-layout-sider>

      <a-layout style="padding: 0 2px 2px">
        <a-layout-content class="content-layout">
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

.main-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.body-layout {
  flex: 1;
  display: flex;
}

.content-layout {
  padding: 20px;
  margin: 0;
  flex: 1;
  background: #fff;
  min-height: 0;
  overflow: auto;
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
  flex-direction: column;
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