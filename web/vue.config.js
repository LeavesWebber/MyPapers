const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false, // 关闭eslint校验
  devServer: {
    proxy: {
      '/mypapers': {
        target: 'http://127.0.0.1:8887/mypapers/',// 后端接口
        changeOrigin: true, // 是否跨域
        pathRewrite: {
          '/mypapers': ''
        }
      }
    },
    client: {
      overlay: false,
    },
    allowedHosts: 'all',  // 允许所有的 Host 请求
  },
  // 生产环境配置
  productionSourceMap: false, // 关闭生产环境的sourcemap
  publicPath: '/', // 部署应用包时的基本URL
})