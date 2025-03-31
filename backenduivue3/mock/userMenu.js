import Mock from 'mockjs'

// 生成固定菜单数据
const generateFixedMenus = () => {
  return [
    {
      id: 1,
      name: '仪表盘',
      path: '/dashboard',
      component: 'dashboard/Index',
      icon: 'IconDashboard',
      sort: 1,
      status: 'active'
    },
    {
      id: 2,
      name: '会员管理',
      path: '/members',
      component: 'members/Index',
      icon: 'IconUserGroup',
      sort: 2,
      status: 'active'
    },
    {
      id: 3,
      name: '订单管理',
      path: '/orders',
      component: 'Layout',
      icon: 'IconOrderedList',
      sort: 3,
      status: 'active',
      children: [
        {
          id: 31,
          name: '平台订单',
          path: '/orders/manage',
          component: 'orders/Index',
          icon: 'IconShoppingCart',
          sort: 1,
          status: 'active'
        },
        {
          id: 32,
          name: '我的订单',
          path: '/orders/my',
          component: 'memberOrders/Index',
          icon: 'IconFile',
          sort: 2,
          status: 'active'
        }
      ]
    },
    {
      id: 4,
      name: '产品管理',
      path: '/products',
      component: 'Layout',
      icon: 'IconApps',
      sort: 4,
      status: 'active',
      children: [
        {
          id: 41,
          name: '产品分类',
          path: '/products/category',
          component: 'products/Category',
          icon: 'IconTag',
          sort: 1,
          status: 'active'
        },
        {
          id: 42,
          name: '产品列表',
          path: '/products/list',
          component: 'products/List',
          icon: 'IconUnorderedList',
          sort: 2,
          status: 'active'
        }
      ]
    },
    {
      id: 5,
      name: '系统管理',
      path: '/system',
      component: 'Layout',
      icon: 'IconSettings',
      sort: 5,
      status: 'active',
      children: [
        {
          id: 51,
          name: '用户管理',
          path: '/system/users',
          component: 'system/users/index',
          icon: 'IconUser',
          sort: 1,
          status: 'active'
        },
        {
          id: 52,
          name: '角色管理',
          path: '/system/roles',
          component: 'system/roles/index',
          icon: 'IconTeam',
          sort: 2,
          status: 'active'
        },
        {
          id: 53,
          name: '权限管理',
          path: '/system/permissions',
          component: 'system/permissions/index',
          icon: 'IconLock',
          sort: 3,
          status: 'active'
        },
        {
          id: 54,
          name: '菜单管理',
          path: '/system/menus',
          component: 'system/menus/index',
          icon: 'IconMenu',
          sort: 4,
          status: 'active'
        },
        {
          id: 55,
          name: '系统日志',
          path: '/system/logs',
          component: 'system/logs/index',
          icon: 'IconHistory',
          sort: 5,
          status: 'active'
        }
      ]
    }
  ]
}

// 生成初始数据
const authMenus = generateFixedMenus()

// 根据角色过滤菜单
const filterMenusByRole = (menus, role) => {
  // 这里可以根据角色来过滤菜单
  // 示例中先返回所有菜单
  return menus
}

export default [
  // 获取用户菜单
  {
    url: '/api/user/menus',
    method: 'get',
    response: ({ headers }) => {
      const role = headers['x-role'] || 'admin'
      const filteredMenus = filterMenusByRole(authMenus, role)
      
      return {
        code: 200,
        data: {
          list: filteredMenus,
          total: filteredMenus.length
        },
        message: '获取菜单成功'
      }
    }
  }
] 