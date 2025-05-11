<template>
  <div class="upload-paper">
    <div class="container">
      <h1>Upload Published Paper</h1>
      <el-form :model="paperForm" :rules="rules" ref="paperForm" label-width="180px">
        <el-form-item label="Paper Type" prop="paperType">
          <el-radio-group v-model="paperForm.paperType">
            <el-radio label="journal">Journal Paper</el-radio>
            <el-radio label="conference">Conference Paper</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="Paper Title" prop="title">
          <el-input v-model="paperForm.title" placeholder="Enter the full title of your paper"></el-input>
        </el-form-item>

        <el-form-item label="Authors" prop="authors">
          <el-tag
            :key="tag"
            v-for="tag in paperForm.authorTags"
            closable
            :disable-transitions="false"
            @close="handleClose(tag)"
          >
            {{ tag }}
          </el-tag>
          <el-input
            class="input-new-tag"
            v-if="inputVisible"
            v-model="inputValue"
            ref="saveTagInput"
            size="small"
            @keyup.enter.native="handleInputConfirm"
            @blur="handleInputConfirm"
          >
          </el-input>
          <el-button
            v-else
            class="button-new-tag"
            size="small"
            @click="showInput"
          >+ Author Name</el-button>
        </el-form-item>

        <el-form-item label="Keywords" prop="keywords">
          <el-tag
            :key="tag"
            v-for="tag in paperForm.keywordTags"
            closable
            :disable-transitions="false"
            @close="handleClose2(tag)"
          >
            {{ tag }}
          </el-tag>
          <el-input
            class="input-new-tag"
            v-if="inputVisible2"
            v-model="inputValue2"
            ref="saveTagInput2"
            size="small"
            @keyup.enter.native="handleInputConfirm2"
            @blur="handleInputConfirm2"
          >
          </el-input>
          <el-button
            v-else
            class="button-new-tag"
            size="small"
            @click="showInput2"
          >+ Key Word</el-button>
        </el-form-item>

        <template v-if="paperForm.paperType === 'journal'">
          <el-form-item label="Journal Name" prop="journalName">
            <el-input v-model="paperForm.journalName" placeholder="Enter journal name"></el-input>
          </el-form-item>

          <el-form-item label="Volume & Issue" prop="volumeAndIssue">
            <el-input v-model="paperForm.volumeAndIssue" placeholder="e.g., Volume: 34, Issue: 6"></el-input>
          </el-form-item>

          <el-form-item label="Date of Publication" prop="publicationDate">
            <el-date-picker
              v-model="paperForm.publicationDate"
              type="date"
              placeholder="Select publication date">
            </el-date-picker>
          </el-form-item>
        </template>

        <template v-if="paperForm.paperType === 'conference'">
          <el-form-item label="Conference Name" prop="conferenceName">
            <el-input v-model="paperForm.conferenceName" placeholder="Enter conference name"></el-input>
          </el-form-item>

          <el-form-item label="Conference Date" prop="conferenceDate">
            <el-date-picker
              v-model="paperForm.conferenceDate"
              type="daterange"
              range-separator="to"
              start-placeholder="Start date"
              end-placeholder="End date"
              class="conference-date-picker">
            </el-date-picker>
          </el-form-item>

          <el-form-item label="Conference Location" prop="conferenceLocation">
            <el-input v-model="paperForm.conferenceLocation" placeholder="Enter conference location"></el-input>
          </el-form-item>
        </template>

        <el-form-item label="Pages" prop="pages">
          <el-input v-model="paperForm.pages" placeholder="e.g., 266-271"></el-input>
        </el-form-item>

        <el-form-item label="ISSN" prop="issn">
          <el-input v-model="paperForm.issn" placeholder="Enter ISSN number"></el-input>
        </el-form-item>

        <el-form-item label="Paper Link" prop="paperLink">
          <el-input v-model="paperForm.paperLink" placeholder="Enter paper URL"></el-input>
        </el-form-item>

        <el-form-item label="Corresponding Author Email" prop="correspondingEmail">
          <div class="email-verification-group">
            <el-input 
              v-model="paperForm.correspondingEmail" 
              placeholder="Enter corresponding author's email"
              @input="handleEmailInput"
              :disabled="emailVerified || sendingCode">
              <template slot="append">
                <el-button 
                  @click="sendVerificationCode" 
                  :disabled="!canSendCode || sendingCode"
                  :loading="sendingCode"
                  :class="{ 'is-primary': canSendCode }"
                  size="small">
                  {{ countdown > 0 ? `Resend (${countdown}s)` : (sendingCode ? 'Sending...' : 'Get Code') }}
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
              style="width: 230px">
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

        <el-form-item label="Paper File" prop="paperFile">
          <template v-if="!fileList.length">
            <el-upload
              class="upload-demo"
              ref="upload"
              :on-change="handleUploadChange"
              action=""
              :file-list="fileList"
              :auto-upload="false"
              :limit="1"
              :on-exceed="handleExceed"
              drag
            >
              <i class="el-icon-upload"></i>
              <div class="el-upload__text">Drop file here or <em>click to upload</em></div>
              <div class="el-upload__tip" slot="tip">Only PDF files are allowed (max 15MB)</div>
            </el-upload>
          </template>
          <template v-else>
            <div class="uploaded-file-info">
              <div class="file-info-item">
                <span class="label">文件名：</span>
                <span class="value">{{ fileList[0].name }}</span>
              </div>
              <div class="file-info-item">
                <span class="label">文件大小：</span>
                <span class="value">{{ formatFileSize(fileList[0].size) }}</span>
              </div>
              <div class="file-info-item">
                <span class="label">文件哈希：</span>
                <span class="value">{{ form.hash }}</span>
              </div>
              <div class="file-info-item">
                <span class="label">区块地址：</span>
                <span class="value">{{ form.block_address }}</span>
              </div>
              <div class="file-info-item">
                <span class="label">交易地址：</span>
                <span class="value">{{ form.paper_transaction_address }}</span>
              </div>
              <el-button 
                type="text" 
                @click="clearFile"
                class="clear-file-btn"
                :disabled="isProcessing"
              >
                重新选择文件
              </el-button>
            </div>
            <div v-if="isProcessing" class="transaction-status">
              <el-alert
                :title="transactionStatus.title"
                :type="transactionStatus.type"
                :description="transactionStatus.description"
                :closable="false"
                show-icon
              >
                <template v-if="transactionStatus.showLoading">
                  <el-progress 
                    :percentage="transactionStatus.progress" 
                    :status="transactionStatus.progressStatus"
                    :format="progressFormat"
                  />
                </template>
              </el-alert>
            </div>
          </template>
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
import { ERC20contractInstance } from "@/constant";
const contractInstance = ERC20contractInstance;

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
        paperType: 'journal',
        title: '',
        authors: '',
        authorTags: [],
        keywordTags: [],
        journalName: '',
        volumeAndIssue: '',
        publicationDate: '',
        conferenceName: '',
        conferenceDate: '',
        conferenceLocation: '',
        pages: '',
        issn: '',
        paperLink: '',
        correspondingEmail: '',
        venueType: '',
        venueName: '',
        paperFile: null
      },
      rules: {
        paperType: [
          { required: true, message: 'Please select paper type', trigger: 'change' }
        ],
        title: [
          { required: true, message: 'Please enter paper title', trigger: 'blur' },
          { min: 1, max: 200, message: 'Title should be between 1 and 200 characters', trigger: 'blur' }
        ],
        authorTags: [
          { required: true, message: 'Please enter at least one author', trigger: 'blur' }
        ],
        keywordTags: [
          { required: true, message: 'Please enter at least one keyword', trigger: 'blur' }
        ],
        journalName: [
          { required: true, message: 'Please enter journal name', trigger: 'blur' }
        ],
        volumeAndIssue: [
          { required: true, message: 'Please enter volume and issue', trigger: 'blur' }
        ],
        publicationDate: [
          { required: true, message: 'Please select publication date', trigger: 'change' }
        ],
        conferenceName: [
          { required: true, message: 'Please enter conference name', trigger: 'blur' }
        ],
        conferenceDate: [
          { required: true, message: 'Please select conference date', trigger: 'change' },
          { type: 'array', message: 'Please select start and end dates', trigger: 'change' }
        ],
        conferenceLocation: [
          { required: true, message: 'Please enter conference location', trigger: 'blur' }
        ],
        pages: [
          { required: false }
        ],
        issn: [
          { required: false },
          { pattern: /^[0-9]{4}-[0-9]{3}[0-9X]$/, message: 'Please enter a valid ISSN number', trigger: 'blur' }
        ],
        paperLink: [
          { required: false, type: 'url', message: 'Please enter a valid URL', trigger: 'blur' }
        ],
        correspondingEmail: [
          { required: true, message: 'Please enter corresponding author email', trigger: 'blur' },
          { type: 'email', message: 'Please enter valid email address', trigger: 'blur' }
        ],
        verificationCode: [
          { validator: validateVerificationCode, trigger: 'blur' }
        ],
        paperFile: [
          { 
            validator: (rule, value, callback) => {
              if (!this.fileList || this.fileList.length === 0) {
                callback(new Error('Please upload paper file'));
              } else {
                callback();
              }
            },
            trigger: 'change'
          }
        ]
      },
      verificationCode: '',
      loading: false,
      sendingCode: false,
      verifying: false,
      countdown: 0,
      timer: null,
      showVerificationTip: false,
      emailVerified: false,
      isValidEmail: false,
      fileList: [],
      form: {
        hash: "",
        block_address: "",
        paper_transaction_address: "",
      },
      inputVisible: false,
      inputValue: '',
      inputVisible2: false,
      inputValue2: '',
      isProcessing: false,
      transactionStatus: {
        title: '',
        type: 'info',
        description: '',
        showLoading: false,
        progress: 0,
        progressStatus: 'success'
      },
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
/*       console.log('canSendCode conditions:', {
        hasEmail: Boolean(this.paperForm.correspondingEmail),
        isValidEmail: this.isValidEmail,
        noCountdown: !this.countdown,
        notVerified: !this.emailVerified,
        finalResult: result
      }) */
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
      // Only allow numbers
      this.verificationCode = value.replace(/[^\d]/g, '')
    },
    validateEmail(email) {
      if (!email) return false
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      return emailRegex.test(email)
    },
    async verifyCode() {
      if (!this.verificationCode) {
        this.$message.error('请输入验证码');
        return;
      }

      this.verifying = true;
      try {
        const response = await this.$http.post('/mypapers/user/VerifyMail', {
          email: this.paperForm.correspondingEmail,
          code: this.verificationCode
        });

        if (response.data.code === 1000) {
          // 验证成功后，将验证信息存储在 localStorage 中
          const verificationData = {
            email: this.paperForm.correspondingEmail,
            code: this.verificationCode,
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
      if (!this.paperForm.correspondingEmail || !this.isValidEmail) {
        this.$message.error('Please enter a valid email address')
        return
      }

      this.sendingCode = true
      this.$http.post('/mypapers/user/SendMail', {
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
        this.sendingCode = false
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
    async handleUploadChange(file, fileList) {
      console.log('File info:', {
        name: file.name,
        type: file.type,
        size: file.size,
        raw: file.raw
      });
      
      if (this.beforeUpload(file) === false) {
        console.log('File validation failed');
        return;
      }
      this.fileList = fileList;
      this.paperForm.paperFile = file.raw;
      try {
        // 检查file.raw是否存在
        if (!file.raw) {
          this.$message.error('文件对象无效，请重新选择文件');
          this.fileList = [];
          return;
        }
        
        // 计算文件的哈希值
        const fileData = await this.readFileAsArrayBuffer(file.raw);
        // 检查 crypto API 是否可用
        if (!window.crypto || !window.crypto.subtle) {
          this.$message.error('您的浏览器不支持Web Cryptography API，请使用最新版Chrome或Firefox浏览器');
          this.fileList = [];
          return;
        }
        
        const hashBuffer = await window.crypto.subtle.digest("SHA-256", fileData);
        const hashArray = Array.from(new Uint8Array(hashBuffer));
        const fileHash = hashArray
          .map((byte) => byte.toString(16).padStart(2, "0"))
          .join("");
        console.log("fileHash:", fileHash);
        
        // 检查哈希值是否已存在
        try {
          const existingAddress = await contractInstance.methods.getRecipientByHash(fileHash).call();
          if (existingAddress !== '0x0000000000000000000000000000000000000000') {
            this.$message.error('该论文已经被上传过，请勿重复上传');
            this.fileList = [];
            return;
          }
        } catch (error) {
          console.error('检查哈希值失败:', error);
          this.$message.error('检查论文哈希值失败，请稍后重试');
          this.fileList = [];
          return;
        }

        // 调用合约函数
        this.form.hash = fileHash;
        await this.callSmartContract(fileHash);
      } catch (error) {
        console.error("Error:", error);
        this.$message.error('文件处理失败，请重试');
        this.fileList = [];
      }
    },
    beforeUpload(file) {
      // 检查文件类型
      const isPDF = (file.type === "application/pdf") || 
                   (typeof file.name === 'string' && file.name.toLowerCase().endsWith('.pdf'));
      const isLt15M = file.size / 1024 / 1024 < 15;

      if (!isPDF) {
        this.$message.error("Upload file must be PDF format!");
        // 清除添加文件
        this.fileList = [];
        return false;
      }
      if (!isLt15M) {
        this.$message.error("Upload file size can not exceed 15MB!");
        // 清除添加文件
        this.fileList = [];
        return false;
      }

      return true;
    },
    readFileAsArrayBuffer(file) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = () => resolve(reader.result);
        reader.onerror = reject;
        reader.readAsArrayBuffer(file);
      });
    },
    async callSmartContract(fileHash) {
      try {
        this.isProcessing = true;
        this.updateTransactionStatus({
          title: '正在连接钱包...',
          type: 'info',
          description: '请确保MetaMask钱包已安装并已连接到正确的网络',
          showLoading: true,
          progress: 0
        });

        // 检查 MetaMask 是否已连接
        if (!window.ethereum || !window.ethereum.selectedAddress) {
          this.updateTransactionStatus({
            title: 'MetaMask未连接',
            type: 'warning',
            description: '请先连接MetaMask钱包',
            showLoading: false
          });
          throw new Error('请先连接 MetaMask');
        }

        // 获取当前账户
        const accounts = await window.ethereum.request({ method: 'eth_accounts' });
        if (!accounts || accounts.length === 0) {
          this.updateTransactionStatus({
            title: '未找到账户',
            type: 'warning',
            description: '未找到MetaMask账户，请确保已登录',
            showLoading: false
          });
          throw new Error('未找到 MetaMask 账户');
        }

        // 检查当前网络
        const chainId = await window.ethereum.request({ method: 'eth_chainId' });
        if (chainId !== '0x198') {
          this.updateTransactionStatus({
            title: '网络错误',
            type: 'warning',
            description: '请切换到正确的网络 (Chain ID: 408)',
            showLoading: false
          });
          throw new Error('请切换到正确的网络 (Chain ID: 408)');
        }

        this.updateTransactionStatus({
          title: '正在检查文件...',
          type: 'info',
          description: '正在验证文件哈希值',
          showLoading: true,
          progress: 30
        });

        // 检查哈希值是否已存在
        try {
          const existingAddress = await contractInstance.methods.getRecipientByHash(fileHash).call();
          if (existingAddress !== '0x0000000000000000000000000000000000000000' && existingAddress !== accounts[0]) {
            this.updateTransactionStatus({
              title: '文件已存在',
              type: 'warning',
              description: '该论文已经被其他地址上传过，请勿重复上传',
              showLoading: false
            });
            this.fileList = [];
            return;
          }
        } catch (error) {
          this.updateTransactionStatus({
            title: '检查失败',
            type: 'error',
            description: '检查论文哈希值失败，请稍后重试',
            showLoading: false
          });
          this.fileList = [];
          return;
        }

        this.updateTransactionStatus({
          title: '准备交易...',
          type: 'info',
          description: '即将打开MetaMask钱包，请确认交易',
          showLoading: true,
          progress: 60
        });

        // 获取当前gas价格
        const gasPrice = await window.ethereum.request({
          method: 'eth_gasPrice',
          params: []
        });
        const hexGasPrice = gasPrice.result;

        // 调用合约函数
        this.updateTransactionStatus({
          title: 'MetaMask钱包正在确认交易，请稍等...',
          type: 'info',
          description: '请在MetaMask钱包中确认交易',
          showLoading: true,
          progress: 80
        });

        const result = await contractInstance.methods.storeHash(fileHash).send({
          from: accounts[0],
          gas: "0x186A0",
          gasPrice: hexGasPrice,
        });
        
        this.updateTransactionStatus({
          title: '交易成功',
          type: 'success',
          description: '文件已成功上传到区块链',
          showLoading: true,
          progress: 100,
          progressStatus: 'success'
        });

        this.form.block_address = result.blockHash;
        this.form.paper_transaction_address = result.transactionHash;

        // 3秒后清除状态提示
        setTimeout(() => {
          this.isProcessing = false;
          this.resetTransactionStatus();
        }, 3000);

      } catch (error) {
        this.fileList = [];
        let errorMessage = '交易失败，请检查合约状态或联系管理员';
        
        if (error.message.includes("user denied")) {
          errorMessage = "您已取消交易";
        } else if (error.message.includes("insufficient funds")) {
          errorMessage = "账户余额不足，请确保有足够的 ETH 支付 gas 费用";
        } else if (error.message.includes("Internal JSON-RPC error")) {
          errorMessage = "区块链节点连接问题，请检查网络连接或联系管理员";
        } else if (error.message.includes("Hash already stored")) {
          errorMessage = "该文件哈希值已被存储，请勿重复上传";
        } else if (error.message.includes("请切换到正确的网络")) {
          errorMessage = "请切换到正确的网络 (Chain ID: 408)";
        }

        this.updateTransactionStatus({
          title: '交易失败',
          type: 'error',
          description: errorMessage,
          showLoading: false
        });

        // 5秒后清除错误提示
        setTimeout(() => {
          this.isProcessing = false;
          this.resetTransactionStatus();
        }, 5000);
      }
    },
    handleExceed(files, fileList) {
      this.$message.warning(
        `当前限制选择 1 个文件，本次选择了 ${files.length} 个文件，共选择了 ${files.length + fileList.length} 个文件`
      );
    },
    async submitForm(formName) {
      this.$refs[formName].validate(async (valid) => {
        if (!valid) {
          return false;
        }

        // 检查作者是否为空
        if (!this.paperForm.authorTags || this.paperForm.authorTags.length === 0) {
          this.$message.error('请至少添加一个作者');
          return false;
        }

        // 检查关键词是否为空
        if (!this.paperForm.keywordTags || this.paperForm.keywordTags.length === 0) {
          this.$message.error('请至少添加一个关键词');
          return false;
        }

        // 即使前端显示已验证，也要在后端重新验证
        if (!this.emailVerified || !this.verificationCode) {
          this.$message.error('请先完成邮箱验证');
          return false;
        }

        // 在提交前再次验证邮箱
        try {
          const response = await this.$http.post('/mypapers/user/VerifyMail', {
            email: this.paperForm.correspondingEmail,
            code: this.verificationCode
          });

          if (response.data.code !== 1000) {
            this.$message.error('邮箱验证已过期，请重新验证');
            this.emailVerified = false;
            return false;
          }
        } catch (error) {
          this.$message.error('邮箱验证失败，请重新验证');
          this.emailVerified = false;
          return false;
        }

        // 检查MetaMask连接状态
        if (!await this.checkMetaMaskConnection()) {
          return false;
        }

        // 检查文件是否上传
        if (!this.fileList.length) {
          this.$message.error('Please upload your paper');
          return false;
        }

        // 邮箱已验证，提交表单
        this.submitPaper();
      });
    },
    async checkMetaMaskConnection() {
      if (!window.ethereum) {
        this.$message.error('请安装MetaMask钱包插件');
        return false;
      }
      
      try {
        // 请求账户授权
        const accounts = await window.ethereum.request({
          method: 'eth_requestAccounts'
        });
        
        if (accounts.length === 0) {
          this.$message.error('请连接MetaMask钱包');
          return false;
        }
        
        // 检查网络ID
        const chainId = await window.ethereum.request({
          method: 'eth_chainId'
        });
        
        if (chainId !== '0x198') { // 408的十六进制
          this.$message.warning('请切换到Papers Chain网络 (ID: 408)');
          return false;
        }
        
        return true;
      } catch (error) {
        console.error('MetaMask连接检查失败:', error);
        this.$message.error(`MetaMask连接错误: ${error.message}`);
        return false;
      }
    },
    submitPaper() {
      // Show loading state
      this.loading = true;

      // 获取验证token
      const verificationData = JSON.parse(localStorage.getItem('emailVerification') || '{}');
      if (!verificationData.token || Date.now() > verificationData.expiresAt * 1000) {
        this.$message.error('邮箱验证已过期，请重新验证');
        this.loading = false;
        return;
      }

      // Build form data
      const formData = new FormData();
      formData.append('title', this.paperForm.title);
      formData.append('authors', this.paperForm.authorTags.join(','));
      formData.append('keywords', this.paperForm.keywordTags.join(','));
      formData.append('corresponding_email', this.paperForm.correspondingEmail);
      formData.append('paper_type', this.paperForm.paperType);
      
      // Add paper type specific fields
      if (this.paperForm.paperType === 'journal') {
        formData.append('journal_name', this.paperForm.journalName);
        formData.append('volume_and_issue', this.paperForm.volumeAndIssue);
        formData.append('publication_date', this.paperForm.publicationDate);
      } else {
        formData.append('conference_name', this.paperForm.conferenceName);
        formData.append('conference_date', this.paperForm.conferenceDate);
        formData.append('conference_location', this.paperForm.conferenceLocation);
      }
      
      // Add optional fields
      if (this.paperForm.pages) {
        formData.append('pages', this.paperForm.pages);
      }
      if (this.paperForm.issn) {
        formData.append('issn', this.paperForm.issn);
      }
      if (this.paperForm.paperLink) {
        formData.append('paper_link', this.paperForm.paperLink);
      }

      // Add blockchain related fields
      formData.append('hash', this.form.hash);
      formData.append('block_address', this.form.block_address);
      formData.append('paper_transaction_address', this.form.paper_transaction_address);

      if (this.fileList[0]) {
        formData.append('data', this.fileList[0].raw);
      }

      // Call backend API to submit paper
      this.$http.post('/mypapers/paper/uploadPublished', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
          'X-Verification-Token': verificationData.token,
          'Authorization': 'Bearer ' + localStorage.getItem('token')
        }
      }).then(response => {
        if (response.data.code === 1000) {
          this.$message.success('论文提交成功');
          this.$router.push('/center/papers');
        } else {
          this.$message.error(response.data.msg || '提交失败');
        }
      }).catch(error => {
        this.$message.error('提交失败: ' + (error.response?.data?.msg || error.message));
      }).finally(() => {
        this.loading = false;
      });
    },
    handleClose(tag) {
      this.paperForm.authorTags.splice(this.paperForm.authorTags.indexOf(tag), 1);
    },
    handleClose2(tag) {
      this.paperForm.keywordTags.splice(this.paperForm.keywordTags.indexOf(tag), 1);
    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick((_) => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    showInput2() {
      this.inputVisible2 = true;
      this.$nextTick((_) => {
        this.$refs.saveTagInput2.$refs.input.focus();
      });
    },
    handleInputConfirm() {
      let inputValue = this.inputValue;
      if (inputValue) {
        this.paperForm.authorTags.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = '';
    },
    handleInputConfirm2() {
      let inputValue2 = this.inputValue2;
      if (inputValue2) {
        this.paperForm.keywordTags.push(inputValue2);
      }
      this.inputVisible2 = false;
      this.inputValue2 = '';
    },
    mounted() {
      // 检查是否有已存储的验证信息
      const storedVerification = localStorage.getItem('emailVerification');
      if (storedVerification) {
        const verificationData = JSON.parse(storedVerification);
        // 检查验证是否过期
        if (Date.now() < verificationData.expiresAt * 1000) {
          this.paperForm.correspondingEmail = verificationData.email;
          this.verificationCode = verificationData.code;
          this.emailVerified = true;
        } else {
          // 验证已过期，清除存储
          localStorage.removeItem('emailVerification');
        }
      }
    },
    formatFileSize(bytes) {
      if (bytes === 0) return '0 B';
      const k = 1024;
      const sizes = ['B', 'KB', 'MB', 'GB'];
      const i = Math.floor(Math.log(bytes) / Math.log(k));
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
    },
    clearFile() {
      this.fileList = [];
      this.form.hash = '';
      this.form.block_address = '';
      this.form.paper_transaction_address = '';
      this.paperForm.paperFile = null;
    },
    updateTransactionStatus(status) {
      this.transactionStatus = {
        ...this.transactionStatus,
        ...status
      };
    },
    resetTransactionStatus() {
      this.transactionStatus = {
        title: '',
        type: 'info',
        description: '',
        showLoading: false,
        progress: 0,
        progressStatus: 'success'
      };
    },
    progressFormat(percentage) {
      return percentage === 100 ? '完成' : `${percentage}%`;
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
    position: relative;
    margin-bottom: 20px;

    .verify-button {
      margin-left: 10px;
      min-width: 100px;
    }

    .verification-tip {
      font-size: 12px;
      color: #909399;
      position: absolute;
      top: 100%;
      left: 0;
      z-index: 10;
      background-color: #fff;
      padding: 2px 5px;
      box-shadow: 0 2px 4px rgba(0,0,0,0.1);
      border-radius: 4px;
      white-space: nowrap;
    }
  }

  .el-upload {
    width: 100%;
    margin-top: 10px;
  }

  .el-upload-dragger {
    width: 100%;
  }

  .conference-date-picker {
    :deep(.el-range-separator) {
      padding: 0 8px;  // 增加分隔符两侧的间距
      width: auto !important;  // 允许分隔符宽度自适应
    }
    
    :deep(.el-range-input) {
      // 调整输入框的宽度，确保有足够空间显示占位符
      width: 130px !important;
    }
  }

  .el-tag {
    margin-right: 10px;
    margin-bottom: 10px;
  }

  .button-new-tag {
    margin-right: 10px;
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
  }

  .input-new-tag {
    width: 90px;
    margin-right: 10px;
    vertical-align: bottom;
  }

  .el-form-item {
    margin-bottom: 22px;
  }

  .el-button--primary {
    margin-right: 10px;
  }

  .uploaded-file-info {
    padding: 20px;
    background-color: #f5f7fa;
    border-radius: 4px;
    border: 1px dashed #d9d9d9;
    
    .file-info-item {
      margin-bottom: 10px;
      line-height: 1.5;
      
      .label {
        color: #606266;
        font-weight: bold;
        margin-right: 10px;
      }
      
      .value {
        color: #303133;
        word-break: break-all;
      }
    }
    
    .clear-file-btn {
      margin-top: 10px;
      color: #f56c6c;
      
      &:hover {
        color: #f78989;
      }
    }
  }

  .transaction-status {
    margin-top: 20px;
    
    .el-alert {
      margin-bottom: 10px;
    }
    
    .el-progress {
      margin-top: 10px;
    }
  }
}
</style>