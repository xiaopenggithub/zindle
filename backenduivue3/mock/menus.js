import Mock from 'mockjs'

// 生成菜单数据
const generateMenus = () => {
  const fixedMenus = [
    {
      id: 1,
      name: 'Dashboard',
      path: '/dashboard',
      component: 'Dashboard',
      icon: 'IconDashboard',
      sort: 1,
      status: 'active'
    },
    {
      id: 2,
      name: 'System Management',
      path: '/system',
      component: 'Layout',
      icon: 'IconSettings',
      sort: 2,
      status: 'active',
      children: [
        {
          id: 21,
          name: 'User Management',
          path: '/system/users',
          component: 'Users',
          icon: 'IconUser',
          sort: 1,
          status: 'active'
        },
        {
          id: 22,
          name: 'Role Management',
          path: '/system/roles',
          component: 'Roles',
          icon: 'IconUserGroup',
          sort: 2,
          status: 'active'
        },
        {
          id: 23,
          name: 'Permission Management',
          path: '/system/permissions',
          component: 'Permissions',
          icon: 'IconLock',
          sort: 3,
          status: 'active'
        },
        {
          id: 24,
          name: 'Menu Management',
          path: '/system/menus',
          component: 'Menus',
          icon: 'IconMenu',
          sort: 4,
          status: 'active'
        }
      ]
    }
  ]

  // 生成随机菜单
  const randomMenus = Array.from({ length: 3 }, (_, index) => {
    const id = index + 3
    return {
      id,
      name: Mock.Random.ctitle(4, 8),
      path: `/${Mock.Random.word(4, 8)}`,
      component: Mock.Random.capitalize(Mock.Random.word(4, 8)),
      icon: `Icon${Mock.Random.capitalize(Mock.Random.word(4, 8))}`,
      sort: id,
      status: Mock.Random.pick(['active', 'inactive']),
      children: Array.from({ length: Mock.Random.integer(2, 4) }, (_, childIndex) => ({
        id: id * 10 + childIndex + 1,
        name: Mock.Random.ctitle(4, 8),
        path: `/${Mock.Random.word(4, 8)}/${Mock.Random.word(4, 8)}`,
        component: Mock.Random.capitalize(Mock.Random.word(4, 8)),
        icon: `Icon${Mock.Random.capitalize(Mock.Random.word(4, 8))}`,
        sort: childIndex + 1,
        status: Mock.Random.pick(['active', 'inactive'])
      }))
    }
  })

  return [...fixedMenus, ...randomMenus]
}

// 生成初始数据
const allMenus = generateMenus()

export default [
  // 获取菜单列表
  {
    url: '/api/menus',
    method: 'get',
    response: ({ query }) => {
      const { keyword = '' } = query
      let filteredMenus = [...allMenus]
      
      if (keyword) {
        const loweredKeyword = keyword.toLowerCase()
        filteredMenus = filteredMenus.filter(menu => 
          menu.name.toLowerCase().includes(loweredKeyword) ||
          menu.path.toLowerCase().includes(loweredKeyword)
        )
      }

      return {
        code: 200,
        data: filteredMenus
      }
    }
  },

  // 删除菜单
  {
    url: '/api/menus/:id',
    method: 'delete',
    response: (options) => {
      const id = options.url.match(/\/api\/menus\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的菜单ID'
        }
      }

      const index = allMenus.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allMenus.splice(index, 1)
        return {
          code: 200,
          message: '删除成功'
        }
      }
      return {
        code: 404,
        message: '菜单不存在'
      }
    }
  },

  // 添加菜单
  {
    url: '/api/menus',
    method: 'post',
    response: ({ body }) => {
      const newMenu = {
        id: allMenus.length + 1,
        ...body,
        status: body.status || 'active'
      }
      allMenus.push(newMenu)
      return {
        code: 200,
        message: '添加成功',
        data: newMenu
      }
    }
  },

  // 更新菜单
  {
    url: '/api/menus/:id',
    method: 'put',
    response: ({ body, url }) => {
      const id = url.match(/\/api\/menus\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的菜单ID'
        }
      }

      const index = allMenus.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allMenus[index] = {
          ...allMenus[index],
          ...body
        }
        return {
          code: 200,
          message: '更新成功',
          data: allMenus[index]
        }
      }
      return {
        code: 404,
        message: '菜单不存在'
      }
    }
  }
] 