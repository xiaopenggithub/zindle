import request from '@/utils/request'

export function getLogs(params) {
  return request({
    url: '/api/logs',
    method: 'get',
    params
  })
}

export function deleteLog(id) {
  return request({
    url: `/api/logs/${id}`,
    method: 'delete'
  })
}

export function exportLogs() {
  return request({
    url: '/api/logs/export',
    method: 'get',
    responseType: 'blob'
  })
} 