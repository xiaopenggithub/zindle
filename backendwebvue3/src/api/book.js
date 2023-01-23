import request from '@/utils/request'
// 图书管理
export function bookList(data) {
    return request({
        url: '/book/list',
        method: 'post',
        data
    })
}

export function bookDelete(data) {
    return request({
        url: '/book/delete',
        method: 'delete',
        data
    })
}

export function bookDeleteBatch(data) {
    return request({
        url: '/book/deleteBatch',
        method: 'delete',
        data
    })
}

export function bookOne(data) {
    return request({
        url: '/book/find',
        method: 'post',
        data
    })
}



export function bookAdd_odd(data) {
    return request({
        url: '/book/add',
        method: 'post',
        data
    })
}

export const bookAdd = (data) => {
    return request({
        url: '/book/add',
        method: 'post',
        donNotShowLoading: true,
        headers: { 'Content-Type': 'multipart/form-data' },
        data
    })
}

export function bookUpdate_old(data) {
    return request({
        url: '/book/update',
        method: 'put',
        data
    })
}

export const bookUpdate = (data) => {
    return request({
        url: '/book/update',
        method: 'put',
        donNotShowLoading: true,
        headers: { 'Content-Type': 'multipart/form-data' },
        data
    })
}
export function counts(data) {
    return request({
        url: '/book/counts',
        method: 'post',
        data
    })
}






