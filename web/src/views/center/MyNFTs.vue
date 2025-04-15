<template>
  <div>
    <el-dialog
      title="Change Price"
      :visible.sync="issueDialogVisible"
      width="30%"
      @close="cancel"
    >
      <!-- Issue的表单信息 -->
      <el-form
        ref="formData"
        :inline="true"
        :model="formData"
        label-width="180px"
      >
        <el-form-item label="Download Price" prop="download_price">
          <el-input
            style="width: 220px"
            v-model="formData.download_price"
          ></el-input>
        </el-form-item>
        <el-form-item
          label="Copyrigrt Trading Price"
          prop="copyright_trading_price"
        >
          <el-input
            style="width: 220px"
            v-model="formData.copyright_trading_price"
          ></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">Cancel</el-button>
        <el-button type="primary" @click="submit">Change </el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="Confirm Price"
      :visible.sync="issueDialogVisible2"
      width="30%"
      @close="cancel2"
    >
      <p>
        Copyrigrt Trading Price: {{ copy_right_trading_prices[index] }} MPER
      </p>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel2">Cancel</el-button>
        <el-button type="primary" @click="confirmSell">Confirm </el-button>
      </span>
    </el-dialog>
    <el-row>
      <div class="title">My NFTs</div>
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
            <el-button type="primary" size="mini" @click="changePrice(index)"
              >Change Price</el-button
            >
            <el-button type="primary" size="mini" @click="sell(index)"
              >Sell</el-button
            >
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import { getMyNFTs, updatePrice } from "../../api";
import { MarketplacecontractAddress, ERC721contractInstance } from "@/constant";
import Web3 from "web3";
import BigNumber from "bignumber.js";
const web3 = new Web3(window.ethereum);
const contractAddress = MarketplacecontractAddress;
const contractInstance = ERC721contractInstance;
export default {
  data() {
    return {
      userInfo: {},
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
      issueDialogVisible2: false,
      index: 0,
    };
  },
  methods: {
    changePrice(index) {
      this.index = index;
      this.issueDialogVisible = true;
    },
    sell(index) {
      this.index = index;
      this.issueDialogVisible2 = true;
    },
    cancel() {
      this.issueDialogVisible = false;
    },
    cancel2() {
      this.issueDialogVisible2 = false;
    },
    submit() {
      this.formData.paper_id = this.paper_ids[this.index];
      this.formData.download_price = Number(this.formData.download_price);
      this.formData.copyright_trading_price = Number(
        this.formData.copyright_trading_price
      );

      console.log(this.formData, "this.formData");
      updatePrice(this.formData).then((res) => {
        console.log(res.data, "res.data");
        if (res.data.code === 1000) {
          this.$message({
            message: "Change successfully",
            type: "success",
          });
          // 刷新页面
          this.$router.go(0);
        } else {
          this.$message({
            message: "Change failed",
            type: "error",
          });
        }
      });
    },
    confirmSell() {
      if (
        this.userInfo.block_chain_address !== window.ethereum.selectedAddress
      ) {
        this.$message({
          message: "Please use the your account_address",
          type: "warning",
        });
        return;
      }
      
      try {
        // 拿到当前token_id，当前用户的地址，交易合约地址，交易价格
        const token_id = this.tokenIds[this.index];
        const user_address = window.ethereum.selectedAddress;
        
        // 交易价格（转换为wei单位）
        const trading_price = this.copy_right_trading_prices[this.index] * 10 ** 18;
        console.log("交易价格:", trading_price);
        
        // 按照合约要求，我们需要将价格数据正确编码
        // 合约中使用了onERC721Received方法来处理通过safeTransferFrom接收的NFT
        // 价格数据需要以字节形式传递
        const priceData = web3.eth.abi.encodeParameter('uint256', trading_price);
        console.log("编码后的价格数据:", priceData);
        
        // safeTransferFrom函数的最后一个参数应该是bytes类型的数据
        const functionArgs = [
          user_address,          // from地址
          contractAddress,       // to地址（合约地址）
          token_id,              // tokenId
          priceData              // 数据（价格）
        ];
        
        console.log("调用参数:", functionArgs);
        console.log("调用地址:", window.ethereum.selectedAddress);
        
        // 调用合约方法
        const result = contractInstance.methods["safeTransferFrom"](
          ...functionArgs
        ).send({
          from: window.ethereum.selectedAddress
        });
        
        // 处理Promise返回结果
        result.then(tx => {
          console.log("交易成功:", tx);
          this.$message({
            message: "NFT已成功挂单出售",
            type: "success"
          });
          // 延迟关闭对话框
          setTimeout(() => {
            this.issueDialogVisible2 = false;
            // 刷新数据
            this.$router.go(0);
          }, 1000);
        }).catch(error => {
          console.error("交易失败:", error);
          this.$message.error("NFT出售失败: " + error.message);
        });
      } catch (error) {
        console.error("准备交易时出错:", error);
        this.$message.error("准备NFT出售失败: " + error.message);
      }
    },
  },
  mounted() {
    getMyNFTs().then((res) => {
      this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
      // 把this.userInfo.block_chain_address中的大小字母转换成小写
      this.userInfo.block_chain_address =
        this.userInfo.block_chain_address.toLowerCase();
      console.log(res.data.data);
      // data有多个图片的uri
      for (let i = 0; i < res.data.data.length; i++) {
        this.image_urls.push(res.data.data[i].image_uri);
        this.tokenIds.push(res.data.data[i].token_id);
        this.copy_right_trading_prices.push(
          res.data.data[i].copy_right_trading_price
        );
        this.paper_ids.push(res.data.data[i].paper_id);
        this.image_cids.push(res.data.data[i].image_cid);
        this.transaction_hashs.push(res.data.data[i].transaction_hash);
      }
    });
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
