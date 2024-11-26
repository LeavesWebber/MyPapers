<template>
  <div class="logo">
    <img src="../../images/logo.jpg" class="image" />
    <div class="info">My Papers My Rules</div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      msg: "Welcome to Your Vue.js App",
    };
  },
};
</script>
<style lang="less" scoped>
.logo {
  width: 80%;
  height: 80px;
  // 对于页面居中
  margin: 20px auto;
  //   margin-top: 20px;
  // 字号
  font-size: 20px;
  // 字体颜色
  color: #072e5b;
  // 字体加粗
  font-weight: bold;
  //   background-color: #f31c1c;
  .image {
    width: 13%;
    // height: 90px;
    float: left;
  }
  .info {
    // 行内块
    display: inline-block;
    // 宽高
    width: 80%;
    height: 80px;
    // 垂直居中
    line-height: 80px;
    // 文字居中
    text-align: center;
  }
}
</style>