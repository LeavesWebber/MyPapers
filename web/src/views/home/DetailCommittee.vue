<template>
  <div class="box">
    <div class="box1">
      <el-row>
        <el-col :span="22"
          ><div class="grid-content bg-purple-dark name">
            <span>{{ committeeDetail.name }}</span>
          </div></el-col
        >
        <el-col :span="22"
          ><div class="grid-content bg-purple-dark author">
            {{ committeeDetail.authors }}
          </div></el-col
        >
        <el-col :span="4"
          ><div class="grid-content bg-purple-dark key">
            Description:
          </div></el-col
        >
        <el-col :span="20"
          ><div class="grid-content bg-purple-dark value">
            {{ committeeDetail.description }}
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
              :data="committeeDetail.presidents"
              border
              stripe
              style="width: 100%"
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
              :data="committeeDetail.vice_presidents"
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
              :data="committeeDetail.members"
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
      <el-row>
        <h1 style="text-align: center">Conferences</h1>
        <div class="query">
          <span> Query: </span>
          <el-input
            style="width: 260px"
            v-model="searchConferenceKeyword"
            @input="searchConferenceData"
            placeholder="Enter name keywords to query"
          ></el-input>
        </div>
        <el-table
          :empty-text="emptyConferencesText"
          :data="displayedConferencesData"
          height="400"
          border
          stripe
          style="width: 100%"
          :header-cell-style="{ color: '#000000', background: '#F7FBFF' }"
        >
          <el-table-column fixed prop="ID" label="ID" width="80" sortable>
          </el-table-column>
          <el-table-column prop="name" label="Name" width="350">
          </el-table-column>
          <el-table-column prop="location" label="Location" width="200">
          </el-table-column>
          <el-table-column prop="start_time" label="Start Time" width="200">
          </el-table-column>
          <el-table-column prop="end_time" label="End Time"> </el-table-column>
          <el-table-column fixed="right" label="" width="100">
            <template slot-scope="scope">
              <el-button @click="details(scope.row)" type="text" size="small"
                >Details</el-button
              >
            </template>
          </el-table-column>
        </el-table>
        <div class="block">
          <el-pagination
            @current-change="changeConferencePage"
            :current-page="currentConferencePage"
            :page-size="pageConferenceSize"
            :hide-on-single-page="true"
            layout="prev, pager, next"
            :total="conferencesCount"
          >
          </el-pagination>
        </div>
      </el-row>
      <el-row>
        <h1 style="text-align: center">Journals</h1>
        <div class="query">
          <span> Query: </span>
          <el-input
            style="width: 260px"
            v-model="searchJournalKeyword"
            @input="searchJournalData"
            placeholder="Enter name keywords to query"
          ></el-input>
        </div>
        <el-table
          :empty-text="emptyJournalsText"
          :data="displayedJournalsData"
          height="400"
          border
          stripe
          style="width: 100%"
          :header-cell-style="{ color: '#000000', background: '#F7FBFF' }"
        >
          <el-table-column fixed prop="ID" label="ID" width="80" sortable>
          </el-table-column>
          <el-table-column prop="name" label="Name" width="350">
          </el-table-column>
          <el-table-column prop="CreatedAt" label="Created Time" width="200">
          </el-table-column>
          <el-table-column prop="category" label="Category"> </el-table-column>
          <el-table-column fixed="right" label="" width="100">
            <template slot-scope="scope">
              <el-button @click="details(scope.row)" type="text" size="small"
                >Details</el-button
              >
            </template>
          </el-table-column>
        </el-table>
        <div class="block">
          <el-pagination
            @current-change="changeJournalPage"
            :current-page="currentJournalPage"
            :page-size="pageJournalSize"
            :hide-on-single-page="true"
            layout="prev, pager, next"
            :total="journalsCount"
          >
          </el-pagination>
        </div>
      </el-row>
    </div>
  </div>
