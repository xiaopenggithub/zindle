import request from '@/utils/request'

export function getMemberLevels(params) {
  return request({
    url: '/api/ums/memberLevel/list',
    method: 'get',
    params
  })
}

export function addMemberLevel(data) {
  return request({
    url: '/api/ums/memberLevel/add',
    method: 'post',
    data
  })
}

export function updateMemberLevel(data) {
  return request({
    url: `/api/ums/memberLevel/update`,
    method: 'put',
    data
  })
}

export function deleteMemberLevel(data) {
  return request({
    url: `/api/ums/memberLevel/delete`,
    method: 'delete',
    data
  })
}

export function exportMembers() {
  return request({
    url: '/api/memberLevel/export',
    method: 'get'
  })
} 