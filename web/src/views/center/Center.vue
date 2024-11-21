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
.el-main {
  /* background-color: #ECF5FF;
  height: 40px !important; */
  line-height: 40px;
  text-align: center;
}
</style>