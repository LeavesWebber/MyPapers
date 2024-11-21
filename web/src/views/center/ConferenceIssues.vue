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
    <el-dialog
      title="Issue"
      :visible.sync="issueDialogVisible"
      width="60%"
      @close="cancel"
    >
      <!-- Issue的表单信息 -->
      <el-form
        ref="formData"
        :inline="true"
        :model="formData"
        label-width="95px"
        :rules="rules"
      >
        <el-form-item label="Name" prop="name">
          <el-input style="width: 220px" v-model="formData.name"></el-input>
        </el-form-item>
        <el-form-item label="Volume" prop="volume">
          <el-input style="width: 220px" v-model="formData.volume"></el-input>
        </el-form-item>
        <el-form-item label="Start Time" prop="submission_start_time">
          <el-date-picker
            v-model="formData.submission_start_time"
            type="date"
            :picker-options="dateTimeStartFunc"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item label="End Time" prop="submission_end_time">
          <el-date-picker
            v-model="formData.submission_end_time"
            type="date"
            :picker-options="dateTimeEndFunc"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item label="Description" prop="description">
          <el-input
            style="width: 540px"
            v-model="formData.description"
          ></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">Cancel</el-button>
        <el-button type="primary" @click="submit('formData')">OK </el-button>
      </span>
    </el-dialog>

    <div class="button">
      <el-button
        size="small"
        type="primary"
        @click="popupAddIssue"
        v-if="showButton"
        >Add Issue</el-button
      >
    </div>
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
      :data="displayedData"
      height="550"
      stripe
      border
      :header-cell-style="{
        color: '#000000',
        background: '#F7FBFF',
      }"
      :default-sort="{ prop: 'volume', order: 'descending' }"
    >
      <el-table-column fixed prop="volume" label="Volume" sortable width="100">
      </el-table-column>
      <el-table-column prop="name" label="name" width="120"> </el-table-column>
      <el-table-column prop="year" label="Year" width="120"> </el-table-column>
      <el-table-column
        prop="submission_start_time"
        label="Start Time"
        width="140"
      >
      </el-table-column>
      <el-table-column prop="submission_end_time" label="End Time" width="140">
      </el-table-column>
      <el-table-column prop="description" label="Description" width="330">
      </el-table-column>
      <el-table-column fixed="right" label="Options">
        <template slot-scope="scope">
          <el-button
            @click="handleUpdate(scope.row)"
            type="primary"
            size="mini"
            v-if="showButton"
            >Update</el-button
          >
          <el-button
            @click="popup(scope.row)"
            type="danger"
            size="mini"
            v-if="showButton"
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
import {
  getConferenceIssues,
  getConferenceLevel,
  createConferenceIssue,
  updateConferenceIssue,
  deleteConferenceIssue,
} from "../../api";
export default {
  data() {
    var validateWordCount = (rule, value, callback) => {
      const wordCount = value.trim().split(/\s+/).length;
      if (wordCount < 1 || wordCount > 20) {
        callback(new Error("Should contain 1 to 10 words"));
      } else {
        callback();
      }
    };
    return {
      formData: {
        ID: 0,
        conference_id: "",
        name: "",
        description: "",
        volume: "",
        submission_start_time: "",
        submission_end_time: "",
      },

      tableData: [
        {
          ID: 0,
          CreatedAt: "",
          conference_id: "",
          name: "",
          submission_start_time: "",
          submission_end_time: "",
          description: "",
          year: "",
          volume: 0,
        },
      ],
      rules: {
        name: [
          { required: true, message: "Please input name", trigger: "blur" },
          { validator: validateWordCount, trigger: "blur" },
        ],
        volume: [
          { required: true, message: "Please input volume", trigger: "blur" },
          // 限制输入的是数字
          {
            pattern: /^[0-9]*$/,
            message: "Please input number",
            trigger: "blur",
          },
        ],
        submission_start_time: [
          {
            required: true,
            message: "Please input submission start time",
            trigger: "blur",
          },
        ],
        submission_end_time: [
          {
            required: true,
            message: "Please input submission end time",
            trigger: "blur",
          },
        ],
      },
      modalType: 0, // 0表示新增弹窗，1表示编辑
      selectRow: null,
      showButton: false,
      userInfo: {},
      centerDialogVisible: false,
      issueDialogVisible: false,
      currentPage: 1,
      pageSize: 7,
      searchKeyword: "",
    };
  },
  methods: {
    updateIssue(row) {
      console.log(row.ID);
      this.$router.push({
        path: "/center/detailConference",
        query: {
          conference_id: row.ID,
        },
      });
    },
    popup(row) {
      this.centerDialogVisible = true;
      this.selectRow = row;
    },
    popupIssue(row) {
      this.issueDialogVisible = true;
      this.selectRow = row;
    },
    popupAddIssue() {
      this.modalType = 0;
      this.issueDialogVisible = true;
    },
    cancel() {
      // 重置表单
      this.formData = {
        ID: 0,
        conference_id: "",
        name: "",
        description: "",
        volume: "",
        submission_start_time: "",
        submission_end_time: "",
      };
      this.issueDialogVisible = false;
    },
    handleDelete(row) {
      console.log(row.ID);
      deleteConferenceIssue({ issue_id: row.ID }).then((res) => {
        console.log(res);
        // 先弹窗提示是否删除
        if (res.data.code === 1000) {
          this.getIssues();
          this.$message({
            message: "Delete Success",
            type: "success",
          });
        } else {
          this.$message({
            message: "Delete Failed",
            type: "error",
          });
        }
        this.centerDialogVisible = false;
        // 刷新页面
        this.getIssues();
      });
    },
    getIssues() {
      getConferenceIssues({
        conference_id: this.$route.query.conference_id,
      }).then((res) => {
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
      });
      // 获取用户在期刊的level
      getConferenceLevel({
        conference_id: this.$route.query.conference_id,
      }).then((res) => {
        console.log(res.data.data, "res.data.data.level");
        // res.data.data 为president,vice_president时，显示按钮
        if (
          res.data.data === "president" ||
          res.data.data === "vice_president"
        ) {
          this.showButton = true;
        }
      });
    },
    // 提交表单
    submit(formData) {
      this.$refs[formData].validate((valid) => {
        if (valid) {
          this.formData.conference_id = this.$route.query.conference_id;
          // conference_id转为数字类型
          this.formData.conference_id = parseInt(this.formData.conference_id);
          this.formData.volume = parseInt(this.formData.volume);
          // 后续对表单数据的处理
          if (this.modalType === 0) {
            console.log(this.formData, "before formData");
            createConferenceIssue(this.formData).then(() => {
              // 重新获取列表的接口
              this.getIssues();
            });
          } else {
            // 如果时间格式为2022-08-24，要转为2022-08-03T15:04:05Z格式
            if (this.formData.submission_start_time.length === 10) {
              this.formData.submission_start_time =
                this.formData.submission_start_time + "T00:00:00Z";
            }
            if (this.formData.submission_end_time.length === 10) {
              this.formData.submission_end_time =
                this.formData.submission_end_time + "T00:00:00Z";
            }
            console.log(this.formData, "before formData");
            updateConferenceIssue(this.formData).then(() => {
              // 重新获取列表的接口
              this.getIssues();
            });
          }
          // 重置表单
          this.formData = {
            ID: 0,
            conference_id: "",
            name: "",
            description: "",
            volume: "",
            submission_start_time: "",
            submission_end_time: "",
          };
          //   this.$refs.formData.resetFields();
          console.log(this.formData, "formData");
          // 关闭弹窗
          this.issueDialogVisible = false;
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    // 修改
    handleUpdate(row) {
      this.modalType = 1;
      this.issueDialogVisible = true;
      //等弹窗渲染完后（form初始化为空的）再赋值
      console.log(row, "row");
      this.$nextTick(() => {
        // 注意需要对当前行数据进行深拷贝，否则会出错
        this.formData = JSON.parse(JSON.stringify(row));
      });

      // this.form = row;
      console.log(this.formData, "handleUpdate");
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
    dateTimeStartFunc() {
      return {
        disabledDate: (time) => {
          if (this.formData.submission_end_time) {
            return (
              time.getTime() >
              new Date(this.formData.submission_end_time).getTime() - 8.64e7
            );
          }
        },
      };
    },
    dateTimeEndFunc() {
      return {
        disabledDate: (time) => {
          if (this.formData.submission_start_time) {
            return (
              time.getTime() <
              new Date(this.formData.submission_start_time).getTime() + 8.64e7
            );
          }
        },
      };
    },
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
    this.getIssues();
  },
};
</script>
<style lang="less" scoped>
.button {
  // 靠右
  float: right;
}
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
  margin-right: 20px;
  // 靠右
  float: right;
}
</style>