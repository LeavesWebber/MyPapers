<template>
  <div>
    <el-row>
      <el-col :span="2"
        ><div class="grid-content bg-purple-dark name">
          <el-button size="small" type="primary" @click="issues()"
            >Issues</el-button
          >
        </div></el-col
      >
      <el-col :span="2"
        ><div class="grid-content bg-purple-dark name">
          <el-button
            size="small"
            type="primary"
            @click="updateConference()"
            v-if="showUpdateButton"
            >Update</el-button
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
        ><div class="grid-content bg-purple-dark key">Description:</div></el-col
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
        ><div class="grid-content bg-purple-dark key">Category:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ conferenceDetail.category }}
        </div></el-col
      >

      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Start Time:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ conferenceDetail.start_time }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">End Time:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ conferenceDetail.end_time }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Presidents:</div></el-col
      >
      <el-col :span="24"
        ><div class="grid-content bg-purple-dark value">
          <el-table
            :data="conferenceDetail.presidents"
            border
            stripe
            :header-cell-style="{
              color: '#000000',
              background: '#F7FBFF',
            }"
          >
            <el-table-column fixed prop="name" label="Name" width="130px">
            </el-table-column>
            <el-table-column prop="first_name" label="FirstName" width="130px">
            </el-table-column>
            <el-table-column prop="last_name" label="LastName" width="130px">
            </el-table-column>
            <el-table-column prop="start_time" label="StartTime" width="130px">
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
            stripe
            :header-cell-style="{
              color: '#000000',
              background: '#F7FBFF',
            }"
          >
            <el-table-column fixed prop="name" label="Name" width="130px">
            </el-table-column>
            <el-table-column prop="first_name" label="FirstName" width="130px">
            </el-table-column>
            <el-table-column prop="last_name" label="LastName" width="130px">
            </el-table-column>
            <el-table-column prop="start_time" label="StartTime" width="130px">
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
            stripe
            :header-cell-style="{
              color: '#000000',
              background: '#F7FBFF',
            }"
          >
            <el-table-column fixed prop="name" label="Name" width="130px">
            </el-table-column>
            <el-table-column prop="first_name" label="FirstName" width="130px">
            </el-table-column>
            <el-table-column prop="last_name" label="LastName" width="130px">
            </el-table-column>
            <el-table-column prop="start_time" label="StartTime" width="130px">
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
</template>
<script>
import { getConferenceDetail } from "../../api";
export default {
  data() {
    return {
      conferenceDetail: {
        creator_id: 0,
        name: "",
        description: "",
        category: "",
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
      showUpdateButton: false,
      userInfo: {},
    };
  },
  methods: {
    issues() {
      this.$router.push({
        path: "/center/conferenceIssues",
        query: {
          conference_id: this.$route.query.conference_id,
        },
      });
    },
    updateConference() {
      this.$router.push({
        path: "/center/updateConference",
        query: {
          conference_id: this.$route.query.conference_id,
        },
      });
    },
  },
  mounted() {
    // console.log(this.$route.query_id);
    getConferenceDetail({
      conference_id: this.$route.query.conference_id,
    }).then((res) => {
      console.log(res.data.data);
      this.conferenceDetail = res.data.data;
      // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
      this.conferenceDetail.start_time =
        this.conferenceDetail.start_time.substring(0, 10);
      this.conferenceDetail.end_time = this.conferenceDetail.end_time.substring(
        0,
        10
      );
      for (let i = 0; i < this.conferenceDetail.presidents.length; i++) {
        this.conferenceDetail.presidents[i].start_time =
          this.conferenceDetail.presidents[i].start_time.substring(0, 10);
        this.conferenceDetail.presidents[i].end_time =
          this.conferenceDetail.presidents[i].end_time.substring(0, 10);
      }
      for (let i = 0; i < this.conferenceDetail.vice_presidents.length; i++) {
        this.conferenceDetail.vice_presidents[i].start_time =
          this.conferenceDetail.vice_presidents[i].start_time.substring(0, 10);
        this.conferenceDetail.vice_presidents[i].end_time =
          this.conferenceDetail.vice_presidents[i].end_time.substring(0, 10);
      }
      for (let i = 0; i < this.conferenceDetail.members.length; i++) {
        this.conferenceDetail.members[i].start_time =
          this.conferenceDetail.members[i].start_time.substring(0, 10);
        this.conferenceDetail.members[i].end_time =
          this.conferenceDetail.members[i].end_time.substring(0, 10);
      }
      // 从localStorage里面获取用户信息
      this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
      console.log(this.userInfo, "userInfo");
      // 遍历presidents里面是否有自己的名字
      for (let i = 0; i < this.conferenceDetail.presidents.length; i++) {
        if (
          this.conferenceDetail.presidents[i].name === this.userInfo.username ||
          this.conferenceDetail.creator_id === this.userInfo.ID
        ) {
          this.showUpdateButton = true;
        }
      }
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
</style>