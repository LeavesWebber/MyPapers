<template>
  <el-row>
    <el-col :span="24">
      <el-col :span="4">
        <el-upload
          class="avatar-uploader"
          :on-change="handleUploadChange"
          action=""
          :auto-upload="false"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
        >
          <img
            v-if="userForm.headerImg"
            :src="userForm.headerImg"
            class="avatar"
          />
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
      </el-col>
      <el-col :span="12"> Username: {{ userForm.username }} </el-col>
      <el-col :span="12"> First Name: {{ userForm.first_name }} </el-col>
      <el-col :span="12"> Last Name: {{ userForm.last_name }} </el-col>
      <el-col :span="19">
        Blockchain Address: {{ userForm.block_chain_address }}
      </el-col>
      <el-button type="primary" plain @click="update">Update</el-button>
    </el-col>

    <el-col :span="12"> Email: {{ userForm.email }} </el-col>
    <el-col :span="12"> Department: {{ userForm.department }} </el-col>
    <el-col :span="12"> Phone: {{ userForm.phone }} </el-col>
    <el-col :span="12"> Address: {{ userForm.address }} </el-col>
    <el-col :span="12"> Education: {{ userForm.education }} </el-col>
    <el-col :span="12"> Title: {{ userForm.title }} </el-col>
    <el-col :span="12"> Research: {{ userForm.research }} </el-col>

    <el-col :span="12"> Affiliation: {{ userForm.affiliation }} </el-col>
    <el-col :span="12">
      Affiliation Type: {{ userForm.affiliation_type }}
    </el-col>
  </el-row>
</template>
<script>
import { getSelfInfo, changeHeaderImg } from "../../api";
export default {
  data() {
    return {
      userForm: {
        id: 0,
        username: "",
        first_name: "",
        last_name: "",
        sex: "",
        email: "",
        departmentd: "",
        phone: "",
        address: "",
        education: "",
        title: "",
        research: "",
        block_chain_address: "",
        affiliations: "",
        affiliation_type: "",
        headerImg: "",
      },
    };
  },
  methods: {
    update() {
      this.$router.push("/center/updateInformation");
    },
    handleAvatarSuccess(res, file) {
      this.headerImg = URL.createObjectURL(file.raw);
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isPDF = file.type === "application/pdf";
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error("Only JPG Or PDF Format Allowed!");
      }
      if (!isPDF) {
        this.$message.error("Only JPG Or PDF Format Allowed!");
      }
      if (!isLt2M) {
        this.$message.error("Image Size Can't Exceed 2MB!");
      }
      return isJPG && isLt2M;
    },
    async handleUploadChange(file, fileList) {
      // 调用修改头像接口
      const formData = new FormData();
      formData.append("file_name", fileList[0].name);
      formData.append("data", fileList[0].raw);
      await changeHeaderImg(formData).then((res) => {
        console.log(res);
        if (res.data.code === 1000) {
          this.userForm.headerImg = res.data.data;
          // 清空fileList
          fileList.splice(0, 1);
          this.$message.success("Change Success");
        } else {
          // 清空fileList
          fileList.splice(0, 1);
          this.$message({
            message: "Change failed",
            type: "error",
          });
        }
      });
    },
  },
  mounted() {
    getSelfInfo().then((res) => {
      this.userForm = res.data.data;
      console.log(res.data);
    });
  },
};
</script>
<style  lang="less" scoped>
.el-row {
  padding: 50px;
  margin-bottom: 20px;
  background: #ecf5ff;
  height: 550px;
  border-radius: 8px;
  &:last-child {
    margin-bottom: 0;
  }
}
.el-col {
  // 内容靠左
  text-align: left;
  border-radius: 4px;
}
.box {
  // 内边距
  padding: 50px;
  background: #ecf5ff;
  border-radius: 8px;
  height: 570px;
}
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  background-color: rgb(226, 237, 247);
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 132px;
  height: 178px;
  display: block;
}
.avatar-box {
  width: 132px;
}
</style>