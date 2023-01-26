import request from '@/utils/request'
// 订单详情
export function activityOrdersList(data) {
    return request({
        url: '/activityOrders/list',
        method: 'post',
        data
    })
}

export function activityOrdersDelete(data) {
    return request({
        url: '/activityOrders/delete',
        method: 'delete',
        data
    })
}

export function activityOrdersDeleteBatch(data) {
    return request({
        url: '/activityOrders/deleteBatch',
        method: 'delete',
        data
    })
}

export function activityOrdersOne(data) {
    return request({
        url: '/activityOrders/find',
        method: 'post',
        data
    })
}



export function activityOrdersAdd(data) {
    return request({
        url: '/activityOrders/add',
        method: 'post',
        data
    })
}

export function activityOrdersUpdate(data) {
    return request({
        url: '/activityOrders/update',
        method: 'put',
        data
    })
}