<!-- 作用：评审论文的详细信息页面
功能：
显示论文详情
提交评审意见
查看评审历史
PDF预览相关：包含论文PDF文件的预览链接(通过paperDetail.filepath访问) -->
<template>
  <div>
    <el-dialog :visible.sync="dialogVisible" width="60%" append-to-body>
      <el-form ref="reviewForm" :model="reviewForm" label-width="160px">
        <el-form-item label="ReviewStatus">
          <el-radio-group v-model="reviewForm.status">
            <el-radio :label="1">Accept</el-radio>
            <el-radio :label="2">Minor Revisions</el-radio>
            <el-radio :label="3">Major Revisions</el-radio>
            <el-radio :label="4">Reject and Resubmit</el-radio>
            <el-radio :label="5">Reject</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="ReviewComments">
          <el-input
            type="textarea"
            :rows="5"
            placeholder="Please input"
            v-model="reviewForm.comment"
          ></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="reviewSubmit">Review</el-button>
      </span>
    </el-dialog>

    <el-row>
      <el-col :span="2"
        ><div class="grid-content bg-purple-dark title">
          <el-button
            size="small"
            type="primary"
            v-if="
              paperDetail.status === 'UnReview' ||
              paperDetail.status === 'InReview'
            "
            @click="popup"
            >Review</el-button
          >
        </div></el-col
      >
      <el-col :span="24"
        ><div class="grid-content bg-purple-dark title">
          <span>{{ paperDetail.title }}</span>
        </div></el-col
      >

      <!-- <el-col :span="24"
      ><div
        class="grid-content bg-purple-dark"
        v-for="user in paperDetail.user"
        :key="paperDetail.user.ID"
      >
        {{ user.first_name }} {{ user.last_name }}
      </div></el-col
    > -->
      <el-col :span="22"
        ><div class="grid-content bg-purple-dark author">
          {{ paperDetail.authors }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Abstract:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.abstract }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Key Words:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.key_words }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Paper Type:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.paper_type }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">
          Subject Category:
        </div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.subject_category }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">
          Corresponding Author:
        </div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.cor_author }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">
          Block Address:
        </div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.block_address }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Paper:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          <!-- paper链接 -->
          <a :href="paperDetail.filepath" target="paper">{{
            paperDetail.title
          }}</a>
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Cid:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.cid }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Hash:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.hash }}
        </div></el-col
      >
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Version ID:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.version_id }}
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Reviews:</div></el-col
      >
      <el-col :span="24"
        ><div class="grid-content bg-purple-dark value">
          <el-table
            :data="review_infos"
            height="300"
            border
            stripe
            style="width: 100%"
            :header-cell-style="{ color: '#000000', background: '#F7FBFF' }"
          >
            <el-table-column
              fixed
              prop="reviewer_name"
              label="Reviewer"
              width="200px"
            >
            </el-table-column>
            <el-table-column prop="comment" label="Comment"> </el-table-column>
            <el-table-column prop="status" label="Result" width="200px">
            </el-table-column>
          </el-table></div
      ></el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">
          Historical Version:
        </div></el-col
      >
      <el-col :span="24"
        ><div class="grid-content bg-purple-dark value">
          <el-table
            :data="tableData"
            height="300"
            border
            stripe
            style="width: 100%"
            :header-cell-style="{
              color: '#000000',
              background: '#F7FBFF',
            }"
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
            <el-table-column fixed="right" label="Options" width="80">
              <template slot-scope="scope">
                <el-button
                  @click="viewPaper(scope.row)"
                  type="primary"
                  size="mini"
                  >View</el-button
                >
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import { getDetailPapers, submitReview, getPaperVersions } from "../../api";
import { ERC20contractInstance } from "@/constant";
const contractInstance = ERC20contractInstance;
export default {
  data() {
    return {
      userInfo: {},
      paperDetail: {
        ID: 0,
        version_id: 0,
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
      review_infos: [
        {
          reviewer_name: "",
          comment: "",
          result: "",
        },
      ],
      dialogVisible: false,
      reviewForm: {
        paper_id: 0,
        status: 0,
        comment: "",
      },
      usersBlockChainAddress: [],
    };
  },
  methods: {
    popup() {
      this.dialogVisible = true;
    },
    async reviewSubmit() {
      this.reviewForm.paper_id = this.paperDetail.ID;
      console.log(this.reviewForm, "reviewForm");
      // reviewForm的status要转换成对应的字符串
      if (this.reviewForm.status === 1) {
        this.reviewForm.status = "Accept";
      } else if (this.reviewForm.status === 2) {
        this.reviewForm.status = "Minor Revisions";
      } else if (this.reviewForm.status === 3) {
        this.reviewForm.status = "Major Revisions";
      } else if (this.reviewForm.status === 4) {
        this.reviewForm.status = "Reject and Resubmit";
      } else if (this.reviewForm.status === 5) {
        this.reviewForm.status = "Reject";
      }
      // await submitReview(this.reviewForm).then((res) => {
      //   console.log(res.data, "res");
      //   if (res.data.code === 1000) {
      //     this.dialogVisible = false;
      //     // 刷新页面
      //     this.$router.go(0);
      //     this.$message({
      //       message: "Review Success",
      //       type: "success",
      //     });
      //   } else {
      //     this.$message({
      //       message: "Review Failed",
      //       type: "error",
      //     });
      //   }
      // });
      // 把comment、status、paperd的title整合成一个字符串存储到区块链
      // 格式为：title：status：comment
      let reviewInfo = `title: ${this.paperDetail.title} comment: ${this.reviewForm.comment} status: ${this.reviewForm.status}`;
      // reviewInfo = "good good morning";
      console.log(reviewInfo, "reviewInfo");
      try {
        // 调用智能合约函数
        if (
          this.userInfo.block_chain_address !== window.ethereum.selectedAddress
        ) {
          this.fileList = [];
          this.$message({
            message: "Please use the your account_address",
            type: "warning",
          });
          return;
        }
        const functionArgs = [reviewInfo];
        const functionName = "storeReview";
        const result = await contractInstance.methods[functionName](
          ...functionArgs
        ).send({
          from: window.ethereum.selectedAddress,
          gasPrice: "0",
        });

        // 输出结果
        console.log("Transaction result:", result);
      } catch (error) {
        // 处理错误
        console.error("Error:", error);
        return;
      }
      try {
        // const resc = await this.callSmartContractStoreReviewFunction();
        // console.log(resc, "resc");
        const res = await submitReview(this.reviewForm);
        console.log(res.data, "res");
        if (res.data.code === 1000) {
          // 从localStorage中获取用户信息
          // const userInfo = JSON.parse(localStorage.getItem("userInfo"));
          // 老师说审核通过还不能发币
          // if (
          //   this.reviewForm.status === "Accept" &&
          //   userInfo.authorityId === 102
          // ) {
          //   await this.callSmartContractMintFunction();
          // }
          this.dialogVisible = false;
          // 刷新页面
          this.$router.go(0);
          this.$message({
            message: "Review Success",
            type: "success",
          });
        } else {
          this.$message({
            message: "Review Failed",
            type: "error",
          });
        }
      } catch (error) {
        console.error("Error submitting review:", error);
        this.$message({
          message: "An error occurred while submitting the review.",
          type: "error",
        });
      }
    },
    async callSmartContractMintFunction() {
      try {
        // 调用智能合约函数
        if (
          this.userInfo.block_chain_address !== window.ethereum.selectedAddress
        ) {
          this.fileList = [];
          this.$message({
            message: "Please use the your account_address",
            type: "warning",
          });
          return;
        }
        const functionName = "mint";
        console.log(
          "this.usersBlockChainAddress:",
          this.usersBlockChainAddress
        );
        const functionArgs = [
          this.usersBlockChainAddress,
          "100000000000000000000",
        ];
        console.log(
          "window.ethereum.selectedAddress:",
          window.ethereum.selectedAddress
        );

        const result = await contractInstance.methods[functionName](
          ...functionArgs
        ).send({
          from: window.ethereum.selectedAddress,
          gasPrice: "0",
        });

        // 输出结果
        console.log("Transaction result:", result);
      } catch (error) {
        // 处理错误
        console.error("Error:", error);
      }
    },
    async callSmartContractStoreReviewFunction() {
      // 把comment、status、paperd的title整合成一个字符串存储到区块链
      // 格式为：title：status：comment
      let reviewInfo = `${this.reviewForm.comment}:${this.reviewForm.status}:${this.paperDetail.title}`;
      console.log(reviewInfo, "reviewInfo");
      try {
        // 调用智能合约函数
        if (
          this.userInfo.block_chain_address !== window.ethereum.selectedAddress
        ) {
          this.fileList = [];
          this.$message({
            message: "Please use the your account_address",
            type: "warning",
          });
          return;
        }
        const functionArgs = [reviewInfo];
        const functionName = "storeReview";
        const result = await contractInstance.methods[functionName](
          ...functionArgs
        ).send({
          from: window.ethereum.selectedAddress,
          gasPrice: "0",
        });

        // 输出结果
        console.log("Transaction result:", result);
      } catch (error) {
        // 处理错误
        console.error("Error:", error);
      }
    },
    getPapers() {
      getPaperVersions({ version_id: this.paperDetail.version_id }).then(
        (res) => {
          // 过滤掉当前的paper
          res.data.data = res.data.data.filter(
            (item) => item.ID !== this.paperDetail.ID
          );
          // 格式化时间 2023-09-30T16:50:21.503+08:00 变成 2023-09-30
          for (let i = 0; i < res.data.data.length; i++) {
            res.data.data[i].CreatedAt = res.data.data[i].CreatedAt.substring(
              0,
              10
            );
          }

          this.tableData = res.data.data;
          // 把tableData里面的user数组里面的first_name和last_name拼接起来
          // for (let i = 0; i < this.tableData.length; i++) {
          //   for (let j = 0; j < this.tableData[i].user.length; j++) {
          //     // 如果authors没有初始化，就初始化为空
          //     if (this.tableData[i].authors === undefined) {
          //       this.tableData[i].authors = "";
          //     }
          //     // 如果是最后一个作者，就不加逗号
          //     if (j === this.tableData[i].user.length - 1) {
          //       this.tableData[i].authors +=
          //         this.tableData[i].user[j].first_name +
          //         " " +
          //         this.tableData[i].user[j].last_name;
          //     } else {
          //       this.tableData[i].authors +=
          //         this.tableData[i].user[j].first_name +
          //         " " +
          //         this.tableData[i].user[j].last_name +
          //         ",";
          //     }
          //   }
          // }
          console.log(this.tableData, "tableData");
        }
      );
    },
    cellStyle(row, column, rowIndex, columnIndex) {
      const status = row.row.status;
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
    getDetail() {
      this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
      // 把this.userInfo.block_chain_address中的大小字母转换成小写
      this.userInfo.block_chain_address =
        this.userInfo.block_chain_address.toLowerCase();
      getDetailPapers({ paper_id: this.$route.query.paper_id }).then((res) => {
        console.log(res.data.data);
        this.paperDetail = res.data.data.paper;
        this.review_infos = res.data.data.review_infos;
        this.paperDetail.filepath = this.paperDetail.filepath;
        // "http://172.16.170.195:8888/" + this.paperDetail.filepath;
        // "http://localhost:8888/" + this.paperDetail.filepath;
        // 把res.data.data.paper.user的first_name和last_name拼接起来 放在authors
        // for (let j = 0; j < this.paperDetail.user.length; j++) {
        //   this.usersBlockChainAddress.push(
        //     // 保存block_chain_address
        //     this.paperDetail.user[j].block_chain_address
        //   );
        //   // 如果authors没有初始化，就初始化为空
        //   if (this.paperDetail.authors === undefined) {
        //     this.paperDetail.authors = "";
        //   }
        //   // 如果是最后一个作者，就不加逗号
        //   if (j === this.paperDetail.user.length - 1) {
        //     this.paperDetail.authors +=
        //       this.paperDetail.user[j].first_name +
        //       " " +
        //       this.paperDetail.user[j].last_name;
        //   } else {
        //     this.paperDetail.authors +=
        //       this.paperDetail.user[j].first_name +
        //       " " +
        //       this.paperDetail.user[j].last_name +
        //       ", ";
        //   }
        // }
        // 解析key_words ["[\"Decentralization\",\"Transparency\",\"Cryptography\"]"]
        // this.paperDetail.key_words = JSON.parse(
        //   this.paperDetail.key_words
        // ).join(",");
        // // 解析subject_category ["[\"Private\",\"Blockchain\"]"]
        // this.paperDetail.subject_category = JSON.parse(
        //   this.paperDetail.subject_category
        // ).join(",");
        this.getPapers();
      });
    },
    viewPaper(row) {
      console.log(row.ID);
      const currentPath = this.$route.path;
      this.$router.replace({
        path: "/dummy", // 使用一个虚拟的路径
        query: {
          paper_id: row.ID,
        },
      });

      // 在下一个事件循环中跳转回原来的路径
      this.$nextTick(() => {
        this.$router.replace({
          path: currentPath,
          query: {
            paper_id: row.ID,
          },
        });
      });
    },
  },
  mounted() {
    this.getDetail();
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
.title {
  font-size: 25px;
  font-weight: bold;
  text-align: center;
}
.author {
  // 字体加粗
  font-weight: bold;
  font-size: 15px;
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