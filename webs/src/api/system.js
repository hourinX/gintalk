import request from '@/utils/request'

export function refresh(data) {
    return request({
        url: '/system/refresh',
        method: 'post',
        data: data
    })
}