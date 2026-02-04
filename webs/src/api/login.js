import request from '@/utils/request'

// register user
export function register(data) {
    return request({
        url: '/auth/register',
        method: 'post',
        data: data
    })
}

// user login
export function login(data) {
    return request({
        url: '/auth/login',
        method: 'post',
        data: data
    })
}

// get login captcha
export function captcha(params) {
    return request({
        url: '/auth/captcha',
        method: 'get',
        params: params
    })
}

// get login aeskey
export function getPublicKey(params) {
    return request({
        url: '/auth/cryptoPk',
        method: 'get',
        params: params
    })
}