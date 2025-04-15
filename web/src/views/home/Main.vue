<template>
  <div class="main-container">
    <el-container class="content-wrapper">
      <el-header>
        <common-header />
      </el-header>
      <div class="main-content">
        <logo />
        <navbar />
        <!-- 路由出口 -->
        <!-- 路由匹配到的组件将渲染在这里 -->
        <!-- 子路由出口 -->
        <router-view></router-view>
      </div>
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
<style lang="less" scoped>
.main-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.content-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.el-header {
  background-color: #ecf5ff;
  height: 40px !important;
  line-height: 40px;
  padding: 0;
}

.el-footer {
  background-color: #ecf5ff;
  height: 40px !important;
  line-height: 40px;
  text-align: center;
  margin-top: auto;
}
</style>