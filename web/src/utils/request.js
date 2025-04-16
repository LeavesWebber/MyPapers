import axios from "axios";

const http = axios.create({
    // 通用请求的地址前缀
    baseURL: '/mypapers',
    // 超时时间
    timeout: 1000000,
});

// 添加请求拦截器
http.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    // 设置token start
    let token = localStorage.getItem('token');
    if (token && token !== '') {
        config.headers.Authorization = "Bearer " + token;
    }
    return config;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});

// 添加响应拦截器
http.interceptors.response.use(function (response) {
    // 对响应数据做点什么
    if (response.status === 200) {
        return response;
    }
    return Promise.reject(response);
}, function (error) {
    // 对响应错误做点什么
    console.error('Request error:', error);
    return Promise.reject(error);
});

export default http