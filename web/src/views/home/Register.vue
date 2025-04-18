<template>
  <div class="background">
    <!-- <div>
      <img class="image" src="../../images/1.jpg" />
    </div> -->
    <div class="box">
      <el-form
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"
        label-width="140px"
        class="demo-ruleForm"
        size="small"
      >
        <!-- Account Information -->
        <div class="form-section">
          <h3 class="section-title">Account Information</h3>
          <el-row :gutter="20">
            <el-col :span="24">
              <el-form-item label="Username" prop="username">
                <el-input v-model="ruleForm.username" placeholder="Enter username"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="Password" prop="password">
                <el-input
                  type="password"
                  v-model="ruleForm.password"
                  placeholder="Enter password"
                  autocomplete="off"
                ></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Confirm" prop="confirmPassword">
                <el-input
                  type="password"
                  v-model="ruleForm.confirmPassword"
                  placeholder="Confirm password"
                  autocomplete="off"
                ></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="16">
              <el-form-item label="Email" prop="email">
                <el-input 
                  v-model.lazy="ruleForm.email" 
                  placeholder="Enter email"
                  :disabled="emailVerified"
                ></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-button 
                @click="sendVerificationCode"
                :disabled="!canSendCode||sendingCode||emailVerified"
                :loading="sendingCode"
                class="send-code-btn verification-btn"
                :class="{
                  'btn-active': isValidEmail && !emailVerified,
                  'btn-verified': emailVerified
                }"
                plain
              >
                {{ emailVerified ? 'Verified' : 
                   (countdown > 0 ? `Resend (${countdown}s)` : 
                   (sendingCode ? 'Sending...' : 'Get Code')) }}
              </el-button>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="16">
              <el-form-item label="Email Code" prop="EmailCode">
                <el-input 
                  v-model="ruleForm.EmailCode" 
                  placeholder="Enter verification code"
                  :disabled="emailVerified"
                >
                  <template #suffix v-if="emailVerified">
                    <i class="el-icon-check" style="color: #67C23A;"></i>
                  </template>
                </el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-button 
                @click="verifyCode"
                :disabled="!VerificationSended || emailVerified"
                class="verify-btn verification-btn"
                :class="{
                  'btn-active': VerificationSended && !emailVerified,
                  'btn-verified': emailVerified
                }"
                plain
              >
                {{ emailVerified ? 'Verified' : 'Verify Code' }}
              </el-button>
            </el-col>
          </el-row>
        </div>

        <!-- Personal Information -->
        <div class="form-section">
          <h3 class="section-title">Personal Information</h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="First Name" prop="first_name">
                <el-input v-model="ruleForm.first_name" placeholder="Enter first name"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Last Name" prop="last_name">
                <el-input v-model="ruleForm.last_name" placeholder="Enter last name"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="Phone" prop="phone">
                <el-input v-model="ruleForm.phone" placeholder="Enter phone number"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Address" prop="address">
                <el-input v-model="ruleForm.address" placeholder="Enter address"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="24">
              <el-form-item label="Blockchain Addr" prop="block_chain_address">
                <el-input v-model="ruleForm.block_chain_address" placeholder="Enter blockchain address"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- Professional Information -->
        <div class="form-section">
          <h3 class="section-title">Professional Information</h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="Education" prop="education">
                <el-input v-model="ruleForm.education" placeholder="Enter education"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Title" prop="title">
                <el-input v-model="ruleForm.title" placeholder="Enter title"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="24">
              <el-form-item label="Research" prop="research">
                <el-input v-model="ruleForm.research" placeholder="Enter research field"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="Affiliations" prop="affiliation">
                <el-input v-model="ruleForm.affiliation" placeholder="Enter affiliations"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Affiliation Type" prop="affiliation_type">
                <el-input v-model="ruleForm.affiliation_type" placeholder="Enter affiliation type"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- Action Buttons -->
        <div class="form-actions">
          <el-button type="primary" round @click="submitForm('ruleForm')"
            :disabled="!(emailVerified && AllNeed)"
          >Register</el-button>
          <el-button type="default" round @click="resetForm('ruleForm')"
          >Reset</el-button>
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
    var validateConfirmPassword = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.ruleForm.password) {
        callback(new Error('两次输入密码不一致!'));
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
        confirmPassword: "",
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
          { min: 6, max: 20, message: "密码长度在6到20个字符之间", trigger: "blur" }
        ],
        confirmPassword: [
          { required: true, trigger: "blur", message: "please confirm password" },
          { validator: validateConfirmPassword, trigger: "blur" }
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
          { required: true, trigger: "blur", message: "please input email" },
          { type: "email", message: "请输入正确的邮箱地址", trigger: "blur" }
        ],
        block_chain_address: [
          {
            required: true,
            trigger: "blur",
            message: "please input block chain address",
          },
          { pattern: /^0x[a-fA-F0-9]{40}$/, message: "请输入有效的区块链地址", trigger: "blur" }
        ],
        EmailCode: [
          { required: true, trigger: "blur", message: "please input the Email Code" },
          { pattern: /^\d{6}$/, message: "验证码必须是6位数字", trigger: "blur" }
        ],
      },
    };
  },
  methods: {
    async verifyCode() {
      try {
        this.verifying = true;
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
          this.$message({
            message: 'Email verification successful',
            type: 'success',
            duration: 2000
          });
        } else {
          this.emailVerified = false;
          this.$message.error(response.data.msg || 'Invalid verification code');
        }
      } catch (error) {
        this.emailVerified = false;
        this.$message.error('Verification failed: ' + (error.response?.data?.msg || error.message));
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
    resetForm(formName) {
      this.$refs[formName].resetFields();
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
  overflow-y: auto;
  min-height: 100vh;
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
  position: relative;
  width: 900px;
  margin: 40px auto;
  background-color: rgba(255, 255, 255, 0.9);
  border-radius: 8px;
  padding: 30px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  .button {
    margin-left: 330px;
  }
}
.error-message {
  color: red;
  margin-top: 10px;
  font-size: 14px;
}

.form-section {
  margin-bottom: 30px;
  padding: 20px;
  background-color: rgba(255, 255, 255, 0.7);
  border-radius: 4px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);

  .section-title {
    margin: 0 0 20px 0;
    padding-bottom: 10px;
    border-bottom: 1px solid #eee;
    color: #409EFF;
    font-size: 16px;
    font-weight: 500;
  }
}

.form-actions {
  text-align: center;
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid #eee;

  .el-button {
    padding: 12px 35px;
    margin: 0 10px;
  }
}

.verification-btn {
  margin-top: 4px;
  width: 100%;
  height: 32px;
  border: 1px solid #DCDFE6;
  background: transparent;
  color: #909399;
  font-size: 12px;
  transition: all 0.3s ease;
  padding: 8px 15px;

  &:hover {
    background-color: #F5F7FA;
    border-color: #DCDFE6;
    color: #606266;
  }

  &.btn-active {
    border-color: #409EFF;
    color: #409EFF;

    &:hover {
      background-color: #ecf5ff;
    }

    &:not(:disabled):hover {
      background-color: #ecf5ff;
      border-color: #409EFF;
      color: #409EFF;
    }
  }

  &.btn-verified {
    background-color: #F0F9EB;
    border-color: #67C23A;
    color: #67C23A;

    &:hover, &:focus {
      background-color: #F0F9EB;
      border-color: #67C23A;
      color: #67C23A;
      cursor: default;
    }

    &.is-disabled {
      background-color: #F0F9EB;
      border-color: #67C23A;
      color: #67C23A;
      opacity: 0.8;
    }
  }

  &:disabled {
    background-color: transparent;
    border-color: #DCDFE6;
    color: #C0C4CC;
    cursor: not-allowed;

    &:hover {
      background-color: transparent;
      border-color: #DCDFE6;
      color: #C0C4CC;
    }
  }
}

:deep(.el-input.is-disabled .el-input__inner) {
  background-color: #F5F7FA;
  border-color: #E4E7ED;
  color: #909399;
  cursor: not-allowed;
}
</style>
