<template>
  <div>
    <el-container>
      <el-aside width="auto">
        <center-aside />
      </el-aside>
      <el-container>
        <el-header height="64px" class="header-wrapper">
          <center-header />
        </el-header>
        <el-main class="content-main">
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
  box-shadow: 1px 0 4px rgba(0, 0, 0, 0.08);
}

.header-wrapper {
  padding: 0;
  line-height: 1;
}

.content-main {
  padding: 20px;
  line-height: 1.5;
}

@media screen and (max-width: 1200px) {
  .content-main {
    margin-left: 260px;
  }
}

@media screen and (max-width: 768px) {
  .content-main {
    margin-left: 250px;
    padding: 15px;
  }
}
</style>