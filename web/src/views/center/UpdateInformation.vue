<template>
  <div class="background">
    <div class="box">
      <el-form
        :model="userForm"
        :rules="rules"
        ref="userForm"
        label-width="140px"
        class="demo-userForm"
        size="small"
      >
        <!-- Account Information -->
        <div class="form-section">
          <h3 class="section-title">Account Information</h3>
          <el-row :gutter="20">
            <el-col :span="24">
              <el-form-item label="Username" prop="username">
                <el-input v-model="userForm.username" placeholder="Enter username" disabled></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="Password" prop="password">
                <el-input
                  type="password"
                  v-model="userForm.password"
                  placeholder="Enter new password"
                  autocomplete="off"
                  @focus="passwordFocused = true"
                  @blur="passwordFocused = false"
                ></el-input>
                <div class="password-tips" v-if="passwordError && userForm.password">
                  <div class="tips-title">Password must contain:</div>
                  <ul class="tips-list">
                    <li class="tip-item" :class="{ 'fulfilled': userForm.password.length >= 6 }">
                      At least 6 characters
                    </li>
                    <li class="tip-item" :class="{ 'fulfilled': /[A-Z]/.test(userForm.password) }">
                      At least one uppercase letter
                    </li>
                    <li class="tip-item" :class="{ 'fulfilled': /[a-z]/.test(userForm.password) }">
                      At least one lowercase letter
                    </li>
                    <li class="tip-item" :class="{ 'fulfilled': /[0-9]/.test(userForm.password) }">
                      At least one number
                    </li>
                  </ul>
                </div>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Confirm" prop="confirmPassword">
                <el-input
                  type="password"
                  v-model="userForm.confirmPassword"
                  placeholder="Confirm new password"
                  autocomplete="off"
                ></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="16">
              <el-form-item label="Email" prop="email">
                <el-input 
                  v-model.lazy="userForm.email" 
                  placeholder="Enter email"
                >
                </el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label=" ">
                <el-button 
                  @click="sendVerificationCode"
                  :disabled="!canSendCode || sendingCode"
                  :loading="sendingCode"
                  class="verification-btn"
                  :class="{
                    'btn-active': isValidEmail && userForm.email !== originalEmail,
                    'btn-verified': emailVerified
                  }"
                  size="small"
                  style="width: 100%;"
                >
                  {{ emailVerified ? 'Verified' : 
                     (countdown > 0 ? `Resend (${countdown}s)` : 
                     (sendingCode ? 'Sending...' : 'Get Code')) }}
                </el-button>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="16">
              <el-form-item label="Verification Code" prop="EmailCode">
                <el-input 
                  v-model="userForm.EmailCode" 
                  placeholder="Enter verification code"
                  :disabled="emailVerified || userForm.email === originalEmail"
                >
                  <template #suffix v-if="emailVerified || userForm.email === originalEmail">
                    <i class="el-icon-check" style="color: #67C23A;"></i>
                  </template>
                </el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label=" ">
                <el-button 
                  @click="verifyCode"
                  :disabled="!VerificationSended || emailVerified || userForm.email === originalEmail"
                  class="verification-btn"
                  :class="{
                    'btn-active': VerificationSended && !emailVerified && userForm.email !== originalEmail,
                    'btn-verified': emailVerified || userForm.email === originalEmail
                  }"
                  size="small"
                  style="width: 100%;"
                >
                  {{ emailVerified || userForm.email === originalEmail ? 'Verified' : 'Verify' }}
                </el-button>
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- Personal Information -->
        <div class="form-section">
          <h3 class="section-title">Personal Information</h3>
          <el-row :gutter="20">
            <el-col :span="8">
              <el-form-item label="First Name" prop="first_name">
                <el-input v-model="userForm.first_name" placeholder="First name"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="Middle Name" prop="middle_name">
                <el-input v-model="userForm.middle_name" placeholder="Middle name"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="Last Name" prop="last_name">
                <el-input v-model="userForm.last_name" placeholder="Last name"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="Phone" prop="phone">
                <el-input v-model="userForm.phone" placeholder="Enter phone number"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Address" prop="address">
                <el-input v-model="userForm.address" placeholder="Enter address"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="24">
              <el-form-item label="Blockchain Add." prop="block_chain_address">
                <el-input v-model="userForm.block_chain_address" placeholder="Enter blockchain address"></el-input>
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
                <el-select v-model="userForm.education" placeholder="Select education level" style="width: 100%">
                  <el-option label="Bachelor" value="Bachelor"></el-option>
                  <el-option label="Master" value="Master"></el-option>
                  <el-option label="Doctor" value="Doctor"></el-option>
                  <el-option label="Other" value="Other"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Position/Title" prop="title">
                <el-select v-model="userForm.title" placeholder="Select position/title" style="width: 100%">
                  <el-option label="Professor" value="Professor"></el-option>
                  <el-option label="Associate Professor" value="Associate Professor"></el-option>
                  <el-option label="Assistant Professor" value="Assistant Professor"></el-option>
                  <el-option label="Research Fellow" value="Research Fellow"></el-option>
                  <el-option label="Postdoctoral Researcher" value="Postdoctoral Researcher"></el-option>
                  <el-option label="Doctoral Student" value="Doctoral Student"></el-option>
                  <el-option label="Master's Student" value="Master's Student"></el-option>
                  <el-option label="Undergraduate Student" value="Undergraduate Student"></el-option>
                  <el-option label="Other" value="Other"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="24">
              <el-form-item label="Research Fields" prop="research">
                <el-input v-model="userForm.research" placeholder="Enter research fields (optional)"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="Organization" prop="affiliation">
                <el-input v-model="userForm.affiliation" placeholder="Enter organization"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Department" prop="affiliation_type">
                <el-input v-model="userForm.affiliation_type" placeholder="Enter department"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- Action Buttons -->
        <div class="form-actions">
          <el-button type="primary" round @click="setSelfInfo">Update</el-button>
          <el-button type="default" round @click="resetForm">Reset</el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>
