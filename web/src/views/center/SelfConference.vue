<template>
  <div>
    <el-dialog
      title="Warn"
      :visible.sync="centerDialogVisible"
      width="30%"
      append-to-body
    >
      <span>Are you sure to delete it?</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="centerDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="handleDelete(selectRow)"
          >Sure</el-button
        >
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
      style="width: 100%"
      :header-cell-style="{ color: '#000000', background: '#F7FBFF' }"
      :default-sort="{ prop: 'CreatedAt', order: 'descending' }"
    >
      <el-table-column fixed prop="ID" label="ID" width="90" sortable> </el-table-column>
      <el-table-column prop="name" label="Name" width="330"> </el-table-column>
      <el-table-column prop="location" label="Location">
      </el-table-column>
      <el-table-column prop="start_time" label="Start Time" width="120">
      </el-table-column>
      <el-table-column prop="end_time" label="End Time" width="120">
      </el-table-column>
      
      <!-- <el-table-column prop="category" label="Category" width="120">
      </el-table-column> -->
      <el-table-column fixed="right" label="Options" width="165">
        <template slot-scope="scope">
          <el-button @click="details(scope.row)" type="primary" size="mini"
            >Details</el-button
          >
          <el-button
            @click="popup(scope.row)"
            type="danger"
            size="mini"
            v-if="scope.row.creator_id === userInfo.ID"
            >Delete</el-button
          >
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
import { getSelfConferenceList, deleteConference } from "../../api";
export default {
  data() {
    return {
      tableData: [
        {
          ID: 0,
          name: "",
          description: "",
          location: "",
          start_time: "",
          end_time: "",
          creator_id: 0,
        },
      ],
      selectRow: null,
      isPresident: false,
      userInfo: {},
      centerDialogVisible: false,
      currentPage: 1,
      pageSize: 7,
      searchKeyword: "",
    };
  },
  methods: {
    details(row) {
      console.log(row.ID);
      this.$router.push({
        path: "/center/detailConference",
        query: {
          conference_id: row.ID,
        },
      });
    },
    popup(row) {
      this.centerDialogVisible = true;
      this.selectRow = row;
    },
    handleDelete(row) {
      console.log(row.ID);
      deleteConference({ conference_id: row.ID }).then((res) => {
        console.log(res);
        // 先弹窗提示是否删除

        if (res.data.code === 1000) {
          this.getConferences();
          this.$message({
            message: "Delete Success",
            type: "success",
          });
        } else {
          this.$message({
            message: "Delete Failed",
            type: "error",
          });
        }
        this.centerDialogVisible = false;
        // 刷新页面
        this.getConferences();
      });
    },
    getConferences() {
      getSelfConferenceList().then((res) => {
        this.tableData = [];
        // 如果返回值为空就不往下执行了
        if (res.data.data === null) {
          return;
        }
        // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
        for (let i = 0; i < res.data.data.length; i++) {
          res.data.data[i].CreatedAt = res.data.data[i].CreatedAt.substring(
            0,
            10
          );
          res.data.data[i].start_time = res.data.data[i].start_time.substring(
            0,
            10
          );
          res.data.data[i].end_time = res.data.data[i].end_time.substring(
            0,
            10
          );
        }

        this.tableData = res.data.data;
        console.log(this.tableData, "tableDate");
      });
      // 从localStorage里面获取用户信息
      this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
      console.log(this.userInfo, "userInfo");
      // 如果用户是主席角色，就显示删除按钮
      // if (this.userInfo.authorityId === 102) {
      //   this.isPresident = true;
      // }
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
        item.name.toLowerCase().includes(this.searchKeyword.toLowerCase())
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
  },
  mounted() {
    this.getConferences();
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
</style>