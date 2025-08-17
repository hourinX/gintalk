import { createRouter, createWebHashHistory } from 'vue-router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

NProgress.configure({
  showSpinner: false
})

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      meta: {
        title: 'GinTalk',
      },
      component:() => import('@/views/home/index.vue')
    },
    {
      path: '/login',
      name: 'Login',
      meta: {
        title: 'Welcome to GinTalk',
        requiresAuth: false,
      },
      component:() => import('@/views/login/index.vue')
    },
  ],
})

router.beforeEach((to,from,next) => {
  NProgress.start()

  const accessToken = localStorage.getItem('access_token')
  const expireTime = localStorage.getItem('expire_time')
  const refreshToken = localStorage.getItem('refresh_token')
  const isLoggedIn = accessToken && expireTime && refreshToken && isTokenValid()

  if (to.meta.requiresAuth === false) {
    if (to.name === 'Login' && isLoggedIn) {
      next({ name: 'Home' })
      NProgress.done()
    } else {
      next()
    }
  } else {
    if (isLoggedIn) {
      next()
    } else {
      message.warning('Token has expired, please re-login first...')
      clearAuth()
      next({ name: 'Login' })
      NProgress.done()
    }
  }
})

const isTokenValid = () => {
  const token = localStorage.getItem('access_token')
  const expireTime = localStorage.getItem('expire_time')

  if (!token || !expireTime) return false

  const now = Date.now() / 1000
  return now < expireTime
}

const clearAuth = () => {
  localStorage.removeItem('access_token')
  localStorage.removeItem('expire_time')
  localStorage.removeItem('refresh_token')
}

router.afterEach((to) => {
  NProgress.done()
  document.title = `${to.meta.title || import.meta.env.BASE_GINTALK_NAME}`
})

export default router
