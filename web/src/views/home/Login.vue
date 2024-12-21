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
          login(this.ruleForm).then(({ data }) => {
            if (data.code === 1000) {
              localStorage.setItem(
                "userInfo",
                JSON.stringify(data.data.userInfo)
              );
              localStorage.setItem("token", data.data.token);
              
              getMenu()
                .then(({ data: menuData }) => {
                  console.log('Menu response:', menuData);
                  
                  if (menuData.code === 1000) {
                    const defaultMenu = [
                      {
                        path: "/center",
                        name: "information",
                        label: "Information",
                        icon: "user",
                        url: "Information",
                      }
                    ];
                    
                    const menuToUse = menuData.data || defaultMenu;
                    
                    localStorage.setItem("menu", JSON.stringify(menuToUse));
                    this.$store.commit("setMenu", menuToUse);
                    
                    try {
                      this.$store.commit("addMenu", this.$router);
                    } catch (error) {
                      console.error('Failed to add routes:', error);
                      throw new Error('Failed to add routes');
                    }

                    this.$notify({
                      title: "Login Success!",
                      type: "success",
                      position: "bottom-right",
                    });
                    
                    setTimeout(() => {
                      this.$router.push("/home");
                    }, 100);
                    
                  } else {
                    throw new Error(menuData.msg || 'Failed to get menu data');
                  }
                })
                .catch(error => {
                  console.error('Menu error:', error);
                  this.$notify({
                    title: "Login Error",
                    message: error.message || "Failed to load menu data",
                    type: "error",
                    position: "bottom-right",
                  });
                });
            } else if (data.code === 1005) {
              this.$alert("Username or Password incorrect", {
                confirmButtonText: "OK",
              });
            } else {
              this.$alert(data.msg || "Login failed", {
                confirmButtonText: "OK",
              });
            }
          }).catch(error => {
            console.error('Login error:', error);
            this.$alert("Login failed. Please try again.", {
              confirmButtonText: "OK",
            });
          });
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
