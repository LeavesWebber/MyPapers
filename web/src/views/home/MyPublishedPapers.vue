<template>
  <div class="my-papers">
    <div class="container">
      <h1>My Published Papers</h1>
      
      <!-- 搜索和批量操作区域 -->
      <div class="operation-area">
        <div class="search-box">
          <el-input
            v-model="searchQuery"
            placeholder="Search by title, authors or venue"
            @keyup.enter.native="handleSearch"
            clearable
            @clear="handleSearch">
            <el-button slot="append" icon="el-icon-search" @click="handleSearch"></el-button>
          </el-input>
        </div>
        
        <div class="batch-actions" v-show="selectedPapers.length > 0">
          <el-button 
            type="primary" 
            icon="el-icon-download"
            @click="handleBatchDownload"
            :loading="downloading">
            Download Selected ({{ selectedPapers.length }})
          </el-button>
        </div>
      </div>

      <!-- 论文列表 -->
      <el-table
        ref="paperTable"
        v-loading="loading"
        :data="papers"
        @selection-change="handleSelectionChange"
        style="width: 100%; margin-top: 20px;">
        
        <el-table-column type="selection" width="55"></el-table-column>
        
        <el-table-column prop="title" label="Title" min-width="250">
          <template slot-scope="scope">
            <div class="paper-title">
              <el-tooltip :content="scope.row.title" placement="top" :disabled="scope.row.title.length < 50">
                <span>{{ scope.row.title }}</span>
              </el-tooltip>
            </div>
          </template>
        </el-table-column>
        
        <el-table-column prop="authors" label="Authors" min-width="200">
          <template slot-scope="scope">
            <el-tooltip :content="scope.row.authors" placement="top">
              <span>{{ scope.row.authors }}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        
        <el-table-column prop="venue" label="Venue" min-width="150">
          <template slot-scope="scope">
            <span>{{ scope.row.venueType === 'journal' ? 'Journal' : 'Conference' }}:</span>
            <span>{{ scope.row.venueName }}</span>
          </template>
        </el-table-column>
        
        <el-table-column prop="publicationDate" label="Publication Date" width="120">
          <template slot-scope="scope">
            {{ formatDate(scope.row.publicationDate) }}
          </template>
        </el-table-column>
        
        <el-table-column prop="status" label="Status" width="120">
          <template slot-scope="scope">
            <el-tag :type="getStatusType(scope.row.status)">
              {{ scope.row.status }}
            </el-tag>
          </template>
        </el-table-column>
        
        <el-table-column label="NFT" width="100">
          <template slot-scope="scope">
            <el-button 
              type="text" 
              @click="viewNFT(scope.row)"
              :disabled="!scope.row.nftGenerated">
              {{ scope.row.nftGenerated ? 'View' : 'Pending' }}
            </el-button>
          </template>
        </el-table-column>
        
        <el-table-column label="Operations" width="150">
          <template slot-scope="scope">
            <el-button 
              type="text" 
              icon="el-icon-download"
              @click="handleDownload(scope.row)">
              Download
            </el-button>
            <el-button 
              type="text" 
              icon="el-icon-view"
              @click="handleView(scope.row)">
              View
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination-container">
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          :page-sizes="[10, 20, 50, 100]"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total">
        </el-pagination>
      </div>
    </div>

    <!-- NFT详情对话框 -->
    <el-dialog
      title="NFT Certificate Details"
      :visible.sync="nftDialogVisible"
      width="50%">
      <div v-if="currentNFT" class="nft-details">
        <div class="nft-info">
          <p><strong>Token ID:</strong> {{ currentNFT.tokenId }}</p>
          <p><strong>Contract Address:</strong> {{ currentNFT.contractAddress }}</p>
          <p><strong>Creation Date:</strong> {{ formatDate(currentNFT.creationDate) }}</p>
          <p><strong>Blockchain:</strong> {{ currentNFT.blockchain }}</p>
        </div>
        <div class="nft-image">
          <img :src="currentNFT.imageUrl" alt="NFT Certificate">
        </div>
        <div class="nft-actions">
          <el-button type="primary" @click="viewOnBlockchain">View on Blockchain</el-button>
          <el-button type="success" @click="downloadCertificate">Download Certificate</el-button>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'MyPublishedPapers',
  data() {
    return {
      searchQuery: '',
      papers: [],
      selectedPapers: [],
      loading: false,
      downloading: false,
      currentPage: 1,
      pageSize: 10,
      total: 0,
      nftDialogVisible: false,
      currentNFT: null
    }
  },
  created() {
    this.fetchPapers()
  },
  methods: {
    async fetchPapers() {
      this.loading = true
      try {
        const response = await this.$http.get('/api/published-papers', {
          params: {
            page: this.currentPage,
            pageSize: this.pageSize,
            query: this.searchQuery
          }
        })
        
        if (response.data.code === 1000) {
          this.papers = response.data.data.papers
          this.total = response.data.data.total
        } else {
          this.$message.error(response.data.msg || 'Failed to fetch papers')
        }
      } catch (error) {
        this.$message.error('Failed to fetch papers: ' + error.message)
      } finally {
        this.loading = false
      }
    },
    
    handleSearch() {
      this.currentPage = 1
      this.fetchPapers()
    },
    
    handleSelectionChange(selection) {
      this.selectedPapers = selection
    },
    
    async handleBatchDownload() {
      if (this.selectedPapers.length === 0) {
        this.$message.warning('Please select papers to download')
        return
      }

      this.downloading = true
      try {
        const paperIds = this.selectedPapers.map(paper => paper.id)
        const response = await this.$http.post('/api/published-papers/batch-download', {
          paperIds
        }, {
          responseType: 'blob'
        })
        
        // 创建下载链接
        const url = window.URL.createObjectURL(new Blob([response.data]))
        const link = document.createElement('a')
        link.href = url
        link.setAttribute('download', 'selected_papers.zip')
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        window.URL.revokeObjectURL(url)
        
        this.$message.success('Papers downloaded successfully')
      } catch (error) {
        this.$message.error('Failed to download papers: ' + error.message)
      } finally {
        this.downloading = false
      }
    },
    
    async handleDownload(paper) {
      try {
        const response = await this.$http.get(`/api/published-papers/${paper.id}/download`, {
          responseType: 'blob'
        })
        
        const url = window.URL.createObjectURL(new Blob([response.data]))
        const link = document.createElement('a')
        link.href = url
        link.setAttribute('download', `${paper.title}.pdf`)
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        window.URL.revokeObjectURL(url)
        
        this.$message.success('Paper downloaded successfully')
      } catch (error) {
        this.$message.error('Failed to download paper: ' + error.message)
      }
    },
    
    handleView(paper) {
      window.open(`/api/published-papers/${paper.id}/view`, '_blank')
    },
    
    async viewNFT(paper) {
      try {
        const response = await this.$http.get(`/api/published-papers/${paper.id}/nft`)
        if (response.data.code === 1000) {
          this.currentNFT = response.data.data
          this.nftDialogVisible = true
        } else {
          this.$message.error(response.data.msg || 'Failed to fetch NFT details')
        }
      } catch (error) {
        this.$message.error('Failed to fetch NFT details: ' + error.message)
      }
    },
    
    viewOnBlockchain() {
      if (this.currentNFT && this.currentNFT.blockchainUrl) {
        window.open(this.currentNFT.blockchainUrl, '_blank')
      }
    },
    
    downloadCertificate() {
      if (this.currentNFT && this.currentNFT.certificateUrl) {
        window.open(this.currentNFT.certificateUrl, '_blank')
      }
    },
    
    handleSizeChange(val) {
      this.pageSize = val
      this.fetchPapers()
    },
    
    handleCurrentChange(val) {
      this.currentPage = val
      this.fetchPapers()
    },
    
    formatDate(date) {
      if (!date) return ''
      const d = new Date(date)
      return d.toLocaleDateString()
    },
    
    getStatusType(status) {
      const statusMap = {
        'pending': 'warning',
        'verified': 'success',
        'rejected': 'danger'
      }
      return statusMap[status] || 'info'
    }
  }
}
</script>

<style lang="less" scoped>
.my-papers {
  padding: 40px 0;
  
  .container {
    width: 90%;
    max-width: 1400px;
    margin: 0 auto;
    background-color: #fff;
    padding: 40px;
    border-radius: 8px;
    box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);

    h1 {
      color: #2c3e50;
      margin-bottom: 30px;
      text-align: center;
    }
  }

  .operation-area {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;

    .search-box {
      width: 300px;
    }

    .batch-actions {
      display: flex;
      gap: 10px;
    }
  }

  .paper-title {
    font-weight: bold;
  }

  .pagination-container {
    margin-top: 20px;
    text-align: right;
  }
}

.nft-details {
  .nft-info {
    margin-bottom: 20px;
    
    p {
      margin: 10px 0;
    }
  }

  .nft-image {
    text-align: center;
    margin: 20px 0;
    
    img {
      max-width: 100%;
      max-height: 400px;
      object-fit: contain;
    }
  }

  .nft-actions {
    text-align: center;
    margin-top: 20px;
  }
}
</style> 