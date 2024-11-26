<template>
  <div class="box">
    <div class="box1">
      <el-form ref="form" :model="form" label-width="200px" :rules="rules">
        <el-form-item label="Upload your paper here:">
          <el-upload
            class="upload-demo"
            ref="upload"
            :on-change="handleUploadChange"
            action=""
            :file-list="fileList"
            :auto-upload="false"
            :limit="1"
            :on-exceed="handleExceed"
          >
            <el-button slot="trigger" size="small" type="primary"
              >Upload</el-button
            >
          </el-upload>
        </el-form-item>
        <el-form-item label="Title" prop="title">
          <el-input v-model="form.title" style="width: 600px"></el-input>
        </el-form-item>
        
        <el-form-item label="Abstract" prop="abstract">
          <el-input
            style="width: 600px"
            type="textarea"
            v-model="form.abstract"
            :rows="10"
          ></el-input>
        </el-form-item>

        
        </el-form-item>
        <el-form-item label="Paper Type" prop="paper_type">
          <el-select
            style="width: 200px"
            v-model="form.paper_type"
            placeholder="Paper Type"
          >
            <el-option label="Regular Paper" value="Regular Paper"></el-option>
            <el-option label="Survey Paper" value="Survey Paper"></el-option>
            <el-option label="Review Paper" value="Review Paper"></el-option>
            <el-option label="Short Paper" value="Short Paper"></el-option>
            <el-option
              label="Technical Note"
              value="Technical Note"
            ></el-option>
            <el-option label="Letter" value="Letter"></el-option>
            <el-option
              label="Conference Paper"
              value="Conference Paper"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Subject Category" prop="subject_category">
          <el-select
            style="width: 200px"
            v-model="form.subject_category"
            placeholder="Subject Category"
          >
            <el-option
              v-for="option in form.options"
              :key="option.value"
              :label="option.label"
              :value="option.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Authors" prop="authorTags">
          <el-tag
            :key="tag"
            v-for="tag in form.authorTags"
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
            >+ Author Name</el-button
          >
        </el-form-item>

        <!-- <el-form-item label="Keywords" prop="keywords">
          <el-input style="width: 360px" v-model="form.keywords"></el-input>
        </el-form-item> -->
        <el-form-item label="Keywords" prop="keywordTags">
          <el-tag
            :key="tag"
            v-for="tag in form.keywordTags"
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
            >+ Key Word</el-button
          >
        </el-form-item>
        <el-form-item label="Corresponding Author" prop="cor_author">
          <el-input style="width: 360px" v-model="form.cor_author"></el-input>
        </el-form-item>
        
        <el-form-item label="Unique Contribution" prop="unique_contribution">
          <el-input
            style="width: 360px"
            type="textarea"
            v-model="form.unique_contribution"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="validateForm">Submit</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script>
