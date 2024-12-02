<template>
  <div class="box">
    <div class="box1">
      <h1 style="text-align: center">Journal Papers</h1>
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
        :empty-text="emptyText"
        height="550"
        border
        stripe
        style="width: 100%"
        :header-cell-style="{ color: '#000000', background: '#F7FBFF' }"
        :default-sort="{ prop: 'CreatedAt', order: 'descending' }"
      >
        <el-table-column fixed prop="title" label="Title"> </el-table-column>
        <el-table-column prop="block_address" label="Block Address">
        </el-table-column>
        <el-table-column prop="authors" label="Author"> </el-table-column>
        <el-table-column
          prop="CreatedAt"
          label="Publication Time"
          width="165px"
          sortable
        >
        </el-table-column>
        <!-- <el-table-column prop="paper_type" label="Type"> </el-table-column> -->
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
      <h1 style="text-align: center">Conference Papers</h1>
      <div class="query">
        <span> Query: </span>
        <el-input
          style="width: 260px"
          v-model="searchKeyword2"
          @input="searchData2"
          placeholder="Enter title keywords to query"
        ></el-input>
      </div>
      <el-table
        :data="displayedData2"
        height="550"
        border
        stripe
        style="width: 100%"
        :header-cell-style="{ color: '#000000', background: '#F7FBFF' }"
        :default-sort="{ prop: 'CreatedAt', order: 'descending' }"
      >
        <el-table-column fixed prop="title" label="Title"> </el-table-column>
        <el-table-column prop="block_address" label="Block Address">
        </el-table-column>
        <el-table-column prop="authors" label="Author"> </el-table-column>
        <el-table-column
          prop="CreatedAt"
          label="Publication Time"
          width="165px"
          sortable
        >
        </el-table-column>
        <!-- <el-table-column prop="paper_type" label="Type"> </el-table-column> -->
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
      <div class="block">
        <el-pagination
          @current-change="changePage2"
          :current-page="currentPage2"
          :page-size="pageSize2"
          :hide-on-single-page="true"
          layout="prev, pager, next"
          :total="tableDataCount2"
        >
        </el-pagination>
      </div>
    </div>
  </div>
</template>
<script>
import { getAllAcceptPapers } from "../../api";
export default {
  data() {
    return {
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
      conferencePapers: [
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
      currentPage: 1,
      pageSize: 5,
      searchKeyword: "",
      currentPage2: 1,
      pageSize2: 5,
      searchKeyword2: "",
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

    getPapers() {
      getAllAcceptPapers().then((res) => {
        this.journalPapers = [];
        this.conferencePapers = [];
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
        // 根据conference_id和journal_id判断是会议论文还是期刊论文
        for (let i = 0; i < res.data.data.length; i++) {
          if (res.data.data[i].conference_id === 0) {
            this.journalPapers.push(res.data.data[i]);
          } else {
            this.conferencePapers.push(res.data.data[i]);
          }
        }
        // 把journalPapers里面的user数组里面的first_name和last_name拼接起来
        // for (let i = 0; i < this.journalPapers.length; i++) {
        //   for (let j = 0; j < this.journalPapers[i].user.length; j++) {
        //     // 如果authors没有初始化，就初始化为空
        //     if (this.journalPapers[i].authors === undefined) {
        //       this.journalPapers[i].authors = "";
        //     }
        //     // 如果是最后一个作者，就不加逗号
        //     if (j === this.journalPapers[i].user.length - 1) {
        //       this.journalPapers[i].authors +=
        //         this.journalPapers[i].user[j].first_name +
        //         " " +
        //         this.journalPapers[i].user[j].last_name;
        //     } else {
        //       this.journalPapers[i].authors +=
        //         this.journalPapers[i].user[j].first_name +
        //         " " +
        //         this.journalPapers[i].user[j].last_name +
        //         ",";
        //     }
        //   }
        // }
        // 把conferencePapers里面的user数组里面的first_name和last_name拼接起来
        // for (let i = 0; i < this.conferencePapers.length; i++) {
        //   for (let j = 0; j < this.conferencePapers[i].user.length; j++) {
        //     // 如果authors没有初始化，就初始化为空
        //     if (this.conferencePapers[i].authors === undefined) {
        //       this.conferencePapers[i].authors = "";
        //     }
        //     // 如果是最后一个作者，就不加逗号
        //     if (j === this.conferencePapers[i].user.length - 1) {
        //       this.conferencePapers[i].authors +=
        //         this.conferencePapers[i].user[j].first_name +
        //         " " +
        //         this.conferencePapers[i].user[j].last_name;
        //     } else {
        //       this.conferencePapers[i].authors +=
        //         this.conferencePapers[i].user[j].first_name +
        //         " " +
        //         this.conferencePapers[i].user[j].last_name +
        //         ",";
        //     }
        //   }
        // }
        console.log(this.journalPapers, "journalPapers");
        console.log(this.conferencePapers, "conferencePapers");
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
    searchData2() {
      // 当搜索关键字发生变化时，重置当前页码为第一页
      this.currentPage2 = 1;
    },
    changePage2(currentPage2) {
      this.currentPage2 = currentPage2;
      console.log(currentPage2, this.currentPage2, "currentPage2");
    },
  },
  computed: {
    tableDataCount() {
      return this.journalPapers.length;
    },
    displayedData() {
      // 进行查询过滤
      const filteredData = this.journalPapers.filter((item) =>
        item.title.toLowerCase().includes(this.searchKeyword.toLowerCase())
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
    tableDataCount2() {
      return this.conferencePapers.length;
    },
    displayedData2() {
      // 进行查询过滤
      const filteredData = this.conferencePapers.filter((item) =>
        item.title.toLowerCase().includes(this.searchKeyword2.toLowerCase())
      );
      // 进行分页处理
      const startIndex = (this.currentPage2 - 1) * this.pageSize2;

      const endIndex = startIndex + this.pageSize2;
      console.log(startIndex, endIndex, "startIndex, endIndex");
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
.query {
  margin-bottom: 10px;
  text-align: right;
}
.block {
  text-align: center;
}
</style>
