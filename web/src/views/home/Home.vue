<template>
  <div class="box">
    <div class="hero-section">
      <img src="../../images/xmut.jpg" class="image" alt="厦门理工学院" />
    </div>
    
    <div class="content-section">
      <div class="link-section">
        <h2 class="section-title">项目地址</h2>
        <div class="link-item">
          <div class="link-label">Gitee 仓库</div>
          <a href="https://gitee.com/xmrchen/MyPapers" target="_blank" class="link">
            https://gitee.com/xmrchen/MyPapers
          </a>
        </div>
        <div class="link-item">
          <div class="link-label">Github 仓库</div>
          <a href="https://github.com/xmrchen/MyPapers" target="_blank" class="link">
            https://github.com/xmrchen/MyPapers
          </a>
        </div>
      </div>

      <div class="teacher-section">
        <h2 class="section-title">指导老师</h2>
        <div class="teacher-info">
          <span class="school">厦门理工学院</span>
          <span class="teacher-name">陈仁</span>
        </div>
      </div>
    </div>
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
  methods: {
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
.box {
  width: 100%;
  min-height: 100vh;
  background-color: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
  
  .hero-section {
    width: 100%;
    position: relative;
    padding: 3rem 0;
    background: linear-gradient(to bottom, #f8f9fa, #fff);
    display: flex;
    justify-content: center;
    
    &::after {
      content: '';
      position: absolute;
      bottom: 0;
      left: 0;
      right: 0;
      height: 100px;
      background: linear-gradient(to bottom, transparent, #fff);
      pointer-events: none;
    }
    
    .image {
      width: 1200px;
      height: 430px;
      object-fit: cover;
      border-radius: 16px;
      box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
      transition: transform 0.3s ease;
      
      &:hover {
        transform: translateY(-5px);
      }
    }
  }

  .content-section {
    position: relative;
    z-index: 1;
    width: 1200px;
    margin: -2rem auto 2rem;
    padding: 2.5rem;
    background-color: #fff;
    border-radius: 16px;
    box-shadow: 0 0 30px rgba(0, 0, 0, 0.06);

    .section-title {
      font-size: 1.6rem;
      color: #2c3e50;
      margin-bottom: 1.8rem;
      padding-bottom: 0.8rem;
      border-bottom: 2px solid #3498db;
      position: relative;
      
      &::after {
        content: '';
        position: absolute;
        bottom: -2px;
        left: 0;
        width: 50px;
        height: 2px;
        background-color: #2980b9;
      }
    }

    .link-section {
      margin-bottom: 3rem;

      .link-item {
        margin-bottom: 1.2rem;
        padding: 1.2rem;
        background-color: #f8f9fa;
        border-radius: 12px;
        transition: all 0.3s ease;
        border: 1px solid transparent;

        &:hover {
          transform: translateY(-3px);
          border-color: #e8e8e8;
          box-shadow: 0 8px 24px rgba(0, 0, 0, 0.05);
        }

        .link-label {
          font-size: 1.1rem;
          color: #2c3e50;
          margin-bottom: 0.8rem;
          font-weight: 500;
        }

        .link {
          color: #3498db;
          text-decoration: none;
          font-size: 1rem;
          display: inline-block;
          padding: 0.3rem 0;
          border-bottom: 1px solid transparent;
          
          &:hover {
            color: #2980b9;
            border-bottom-color: #2980b9;
          }
        }
      }
    }

    .teacher-section {
      .teacher-info {
        padding: 1.5rem;
        background-color: #f8f9fa;
        border-radius: 12px;
        display: flex;
        align-items: center;
        gap: 1.2rem;
        transition: all 0.3s ease;
        border: 1px solid transparent;

        &:hover {
          border-color: #e8e8e8;
          box-shadow: 0 8px 24px rgba(0, 0, 0, 0.05);
        }

        .school {
          color: #2c3e50;
          font-size: 1.1rem;
          position: relative;
          padding-right: 1.2rem;

          &::after {
            content: '·';
            position: absolute;
            right: 0;
            color: #606266;
          }
        }

        .teacher-name {
          color: #2c3e50;
          font-weight: 600;
          font-size: 1.1rem;
        }
      }
    }
  }
}
</style>
