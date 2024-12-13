<template>
  <el-menu
    default-active="1-4-1"
    class="el-menu-vertical-demo"
    @open="handleOpen"
    @close="handleClose"
    :collapse="isCollapse"
     background-color="#191A23"
      text-color="#fff"
      active-text-color="#fff"
  >
  <h4>{{ isCollapse ? 'Center': 'Personal Center' }}</h4>
    <el-menu-item @click="clickMenu(item)" v-for="item in noChildren" :key="item.title" :index="item.title">
      <i :class="`el-icon-${item.icon}`" style="color:#fff"></i>
      <span slot="title">{{item.name}}</span>
    </el-menu-item>
    <el-submenu v-for="item in hasChildren" :key="item.title" :index="item.title">
      <template slot="title">
        <i :class="`el-icon-${item.icon}`" style="color:#fff"></i>
        <span slot="title">{{item.label}}</span>
      </template>
      <el-menu-item-group v-for="subItem in item.children" :key="subItem.path">
        <el-menu-item  @click="clickMenu(subItem)" :index="subItem">{{subItem.label}}</el-menu-item>
      </el-menu-item-group>
      </el-submenu>
  </el-menu>
</template>

<style lang="less" scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 25%;
  min-width: 280px;
  max-width: 400px;
  min-height: 100vh;
}

.el-menu {
  height: 100vh;
  border-right: none;
  
  h4 {
    color: #fff;
    text-align: center;
    height: 8vh;
    line-height: 8vh;
    font-size: clamp(14px, 2vh, 18px);
    font-weight: 400;
  }

  :deep(.el-submenu__title),
  :deep(.el-menu-item) {
    height: auto;
    min-height: 7vh;
    line-height: 1.5;
    padding: 1vh 1.5vw;
    white-space: normal;
    word-wrap: break-word;
    display: flex;
    align-items: center;
    font-size: clamp(13px, 1.8vh, 16px);
  }

  :deep(.el-submenu .el-menu-item) {
    min-height: 6vh;
    height: auto;
    padding: 1vh 1.5vw 1vh 3vw;
  }

  :deep(.el-menu-item-group__title) {
    padding: 0;
  }

  :deep(.el-submenu__icon-arrow) {
    right: 1.5vw;
    margin-top: 0;
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
  }

  // 图标样式
  :deep([class^="el-icon-"]) {
    margin-right: 1vw;
    font-size: clamp(16px, 2vh, 20px);
    width: 2em;
    text-align: center;
    flex-shrink: 0;
  }

  // 文本内容样式
  :deep(.el-submenu__title span),
  :deep(.el-menu-item span) {
    flex: 1;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
  }
}

// 响应式布局
@media screen and (max-width: 1200px) {
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 30%;
  }
  
  :deep(.el-submenu__title),
  :deep(.el-menu-item) {
    padding: 1vh 1vw;
  }
}

@media screen and (max-width: 768px) {
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 35%;
    min-width: 250px;
  }
}
</style>

<script>
export default {
  data() {
    return {
      menuData: [
        {
          path: "/center/information",
          name: "information",
          label: "Personal Information",
          icon: "user",
          title: "1",
          url: "Center",
        },
        {
          label: "Committee Management",
          icon: "s-custom",
          title: "2",
          children: [
            {
              path:"/center/createcommittees",
              name:"createCommittee",
              label: "Create Committee",
            },
            {
              path: "/center/selfCommittee",
              name:"selfCommittee",
              label: "Manage Committee",
              icon: "el-icon-plus",
            }
          ]
        },
        {
          label: "Conference Management",
          icon: "s-marketing",
          title: "3", 
          children: [
            {
              path: "/center/createConference",
              name: "createConference",
              label: "Create Conference",
              url: "Center/CreateConference",
            },
            {
              path: "/center/selfConference", 
              label: "Manage Conference",
              name: "selfConference",
              url: "center/selfConference",
            },
            {
              path: "/center/conferenceIssues",
              label: "Conference Issues",
              name: "conferenceIssues",
              url: "center/conferenceIssues",
            }
          ]
        },
        {
          label: "Journal Management",
          icon: "document",
          title: "4",
          children: [
            {
              path: "/center/createJournal",
              label: "Create Journal",
              name: "createJournal",
            },
            {
              path: "/center/selfJournal",
              label: "Manage Journal",
              name: "selfJournal",
            },
            {
              path: "/center/journalIssues",
              label: "Journal Issues",
              name: "journalIssues",
            }
          ]
        },
        {
          label: "Paper Management",
          icon: "document-copy",
          title: "5",
          children: [
            {
              path: "/center/papers",
              label: "My Papers",
              name:"papers",
            },
            {
              path:"/center/ReviewedPapers",
              name:"reviewedPapers",
              label: "My Papers (Reviewed)",
            },
            {
              path:"/center/InReviewPapers",
              name:"inReviewPapers",
              label: "My Papers (Under Review)",
            },
            {
              path: "/center/inReview",
              name:"inReview",
              label: "Papers Under Review",
            },
            {
              path: "/center/Reviewed",
              label: "Papers I Reviewed",
              name:"Reviewed",
            },
            {
              path: "/center/Reviews",
              label: "Papers To Review",
              name:"Reviews",
            }
          ]
        },
        {
          label: "NFT Management",
          icon: "trophy-1",
          title: "6",
          children: [
            {
              path: "/center/MyNFTs",
              label: "My NFTs",
              name: "myNFTs",
            },
            {
              path: "/center/NFTSelling",
              label: "NFT Trading",
              name: "nftSelling",
            },
            {
              path: "/center/Mint",
              label: "Mint NFT",
              name: "mint",
            }
          ]
        },
        {
          path: "/center/Users",
          name: "users",
          label: "User Management",
          icon: "user-solid",
          title: "7"
        }
      ]
    };
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },
    clickMenu(item) {
      // 当页面的路由域跳转的路由不一致才允许跳转
      if (
        this.$route.path !== item.path &&
        !(this.$route.path === "/center/information" && item.path === "/center")
      ) {
        this.$router.push(item.path);
      }
    },
  },
  computed: {
    // 没有子菜单
    noChildren() {
      // 对当前的menuData进行遍历，item会拿到其中的某一项，根据true和false来判断当前的数据是否返回
      return this.menuData.filter((item) => !item.children);
    },
    // 有子菜单
    hasChildren() {
      return this.menuData.filter((item) => item.children);
    },
    menuData() {
      // 判断当前数据，如果缓存中没有，当前store中去获取
      // return JSON.parse(Cookies.get("menu")) || this.$store.state.tab.menu;
      // 从local store中获取
      return (
        JSON.parse(localStorage.getItem("menu")) || this.$store.state.tab.menu
      );
    },
    isCollapse() {
      return this.$store.state.tab.isCollapse;
    },
  },
};
</script>