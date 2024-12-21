const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false, // 关闭eslint校验
  devServer: {
    host: '0.0.0.0',  // 确保监听所有 IP 地址
    port: 8080,        // 设置端口号
    proxy: {
      '/mypapers': {
        target: 'http://localhost:8887/mypapers/',  // 改回http
        changeOrigin: true,
        pathRewrite: {
          '/mypapers': ''
        }
      }
    },
    client: {
      webSocketURL: {
        hostname: 'localhost',
        pathname: '/ws',
        port: 8080,
        protocol: 'ws',
      }
    },
    allowedHosts: 'all'
  }
})
