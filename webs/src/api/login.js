import request from '@/utils/request'

export function register(data) {
    return request({
        url: '/auth/register',
        method: 'post',
        data: data
    })
}

export function login(data) {
    return request({
        url: '/auth/login',
        method: 'post',
        data: data
    })
}

export function getPublicKey(params) {
    return request({
        url: '/auth/cryptoPk',
        method: 'get',
        params: params
    })
}