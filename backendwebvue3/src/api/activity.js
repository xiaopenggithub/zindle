import request from '@/utils/request'
// 付款信息
export function activityList(data) {
    return request({
        url: '/activity/list',
        method: 'post',
        data
    })
}

export function activityDelete(data) {
    return request({
        url: '/activity/delete',
        method: 'delete',
        data
    })
}

export function activityDeleteBatch(data) {
    return request({
        url: '/activity/deleteBatch',
        method: 'delete',
        data
    })
}

export function activityOne(data) {
    return request({
        url: '/activity/find',
        method: 'post',
        data
    })
}



export function activityAdd(data) {
    return request({
        url: '/activity/add',
        method: 'post',
        data
    })
}

export function activityUpdate(data) {
    return request({
        url: '/activity/update',
        method: 'put',
        data
    })
}






