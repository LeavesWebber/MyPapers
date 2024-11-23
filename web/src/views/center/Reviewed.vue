<!--作用：显示已评审论文的管理页面
功能：
分配评审员
查看评审状态和评论
处理评审结果-->
<template>
  <div>
    <el-dialog
      title="Please input reviewer name"
      :visible.sync="centerDialogVisible"
      width="30%"
      append-to-body
    >
      <el-tag
        :key="tag"
        v-for="tag in reviewerNames"
        closable
        :disable-transitions="false"
        @close="handleClose(tag)"
      >
        {{ tag }}
      </el-tag>
      <el-input
        class="input-new-tag"
        v-if="inputVisible"
        v-model="inputValue"
        ref="saveTagInput"
        size="small"
        @keyup.enter.native="handleInputConfirm"
        @blur="handleInputConfirm"
      >
      </el-input>
      <el-button v-else class="button-new-tag" size="small" @click="showInput"
        >+ Reviewer Name</el-button
      >
      <span slot="footer" class="dialog-footer">
        <el-button @click="centerDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="handleAllot(selectRow)"
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
      height="490"
      border
      style="width: 100%"
      :header-cell-style="{ color: '#000000', background: '#F7FBFF' }"
      :cell-style="cellStyle"
      :default-sort="{ prop: 'CreatedAt', order: 'descending' }"
    >
      <el-table-column fixed prop="paper.title" label="Title" width="240px">
      </el-table-column>
      <el-table-column prop="paper.block_address" label="Block Address">
      </el-table-column>
      <!-- <el-table-column prop="paper.authors" label="Author"> </el-table-column> -->
      <el-table-column
        prop="paper.CreatedAt"
        label="Publication Time"
        width="180px"
        sortable
      >
      </el-table-column>
      <!-- <el-table-column prop="paper.paper_type" label="Type"> </el-table-column> -->
      <el-table-column prop="paper.status" label="Status" width="150px">
      </el-table-column>
      <el-table-column fixed="right" label="Options" width="160">
        <template slot-scope="scope">
          <el-button @click="viewPaper(scope.row)" type="primary" size="mini"
            >View</el-button
          >
          <!-- <el-button
            @click="popup(scope.row)"
            type="success"
            size="mini"
            v-if="isPresident"
            >Allot</el-button
          > -->
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
import { getReviews, allotReviewers } from "../../api";
export default {
  data() {
    return {
      tableData: [
        {
          paper: {
            ID: 0,
            authors: "",
            CreatedAt: "",
            UpdatedAt: "",
            conference_id: 4,
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
          review_infos: [
            {
              reviewer_name: "",
              comment: "",
              status: "",
            },
          ],
        },
      ],
      centerDialogVisible: false,
      selectRow: null,
      isPresident: false,
      userInfo: {},
      reviewerNames: [],
      inputVisible: false,
      inputValue: "",
      param: {
        paper_id: 0,
        reviewer_names: [],
      },
      currentPage: 1,
      pageSize: 6,
      searchKeyword: "",
    };
  },
  methods: {
    viewPaper(row) {
      console.log(row.paper.ID);
      this.$router.push({
        path: "/center/reviewPaperDetail",
        query: {
          paper_id: row.paper.ID,
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
      console.log(this.selectRow, "selectRow");
    },
    getPapers() {
      getReviews({ filter: "Reviewed" }).then((res) => {
        this.tableData = [];
        // 如果返回值为空就不往下执行了
        if (res.data.data === null) {
          return;
        }
        // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
        for (let i = 0; i < res.data.data.length; i++) {
          res.data.data[i].paper.CreatedAt = res.data.data[
            i
          ].paper.CreatedAt.substring(0, 10);
        }

        this.tableData = res.data.data;
        // 把paper里面的user数组里面的first_name和last_name拼接起来
        for (let j = 0; j < this.tableData.length; j++) {
          let authors = "";
          for (let i = 0; i < this.tableData[j].paper.user.length; i++) {
            authors +=
              this.tableData[j].paper.user[i].first_name +
              " " +
              this.tableData[j].paper.user[i].last_name +
              ", ";
          }
          this.tableData[j].paper.authors = authors.substring(
            0,
            authors.length - 2
          );
        }
        console.log(this.tableData, "tableDate");
      });
      // 从localStorage里面获取用户信息
      this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
      console.log(this.userInfo, "userInfo");
      // 如果用户是主席，就显示分配按钮，遍历this.userInfo.authorities
      // for (let i = 0; i < this.userInfo.authorities.length; i++) {
      //   if (
      //     this.userInfo.authorityId === 101 ||
      //     this.userInfo.authorityId === 102
      //   ) {
      //     this.isPresident = true;
      //   }
      // }
    },
    handleAllot(row) {
      console.log(row, "row");
      console.log(this.reviewerNames, "reviewerNames1");
      this.param.paper_id = row.paper.ID;
      this.param.reviewer_names = this.reviewerNames;
      allotReviewers(this.param).then((res) => {
        console.log(res);
        if (res.data.code === 1000) {
          this.centerDialogVisible = false;
          this.getPapers();
          this.$message({
            message: "Allot Success",
            type: "success",
          });
        } else {
          this.$message({
            message: "Allot Failed",
            type: "error",
          });
        }
      });
    },
    handleClose(tag) {
      this.reviewerNames.splice(this.reviewerNames.indexOf(tag), 1);
    },

    showInput() {
      this.inputVisible = true;
      this.$nextTick((_) => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },

    handleInputConfirm() {
      let inputValue = this.inputValue;
      if (inputValue) {
        this.reviewerNames.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = "";
    },
    searchData() {
      // 当搜索关键字发生变化时，重置当前页码为第一页
      this.currentPage = 1;
    },
    changePage(currentPage) {
      this.currentPage = currentPage;
      console.log(currentPage, this.currentPage, "currentPage");
    },
    cellStyle(row, column, rowIndex, columnIndex) {
      const status = row.row.paper.status;
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
        item.paper.title
          .toLowerCase()
          .includes(this.searchKeyword.toLowerCase())
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