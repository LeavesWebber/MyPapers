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
      // 拿到当前token_id，当前用户的地址，交易合约地址，交易价格
      // token_id
      const token_id = this.tokenIds[this.index];
      // 从狐狸钱包中获取当前用户的地址
      const user_address = window.ethereum.selectedAddress;
      console.log(this.copy_right_trading_prices[this.index], "this.index");
      // 交易价格
      const trading_price =
        this.copy_right_trading_prices[this.index] * 10 ** 18;
      console.log(trading_price, "trading_price");
      // 转为0x开头的16进制,像这样0x0000000000000000000000000000000000000000000000000001c6bf52634000
      // 将输入转换为十六进制，并移除0x前缀
      // const hexString = web3.utils.toHex(trading_price).substr(2);
      // // 使用web3.js提供的padLeft方法添加左填充
      // const paddedString = web3.utils.padLeft(hexString, 64);
      // // 加上0x前缀
      // console.log("paddedString:", paddedString);

      // const resultWithPrefix = "0x" + paddedString;
      // console.log("resultWithPrefix :", resultWithPrefix);

      // 将交易价格转换为 BigNumber 对象
      const bigNum = new BigNumber(trading_price.toString());
      // 转换为十六进制字符串
      const resultWithPrefix = "0x" + bigNum.toString(16).padStart(64, "0");

      console.log(resultWithPrefix, "trading_price_hex");

      const functionArgs = [
        user_address,
        contractAddress,
        token_id,
        resultWithPrefix,
      ];
      console.log(
        "window.ethereum.selectedAddress:",
        window.ethereum.selectedAddress
      );
      const result = contractInstance.methods["safeTransferFrom"](
        ...functionArgs
      ).send({
        from: window.ethereum.selectedAddress,
        gasPrice: "0",
      });
      console.log("result:", result);
      this.issueDialogVisible2 = false;
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
