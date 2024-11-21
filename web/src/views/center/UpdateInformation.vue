<template>
  <el-row>
    <el-col :span="24"
      ><div class="box">
        <el-form
          :model="userForm"
          :inline="true"
          status-icon
          ref="userForm"
          label-width="180px"
          class="demo-userForm"
          size="mini"
        >
          <el-form-item label="UserName" prop="data">
            <el-input v-model="userForm.username"></el-input>
          </el-form-item>
          <el-form-item label="FirstName" prop="data">
            <el-input v-model="userForm.first_name"></el-input>
          </el-form-item>
          <el-form-item label="LastName" prop="data">
            <el-input v-model="userForm.last_name"></el-input>
          </el-form-item>
          <!-- <el-form-item label="Sex" prop="data">
          <el-input v-model="userForm.sex"></el-input>
        </el-form-item> -->
          <el-form-item label="Email" prop="data">
            <el-input v-model="userForm.email"></el-input>
          </el-form-item>
          <el-form-item label="Department" prop="data">
            <el-input v-model="userForm.department"></el-input>
          </el-form-item>
          <el-form-item label="Phone" prop="data">
            <el-input v-model="userForm.phone"></el-input>
          </el-form-item>
          <el-form-item label="Address" prop="data">
            <el-input v-model="userForm.address"></el-input>
          </el-form-item>
          <el-form-item label="Education" prop="data">
            <el-input v-model="userForm.education"></el-input>
          </el-form-item>
          <el-form-item label="Title" prop="data">
            <el-input v-model="userForm.title"></el-input>
          </el-form-item>
          <el-form-item label="Research" prop="data">
            <el-input v-model="userForm.research"></el-input>
          </el-form-item>
          <el-form-item label="Blockchain Addr" prop="data">
            <el-input v-model="userForm.block_chain_address"></el-input>
          </el-form-item>
          <el-form-item label="Affiliations" prop="data">
            <el-input v-model="userForm.affiliation"></el-input>
          </el-form-item>
          <el-form-item label="AffiliationType" prop="data">
            <el-input v-model="userForm.affiliation_type"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" class="button" @click="setSelfInfo"
              >Update</el-button
            >
          </el-form-item>
        </el-form>
      </div></el-col
    >
  </el-row>
</template>
<script>
import { getSelfInfo, setSelfInfo } from "../../api";
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
      },
    };
  },
  methods: {
    setSelfInfo() {
      setSelfInfo(this.userForm).then((res) => {
        console.log(res.data);
        if (res.data.code === 1000) {
          // localStorage.setItem("token", data.data.token); // 用localStorage缓存token值
          this.$message({
            message: "Update Success",
            type: "success",
          });
          this.$router.push("/center/information");
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
  margin-bottom: 20px;
  &:last-child {
    margin-bottom: 0;
  }
}
.el-col {
  border-radius: 4px;
}
.box {
  // 上内边距
  padding: 50px;
  //左内边距
  padding-left: 40px;
  //右内边距
  padding-right: 100px;
  //   background: #ECF5FF;
  border-radius: 8px;
  height: 570px;
  .button {
    margin-left: 289px;
  }
}
</style>