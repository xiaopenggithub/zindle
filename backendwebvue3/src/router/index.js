import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
  path: '/',
  redirect: '/login'
},
{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index.vue')
},
{
  path: '/login',
  name: 'Login',
  component: () => import('@/view/login/index.vue')
},
{
  path: '/tmproute',
  name: 'tmproute',
  component: () => import('@/view/login/tmproute.vue')
}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
