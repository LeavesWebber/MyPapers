<template>
  <div class="nft-marketplace-container">
    <!-- 搜索和筛选 -->
    <el-card class="filter-card">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-input
            v-model="searchQuery"
            placeholder="搜索NFT..."
            prefix-icon="el-icon-search"
            @input="handleSearch"
          ></el-input>
        </el-col>
        <el-col :span="4">
          <el-select v-model="sortBy" placeholder="排序方式" @change="handleSort">
            <el-option label="最新" value="createdAt"></el-option>
            <el-option label="价格从低到高" value="priceAsc"></el-option>
            <el-option label="价格从高到低" value="priceDesc"></el-option>
            <el-option label="最热门" value="popular"></el-option>
          </el-select>
        </el-col>
        <el-col :span="4">
          <el-select v-model="category" placeholder="分类" @change="handleFilter">
            <el-option label="全部" value=""></el-option>
            <el-option label="论文证书" value="paper"></el-option>
            <el-option label="研究奖项" value="award"></el-option>
            <el-option label="会议徽章" value="badge"></el-option>
          </el-select>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="refreshData">刷新</el-button>
        </el-col>
      </el-row>
    </el-card>

    <!-- NFT列表 -->
    <el-row :gutter="20" class="nft-grid">
      <el-col 
        :span="6" 
        v-for="nft in nftList" 
        :key="nft.tokenId"
        class="nft-item"
      >
        <el-card class="nft-card" shadow="hover">
          <div class="nft-image">
            <img :src="nft.image || '/default-nft.png'" :alt="nft.name" />
            <div class="nft-overlay">
              <el-button type="primary" size="small" @click="viewNFT(nft)">
                查看详情
              </el-button>
              <el-button 
                v-if="nft.isForSale" 
                type="success" 
                size="small" 
                @click="buyNFT(nft)"
              >
                购买
              </el-button>
            </div>
          </div>
          <div class="nft-info">
            <h3>{{ nft.name }}</h3>
            <p class="nft-owner">所有者: {{ formatAddress(nft.owner) }}</p>
            <p class="nft-price" v-if="nft.isForSale">
              价格: {{ nft.price }} ETH
            </p>
            <p class="nft-status" v-else>
              <span class="not-for-sale">未出售</span>
            </p>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 分页 -->
    <div class="pagination-container">
      <el-pagination
        @current-change="handlePageChange"
        :current-page="currentPage"
        :page-size="pageSize"
        :total="total"
        layout="total, prev, pager, next, jumper"
      ></el-pagination>
    </div>

    <!-- NFT详情对话框 -->
    <el-dialog
      title="NFT详情"
      :visible.sync="nftDetailVisible"
      width="600px"
    >
      <div v-if="selectedNFT" class="nft-detail">
        <div class="detail-image">
          <img :src="selectedNFT.image || '/default-nft.png'" :alt="selectedNFT.name" />
        </div>
        <div class="detail-info">
          <h2>{{ selectedNFT.name }}</h2>
          <p><strong>Token ID:</strong> {{ selectedNFT.tokenId }}</p>
          <p><strong>所有者:</strong> {{ formatAddress(selectedNFT.owner) }}</p>
          <p><strong>创建时间:</strong> {{ formatDate(selectedNFT.createdAt) }}</p>
          <p v-if="selectedNFT.isForSale">
            <strong>价格:</strong> {{ selectedNFT.price }} ETH
          </p>
          <p><strong>描述:</strong> {{ selectedNFT.description || '暂无描述' }}</p>
          
          <div class="detail-actions">
            <el-button 
              v-if="selectedNFT.isForSale" 
              type="primary" 
              @click="buyNFT(selectedNFT)"
            >
              购买NFT
            </el-button>
            <el-button @click="viewOnExplorer(selectedNFT.tokenId)">
              在区块浏览器中查看
            </el-button>
          </div>
        </div>
      </div>
    </el-dialog>

    <!-- 购买确认对话框 -->
    <el-dialog
      title="确认购买"
      :visible.sync="buyDialogVisible"
      width="400px"
    >
      <div v-if="buyingNFT" class="buy-confirmation">
        <p>您确定要购买以下NFT吗？</p>
        <div class="buy-nft-info">
          <img :src="buyingNFT.image || '/default-nft.png'" :alt="buyingNFT.name" />
          <div>
            <h3>{{ buyingNFT.name }}</h3>
            <p>价格: {{ buyingNFT.price }} ETH</p>
            <p>Token ID: {{ buyingNFT.tokenId }}</p>
          </div>
        </div>
        <div class="buy-actions">
          <el-button @click="buyDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmBuy" :loading="buying">
            {{ buying ? '购买中...' : '确认购买' }}
          </el-button>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { 
  getNFTMarketplace, 
  getNFTByTokenId, 
  buyNFT as buyNFTApi 
} from '@/api/nft'

