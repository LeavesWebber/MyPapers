<template>
  <div class="published-paper-submit">
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>已出版论文上传</span>
      </div>
      
      <el-form :model="form" :rules="rules" ref="form" label-width="120px">
        <!-- 论文基本信息 -->
        <el-form-item label="论文标题" prop="title">
          <el-input v-model="form.title" placeholder="请输入论文标题"></el-input>
        </el-form-item>

        <el-form-item label="作者" prop="authors">
          <el-tag
            v-for="tag in form.authorTags"
            :key="tag"
            closable
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
          >
          </el-input>
          <el-button v-else class="button-new-tag" size="small" @click="showInput"
            >+ 添加作者</el-button
          >
        </el-form-item>

        <el-form-item label="关键词" prop="keywords">
          <el-tag
            v-for="tag in form.keywordTags"
            :key="tag"
            closable
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
          >
          </el-input>
          <el-button v-else class="button-new-tag" size="small" @click="showInput2"
            >+ 添加关键词</el-button
          >
        </el-form-item>

        <el-form-item label="论文摘要" prop="abstract">
          <el-input
            type="textarea"
            v-model="form.abstract"
            :rows="4"
            placeholder="请输入论文摘要"
          ></el-input>
        </el-form-item>

        <!-- 论文类型选择 -->
        <el-form-item label="论文类型" prop="paper_type">
          <el-radio-group v-model="form.paper_type">
            <el-radio label="journal">期刊论文</el-radio>
            <el-radio label="conference">会议论文</el-radio>
          </el-radio-group>
        </el-form-item>

        <!-- 期刊论文特有字段 -->
        <template v-if="form.paper_type === 'journal'">
          <el-form-item label="期刊名称" prop="journal_name">
            <el-input v-model="form.journal_name" placeholder="请输入期刊名称"></el-input>
          </el-form-item>

          <el-form-item label="期数" prop="volume_issue">
            <el-input v-model="form.volume_issue" placeholder="例如: Volume: 34, Issue: 6"></el-input>
          </el-form-item>

          <el-form-item label="出版时间" prop="publication_date">
            <el-date-picker
              v-model="form.publication_date"
              type="date"
              placeholder="选择出版日期"
              format="yyyy-MM-dd"
              value-format="yyyy-MM-dd"
            ></el-date-picker>
          </el-form-item>
        </template>

        <!-- 会议论文特有字段 -->
        <template v-if="form.paper_type === 'conference'">
          <el-form-item label="会议名称" prop="conference_name">
            <el-input v-model="form.conference_name" placeholder="请输入会议名称"></el-input>
          </el-form-item>

          <el-form-item label="会议时间" prop="conference_date">
            <el-date-picker
              v-model="form.conference_date"
              type="daterange"
              range-separator="至"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              format="yyyy-MM-dd"
              value-format="yyyy-MM-dd"
            ></el-date-picker>
          </el-form-item>

          <el-form-item label="会议地点" prop="conference_location">
            <el-input v-model="form.conference_location" placeholder="请输入会议地点"></el-input>
          </el-form-item>
        </template>

        <!-- 共同字段 -->
        <el-form-item label="论文页数" prop="pages">
          <el-input v-model="form.pages" placeholder="例如: 266-271"></el-input>
        </el-form-item>

        <el-form-item label="ISSN号" prop="issn">
          <el-input v-model="form.issn" placeholder="请输入ISSN号"></el-input>
        </el-form-item>

        <el-form-item label="论文链接" prop="paper_link">
          <el-input v-model="form.paper_link" placeholder="请输入论文链接"></el-input>
        </el-form-item>

        <!-- 通讯作者邮箱验证 -->
        <el-form-item label="通讯作者邮箱" prop="corresponding_email">
          <el-input v-model="form.corresponding_email" placeholder="请输入通讯作者邮箱"></el-input>
        </el-form-item>

        <!-- 邮箱验证 -->
        <el-form-item label="邮箱验证" v-if="form.corresponding_email">
          <div class="email-verification-group">
            <el-input
              v-model="form.verification_code"
              placeholder="请输入验证码"
              style="width: 200px"
            ></el-input>
            <el-button
              type="primary"
              @click="sendVerificationCode"
              :disabled="countdown > 0"
            >
              {{ countdown > 0 ? `${countdown}s后重试` : '获取验证码' }}
            </el-button>
          </div>
          <el-button
            type="success"
            @click="verifyEmail"
            :disabled="!form.verification_code || emailVerified"
            style="margin-top: 10px"
          >
            {{ emailVerified ? '已验证' : '验证邮箱' }}
          </el-button>
        </el-form-item>

        <!-- 文件上传 -->
        <el-form-item label="论文文件" prop="file">
          <el-upload
            class="upload-demo"
            action="#"
            :auto-upload="false"
            :on-change="handleUploadChange"
            :before-upload="beforeUpload"
            :limit="1"
            :on-exceed="handleExceed"
            :file-list="fileList"
          >
            <el-button size="small" type="primary">点击上传</el-button>
            <div slot="tip" class="el-upload__tip">
              只能上传PDF文件，且不超过15MB
            </div>
          </el-upload>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="validateForm">提交</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { submitPublishedPaper } from "../../api/paper";
