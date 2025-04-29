import request from '@/utils/request'

export function getCategoryStocks(params) {
  return request({
    url: '/api/stock/stockCategoryStock/list',
    method: 'get',
    params
  })
}

export function addCategoryStock(data) {
  return request({
    url: '/api/stock/stockCategoryStock/add',
    method: 'post',
    data
  })
}

export function updateCategoryStock(data) {
  return request({
    url: `/api/stock/stockCategoryStock/update`,
    method: 'put',
    data
  })
}

export function deleteCategoryStock(data) {
  return request({
    url: `/api/stock/stockCategoryStock/delete`,
    method: 'delete',
    data
  })
}