<template>
  <div class="user-profile-container">
    <div class="profile-header">
      <div class="avatar-section">
        <div class="avatar-wrapper">
          <img
            v-if="userForm && userForm.headerImg"
            :src="userForm.headerImg"
            class="avatar"
            alt="User Avatar"
          />
          <i v-else class="el-icon-user-solid avatar-placeholder"></i>
        </div>
        <el-button 
          type="primary" 
          size="small" 
          @click="showImageDialog = true"
          class="change-avatar-btn"
        >
          Change Avatar
        </el-button>
      </div>
      <div class="user-name">
        <h2>{{ userForm?.first_name || '' }} {{ userForm?.middle_name || '' }} {{ userForm?.last_name || '' }}</h2>
        <div v-if="userForm?.block_chain_address" class="blockchain-address-section" @click="copyBlockchainAddress">
          <span class="blockchain-address">{{ formatBlockchainAddress(userForm.block_chain_address) }}</span>
          <i class="el-icon-document-copy copy-icon"></i>
        </div>
        <p class="username">@{{ userForm?.username || '' }}</p>
      </div>
    </div>

    <el-tabs v-model="activeTab" class="profile-tabs">
      <el-tab-pane label="Information" name="basic">
        <div class="info-section">
          <h3 class="section-title">Account Information</h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <div class="info-item">
                <span class="info-label">Username</span>
                <span class="info-value">{{ userForm?.username || 'Not Set' }}</span>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="info-item">
                <span class="info-label">Email</span>
                <span class="info-value">{{ userForm?.email || 'Not Set' }}</span>
              </div>
            </el-col>
          </el-row>
        </div>

        <div class="info-section">
          <h3 class="section-title">Personal Information</h3>
          <el-row :gutter="20">
            <el-col :span="8">
              <div class="info-item">
                <span class="info-label">First Name</span>
                <span class="info-value">{{ userForm?.first_name || 'Not Set' }}</span>
              </div>
            </el-col>
            <el-col :span="8">
              <div class="info-item">
                <span class="info-label">Middle Name</span>
                <span class="info-value">{{ userForm?.middle_name || 'Not Set' }}</span>
              </div>
            </el-col>
            <el-col :span="8">
              <div class="info-item">
                <span class="info-label">Last Name</span>
                <span class="info-value">{{ userForm?.last_name || 'Not Set' }}</span>
              </div>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <div class="info-item">
                <span class="info-label">Phone</span>
                <span class="info-value">{{ userForm?.phone || 'Not Set' }}</span>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="info-item">
                <span class="info-label">Address</span>
                <span class="info-value">{{ userForm?.address || 'Not Set' }}</span>
              </div>
            </el-col>
          </el-row>
        </div>

        <div class="info-section">
          <h3 class="section-title">Professional Information</h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <div class="info-item">
                <span class="info-label">Education</span>
                <span class="info-value">{{ userForm?.education || 'Not Set' }}</span>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="info-item">
                <span class="info-label">Position/Title</span>
                <span class="info-value">{{ userForm?.title || 'Not Set' }}</span>
              </div>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="24">
              <div class="info-item">
                <span class="info-label">Research Fields</span>
                <span class="info-value">{{ userForm?.research || 'Not Set' }}</span>
              </div>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <div class="info-item">
                <span class="info-label">Organization</span>
                <span class="info-value">{{ userForm?.affiliation || 'Not Set' }}</span>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="info-item">
                <span class="info-label">Department</span>
                <span class="info-value">{{ userForm?.affiliation_type || 'Not Set' }}</span>
              </div>
            </el-col>
          </el-row>
        </div>
      </el-tab-pane>
    </el-tabs>

    <div class="action-buttons">
      <el-button type="primary" @click="update">Edit Information</el-button>
    </div>

    <!-- Avatar Selection Dialog -->
    <el-dialog 
      :visible.sync="showImageDialog" 
      title="Select Avatar" 
      width="600px"
      custom-class="avatar-dialog"
    >
      <div class="image-container">
        <div 
          v-for="(image, index) in existingImages" 
          :key="index" 
          class="image-item"
          @click="selectExistingImage(image)"
        >
          <img :src="image" alt="Avatar Option" />
        </div>
      </div>
      
      <div class="upload-section">
        <el-upload
          class="avatar-uploader"
          :on-change="handleUploadChange"
          action=""
          :auto-upload="false"
          :show-file-list="false"
          :before-upload="beforeAvatarUpload"
        >
          <el-button type="primary" plain>Upload Local Image</el-button>
        </el-upload>
      </div>
      
      <span slot="footer" class="dialog-footer">
        <el-button @click="showImageDialog = false">Cancel</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { getSelfInfo, changeHeaderImg } from "../../api";

