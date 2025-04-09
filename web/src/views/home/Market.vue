<template>
  <div class="box">
    <div class="box1">
      <el-dialog
        title="Confirm"
        :visible.sync="issueDialogVisible"
        width="30%"
        @close="cancel"
      >
        <p>Are you sure you want to purchase this NFT</p>
        <span slot="footer" class="dialog-footer">
          <el-button @click="cancel">Cancel</el-button>
          <el-button type="primary" @click="confirm">Confirm </el-button>
        </span>
      </el-dialog>
      <el-row>
        <div class="title">NFT Copyright trading market</div>
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
              <el-button type="primary" size="mini" @click="buy(index)"
                >Buy</el-button
              >
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>
<script>
import { getNFTInfoByTokenId, updatePaperUserId } from "../../api";
import {
  MarketplacecontractInstance,
  ERC20contractInstance,
  MarketplacecontractAddress,
} from "@/constant";
const contractInstance = MarketplacecontractInstance;
const ERC20contract = ERC20contractInstance;
export default {
  data() {
    return {
      userInfo: {},
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
    };
  },
  methods: {
    buy(index) {
      this.index = index;
      this.issueDialogVisible = true;
    },
    cancel() {
      this.issueDialogVisible = false;
    },
    async confirm() {
      console.log("this.index:", this.index);
      console.log("this.tokenIds[this.index]:", this.tokenIds[this.index]);
      //   Promise.all([
      //     // 调用ERC20合约的approve方法
      //     ERC20contract.methods["approve"](
      //       MarketplacecontractAddress,
      //       this.copy_right_trading_prices[this.index]
      //     ).send({
      //       from: window.ethereum.selectedAddress,
      //       gasPrice: "0",
      //     }),
      //     // 调用合约buyNFT方法，购买NFT
      //     contractInstance.methods["buy"](this.tokenIds[this.index]).send({
      //       from: window.ethereum.selectedAddress,
      //       gasPrice: "0",
      //     }),
      //   ]).then((res) => {
      //     console.log("res:", res);
      //   });

      const price = this.copy_right_trading_prices[this.index] * 10 ** 18;

      try {
        // 调用ERC20合约的approve方法
        if (
          this.userInfo.block_chain_address !== window.ethereum.selectedAddress
        ) {
          this.$message({
            message: "Please use the your account_address",
            type: "warning",
          });
          return;
        }
        const result = ERC20contract.methods["approve"](
          MarketplacecontractAddress,
          price
        ).send({
          from: window.ethereum.selectedAddress,
          gasPrice: "0",
        });
        console.log("ERC20contract result:", result);
        // 调用合约buyNFT方法，购买NFT
        const res = await contractInstance.methods["buy"](
          this.tokenIds[this.index]
        ).send({ from: window.ethereum.selectedAddress, gasPrice: "0" });
        console.log("res:", res);
        // 后端修改paper_user对应关系
        updatePaperUserId({
          paper_id: this.paper_ids[this.index],
        }).then((res) => {
          console.log("res:", res);
        });
      } catch (error) {
        // 处理错误
        console.error("Error:", error);
      }
      // 刷新页面
      this.$router.go(0);
    },
    // 调用合约getAllNFTs方法，获取正在出售的NFT
    async getAllNFTs() {
      try {
        // 从localStorage中获取用户信息
        this.userInfo = JSON.parse(localStorage.getItem("userInfo"));
        // 把this.userInfo.block_chain_address中的大小字母转换成小写
        this.userInfo.block_chain_address =
          this.userInfo.block_chain_address.toLowerCase();
        
        // 检查网络连接
        const networkId = await window.ethereum.request({ method: 'eth_chainId' });
        console.log("Current network ID:", networkId);
        
        if (networkId !== '0x198') {
          this.$message({
            message: "请切换到Papers Chain网络 (chainId: 0x198)",
            type: "warning"
          });
          return;
        }
          
        // 检查合约是否正确初始化
        console.log("Contract address:", MarketplacecontractAddress);
        
        try {
          // 方法一：尝试使用call方法直接获取NFT数据
          const result = await contractInstance.methods
            .getAllNFTs()
            .call({ from: window.ethereum.selectedAddress });

          console.log("Transaction result:", result);
          
          if (!result || result.length === 0) {
            console.log("没有可用的NFT");
            return;
          }
          
          // 拿到tokenId
          for (let i = 0; i < result.length; i++) {
            // 检查result[i]的结构，适应不同版本Web3返回的结果格式
            const tokenId = result[i].tokenId || result[i][1];
            if (tokenId) {
              this.token_ids += tokenId.toString() + ",";
            }
          }
          
          // 去掉最后一个逗号
          if (this.token_ids) {
            this.token_ids = this.token_ids.substring(0, this.token_ids.length - 1);
            console.log("this.token_ids:", this.token_ids);
            
            // 根据token_ids去后端查询NFT信息
            getNFTInfoByTokenId({ token_ids: this.token_ids }).then((res) => {
              console.log("res:", res.data.data);
              if (res.data && res.data.data && res.data.data.length > 0) {
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
              } else {
                console.log("没有找到相关的NFT信息");
              }
            }).catch(err => {
              console.error("获取NFT信息失败:", err);
            });
          }
        } catch (contractError) {
          console.error("Contract call error details:", {
            message: contractError.message,
            code: contractError.code,
            data: contractError.data
          });
          
          // 尝试获取具体的合约错误信息
          if (contractError.data) {
            try {
              // Web3 v1.7.4中获取错误细节的方式
              console.log("错误详细数据:", contractError.data);
            } catch (e) {
              console.log("无法解析错误详情:", e);
            }
          }
          
          // 尝试备选方案：使用getOrderLength和orders方法单独获取每个NFT
          try {
            console.log("尝试备选方案获取NFT数据...");
            const orderLength = await contractInstance.methods.getOrderLength().call();
            console.log("订单总数:", orderLength);
            
            if (orderLength > 0) {
              for (let i = 0; i < orderLength; i++) {
                const order = await contractInstance.methods.orders(i).call();
                console.log(`订单 ${i}:`, order);
                
                if (order && order.tokenId) {
                  this.token_ids += order.tokenId.toString() + ",";
                }
              }
              
              if (this.token_ids) {
                this.token_ids = this.token_ids.substring(0, this.token_ids.length - 1);
                console.log("备选方案 - tokenIds:", this.token_ids);
                
                getNFTInfoByTokenId({ token_ids: this.token_ids }).then((res) => {
                  if (res.data && res.data.data) {
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
                  }
                });
              }
            } else {
              console.log("没有可用的NFT订单");
            }
          } catch (backupError) {
            console.error("备选方案获取NFT数据失败:", backupError);
            
            // 显示用户友好的错误信息
            this.$message({
              message: "获取NFT数据失败，请检查网络连接或合约状态",
              type: "error"
            });
          }
        }
      } catch (error) {
        // 处理错误
        console.error("General error:", error);
        this.$message({
          message: "操作失败: " + error.message,
          type: "error"
        });
      }
    },
  },
  mounted() {
    this.getAllNFTs();
  },
};
</script>
<style lang="less" scoped>
.box {
  display: inline-block;
  text-align: center;
  .box1 {
    // 取消内容水平居中
    // text-align: left;
    width: 80%;
    height: 950px;
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
.title {
  font-size: 30px;
  font-weight: bold;
  margin-bottom: 20px;
}
.cid {
  font-size: 14px;
}
.transactionHash {
  font-size: 13px;
}
.image {
  width: 100%;
  display: block;
}
</style>
