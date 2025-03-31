import request from '@/utils/request'
// 获取用户菜单
export function getUserMenus() {
  return request({
    url: '/api/user/menus',
    method: 'get',
    headers: {
      'x-role': localStorage.getItem('userRoles') || 'admin'
    }
  })
} 