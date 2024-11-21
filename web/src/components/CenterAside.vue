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
           path: "/center",
           name: "information",
           label: "Information",
           icon: "user",
           url: "Center",
         },
         {
           
           name: "committees",
           label: "committees",
           icon: "user",
           children:[
            {
              path:"/center/createcommittees",
              name:"Create Committee",
              label:"Create Committee",
            },
            {
              path:"/center/mycommittees",
              name:"My Committees",
              label:"My Committees",
            }
           ]
         },
         {
           label: "Conference",
           icon: "location",
           children: [
             {
               path: "/center/myconference",
               name: "My conferences",
               label: "My Conferences",
               icon: "setting",
               url: "Center/MyConference",
             },
             {
               path: "/center/createConference",
               name: "Create Conference",
               label: "Create Conference",
               icon: "setting",
               url: "Center/CreateConference",
             },
           ],
         },
       ],
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