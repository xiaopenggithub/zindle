import Mock from 'mockjs'

// 生成初始的会员数据
const generateMembers = (count) => {
  const members = []
  for (let i = 0; i < count; i++) {
    members.push({
      id: i + 1,
      username: `user${i + 1}`,
      nickname: Mock.Random.cname(),
      mobile: Mock.Random.string('1234567890', 11),
      email: Mock.Random.email(),
      status: Mock.Random.pick([0, 1]),
      registerTime: Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
    })
  }
  return members
}

// 生成100条初始数据
const allMembers = generateMembers(100)

export default [
  {
    url: '/api/members',
    method: 'get',
    response: ({ query }) => {
      const { current = 1, pageSize = 10, searchQuery = '' } = query
      
      let filteredMembers = [...allMembers]
      
      // 搜索过滤
      if (searchQuery) {
        filteredMembers = filteredMembers.filter(member => 
          member.username.toLowerCase().includes(searchQuery.toLowerCase()) ||
          member.nickname.toLowerCase().includes(searchQuery.toLowerCase()) ||
          member.mobile.includes(searchQuery) ||
          member.email.toLowerCase().includes(searchQuery.toLowerCase())
        )
      }

      // 计算分页
      const start = (parseInt(current) - 1) * parseInt(pageSize)
      const end = start + parseInt(pageSize)
      const paginatedMembers = filteredMembers.slice(start, end)

      return {
        code: 200,
        data: {
          list: paginatedMembers,
          total: filteredMembers.length
        }
      }
    }
  },
  {
    url: '/api/members',
    method: 'post',
    response: ({ body }) => {
      const newMember = {
        id: allMembers.length + 1,
        ...body,
        registerTime: Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
      }
      allMembers.unshift(newMember)
      
      return {
        code: 200,
        message: '添加成功',
        data: newMember
      }
    }
  },
  {
    url: '/api/members/:id',
    method: 'put',
    response: (options) => {
      const { body } = options
      const id = options.url.match(/\/api\/members\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的会员ID'
        }
      }

      const index = allMembers.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allMembers[index] = { ...allMembers[index], ...body }
        return {
          code: 200,
          message: '更新成功',
          data: allMembers[index]
        }
      }
      return {
        code: 404,
        message: '会员不存在'
      }
    }
  },
  {
    url: '/api/members/:id',
    method: 'delete',
    response: (options) => {
      const id = options.url.match(/\/api\/members\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的会员ID'
        }
      }

      const index = allMembers.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allMembers.splice(index, 1)
        return {
          code: 200,
          message: '删除成功'
        }
      }
      return {
        code: 404,
        message: '会员不存在'
      }
    }
  },
  {
    url: '/api/members/export',
    method: 'get',
    response: () => {
      return {
        code: 200,
        message: '导出成功',
        data: allMembers
      }
    }
  }
] 