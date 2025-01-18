<template>
  <div class="header-container">
    <div class="l-content">
      <el-button
        style="margin-right: 20px; color: #40a2fe"
        @click="handleMenu"
        icon="el-icon-menu"
        size="mini"
      ></el-button>
    </div>
    <el-button class="logout" type="text" @click="logout">Logout</el-button>
    <el-button class="home" type="text" @click="home">Home</el-button>
    <span class="username">
      {{ userInfo.username }}
    </span>
    <el-dropdown @command="handleCommand" v-if="hasAuthority()">
      <span class="el-dropdown-link">
        Authority <i class="el-icon-arrow-down el-icon--right"></i>
      </span>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item
          v-for="authority in userInfo.authorities"
          :key="authority.authorityId"
          :command="authority.authorityId"
        >
          <span>{{ authority.authorityName }}</span>
          <i
            v-if="authority.authorityId === userInfo.authorityId"
            class="el-icon-check"
          ></i>
        </el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
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
            message: "Create successfully",
            type: "success",
          });
        } else {
          this.$message({
            message: "Change Authority failed",
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
  },
  mounted() {
    // 从localStorage中获取用户信息
    this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
  },
};
</script>
<style lang="less" scoped>
.header-container {
  background-color: #fff;
  border-bottom: 1px solid #f3f4f6;
  height: 40px;
  padding: 0 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
  .text {
    color: #fff;
    font-size: 14px;
    margin-left: 10px;
  }
  .r-content {
    .user {
      width: 40px;
      height: 40px;
      border-radius: 50%;
    }
  }
  .l-content {
    display: flex;
    align-items: center;
    /deep/.el-breadcrumb__item {
      .el-breadcrumb__inner {
        font-weight: normal;
        &.is-link {
          color: #666;
        }
      }
      &:last-child {
        .el-breadcrumb__inner {
          color: #fff;
        }
      }
    }
  }
  .logout {
    position: absolute;
    right: 180px;
    color: #40a2fe;
    // 鼠标悬停时显示下划线
    &:hover {
      text-decoration: underline;
    }
  }
  .home {
    position: absolute;
    right: 240px;
    color: #40a2fe;
    // 鼠标悬停时显示下划线
    &:hover {
      text-decoration: underline;
    }
  }
  .username {
    position: absolute;
    right: 120px;
    color: #40a2fe;
  }
}
.el-dropdown-link {
  cursor: pointer;
  color: #409eff;
}
.el-icon-arrow-down {
  font-size: 12px;
}
</style>