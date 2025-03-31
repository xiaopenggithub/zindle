import { createRouter, createWebHashHistory } from 'vue-router'
import baseRoutes, { addDynamicRoutes } from './routes'
import { getToken, removeToken } from '@/utils/auth'
import { Message } from '@arco-design/web-vue'
import { useStore } from '@/store'

const router = createRouter({
  history: createWebHashHistory(),
  routes: baseRoutes
})

// 记录是否已添加动态路由
let hasAddedRoutes = false

// 路由守卫
router.beforeEach(async (to, from, next) => {
  const token = getToken()
  
  // 如果已登录且访问登录页，确保动态路由已加载后再跳转到上次访问的页面或首页
  if (token && (to.path === '/login' || to.path === '/register')) {
    if (!hasAddedRoutes) {
      try {
        const store = useStore()
        const success = await store.getUserInfo()
        if (success) {
          addDynamicRoutes(router)
          hasAddedRoutes = true
          // 获取上次访问的页面
          const lastPath = localStorage.getItem('lastPath')
          // 如果有上次访问的页面且不是登录相关页面，则跳转到该页面
          if (lastPath && !['/login', '/register', '/'].includes(lastPath)) {
            next(lastPath)
          } else {
            next('/dashboard')
          }
          return
        } else {
          removeToken()
          next()
          return
        }
      } catch (error) {
        console.error('Error getting user info:', error)
        removeToken()
        next()
        return
      }
    } else {
      // 获取上次访问的页面
      const lastPath = localStorage.getItem('lastPath')
      // 如果有上次访问的页面且不是登录相关页面，则跳转到该页面
      if (lastPath && !['/login', '/register', '/'].includes(lastPath)) {
        next(lastPath)
      } else {
        next('/dashboard')
      }
      return
    }
  }
  
  // 需要登录的页面
  if (to.meta.requiresAuth === true) {
    if (!token) {
      Message.error('请先登录')
      // 保存尝试访问的页面路径
      if (to.path !== '/') {
        localStorage.setItem('lastPath', to.fullPath)
      }
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
      return
    }

    // 如果已登录但还没有添加动态路由
    if (!hasAddedRoutes) {
      try {
        const store = useStore()
        // 获取用户信息和菜单，并添加动态路由
        const success = await store.getUserInfo()
        if (success) {
          addDynamicRoutes(router)
          hasAddedRoutes = true
          // 重新进入当前路由
          next({ ...to, replace: true })
          return
        } else {
          Message.error('获取用户信息失败')
          next('/login')
          return
        }
      } catch (error) {
        console.error('Error getting user info:', error)
        Message.error('获取用户信息失败')
        next('/login')
        return
      }
    }
    
    // 检查角色权限
    if (to.meta.roles) {
      const userRoles = JSON.parse(localStorage.getItem('userRoles') || '[]')
      if (!to.meta.roles.some(role => userRoles.includes(role))) {
        Message.error('没有访问权限')
        next('/403')
        return
      }
    }
  }
  
  // 记录最后访问的页面路径（排除登录、注册等页面和根路径）
  if (!['Login', 'Register', 'NotFound'].includes(to.name) && to.path !== '/') {
    localStorage.setItem('lastPath', to.fullPath)
  }
  
  next()
})

// 重置路由状态
export function resetRouter() {
  hasAddedRoutes = false
  // 移除所有动态添加的路由
  router.getRoutes().forEach(route => {
    if (route.name && !['Layout', 'Login', 'Register', 'NotFound', 'Profile'].includes(route.name)) {
      router.removeRoute(route.name)
    }
  })
}

export default router 