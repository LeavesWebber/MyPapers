<template>
  <div class="navbar">
    <div class="box">
      <a class="a" href="#"
        >About
        <p class="p">
          <a class="a2" href="#">Mission and version</a>
          <a class="a2" href="#">Organization of mypapers</a>
          <a class="a2" href="#">History of mypapers</a>
        </p>
      </a>
      <a class="a" href="#"
        >Membership
        <p class="p">
          <a class="a2" href="#">Join mypapers</a>
          <a class="a2" href="#">Join an society</a>
          <a class="a2" href="#">Member benefits</a>
        </p>
      </a>
      <a class="a" href="committees"
        >Committees
        <p class="p">
          <a class="a2" href="#">Mission and version</a>
          <a class="a2" href="#">Organization of mypapers</a>
          <a class="a2" href="#">History of mypapers</a>
        </p>
      </a>
      <a class="a" href="conferences"
        >Conferences
        <p class="p">
          <a class="a2" href="#">Mission and version</a>
          <a class="a2" href="#">Organization of mypapers</a>
          <a class="a2" href="#">History of mypapers</a>
        </p>
      </a>
      <a class="a" href="journals"
        >Journals
        <p class="p">
          <a class="a2" href="#">Mission and version</a>
          <a class="a2" href="#">Organization of mypapers</a>
          <a class="a2" href="#">History of mypapers</a>
        </p>
      </a>
      <a class="a" href="market"
        >Market
        <p class="p">
          <a class="a2" href="#">Mission and version</a>
          <a class="a2" href="#">Organization of mypapers</a>
          <a class="a2" href="#">History of mypapers</a>
        </p>
      </a>
      <a class="a" href="#"
        >Education
        <p class="pr">
          <a class="a2" href="#">Mission and version</a>
          <a class="a2" href="#">Organization of mypapers</a>
          <a class="a2" href="#">History of mypapers</a>
        </p>
      </a>
      <a class="a" href="#"
        >MyArXiv
        <p class="pr">
          <a class="a2" href="#">Mission and version</a>
          <a class="a2" href="#">Organization of mypapers</a>
          <a class="a2" href="#">History of mypapers</a>
        </p>
      </a>
    </div>
  </div>
</template>
<script>
// import { getHomeData } from "../../api";
export default {
  data() {
    return {
      msg: "Welcome to Your Vue.js App",
    };
  },
  // mounted() {
  //   getHomeData().then((res) => {
  //     console.log(res);
  //   });
  // },
};
</script>
<style lang="less" scoped>
// 子绝父相
.navbar {
  //   position: relative;
  width: 100%;
  // 盒子里面的盒子水平居中
  text-align: center;
  margin-bottom: 20px;
  //   background-color: #f31c1c;
  .box {
    // 宽度被内容撑开
    display: inline-block;
  }
  .a {
    position: relative;
    // width: 80px;
    // 让文字撑开盒子,左右20px，上下0px
    padding: 0 20px;
    height: 40px;

    // 行内块，一行排多个，宽高生效
    display: inline-block;
    // 水平居中
    text-align: center;
    // 垂直居中
    line-height: 40px;
    // 字号
    font-size: 16px;
    // 字体颜色
    color: #000000;
    // 字体加粗
    font-weight: bold;
    // 去掉下划线
    text-decoration: none;
    // 下边框
    border-bottom: 1px solid #ecf5ff;
    // 鼠标悬停时的样式
    &:hover {
      background-color: #ecf5ff;
      color: #40a2fe;
      z-index: 1000; /* 设置一个较高的 z-index 值 */
    }
    .p {
      background-color: #ecf5ff;
      // 显示在所有盒子的上面
      z-index: 1;
      position: absolute;
      left: 0;
      width: 200px;
      top: 40px;
      // 隐藏
      display: none;
      .a2 {
        // 文字靠左
        text-align: left;
        padding-left: 20px;
        // 独占一行
        display: block;
        // 字号
        font-size: 10px;
        // 字体颜色
        color: #40a2fe;
        // 字体不加粗
        font-weight: normal;
        // 去掉下划线
        text-decoration: none;
        // 下边框
        border-bottom: 1px solid #072e5b;
        // 鼠标悬停时显示下划线
        &:hover {
          text-decoration: underline;
        }
      }
    }
    // 鼠标悬停时显示
    // &:hover p {
    //   display: block;
    // }
    .pr {
      background-color: #ecf5ff;
      position: absolute;
      right: 0;
      width: 200px;
      top: 40px;
      // 隐藏
      display: none;
      .a2 {
        // 文字靠右
        text-align: right;
        padding-right: 20px;
        // 独占一行
        display: block;
        // 字号
        font-size: 10px;
        // 字体颜色
        color: #40a2fe;
        // 字体不加粗
        font-weight: normal;
        // 去掉下划线
        text-decoration: none;
        // 下边框
        border-bottom: 1px solid #072e5b;
        // 鼠标悬停时显示下划线
        &:hover {
          text-decoration: underline;
        }
      }
    }
  }
}
</style>