import { createRouter, createWebHistory } from 'vue-router'


const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    name: 'Login',
    path: '/login',
    component: () => import('@/pages/login/index.vue')
  }, {
    name: 'Chat',
    path: '/chat',
    component: () => import('@/pages/chat/index.vue'),
  }

]


const router = createRouter({
  history: createWebHistory(),
  routes,
})
router.beforeEach((to) => {
  const token = localStorage.getItem('token')
  if (to.fullPath !== '/login') {
    return token ? true : { name: 'Login' }
  } else {
    return token ? { name: 'Chat' } : true
  }
})

export default router