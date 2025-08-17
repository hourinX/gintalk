
import { message } from 'ant-design-vue'
import axios from 'axios'
import { refresh } from '@/api/system'

const service = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    timeout: 10000,
    headers: { 'Content-Type': 'application/json;charset=UTF-8'},
})

const isTokenValid = (expireTime) => {
    if (!expireTime) return false
    const currentTime = Math.floor(Date.now() / 1000)
    return currentTime < expireTime
}

const refreshToken = async (token) => {
    const refresh_token = localStorage.getItem('refresh_token')
    if (!refresh_token) {
        return false
    }

    try {
        let data = {
            refresh_token: refresh_token
        }
        const res = await refresh(data)

        if (res.code === 10000) {
            const { access_token, expire_time } = res.data
            localStorage.setItem('access_token', access_token)
            localStorage.setItem('expire_time', expire_time)
            return true
        }    
    } catch (error) {
        console.error('Error refreshing token:', error)
        return false
    }
}

service.interceptors.request.use(
    async (config) => {
        const token = localStorage.getItem('access_token')
        const expireTime = localStorage.getItem('expire_time')

        if (token && isTokenValid(expireTime)) {
            config.headers['Authorization'] = `Bearer ${token}`
        } else if (token && !isTokenValid(expireTime)) {
            const refreshed = await refreshToken(token)
            if (refreshed) {
                const newToken = localStorage.getItem('access_token')
                config.headers['Authorization'] = `Bearer ${newToken}`
            } else {
                message.warning('token has expired, please relogin...')
                localStorage.removeItem('access_token')
                localStorage.removeItem('refresh_token')
                localStorage.removeItem('expire_time')
                router.push('/login')
            }
        }
        return config
    },
    error => {
        console.error("request interceptor error:",error)
        message.error('send request error')
        return Promise.reject(error)
    }
)

service.interceptors.response.use(
    response => {
        const res = response.data
        if (res.code !== 10000) {
            message.error(res.msg || 'Error')
            if (res.code === 401) {
                message.warning('token has expired, please relogin...')
                localStorage.removeItem('token')
                router.push('/login')
            }
            return Promise.reject(new Error(res.msg || 'Error'))
        } else {
            message.success(res.msg)
            return res
        }
    },
    error => {
        console.error('response interceptors error:',error)
        let errorMsg = 'unknow error'
        if (error.response) {
            switch (error.response.status) {
                case 401:
                    errorMsg = 'token has expired,please relogin...'
                    localStorage.removeItem('token')
                    router.push('/login')
                    break
                default:
                    errorMsg = `connector error(${error.response.status})!`
            }
        } else {
            if (error.message.includes('timeout')) {
                errorMsg = 'request timeout,please retry later...'
            } else {
                errorMsg = 'connect internet error,please check out net'
            }
        }
        message.error(errorMsg)
        return Promise.reject(error)
    }
)

export default service