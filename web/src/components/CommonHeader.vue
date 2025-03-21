<template>
  <div class="header">
    <el-link
      class="login"
      :underline="false"
      :style="{ display: visibleLogin }"
      href="login"
      >Login</el-link
    >
    <el-button
      class="login"
      :style="{ display: visibleUserInfo }"
      type="text"
      @click="logout"
      >Logout</el-button
    >
    <el-link
      class="register"
      :underline="false"
      :style="{ display: visibleLogin }"
      href="register"
      >Register</el-link
    >
    <el-link
      class="register"
      :underline="false"
      :style="{ display: visibleUserInfo }"
      href="center"
      target="_blank"
      >Center</el-link
    >
    <router-link class="mypapers" to="/" tag="a">mypapers.io</router-link>
  </div>
</template>

<script>
export default {
  data() {
    return {
      visibleUserInfo: "",
      visibleLogin: "",
      activeIndex: "1",
      activeIndex2: "1",
    };
  },
  methods: {
    logout() {
      localStorage.removeItem("token");
      localStorage.removeItem("menu");
      this.visibleUserInfo = "none";
      this.visibleLogin = "";
      // 跳转到主页面
      this.$router
        .push({
          path: "/home",
        })
        .catch((err) => {});
    },
  },
  mounted() {
    // 判断是否登录
    let token = localStorage.getItem("token");
    if (!token) {
      this.visibleUserInfo = "none";
    } else {
      this.visibleLogin = "none";
    }
  },
};
</script>

<style lang="less" scoped>
// 子绝父相
.header {
  position: relative;
  height: 40px;
  line-height: 40px;
  // background-color: #491616;
  .mypapers {
    position: absolute;
    left: 100px;
    color: #40a2fe;
    // 鼠标悬停时显示下划线
    text-decoration: none;
    &:hover {
      text-decoration: underline;
    }
  }
  .login {
    position: absolute;
    right: 100px;
    color: #40a2fe;
    // 鼠标悬停时显示下划线
    &:hover {
      text-decoration: underline;
    }
  }
  .register {
    position: absolute;
    right: 150px;
    color: #40a2fe;
    // 鼠标悬停时显示下划线
    &:hover {
      text-decoration: underline;
    }
  }
}
</style>