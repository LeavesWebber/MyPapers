<template>
  <div class="upload-paper">
    <div class="container">
      <h1>Upload Published Paper</h1>
      <el-form :model="paperForm" :rules="rules" ref="paperForm" label-width="180px">
        <el-form-item label="Paper Title" prop="title">
          <el-input v-model="paperForm.title" placeholder="Enter the full title of your paper"></el-input>
        </el-form-item>

        <el-form-item label="Authors" prop="authors">
          <el-input
            type="textarea"
            v-model="paperForm.authors"
            placeholder="Enter authors' names (separated by commas)"
            :rows="2"
          ></el-input>
        </el-form-item>

        <el-form-item label="Abstract" prop="abstract">
          <el-input
            type="textarea"
            v-model="paperForm.abstract"
            placeholder="Enter paper abstract"
            :rows="4"
          ></el-input>
        </el-form-item>

        <el-form-item label="Corresponding Author Email" prop="correspondingEmail">
          <div class="email-verification-group">
            <el-input 
              v-model="paperForm.correspondingEmail" 
              placeholder="Enter corresponding author's email"
              @input="handleEmailInput"
              :disabled="emailVerified">
              <template slot="append">
                <el-button 
                  @click="sendVerificationCode" 
                  :disabled="!canSendCode"
                  :class="{ 'is-primary': canSendCode }"
                  size="small">
                  {{ countdown > 0 ? `Resend (${countdown}s)` : 'Get Code' }}
                </el-button>
              </template>
            </el-input>
          </div>
        </el-form-item>

        <el-form-item label="Access Code" prop="verificationCode">
          <div class="verification-code-group">
            <el-input 
              v-model="verificationCode" 
              placeholder="Enter 6-digit verification code"
              maxlength="6"
              :disabled="emailVerified"
              @input="handleVerificationCodeInput"
              style="width: 200px">
            </el-input>
            <el-button 
              @click="verifyCode"
              :disabled="!canVerify"
              :loading="verifying"
              :type="getVerifyButtonType"
              class="verify-button">
              {{ getVerifyButtonText }}
            </el-button>
            <div class="verification-tip" v-if="showVerificationTip">
              Please check your email for the verification code
            </div>
          </div>
        </el-form-item>

        <el-form-item label="Publication Venue" prop="venue">
          <el-select v-model="paperForm.venueType" placeholder="Select venue type" style="width: 150px; margin-right: 10px">
            <el-option label="Journal" value="journal"></el-option>
            <el-option label="Conference" value="conference"></el-option>
          </el-select>
          <el-input 
            v-model="paperForm.venueName" 
            placeholder="Enter journal/conference name"
            style="width: calc(100% - 160px)"
          ></el-input>
        </el-form-item>

        <el-form-item label="Publication Date" prop="publicationDate">
          <el-date-picker
            v-model="paperForm.publicationDate"
            type="date"
            placeholder="Select publication date">
          </el-date-picker>
        </el-form-item>

        <el-form-item label="Paper File" prop="paperFile">
          <el-upload
            class="upload-demo"
            drag
            action="/api/upload"
            :on-success="handleUploadSuccess"
            :on-error="handleUploadError"
            :before-upload="beforeUpload"
            accept=".pdf"
          >
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">Drop file here or <em>click to upload</em></div>
            <div class="el-upload__tip" slot="tip">Only PDF files are allowed</div>
          </el-upload>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm('paperForm')" :loading="loading">Submit</el-button>
          <el-button @click="resetForm('paperForm')">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'UploadPublishedPaper',
  data() {
    const validateVerificationCode = (rule, value, callback) => {
      if (!this.showVerificationInput) {
        callback();
      } else if (!value) {
        callback(new Error('Please enter verification code'));
      } else {
        callback();
      }
    };

    return {
      paperForm: {
        title: '',
        authors: '',
        abstract: '',
        correspondingEmail: '',
        venueType: '',
        venueName: '',
        publicationDate: '',
        paperFile: null
      },
      rules: {
        title: [
          { required: true, message: 'Please enter paper title', trigger: 'blur' }
        ],
        authors: [
          { required: true, message: 'Please enter authors', trigger: 'blur' }
        ],
        abstract: [
          { required: true, message: 'Please enter abstract', trigger: 'blur' }
        ],
        correspondingEmail: [
          { required: true, message: 'Please enter corresponding author email', trigger: 'blur' },
          { type: 'email', message: 'Please enter valid email address', trigger: 'blur' }
        ],
        verificationCode: [
          { validator: validateVerificationCode, trigger: 'blur' }
        ],
        venueType: [
          { required: true, message: 'Please select venue type', trigger: 'change' }
        ],
        venueName: [
          { required: true, message: 'Please enter venue name', trigger: 'blur' }
        ],
        publicationDate: [
          { required: true, message: 'Please select publication date', trigger: 'change' }
        ],
        paperFile: [
          { required: true, message: 'Please upload paper file', trigger: 'change' }
        ]
      },
      verificationCode: '',
      loading: false,
      verifying: false,
      countdown: 0,
      timer: null,
      showVerificationTip: false,
      emailVerified: false,
      isValidEmail: false
    }
  },
  computed: {
    canVerify() {
      return !this.emailVerified && this.verificationCode.length === 6 && !this.verifying
    },
    getVerifyButtonType() {
      if (this.emailVerified) return 'success'
      if (this.verifying) return 'primary'
      if (this.verificationCode.length === 6) return 'primary'
      return ''
    },
    getVerifyButtonText() {
      if (this.emailVerified) return 'Verified'
      if (this.verifying) return 'Verifying'
      return 'Verify'
    },
    canSendCode() {
      const result = this.paperForm.correspondingEmail && this.isValidEmail && !this.countdown && !this.emailVerified
      console.log('canSendCode conditions:', {
        hasEmail: Boolean(this.paperForm.correspondingEmail),
        isValidEmail: this.isValidEmail,
        noCountdown: !this.countdown,
        notVerified: !this.emailVerified,
        finalResult: result
      })
      return result
    }
  },
  watch: {
    'paperForm.correspondingEmail': {
      immediate: true,
      handler(value) {
        this.isValidEmail = this.validateEmail(value)
        console.log('Watch triggered:', {
          value,
          isValidEmail: this.isValidEmail
        })
      }
    }
  },
  methods: {
    handleEmailInput(value) {
      const trimmedValue = value.trim()
      this.paperForm.correspondingEmail = trimmedValue
      this.isValidEmail = this.validateEmail(trimmedValue)
      console.log('Email validation:', {
        email: trimmedValue,
        isValid: this.isValidEmail,
        correspondingEmail: this.paperForm.correspondingEmail
      })
      if (this.emailVerified) {
        this.emailVerified = false
        this.verificationCode = ''
      }
    },
    handleVerificationCodeInput(value) {
      // 只允许输入数字
      this.verificationCode = value.replace(/[^\d]/g, '')
    },
    validateEmail(email) {
      if (!email) return false
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      return emailRegex.test(email)
    },
    async verifyCode() {
      if (!this.verificationCode) {
        this.$message.error('Please enter verification code')
        return
      }

      this.verifying = true
      try {
        const response = await this.$http.post('/api/published-papers/verify-code', {
          email: this.paperForm.correspondingEmail,
          code: this.verificationCode
        })

        if (response.data.code === 1000) {
          this.emailVerified = true
          this.$message.success('Email verified successfully')
        } else {
          this.$message.error(response.data.msg || 'Invalid verification code')
        }
      } catch (error) {
        this.$message.error('Verification failed: ' + (error.response?.data?.msg || error.message))
      } finally {
        this.verifying = false
      }
    },
    sendVerificationCode() {
      if (!this.paperForm.correspondingEmail || !this.isValidEmail) {
        this.$message.error('Please enter a valid email address')
        return
      }

      this.loading = true
      this.$http.post('/api/published-papers/verify-email', {
        email: this.paperForm.correspondingEmail
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
        this.loading = false
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
    resetForm(formName) {
      this.$refs[formName].resetFields()
      this.verificationCode = ''
      this.showVerificationTip = false
      this.emailVerified = false
      this.isValidEmail = false
      if (this.timer) {
        clearInterval(this.timer)
        this.countdown = 0
      }
    },
    handleUploadSuccess(response, file) {
      this.paperForm.paperFile = file;
      this.$message.success('File uploaded successfully');
    },
    handleUploadError() {
      this.$message.error('File upload failed');
    },
    beforeUpload(file) {
      const isPDF = file.type === 'application/pdf';
      const isLt50M = file.size / 1024 / 1024 < 50;

      if (!isPDF) {
        this.$message.error('Only PDF files are allowed');
        return false;
      }
      if (!isLt50M) {
        this.$message.error('File size cannot exceed 50MB');
        return false;
      }

      // 检查文件名是否包含特殊字符
      const fileName = file.name;
      const reg = /[\\/:*?"<>|]/g;
      if (reg.test(fileName)) {
        this.$message.error('File name cannot contain special characters: \\ / : * ? " < > |');
        return false;
      }

      return true;
    },
    submitForm(formName) {
      this.$refs[formName].validate(async (valid) => {
        if (!valid) {
          return false;
        }

        if (!this.emailVerified) {
          // 如果邮箱未验证，先验证邮箱
          await this.verifyCode();
          if (!this.emailVerified) {
            return false;
          }
        }

        // 邮箱已验证，提交表单
        this.submitPaper();
      });
    },
    submitPaper() {
      // 显示加载状态
      this.loading = true;

      // 构建表单数据
      const formData = new FormData();
      formData.append('title', this.paperForm.title);
      formData.append('authors', this.paperForm.authors);
      formData.append('abstract', this.paperForm.abstract);
      formData.append('correspondingEmail', this.paperForm.correspondingEmail);
      formData.append('venueType', this.paperForm.venueType);
      formData.append('venueName', this.paperForm.venueName);
      formData.append('publicationDate', this.paperForm.publicationDate);
      if (this.paperForm.paperFile) {
        formData.append('paperFile', this.paperForm.paperFile.raw);
      }

      // 调用后端API提交论文
      this.$http.post('/api/published-papers/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then(response => {
        if (response.data.code === 1000) {
          this.$message.success('Paper submitted successfully');
          this.$router.push('/published-papers/my-papers');
        } else {
          this.$message.error(response.data.msg || 'Failed to submit paper');
        }
      }).catch(error => {
        this.$message.error('Failed to submit paper: ' + (error.response?.data?.msg || error.message));
      }).finally(() => {
        this.loading = false;
      });
    }
  },
  beforeDestroy() {
    if (this.timer) {
      clearInterval(this.timer);
    }
  }
}
</script>

<style lang="less" scoped>
.upload-paper {
  padding: 40px 0;
  
  .container {
    width: 80%;
    max-width: 1200px;
    margin: 0 auto;
    background-color: #fff;
    padding: 40px;
    border-radius: 8px;
    box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);

    h1 {
      color: #2c3e50;
      margin-bottom: 30px;
      text-align: center;
    }
  }

  .email-verification-group {
    display: flex;
    align-items: center;

    /deep/ .el-input-group__append {
      padding: 0;
      
      .el-button {
        margin: 0;
        border: none;
        height: 100%;
        min-width: 90px;

        &.is-primary {
          background-color: #409EFF;
          color: white;
          
          &:hover {
            background-color: #66b1ff;
          }
        }
      }
    }
  }

  .verification-code-group {
    display: flex;
    align-items: flex-start;
    gap: 10px;

    .verify-button {
      margin-left: 10px;
    }

    .verification-tip {
      font-size: 12px;
      color: #909399;
      margin-top: 5px;
      position: absolute;
      top: 100%;
      left: 0;
    }
  }

  .el-upload {
    width: 100%;
  }

  .el-upload-dragger {
    width: 100%;
  }
}
</style> 