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
  if (to.meta.requiresAuth === false) {
    next()
  } else {
    const token = localStorage.getItem('token')
    console.log('token', token)
    if (!token) {
      next({ name: 'Login' })
      NProgress.done()
    } else {
      next()
    }
  }
})

router.afterEach((to) => {
  NProgress.done()
  document.title = `${to.meta.title || import.meta.env.BASE_GINTALK_NAME}`
})

export default router
