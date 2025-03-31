import request from '@/utils/request'

export function getOrders(params) {
  return request({
    url: '/api/orders',
    method: 'get',
    params
  })
}

export function payOrder(id) {
  return request({
    url: `/api/orders/${id}/pay`,
    method: 'post'
  })
}

export function cancelOrder(id) {
  return request({
    url: `/api/orders/${id}/cancel`,
    method: 'post'
  })
}

export function getMemberInfo(id) {
  return request({
    url: `/api/members/${id}`,
    method: 'get'
  })
} 