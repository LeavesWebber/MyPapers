<template>
  <div class="logo">
    <router-link to="/">
      <img src="../../images/logo.jpg" class="image" />
    </router-link>
    <div class="info">My Papers, My Rules</div>
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
  margin: 20px auto;
  position: relative;
  display: flex;
  align-items: center;
  
  .image {
    width: 10%;
    float: left;
  }
  
  .info {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    font-size: 25px;
    font-weight: 600;
    color: #2c3e50;
    letter-spacing: 1px;
    white-space: nowrap;
  }
}
</style>