<template>
  <div class="box">
    <div class="image-slider">
      
      <div class="image">
        <img :src="currentImage" alt="slider - image" class="image">
        <button @click="prevImage" v-if="imageIndex > 0" class="prev-button"><</button>
      
      <button @click="nextImage" v-if="imageIndex < images.length - 1" class="next-button">></button>
      </div>
      
      <div class="image-indicators">
        <span
        v-for="(image, index) in images"
        :key="index"
        :class="{active: index === imageIndex}"
        @click="switchToImage(index)"
      >
        {{index + 1}}
      </span>
      </div>
    </div>
    
    
    

    <!--
    ><div class="box1">
      <div class="box2">
        <h3 class="h3">The contract address used by the platform</h3>
        <li class="li">
          ERC20 Token Contract Address:
          0x10A62A42D44FA1274F70E016b20B5065Db0F5327
        </li>
        <li class="li">
          ERC721 NFT Contract Address:
          0x0BF4bb730bF29115cE6D6C0cDE6A2F314fba151e
        </li>
        <li class="li">
          Market Contract Address: 0x2b39296Ea695586A341D510C9EAFBECd5664bb7e
        </li>
        <h3 class="h1">
          Verify the author on the consortium chain through a paper hash
        </h3>
        <el-row>
          <el-col :span="10">
            <el-input v-model="input" placeholder="Paper Hash"></el-input>
          </el-col>
          <el-col :span="2">
            <el-button type="primary" @click="callContract">Verify</el-button>
          </el-col>
          <el-col :span="10"> Author Address: {{ author_address }} </el-col>
        </el-row>
        <h3 class="h1">Verify Reviewer through Comment Content</h3>
        <el-row>
          <el-col :span="10">
            <el-input v-model="title" placeholder="Paper Title"></el-input>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-input v-model="comment" placeholder="Paper Comment"></el-input>
          </el-col>

          <el-col :span="4">
            <el-input v-model="status" placeholder="Status"></el-input>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="2">
            <el-button type="primary" @click="getReviewByHash"
              >Verify</el-button
            >
          </el-col>
          <el-col :span="24"> Reviewer Address: {{ reviewer_address }} </el-col>
        </el-row>
        <!-- <h3 class="h1">The latest from MyPapers</h3>
        <hr class="hr" size="8px" color="#ECF5FF" width="110px" />
        <h3 class="h3">Some recent situations</h3>
        <li class="li">The First</li>
        <li class="li">The Second</li> 
      </div>
      <div>
        <el-button @click="showERC20Contract">ERC20 Token Contract</el-button>
        <el-button @click="showERC721Contract">ERC721 NFT Contract</el-button>
        <el-button @click="showMarketContract">Market Contract</el-button>
      </div>
      <div v-highlight v-if="change === 0">
        <h3 class="h1">ERC20 Token Contract</h3>
        <pre :key="'erc20Contract'">
          <code>
            {{erc20Contract}}
          </code>
        </pre>
      </div>

      <div v-highlight v-if="change === 1">
        <h3 class="h1">ERC721 NFT Contract</h3>
        <pre :key="'erc721Contract'">
          <code>
            {{erc721Contract}}
          </code>
        </pre>
      </div>

      <div v-highlight v-if="change === 2">
        <h3 class="h1">Market Contract</h3>
        <pre :key="'marketContract'">
          <code>
            {{marketContract}}
          </code>
        </pre>
      </div>-->

      <!-- <div class="box3">
        <div class="content1">
          <h3 class="title">MyPapers Annual Election</h3>
          <img src="../../images/1.jpg" width="90%" height="60%" />
          <div class="text">
            The 2023 MypPapers Election is underway. Be sure to cast your vote
            for the leaders who will inform MypPapers's strategy for the future.
          </div>
          <li class="li">The First</li>
          <li class="li">The Second</li>
        </div>
        <div class="link-item">
          <div class="link-label">Github 仓库</div>
          <a href="https://github.com/xmrchen/MyPapers" target="_blank" class="link">
            https://github.com/xmrchen/MyPapers
          </a>
        </div>
      </div>
      <div class="box4">
        <div class="content2">
          <h6>Smart Cities infrastructure</h6>
          <img src="../../images/2.jpg" width="90%" height="60%" />
          <div class="title">
            Smart Cities help urban environments by using moderntechnology to
            provide sustainable, resilient, equitable,and privacy-respecting
            communities
          </div>
        </div>
        <div class="content2">
          <h6>Smart Cities infrastructure</h6>
          <img src="../../images/2.jpg" width="90%" height="60%" />
          <div class="title">
            Smart Cities help urban environments by using moderntechnology to
            provide sustainable, resilient, equitable,and privacy-respecting
            communities
          </div>
        </div>
        <div class="content2">
          <h6>Smart Cities infrastructure</h6>
          <img src="../../images/2.jpg" width="90%" height="60%" />
          <div class="title">
            Smart Cities help urban environments by using moderntechnology to
            provide sustainable, resilient, equitable,and privacy-respecting
            communities
          </div>
        </div>
        <div class="content2">
          <h6>Smart Cities infrastructure</h6>
          <img src="../../images/2.jpg" width="90%" height="60%" />
          <div class="title">
            Smart Cities help urban environments by using moderntechnology to
            provide sustainable, resilient, equitable,and privacy-respecting
            communities
          </div>
        </div>
      </div> 
    </div>-->
  </div>
</template>

<script>
/* import {
  ERC20contractInstance,
  ERC20contract,
  ERC721contract,
  Marketcontract,
} from "@/constant";
import hljs from "highlight.js";
 */
// import { getHomeData } from "../../api";