import { contractInstance } from "../../utils/contract";

export default {
  name: "PublishedPaperSubmit",
  data() {
    return {
      form: {
        title: "",
        authorTags: [],
        keywordTags: [],
        abstract: "",
        paper_type: "journal",
        journal_name: "",
        volume_issue: "",
        publication_date: "",
        conference_name: "",
        conference_date: [],
        conference_location: "",
        pages: "",
        issn: "",
        paper_link: "",
        corresponding_email: "",
        block_address: "",
        paper_transaction_address: "",
        hash: "",
        verification_code: '',
      },
      rules: {
        title: [{ required: true, message: "请输入论文标题", trigger: "blur" }],
        authorTags: [{ required: true, message: "请输入作者信息", trigger: "blur" }],
        keywordTags: [{ required: true, message: "请输入关键词", trigger: "blur" }],
        abstract: [{ required: true, message: "请输入摘要", trigger: "blur" }],
        paper_type: [{ required: true, message: "请选择论文类型", trigger: "change" }],
        journal_name: [
          { required: true, message: "请输入期刊名称", trigger: "blur" },
          { validator: this.validateJournalName, trigger: "blur" }
        ],
        volume_issue: [
          { required: true, message: "请输入期刊期数", trigger: "blur" },
          { pattern: /^Volume:\s*\d+,\s*Issue:\s*\d+$/, message: "格式应为 Volume: 数字, Issue: 数字", trigger: "blur" }
        ],
        publication_date: [{ required: true, message: "请选择出版日期", trigger: "change" }],
        conference_name: [
          { required: true, message: "请输入会议名称", trigger: "blur" },
          { validator: this.validateConferenceName, trigger: "blur" }
        ],
        conference_date: [{ required: true, message: "请选择会议日期", trigger: "change" }],
        conference_location: [{ required: true, message: "请输入会议地点", trigger: "blur" }],
        corresponding_email: [
          { required: true, message: "请输入通讯作者邮箱", trigger: "blur" },
          { type: "email", message: "请输入正确的邮箱地址", trigger: "blur" }
        ],
        block_address: [{ required: true, message: "请输入区块链地址", trigger: "blur" }],
        paper_transaction_address: [{ required: true, message: "请输入交易地址", trigger: "blur" }],
      },
      fileList: [],
      inputVisible: false,
      inputValue: "",
      inputVisible2: false,
      inputValue2: "",
      userInfo: {},
      emailVerified: false,
      countdown: 0,
      timer: null,
    };
  },
  methods: {
    // 检查MetaMask连接状态
    async checkMetaMaskConnection() {
      if (!window.ethereum) {
        this.$message.error("请安装MetaMask钱包插件");
        return false;
      }
      
      try {
        // 请求账户授权
        const accounts = await window.ethereum.request({
          method: "eth_requestAccounts",
        });
        
        if (accounts.length === 0) {
          this.$message.error("请连接MetaMask钱包");
          return false;
        }
        
        // 检查网络ID
        const chainId = await window.ethereum.request({
          method: "eth_chainId",
        });
        
        if (chainId !== "0x198") { // 408的十六进制
          this.$message.warning("请切换到Papers Chain网络 (ID: 408)");
          return false;
        }
        
        return true;
      } catch (error) {
        console.error("MetaMask连接检查失败:", error);
        this.$message.error(`MetaMask连接错误: ${error.message}`);
        return false;
      }
    },

    handleExceed(files, fileList) {
      this.$message.warning(`只能选择一个文件`);
    },

    async readFileAsArrayBuffer(file) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader();

        reader.onload = () => {
          resolve(reader.result);
        };

        reader.onerror = (error) => {
          reject(error);
        };

        reader.readAsArrayBuffer(file);
      });
    },

    // 调用智能合约存储哈希值
    async callSmartContract(fileHash) {
      // 检查MetaMask连接
      const isConnected = await this.checkMetaMaskConnection();
      if (!isConnected) {
        this.fileList = [];
        return;
      }
      
      try {
        const functionName = "storeHash";
        const functionArgs = [fileHash];
        
        this.userInfo.block_chain_address = this.userInfo.block_chain_address.toLowerCase();

        if (this.userInfo.block_chain_address !== window.ethereum.selectedAddress) {
          this.fileList = [];
          this.$message({
            message: "请使用您的账户地址",
            type: "warning",
          });
          return;
        }

        // 获取当前网络的gas价格
        const gasPrice = await window.ethereum.request({
          method: 'eth_gasPrice'
        });
        
        // 向上调整gas价格以确保交易能被确认
        const adjustedGasPrice = parseInt(gasPrice, 16) * 1.1;
        const hexGasPrice = '0x' + Math.floor(adjustedGasPrice).toString(16);

        const result = await contractInstance.methods[functionName](...functionArgs).send({
          from: window.ethereum.selectedAddress,
          gasPrice: hexGasPrice,
        });

        this.form.block_address = result.blockHash;
        this.form.paper_transaction_address = result.transactionHash;
      } catch (error) {
        this.fileList = [];
        this.$message({
          message: "区块链上传失败",
          type: "error",
        });
        console.error("Error:", error);
      }
    },

    // 文件上传前的验证
    beforeUpload(file) {
      const isPDF = file.raw.type === "application/pdf";
      const isLt15M = file.raw.size / 1024 / 1024 < 15;

      if (!isPDF) {
        this.$message.error("只能上传PDF格式的文件!");
        this.fileList = [];
      }
      if (!isLt15M) {
        this.$message.error("文件大小不能超过15MB!");
        this.fileList = [];
      }

      return isPDF && isLt15M;
    },

    // 处理文件上传变化
    async handleUploadChange(file, fileList) {
      if (this.beforeUpload(file) === false) {
        return;
      }
      this.fileList = fileList;
      try {
        const fileData = await this.readFileAsArrayBuffer(this.fileList[0].raw);
        const hashBuffer = await crypto.subtle.digest("SHA-256", fileData);
        const hashArray = Array.from(new Uint8Array(hashBuffer));
        const fileHash = hashArray
          .map((byte) => byte.toString(16).padStart(2, "0"))
          .join("");
        await this.callSmartContract(fileHash);
      } catch (error) {
        console.error("Error:", error);
      }
    },

    // 作者标签相关方法
    handleClose(tag) {
      this.form.authorTags.splice(this.form.authorTags.indexOf(tag), 1);
      this.$forceUpdate();
    },

    showInput() {
      this.inputVisible = true;
      this.$nextTick((_) => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },

    handleInputConfirm() {
      let inputValue = this.inputValue;
      if (inputValue) {
        this.form.authorTags.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = "";
    },

    // 关键词标签相关方法
    handleClose2(tag) {
      this.form.keywordTags.splice(this.form.keywordTags.indexOf(tag), 1);
      this.$forceUpdate();
    },

    showInput2() {
      this.inputVisible2 = true;
      this.$nextTick((_) => {
        this.$refs.saveTagInput2.$refs.input.focus();
      });
    },

    handleInputConfirm2() {
      let inputValue2 = this.inputValue2;
      if (inputValue2) {
        this.form.keywordTags.push(inputValue2);
      }
      this.inputVisible2 = false;
      this.inputValue2 = "";
    },

    // 表单验证和提交
    validateForm() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.onSubmit();
        } else {
          console.log("表单验证失败!");
        }
      });
    },

    // 重置表单
    resetForm() {
      this.$refs.form.resetFields();
      this.fileList = [];
      this.form.authorTags = [];
      this.form.keywordTags = [];
    },

    // 验证期刊名称
    validateJournalName(rule, value, callback) {
      if (this.form.paper_type === 'journal' && !value) {
        callback(new Error('请输入期刊名称'));
      } else {
        callback();
      }
    },

    // 验证会议名称
    validateConferenceName(rule, value, callback) {
      if (this.form.paper_type === 'conference' && !value) {
        callback(new Error('请输入会议名称'));
      } else {
        callback();
      }
    },

    // 发送验证码
    async sendVerificationCode() {
      if (!this.form.corresponding_email) {
        this.$message.error('请先输入邮箱地址');
        return;
      }
      
      try {
        const res = await this.$http.post('/api/email/verify', {
          email: this.form.corresponding_email
        });
        
        if (res.data.code === 1000) {
          this.$message.success('验证码已发送到您的邮箱');
          this.startCountdown();
        } else {
          this.$message.error(res.data.msg || '发送验证码失败');
        }
      } catch (error) {
        this.$message.error('发送验证码失败: ' + error.message);
      }
    },

    // 开始倒计时
    startCountdown() {
      this.countdown = 60;
      this.timer = setInterval(() => {
        if (this.countdown > 0) {
          this.countdown--;
        } else {
          clearInterval(this.timer);
        }
      }, 1000);
    },

    // 验证邮箱
    async verifyEmail() {
      if (!this.form.verification_code) {
        this.$message.error('请输入验证码');
        return;
      }

      try {
        const res = await this.$http.post('/api/email/verify-code', {
          email: this.form.corresponding_email,
          code: this.form.verification_code
        });

        if (res.data.code === 1000) {
          this.$message.success('邮箱验证成功');
          this.emailVerified = true;
        } else {
          this.$message.error(res.data.msg || '验证码错误');
        }
      } catch (error) {
        this.$message.error('验证失败: ' + error.message);
      }
    },

    // 修改提交方法
    async onSubmit() {
      if (!this.emailVerified) {
        this.$message.error('请先验证邮箱');
        return;
      }

      if (this.fileList.length === 0) {
        this.$message({
          message: "请上传论文文件",
          type: "warning",
        });
        return;
      }

      const formData = new FormData();
      formData.append("title", this.form.title);
      formData.append("authors", this.form.authorTags);
      formData.append("keywords", this.form.keywordTags);
      formData.append("abstract", this.form.abstract);
      formData.append("paper_type", this.form.paper_type);
      
      // 根据论文类型添加不同字段
      if (this.form.paper_type === "journal") {
        formData.append("journal_name", this.form.journal_name);
        formData.append("volume_issue", this.form.volume_issue);
        formData.append("publication_date", this.form.publication_date);
      } else {
        formData.append("conference_name", this.form.conference_name);
        formData.append("conference_date", this.form.conference_date);
        formData.append("conference_location", this.form.conference_location);
      }

      // 添加可选字段
      if (this.form.pages) {
        formData.append("pages", this.form.pages);
      }
      if (this.form.issn) {
        formData.append("issn", this.form.issn);
      }
      if (this.form.paper_link) {
        formData.append("paper_link", this.form.paper_link);
      }

      formData.append("corresponding_email", this.form.corresponding_email);
      formData.append("block_address", this.form.block_address);
      formData.append("paper_transaction_address", this.form.paper_transaction_address);

      // 计算并添加文件哈希
      try {
        const fileData = await this.readFileAsArrayBuffer(this.fileList[0].raw);
        const hashBuffer = await crypto.subtle.digest("SHA-256", fileData);
        const hashArray = Array.from(new Uint8Array(hashBuffer));
        const fileHash = hashArray
          .map((byte) => byte.toString(16).padStart(2, "0"))
          .join("");
        formData.append("hash", fileHash);
      } catch (error) {
        this.$message.error('计算文件哈希失败: ' + error.message);
        return;
      }

      formData.append("data", this.fileList[0].raw);

      try {
        const res = await submitPublishedPaper(formData);
        if (res.data.code === 1000) {
          this.$message({
            message: "提交成功",
            type: "success",
          });
          this.$router.push("/center/publishedPapers");
        } else {
          this.$message({
            message: res.data.msg || "提交失败",
            type: "error",
          });
        }
      } catch (error) {
        console.error("提交失败:", error);
        this.$message({
          message: "提交失败: " + error.message,
          type: "error",
        });
      }
    },
  },
  mounted() {
    this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
  },
  beforeDestroy() {
    if (this.timer) {
      clearInterval(this.timer);
    }
  }
};
</script>

<style lang="less" scoped>
.published-paper-submit {
  padding: 20px;
  
  .box-card {
    width: 100%;
    
    .el-upload__tip {
      line-height: 1.2;
      padding-top: 5px;
      color: #909399;
    }
  }
  
  .input-new-tag {
    width: 90px;
    margin-left: 8px;
    vertical-align: bottom;
  }
  
  .button-new-tag {
    margin-left: 8px;
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
  }

  .email-verification-group {
    display: flex;
    align-items: center;
    gap: 10px;
  }
}
</style> 