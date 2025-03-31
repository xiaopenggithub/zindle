import request from '@/utils/request'

export function getMemberOrders(params) {
  return request({
    url: '/api/member/orders',
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