import { submitPaper, getConferenceDetail } from "../../api";
import { ERC20contractInstance } from "@/constant";
const contractInstance = ERC20contractInstance;
export default {
  data() {
    var validateWordCount = (rule, value, callback) => {
      const wordCount = value.trim().split(/\s+/).length;
      if (wordCount < 1 || wordCount > 20) {
        callback(new Error("Should contain 1 to 20 words"));
      } else {
        callback();
      }
    };
    var validateAbstractWordCount = (rule, value, callback) => {
      const wordCount = value.trim().split(/\s+/).length;
      if (wordCount < 150 || wordCount > 300) {
        callback(new Error("Should contain 150 to 250 words"));
      } else {
        callback();
      }
    };
    return {
      form: {
        conference_id: 0,
        title: "",
        paper_type: "",
        abstract: "",
        authors: [],
        keywords: [],
        subject_category: "",
        informed_consent: "",
        animal_subjects: "",
        cor_author: "",
        manuscript_type: "",
        unique_contribution: "",
        hash: "",
        block_address: "",
        paper_transaction_address: "",
        authorTags: [],
        keywordTags: [],
        options: [{ label: "", value: "" }],
      },
      rules: {
        title: [
          { required: true, message: "Please input title", trigger: "blur" },
          { validator: validateWordCount, trigger: "blur" },
        ],
        paper_type: [
          {
            required: true,
            message: "Please select paper type",
            trigger: "blur",
          },
        ],
        abstract: [
          { required: true, message: "Please input abstract", trigger: "blur" },
          { validator: validateAbstractWordCount, trigger: "blur" },
        ],
        block_address: [
          {
            required: true,
            message: "Please input block address",
            trigger: "blur",
          },
          {
            // 0x7d2e0f9c8856a52fbfcca0fe3db1891c53c4f9cc7176127ea4a97cfb26ae525a
            pattern: /^0x[a-fA-F0-9]{64}$/,
            message: "Please input valid block address",
            trigger: "blur",
          },
        ],
        authorTags: [
          { required: true, message: "Please input authors", trigger: "blur" },
        ],
        cor_author: [
          {
            required: true,
            message: "Please input corresponding author",
            trigger: "blur",
          },
          // {
          //   // 任意类型的邮箱都可以
          //   pattern:
          //     /^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/,
          //   message: "Please input valid email address",
          //   trigger: "blur",
          // },
        ],
        // keywords: [
        //   { required: true, message: "Please input keywords", trigger: "blur" },
        // ],
        keywordTags: [
          { required: true, message: "Please input keywords", trigger: "blur" },
        ],
        subject_category: [
          {
            required: true,
            message: "Please input subject category",
            trigger: "blur",
          },
        ],
        animal_subjects: [
          {
            required: true,
            message: "Please select informed consent",
            trigger: "blur",
          },
        ],
        unique_contribution: [
          {
            required: true,
            message: "Please input unique contribution",
            trigger: "blur",
          },
        ],
      },
      fileList: [],
      // authorTags: [],
      inputVisible: false,
      inputValue: "",
      // keywordTags: [],
      inputVisible2: false,
      inputValue2: "",
      userInfo: {},
    };
  },

  methods: {
    handleExceed(files, fileList) {
      this.$message.warning(`Only one file can be selected`);
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
    // 合约函数
    async callSmartContract(fileHash) {
      try {
        // 调用智能合约函数
        const functionName = "storeHash";
        const functionArgs = [fileHash];
        console.log(
          "window.ethereum.selectedAddress:",
          window.ethereum.selectedAddress
        );
        // 把this.userInfo.block_chain_address中的大小字母转换成小写
        this.userInfo.block_chain_address =
          this.userInfo.block_chain_address.toLowerCase();

        if (
          this.userInfo.block_chain_address !== window.ethereum.selectedAddress
        ) {
          this.fileList = [];
          this.$message({
            message: "Please use the your account_address",
            type: "warning",
          });
          return;
        }

        const result = await contractInstance.methods[functionName](
          ...functionArgs
        ).send({
          from: window.ethereum.selectedAddress,
          gasPrice: "0",
        });
        // 输出结果
        console.log("Transaction result:", result);
        this.form.block_address = result.blockHash;
        this.form.paper_transaction_address = result.transactionHash;
      } catch (error) {
        this.fileList = [];
        this.$message({
          message: "Upload to blockchain failed",
          type: "error",
        });
        // 处理错误
        console.error("Error:", error);
      }
    },
    validateForm() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.onSubmit();
        } else {
          console.log("error submit!!");
        }
      });
    },
    async onSubmit() {
      if (this.fileList.length == 0) {
        this.$message({
          message: "Please upload your paper",
          type: "warning",
        });
        return;
      }
      console.log(this.form);
      const formData = new FormData();
      formData.append("conference_id", this.$route.query.id);
      formData.append("paper_type", this.form.paper_type);
      formData.append("title", this.form.title);
      formData.append("abstract", this.form.abstract);
      formData.append("authors", this.form.authorTags);
      // formData.append("keywords", this.form.keywords);
      formData.append("keywords", this.form.keywordTags);
      formData.append("subject_category", this.form.subject_category);
      formData.append("informed_consent", this.form.informed_consent);
      formData.append("animal_subjects", this.form.animal_subjects);
      formData.append("cor_author", this.form.cor_author);
      formData.append("manuscript_type", this.form.manuscript_type);
      formData.append("unique_contribution", this.form.unique_contribution);
      // 计算文件的哈希值
      const fileData = await this.readFileAsArrayBuffer(this.fileList[0].raw);
      const hashBuffer = await crypto.subtle.digest("SHA-256", fileData);
      const hashArray = Array.from(new Uint8Array(hashBuffer));
      const fileHash = hashArray
        .map((byte) => byte.toString(16).padStart(2, "0"))
        .join("");
      formData.append("hash", fileHash);
      formData.append("block_address", this.form.block_address);
      formData.append(
        "paper_transaction_address",
        this.form.paper_transaction_address
      );
      formData.append("data", this.fileList[0].raw);
      submitPaper(formData).then((res) => {
        console.log(res);
        if (res.data.code === 1000) {
          this.$message({
            message: "Submit successfully",
            type: "success",
          });
          this.$router.push({
            path: "/center/inReviewPapers",
          });
        } else {
          this.$message({
            message: "Submit failed",
            type: "error",
          });
        }
      });
    },
    // 限制添加文件的格式为PDF和大小小于15MB
    beforeUpload(file) {
      const isPDF = file.raw.type === "application/pdf";
      const isLt15M = file.raw.size / 1024 / 1024 < 15;

      if (!isPDF) {
        this.$message.error("Upload file must be PDF format!");
        // 清除添加文件
        this.fileList = [];
      }
      if (!isLt15M) {
        this.$message.error("Upload file size can not exceed 15MB!");
        // 清除添加文件
        this.fileList = [];
      }

      return isPDF && isLt15M;
    },
    async handleUploadChange(file, fileList) {
      if (this.beforeUpload(file) === false) {
        return;
      }
      this.fileList = fileList;
      try {
        // 计算文件的哈希值
        const fileData = await this.readFileAsArrayBuffer(this.fileList[0].raw);
        const hashBuffer = await crypto.subtle.digest("SHA-256", fileData);
        const hashArray = Array.from(new Uint8Array(hashBuffer));
        const fileHash = hashArray
          .map((byte) => byte.toString(16).padStart(2, "0"))
          .join("");
        console.log("fileHash:", fileHash);
        // 调用合约函数
        await this.callSmartContract(fileHash);
      } catch (error) {
        console.error("Error:", error);
      }
    },
    handleClose(tag) {
      console.log(this.form.authorTags.indexOf(tag), "this.form.authorTags");
      console.log(this.form.authorTags, "this.form.authorTags");

      this.form.authorTags.splice(this.form.authorTags.indexOf(tag), 1);
    },
    handleClose2(tag) {
      this.form.keywordTags.splice(this.form.keywordTags.indexOf(tag), 1);
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
        this.form.authorTags.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = "";
    },
    handleInputConfirm2() {
      let inputValue2 = this.inputValue2;
      if (inputValue2) {
        this.form.keywordTags.push(inputValue2);
      }
      this.inputVisible2 = false;
      this.inputValue2 = "";
    },
  },
  mounted() {
    // 获取conference中的category
    getConferenceDetail({
      conference_id: this.$route.query.id,
    }).then((res) => {
      this.form.options = [];
      res.data.data.category.split(",").forEach((item) => {
        this.form.options.push({ label: item, value: item });
      });
      // 从localStorage中获取用户信息
      this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
      // 把this.userInfo.block_chain_address中的大小字母转换成小写
      this.userInfo.block_chain_address =
        this.userInfo.block_chain_address.toLowerCase();
      console.log(this.form.options, "options");
    });
  },
};
</script>
<style lang="less" scoped>
.box {
  // width: 100%;
  // height: 3000px;
  // 盒子里面的内容水平居中
  // 宽高被盒子撑开
  display: inline-block;
  text-align: center;
  .image {
    width: 80%;
    height: 430px;
  }
  .box1 {
    // 取消内容水平居中
    text-align: left;
    width: 80%;
    height: 1000px;
    // background-color: #e2f1fb;
    background-color: #ecf5ff;
    // 圆角
    border-radius: 10px;
    // 水平居中
    margin: 0 auto;
    margin-top: 20px;
    // 上下左右panding
    padding: 20px 20px;
    color: #072e5b;
    .upload-demo {
      margin-left: 10px;
      margin-bottom: 10px;
    }
  }
}
.el-tag + .el-tag {
  margin-left: 10px;
}
.button-new-tag {
  margin-left: 10px;
  height: 32px;
  line-height: 30px;
  padding-top: 0;
  padding-bottom: 0;
}
.input-new-tag {
  width: 90px;
  margin-left: 10px;
  vertical-align: bottom;
}
</style>
