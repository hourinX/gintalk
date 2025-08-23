import request from '@/utils/request'

export function getLoginLogList(params) {
    return request({
        url: '/log/getLoginLogList',
        method: 'get',
        params: params
    })
}