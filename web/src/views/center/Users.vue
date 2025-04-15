<template>
  <div>
    <el-dialog
      title="Edit Rule"
      :visible.sync="centerDialogVisible"
      width="30%"
      append-to-body
    >
      <el-select v-model="value" placeholder="Select">
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        >
        </el-option>
      </el-select>
      <span slot="footer" class="dialog-footer">
        <el-button @click="centerDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="edit(selectRow)">OK</el-button>
      </span>
    </el-dialog>
    <div class="query">
      <span> Query: </span>
      <el-input
        style="width: 260px"
        v-model="searchKeyword"
        @input="searchData"
        placeholder="Enter name keywords to query"
      ></el-input>
    </div>
    <el-table
      :data="displayedData"
      height="550"
      border
      stripe
      :header-cell-style="{
        color: '#000000',
        background: '#F7FBFF',
      }"
      :default-sort="{ prop: 'ID', order: 'descending' }"
    >
      <el-table-column fixed prop="id" label="ID" width="100" sortable>
      </el-table-column>
      <el-table-column prop="authority_id" label="Rule" width="100">
        <template slot-scope="scope">
          <el-tag
            v-if="scope.row.authority_id === 101"
            type="success"
            close-transition
            >Admin</el-tag
          >
          <el-tag
            v-else-if="scope.row.authority_id === 102"
            type="info"
            close-transition
            >President</el-tag
          >
          <el-tag
            v-else-if="scope.row.authority_id === 103"
            type="warning"
            close-transition
            >Member</el-tag
          >
          <el-tag
            v-else-if="scope.row.authority_id === 104"
            type="danger"
            close-transition
            >User</el-tag
          >
        </template>
      </el-table-column>
      <el-table-column prop="username" label="UserName" width="200">
      </el-table-column>
      <el-table-column prop="first_name" label="First Name" width="100">
      </el-table-column>
      <el-table-column prop="last_name" label="Last Name" width="100">
      </el-table-column>
      <el-table-column prop="email" label="Email" width="300">
      </el-table-column>
      <el-table-column prop="department" label="Department" width="200">
      </el-table-column>
      <el-table-column prop="phone" label="Phone" width="200">
      </el-table-column>
      <el-table-column prop="address" label="Address" width="300">
      </el-table-column>
      <el-table-column fixed="right" label="Options">
        <template slot-scope="scope">
          <el-button
            @click="popup(scope.row)"
            type="primary"
            size="mini"
            :disabled="!isAdmin"
          >
            Edit Rule
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="block">
      <el-pagination
        @current-change="changePage"
        :current-page="currentPage"
        :page-size="pageSize"
        :hide-on-single-page="true"
        layout="prev, pager, next"
        :total="tableDataCount"
      >
      </el-pagination>
    </div>
  </div>
</template>

<script>
import { getUserList, setUserInfo } from "../../api";
export default {
  data() {
    return {
      tableData: [
        {
          id: 0,
          authority_id: 0,
          username: "",
          first_name: "",
          last_name: "",
          email: "",
          department: "",
          phone: "",
          address: "",
        },
      ],
      selectRow: null,
      userInfo: {},
      reviewerNames: [],
      inputVisible: false,
      centerDialogVisible: false,
      currentPage: 1,
      pageSize: 7,
      searchKeyword: "",
      options: [
        {
          value: '102',
          label: 'President'
        },
        {
          value: '103',
          label: 'Member'
        },
        {
          value: '104',
          label: 'User'
        }
      ],
      value: '',
    };
  },
  methods: {
    async edit(row) {
      try {
        console.log('Edit user data:', {
          currentUser: this.userInfo,
          targetUser: row,
          newAuthority: this.value
        });

        const requestData = {
          id: Number(row.id),
          authority_id: Number(this.value)
        };


        const res = await setUserInfo(requestData);

        if (res.data.code === 1000) {
          this.$message.success('修改成功');
          await this.getUsers();
          this.centerDialogVisible = false;
        } else {
          this.$message.error(res.data.msg || '修改失败');
        }
      } catch (err) {
        this.$message.error('操作失败');
      }
    },
    popup(row) {
      if (!this.isAdmin) {
        this.$message.warning('只有管理员可以编辑权限');
        return;
      }
      this.centerDialogVisible = true;
      this.selectRow = row;
      this.value = row.authority_id.toString();
    },
    async initUserInfo() {
      try {
        const userInfoStr = localStorage.getItem("userInfo");
        if (!userInfoStr) return;
        const userInfo = JSON.parse(userInfoStr);
        if (userInfo) {
          this.userInfo = userInfo;
        }
      } catch (error) {
        this.$message.error('获取用户信息失败');
      }
    },
    async getUsers() {
      try {
        const res = await getUserList();
        this.tableData = res.data.data;
        await this.initUserInfo();
      } catch (error) {
        console.error('Failed to get users:', error);
        this.$message.error('获取用户列表失败');
      }
    },
    searchData() {
      // 当搜索关键字发生变化时，重置当前页码为第一页
      this.currentPage = 1;
    },
    changePage(currentPage) {
      this.currentPage = currentPage;
      console.log(currentPage, this.currentPage, "currentPage");
    },
  },
  computed: {
    tableDataCount() {
      return this.tableData.length;
    },
    displayedData() {
      // 进行查询过滤
      const filteredData = this.tableData.filter((item) =>
        item.username.toLowerCase().includes(this.searchKeyword.toLowerCase())
      );
      // 进行分页处理
      const startIndex = (this.currentPage - 1) * this.pageSize;

      const endIndex = startIndex + this.pageSize;
      console.log(startIndex, endIndex, "startIndex, endIndex");
      if (endIndex > filteredData.length) {
        return filteredData.slice(startIndex, filteredData.length);
      }
      return filteredData.slice(startIndex, endIndex);
    },
    isAdmin() {
      const authorityId = Number(this.userInfo?.authorityId);
      return authorityId === 101;
    },
  },
  created() {
    this.initUserInfo();
  },
  mounted() {
    this.getUsers();
  },
};
</script>
<style lang="less" scoped>
.el-tag + .el-tag {
  margin-left: 10px;
}
.button-new-tag {
  margin-left: 10px;
  height: 32px;
  line-height: 30px;
  padding-top: 0;
  padding-bottom: 0;
}
.input-new-tag {
  width: 90px;
  margin-left: 10px;
  vertical-align: bottom;
}
.query {
  margin-bottom: 10px;
  text-align: right;
}
.el-button.is-disabled {
  cursor: not-allowed;
  opacity: 0.6;
}
</style>