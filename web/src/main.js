import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import router from './router'
import store from './store'
// 引入 highlight.js 代码高亮工具
import hljs from "highlight.js";
// 使用样式，有多种样式可选
import "highlight.js/styles/github.css";
// 增加自定义命令v-highlight
Vue.directive("highlight", function (el) {
  let blocks = el.querySelectorAll("pre code");
  blocks.forEach(block => {
    hljs.highlightBlock(block);
  });
});
// 增加组定义属性，用于在代码中预处理代码格式
Vue.prototype.$hljs = hljs;
Vue.config.productionTip = false

Vue.use(ElementUI);

// 添加全局前置导航守卫
router.beforeEach((to, from, next) => {
  // 给部分路由添加权限
  // console.log(to.meta.requireAuth)
  if (to.meta.requireAuth) {
    // // 判断token 存不存在
    let token = localStorage.getItem('token')
    // token不存在，说明当前用户未登录，应该跳转至登录页
    if (!token && to.name !== 'login') {
      next({ name: 'login' })
    } else if (token && to.name === 'login') { // token存在，说明用户登录，此时跳转至首页
      console.log(to.name)
      next({ name: 'home' })
    } else {
      next()
    }
  } else {
    next()
  }
})

new Vue({
  router, // 注入到根实例中
  store,
  render: h => h(App),
  created() {
    store.commit('addMenu', router)
  }
}).$mount('#app')
