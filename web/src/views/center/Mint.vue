<template>
  <div class="box">
    <el-form ref="form" :model="form" label-width="200px">
      <el-form-item label="Upload your paper here:">
        <el-upload
          class="upload-demo"
          ref="upload"
          :on-change="handleUploadChange"
          action=""
          :file-list="fileList"
          :auto-upload="false"
          :limit="1"
        >
          <el-button slot="trigger" size="small" type="primary"
            >Upload</el-button
          >
        </el-upload>
      </el-form-item>
      <el-form-item label="Title">
        <el-input v-model="form.title"></el-input>
      </el-form-item>
      <el-form-item label="Paper Type">
        <el-input v-model="form.paper_type" style="width: 360px"></el-input>
      </el-form-item>
      <el-form-item label="Abstract">
        <el-input
          style="width: 360px"
          type="textarea"
          v-model="form.abstract"
        ></el-input>
      </el-form-item>
      <el-form-item label="Block Address">
        <el-input style="width: 700px" v-model="form.block_address"></el-input>
      </el-form-item>
      <el-form-item label="Authors">
        <el-tag
          :key="tag"
          v-for="tag in authorTags"
          closable
          :disable-transitions="false"
          @close="handleClose(tag)"
        >
          {{ tag }}
        </el-tag>
        <el-input
          style="width: 160px"
          class="input-new-tag"
          v-if="inputVisible"
          v-model="inputValue"
          ref="saveTagInput"
          size="small"
          @keyup.enter.native="handleInputConfirm"
          @blur="handleInputConfirm"
        >
        </el-input>
        <el-button v-else class="button-new-tag" size="small" @click="showInput"
          >+ Author Name</el-button
        >
      </el-form-item>
      <el-form-item label="Corresponding Author">
        <el-input style="width: 360px" v-model="form.cor_author"></el-input>
      </el-form-item>

      <el-form-item label="Key Words">
        <el-input style="width: 360px" v-model="form.key_words"></el-input>
      </el-form-item>
      <el-form-item label="Subject Category">
        <el-input
          style="width: 360px"
          v-model="form.subject_category"
        ></el-input>
      </el-form-item>
    </el-form>
    <el-form>
      <el-form-item>
        <el-button type="primary" @click="revise">Revise</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getDetailPapers, submitPaper } from "../../api";
import { ERC20contractInstance } from "@/constant";
const contractInstance = ERC20contractInstance;
export default {
  data() {
    return {
      Paper_id: this.$route.query.Paper_id,
      form: {
        ID: 0,
        version_id: "",
        authors: "",
        CreatedAt: "",
        UpdatedAt: "",
        conference_id: 0,
        journal_id: 0,
        title: "",
        paper_type: "",
        abstract: "",
        key_words: "",
        subject_category: "",
        manuscript_id: "",
        informed_consent: "",
        animal_subjects: "",
        cor_author: "",
        manuscript_type: "",
        unique_contribution: "",
        hash: "",
        block_address: "",
        filepath: "",
        cid: "",
        status: "",
        user: [
          {
            ID: 0,
            CreatedAt: "",
            UpdatedAt: "",
            uuid: 0,
            authorityId: 0,
            sex: 0,
            username: "",
            first_name: "",
            last_name: "",
            headerImg: "",
            email: "",
            phone: "",
            address: "",
            education: "",
            title: "",
            research: "",
            block_chain_address: "",
            affiliation: "",
            affiliation_type: "",
          },
        ],
      },
      fileList: [],
      authorTags: [],
      inputVisible: false,
      inputValue: "",
    };
  },
  methods: {
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

        const result = await contractInstance.methods[functionName](
          ...functionArgs
        ).send({
          from: window.ethereum.selectedAddress,
          gasPrice: "0",
        });

        // 输出结果
        console.log("Transaction result:", result);
      } catch (error) {
        // 处理错误
        console.error("Error:", error);
      }
    },
    async revise() {
      console.log(this.form, "revise this.form");
      const formData = new FormData();
      if (this.form.journal_id !== 0) {
        formData.append("journal_id", this.form.journal_id);
      } else {
        formData.append("conference_id", this.form.conference_id);
      }
      formData.append("id", this.$route.query.paper_id);
      formData.append("version_id", this.form.version_id);
      formData.append("paper_type", this.form.paper_type);
      formData.append("title", this.form.title);
      formData.append("abstract", this.form.abstract);
      formData.append("authors", this.authorTags);
      formData.append("keywords", this.form.key_words);
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
      formData.append("data", this.fileList[0].raw);
      console.log(formData, "revise formData");
      submitPaper(formData).then((res) => {
        if (res.data.code === 1000) {
          // localStorage.setItem("token", data.data.token); // 用localStorage缓存token值
          this.$message({
            message: "Revise Success",
            type: "success",
          });
          this.$router.push({
            path: "/center/detailPaper",
            query: {
              paper_id: res.data.data.ID,
            },
          });
        } else {
          this.$message({
            message: "Revise Failed",
            type: "error",
          });
        }
      });
    },
    async handleUploadChange(file, fileList) {
      this.fileList = fileList;
      try {
        // 计算文件的哈希值
        const fileData = await this.readFileAsArrayBuffer(this.fileList[0].raw);
        const hashBuffer = await crypto.subtle.digest("SHA-256", fileData);
        const hashArray = Array.from(new Uint8Array(hashBuffer));
        const fileHash = hashArray
          .map((byte) => byte.toString(16).padStart(2, "0"))
          .join("");
        // 调用合约函数
        await this.callSmartContract(fileHash);
      } catch (error) {
        console.error("Error:", error);
      }
    },
    handleClose(tag) {
      this.authorTags.splice(this.authorTags.indexOf(tag), 1);
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
        this.authorTags.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = "";
    },
  },
  mounted() {
    getDetailPapers({ paper_id: this.$route.query.paper_id }).then((res) => {
      console.log(res.data.data);
      this.form = res.data.data.paper;
      // form里面的user的username加到authorTags
      this.form.user.forEach((item) => {
        this.authorTags.push(item.username);
      });
      console.log(this.form, "this.form");
      console.log(this.authorTags), "this.authorTags";
    });
  },
};
</script>
<style lang="less" scoped>
.box {
  // 内边距
  padding: 40px;
  // 取消父组件的居中
  text-align: left;
}
</style>