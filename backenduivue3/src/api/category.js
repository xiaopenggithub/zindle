import request from '@/utils/request'

export function getCategorys(params) {
  return request({
    url: '/api/stock/stockCategory/list',
    method: 'get',
    params
  })
}

export function addCategory(data) {
  return request({
    url: '/api/stock/stockCategory/add',
    method: 'post',
    data
  })
}

export function updateCategory(data) {
  return request({
    url: `/api/stock/stockCategory/update`,
    method: 'put',
    data
  })
}

export function deleteCategory(data) {
  return request({
    url: `/api/stock/stockCategory/delete`,
    method: 'delete',
    data
  })
}