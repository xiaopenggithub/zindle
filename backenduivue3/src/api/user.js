import request from '@/utils/request'

// 获取验证码
export function getCaptcha() {
  return request({
    url: '/api/sys/user/captcha',
    method: 'get'
  })
}

// 用户登录
export function login(data) {
  return request({
    url: '/api/sys/user/login',
    method: 'post',
    data: {
      username: data.username,
      password: data.password,
      captcha_id: data.captcha_id,
      captcha_code: data.captcha
    }
  })
}

// 修改密码
export function changePassword(data) {
  return request({
    url: '/api/sys/user/password',
    method: 'post',
    data: {
      old_password: data.currentPassword,
      password: data.newPassword,
      confirm_password: data.confirmPassword
    }
  })
}


// 修改信息
export function updateInfo(data) {
  return request({
    url: '/api/sys/user/updateInfo',
    method: 'post',
    data
  })
}

// 用户列表
export function list(data) {
  return request({
    url: '/api/sys/user/list',
    method: 'get',
    params: data
  })
}

// 删除管理员
export function deleteUser(data) {
  return request({
    url: '/api/sys/user/delete',
    method: 'post',
    data
  })
}

// 添加管理员
export function addUser(data) {
  return request({
    url: '/api/sys/user/create',
    method: 'post',
    data
  })
}

// 修改管理员
export function updateUser(data) {
  return request({
    url: '/api/sys/user/updateUserInfo',
    method: 'post',
    data
  })
}
// 获取用户菜单
export function getUserMenus() {
  return request({
    url: '/api/sys/user/menus',
    method: 'get'
  })
} 