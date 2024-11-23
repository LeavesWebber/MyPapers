<!--作用：显示用户自己提交的正在评审中的论文
功能：
查看论文详情
删除未评审的论文
分页显示
-->
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
        placeholder="Enter title keywords to query"
      ></el-input>
    </div>
    <el-table
      :data="displayedData"
      height="500"
      border
      stripe
      style="width: 100%"
      :header-cell-style="{ color: '#000000', background: '#F7FBFF' }"
      :cell-style="cellStyle"
      :default-sort="{ prop: 'CreatedAt', order: 'descending' }"
    >
      <el-table-column fixed prop="title" label="Title" width="240px">
      </el-table-column>
      <el-table-column prop="block_address" label="Block Address">
      </el-table-column>
      <!-- <el-table-column prop="authors" label="Author"> </el-table-column> -->
      <el-table-column
        prop="CreatedAt"
        label="Publication Time"
        width="180px"
        sortable
      >
      </el-table-column>
      <!-- <el-table-column prop="paper_type" label="Type"> </el-table-column> -->
      <el-table-column prop="status" label="Status" width="150px">
      </el-table-column>
      <el-table-column fixed="right" label="Options" width="160">
        <template slot-scope="scope">
          <el-button @click="viewPaper(scope.row)" type="primary" size="mini"
            >View</el-button
          >
          <el-button
            @click="popup(scope.row)"
            type="danger"
            size="mini"
            v-if="scope.row.status === 'UnReview'"
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
import { getSelfPapers, deletePaper } from "../../api";
export default {
  data() {
    return {
      tableData: [
        {
          ID: 0,
          authors: "",
          CreatedAt: "",
          UpdatedAt: "",
          conference_id: 0,
          journal_id: 0,
          title: "",
          paper_type: "",
          abstract: "",
          key_words: "",
          subject_category: "",
          manuscript_id: "",
          informed_consent: "",
          animal_subjects: "",
          cor_author: "",
          manuscript_type: "",
          unique_contribution: "",
          hash: "",
          block_address: "",
          filepath: "",
          cid: "",
          status: "",
          user: [
            {
              ID: 0,
              CreatedAt: "",
              UpdatedAt: "",
              uuid: 0,
              authorityId: 0,
              sex: 0,
              username: "",
              first_name: "",
              last_name: "",
              headerImg: "",
              email: "",
              phone: "",
              address: "",
              education: "",
              title: "",
              research: "",
              block_chain_address: "",
              affiliation: "",
              affiliation_type: "",
            },
          ],
        },
      ],
      centerDialogVisible: false,
      selectRow: null,
      currentPage: 1,
      pageSize: 10,
      searchKeyword: "",
    };
  },
  methods: {
    viewPaper(row) {
      console.log(row.ID);
      this.$router.push({
        path: "/center/detailPaper",
        query: {
          paper_id: row.ID,
        },
      });
    },
    indexMethod(index) {
      console.log(index);
      return index + 1;
    },
    popup(row) {
      this.centerDialogVisible = true;
      this.selectRow = row;
    },
    getPapers() {
      getSelfPapers().then((res) => {
        if (res.data.data === null) {
          this.tableData = [];
          return;
        }
        // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
        for (let i = 0; i < res.data.data.length; i++) {
          res.data.data[i].CreatedAt = res.data.data[i].CreatedAt.substring(
            0,
            10
          );
        }

        this.tableData = res.data.data;
        // 把tableData里面的user数组里面的first_name和last_name拼接起来
        for (let i = 0; i < this.tableData.length; i++) {
          for (let j = 0; j < this.tableData[i].user.length; j++) {
            // 如果authors没有初始化，就初始化为空
            if (this.tableData[i].authors === undefined) {
              this.tableData[i].authors = "";
            }
            // 如果是最后一个作者，就不加逗号
            if (j === this.tableData[i].user.length - 1) {
              this.tableData[i].authors +=
                this.tableData[i].user[j].first_name +
                " " +
                this.tableData[i].user[j].last_name;
            } else {
              this.tableData[i].authors +=
                this.tableData[i].user[j].first_name +
                " " +
                this.tableData[i].user[j].last_name +
                ",";
            }
          }
        }
        console.log(this.tableData);
      });
    },
    handleDelete(row) {
      deletePaper({ paper_id: row.ID }).then((res) => {
        console.log(res.data);
        if (res.data.code === 1000) {
          this.$message({
            message: "Delete Success",
            type: "success",
          });
          this.centerDialogVisible = false;
          // 刷新页面
          this.getPapers();
        }
      });
    },
    searchData() {
      // 当搜索关键字发生变化时，重置当前页码为第一页
      this.currentPage = 1;
    },
    changePage(currentPage) {
      console.log(currentPage, this.currentPage, "currentPage");
      this.currentPage = currentPage;
      console.log(currentPage, this.currentPage, "currentPage");
    },
    cellStyle(row, column, rowIndex, columnIndex) {
      const status = row.row.status;
      console.log(status, row.column.label, "row.row.Status");
      // Accept（绿色）、Reject或Reject and Resubmit（红色）、UnReview（橙色）、InReview（蓝色）、Minor Revisions或者Major Revisions（紫色）
      if (status === "Accept" && row.column.label === "Status") {
        return "color :#13CE66";
      } else if (
        (status === "Reject" || status === "Reject and Resubmit") &&
        row.column.label === "Status"
      ) {
        return "color :#FF4949";
      } else if (status === "UnReview" && row.column.label === "Status") {
        return "color :#F7BA2A";
      } else if (status === "InReview" && row.column.label === "Status") {
        return "color :#409EFF";
      } else if (
        (status === "Minor Revisions" || status === "Major Revisions") &&
        row.column.label === "Status"
      ) {
        return "color :#9C6ADE";
      }
    },
  },
  computed: {
    tableDataCount() {
      return this.tableData.length;
    },
    displayedData() {
      // 进行查询过滤
      const filteredData = this.tableData.filter((item) =>
        item.title.toLowerCase().includes(this.searchKeyword.toLowerCase())
      );
      // 进行分页处理
      const startIndex = (this.currentPage - 1) * this.pageSize;

      const endIndex = startIndex + this.pageSize;
      if (endIndex > filteredData.length) {
        return filteredData.slice(startIndex, filteredData.length);
      }
      return filteredData.slice(startIndex, endIndex);
    },
  },
  mounted() {
    this.getPapers();
  },
};
</script>
<style lang="less" scoped>
.query {
  margin-bottom: 10px;
  text-align: right;
}
</style>