import axios from 'axios'

// 创建axios实例
const service = axios.create({
    baseURL: "http://localhost:8080", // 从环境变量获取基础URL
    timeout: 10000, // 请求超时时间
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
})

// 请求拦截器
service.interceptors.request.use(
    config => {
        // 在发送请求之前做些什么
        const token = "fakeToken"
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`
        }
        return config
    },
    error => {
        // 对请求错误做些什么
        console.log('Request Error:', error)
        return Promise.reject(error)
    }
)

// 响应拦截器
service.interceptors.response.use(
    response => {
        console.log(response)
        // 对响应数据做点什么
        const res = response.data

        // 假设业务状态码不是20000则为错误
        if (res.code !== 1000) {
            message.error(res.message || 'Error')

            // 特殊状态码处理
            if (res.code === 50008 || res.code === 50012 || res.code === 50014) {
                // 重新登录
                message.error('会话已过期，请重新登录')
                // 这里可以调用登出方法
            }
            return Promise.reject(new Error(res.message || 'Error'))
        } else {
            return res
        }
    },
    error => {
        // 对响应错误做点什么
        console.log('Response Error:', error)

        if (error.response) {
            switch (error.response.status) {
                case 400:
                    error.message = '请求错误'
                    break
                case 401:
                    error.message = '未授权，请登录'
                    // 这里可以跳转到登录页面
                    break
                case 403:
                    error.message = '拒绝访问'
                    break
                case 404:
                    error.message = `请求地址出错: ${error.response.config.url}`
                    break
                case 408:
                    error.message = '请求超时'
                    break
                case 500:
                    error.message = '服务器内部错误'
                    break
                case 501:
                    error.message = '服务未实现'
                    break
                case 502:
                    error.message = '网关错误'
                    break
                case 503:
                    error.message = '服务不可用'
                    break
                case 504:
                    error.message = '网关超时'
                    break
                case 505:
                    error.message = 'HTTP版本不受支持'
                    break
                default:
                    error.message = `连接错误 ${error.response.status}`
            }
        } else if (error.message.includes('timeout')) {
            error.message = '请求超时'
        } else if (error.message === 'Network Error') {
            error.message = '网络异常，请检查网络连接'
        }

        alert(error.message)
        return Promise.reject(error)
    }
)

// 封装GET请求
export function get(url, params = {}, config = {}) {
    return service({
        url,
        method: 'get',
        params,
        ...config
    })
}

// 封装POST请求
export function post(url, data = {}, config = {}) {
    return console.log()

    return service({
        url,
        method: 'post',
        data,
        ...config
    })
}

export default service