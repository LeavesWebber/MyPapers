<template>
  <div class="nft-mint-container">
    <el-card class="mint-card">
      <div slot="header" class="card-header">
        <span>铸造NFT</span>
      </div>
      
      <el-form ref="mintForm" :model="mintForm" :rules="rules" label-width="120px">
        <el-form-item label="接收地址" prop="to">
          <el-input 
            v-model="mintForm.to" 
            placeholder="请输入接收NFT的以太坊地址"
            style="width: 400px"
          ></el-input>
        </el-form-item>
        
        <el-form-item label="Token ID" prop="tokenId">
          <el-input-number 
            v-model="mintForm.tokenId" 
            :min="0"
            placeholder="请输入Token ID"
            style="width: 200px"
          ></el-input-number>
        </el-form-item>
        
        <el-form-item label="元数据URI" prop="uri">
          <el-input 
            v-model="mintForm.uri" 
            placeholder="请输入NFT元数据URI (如: ipfs://QmXYZ...)"
            style="width: 400px"
          ></el-input>
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="mintNFT" :loading="minting">
            {{ minting ? '铸造中...' : '铸造NFT' }}
          </el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
      
      <!-- 铸造结果 -->
      <div v-if="mintResult" class="mint-result">
        <el-alert
          :title="mintResult.success ? '铸造成功' : '铸造失败'"
          :type="mintResult.success ? 'success' : 'error'"
          :description="mintResult.message"
          show-icon
        >
          <div v-if="mintResult.success && mintResult.txHash" slot="description">
            <p>交易哈希: {{ mintResult.txHash }}</p>
            <p>Token ID: {{ mintResult.tokenId }}</p>
            <p>接收地址: {{ mintResult.to }}</p>
            <el-button size="small" @click="viewOnExplorer(mintResult.txHash)">
              在区块浏览器中查看
            </el-button>
          </div>
        </el-alert>
      </div>
    </el-card>
    
    <!-- NFT预览 -->
    <el-card class="preview-card" v-if="mintForm.uri">
      <div slot="header" class="card-header">
        <span>NFT预览</span>
      </div>
      <div class="nft-preview">
        <div class="nft-image">
          <img v-if="previewImage" :src="previewImage" alt="NFT Preview" />
          <div v-else class="no-image">暂无图片</div>
        </div>
        <div class="nft-info">
          <h3>MyNFT #{{ mintForm.tokenId }}</h3>
          <p>接收地址: {{ mintForm.to }}</p>
          <p>元数据URI: {{ mintForm.uri }}</p>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script>
import { mintNFT } from '@/api/nft'

export default {
  name: 'NFTMint',
  data() {
    return {
      mintForm: {
        to: '',
        tokenId: 0,
        uri: ''
      },
      rules: {
        to: [
          { required: true, message: '请输入接收地址', trigger: 'blur' },
          { pattern: /^0x[a-fA-F0-9]{40}$/, message: '请输入有效的以太坊地址', trigger: 'blur' }
        ],
        tokenId: [
          { required: true, message: '请输入Token ID', trigger: 'blur' },
          { type: 'number', min: 0, message: 'Token ID必须大于等于0', trigger: 'blur' }
        ],
        uri: [
          { required: true, message: '请输入元数据URI', trigger: 'blur' }
        ]
      },
      minting: false,
      mintResult: null,
      previewImage: null
    }
  },
  watch: {
    'mintForm.uri': {
      handler(newUri) {
        this.loadPreview(newUri)
      },
      immediate: false
    }
  },
  methods: {
    async mintNFT() {
      try {
        await this.$refs.mintForm.validate()
        
        this.minting = true
        this.mintResult = null
        
        const response = await mintNFT(this.mintForm)
        
        this.mintResult = {
          success: true,
          message: 'NFT铸造成功！',
          txHash: response.data.txHash,
          tokenId: response.data.tokenId,
          to: response.data.to
        }
        
        this.$message.success('NFT铸造成功！')
        
        // 重置表单
        this.resetForm()
        
      } catch (error) {
        console.error('铸造NFT失败:', error)
        
        this.mintResult = {
          success: false,
          message: error.response?.data?.message || error.message || '铸造失败'
        }
        
        this.$message.error('铸造NFT失败: ' + this.mintResult.message)
      } finally {
        this.minting = false
      }
    },
    
    resetForm() {
      this.$refs.mintForm.resetFields()
      this.mintResult = null
      this.previewImage = null
    },
    
    async loadPreview(uri) {
      if (!uri) {
        this.previewImage = null
        return
      }
      
      try {
        // 如果是IPFS URI，转换为HTTP URL
        if (uri.startsWith('ipfs://')) {
          const ipfsHash = uri.replace('ipfs://', '')
          this.previewImage = `https://ipfs.io/ipfs/${ipfsHash}`
        } else if (uri.startsWith('http')) {
          this.previewImage = uri
        } else {
          this.previewImage = null
        }
      } catch (error) {
        console.error('加载预览失败:', error)
        this.previewImage = null
      }
    },
    
    viewOnExplorer(txHash) {
      // 根据网络打开对应的区块浏览器
      const network = this.$store.state.network || 'localhost'
      let explorerUrl = ''
      
      if (network === 'localhost') {
        explorerUrl = `http://localhost:8545/tx/${txHash}`
      } else if (network === 'sepolia') {
        explorerUrl = `https://sepolia.etherscan.io/tx/${txHash}`
      } else if (network === 'mainnet') {
        explorerUrl = `https://etherscan.io/tx/${txHash}`
      }
      
      if (explorerUrl) {
        window.open(explorerUrl, '_blank')
      }
    }
  },
  
  mounted() {
    // 设置默认接收地址为当前用户地址
    const userInfo = JSON.parse(localStorage.getItem('userInfo') || '{}')
    if (userInfo.block_chain_address) {
      this.mintForm.to = userInfo.block_chain_address
    }
  }
}
</script>

<style lang="less" scoped>
.nft-mint-container {
  padding: 20px;
  
  .mint-card {
    margin-bottom: 20px;
    
    .card-header {
      font-size: 18px;
      font-weight: bold;
    }
    
    .mint-result {
      margin-top: 20px;
    }
  }
  
  .preview-card {
    .card-header {
      font-size: 16px;
      font-weight: bold;
    }
    
    .nft-preview {
      display: flex;
      gap: 20px;
      
      .nft-image {
        width: 200px;
        height: 200px;
        border: 1px solid #ddd;
        border-radius: 8px;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
        
        img {
          width: 100%;
          height: 100%;
          object-fit: cover;
        }
        
        .no-image {
          color: #999;
          font-size: 14px;
        }
      }
      
      .nft-info {
        flex: 1;
        
        h3 {
          margin: 0 0 10px 0;
          color: #333;
        }
        
        p {
          margin: 5px 0;
          color: #666;
          font-size: 14px;
        }
      }
    }
  }
}
</style> 