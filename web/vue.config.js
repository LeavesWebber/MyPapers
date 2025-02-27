const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false, // 关闭eslint校验
  devServer: {
    proxy: {
      '/mypapers': {
        target: 'http://localhost:8887/mypapers/',// 后端接口
        changeOrigin: true, // 是否跨域
        pathRewrite: {
          '/mypapers': ''
        }
      }
    },
    client: {
      overlay: false,
    },
    allowedHosts: 'all',  // 新增这一行，允许所有的 Host 请求
  }
})