</template>
<script>
import {
  getCommitteeDetail,
  getConferencelistByCommittee,
  getJournallistByCommittee,
} from "../../api";
export default {
  data() {
    return {
      committeeDetail: {
        name: "",
        description: "",
        presidents: [
          {
            name: "",
            first_name: "",
            last_name: "",
            header_img: "",
            position: "",
            start_time: "",
            end_time: "",
            level: "",
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
            level: "",
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
            level: "",
          },
        ],
      },
      conferences: [
        {
          ID: 0,
          name: "",
          location: "",
          start_time: "",
          end_time: "",
        },
      ],
      currentConferencePage: 1,
      pageConferenceSize: 10,
      searchConferenceKeyword: "",
      emptyConferencesText: "No Conferences",
      journals: [
        {
          ID: 0,
          name: "",
          CreatedAt: "",
          category: "",
        },
      ],
      currentJournalPage: 1,
      pageJournalSize: 10,
      searchJournalKeyword: "",
      emptyJournalsText: "No Journals",
    };
  },
  methods: {
    details(row) {
      this.$router.push({
        path: "detailConference",
        query: {
          conference_id: row.ID,
        },
      });
    },
    searchConferenceData() {
      // 当搜索关键字发生变化时，重置当前页码为第一页
      this.currentPage = 1;
    },
    changeConferencePage(currentPage) {
      this.currentPage = currentPage;
    },
    searchJournalData() {
      // 当搜索关键字发生变化时，重置当前页码为第一页
      this.currentPage = 1;
    },
    changeJournalPage(currentPage) {
      this.currentPage = currentPage;
    },
  },
  computed: {
    conferencesCount() {
      return this.conferences.length;
    },
    displayedConferencesData() {
      // 进行查询过滤
      const filteredData = this.conferences.filter((item) =>
        item.name
          .toLowerCase()
          .includes(this.searchConferenceKeyword.toLowerCase())
      );
      // 进行分页处理
      const startIndex =
        (this.currentConferencePage - 1) * this.pageConferenceSize;

      const endIndex = startIndex + this.pageConferenceSize;
      if (endIndex > filteredData.length) {
        return filteredData.slice(startIndex, filteredData.length);
      }
      return filteredData.slice(startIndex, endIndex);
    },
    journalsCount() {
      return this.journals.length;
    },
    displayedJournalsData() {
      // 进行查询过滤
      const filteredData = this.journals.filter((item) =>
        item.name
          .toLowerCase()
          .includes(this.searchJournalKeyword.toLowerCase())
      );
      // 进行分页处理
      const startIndex = (this.currentJournalPage - 1) * this.pageJournalSize;

      const endIndex = startIndex + this.pageJournalSize;
      if (endIndex > filteredData.length) {
        return filteredData.slice(startIndex, filteredData.length);
      }
      return filteredData.slice(startIndex, endIndex);
    },
  },
  mounted() {
    getCommitteeDetail({ committee_id: this.$route.query.committee_id }).then(
      (res) => {
        this.committeeDetail = res.data.data;
        // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
        for (let i = 0; i < this.committeeDetail.presidents.length; i++) {
          this.committeeDetail.presidents[i].start_time =
            this.committeeDetail.presidents[i].start_time.substring(0, 10);
          this.committeeDetail.presidents[i].end_time =
            this.committeeDetail.presidents[i].end_time.substring(0, 10);
        }
        for (let i = 0; i < this.committeeDetail.vice_presidents.length; i++) {
          this.committeeDetail.vice_presidents[i].start_time =
            this.committeeDetail.vice_presidents[i].start_time.substring(0, 10);
          this.committeeDetail.vice_presidents[i].end_time =
            this.committeeDetail.vice_presidents[i].end_time.substring(0, 10);
        }
        for (let i = 0; i < this.committeeDetail.members.length; i++) {
          this.committeeDetail.members[i].start_time =
            this.committeeDetail.members[i].start_time.substring(0, 10);
          this.committeeDetail.members[i].end_time =
            this.committeeDetail.members[i].end_time.substring(0, 10);
        }
      },
      getConferencelistByCommittee({
        committee_id: this.$route.query.committee_id,
      }).then((res) => {
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
        this.conferences = res.data.data;
      }),
      getJournallistByCommittee({
        committee_id: this.$route.query.committee_id,
      }).then((res) => {
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
        }
        this.journals = res.data.data;
      })
    );
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
.query {
  margin-bottom: 10px;
  text-align: right;
}
.block {
  text-align: center;
}
</style>
