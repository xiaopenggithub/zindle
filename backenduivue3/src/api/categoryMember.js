import request from '@/utils/request'

export function getStockCategoryMember(params) {
  return request({
    url: '/api/stock/stockCategoryMember/list',
    method: 'get',
    params
  })
}

export function assignCategoryMember(data) {
  return request({
    url: '/api/stock/stockCategoryMember/assignCategoryMember',
    method: 'post',
    data
  })
}
