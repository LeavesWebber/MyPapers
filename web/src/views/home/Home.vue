<template>
  <div class="box">
    <div class="image-slider">

      <div class="image">
        <div class="slider-image-container">
          <transition name="fade">
            <img :src="currentImage" alt="slider-image" class="slider-image" :key="imageIndex">
          </transition>
          <button @click="prevImage" v-if="imageIndex > 0" class="prev-button"><</button>

        <button @click="nextImage" v-if="imageIndex < images.length - 1" class="next-button">></button>
        </div>
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
  </div>
</template>
<script>
import {
  ERC20contractInstance,
  ERC20contract,
  ERC721contract,
  Marketcontract,
} from "@/constant";
import hljs from "highlight.js";

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
      autoPlayTimer: null,
      imageAspectRatio: 16/9, // 默认宽高比
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
    startAutoPlay() {
      this.autoPlayTimer = setInterval(() => {
        if (this.imageIndex < this.images.length - 1) {
          this.imageIndex++;
        } else {
          this.imageIndex = 0;
        }
      }, 3000); // 每3秒切换一次图片
    },
    stopAutoPlay() {
      if (this.autoPlayTimer) {
        clearInterval(this.autoPlayTimer);
      }
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
    this.startAutoPlay();
    this.connectWallet();
    // 计算第一张图片的宽高比
    const img = new Image();
    img.src = this.images[0];
    img.onload = () => {
      this.imageAspectRatio = img.width / img.height;
    };
  },
  beforeDestroy() {
    this.stopAutoPlay();
  },
};
</script>
<style lang="less" scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}

.image-slider {
  text-align: center;
  position: relative;
  width: 80%;
  margin: 0 auto;
}

.slider-image-container {
  position: relative;
  width: 100%;
  padding-bottom: 56.25%; /* 默认16:9的宽高比 */
  overflow: hidden;
}

.slider-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: contain;
  background-color: #f5f5f5; /* 添加背景色，使空白区域不那么突兀 */
}

.prev-button,
.next-button {
  position: absolute;
  top: 0;
  height: 100%;
  width: 10%;
  background-color: rgba(0, 0, 0, 0.3);
  color: white;
  font-size: 24px;
  border: none;
  cursor: pointer;
  opacity: 0.2;
  z-index: 1;
  transition: opacity 0.3s ease;
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

.image-indicators {
  margin-top: 10px;
  position: relative;
  z-index: 2;
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
  margin-bottom: 50px;
  }

.image-indicators span.active {
  background-color: #4CAF50;
}

.box {
  display: inline-block;
  text-align: center;
  width: 100%;
  .image {
    width: 80%;
    margin: 0 auto;
  }
  .box1 {
    text-align: left;
    width: 80%;
    background-color: #ecf5ff;
    margin: 0 auto;
    margin-top: 20px;
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
      .content1 {
        width: 50%;
        height: 100%;
        float: left;
        .title {
          font-size: 15px;
          margin-bottom: 10px;
        }
        .bottom {
          margin-bottom: 10px;
        }
        .text {
          margin-bottom: 10px;
          width: 85%;
          font-size: 12px;
        }
        .li {
          margin-bottom: 10px;
          font-size: 12px;
        }
      }
    }
    .box4 {
      width: 100%;
      height: 300px;
      .content2 {
        float: left;
        width: 25%;
        height: 100%;
        .title {
          font-size: 8px;
        }
      }
    }
  }
}
</style>
