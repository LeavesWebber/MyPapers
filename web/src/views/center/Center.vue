<template>
  <div>
    <el-container>
      <el-aside width="auto">
        <center-aside />
      </el-aside>
      <el-container>
        <el-header>
          <center-header />
        </el-header>
        <el-main>
          <!-- 子路由出口 -->
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script>
import CenterAside from "../../components/CenterAside.vue";
import CenterHeader from "../../components/CenterHeader.vue";
export default {
  data() {
    return {};
  },
  components: {
    CenterAside,
    CenterHeader,
  },
};
</script>
<style>
.el-container {
  min-height: 100vh;
}

.el-aside {
  position: relative;
  z-index: 1;
}

.el-main {
  padding: 20px;
  line-height: 40px;
  text-align: center;
}

@media screen and (max-width: 1200px) {
  .el-main {
    margin-left: 260px;
  }
}

@media screen and (max-width: 768px) {
  .el-main {
    margin-left: 250px;
  }
}
</style>