export default {
  data() {
    return {
      images: [
        require('@/images/1.png'),
        require('@/images/2.png'),
        require('@/images/3.png'),
        require('@/images/4.png'),
        require('@/images/5.png'),
        require('@/images/6.png')
      ],
      imageIndex: 0,
      msg: "Welcome to Your Vue.js App",
      input: "",
      author_address: "",
      reviewer_address: "",
      comment: "",
      title: "",
      status: "",
      erc20Contract: ERC20contract,
      erc721Contract: ERC721contract,
      marketContract: Marketcontract,
      change: 0,
    };
  },
  computed: {
    currentImage() {
      return this.images[this.imageIndex];
    }
  },
  methods: {
    prevImage() {
      this.imageIndex--;
    },
    nextImage() {
      this.imageIndex++;
    },
    switchToImage(index) {
      this.imageIndex = index;
    },
    async callContract() {
      this.author_address = await ERC20contractInstance.methods
        .getRecipientByHash(this.input)
        .call({ from: window.ethereum.selectedAddress });

      console.log("Transaction result:", this.author_address);
    },
    async getReviewByHash() {
      let reviewInfo = `title: ${this.title} comment: ${this.comment} status: ${this.status}`;
      this.reviewer_address = await ERC20contractInstance.methods
        .getReviewByHash(reviewInfo)
        .call({ from: window.ethereum.selectedAddress });
      console.log(reviewInfo, "reviewInfo");

      console.log("Transaction result:", this.reviewer_address);
    },
    async connectWallet() {
      try {
        // 检查是否存在 window.ethereum 对象
        if (!window.ethereum) {
          alert("Please install the MetaMask plugin to continue the operation");
          return;
        }

        // 尝试请求用户授权
        const accounts = await window.ethereum.request({
          method: "eth_requestAccounts",
        });

        if (accounts.length === 0) {
          alert("请连接到以太坊网络以继续操作");
        } else {
          // 用户已连接，可以执行其他操作或初始化
          console.log("用户已连接，可以执行其他操作");
        }
      } catch (error) {
        // 处理错误
        console.error("连接钱包时出错:", error);
      }
    },
    showERC20Contract() {
      this.change = 0;
      this.highlightCodeBlocks();
    },
    showERC721Contract() {
      this.change = 1;
      this.highlightCodeBlocks();
    },
    showMarketContract() {
      this.change = 2;
      this.highlightCodeBlocks();
    },
    highlightCodeBlocks() {
      // 在页面更新后重新高亮所有代码块
      this.$nextTick(() => {
        let blocks = this.$el.querySelectorAll("pre code");
        blocks.forEach((block) => {
          hljs.highlightBlock(block);
        });
      });
    },
  },
  mounted() {
    // getHomeData().then((res) => {
    //   console.log(res);
    // });
    this.connectWallet(); // 在组件创建时尝试连接钱包
  },
};
</script>

<style lang="less" scoped>
.image-slider {
  text-align: center;
  position: relative;
}

.slider-container {
  position: relative;
  display: inline-block;
}

.prev-button {
  position: absolute;
  top: 0;
  height: 430px;
  width: 10%; /* 每个按钮占图片宽度的一半 */
  background-color: rgba(0, 0, 0, 0.3); /* 半透明背景 */
  margin-left: 18%;
  color: white;
  font-size: 24px;
  border: none;
  cursor: pointer;
  opacity: 0.2;
  z-index: 1;
}

.next-button {
  position: absolute;
  top: 0;
  height: 430px;
  width: 10%; 
  background-color: rgba(0, 0, 0, 0.3); /* 半透明背景 */
  margin-right: 18%;
  color: white;
  font-size: 24px;
  border: none;
  cursor: pointer;
  opacity: 0.2;
  z-index: 1;
}
.prev-button {
  left: 0;
}

.next-button {
  right: 0;
}

.prev-button:hover,
.next-button:hover {
  opacity: 1;
}
 .image-slider {
  text-align: center;
}
  .image-indicators {
  margin-top: 10px;
}

    .image-indicators span {
  display: inline-block;
  width: 20px;
  height: 20px;
  line-height: 20px;
  border-radius: 50%;
  background-color: #ccc;
  color: white;
  margin: 0 5px;
  cursor: pointer;
  }

    .image-indicators span.active {
  background-color: #4CAF50;
  }

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
    // 居中
    margin: 0 auto;
  }
  .box1 {
    // 取消内容水平居中
    text-align: left;
    width: 80%;
    // height: 950px;
    // background-color: #e2f1fb;
    background-color: #ecf5ff;
    // 水平居中
    margin: 0 auto;
    margin-top: 20px;
    // 上下左右panding
    padding: 20px 20px;
    color: #072e5b;
    .box2 {
      width: 100%;
      margin-bottom: 40px;
      .h1 {
        margin-bottom: 10px;
      }
      .h3 {
        margin-top: 10px;
        margin-bottom: 10px;
      }
      .li {
        margin-bottom: 10px;
      }
    }
    .box3 {
      width: 100%;
      height: 450px;
      // background-color: #ECF5FF;
      // margin-bottom: 10px;
      .content1 {
        width: 50%;
        height: 100%;
        // background-color: #d9eddd;
        float: left;
        .title {
          // 字号
          font-size: 15px;
          margin-bottom: 10px;
        }
        .bottom {
          margin-bottom: 10px;
        }
        .text {
          margin-bottom: 10px;
          width: 85%;
          // 字号
          font-size: 12px;
        }
        .li {
          margin-bottom: 10px;
          // 字号
          font-size: 12px;
        }
      }
    }
    .box4 {
      width: 100%;
      height: 300px;
      // background-color: #ECF5FF;
      .content2 {
        float: left;
        width: 25%;
        height: 100%;
        // background-color: #eff5f0;
        .title {
          // 字号
          font-size: 8px;
        }
      }
    }
  }
}
</style>
