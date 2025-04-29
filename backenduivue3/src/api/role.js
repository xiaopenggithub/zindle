import request from '@/utils/request'

// 查询角色
export function find(data) {
  return request({
    url: '/api/sys/roles/find',
    method: 'get',
    params: data
  })
}

// 角色列表
export function getRoles(data) {
  return request({
    url: '/api/sys/roles/list',
    method: 'get',
    params: data
  })
}

// 创建角色
export function addRole(data) {
  return request({
    url: '/api/sys/roles/add',
    method: 'post',
    data
  })
}

// 修改角色
export function updateRole(data) {
  return request({
    url: '/api/sys/roles/update',
    method: 'put',
    data
  })
}

// 删除角色
export function deleteRole(data) {
  return request({
    url: '/api/sys/roles/delete',
    method: 'delete',
    data
  })
}
// 获取角色的权限列表
export function gePermissions(data) {  
  return request({
    url: '/api/sys/roles/getPermissions',
    method: 'get',
    params: data
  })
}

// 分配角色权限
export function assignPermissions(data) {
  return request({
    url: '/api/sys/roles/assignPermissions',
    method: 'post',
    data
  })
}
