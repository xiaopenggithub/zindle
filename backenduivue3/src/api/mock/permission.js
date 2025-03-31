import request from '@/utils/request'

// 获取权限列表
export function getPermissions(params) {
  return request({
    url: '/api/permissions',
    method: 'get',
    params
  })
}

// 删除权限
export function deletePermission(id) {
  return request({
    url: `/api/permissions/${id}`,
    method: 'delete'
  })
}

// 添加权限
export function addPermission(data) {
  return request({
    url: '/api/permissions',
    method: 'post',
    data
  })
}

// 更新权限
export function updatePermission(id, data) {
  return request({
    url: `/api/permissions/${id}`,
    method: 'put',
    data
  })
} 