const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false, // 关闭eslint校验
  devServer: {
    host: '0.0.0.0',  // 确保监听所有 IP 地址
    port: 8080,        // 设置端口号
    proxy: {
      '/mypapers': {
        target: 'http://127.0.0.1:8887/mypapers/', // 后端接口
        changeOrigin: true, // 是否跨域
        pathRewrite: {
          '/mypapers': ''
        }
      }
    },
    client: {
      overlay: false,
    },
  }
})
