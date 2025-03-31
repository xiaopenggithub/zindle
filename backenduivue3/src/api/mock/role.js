import request from '@/utils/request'

// 获取角色列表
export function getRoles(params) {
  return request({
    url: '/api/roles',
    method: 'get',
    params
  })
}

// 删除角色
export function deleteRole(id) {
  return request({
    url: `/api/roles/${id}`,
    method: 'delete'
  })
}

// 添加角色
export function addRole(data) {
  return request({
    url: '/api/roles',
    method: 'post',
    data
  })
}

// 更新角色
export function updateRole(id, data) {
  return request({
    url: `/api/roles/${id}`,
    method: 'put',
    data
  })
} 