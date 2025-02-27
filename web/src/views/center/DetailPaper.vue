<!--作用：显示论文的详细信息页面
功能：
显示论文的基本信息(标题、作者、摘要等)
显示论文的区块链信息
提供论文更新、修改和发布功能
PDF预览相关：包含论文PDF文件的预览链接(通过paperDetail.filepath访问)-->
<template>
  <div>
    <el-dialog :visible.sync="dialogVisible" width="740px" append-to-body>
      <!-- 显示证书 -->
      <!-- <img :src="url" class="image" /> -->
      <el-image
        style="width: 700px; height: 500px"
        :src="url"
        fit="fill"
      ></el-image>
      <el-form
        ref="formData"
        :inline="true"
        :model="formData"
        label-width="220px"
      >
        <el-form-item label="Download Token Price" prop="download_price">
          <el-input
            style="width: 120px"
            v-model="formData.download_price"
          ></el-input>
        </el-form-item>
        <el-form-item
          label="Copyright Trading Token Price"
          prop="copyright_trading_price"
        >
          <el-input
            style="width: 120px"
            v-model="formData.copyright_trading_price"
          ></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="publish">Publish</el-button>
      </span>
    </el-dialog>

    <el-row>
      <el-col :span="2"
        ><div class="grid-content bg-purple-dark title">
          <el-button
            size="small"
            type="primary"
            v-if="paperDetail.status === 'UnReview'"
            @click="handleUpdate"
            >Update</el-button
          >
        </div>
      </el-col>
      <el-col :span="2"
        ><div class="grid-content bg-purple-dark title">
          <el-button
            size="small"
            type="primary"
            v-if="
              paperDetail.status === 'Minor Revisions' ||
              paperDetail.status === 'Major Revisions'
            "
            @click="handleRevise"
            >Revise</el-button
          >
        </div>
      </el-col>
      <el-col :span="2"
        ><div class="grid-content bg-purple-dark title">
          <el-button
            size="small"
            type="primary"
            v-if="paperDetail.status === 'Accept'"
            @click="showHonoraryCertificate"
            >Publish</el-button
          >
        </div>
      </el-col>
      <el-col :span="24"
        ><div class="grid-content bg-purple-dark title">
          <span>{{ paperDetail.title }}</span>
        </div>
      </el-col>

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
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Abstract:</div>
      </el-col>
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.abstract }}
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Key Words:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.key_words }}
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Paper Type:</div>
      </el-col>
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.paper_type }}
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Subject Category:</div>
      </el-col>
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.subject_category }}
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">
          Corresponding Author:
        </div>
      </el-col>
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.cor_author }}
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">
          Block Address:
        </div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.block_address }}
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Paper:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          <!-- paper链接 -->
          <a :href="paperDetail.filepath" target="paper">{{
            paperDetail.title
          }}</a>
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Cid:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.cid }}
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Paper Hash:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.hash }}
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">Version ID:</div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.version_id }}
        </div>
      </el-col>

      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">
          NFT Transaction Hash:
        </div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.transaction_hash }}
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">
          Download Price:
        </div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.download_price }} MPER Token
        </div>
      </el-col>
      <el-col :span="4"
        ><div class="grid-content bg-purple-dark key">
          Copyright Trading Price:
        </div></el-col
      >
      <el-col :span="20"
        ><div class="grid-content bg-purple-dark value">
          {{ paperDetail.copyright_trading_price }} MPER Token
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
            :header-cell-style="{
              color: '#000000',
              background: '#F7FBFF',
            }"
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
          </el-table>
        </div>
      </el-col>
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
              label="Publishation Time"
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
import { from } from "vue-buffer";
import {
  getDetailPapers,
  getPaperVersions,
  getHonoraryCertificate,
  publishPaper,
  getConferenceDetail,
  getJournalDetail,
} from "../../api";
import { ERC721contractInstance } from "@/constant";
const contractInstance = ERC721contractInstance;
export default {
  data() {
    return {
      paperDetail: {
        ID: 0,
        version_id: 0,
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
          download_price: "",
          copyright_trading_price: "",
          transaction_hash: "",
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
      url: "",
      image_uri: "",
      metadata_uri: "",
      cid: "",
      userInfo: {
        block_chain_address: "",
      },
      formData: {
        paper_id: 0,
        download_price: "",
        copyright_trading_price: "",
        transaction_hash: "",
        token_id: "",
      },
      transctionResoult: {
        blockHash: "",
        events: {
          Transfer: {
            returnValues: {
              tokenId: "",
            },
          },
        },
        transactionHash: "",
      },
    };
  },
  methods: {
    async handleUpdate() {
      let catageory = "";
      if (this.paperDetail.conference_id !== 0) {
        await getConferenceDetail({
          conference_id: this.paperDetail.conference_id,
        }).then((res) => {
          console.log(res.data.data, "res.data.data");
          catageory = res.data.data.category;
        });
      } else {
        await getJournalDetail({
          journal_id: this.paperDetail.journal_id,
        }).then((res) => {
          console.log(res.data.data, "res.data.data");
          catageory = res.data.data.category;
        });
      }
      console.log(catageory, "catageory");
      this.$router.push({
        path: "/center/updatePaper",
        query: {
          paper_id: this.$route.query.paper_id,
          subject_category: catageory,
        },
      });
    },
    async handleRevise() {
      let catageory = "";
      if (this.paperDetail.conference_id !== 0) {
        await getConferenceDetail({
          conference_id: this.paperDetail.conference_id,
        }).then((res) => {
          console.log(res.data.data, "res.data.data");
          catageory = res.data.data.category;
        });
      } else {
        await getJournalDetail({
          journal_id: this.paperDetail.journal_id,
        }).then((res) => {
          console.log(res.data.data, "res.data.data");
          catageory = res.data.data.category;
        });
      }
      console.log(catageory, "catageory");
      await this.$router.push({
        path: "/center/revisePaper",
        query: {
          paper_id: this.$route.query.paper_id,
          subject_category: catageory,
        },
      });
    },
    showHonoraryCertificate() {
      // 请求后端获得荣誉证书，弹出荣誉证书预览点击确定后，调用后端接口发布论文，然后刷新页面
      this.dialogVisible = true;
      getHonoraryCertificate({
        paper_id: this.$route.query.paper_id,
      }).then((res) => {
        console.log(res.data.data, "res.data.data");
        this.url = res.data.data.url;
        this.image_uri = res.data.data.image_uri;
        this.metadata_uri = res.data.data.metadata_uri;
        this.cid = res.data.data.cid;
        // console.log(this.url, "this.url");
      });
    },
    async publish() {
      // 调用智能合约的铸币函数，将荣誉证书铸造到用户的钱包中
      try {
        // 调用智能合约函数
        if (
          this.userInfo.block_chain_address !== window.ethereum.selectedAddress
        ) {
          this.$message({
            message: "Please use the your account_address",
            type: "warning",
          });
          return;
        }
        const functionName = "safeMint";
        console.log(
          "this.usersBlockChainAddress:",
          this.userInfo.block_chain_address
        );
        const functionArgs = [
          this.userInfo.block_chain_address,
          this.metadata_uri,
        ];
        console.log(this.metadata_uri)
        console.log(
          "window.ethereum.selectedAddress:",
          window.ethereum.selectedAddress
        );

        const result = await contractInstance.methods[functionName](
          ...functionArgs
        ).send({
          from: window.ethereum.selectedAddress,
          gasPrice: "1000",
        });
        
        // 添加结果检查
        if (!result.events || !result.events.Transfer) {
          this.$message({
            message: "交易成功但未获取到Token ID，请检查交易记录",
            type: "warning"
          });
          return;
        }

        this.transctionResoult = result;
        this.formData.transaction_hash = result.transactionHash;
        this.formData.paper_id = parseInt(this.$route.query.paper_id);
        this.formData.download_price = parseInt(this.formData.download_price);
        this.formData.token_id = result.events.Transfer.returnValues.tokenId.toString();
        this.formData.copyright_trading_price = parseInt(
          this.formData.copyright_trading_price
        );
        console.log(this.formData, "this.formData");
        publishPaper(this.formData).then((res) => {
          console.log(res.data, "res.data");
          if (res.data.code === 1000) {
            this.$message({
              message: "Publish successfully",
              type: "success",
            });
            // 跳转到论文列表页面
            this.$router.push("/center/reviewedPapers");
          } else {
            this.$message({
              message: "Publish failed",
              type: "error",
            });
          }
        });
      } catch (error) {
        // 处理错误
        console.error("Error:", error);
        this.$message({
          message: "发布失败: " + error.message,
          type: "error"
        });
      }
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
    getDetail() {
      // console.log(this.$route.query.paper_id);
      getDetailPapers({ paper_id: this.$route.query.paper_id }).then((res) => {
        console.log(res.data.data, "res.data.data=====");
        this.paperDetail = res.data.data.paper;
        this.review_infos = res.data.data.review_infos;
        // 把res.data.data.paper.user的first_name和last_name拼接起来 放在authors
        // for (let j = 0; j < this.paperDetail.user.length; j++) {
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
        console.log(this.paperDetail, "this.paperDetail");
        this.getPapers();
        // 从localStorage中获取用户的block_chain_address
        this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
        // 把this.userInfo.block_chain_address中的大小字母转换成小写
        this.userInfo.block_chain_address =
          this.userInfo.block_chain_address.toLowerCase();
      });
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
.image {
  width: 700px;
  height: 500px;
  // 居中
  margin: 0 auto;
}
</style>