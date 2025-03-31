import Mock from 'mockjs'

// 生成初始的产品数据
const generateProducts = () => {
  // 固定的产品数据
  const fixedProducts = [
    {
      id: 1,
      name: 'iPhone 15 Pro',
      category_id: 1,
      category_name: '手机数码',
      price: 8999.00,
      stock: 100,
      status: 1,
      description: '最新款iPhone，搭载A17芯片',
      created_at: '2024-03-01 10:00:00',
      updated_at: '2024-03-01 10:00:00'
    },
    {
      id: 2,
      name: 'MacBook Pro 14',
      category_id: 2,
      category_name: '电脑办公',
      price: 14999.00,
      stock: 50,
      status: 1,
      description: '搭载M3芯片的专业笔记本',
      created_at: '2024-03-01 10:30:00',
      updated_at: '2024-03-01 10:30:00'
    },
    {
      id: 3,
      name: '海尔冰箱',
      category_id: 3,
      category_name: '家用电器',
      price: 3999.00,
      stock: 30,
      status: 1,
      description: '对开门变频节能冰箱',
      created_at: '2024-03-01 11:00:00',
      updated_at: '2024-03-01 11:00:00'
    },
    {
      id: 4,
      name: 'Nike运动鞋',
      category_id: 4,
      category_name: '服装鞋包',
      price: 699.00,
      stock: 200,
      status: 1,
      description: '舒适透气运动鞋',
      created_at: '2024-03-01 11:30:00',
      updated_at: '2024-03-01 11:30:00'
    },
    {
      id: 5,
      name: '进口牛肉',
      category_id: 5,
      category_name: '食品生鲜',
      price: 199.00,
      stock: 50,
      status: 1,
      description: '澳洲进口牛肉',
      created_at: '2024-03-01 12:00:00',
      updated_at: '2024-03-01 12:00:00'
    },
    {
      id: 6,
      name: '兰蔻面霜',
      category_id: 6,
      category_name: '美妆护肤',
      price: 1299.00,
      stock: 80,
      status: 1,
      description: '高端护肤面霜',
      created_at: '2024-03-01 12:30:00',
      updated_at: '2024-03-01 12:30:00'
    }
  ]

  // 生成随机产品数据
  const randomProducts = Array.from({ length: 94 }, (_, index) => {
    const categoryId = Mock.Random.integer(1, 6)
    const categoryMap = {
      1: '手机数码',
      2: '电脑办公',
      3: '家用电器',
      4: '服装鞋包',
      5: '食品生鲜',
      6: '美妆护肤'
    }

    return {
      id: index + 7,
      name: Mock.Random.ctitle(3, 10),
      category_id: categoryId,
      category_name: categoryMap[categoryId],
      price: Mock.Random.float(10, 10000, 2, 2),
      stock: Mock.Random.integer(0, 1000),
      status: Mock.Random.integer(0, 1),
      description: Mock.Random.cparagraph(1, 3),
      created_at: Mock.Random.datetime('yyyy-MM-dd HH:mm:ss'),
      updated_at: Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
    }
  })

  return [...fixedProducts, ...randomProducts]
}

// 生成初始数据
const allProducts = generateProducts()

export default [
  {
    url: '/api/products',
    method: 'get',
    response: ({ query }) => {
      const { current = 1, pageSize = 10, name = '', category_id } = query
      
      let filteredProducts = [...allProducts]
      
      // 按名称搜索
      if (name) {
        filteredProducts = filteredProducts.filter(item => 
          item.name.toLowerCase().includes(name.toLowerCase())
        )
      }

      // 按分类筛选
      if (category_id) {
        filteredProducts = filteredProducts.filter(item => 
          item.category_id === parseInt(category_id)
        )
      }

      // 计算分页
      const start = (parseInt(current) - 1) * parseInt(pageSize)
      const end = start + parseInt(pageSize)
      const paginatedProducts = filteredProducts.slice(start, end)

      return {
        code: 200,
        data: {
          list: paginatedProducts,
          total: filteredProducts.length
        }
      }
    }
  },
  {
    url: '/api/products',
    method: 'post',
    response: ({ body }) => {
      const categoryMap = {
        1: '手机数码',
        2: '电脑办公',
        3: '家用电器',
        4: '服装鞋包',
        5: '食品生鲜',
        6: '美妆护肤'
      }

      const now = Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
      const newProduct = {
        id: allProducts.length + 1,
        ...body,
        category_name: categoryMap[body.category_id],
        created_at: now,
        updated_at: now
      }
      allProducts.unshift(newProduct)
      
      return {
        code: 200,
        message: '添加成功',
        data: newProduct
      }
    }
  },
  {
    url: '/api/products/:id',
    method: 'put',
    response: (options) => {
      const { body } = options
      const id = options.url.match(/\/api\/products\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的商品ID'
        }
      }

      const index = allProducts.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        const categoryMap = {
          1: '手机数码',
          2: '电脑办公',
          3: '家用电器',
          4: '服装鞋包',
          5: '食品生鲜',
          6: '美妆护肤'
        }

        allProducts[index] = {
          ...allProducts[index],
          ...body,
          category_name: categoryMap[body.category_id],
          updated_at: Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
        }
        return {
          code: 200,
          message: '更新成功',
          data: allProducts[index]
        }
      }
      return {
        code: 404,
        message: '商品不存在'
      }
    }
  },
  {
    url: '/api/products/:id',
    method: 'delete',
    response: (options) => {
      const id = options.url.match(/\/api\/products\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的商品ID'
        }
      }

      const index = allProducts.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allProducts.splice(index, 1)
        return {
          code: 200,
          message: '删除成功'
        }
      }
      return {
        code: 404,
        message: '商品不存在'
      }
    }
  },
  {
    url: '/api/products/export',
    method: 'get',
    response: () => {
      return {
        code: 200,
        message: '导出成功',
        data: allProducts
      }
    }
  }
] 