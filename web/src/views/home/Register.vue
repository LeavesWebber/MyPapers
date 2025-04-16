<template>
  <div class="background">
    <!-- <div>
      <img class="image" src="../../images/1.jpg" />
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
        size="mini"
      >
        <el-form-item label="UserName" prop="username">
          <el-input v-model="ruleForm.username"></el-input>
        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input
            type="password"
            v-model="ruleForm.password"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="FirstName" prop="first_name">
          <el-input v-model="ruleForm.first_name"></el-input>
        </el-form-item>
        <el-form-item label="LastName" prop="last_name">
          <el-input v-model="ruleForm.last_name"></el-input>
        </el-form-item>
        <!-- <el-form-item label="Sex" prop="data">
          <el-input v-model="ruleForm.sex"></el-input>
        </el-form-item> -->
        <el-form-item label="Email" prop="email" style="margin-right: 0.75%;">
          <el-input v-model.lazy="ruleForm.email"></el-input>
          
        </el-form-item>
        <el-form-item>
           <el-button @click="sendVerificationCode"
           :disabled="!canSendCode||sendingCode"
           :loading="sendingCode"
           style="width: 85px;align-content: center; font-size:10px;height: 28px;"
           >
           {{ countdown > 0 ? `Resend (${countdown}s)` : (sendingCode ? 'Sending...' : 'Get Code') }}

          </el-button>
        </el-form-item>
       
        <el-form-item label-width="auto" label="Department" prop="department">
          <el-input v-model="ruleForm.department"></el-input>
        </el-form-item>
        <el-form-item label="Phone" prop="phone">
          <el-input v-model="ruleForm.phone">
          </el-input>
        </el-form-item>
        <el-form-item label="Address" prop="address">
          <el-input v-model="ruleForm.address"></el-input>
        </el-form-item>
        <el-form-item label="Education" prop="education">
          <el-input v-model="ruleForm.education"></el-input>
        </el-form-item>
        <el-form-item label="Title" prop="title">
          <el-input v-model="ruleForm.title"></el-input>
        </el-form-item>
        <el-form-item label="Research" prop="research">
          <el-input v-model="ruleForm.research"></el-input>
        </el-form-item>
        <el-form-item label="Blockchain Addr" prop="block_chain_address">
          <el-input v-model="ruleForm.block_chain_address"></el-input>
        </el-form-item>
        <el-form-item label="Affiliations" prop="affiliation">
          <el-input v-model="ruleForm.affiliation"></el-input>
        </el-form-item>
        <el-form-item label="AffiliationType" prop="affiliation_type">
          <el-input v-model="ruleForm.affiliation_type"></el-input>
        </el-form-item>
        <el-form-item label="EmailCode" prop="EmailCode">
          <el-input v-model="ruleForm.EmailCode"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="verifyCode"
          :disabled="!VerificationSended">
            Verificate
          </el-button>
        </el-form-item>
        <div class="button">
          <el-button type="primary" round @click="submitForm('ruleForm')"
          :disabled="!(emailVerified && AllNeed)"
          >Register</el-button
          >
          <div class="modal" v-show="showModal">
    </div>
          <el-button type="primary" plain round @click="resetForm('ruleForm')"
            >Reset</el-button
          >
        </div>
        
      </el-form>
    </div>
  </div>
</template>

