<template>
  <div class="box">
    <div class="box1">
      <!-- <el-button type="primary" @click="currentIssue()">
        Current Issue
      </el-button>
      <el-button type="primary" @click="AllIssues()"> All Issues </el-button> -->
      <div>
        <el-button
          v-for="(year, index) in years"
          type="text"
          :key="index"
          :class="{ highlighted: index === highlightedYear }"
          @click="yearClick(index)"
        >
          {{ year }}
        </el-button>
      </div>
      <div>
        <el-button
          v-for="(issue, index) in issues"
          type="text"
          :key="index"
          :class="{ highlighted: index === highlightedIssue }"
          @click="issueClick(issue, index)"
        >
          {{ issue.name }}
        </el-button>
      </div>
      <!-- 显示开始和结束时间 -->
      <div style="text-align: center">
        <h3>
          Time: {{ currentIssue.submission_start_time }} ——
          {{ currentIssue.submission_end_time }}
        </h3>
        <h3>Volume: {{ currentIssue.volume }}</h3>
      </div>
      <el-table
        :data="journalPapers"
        :empty-text="emptyText"
        height="550"
        border
        style="width: 100%"
        stripe
        :header-cell-style="{
          color: '#000000',
          background: '#F7FBFF',
        }"
      >
        <el-table-column fixed prop="title" label="Title"> </el-table-column>
        <el-table-column prop="block_address" label="Block Address">
        </el-table-column>
        <el-table-column prop="authors" label="Author"> </el-table-column>
        <el-table-column
          prop="CreatedAt"
          label="Publication Time"
          width="140px"
        >
        </el-table-column>
        <el-table-column prop="paper_type" label="Type"> </el-table-column>
        <!-- <el-table-column prop="status" label="Status" width="100px">
        </el-table-column> -->
        <el-table-column fixed="right" label="Options" width="100">
          <template slot-scope="scope">
            <el-button @click="details(scope.row)" type="text" size="small"
              >View</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>
<script>
import { getJournalIssues, getJournalIssuePapers } from "../../api";
export default {
  data() {
    return {
      years: [],
      issues: [],
      currentIssue: {},
      tableData: [
        {
          ID: 0,
          CreatedAt: "",
          journal_id: "",
          name: "",
          submission_start_time: "",
          submission_end_time: "",
          description: "",
          year: "",
          volume: 0,
        },
      ],
      journalPapers: [
        {
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
      ],
      highlightedYear: 0,
      highlightedIssue: 0,
      emptyText: "No Papers",
    };
  },
  methods: {
    details(row) {
      console.log(row.ID);
      this.$router.push({
        path: "/detailPaper",
        query: {
          paper_id: row.ID,
        },
      });
    },
    yearClick(index) {
      this.highlightedYear = index;
      this.highlightedIssue = 0;
      // 先清空issues
      this.issues = [];
      // 遍历tableData，得到当前year下面的issues
      for (let i = 0; i < this.tableData.length; i++) {
        // 如果tableData里面的year等于years里面的year，就把这个issue添加到issues里面
        if (this.tableData[i].year === this.years[index]) {
          this.issues.push(this.tableData[i]);
        }
      }
      // issues根据submission_start_time从小到大排序
      this.issues.sort((a, b) => {
        return a.submission_start_time > b.submission_start_time ? 1 : -1;
      });
      this.currentIssue = this.issues[0];
      this.getPapers(this.issues[0]);
    },
    issueClick(issue, index) {
      this.highlightedIssue = index;
      console.log(issue, "issue");
      this.currentIssue = issue;
      this.getPapers(issue);
    },
    getPapers(issue) {
      getJournalIssuePapers({
        journal_id: this.$route.query.id,
        start_time: issue.submission_start_time + "T00:00:00Z",
        end_time: issue.submission_end_time + "T00:00:00Z",
      }).then((res) => {
        this.journalPapers = res.data.data;
        console.log(res.data.data, "res.data.data");
        // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
        for (let i = 0; i < res.data.data.length; i++) {
          res.data.data[i].CreatedAt = res.data.data[i].CreatedAt.substring(
            0,
            10
          );
        }
        // 把journalPapers里面的user数组里面的first_name和last_name拼接起来
        for (let i = 0; i < this.journalPapers.length; i++) {
          for (let j = 0; j < this.journalPapers[i].user.length; j++) {
            // 如果authors没有初始化，就初始化为空
            if (this.journalPapers[i].authors === undefined) {
              this.journalPapers[i].authors = "";
            }
            // 如果是最后一个作者，就不加逗号
            if (j === this.journalPapers[i].user.length - 1) {
              this.journalPapers[i].authors +=
                this.journalPapers[i].user[j].first_name +
                " " +
                this.journalPapers[i].user[j].last_name;
            } else {
              this.journalPapers[i].authors +=
                this.journalPapers[i].user[j].first_name +
                " " +
                this.journalPapers[i].user[j].last_name +
                ",";
            }
          }
        }
      });
    },

    getIssues() {
      getJournalIssues({ journal_id: this.$route.query.id }).then((res) => {
        this.tableData = [];
        // 如果返回值为空就不往下执行了
        console.log(res.data.data, "res.data.data");
        if (res.data.data === null) {
          return;
        }
        // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
        for (let i = 0; i < res.data.data.length; i++) {
          res.data.data[i].submission_start_time = res.data.data[
            i
          ].submission_start_time.substring(0, 10);
          res.data.data[i].submission_end_time = res.data.data[
            i
          ].submission_end_time.substring(0, 10);
        }

        this.tableData = res.data.data;
        console.log(this.tableData, "tableDate");
        // 处理tableData：里面有不同的year，进行分组，得到years，yeas下面有多个issues
        // 遍历tableData，得到不同的year
        for (let i = 0; i < this.tableData.length; i++) {
          // 如果years里面没有这个year，就添加进去
          if (!this.years.includes(this.tableData[i].year)) {
            this.years.push(this.tableData[i].year);
          }
        }
        // years从大到小排序
        this.years.sort((a, b) => {
          return b - a;
        });
        console.log(this.years, "years");
        //最近一年的issues
        for (let i = 0; i < this.tableData.length; i++) {
          if (this.tableData[i].year === this.years[0]) {
            this.issues.push(this.tableData[i]);
          }
        }
        this.currentIssue = this.issues[0];
        console.log(this.issues, "issues");
        this.getPapers(this.issues[0]);
      });
    },
  },
  mounted() {
    this.getIssues();
    // this.handleTableData();
  },
};
</script>
<style lang="less" scoped>
.highlighted {
  background-color: #d3e3fd; // 设置高亮时的背景颜色
}
.el-button {
  margin: 5px; /* 设置按钮之间的间距 */
  padding: 10px 20px; /* 设置按钮内边距 */
  font-size: 16px; /* 设置文字大小 */
  border-radius: 5px; /* 设置按钮圆角 */
  transition: background-color 0.3s, color 0.3s, border 0.3s; /* 添加过渡效果 */
}
.box {
  // width: 100%;
  // height: 3000px;
  // 盒子里面的内容水平居中
  // 宽高被盒子撑开
  display: inline-block;
  text-align: center;
  .box1 {
    // 取消内容水平居中
    text-align: left;
    width: 80%;
    height: 100%;
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
</style>
