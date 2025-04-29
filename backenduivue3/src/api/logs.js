import request from '@/utils/request'

export function getLogs(params) {
  return request({
    url: '/api/sys/logslogin/list',
    method: 'get',
    params
  })
}

export function deleteLog(data) {
  return request({
    url: `/api/sys/logslogin/delete`,
    method: 'delete',
    data
  })
}

export function exportLogs() {
  return request({
    url: '/api/logs/export',
    method: 'get',
    responseType: 'blob'
  })
} 


export function getOperationLogs(params) {
  return request({
    url: '/api/sys/logsoperation/list',
    method: 'get',
    params
  })
}

export function deleteOperationLog(data) {
  return request({
    url: `/api/sys/logsoperation/delete`,
    method: 'delete',
    data
  })
}