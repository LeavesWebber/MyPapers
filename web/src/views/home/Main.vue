<template>
  <div>
    <el-container>
      <el-header>
        <common-header />
      </el-header>
      <logo />
      <navbar />
      <!-- 路由出口 -->
      <!-- 路由匹配到的组件将渲染在这里 -->
      <!-- 子路由出口 -->
      <router-view></router-view>
      <el-footer><common-footer /></el-footer>
    </el-container>
  </div>
</template>
<script>
import CommonFooter from "../../components/CommonFooter.vue";
import CommonHeader from "../../components/CommonHeader.vue";
import Navbar from "../home/Navbar.vue";
import Logo from "../home/Logo.vue";
export default {
  data() {
    return {};
  },
  components: {
    CommonFooter,
    CommonHeader,
    Logo,
    Navbar,
  },
};
</script>
<style>
.el-header {
  background-color: #ecf5ff;
  height: 40px !important;
  line-height: 40px;
  padding: 0;
}
/* .main {
  height: 1000px !important;
} */
.el-footer {
  background-color: #ecf5ff;
  height: 40px !important;
  line-height: 40px;
  text-align: center;
}
</style>