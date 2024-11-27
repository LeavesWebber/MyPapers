<template>
  <div class="background">
    <!-- <div>
    <img class="image" src="../../images/login.jpg" />
  </div> -->
    <div class="box">
      <el-form
        :model="ruleForm"
        :inline="true"
        status-icon
        :rules="rules"
        ref="ruleForm"
        label-width="180px"
        class="demo-ruleForm"
        size="medium"
      >
        <el-form-item label="User Name" prop="username">
          <el-input v-model="ruleForm.username"></el-input>
        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input
            type="password"
            v-model="ruleForm.password"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <div class="button">
          <el-button type="primary" round @click="submitForm('ruleForm')"
            >Login</el-button
          >
          <el-button type="primary" plain round @click="resetForm('ruleForm')"
            >Reset</el-button
          >
        </div>
      </el-form>
    </div>
  </div>
</template>

<script>
import { login, getMenu } from "../../api";
export default {
  data() {
    var checkUserName = (rule, value, callback) => {
      if (!value) {
        return callback(new Error("UserName is required"));
      } else {
        callback();
      }
    };
    var validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("Password is required"));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        username: "",
        password: "",
      },
      rules: {
        username: [
          { required: true, trigger: "blur", message: "please input username" },
        ],
        password: [
          { required: true, trigger: "blur", message: "please input password" },
        ],
      },
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log(this.ruleForm);
          login(this.ruleForm).then(({ data }) => {
            console.log(data.data, "dddddd");
            console.log(data.code, "cccccc");
            if (data.code === 1000) {
              localStorage.setItem(
                "userInfo",
                JSON.stringify(data.data.userInfo)
              ); // 用localStorage缓存用户信息
              localStorage.setItem("token", data.data.token); // 用localStorage缓存token值
              getMenu().then(({ data }) => {
                console.log(data, "mmmmm");
                // 获取菜单的数据，存入store中
                // 用localStorage缓存token值
                localStorage.setItem("menu", JSON.stringify(data.data));
                this.$store.commit("setMenu", data.data);
                this.$store.commit("addMenu", this.$router);
                console.log("----------------------------------------");
                // console.log(this.$router);
                // 跳转至首页
                // this.$router.push("/userinfo");
              });
              this.$notify({
                title: "Login Success!",
                type: "success",
                position: "bottom-right",
              });
              this.$router.push("/home");
            }

            if (data.code === 1005) {
              // this.$alert("User Not Exist!", {
              this.$alert("UserName Or Password False", {
                confirmButtonText: "ok",
              });
            }
          });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
  },
};
</script>
<style lang="less" scoped>
.background {
  // 盒子沾满整个屏幕
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-image: url("../../images/login.jpg");
  background-size: cover;
  // 盒子的背景颜色
  // background-color: #40A2FE;
  // 里面的盒子居中
  display: flex;
  justify-content: center;
  align-items: center;
}
.box {
  // 盒子放在页面中间
  //   position: absolute;
  //   left: 50%;
  //   top: 50%;
  //   transform: translate(-50%, -50%);
  // 盒子的宽度
  width: 500px;
  // 盒子的高度
  height: 300px;
  // 盒子透明
  background-color: rgba(255, 255, 255, 0.5);
  // 盒子的圆角
  border-radius: 8px;
  padding-top: 40px;
  .button {
    // 盒子内的按钮居中
    text-align: center;
    // 放低下
    margin-top: 60px;
  }
}
</style>
