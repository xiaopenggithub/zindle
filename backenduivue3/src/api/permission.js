import request from '@/utils/request'

// 获取权限列表
export function getPermissions(params) {
  return request({
    url: '/api/sys/permissions/list',
    method: 'get',
    params
  })
}

// 删除权限
export function deletePermission(data) {
  return request({
    url: `/api/sys/permissions/delete`,
    method: 'delete',
    data
  })
}

// 添加权限
export function addPermission(data) {
  return request({
    url: '/api/sys/permissions/add',
    method: 'post',
    data
  })
}

// 更新权限
export function updatePermission(data) {
  return request({
    url: `/api/sys/permissions/update`,
    method: 'put',
    data
  })
} 