import Mock from 'mockjs'

// 生成权限数据
const generatePermissions = () => {
  const fixedPermissions = [
    {
      id: 1,
      name: 'User Management',
      code: 'user:manage',
      type: 'Menu',
      description: 'Access to user management module'
    },
    {
      id: 2,
      name: 'Create User',
      code: 'user:create',
      type: 'Button',
      description: 'Create new users'
    },
    {
      id: 3,
      name: 'Delete User',
      code: 'user:delete',
      type: 'Button',
      description: 'Delete existing users'
    },
    {
      id: 4,
      name: 'Role Management',
      code: 'role:manage',
      type: 'Menu',
      description: 'Access to role management module'
    },
    {
      id: 5,
      name: 'Permission Management',
      code: 'permission:manage',
      type: 'Menu',
      description: 'Access to permission management module'
    }
  ]

  // 生成随机权限
  const randomPermissions = Array.from({ length: 15 }, (_, index) => ({
    id: index + 6,
    name: Mock.Random.ctitle(4, 8),
    code: `${Mock.Random.word(4, 8)}:${Mock.Random.word(4, 8)}`,
    type: Mock.Random.pick(['Menu', 'Button', 'API']),
    description: Mock.Random.csentence(10, 20)
  }))

  return [...fixedPermissions, ...randomPermissions]
}

// 生成初始数据
const allPermissions = generatePermissions()

export default [
  // 获取权限列表
  {
    url: '/api/permissions',
    method: 'get',
    response: ({ query }) => {
      const { current = 1, pageSize = 10, keyword = '' } = query
      let filteredPermissions = [...allPermissions]
      
      if (keyword) {
        const loweredKeyword = keyword.toLowerCase()
        filteredPermissions = filteredPermissions.filter(permission => 
          permission.name.toLowerCase().includes(loweredKeyword) ||
          permission.code.toLowerCase().includes(loweredKeyword) ||
          permission.type.toLowerCase().includes(loweredKeyword)
        )
      }

      const start = (parseInt(current) - 1) * parseInt(pageSize)
      const end = start + parseInt(pageSize)
      const paginatedPermissions = filteredPermissions.slice(start, end)

      return {
        code: 200,
        data: {
          list: paginatedPermissions,
          total: filteredPermissions.length
        }
      }
    }
  },

  // 删除权限
  {
    url: '/api/permissions/:id',
    method: 'delete',
    response: (options) => {
      const id = options.url.match(/\/api\/permissions\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的权限ID'
        }
      }

      const index = allPermissions.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allPermissions.splice(index, 1)
        return {
          code: 200,
          message: '删除成功'
        }
      }
      return {
        code: 404,
        message: '权限不存在'
      }
    }
  },

  // 添加权限
  {
    url: '/api/permissions',
    method: 'post',
    response: ({ body }) => {
      const newPermission = {
        id: allPermissions.length + 1,
        ...body
      }
      allPermissions.push(newPermission)
      return {
        code: 200,
        message: '添加成功',
        data: newPermission
      }
    }
  },

  // 更新权限
  {
    url: '/api/permissions/:id',
    method: 'put',
    response: ({ body, url }) => {
      const id = url.match(/\/api\/permissions\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的权限ID'
        }
      }

      const index = allPermissions.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allPermissions[index] = {
          ...allPermissions[index],
          ...body
        }
        return {
          code: 200,
          message: '更新成功',
          data: allPermissions[index]
        }
      }
      return {
        code: 404,
        message: '权限不存在'
      }
    }
  }
] 