import Mock from 'mockjs'

// 生成角色数据
const generateRoles = () => {
  const fixedRoles = [
    {
      id: 1,
      name: 'Administrator',
      code: 'ADMIN',
      description: 'System administrator with full access'
    },
    {
      id: 2,
      name: 'User',
      code: 'USER',
      description: 'Regular user with limited access'
    },
    {
      id: 3,
      name: 'Guest',
      code: 'GUEST',
      description: 'Guest user with minimal access'
    }
  ]

  // 生成随机角色
  const randomRoles = Array.from({ length: 7 }, (_, index) => ({
    id: index + 4,
    name: Mock.Random.ctitle(4, 8),
    code: Mock.Random.upper(Mock.Random.word(4, 8)),
    description: Mock.Random.csentence(10, 20)
  }))

  return [...fixedRoles, ...randomRoles]
}

// 生成初始数据
const allRoles = generateRoles()

export default [
  // 获取角色列表
  {
    url: '/api/roles',
    method: 'get',
    response: ({ query }) => {
      const { current = 1, pageSize = 10, keyword = '' } = query
      let filteredRoles = [...allRoles]
      
      if (keyword) {
        const loweredKeyword = keyword.toLowerCase()
        filteredRoles = filteredRoles.filter(role => 
          role.name.toLowerCase().includes(loweredKeyword) ||
          role.code.toLowerCase().includes(loweredKeyword)
        )
      }

      const start = (parseInt(current) - 1) * parseInt(pageSize)
      const end = start + parseInt(pageSize)
      const paginatedRoles = filteredRoles.slice(start, end)

      return {
        code: 200,
        data: {
          list: paginatedRoles,
          total: filteredRoles.length
        }
      }
    }
  },

  // 删除角色
  {
    url: '/api/roles/:id',
    method: 'delete',
    response: (options) => {
      const id = options.url.match(/\/api\/roles\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的角色ID'
        }
      }

      const index = allRoles.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allRoles.splice(index, 1)
        return {
          code: 200,
          message: '删除成功'
        }
      }
      return {
        code: 404,
        message: '角色不存在'
      }
    }
  },

  // 添加角色
  {
    url: '/api/roles',
    method: 'post',
    response: ({ body }) => {
      const newRole = {
        id: allRoles.length + 1,
        ...body
      }
      allRoles.push(newRole)
      return {
        code: 200,
        message: '添加成功',
        data: newRole
      }
    }
  },

  // 更新角色
  {
    url: '/api/roles/:id',
    method: 'put',
    response: ({ body, url }) => {
      const id = url.match(/\/api\/roles\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的角色ID'
        }
      }

      const index = allRoles.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allRoles[index] = {
          ...allRoles[index],
          ...body
        }
        return {
          code: 200,
          message: '更新成功',
          data: allRoles[index]
        }
      }
      return {
        code: 404,
        message: '角色不存在'
      }
    }
  }
] 