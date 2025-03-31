import Mock from 'mockjs'

// 生成用户数据
const generateUsers = () => {
  const fixedUsers = [
    {
      id: 1,
      username: 'admin',
      email: 'admin@example.com',
      role: 'admin',
      status: 'active',
      createdAt: '2024-03-28 10:00:00'
    },
    {
      id: 2,
      username: 'user1',
      email: 'user1@example.com',
      role: 'user',
      status: 'active',
      createdAt: '2024-03-28 09:00:00'
    },
    {
      id: 3,
      username: 'guest1',
      email: 'guest1@example.com',
      role: 'guest',
      status: 'inactive',
      createdAt: '2024-03-28 08:00:00'
    }
  ]

  // 生成随机用户
  const randomUsers = Array.from({ length: 47 }, (_, index) => ({
    id: index + 4,
    username: Mock.Random.word(5, 10),
    email: Mock.Random.email(),
    role: Mock.Random.pick(['admin', 'user', 'guest']),
    status: Mock.Random.pick(['active', 'inactive']),
    createdAt: Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
  }))

  return [...fixedUsers, ...randomUsers]
}

// 生成初始数据
const allUsers = generateUsers()

export default [
  // 获取用户列表
  {
    url: '/api/users',
    method: 'get',
    response: ({ query }) => {
      const { current = 1, pageSize = 10, keyword = '', role = '' } = query
      let filteredUsers = [...allUsers]
      
      if (keyword) {
        const loweredKeyword = keyword.toLowerCase()
        filteredUsers = filteredUsers.filter(user => 
          user.username.toLowerCase().includes(loweredKeyword) ||
          user.email.toLowerCase().includes(loweredKeyword)
        )
      }

      if (role) {
        filteredUsers = filteredUsers.filter(user => user.role === role)
      }

      const start = (parseInt(current) - 1) * parseInt(pageSize)
      const end = start + parseInt(pageSize)
      const paginatedUsers = filteredUsers.slice(start, end)

      return {
        code: 200,
        data: {
          list: paginatedUsers,
          total: filteredUsers.length
        }
      }
    }
  },

  // 删除用户
  {
    url: '/api/users/:id',
    method: 'delete',
    response: (options) => {
      const id = options.url.match(/\/api\/users\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的用户ID'
        }
      }

      const index = allUsers.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allUsers.splice(index, 1)
        return {
          code: 200,
          message: '删除成功'
        }
      }
      return {
        code: 404,
        message: '用户不存在'
      }
    }
  },

  // 添加用户
  {
    url: '/api/users',
    method: 'post',
    response: ({ body }) => {
      const newUser = {
        id: allUsers.length + 1,
        ...body,
        createdAt: Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
      }
      allUsers.push(newUser)
      return {
        code: 200,
        message: '添加成功',
        data: newUser
      }
    }
  },

  // 更新用户
  {
    url: '/api/users/:id',
    method: 'put',
    response: ({ body, url }) => {
      const id = url.match(/\/api\/users\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的用户ID'
        }
      }

      const index = allUsers.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allUsers[index] = {
          ...allUsers[index],
          ...body
        }
        return {
          code: 200,
          message: '更新成功',
          data: allUsers[index]
        }
      }
      return {
        code: 404,
        message: '用户不存在'
      }
    }
  }
] 