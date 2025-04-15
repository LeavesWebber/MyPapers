<template>
  <div class="box">
    <div class="box1">
      <el-row>
        <el-col :span="2"
          ><div class="grid-content bg-purple-dark name">
            <el-button
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
              size="small"
              type="primary"
              @click="Submit()"
              >Submit Manuscript</el-button
            >
          </div></el-col
        >
        <el-col :span="24"
          ><div class="grid-content bg-purple-dark name">
            <span>{{ conferenceDetail.name }}</span>
          </div></el-col
        >

        <el-col :span="22"
          ><div class="grid-content bg-purple-dark author">
            {{ conferenceDetail.authors }}
          </div></el-col
        >
        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">
            Description:
          </div></el-col
        >
        <el-col :span="20"
          ><div class="grid-content bg-purple-dark value">
            {{ conferenceDetail.description }}
          </div></el-col
        >
        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">Location:</div></el-col
        >
        <el-col :span="20"
          ><div class="grid-content bg-purple-dark value">
            {{ conferenceDetail.location }}
          </div></el-col
        >
        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">StartTime:</div></el-col
        >
        <el-col :span="20"
          ><div class="grid-content bg-purple-dark value">
            {{ conferenceDetail.start_time }}
          </div></el-col
        >
        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">EndTime:</div></el-col
        >
        <el-col :span="20"
          ><div class="grid-content bg-purple-dark value">
            {{ conferenceDetail.end_time }}
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
              :data="conferenceDetail.presidents"
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
              :data="conferenceDetail.vice_presidents"
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
              :data="conferenceDetail.members"
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
import { getConferenceDetail, getConferenceIssues } from "../../api";
export default {
  data() {
    return {
      conferenceDetail: {
        name: "",
        description: "",
        location: "",
        start_time: "",
        end_time: "",
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
        path: "/conferenceSubmit",
        query: {
          id: this.$route.query.conference_id,
        },
      });
    },
    Issues() {
      this.$router.push({
        path: "/conferenceIssues",
        query: {
          id: this.$route.query.conference_id,
        },
      });
    },
    getConference() {
      // console.log(this.$route.query_id);
      getConferenceDetail({
        conference_id: this.$route.query.conference_id,
      }).then((res) => {
        console.log(res.data.data);
        this.conferenceDetail = res.data.data;
        // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
        this.conferenceDetail.start_time =
          this.conferenceDetail.start_time.substring(0, 10);
        this.conferenceDetail.end_time =
          this.conferenceDetail.end_time.substring(0, 10);
        for (let i = 0; i < this.conferenceDetail.presidents.length; i++) {
          this.conferenceDetail.presidents[i].start_time =
            this.conferenceDetail.presidents[i].start_time.substring(0, 10);
          this.conferenceDetail.presidents[i].end_time =
            this.conferenceDetail.presidents[i].end_time.substring(0, 10);
        }
        for (let i = 0; i < this.conferenceDetail.vice_presidents.length; i++) {
          this.conferenceDetail.vice_presidents[i].start_time =
            this.conferenceDetail.vice_presidents[i].start_time.substring(
              0,
              10
            );
          this.conferenceDetail.vice_presidents[i].end_time =
            this.conferenceDetail.vice_presidents[i].end_time.substring(0, 10);
        }
        for (let i = 0; i < this.conferenceDetail.members.length; i++) {
          this.conferenceDetail.members[i].start_time =
            this.conferenceDetail.members[i].start_time.substring(0, 10);
          this.conferenceDetail.members[i].end_time =
            this.conferenceDetail.members[i].end_time.substring(0, 10);
        }
      });
      getConferenceIssues({
        conference_id: this.$route.query.conference_id,
      }).then((res) => {
        if (res.data.data.length !== 0) {
          this.showAllIssueButton = true;
          // 如果最近的一期的submission_end_time大于当前时间，就显示submit按钮
          // 遍历找出最近的一期的submission_end_time
          let current_time = new Date();
          // 定义一个时间最小值
          let latest_time = new Date("1970-01-01");
          console.log(res.data.data, "res.data.data");
          for (let i = 0; i < res.data.data.length; i++) {
            if (new Date(res.data.data[i].submission_end_time) > latest_time) {
              latest_time = new Date(res.data.data[i].submission_end_time);
            }
          }
          console.log(latest_time, current_time, "latest_time", "current_time");
          if (latest_time > current_time) {
            this.showSubmitButton = true;
          }
        }
      });
    },
  },
  mounted() {
    this.getConference();
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
