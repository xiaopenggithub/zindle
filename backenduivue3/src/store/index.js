import { defineStore } from 'pinia'
import { setToken, removeToken } from '@/utils/auth'
import { getUserMenus } from '@/api/mock/userMenu'
import { addDynamicRoutes } from '@/router/routes'
import router from '@/router'

export const useStore = defineStore('main', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('userInfo')) || null,
    token: null,
    isAuthenticated: !!localStorage.getItem('token'),
    theme: localStorage.getItem('theme') || 'light',
    userMenus: JSON.parse(localStorage.getItem('userMenus')) || [],
    users: [
      {
        id: 1,
        username: 'admin',
        password: 'admin123',
        email: 'admin@example.com',
        role: 'admin',
        avatar: 'https://avatars.githubusercontent.com/u/1?v=4',
        status: 1,
        createdAt: '2024-01-01'
      }
    ],
    roles: [
      {
        id: 1,
        name: '管理员',
        code: 'admin',
        description: '系统管理员',
        status: 1,
        createdAt: '2024-01-01'
      }
    ],
    permissions: [
      {
        id: 1,
        name: '系统管理',
        code: 'system:manage',
        type: 'menu',
        status: 1,
        createdAt: '2024-01-01'
      }
    ],
    menus: [
      {
        id: 1,
        name: '仪表盘',
        path: '/dashboard',
        component: 'Dashboard',
        icon: 'dashboard',
        sort: 1,
        status: 1
      }
    ],
    logs: [
      {
        id: 1,
        type: 'login',
        user: 'admin',
        action: '用户登录',
        ip: '127.0.0.1',
        createdAt: '2024-01-01 12:00:00'
      }
    ]
  }),

  actions: {
    // 获取用户信息和菜单
    async getUserInfo() {
      try {
        // 从localStorage获取用户信息
        const userInfo = localStorage.getItem('userInfo')
        if (userInfo) {
          this.user = JSON.parse(userInfo)
          this.isAuthenticated = true
        }

        // 获取用户菜单
        const response = await getUserMenus()
        if (response && response.data && response.data.data && response.data.data.list) {
          this.userMenus = response.data.data.list
          localStorage.setItem('userMenus', JSON.stringify(response.data.data.list))
          // 添加动态路由
          addDynamicRoutes(router)
          return true
        }
        return false
      } catch (error) {
        console.error('Failed to get user info:', error)
        return false
      }
    },

    async login(username, password) {
      return new Promise(async (resolve, reject) => {
        try {
          const user = this.users.find(
            u => u.username === username && u.password === password
          )
          
          if (user) {
            // 生成 token
            const token = btoa(username + ':' + new Date().getTime())
            
            // 更新状态
            this.user = user
            this.token = token
            this.isAuthenticated = true
            
            // 保存到 localStorage
            setToken(token)
            localStorage.setItem('userRoles', JSON.stringify([user.role]))
            localStorage.setItem('userInfo', JSON.stringify(user))
            
            try {
              // 获取用户菜单
              const response = await getUserMenus()              
              if (response && response.data && response.data.data && response.data.data.list) {
                this.userMenus = response.data.data.list
                localStorage.setItem('userMenus', JSON.stringify(response.data.data.list))
                // 添加动态路由
                addDynamicRoutes(router)
              } else {
                this.userMenus = []
                console.warn('No menu data received')
              }
            } catch (error) {
              console.error('Failed to fetch user menus:', error)
              this.userMenus = []
            }
            
            // 添加登录日志
            this.addLog({
              type: 'login',
              user: username,
              action: '用户登录',
              ip: '127.0.0.1'
            })
            
            resolve(true)
          } else {
            reject(new Error('用户名或密码错误'))
          }
        } catch (error) {
          console.error('Login error:', error)
          reject(error)
        }
      })
    },

    logout() {
      // 清除状态
      this.user = null
      this.token = null
      this.isAuthenticated = false
      this.userMenus = []
      
      // 清除本地存储
      removeToken()
      localStorage.removeItem('userRoles')
      localStorage.removeItem('userInfo')
      localStorage.removeItem('userMenus') // Clear menu data on logout
      
      // 添加登出日志
      this.addLog({
        type: 'logout',
        user: this.user?.username || 'unknown',
        action: '用户登出',
        ip: '127.0.0.1'
      })
    },

    register(userData) {
      const newUser = {
        id: this.users.length + 1,
        ...userData,
        status: 1,
        createdAt: new Date().toISOString().split('T')[0]
      }
      this.users.push(newUser)
    },

    // 用户管理
    getUsers() {
      return this.users
    },

    addUser(user) {
      this.users.push({
        id: this.users.length + 1,
        ...user,
        createdAt: new Date().toISOString().split('T')[0]
      })
    },

    updateUser(id, user) {
      const index = this.users.findIndex(u => u.id === id)
      if (index !== -1) {
        this.users[index] = { ...this.users[index], ...user }
      }
    },

    deleteUser(id) {
      const index = this.users.findIndex(u => u.id === id)
      if (index !== -1) {
        this.users.splice(index, 1)
      }
    },

    // 角色管理
    getRoles() {
      return this.roles
    },

    addRole(role) {
      this.roles.push({
        id: this.roles.length + 1,
        ...role,
        createdAt: new Date().toISOString().split('T')[0]
      })
    },

    updateRole(id, role) {
      const index = this.roles.findIndex(r => r.id === id)
      if (index !== -1) {
        this.roles[index] = { ...this.roles[index], ...role }
      }
    },

    deleteRole(id) {
      const index = this.roles.findIndex(r => r.id === id)
      if (index !== -1) {
        this.roles.splice(index, 1)
      }
    },

    // 权限管理
    getPermissions() {
      return this.permissions
    },

    addPermission(permission) {
      this.permissions.push({
        id: this.permissions.length + 1,
        ...permission,
        createdAt: new Date().toISOString().split('T')[0]
      })
    },

    updatePermission(id, permission) {
      const index = this.permissions.findIndex(p => p.id === id)
      if (index !== -1) {
        this.permissions[index] = { ...this.permissions[index], ...permission }
      }
    },

    deletePermission(id) {
      const index = this.permissions.findIndex(p => p.id === id)
      if (index !== -1) {
        this.permissions.splice(index, 1)
      }
    },

    // 菜单管理
    getMenus() {
      return this.menus
    },

    addMenu(menu) {
      this.menus.push({
        id: this.menus.length + 1,
        ...menu,
        createdAt: new Date().toISOString().split('T')[0]
      })
    },

    updateMenu(id, menu) {
      const index = this.menus.findIndex(m => m.id === id)
      if (index !== -1) {
        this.menus[index] = { ...this.menus[index], ...menu }
      }
    },

    deleteMenu(id) {
      const index = this.menus.findIndex(m => m.id === id)
      if (index !== -1) {
        this.menus.splice(index, 1)
      }
    },

    // 日志管理
    getLogs() {
      return this.logs
    },

    addLog(log) {
      this.logs.push({
        id: this.logs.length + 1,
        ...log,
        createdAt: new Date().toISOString()
      })
    },

    // Theme management
    setTheme(theme) {
      this.theme = theme
      localStorage.setItem('theme', theme)
      document.body.setAttribute('arco-theme', theme)
      // 更新 CSS 变量
      if (theme === 'dark') {
        document.body.setAttribute('arco-theme', 'dark')
      } else {
        document.body.removeAttribute('arco-theme')
      }
    }
  }
}) 