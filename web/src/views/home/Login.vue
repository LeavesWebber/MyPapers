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
          console.log('Form data:', this.ruleForm);
          login(this.ruleForm).then((response) => {
            console.log('Login response:', response);
            if (response && response.code === 1000 && response.data) {
              // 存储用户信息
              localStorage.setItem(
                "userInfo",
                JSON.stringify(response.data.userInfo)
              );
              // 存储token
              localStorage.setItem("token", response.data.token);
              
              // 获取菜单数据
              getMenu().then((menuResponse) => {
                console.log('Menu response:', menuResponse);
                if (menuResponse && menuResponse.code === 1000) {
                  // 即使菜单数据为空，也存储一个空数组
                  const menuData = menuResponse.data || [];
                  localStorage.setItem("menu", JSON.stringify(menuData));
                  this.$store.commit("setMenu", menuData);
                  this.$store.commit("addMenu", this.$router);
                  
                  // 登录成功提示
                  this.$notify({
                    title: "登录成功！",
                    type: "success",
                    position: "bottom-right",
                  });
                  
                  // 跳转到首页
                  this.$router.push("/home");
                } else {
                  console.error('获取菜单失败:', menuResponse);
                  // 菜单获取失败也允许登录，使用空菜单
                  localStorage.setItem("menu", JSON.stringify([]));
                  this.$store.commit("setMenu", []);
                  this.$store.commit("addMenu", this.$router);
                  
                  this.$notify({
                    title: "登录成功！",
                    message: "但获取菜单失败，部分功能可能受限",
                    type: "warning",
                    position: "bottom-right",
                  });
                  
                  this.$router.push("/home");
                }
              }).catch(error => {
                console.error('获取菜单错误:', error);
                // 菜单获取出错也允许登录，使用空菜单
                localStorage.setItem("menu", JSON.stringify([]));
                this.$store.commit("setMenu", []);
                this.$store.commit("addMenu", this.$router);
                
                this.$notify({
                  title: "登录成功！",
                  message: "但获取菜单失败，部分功能可能受限",
                  type: "warning",
                  position: "bottom-right",
                });
                
                this.$router.push("/home");
              });
            } else if (response && response.code === 1005) {
              this.$alert("用户名或密码错误", {
                confirmButtonText: "确定",
              });
            } else {
              this.$alert(response?.msg || "登录失败，请重试", {
                confirmButtonText: "确定",
              });
            }
          }).catch(error => {
            console.error('登录请求错误:', error);
            this.$message.error('登录请求失败，请重试');
          });
        } else {
          console.log("表单验证失败");
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
