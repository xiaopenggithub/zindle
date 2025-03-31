import { useStore } from '@/store'

// 将菜单数据转换为路由配置
function generateRoutes(menus) {
  return menus.map(menu => {
    
    // 处理路径，确保子路由路径正确
    const routePath = menu.path.replace(/^\//, '')  // 移除开头的斜杠
    
    // 基础路由配置
    const route = {
      path: routePath,
      name: menu.name,
      meta: {
        title: menu.name,
        icon: menu.icon,
        requiresAuth: true
      }
    }

    // 处理组件和子路由
    if (menu.children && menu.children.length > 0) {
      
      // 生成子路由
      route.children = menu.children.map(child => {
        // 子路由的路径应该是相对于父路由的
        const childPath = child.path.replace(/^\//, '').replace(new RegExp(`^${routePath}/`), '')
        
        
        return {
          path: childPath,
          name: child.name,
          component: loadComponent(child.component),
          meta: {
            title: child.name,
            icon: child.icon,
            requiresAuth: true
          }
        }
      })
      
      // 父级路由使用空组件，但保持router-view
      route.component = () => import('@/layout/SubLayout.vue')
      
      // 添加重定向到第一个子菜单
      if (route.children.length > 0) {
        const firstChild = route.children[0]
        route.redirect = firstChild.path
      }
    } else {
      // 如果没有子菜单，直接加载指定的组件
      route.component = loadComponent(menu.component)
    }
    return route
  })
}

// 动态加载组件
function loadComponent(component) {
  
  // 如果是Layout组件，直接返回布局组件
  if (component === 'Layout') {
    return () => import('@/layout/Index.vue')
  }
  
  // 处理组件路径
  const componentPath = component.replace(/\\/g, '/').replace(/^\//, '')  // 替换反斜杠为正斜杠并移除开头的斜杠
  
  // 构建动态导入路径
  const modules = import.meta.glob('@/views/**/*.vue')
  
  // 尝试不同的路径格式
  const possiblePaths = [
    `/src/views/${componentPath}.vue`,
    `/src/views/${componentPath}/Index.vue`,
    `/src/views/${componentPath}/index.vue`
  ].map(path => path.replace(/\/+/g, '/')) // 规范化路径，移除多余的斜杠
  
  
  // 查找匹配的模块
  for (const path of possiblePaths) {
    const matchedModule = modules[path]
    if (matchedModule) {
      return matchedModule
    }
  }
  
  // 如果都没找到，返回404
  console.error('No matching module found for component:', component, 'Tried paths:', possiblePaths)
  return () => import('@/views/error/404.vue')
}

// 基础路由配置
const baseRoutes = [
  {
    path: '/',
    name: 'Layout',
    component: () => import('@/layout/Index.vue'),
    redirect: () => {
      // 获取上次访问的页面，如果没有则跳转到 dashboard
      const lastPath = localStorage.getItem('lastPath')
      return lastPath || '/dashboard'
    },
    children: [
      {
        path: '/profile',
        name: 'Profile',
        component: () => import('@/views/profile/Index.vue'),
        meta: { 
          requiresAuth: true,
          title: 'common.profile'
        }
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/auth/Login.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/auth/Register.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/forgotPassword',
    name: 'ForgotPassword',
    component: () => import('@/views/auth/ForgotPassword.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/error/404.vue')
  }
]

// 导出路由配置
export default baseRoutes

// 导出添加动态路由的方法
export function addDynamicRoutes(router) {
  const store = useStore()
  const userMenus = store.userMenus
  
  // 如果有菜单数据，生成动态路由
  if (userMenus && userMenus.length > 0) {
    try {
      // 先移除所有动态添加的路由，但保留基础路由
      router.getRoutes().forEach(route => {
        if (route.name && 
            !['Layout', 'Login', 'Register', 'NotFound', 'Profile', 'ForgotPassword'].includes(route.name)) {
          router.removeRoute(route.name)
        }
      })

      const dynamicRoutes = generateRoutes(userMenus)      
      // 找到根路由
      const mainRoute = router.getRoutes().find(route => route.path === '/')
      if (mainRoute) {
        // 直接添加到根路由的children中
        dynamicRoutes.forEach(route => {
          router.addRoute('Layout', route)
        })
      }
    } catch (error) {
      console.error('Error adding dynamic routes:', error)
    }
  } else {
    console.warn('No user menus available')
  }
}