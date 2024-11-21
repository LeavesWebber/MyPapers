<template>
  <div class="box">
    <div class="box1">
      <h1 style="text-align: center">Conferences</h1>
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
        :empty-text="emptyText"
        :data="displayedData"
        height="800"
        border
        stripe
        style="width: 100%"
        :header-cell-style="{ color: '#000000', background: '#F7FBFF' }"
      >
        <el-table-column fixed prop="ID" label="ID" width="80" sortable>
        </el-table-column>
        <el-table-column prop="name" label="Name" width="200">
        </el-table-column>
        <el-table-column prop="location" label="Location" width="300">
        </el-table-column>
        <el-table-column prop="start_time" label="Start Time" width="200">
        </el-table-column>
        <el-table-column prop="end_time" label="End Time" width="200">
        </el-table-column>
        <el-table-column fixed="right" label="">
          <template slot-scope="scope">
            <el-button @click="details(scope.row)" type="text" size="small"
              >Details</el-button
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
  </div>
</template>
<script>
import { getConferenceList } from "../../api";
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
        },
      ],
      currentPage: 1,
      pageSize: 10,
      searchKeyword: "",
      emptyText: "No Conferences",
    };
  },
  methods: {
    details(row) {
      console.log(row.ID);
      this.$router.push({
        path: "detailConference",
        query: {
          conference_id: row.ID,
        },
      });
    },

    getConferences() {
      getConferenceList().then((res) => {
        this.tableData = [];
        // 如果返回值为空就不往下执行了
        if (res.data.data === null) {
          return;
        }
        // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
        for (let i = 0; i < res.data.data.length; i++) {
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
.box {
  // width: 100%;
  // height: 3000px;
  // 盒子里面的内容水平居中
  // 宽高被盒子撑开
  display: inline-block;
  text-align: center;
  .image {
    width: 80%;
    height: 430px;
  }
  .box1 {
    // 取消内容水平居中
    text-align: left;
    width: 80%;
    height: 950px;
    // background-color: #e2f1fb;
    background-color: #ecf5ff;
    // 圆角
    border-radius: 10px;
    // 水平居中
    margin: 0 auto;
    margin-top: 20px;
    // 上下左右panding
    padding: 20px 20px;
    color: #072e5b;
  }
}
.query {
  margin-bottom: 10px;
  text-align: right;
}
.block {
  text-align: center;
}
</style>
