import request from '@/utils/request'

// 获取用户列表
export function getUsers(params) {
  return request({
    url: '/api/users',
    method: 'get',
    params
  })
}

// 删除用户
export function deleteUser(id) {
  return request({
    url: `/api/users/${id}`,
    method: 'delete'
  })
}

// 添加用户
export function addUser(data) {
  return request({
    url: '/api/users',
    method: 'post',
    data
  })
}

// 更新用户
export function updateUser(id, data) {
  return request({
    url: `/api/users/${id}`,
    method: 'put',
    data
  })
} 