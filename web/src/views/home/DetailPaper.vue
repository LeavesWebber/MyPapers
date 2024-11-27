<template>
  <div class="box">
    <el-dialog :visible.sync="centerDialogVisible" width="30%" append-to-body>
      <span
        >The viewing price set by the author is
        {{ paperDetail.download_price }} MPER Token. Do you want to pay to view
        the paper ?</span
      >
      <span slot="footer" class="dialog-footer">
        <el-button @click="centerDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="handlePay">Pay</el-button>
      </span>
    </el-dialog>
    <div class="box1">
      <el-row>
        <el-col :span="22"
          ><div class="grid-content bg-purple-dark title">
            <span v-if="!editMode">{{ paperDetail.title }}</span>
            <el-input v-else v-model="editedTitle"></el-input></div
        ></el-col>
        <el-col :span="2"
          ><div class="grid-content bg-purple-dark title">
            <!-- <el-button
          size="small"
          type="primary"
          v-if="!editMode && paperDetail.status === 'UnReview'"
          @click="enterEditMode"
          >Edit</el-button
        >
        </el-button> -->
            <el-button
              size="small"
              type="primary"
              v-if="editMode"
              @click="saveChanges"
            >
              Save
            </el-button>
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
          ><div class="grid-content bg-purple-dark key">
            Paper Type:
          </div></el-col
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
        <el-col :span="5"
          ><div class="grid-content bg-purple-dark key">
            Corresponding Author:
          </div></el-col
        >
        <el-col :span="19"
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
            <!-- paper链接 
            <a :href="paperDetail.filepath" target="paper">{{
              paperDetail.title
            }}</a>-->
            <el-button type="primary" size="small" @click="viewPaper">
              View Paper
            </el-button>
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
          ><div class="grid-content bg-purple-dark key">
            Paper Hash:
          </div></el-col
        >
        <el-col :span="20"
          ><div class="grid-content bg-purple-dark value">
            {{ paperDetail.hash }}
          </div></el-col
        >
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
        <el-col :span="24"
          ><div class="grid-content bg-purple-dark value">
            <el-table
              :data="review_infos"
              stripe
              height="300"
              border
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
              <el-table-column prop="comment" label="Comment">
              </el-table-column>
              <el-table-column prop="status" label="Result" width="200px">
              </el-table-column>
            </el-table></div
        ></el-col>
      </el-row>
    </div>
  </div>
</template>
<script>
import { getDetailPapers, checkPaperViewer, addPaperViewer } from "../../api";
import { ERC20contractInstance } from "@/constant";
const contractInstance = ERC20contractInstance;
export default {
  data() {
    return {
      userInfo: {},
      paperDetail: {
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
        download_price: 0,
        copyright_trading_price: 0,
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
      review_infos: [
        {
          reviewer_name: "",
          comment: "",
          result: "",
        },
      ],
      editMode: false, // 编辑模式开关
      editedTitle: "", // 用于编辑标题的临时变量
      centerDialogVisible: false,
      addPaperViewer: {
        paper_id: 0,
      },
    };
  },
  methods: {
    viewPaper() {
      // 请求后端查看是否已经购买
      checkPaperViewer({ paper_id: this.$route.query.paper_id }).then((res) => {
        console.log(res.data);
        if (res.data.code === 1000) {
          // 如果已经购买，直接跳转链接浏览器打开
          if (res.data.data === true) {
            window.open(this.paperDetail.filepath);
          } else {
            // 如果没有购买，提示购买
            this.centerDialogVisible = true;
          }
        } else {
          this.$message({
            message: "Server Error",
            type: "warning",
          });
        }
      });
    },
    async handlePay() {
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
        const functionName = "transfer";
        const fee = this.paperDetail.download_price * Math.pow(10, 18);
        const functionArgs = [
          this.paperDetail.user[0].block_chain_address,
          fee,
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
        // 请求后端增加投稿可查看者
        this.addPaperViewer.paper_id = parseInt(this.$route.query.paper_id);
        addPaperViewer(this.addPaperViewer).then((res) => {
          console.log(res.data);
          if (res.data.code === 1000) {
            // 如果购买成功，直接跳转链接浏览器打开
            window.open(this.paperDetail.filepath);
            this.centerDialogVisible = false;
          } else {
            this.$message({
              message: "Purchase failed",
              type: "warning",
            });
          }
        });
      } catch (error) {
        // 处理错误
        console.error("Error:", error);
      }
    },
    // async callSmartContractMintFunction() {
    //   // 在这里调用智能合约的铸币函数
    //   // 例如：smartContract.mintCoins();
    //   // 请根据你的智能合约接口进行相应的调用
    //   try {
    //     // 调用智能合约函数
    //     const functionName = "mint";
    //     console.log(
    //       "this.usersBlockChainAddress:",
    //       this.usersBlockChainAddress
    //     );
    //     const functionArgs = [
    //       this.usersBlockChainAddress,
    //       "100000000000000000000",
    //     ];
    //     console.log(
    //       "window.ethereum.selectedAddress:",
    //       window.ethereum.selectedAddress
    //     );

    //     const result = await contractInstance.methods[functionName](
    //       ...functionArgs
    //     ).send({
    //       from: window.ethereum.selectedAddress,
    //       gasPrice: "0",
    //     });

    //     // 输出结果
    //     console.log("Transaction result:", result);
    //   } catch (error) {
    //     // 处理错误
    //     console.error("Error:", error);
    //   }
    // },
  },
  mounted() {
    this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
    // 把this.userInfo.block_chain_address中的大小字母转换成小写
    this.userInfo.block_chain_address =
      this.userInfo.block_chain_address.toLowerCase();
    // console.log(this.$route.query.paper_id);
    this.addPaperViewer.paper_id = parseInt(this.$route.query.paper_id);
    console.log(this.addPaperViewer.paper_id, "this.addPaperViewer.paper_id");
    getDetailPapers({ paper_id: this.$route.query.paper_id }).then((res) => {
      console.log(this.$route.query.paper_id, "this.$route.query.paper_id");
      console.log(res.data.data);
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
      this.paperDetail.key_words = JSON.parse(this.paperDetail.key_words).join(
        ","
      );
      // 解析subject_category ["[\"Private\",\"Blockchain\"]"]
      this.paperDetail.subject_category = JSON.parse(
        this.paperDetail.subject_category
      ).join(",");
    });
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
    height: 1200px;
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
    .title {
      font-size: 25px;
      font-weight: bold;
      text-align: center;
    }
    .author {
      // 字体加粗
      font-weight: bold;
      font-size: 15px;
      // 居中
      text-align: center;
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