<script>
import { getSelfInfo, setSelfInfo } from "../../api";
export default {
  data() {
    // 密码验证函数
    const validatePass = (rule, value, callback) => {
      if (value === "") {
        callback();
      } else if (value.length < 6) {
        callback(new Error("密码至少需要6个字符"));
      } else if (!/[A-Z]/.test(value)) {
        callback(new Error("密码必须包含至少一个大写字母"));
      } else if (!/[a-z]/.test(value)) {
        callback(new Error("密码必须包含至少一个小写字母"));
      } else if (!/[0-9]/.test(value)) {
        callback(new Error("密码必须包含至少一个数字"));
      } else {
        callback();
      }
    };
    
    // 确认密码验证函数
    const validateConfirmPassword = (rule, value, callback) => {
      if (value === '') {
        callback();
      } else if (value !== this.userForm.password) {
        callback(new Error('两次输入的密码不一致'));
      } else {
        callback();
      }
    };
    
    return {
      userForm: {
        id: 0,
        username: "",
        password: "",
        confirmPassword: "",
        first_name: "",
        middle_name: "",
        last_name: "",
        email: "",
        phone: "",
        address: "",
        education: "",
        title: "",
        research: "",
        block_chain_address: "",
        affiliation: "",
        affiliation_type: "",
        EmailCode: "",
      },
      rules: {
        password: [
          { validator: validatePass, trigger: "blur" }
        ],
        confirmPassword: [
          { validator: validateConfirmPassword, trigger: "blur" }
        ],
        phone: [
          { required: true, trigger: 'blur', message: '请输入11位手机号'},
          { required: true, trigger: 'blur', min: 11, max: 11, message: '长度不符合'},
        ],
        first_name: [
          { required: true, trigger: "blur", message: "请输入名字" },
        ],
        last_name: [
          { required: true, trigger: "blur", message: "请输入姓氏" },
        ],
        email: [
          { required: true, trigger: "blur", message: "请输入邮箱" },
          { type: "email", message: "请输入正确的邮箱地址", trigger: "blur" }
        ],
        block_chain_address: [
          { required: true, trigger: "blur", message: "请输入区块链地址" },
          { pattern: /^0x[a-fA-F0-9]{40}$/, message: "请输入有效的区块链地址", trigger: "blur" }
        ],
        education: [
          { required: true, trigger: "change", message: "请选择教育水平" }
        ],
        title: [
          { required: true, trigger: "change", message: "请选择职位/职称" }
        ],
        affiliation: [
          { required: true, trigger: "blur", message: "请输入组织名称" }
        ],
        affiliation_type: [
          { required: true, trigger: "blur", message: "请输入部门名称" }
        ],
      },
      countdown: 0,
      sendingCode: false,
      emailVerified: false,
      isValidEmail: false,
      VerificationSended: false,
      verifying: false,
      originalEmail: '', // Save original email
      passwordFocused: false,
      passwordError: false,
    };
  },
  methods: {
    setSelfInfo() {
      this.$refs.userForm.validate((valid) => {
        if (valid) {
          if (this.userForm.email !== this.originalEmail && !this.emailVerified) {
            this.$message.warning('Please verify your new email first');
            return;
          }
          
          const updateData = { ...this.userForm };
          delete updateData.confirmPassword;
          delete updateData.EmailCode;
          
          if (!updateData.password) {
            delete updateData.password;
          }
          
          setSelfInfo(updateData).then((res) => {
            console.log("Update response:", res);
            if (res.data && res.data.code === 1000) {
              this.$message({
                message: "Update successful",
                type: "success",
              });
              this.$router.push("/center/information");
            } else {
              this.$message({
                message: res.data?.msg || "Update failed",
                type: "error",
              });
            }
          }).catch(error => {
            console.error("Update failed:", error);
            this.$message({
              message: "Update failed, please try again later",
              type: "error",
            });
          });
        } else {
          this.$message({
            message: "Please fill in all required fields",
            type: "warning",
          });
          return false;
        }
      });
    },
    resetForm() {
      this.$refs.userForm.resetFields();
      // 重新获取用户信息
      this.fetchUserInfo();
    },
    async fetchUserInfo() {
      try {
        const token = localStorage.getItem('token');
        if (!token) {
          this.$message.error("Please login first");
          this.$router.push('/login');
          return;
        }

        const res = await getSelfInfo();
        console.log("Get user info response:", res);
        
        if (res && res.code === 1000 && res.data) {
          const userData = res.data;
          this.originalEmail = userData.email || ''; // Save original email
          this.userForm = {
            id: userData.ID || 0,
            username: userData.username || "",
            password: "", // Password field not displayed
            confirmPassword: "", // Confirm password field not displayed
            first_name: userData.first_name || "",
            middle_name: userData.middle_name || "",
            last_name: userData.last_name || "",
            email: userData.email || "",
            phone: userData.phone || "",
            address: userData.address || "",
            education: userData.education || "",
            title: userData.title || "",
            research: userData.research || "",
            block_chain_address: userData.block_chain_address || "",
            affiliation: userData.affiliation || "",
            affiliation_type: userData.affiliation_type || "",
            EmailCode: "",
          };
          console.log("Processed user info:", this.userForm);
          // Set email as verified initially
          this.emailVerified = true;
        } else {
          console.error("Failed to get user info:", res);
          this.$message.error(res?.msg || "Failed to get user info");
          if (res?.code === 401) {
            this.$router.push('/login');
          }
        }
      } catch (error) {
        console.error("Error getting user info:", error);
        this.$message.error("Failed to get user info, please check your network connection");
        if (error.response?.status === 401) {
          this.$router.push('/login');
        }
      }
    },
    async verifyCode() {
      try {
        this.verifying = true;
        const response = await this.$http.post('/mypapers/user/VerifyMail', {
          email: this.userForm.email,
          code: this.userForm.EmailCode
        });

        if (response.data.code === 1000) {
          const verificationData = {
            email: this.userForm.email,
            code: this.userForm.EmailCode,
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
      if (!this.userForm.email || !this.isValidEmail) {
        this.$message.error('Please enter a valid email address')
        return
      }
      this.sendingCode = true
      this.$http.post('/mypapers/user/SendMail', {
        email: this.userForm.email
      }).then(response => {
        if (response.data.code === 1000) {
          this.startCountdown()
          this.$message.success('Verification code has been sent to your email')
        } else {
          this.$message.error(response.data.msg || 'Failed to send verification code')
        }
      }).catch(error => {
        this.$message.error('Failed to send verification code: ' + (error.response?.data?.msg || error.message))
      }).finally(() => {
        this.sendingCode = false
        this.VerificationSended = true
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
    validateEmail(email) {
      if (!email) return false
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      return emailRegex.test(email)
    },
  },
  watch: {
    'userForm.email': {
      immediate: true,
      handler(value) {
        this.isValidEmail = this.validateEmail(value)
        // If email has changed, reset verification status
        if (value !== this.originalEmail) {
          this.emailVerified = false
          this.VerificationSended = false
          this.userForm.EmailCode = ''
        } else {
          // If email is back to the original, consider it verified
          this.emailVerified = true
        }
      }
    },
    'userForm.password': {
      handler(value) {
        if (!value) {
          this.passwordError = false;
          return;
        }
        
        // Check if password meets all requirements
        const isLengthValid = value.length >= 6;
        const hasUppercase = /[A-Z]/.test(value);
        const hasLowercase = /[a-z]/.test(value);
        const hasNumber = /[0-9]/.test(value);
        
        // Show error if any requirement is not met
        this.passwordError = !(isLengthValid && hasUppercase && hasLowercase && hasNumber);
      }
    }
  },
  mounted() {
    this.fetchUserInfo();
  },
  computed: {
    canSendCode() {
      // Only enable the button if email has changed and is valid
      return this.userForm.email && 
             this.isValidEmail && 
             !this.countdown && 
             !this.emailVerified &&
             this.userForm.email !== this.originalEmail
    }
  },
};
</script>
<style lang="less" scoped>
.background {
  width: 100%;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f7fa;
}

.box {
  width: 80%;
  max-width: 1200px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  padding: 30px;
  margin: 20px auto;
}

.form-section {
  margin-bottom: 30px;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
}

.section-title {
  margin-bottom: 20px;
  color: #303133;
  font-size: 18px;
  font-weight: 500;
}

.form-actions {
  display: flex;
  justify-content: center;
  margin-top: 30px;
  gap: 20px;
}

.el-form-item {
  margin-bottom: 22px;
}

.el-input, .el-select {
  width: 100%;
}

.password-tips {
  margin-top: 10px;
  padding: 10px;
  background-color: #f8f9fa;
  border-radius: 4px;
  border-left: 3px solid #e6e8eb;
  
  .tips-title {
    margin-bottom: 8px;
    font-weight: 500;
    color: #303133;
    font-size: 13px;
  }
  
  .tips-list {
    list-style-type: none;
    padding-left: 0;
    margin: 0;
    
    .tip-item {
      position: relative;
      padding: 4px 0 4px 24px;
      color: #909399;
      font-size: 12px;
      
      &:before {
        content: "•";
        position: absolute;
        left: 10px;
      }
      
      &.fulfilled {
        color: #67C23A;
        
        &:before {
          content: "✓";
          left: 10px;
        }
      }
    }
  }
}

.verification-btn {
  height: 32px;
  border: 1px solid #DCDFE6;
  background: transparent;
  color: #909399;
  font-size: 12px;
  transition: all 0.3s ease;

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
</style>