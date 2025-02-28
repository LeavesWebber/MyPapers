<template>
  <el-row>
    <el-col :span="24">
      <el-col :span="4">
      <!-- 图片选择对话框 -->
      <el-dialog :visible.sync="showImageDialog" title="选择图片">
        
          <!-- 已有图片列表 -->
          <div class="image-container">
            <div v-for="(image, index) in existingImages" :key="index">
              <img :src="image" alt="已有图片" style="width: 100px; height: 100px; margin: 10px; cursor: pointer" @click="selectExistingImage(image)">
            </div>
          </div>
          <!-- 选择本地图片按钮 -->
          <el-upload class="avatar-uploader"
          :on-change="handleUploadChange"
          action=""
          :auto-upload="false"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
          >选择本地图片</el-upload>
        
        <template #footer>
          <el-button @click="showImageDialog = false">取消</el-button>
        </template>
      </el-dialog>
      <!-- 隐藏的文件输入框 -->
      
      <el-button
        @click="showImageDialog=true"
      >
        <img
          v-if="userForm.headerImg"
          :src="userForm.headerImg"
          class="avatar"
        />
        <i v-else class="el-icon-plus avatar-uploader-icon"></i>
      </el-button>
    </el-col>
      <el-col :span="12"> UserName: {{ userForm.username }} </el-col>

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
    <el-col :span="12"> Affiliation Type: {{ userForm.affiliation_type }}</el-col>
    
      
  </el-row>
</template>

<script>
import { getSelfInfo, changeHeaderImg } from "../../api";

export default {
  data() {
    return {
      showImageDialog: false, 
      existingImages: [
        "/headimg/成熟男用户头像1.png",
        "/headimg/成熟男用户头像2.png",
        "/headimg/成熟男用户头像3.png",
        "/headimg/成熟男用户头像4.png",
        "/headimg/成熟男用户头像-欧美1.png",
        "/headimg/成熟男用户头像-欧美2.png",
        "/headimg/成熟男用户头像-欧美3.png",
        "/headimg/成熟男用户头像-欧美4.png",
        "/headimg/成熟女用户头像1.png",
        "/headimg/成熟女用户头像2.png",
        "/headimg/成熟女用户头像3.png",
        "/headimg/成熟女用户头像4.png",
        "/headimg/成熟女用户头像5.png",
        "/headimg/成熟女用户头像6.png",
        "/headimg/成熟女用户头像7.png",
        "/headimg/成熟女用户头像8.png",
        "/headimg/成熟女用户头像-欧美1.png",
        "/headimg/成熟女用户头像-欧美2.png",
        "/headimg/成熟女用户头像-欧美3.png",
        "/headimg/成熟女用户头像-欧美4.png",
        "/headimg/年青男用户头像-欧美1.png",
        "/headimg/年青男用户头像-欧美2.png",
        "/headimg/年青男用户头像-欧美3.png",
        "/headimg/年青男用户头像-欧美4.png",
        "/headimg/年青男用户头像1.png",
        "/headimg/年青男用户头像2.png",
        "/headimg/年青男用户头像3.png",
        "/headimg/年青男用户头像4.png",
        "/headimg/年青男用户头像5.png",
        "/headimg/年青女用户头像-欧美1.png",
        "/headimg/年青女用户头像-欧美2.png",
        "/headimg/年青女用户头像-欧美3.png",
        "/headimg/年青女用户头像-欧美4.png",
        "/headimg/年青女用户头像1.png",
        "/headimg/年青女用户头像2.png",
        "/headimg/年青女用户头像3.png",
        "/headimg/年青女用户头像4.png",
        "/headimg/年青女用户头像5.png",
        "/headimg/年青女用户头像6.png",
        "/headimg/年青女用户头像7.png",
        "/headimg/年青女用户头像8.png",
        "/headimg/年青女用户头像9.png",
      ],
      selectedFile: null,
      headers: {},
      uploadUrl: 'http://localhost:8887/mypapers/user/changeHeaderImg',
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
    async selectExistingImage(image) {
      const blob = await (await fetch(image)).blob();
      const file = new File([blob], 'selected_image.jpg', { type: blob.type });

      const formData = new FormData();
      
      formData.append("file_name", file.name);
      formData.append("data", file);
      await changeHeaderImg(formData).then((res) => {
        console.log(res);
        if (res.data.code === 1000) {
          this.userForm.headerImg = res.data.data;
          // 清空fileList
          this.$message.success("Change Success");
          location.reload()
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
          location.reload();
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
.image-container {
  display: flex;
  flex-wrap: wrap;
}
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
.avatar-uploader .el-button {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-button:hover {
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