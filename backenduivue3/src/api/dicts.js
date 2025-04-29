import request from '@/utils/request'

// 查询角色
export function find(data) {
  return request({
    url: '/api/sys/dicts/find',
    method: 'get',
    params: data
  })
}

// 角色列表
export function getDicts(data) {
  return request({
    url: '/api/sys/dicts/list',
    method: 'get',
    params: data
  })
}

// 创建角色
export function addDict(data) {
  return request({
    url: '/api/sys/dicts/add',
    method: 'post',
    data
  })
}

// 修改角色
export function updateDict(data) {
  return request({
    url: '/api/sys/dicts/update',
    method: 'put',
    data
  })
}

// 删除角色
export function deleteDict(data) {
  return request({
    url: '/api/sys/dicts/delete',
    method: 'delete',
    data
  })
}