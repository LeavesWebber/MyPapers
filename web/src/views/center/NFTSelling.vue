<template>
  <div>
    <el-dialog
      title="Cancel Sell"
      :visible.sync="issueDialogVisible"
      width="30%"
      @close="cancel"
    >
      <p>Are you sure to cancel sell this NFT ?</p>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">Cancel</el-button>
        <el-button type="primary" @click="confirm">Confirm </el-button>
      </span>
    </el-dialog>
    <el-row>
      <div class="title">Selling NFTs</div>
      <el-col
        :span="12"
        v-for="(image_url, index) in image_urls"
        :key="index"
        :offset="0"
      >
        <el-card :body-style="{ padding: '5px' }">
          <img :src="image_url" class="image" />
          <div class="transactionHash">
            {{ transaction_hashs[index] }}
          </div>
          <div class="cid">cid: {{ image_cids[index] }}</div>
          <div>
            tokenId: {{ tokenIds[index] }} &nbsp; &nbsp;&nbsp; Price:
            {{ copy_right_trading_prices[index] }} MPER
          </div>
          <div>
            <el-button type="primary" size="mini" @click="cancelSell(index)"
              >Cancel Sell</el-button
            >
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import { getNFTInfoByTokenId } from "../../api";
import { MarketplacecontractInstance } from "@/constant";
const contractInstance = MarketplacecontractInstance;
export default {
  data() {
    return {
      token_ids: "",
      imageCount: 10,
      image_urls: [],
      tokenIds: [],
      copy_right_trading_prices: [],
      image_cids: [],
      transaction_hashs: [],
      paper_ids: [],
      currentDate: new Date(),
      formData: {
        paper_id: "",
        download_price: 0,
        copyright_trading_price: 0,
      },
      issueDialogVisible: false,
      index: 0,
      userInfo: "",
    };
  },
  methods: {
    cancelSell(index) {
      this.index = index;
      this.issueDialogVisible = true;
    },
    cancel() {
      this.issueDialogVisible = false;
    },
    async confirm() {
      if (
        this.userInfo.block_chain_address !== window.ethereum.selectedAddress
      ) {
        this.$message({
          message: "Please use the your account_address",
          type: "warning",
        });
        return;
      }
      console.log("this.index:", this.index);
      console.log("this.tokenIds[this.index]:", this.tokenIds[this.index]);
      // 调用合约cancelOrder方法，取消出售NFT
      const res = await contractInstance.methods["cancelOrder"](
        this.tokenIds[this.index]
      ).send({ from: window.ethereum.selectedAddress, gasPrice: "0" });
      console.log("res:", res);
      // 刷新页面
      this.$router.go(0);
    },
    // 调用合约getMyNFTs方法，获取正在出售的NFT
    async getMyNFTs() {
      this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
      // 把this.userInfo.block_chain_address中的大小字母转换成小写
      this.userInfo.block_chain_address =
        this.userInfo.block_chain_address.toLowerCase();
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
        const result = await contractInstance.methods
          .getMyNFTs()
          .call({ from: window.ethereum.selectedAddress });

        console.log("Transaction result:", result);
        // 拿到tokenId
        for (let i = 0; i < result.length; i++) {
          this.token_ids += result[i].tokenId.toString() + ",";
        }
        console.log("this.token_ids:", this.token_ids);

        // 去掉最后一个逗号
        this.token_ids = this.token_ids.substring(0, this.token_ids.length - 1);
        console.log("this.token_ids:", this.token_ids);
        if (this.token_ids !== "0") {
          // 根据token_ids去后端查询NFT信息
          getNFTInfoByTokenId({ token_ids: this.token_ids }).then((res) => {
            console.log("res:", res.data.data);
            for (let i = 0; i < res.data.data.length; i++) {
              this.image_urls.push(res.data.data[i].image_url);
              this.tokenIds.push(res.data.data[i].token_id);
              this.copy_right_trading_prices.push(
                res.data.data[i].copy_right_trading_price
              );
              this.paper_ids.push(res.data.data[i].paper_id);
              this.image_cids.push(res.data.data[i].image_cid);
              this.transaction_hashs.push(res.data.data[i].transaction_hash);
            }
          });
          // 展示NFT信息
        }
      } catch (error) {
        // 处理错误
        console.error("Error:", error);
      }
      // try {
      //   // 调用智能合约函数
      //   console.log(
      //     "window.ethereum.selectedAddress:",
      //     window.ethereum.selectedAddress
      //   );

      //   const result = await contractInstance.methods["balanceOf"](
      //     window.ethereum.selectedAddress
      //   ).call({
      //     from: window.ethereum.selectedAddress,
      //     gasPrice: "0",
      //   });

      //   // 输出结果
      //   console.log("Transaction result:", result);
      // } catch (error) {
      //   // 处理错误
      //   console.error("Error:", error);
      // }
    },
  },
  mounted() {
    this.getMyNFTs();
  },
};
</script>
<style lang="less" scoped>
.title {
  font-size: 30px;
  font-weight: bold;
  margin-bottom: 20px;
}
.cid {
  font-size: 14px;
}
.transactionHash {
  font-size: 14px;
}
.image {
  width: 100%;
  display: block;
}
</style>
