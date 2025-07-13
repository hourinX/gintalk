
import { message } from 'ant-design-vue'
import axios from 'axios'

const service = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    timeout: 10000,
    headers: { 'Content-Type': 'application/json;charset=UTF-8'},
})

service.interceptors.request.use(
    config => {
        const token = localStorage.getItem('token')
        if (token) {
            config.headers['Authorization'] = 'Bearer' + token
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