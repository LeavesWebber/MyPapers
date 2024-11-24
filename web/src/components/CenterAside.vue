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
  width: 200px;
  min-height: 100%;
}
.el-menu {
  height: 100vh;
  border-right: none;
  h4 {
    color: #fff;
    text-align: center;
    line-height: 48px;
    font-size: 16px;
    font-weight: 400;
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
          label: "个人信息",
          icon: "user",
          title: "1",
          url: "Center",
        },
        {
          label: "委员会管理",
          icon: "s-custom",
          title: "2",
          children: [
            {
              path:"/center/createcommittees",
              name:"createCommittee",
              label: "创建委员会",
            },
            {
              path: "/center/selfCommittee",
              name:"selfCommittee",
              label: "管理委员会",
            }
          ]
        },
        {
          label: "会议管理",
          icon: "s-marketing",
          title: "3", 
          children: [
            {
              path: "/center/createConference",
              name: "createConference",
              label: "创建会议",
              url: "Center/CreateConference",
            },
            {
              path: "/center/selfConference", 
              label: "管理会议",
              name: "selfConference",
              url: "center/selfConference",
            },
            {
              path: "/center/conferenceIssues",
              label: "会议期刊",
              name: "conferenceIssues",
              url: "center/conferenceIssues",
            }
          ]
        },
        {
          label: "期刊管理",
          icon: "document",
          title: "4",
          children: [
            {
              path: "/center/createJournal",
              label: "创建期刊",
              name: "createJournal",
            },
            {
              path: "/center/selfJournal",
              label: "管理期刊",
              name: "selfJournal",
            },
            {
              path: "/center/journalIssues",
              label: "期刊期次",
              name: "journalIssues",
            }
          ]
        },
        {
          label: "论文管理",
          icon: "document-copy",
          title: "5",
          children: [
            {
              path: "/center/papers",
              label: "我的论文",
              name:"papers",
            },
            {
              path:"/center/ReviewedPapers",
              name:"reviewedPapers",
              label: "我的论文（已审核）",
            },
            {
              path:"/center/InReviewPapers",
              name:"inReviewPapers",
              label: "我的论文（待审核）",
            },
            {
              path: "/center/inReview",
              name:"inReview",
              label: "审核中论文",
            },
            {
              path: "/center/Reviewed",
              label: "我已审核的论文",
              name:"Reviewed",
            },
            {
              path: "/center/Reviews",
              label: "待审核论文",
              name:"Reviews",
            }
          ]
        },
        {
          label: "NFT管理",
          icon: "picture",
          title: "6",
          children: [
            {
              path: "/center/MyNFTs",
              label: "我的NFT",
              name: "myNFTs",
            },
            {
              path: "/center/NFTSelling",
              label: "NFT交易",
              name: "nftSelling",
            },
            {
              path: "/center/Mint",
              label: "铸造NFT",
              name: "mint",
            }
          ]
        },
        {
          path: "/center/Users",
          name: "users",
          label: "用户管理",
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