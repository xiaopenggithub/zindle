import Mock from 'mockjs'

// 生成初始的订单数据
const generateOrders = () => {
  // 固定的订单数据
  const fixedOrders = [
    {
      id: 1,
      orderNo: 'ORDER202403010001',
      memberName: '张三',
      memberId: 1,
      amount: 99.00,
      status: 'paid',
      createTime: '2024-03-01 10:00:00',
      payTime: '2024-03-01 10:05:00',
      cancelTime: null,
      cancelReason: null,
      remark: '无'
    },
    {
      id: 2,
      orderNo: 'ORDER202403010002',
      memberName: '李四',
      memberId: 2,
      amount: 199.00,
      status: 'unpaid',
      createTime: '2024-03-01 11:00:00',
      payTime: null,
      cancelTime: null,
      cancelReason: null,
      remark: '加急订单'
    },
    {
      id: 3,
      orderNo: 'ORDER202403010003',
      memberName: '王五',
      memberId: 3,
      amount: 299.00,
      status: 'cancelled',
      createTime: '2024-03-01 12:00:00',
      payTime: null,
      cancelTime: '2024-03-01 13:00:00',
      cancelReason: '用户取消',
      remark: '无'
    }
  ]

  // 生成随机订单数据
  const randomOrders = Array.from({ length: 97 }, (_, index) => {
    const status = Mock.Random.pick(['paid', 'unpaid', 'cancelled'])
    const createTime = Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
    const amount = Mock.Random.float(10, 10000, 2, 2)
    const memberId = Mock.Random.integer(1, 100)

    return {
      id: index + 4,
      orderNo: `ORDER${Mock.Random.datetime('yyyyMMdd')}${String(index + 4).padStart(4, '0')}`,
      memberName: Mock.Random.cname(),
      memberId,
      amount,
      status,
      createTime,
      payTime: status === 'paid' ? Mock.Random.datetime('yyyy-MM-dd HH:mm:ss') : null,
      cancelTime: status === 'cancelled' ? Mock.Random.datetime('yyyy-MM-dd HH:mm:ss') : null,
      cancelReason: status === 'cancelled' ? Mock.Random.pick(['用户取消', '超时未支付', '库存不足', '其他原因']) : null,
      remark: Mock.Random.boolean() ? Mock.Random.csentence(3, 10) : '无'
    }
  })

  return [...fixedOrders, ...randomOrders]
}

// 生成初始数据
const allOrders = generateOrders()

// 生成会员数据
const generateMemberInfo = (memberId, memberName) => {
  return {
    id: memberId,
    username: memberName,
    nickname: `昵称_${memberName}`,
    mobile: Mock.Random.phone(),
    email: `${Mock.Random.word(5, 10)}@${Mock.Random.domain()}`,
    registerTime: Mock.Random.datetime('yyyy-MM-dd HH:mm:ss'),
    status: Mock.Random.pick(['active', 'inactive'])
  }
}

export default [
  {
    url: '/api/orders',
    method: 'get',
    response: ({ query }) => {
      const { current = 1, pageSize = 10, status = 'all', keyword = '', startTime, endTime } = query
      
      let filteredOrders = [...allOrders]
      
      // 按状态筛选
      if (status !== 'all') {
        filteredOrders = filteredOrders.filter(item => item.status === status)
      }

      // 按关键字搜索
      if (keyword) {
        const loweredKeyword = keyword.toLowerCase()
        filteredOrders = filteredOrders.filter(item => 
          item.orderNo.toLowerCase().includes(loweredKeyword) ||
          item.memberName.toLowerCase().includes(loweredKeyword)
        )
      }

      // 按时间范围筛选
      if (startTime && endTime) {
        const start = new Date(startTime).getTime()
        const end = new Date(endTime).getTime()
        filteredOrders = filteredOrders.filter(item => {
          const createTime = new Date(item.createTime).getTime()
          return createTime >= start && createTime <= end
        })
      }

      // 计算分页
      const start = (parseInt(current) - 1) * parseInt(pageSize)
      const end = start + parseInt(pageSize)
      const paginatedOrders = filteredOrders.slice(start, end)

      return {
        code: 200,
        data: {
          list: paginatedOrders,
          total: filteredOrders.length
        }
      }
    }
  },
  {
    url: '/api/orders/:id/pay',
    method: 'post',
    response: (options) => {
      const id = options.url.match(/\/api\/orders\/(\d+)\/pay/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的订单ID'
        }
      }

      const index = allOrders.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        if (allOrders[index].status !== 'unpaid') {
          return {
            code: 400,
            message: '订单状态不正确'
          }
        }
        allOrders[index].status = 'paid'
        allOrders[index].payTime = Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
        return {
          code: 200,
          message: '支付成功'
        }
      }
      return {
        code: 404,
        message: '订单不存在'
      }
    }
  },
  {
    url: '/api/orders/:id/cancel',
    method: 'post',
    response: (options) => {
      const id = options.url.match(/\/api\/orders\/(\d+)\/cancel/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的订单ID'
        }
      }

      const index = allOrders.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        if (allOrders[index].status !== 'unpaid') {
          return {
            code: 400,
            message: '订单状态不正确'
          }
        }
        allOrders[index].status = 'cancelled'
        allOrders[index].cancelTime = Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
        allOrders[index].cancelReason = '用户取消'
        return {
          code: 200,
          message: '取消成功'
        }
      }
      return {
        code: 404,
        message: '订单不存在'
      }
    }
  },
  {
    url: '/api/members/:id',
    method: 'get',
    response: (options) => {
      const id = options.url.match(/\/api\/members\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的会员ID'
        }
      }

      const order = allOrders.find(item => item.memberId === parseInt(id))
      if (order) {
        return {
          code: 200,
          data: generateMemberInfo(parseInt(id), order.memberName)
        }
      }
      return {
        code: 404,
        message: '会员不存在'
      }
    }
  },
  {
    url: '/api/member/orders',
    method: 'get',
    response: ({ query }) => {
      const { current = 1, pageSize = 10, status = 'all', keyword = '', startTime, endTime } = query
      
      // 模拟当前登录会员的ID为1
      const currentMemberId = 1
      
      let filteredOrders = allOrders.filter(order => order.memberId === currentMemberId)
      
      // 按状态筛选
      if (status !== 'all') {
        filteredOrders = filteredOrders.filter(item => item.status === status)
      }

      // 按关键字搜索
      if (keyword) {
        const loweredKeyword = keyword.toLowerCase()
        filteredOrders = filteredOrders.filter(item => 
          item.orderNo.toLowerCase().includes(loweredKeyword)
        )
      }

      // 按时间范围筛选
      if (startTime && endTime) {
        const start = new Date(startTime).getTime()
        const end = new Date(endTime).getTime()
        filteredOrders = filteredOrders.filter(item => {
          const createTime = new Date(item.createTime).getTime()
          return createTime >= start && createTime <= end
        })
      }

      // 计算分页
      const start = (parseInt(current) - 1) * parseInt(pageSize)
      const end = start + parseInt(pageSize)
      const paginatedOrders = filteredOrders.slice(start, end)

      return {
        code: 200,
        data: {
          list: paginatedOrders,
          total: filteredOrders.length
        }
      }
    }
  }
] 