import request from '@/utils/request'

// 获取菜单列表
export function getMenus(params) {
  return request({
    url: '/api/menus',
    method: 'get',
    params
  })
}

// 删除菜单
export function deleteMenu(id) {
  return request({
    url: `/api/menus/${id}`,
    method: 'delete'
  })
}

// 添加菜单
export function addMenu(data) {
  return request({
    url: '/api/menus',
    method: 'post',
    data
  })
}

// 更新菜单
export function updateMenu(id, data) {
  return request({
    url: `/api/menus/${id}`,
    method: 'put',
    data
  })
} 