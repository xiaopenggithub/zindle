import request from '@/utils/request'

export function getMembers(params) {
  return request({
    url: '/api/ums/members/list',
    method: 'get',
    params
  })
}

export function addMember(data) {
  return request({
    url: '/api/ums/members/add',
    method: 'post',
    data
  })
}

export function updateMember(data) {
  return request({
    url: `/api/ums/members/update`,
    method: 'put',
    data
  })
}

export function deleteMember(data) {
  return request({
    url: `/api/ums/members/delete`,
    method: 'delete',
    data
  })
}

export function exportMembers() {
  return request({
    url: '/api/members/export',
    method: 'get'
  })
} 