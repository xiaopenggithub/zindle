import request from '@/utils/request'

// 查询角色
export function find(data) {
  return request({
    url: '/api/sys/dictdatas/find',
    method: 'get',
    params: data
  })
}

// 角色列表
export function getDictDatas(data) {
  return request({
    url: '/api/sys/dictdatas/list',
    method: 'get',
    params: data
  })
}

// 创建角色
export function addDictData(data) {
  return request({
    url: '/api/sys/dictdatas/add',
    method: 'post',
    data
  })
}

// 修改角色
export function updateDictData(data) {
  return request({
    url: '/api/sys/dictdatas/update',
    method: 'put',
    data
  })
}

// 删除角色
export function deleteDictData(data) {
  return request({
    url: '/api/sys/dictdatas/delete',
    method: 'delete',
    data
  })
}