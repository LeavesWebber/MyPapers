<template>
  <div class="box">
    <div class="box1">
      <el-row>
        <el-col :span="2"
          ><div class="grid-content bg-purple-dark name">
            <el-button
              v-if="showAllIssueButton"
              size="small"
              type="primary"
              @click="Issues()"
              >All Issues</el-button
            >
          </div></el-col
        >
        <el-col :span="3"
          ><div class="grid-content bg-purple-dark name">
            <el-button
              v-if="showSubmitButton"
              size="small"
              type="primary"
              @click="Submit()"
              >Submit Manuscript</el-button
            >
          </div></el-col
        >
        <el-col :span="24"
          ><div class="grid-content bg-purple-dark name">
            <span>{{ journalDetail.name }}</span>
          </div></el-col
        >

        <el-col :span="22"
          ><div class="grid-content bg-purple-dark author">
            {{ journalDetail.authors }}
          </div></el-col
        >
        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">
            Description:
          </div></el-col
        >
        <el-col :span="20"
          ><div class="grid-content bg-purple-dark value">
            {{ journalDetail.description }}
          </div></el-col
        >
        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">Category:</div></el-col
        >
        <el-col :span="20"
          ><div class="grid-content bg-purple-dark value">
            {{ journalDetail.category }}
          </div></el-col
        >
        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">
            Create Time:
          </div></el-col
        >
        <el-col :span="20"
          ><div class="grid-content bg-purple-dark value">
            {{ journalDetail.create_at }}
          </div></el-col
        >

        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">
            Presidents:
          </div></el-col
        >
        <el-col :span="24"
          ><div class="grid-content bg-purple-dark value">
            <el-table
              :data="journalDetail.presidents"
              border
              style="width: 100%"
              stripe
              :header-cell-style="{
                color: '#000000',
                background: '#F7FBFF',
              }"
            >
              <el-table-column fixed prop="name" label="Name" width="130px">
              </el-table-column>
              <el-table-column
                prop="first_name"
                label="FirstName"
                width="130px"
              >
              </el-table-column>
              <el-table-column prop="last_name" label="LastName" width="130px">
              </el-table-column>
              <el-table-column
                prop="start_time"
                label="StartTime"
                width="130px"
              >
              </el-table-column>
              <el-table-column prop="end_time" label="EndTime" width="130px">
              </el-table-column>
              <el-table-column prop="position" label="Position">
              </el-table-column>
              <el-table-column label="Picture">
                <template slot-scope="scope">
                  <img
                    :src="scope.row.header_img"
                    style="width: 120px; height: auto"
                  />
                </template>
              </el-table-column>
            </el-table></div
        ></el-col>
        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">
            Vice Presidents:
          </div></el-col
        >
        <el-col :span="24"
          ><div class="grid-content bg-purple-dark value">
            <el-table
              :data="journalDetail.vice_presidents"
              border
              style="width: 100%"
              stripe
              :header-cell-style="{
                color: '#000000',
                background: '#F7FBFF',
              }"
            >
              <el-table-column fixed prop="name" label="Name" width="130px">
              </el-table-column>
              <el-table-column
                prop="first_name"
                label="FirstName"
                width="130px"
              >
              </el-table-column>
              <el-table-column prop="last_name" label="LastName" width="130px">
              </el-table-column>
              <el-table-column
                prop="start_time"
                label="StartTime"
                width="130px"
              >
              </el-table-column>
              <el-table-column prop="end_time" label="EndTime" width="130px">
              </el-table-column>
              <el-table-column prop="position" label="Position">
              </el-table-column>
              <el-table-column label="Picture">
                <template slot-scope="scope">
                  <img
                    :src="scope.row.header_img"
                    style="width: 120px; height: auto"
                  />
                </template>
              </el-table-column>
            </el-table></div
        ></el-col>
        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">Members:</div></el-col
        >
        <el-col :span="24"
          ><div class="grid-content bg-purple-dark value">
            <el-table
              :data="journalDetail.members"
              border
              style="width: 100%"
              stripe
              :header-cell-style="{
                color: '#000000',
                background: '#F7FBFF',
              }"
            >
              <el-table-column fixed prop="name" label="Name" width="130px">
              </el-table-column>
              <el-table-column
                prop="first_name"
                label="FirstName"
                width="130px"
              >
              </el-table-column>
              <el-table-column prop="last_name" label="LastName" width="130px">
              </el-table-column>
              <el-table-column
                prop="start_time"
                label="StartTime"
                width="130px"
              >
              </el-table-column>
              <el-table-column prop="end_time" label="EndTime" width="130px">
              </el-table-column>
              <el-table-column prop="position" label="Position">
              </el-table-column>
              <el-table-column label="Picture">
                <template slot-scope="scope">
                  <img
                    :src="scope.row.header_img"
                    style="width: 120px; height: auto"
                  />
                </template>
              </el-table-column>
            </el-table></div
        ></el-col>
      </el-row>
    </div>
  </div>
