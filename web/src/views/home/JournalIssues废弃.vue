<template>
  <div class="box">
    <div class="box1">
      <el-button
        v-for="(issue, index) in issueList"
        :key="index"
        @click="handleIssueClick(index)"
      >
        {{ issue }}
      </el-button>
      <!-- 显示开始和结束时间 -->
      <div>
        Time: {{ quarters[index].startDate }} —— {{ quarters[index].endDate }}
      </div>
      <el-table :data="journalPapers" height="550" border style="width: 100%">
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
import { getIssues } from "../../api";
export default {
  data() {
    return {
      issueList: [], // 存储从后端获取的分割后的时间数据
      issue_period: "",
      id: "",
      create_at: "",
      quarters: [{ startDate: "", endDate: "", count: 0 }],
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
      index: 0,
    };
  },
  mounted() {
    this.id = this.$route.query.id;
    this.issue_period = this.$route.query.issue_period;
    this.create_at = this.$route.query.create_at; // 2023-08-13

    let now = new Date();
    // now转为字符串
    now = now.toISOString().split("T")[0];
    this.calculateQuarters(this.create_at, now);
    this.generateIssueList(this.quarters.length);
    // 默认请求最近一期的数据
    this.handleIssueClick(this.issueList.length - 1);
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
    generateIssueList(splitLength) {
      // 生成按钮列表，以 "Issue 1", "Issue 2", ... 的形式
      for (let i = 1; i <= splitLength; i++) {
        this.issueList.push(`Issue ${i}`);
      }
    },
    calculateQuarters(startDate, endDate) {
      const end = new Date(endDate);
      let quarterCount = 0;
      let slowDate = new Date(startDate);
      let fastDate = new Date(startDate);
      this.quarters = [];
      fastDate.setMonth(fastDate.getMonth() + 3);
      while (fastDate < end) {
        this.quarters.push({
          startDate: slowDate.toISOString().split("T")[0],
          endDate: fastDate.toISOString().split("T")[0],
          count: ++quarterCount,
        });

        fastDate.setMonth(fastDate.getMonth() + 3);
        slowDate.setMonth(slowDate.getMonth() + 3);
      }
    },
    handleIssueClick(index) {
      // 点击按钮后，根据index获取对应的季度的开始时间
      let startDate = this.quarters[index].startDate;
      // startDate = 转为2023-10-01T00:00:00Z格式
      startDate = startDate + "T00:00:00Z";
      // 请求后端数据
      getIssues({ journal_id: this.id, start_time: startDate }).then((res) => {
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
      this.index = index;
    },
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
