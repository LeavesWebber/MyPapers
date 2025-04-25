<template>
  <div class="header-container">
    <div class="header-left">
      <el-button
        class="toggle-menu"
        @click="handleMenu"
        icon="el-icon-s-fold"
        type="text"
      ></el-button>
      <div class="brand-title">Mypapers</div>
    </div>
    
    <div class="header-right">
      <div class="nav-item">
        <el-button class="action-button" type="text" @click="home">
          <i class="el-icon-s-home"></i>
        </el-button>
      </div>
      
      <el-dropdown @command="handleCommand" v-if="hasAuthority()" class="nav-item">
        <span class="el-dropdown-link">
          <i class="el-icon-s-tools"></i>
          <span class="dropdown-text">Permissions</span>
        </span>
        <el-dropdown-menu slot="dropdown" class="auth-dropdown">
          <el-dropdown-item
            v-for="authority in userInfo.authorities"
            :key="authority.authorityId"
            :command="authority.authorityId"
            :class="{ 'active-authority': authority.authorityId === userInfo.authorityId }"
          >
            <span>{{ authority.authorityName }}</span>
            <i
              v-if="authority.authorityId === userInfo.authorityId"
              class="el-icon-check"
            ></i>
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
      
      <el-dropdown class="nav-item user-dropdown" trigger="click">
        <div class="user-info">
          <el-avatar :size="32" class="user-avatar">
            {{ userInfo.username.charAt(0).toUpperCase() }}
          </el-avatar>
          <span class="username">{{ userInfo.username }}</span>
          <i class="el-icon-arrow-down"></i>
        </div>
        <el-dropdown-menu slot="dropdown" class="user-dropdown-menu">
          <el-dropdown-item @click.native="goToProfile">
            <i class="el-icon-user"></i> Profile
          </el-dropdown-item>
          <el-dropdown-item divided @click.native="logout">
            <i class="el-icon-switch-button"></i> Logout
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
import { changeAuthority, getMenu } from "../api";
export default {
  data() {
    return {
      userInfo: {
        username: "",
        authorities: [],
        authorityId: 0,
      },
    };
  },
  methods: {
    handleCommand(command) {
      changeAuthority({ authority_id: command }).then((res) => {
        if (res.data.code === 1000) {
          // 改一下userInfo中的authorityId
          this.userInfo.authorityId = command;
          // 更新localStorage中的userInfo
          localStorage.setItem("userInfo", JSON.stringify(this.userInfo));
          // 重新获取菜单
          getMenu({ authorityId: this.userInfo.authorityId }).then(
            ({ data }) => {
              // 更新localStorage中的menu
              // localStorage.setItem("menu", JSON.stringify(data.data));
              this.$store.commit("setMenu", data.data);
              this.$store.commit("addMenu", this.$router);
              // 刷新页面
              window.location.reload();
            }
          );
          this.$message({
            message: "权限切换成功",
            type: "success",
          });
        } else {
          this.$message({
            message: "权限切换失败",
            type: "error",
          });
        }
      });
    },
    logout() {
      localStorage.removeItem("token");
      localStorage.removeItem("menu");
      // 跳转到主页面
      this.$router
        .push({
          path: "/home",
        })
        .catch((err) => {});
    },
    home() {
      // 跳转到主页面
      this.$router
        .push({
          path: "/home",
        })
        .catch((err) => {});
    },
    handleMenu() {
      this.$store.commit("collapseMenu");
    },
    hasAuthority() {
      // return this.userInfo.authorities.some(
      //   (authority) => authority.authorityId === 101
      // );
      // return this.userInfo.authorityId === 101;
      return false;
    },
    goToProfile() {
      this.$router.push('/center/information');
    },
  },
  mounted() {
    // 从localStorage中获取用户信息
    this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
  },
};
</script>
<style lang="less" scoped>
.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
  background-color: #ffffff;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.08);
  padding: 0 20px;
  color: #333;
  position: relative;
  z-index: 100;
}

.header-left {
  display: flex;
  align-items: center;
  
  .toggle-menu {
    color: #606266;
    font-size: 18px;
    margin-right: 16px;
    padding: 0;
    
    &:hover, &:focus {
      color: #409EFF;
      background: transparent;
    }
  }
  
  .brand-title {
    font-size: 16px;
    font-weight: 500;
    color: #303133;
    margin-left: 4px;
    letter-spacing: 0.5px;
  }
}

.header-right {
  display: flex;
  align-items: center;
}

.nav-item {
  margin-left: 8px;
  position: relative;
  padding: 0 8px;
  
  &:not(:last-child)::after {
    content: '';
    position: absolute;
    right: -4px;
    top: 50%;
    transform: translateY(-50%);
    width: 1px;
    height: 16px;
    background-color: #EBEEF5;
  }
}

.action-button {
  color: #606266;
  font-size: 18px;
  
  &:hover, &:focus {
    color: #409EFF;
    background: transparent;
  }
}

.el-dropdown-link {
  cursor: pointer;
  color: #606266;
  display: flex;
  align-items: center;
  font-size: 14px;
  
  i {
    font-size: 16px;
    margin-right: 4px;
  }
  
  .dropdown-text {
    margin-right: 4px;
  }
}

.user-dropdown {
  cursor: pointer;
  margin-left: 20px;
  
  .user-info {
    display: flex;
    align-items: center;
    padding: 0 8px;
    
    &:hover {
      background-color: #F2F6FC;
      border-radius: 4px;
    }
  }
  
  .user-avatar {
    background-color: #409EFF;
    color: #fff;
    border: 1px solid #E4E7ED;
  }
  
  .username {
    margin: 0 8px;
    font-size: 14px;
    max-width: 120px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    color: #606266;
  }
}

.auth-dropdown {
  min-width: 120px;
}

.active-authority {
  color: #409EFF;
  font-weight: 500;
  background-color: #ecf5ff;
}

/deep/ .el-dropdown-menu__item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 16px;
  
  i {
    margin-left: 8px;
    font-size: 14px;
  }
  
  &:hover {
    background-color: #F5F7FA;
  }
}

.user-dropdown-menu {
  min-width: 160px;
  
  /deep/ .el-dropdown-menu__item {
    i {
      margin-right: 8px;
      margin-left: 0;
      font-size: 14px;
    }
  }
}
</style>