export default {
  data() {
    return {
      activeTab: 'basic',
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
      userForm: null,
    };
  },
  methods: {
    async selectExistingImage(image) {
      try {
        const blob = await (await fetch(image)).blob();
        const file = new File([blob], 'selected_image.jpg', { type: blob.type });

        const formData = new FormData();
        formData.append("file_name", file.name);
        formData.append("data", file);
        
        const res = await changeHeaderImg(formData);
        if (res.data.code === 1000) {
          this.userForm.headerImg = res.data.data;
          this.$message.success("Avatar changed successfully");
          this.showImageDialog = false;
          location.reload();
        } else {
          this.$message.error("Failed to change avatar");
        }
      } catch (error) {
        console.error("Error changing avatar:", error);
        this.$message.error("Failed to change avatar");
      }
    },
    update() {
      this.$router.push("/center/updateInformation");
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isPNG = file.type === "image/png";
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG && !isPNG) {
        this.$message.error("Only JPG or PNG format is allowed!");
        return false;
      }
      if (!isLt2M) {
        this.$message.error("Image size cannot exceed 2MB!");
        return false;
      }
      return true;
    },
    async handleUploadChange(file, fileList) {
      if (fileList.length === 0) return;
      
      try {
        const formData = new FormData();
        formData.append("file_name", fileList[0].name);
        formData.append("data", fileList[0].raw);
        
        const res = await changeHeaderImg(formData);
        if (res.data.code === 1000) {
          this.userForm.headerImg = res.data.data;
          this.$message.success("Avatar changed successfully");
          this.showImageDialog = false;
          location.reload();
        } else {
          this.$message.error("Failed to change avatar");
        }
      } catch (error) {
        console.error("Error changing avatar:", error);
        this.$message.error("Failed to change avatar");
      }
    },
    copyBlockchainAddress() {
      if (this.userForm?.block_chain_address) {
        navigator.clipboard.writeText(this.userForm.block_chain_address)
          .then(() => {
            this.$message.success('Blockchain address copied to clipboard');
          })
          .catch(() => {
            this.$message.error('Failed to copy blockchain address');
          });
      }
    },
    formatBlockchainAddress(address) {
      if (!address) return '';
      if (address.length <= 12) return address;
      return `${address.slice(0, 6)}...${address.slice(-6)}`;
    },
  },
  mounted() {
    getSelfInfo().then((res) => {
      if (res && res.code === 1000 && res.data) {
        const userData = res.data;
        this.userForm = {
          id: userData.ID || 0,
          username: userData.username || '',
          first_name: userData.first_name || '',
          middle_name: userData.middle_name || '',
          last_name: userData.last_name || '',
          email: userData.email || '',
          phone: userData.phone || '',
          address: userData.address || '',
          education: userData.education || '',
          title: userData.title || '',
          research: userData.research || '',
          block_chain_address: userData.block_chain_address || '',
          affiliation: userData.affiliation || '',
          affiliation_type: userData.affiliation_type || '',
          headerImg: userData.headerImg || '',
        };
      } else {
        console.error('Failed to get user information:', res);
        this.$message.error('Failed to get user information');
        this.userForm = {
          id: 0,
          username: '',
          first_name: '',
          middle_name: '',
          last_name: '',
          email: '',
          phone: '',
          address: '',
          education: '',
          title: '',
          research: '',
          block_chain_address: '',
          affiliation: '',
          affiliation_type: '',
          headerImg: '',
        };
      }
    }).catch(error => {
      console.error('Error getting user information:', error);
      this.$message.error('Failed to get user information');
      this.userForm = {
        id: 0,
        username: '',
        first_name: '',
        middle_name: '',
        last_name: '',
        email: '',
        phone: '',
        address: '',
        education: '',
        title: '',
        research: '',
        block_chain_address: '',
        affiliation: '',
        affiliation_type: '',
        headerImg: '',
      };
    });
  },
};
</script>