export default {
  name: 'NFTMarketplace',
  data() {
    return {
      searchQuery: '',
      sortBy: 'createdAt',
      category: '',
      currentPage: 1,
      pageSize: 12,
      total: 0,
      nftList: [],
      loading: false,
      nftDetailVisible: false,
      selectedNFT: null,
      buyDialogVisible: false,
      buyingNFT: null,
      buying: false
    }
  },
  methods: {
    async loadNFTs() {
      try {
        this.loading = true
        
        const params = {
          page: this.currentPage,
          limit: this.pageSize,
          category: this.category,
          sortBy: this.sortBy,
          sortOrder: 'desc',
          q: this.searchQuery
        }
        
        const response = await getNFTMarketplace(params)
        this.nftList = response.data.items || []
        this.total = response.data.total || 0
        
      } catch (error) {
        console.error('加载NFT列表失败:', error)
        this.$message.error('加载NFT列表失败')
      } finally {
        this.loading = false
      }
    },
    
    handleSearch() {
      this.currentPage = 1
      this.loadNFTs()
    },
    
    handleSort() {
      this.currentPage = 1
      this.loadNFTs()
    },
    
    handleFilter() {
      this.currentPage = 1
      this.loadNFTs()
    },
    
    handlePageChange(page) {
      this.currentPage = page
      this.loadNFTs()
    },
    
    refreshData() {
      this.loadNFTs()
    },
    
    async viewNFT(nft) {
      try {
        const response = await getNFTByTokenId(nft.tokenId)
        this.selectedNFT = response.data
        this.nftDetailVisible = true
      } catch (error) {
        console.error('获取NFT详情失败:', error)
        this.$message.error('获取NFT详情失败')
      }
    },
    
    buyNFT(nft) {
      this.buyingNFT = nft
      this.buyDialogVisible = true
    },
    
    async confirmBuy() {
      try {
        this.buying = true
        
        const buyData = {
          tokenId: this.buyingNFT.tokenId,
          price: this.buyingNFT.price
        }
        
        const response = await buyNFTApi(buyData)
        
        this.$message.success('购买成功！')
        this.buyDialogVisible = false
        this.buyingNFT = null
        
        // 刷新列表
        this.loadNFTs()
        
      } catch (error) {
        console.error('购买NFT失败:', error)
        this.$message.error('购买NFT失败: ' + (error.response?.data?.message || error.message))
      } finally {
        this.buying = false
      }
    },
    
    formatAddress(address) {
      if (!address) return '未知'
      return address.substring(0, 6) + '...' + address.substring(address.length - 4)
    },
    
    formatDate(dateString) {
      if (!dateString) return '未知'
      return new Date(dateString).toLocaleString()
    },
    
    viewOnExplorer(tokenId) {
      const network = this.$store.state.network || 'localhost'
      let explorerUrl = ''
      
      if (network === 'localhost') {
        explorerUrl = `http://localhost:8545/token/${tokenId}`
      } else if (network === 'sepolia') {
        explorerUrl = `https://sepolia.etherscan.io/token/${tokenId}`
      } else if (network === 'mainnet') {
        explorerUrl = `https://etherscan.io/token/${tokenId}`
      }
      
      if (explorerUrl) {
        window.open(explorerUrl, '_blank')
      }
    }
  },
  
  mounted() {
    this.loadNFTs()
  }
}
</script>

<style lang="less" scoped>
.nft-marketplace-container {
  padding: 20px;
  
  .filter-card {
    margin-bottom: 20px;
  }
  
  .nft-grid {
    margin-bottom: 20px;
    
    .nft-item {
      margin-bottom: 20px;
      
      .nft-card {
        .nft-image {
          position: relative;
          height: 200px;
          overflow: hidden;
          
          img {
            width: 100%;
            height: 100%;
            object-fit: cover;
          }
          
          .nft-overlay {
            position: absolute;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background: rgba(0, 0, 0, 0.7);
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 10px;
            opacity: 0;
            transition: opacity 0.3s;
          }
          
          &:hover .nft-overlay {
            opacity: 1;
          }
        }
        
        .nft-info {
          padding: 10px;
          
          h3 {
            margin: 0 0 5px 0;
            font-size: 16px;
            color: #333;
          }
          
          .nft-owner {
            margin: 5px 0;
            font-size: 12px;
            color: #666;
          }
          
          .nft-price {
            margin: 5px 0;
            font-size: 14px;
            color: #409EFF;
            font-weight: bold;
          }
          
          .nft-status {
            margin: 5px 0;
            
            .not-for-sale {
              color: #F56C6C;
              font-size: 12px;
            }
          }
        }
      }
    }
  }
  
  .pagination-container {
    text-align: center;
    margin-top: 20px;
  }
  
  .nft-detail {
    display: flex;
    gap: 20px;
    
    .detail-image {
      width: 200px;
      height: 200px;
      
      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        border-radius: 8px;
      }
    }
    
    .detail-info {
      flex: 1;
      
      h2 {
        margin: 0 0 15px 0;
        color: #333;
      }
      
      p {
        margin: 8px 0;
        color: #666;
      }
      
      .detail-actions {
        margin-top: 20px;
        
        .el-button {
          margin-right: 10px;
        }
      }
    }
  }
  
  .buy-confirmation {
    .buy-nft-info {
      display: flex;
      gap: 15px;
      margin: 15px 0;
      padding: 15px;
      background: #f5f5f5;
      border-radius: 8px;
      
      img {
        width: 80px;
        height: 80px;
        object-fit: cover;
        border-radius: 4px;
      }
      
      h3 {
        margin: 0 0 5px 0;
        color: #333;
      }
      
      p {
        margin: 3px 0;
        color: #666;
      }
    }
    
    .buy-actions {
      text-align: right;
      margin-top: 20px;
      
      .el-button {
        margin-left: 10px;
      }
    }
  }
}
</style> 