</template>
<script>
import { getJournalDetail, getJournalIssues } from "../../api";
export default {
  data() {
    return {
      journalDetail: {
        name: "",
        description: "",
        category: "",
        create_at: "",
        presidents: [
          {
            name: "",
            first_name: "",
            last_name: "",
            header_img: "",
            position: "",
            start_time: "",
            end_time: "",
          },
        ],
        vice_presidents: [
          {
            name: "",
            first_name: "",
            last_name: "",
            header_img: "",
            position: "",
            start_time: "",
            end_time: "",
          },
        ],
        members: [
          {
            name: "",
            first_name: "",
            last_name: "",
            header_img: "",
            position: "",
            start_time: "",
            end_time: "",
          },
        ],
      },
      showAllIssueButton: false,
      showSubmitButton: false,
    };
  },
  methods: {
    Submit() {
      this.$router.push({
        path: "/journalSubmit",
        query: {
          id: this.$route.query.journal_id,
        },
      });
    },
    Issues() {
      this.$router.push({
        path: "/journalIssues",
        query: {
          id: this.$route.query.journal_id,
        },
      });
    },
    getJournal() {
      // console.log(this.$route.query_id);
      getJournalDetail({
        journal_id: this.$route.query.journal_id,
      }).then((res) => {
        console.log(res.data.data);
        this.journalDetail = res.data.data;
        // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
        this.journalDetail.create_at = this.journalDetail.create_at.substring(
          0,
          10
        );
        for (let i = 0; i < this.journalDetail.presidents.length; i++) {
          this.journalDetail.presidents[i].start_time =
            this.journalDetail.presidents[i].start_time.substring(0, 10);
          this.journalDetail.presidents[i].end_time =
            this.journalDetail.presidents[i].end_time.substring(0, 10);
        }
        for (let i = 0; i < this.journalDetail.vice_presidents.length; i++) {
          this.journalDetail.vice_presidents[i].start_time =
            this.journalDetail.vice_presidents[i].start_time.substring(0, 10);
          this.journalDetail.vice_presidents[i].end_time =
            this.journalDetail.vice_presidents[i].end_time.substring(0, 10);
        }
        for (let i = 0; i < this.journalDetail.members.length; i++) {
          this.journalDetail.members[i].start_time = this.journalDetail.members[
            i
          ].start_time.substring(0, 10);
          this.journalDetail.members[i].end_time = this.journalDetail.members[
            i
          ].end_time.substring(0, 10);
        }
        console.log(this.journalDetail);
      });
      getJournalIssues({ journal_id: this.$route.query.journal_id }).then(
        (res) => {
          if (res.data.data.length !== 0) {
            this.showAllIssueButton = true;
            // 如果最近的一期的submission_end_time大于当前时间，就显示submit按钮
            // 遍历找出最近的一期的submission_end_time
            let current_time = new Date();
            // 定义一个时间最小值
            let latest_time = new Date("1970-01-01");
            console.log(res.data.data, "res.data.data");
            for (let i = 0; i < res.data.data.length; i++) {
              if (
                new Date(res.data.data[i].submission_end_time) > latest_time
              ) {
                latest_time = new Date(res.data.data[i].submission_end_time);
              }
            }
            console.log(
              latest_time,
              current_time,
              "latest_time",
              "current_time"
            );
            if (latest_time > current_time) {
              this.showSubmitButton = true;
            }
          }
        }
      );
    },
  },
  mounted() {
    this.getJournal();
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
    // height: 950px;
    // background-color: #e2f1fb;
    background-color: #ecf5ff;
    // 圆角
    border-radius: 10px;
    // 水平居中
    margin: 0 auto;
    margin-top: 20px;
    // 上下左右panding
    padding: 20px 20px;
    // color: #072e5b;
    .el-row {
      margin-bottom: 20px;
      &:last-child {
        margin-bottom: 0;
      }
    }
    .el-col {
      border-radius: 4px;
    }
    // .bg-purple-dark {
    //   background: #99a9bf;
    // }
    .grid-content {
      border-radius: 4px;
      min-height: 36px;
    }
    .name {
      font-size: 25px;
      font-weight: bold;
      text-align: center;
    }
    .author {
      // 字体加粗
      font-weight: bold;
      font-size: 10px;
    }
    .key {
      // 文字靠左
      text-align: left;
      // 字体加粗
      font-weight: bold;
    }
    .value {
      // 文字靠左
      text-align: left;
      // 盒子默认最小高度
      min-height: 40px;
    }
  }
}
</style>
