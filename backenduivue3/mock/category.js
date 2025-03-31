import Mock from 'mockjs'

// 生成初始的分类数据
const generateCategories = () => [
  {
    id: 1,
    name: '手机数码',
    description: '手机、平板电脑、数码相机等数码产品',
    created_at: '2024-03-01 10:00:00'
  },
  {
    id: 2,
    name: '电脑办公',
    description: '笔记本、台式机、打印机等办公设备',
    created_at: '2024-03-01 10:30:00'
  },
  {
    id: 3,
    name: '家用电器',
    description: '电视、冰箱、洗衣机等家用电器',
    created_at: '2024-03-01 11:00:00'
  },
  {
    id: 4,
    name: '服装鞋包',
    description: '男装、女装、童装、鞋靴、箱包等',
    created_at: '2024-03-01 11:30:00'
  },
  {
    id: 5,
    name: '食品生鲜',
    description: '零食、饮料、生鲜食品等',
    created_at: '2024-03-01 12:00:00'
  },
  {
    id: 6,
    name: '美妆护肤',
    description: '化妆品、护肤品、香水等',
    created_at: '2024-03-01 12:30:00'
  },
  // 添加更多初始数据
  ...Array.from({ length: 14 }, (_, index) => ({
    id: index + 7,
    name: Mock.Random.ctitle(2, 4) + '类',
    description: Mock.Random.csentence(10, 20),
    created_at: Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
  }))
]

// 生成初始数据
const allCategories = generateCategories()

export default [
  {
    url: '/api/categories',
    method: 'get',
    response: ({ query }) => {
      const { current = 1, pageSize = 10 } = query
      
      // 计算分页
      const start = (parseInt(current) - 1) * parseInt(pageSize)
      const end = start + parseInt(pageSize)
      const paginatedCategories = allCategories.slice(start, end)

      return {
        code: 200,
        data: {
          list: paginatedCategories,
          total: allCategories.length
        }
      }
    }
  },
  {
    url: '/api/categories',
    method: 'post',
    response: ({ body }) => {
      const newCategory = {
        id: allCategories.length + 1,
        ...body,
        created_at: Mock.Random.datetime('yyyy-MM-dd HH:mm:ss')
      }
      allCategories.unshift(newCategory)
      
      return {
        code: 200,
        message: '添加成功',
        data: newCategory
      }
    }
  },
  {
    url: '/api/categories/:id',
    method: 'put',
    response: (options) => {
      const { body } = options
      const id = options.url.match(/\/api\/categories\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的分类ID'
        }
      }

      const index = allCategories.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allCategories[index] = { ...allCategories[index], ...body }
        return {
          code: 200,
          message: '更新成功',
          data: allCategories[index]
        }
      }
      return {
        code: 404,
        message: '分类不存在'
      }
    }
  },
  {
    url: '/api/categories/:id',
    method: 'delete',
    response: (options) => {
      const id = options.url.match(/\/api\/categories\/(\d+)/)?.[1]
      if (!id) {
        return {
          code: 400,
          message: '无效的分类ID'
        }
      }

      const index = allCategories.findIndex(item => item.id === parseInt(id))
      if (index > -1) {
        allCategories.splice(index, 1)
        return {
          code: 200,
          message: '删除成功'
        }
      }
      return {
        code: 404,
        message: '分类不存在'
      }
    }
  }
] 