<script>
import { register, SendMail, submitPaper } from "../../api";
import { MPScontractInstance } from "@/constant";
const contractInstance = MPScontractInstance;
export default {
  data() {
    var checkData = (rule, value, callback) => {
      if (!value) {
        return callback(new Error("data is required"));
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
      inputValue:"",
      showModal:null,
      showError:null,
      countdown:60,
      sendingCode: false,
      countdown:0,
      emailVerified:false,
      isValidEmail:false,
      showVerificationTip:undefined,
      VerificationSended:false,
      verifying:false,
      AllNeed:false,
      ruleForm: {
        id: 0,
        username: "",
        password: "",
        first_name: "",
        last_name: "",
        // sex: "",
        email: "",
        department: "",
        phone: "",
        address: "",
        education: "",
        title: "",
        research: "",
        block_chain_address: "",
        affiliation: "",
        affiliation_type: "",
        EmailCode:"",
      },
      SendMails:{
        MailReceiver:"",
        Verification:"",
      } ,
      verificationCode:"",
      rules: {
        username: [
          { required: true, trigger: "blur", message: "please input username" },
        ],
        password: [
          { required: true, trigger: "blur", message: "please input password" },
        ],
        phone: [
          { required: true, trigger: 'blur', message: '请输入11位手机号'},
          { required: true, trigger: 'blur', min: 11, max: 11, message: '长度不符合'},
          
        ],
        first_name: [
          {
            required: true,
            trigger: "blur",
            message: "please input first name",
          },
        ],
        last_name: [
          {
            required: true,
            trigger: "blur",
            message: "please input last name",
          },
        ],
        email: [
          { required: true, trigger: "blur", message: "please input email",type:'email'},
        ],
        block_chain_address: [
          {
            required: true,
            trigger: "blur",
            message: "please input block chain address",
          },
        ],
        EmailCode:[{required:true,trigger:"blur",message:"please input the Email Code"}],
      },
    };
  },
  methods: {
    async verifyCode() {
      try {
        const response = await this.$http.post('/mypapers/user/VerifyMail', {
          email: this.ruleForm.email,
          code: this.ruleForm.EmailCode
        });

        if (response.data.code === 1000) {
          // 验证成功后，将验证信息存储在 localStorage 中
          const verificationData = {
            email: this.ruleForm.email,
            code: this.ruleForm.EmailCode,
            token: response.data.data.token,
            expiresAt: response.data.data.expires_at
          };
          localStorage.setItem('emailVerification', JSON.stringify(verificationData));
          
          this.emailVerified = true;
          this.$message.success('邮箱验证成功');
        } else {
          this.emailVerified = false;
          this.$message.error(response.data.msg || '验证码无效');
        }
      } catch (error) {
        this.emailVerified = false;
        this.$message.error('验证失败: ' + (error.response?.data?.msg || error.message));
      } finally {
        this.verifying = false;
      }
    },
    sendVerificationCode() {
      if (!this.ruleForm.email || !this.isValidEmail) {
        this.$message.error('Please enter a valid email address')
        return
      }
      this.sendingCode = true
      this.$http.post('/mypapers/user/SendMail', {
        email: this.ruleForm.email
      }).then(response => {
        if (response.data.code === 1000) {
          this.showVerificationTip = true
          this.startCountdown()
          this.$message.success('Verification code has been sent to your email')
        } else {
          this.$message.error(response.data.msg || 'Failed to send verification code')
        }
      }).catch(error => {
        this.$message.error('Failed to send verification code: ' + (error.response?.data?.msg || error.message))
      }).finally(() => {
        this.sendingCode = false
        this.VerificationSended=true
      })
    },

    startCountdown() {
      this.countdown = 60
      if (this.timer) {
        clearInterval(this.timer)
      }
      this.timer = setInterval(() => {
        if (this.countdown > 0) {
          this.countdown--
        } else {
          clearInterval(this.timer)
        }
      }, 1000)
    },

    async registe_gift(block_chain_address) {
      const functionArgs = [
          block_chain_address
        ];
      const functionName="registerUser"
      const result = await contractInstance.methods[functionName](
          ...functionArgs
        ).send({
          from: window.ethereum.selectedAddress,
          gasPrice: "0",
        });
      
      },
    fixall(){
      this.AllNeed=this.ruleForm.username 
      && this.ruleForm.password 
      && this.ruleForm.email 
      && this.ruleForm.first_name 
      && this.ruleForm.last_name 
      && this.ruleForm.phone 
      && this.ruleForm.block_chain_address
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          register(this.ruleForm).then(({ data }) => {
            console.log(data.data);
            if (data.code === 1000) {
              // localStorage.setItem("token", data.data.token); // 用localStorage缓存token值
              this.$alert("Register success", {
                confirmButtonText: "ok",
              });
              this.registe_gift(this.ruleForm.block_chain_address)
              this.$router.push("/home");
            }

            if (data.code === 1003) {
              this.$alert("User Existed!", {
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
    validateEmail(email) {
      if (!email) return false
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      return emailRegex.test(email)
    },
  submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          register(this.ruleForm).then(({ data }) => {
            console.log(data.data);
            if (data.code === 1000) {
              // localStorage.setItem("token", data.data.token); // 用localStorage缓存token值
              this.$alert("Register success", {
                confirmButtonText: "ok",
              });
              this.registe_gift(this.ruleForm.block_chain_address)
              this.$router.push("/home");
            }

            if (data.code === 1003) {
              this.$alert("User Existed!", {
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
},
  computed:{
    canVerify() {
      return !this.emailVerified && this.verificationCode.length === 6 && !this.verifying
    },
    canSendCode() {
      const result =this.ruleForm.email && !this.countdown && this.isValidEmail
      console.log(this.isValidEmail)
      console.log(this.countdown)
      console.log(result);
      return result
  }
},
watch:{
  'ruleForm.email': {
      immediate: true,
      handler(value) {
        this.isValidEmail = this.validateEmail(value)
        console.log('Watch triggered:', {
          value,
          isValidEmail: this.isValidEmail
        })
      }
    },
  'ruleForm.username':{
    immediate: true,
    handler(){
      this.fixall()
    }
  },
  'ruleForm.password':{
    immediate: true,
    handler(){
      this.fixall()
    }
  },
  'ruleForm.firstname':{
    immediate: true,
    handler(){
      this.fixall()
    }
  },
  'ruleForm.lastname':{
    immediate: true,
    handler(){
      this.fixall()
    }
  },
  'ruleForm.block_chain_address':{
    immediate: true,
    handler(){
      this.fixall()
    }
  },
  'ruleForm.phone':{
    immediate: true,
    handler(){
      this.fixall()
    }
  },
},
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
  }
  

</script>
<style lang="less" scoped>
.image {
  width: 100%;
  height: 950px;
}
// .box {
//   border-radius: 4px;
//   box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
//   position: absolute;
//   left: 50%;
//   top: 50%;
//   transform: translate(-50%, -50%);
//   background-color: #f2f6fc;
//   height: 550px;
//   width: 70%;
//     .demo-ruleForm {
//       // display: inline-block;
//       padding-right: 60px;
//       margin-top: 80px;
//     }
//   .button {
//     margin-left: 330px;
//   }
// }
.background {
  // 盒子沾满整个屏幕
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-image: url("../../images/login.jpg");
  background-size: cover;
}
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1;
}
.modal-content {
  background-color: #fff;
  padding: 30px;
  border-radius: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
}
.open-modal-button {
  padding: 10px 20px;
  background-color: #007BFF;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
}
.open-modal-button:hover {
  background-color: #0056b3;
}
.input-label {
  margin-bottom: 10px;
  font-size: 18px;
}
.number-input {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 16px;
  width: 200px;
  margin-bottom: 20px;
}
.number-input:focus {
  outline: none;
  border-color: #007BFF;
  box-shadow: 0 0 5px rgba(0, 123, 255, 0.5);
}
.button-container {
  display: flex;
  justify-content: space-around;
  width: 100%;
}
.confirm-button,
.cancel-button {
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
}
.confirm-button {
  background-color: #28a745;
  color: #fff;
}
.confirm-button:hover {
  background-color: #218838;
}
.cancel-button {
  background-color: #dc3545;
  color: #fff;
}
.cancel-button:hover {
  background-color: #c82333;
}
.box {
  // 盒子放在页面中间
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  // 盒子的宽度
  width: 850px;
  // 高度被内容撑开
  display: inline-block;

  // 盒子透明
  background-color: rgba(255, 255, 255, 0.5);
  // 盒子的圆角
  border-radius: 8px;
  padding-top: 25px;
  .button {
    margin-left: 330px;
  }
  padding-bottom: 20px;
}
.error-message {
  color: red;
  margin-top: 10px;
  font-size: 14px;
}

</style>