<style lang="less" scoped>
.user-profile-container {
  max-width: 1000px;
  margin: 30px auto;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  padding: 30px;
}

.profile-header {
  display: flex;
  align-items: flex-start;
  margin-bottom: 30px;
  padding-bottom: 20px;
  border-bottom: 1px solid #ebeef5;
}

.avatar-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-right: 30px;
}

.avatar-wrapper {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  overflow: hidden;
  margin-bottom: 10px;
  border: 2px solid #409EFF;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.avatar {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f5f7fa;
  color: #909399;
  font-size: 50px;
}

.change-avatar-btn {
  margin-top: 10px;
}

.user-name {
  margin-left: 30px;
  text-align: left;
  
  h2 {
    margin: 0 0 15px 0;
    font-size: 24px;
    color: #303133;
  }
  
  .username {
    margin: 10px 0 0 0;
    color: #909399;
    font-size: 14px;
  }

  .blockchain-address-section {
    margin: 15px 0;
    display: inline-flex;
    align-items: center;
    gap: 6px;
    cursor: pointer;
    max-width: 200px;
    transition: all 0.3s;
    font-size: 13px;
    
    .blockchain-address {
      color: #1a1a1a;
      font-family: monospace;
      background-color: #f5f7fa;
      padding: 2px 10px;
      border-radius: 12px;
      box-shadow: 0 1px 2px rgba(0, 0, 0, 0.08);
      
      &:hover {
        background-color: #ecf5ff;
      }
    }
    
    .copy-icon {
      opacity: 0;
      margin-left: 4px;
      transition: all 0.3s;
      color: #409EFF;
      font-size: 14px;
    }
    
    &:hover {
      .copy-icon {
        opacity: 1;
      }
    }
  }
}

.profile-tabs {
  margin-bottom: 30px;
}

.info-section {
  margin-bottom: 30px;
  padding: 20px;
  background-color: #f8f9fa;
  border-radius: 8px;
  
  .section-title {
    margin: 0 0 20px 0;
    padding-bottom: 10px;
    border-bottom: 1px solid #ebeef5;
    color: #409EFF;
    font-size: 18px;
    font-weight: 500;
  }
}

.info-item {
  margin-bottom: 20px;
  
  .info-label {
    display: block;
    color: #606266;
    font-size: 14px;
    margin-bottom: 5px;
  }
  
  .info-value {
    display: block;
    color: #303133;
    font-size: 16px;
    font-weight: 500;
  }
  
  .blockchain-address {
    word-break: break-all;
    font-family: monospace;
  }
}

.action-buttons {
  text-align: center;
  margin-top: 30px;
  
  .el-button {
    padding: 12px 35px;
  }
}

.image-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  max-height: 300px;
  overflow-y: auto;
  padding: 10px;
  background-color: #f8f9fa;
  border-radius: 4px;
}

.image-item {
  width: 80px;
  height: 80px;
  border-radius: 4px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.3s;
  border: 2px solid transparent;
  
  &:hover {
    transform: scale(1.05);
    border-color: #409EFF;
  }
  
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}

.upload-section {
  margin-top: 20px;
  text-align: center;
}

.avatar-dialog {
  .el-dialog__body {
    padding: 20px;
  }
}
</style>