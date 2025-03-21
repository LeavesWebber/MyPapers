<template>
  <div class="published-papers">
    <div class="container">
      <div class="header">
        <h1>Published Papers</h1>
        <p class="description">
          Secure your academic achievements with blockchain technology. Upload your published papers 
          and receive NFT certificates to protect your intellectual property rights.
        </p>
      </div>

      <div class="features">
        <el-row :gutter="30">
          <el-col :span="8">
            <div class="feature-card">
              <div class="icon">
                <i class="el-icon-upload"></i>
              </div>
              <h3>Upload Papers</h3>
              <p>Upload your published academic papers with verified email authentication.</p>
              <el-button 
                type="primary" 
                @click="$router.push('/published-papers/upload')"
                :disabled="!isLoggedIn">
                Upload Now
              </el-button>
            </div>
          </el-col>

          <el-col :span="8">
            <div class="feature-card">
              <div class="icon">
                <i class="el-icon-document"></i>
              </div>
              <h3>My Papers</h3>
              <p>View and manage your uploaded papers, download NFT certificates.</p>
              <el-button 
                type="primary" 
                @click="$router.push('/published-papers/my-papers')"
                :disabled="!isLoggedIn">
                View Papers
              </el-button>
            </div>
          </el-col>

          <el-col :span="8">
            <div class="feature-card">
              <div class="icon">
                <i class="el-icon-info"></i>
              </div>
              <h3>How It Works</h3>
              <p>Learn about our paper verification and NFT certification process.</p>
              <el-button type="primary" @click="showGuide">
                Read Guide
              </el-button>
            </div>
          </el-col>
        </el-row>
      </div>

      <div class="process-section">
        <h2>Publication Process</h2>
        <el-steps :active="4" finish-status="success" align-center>
          <el-step title="Upload Paper" description="Upload your published paper with details"></el-step>
          <el-step title="Email Verification" description="Verify your author email"></el-step>
          <el-step title="Review" description="Our system verifies the paper"></el-step>
          <el-step title="NFT Generation" description="Receive your NFT certificate"></el-step>
        </el-steps>
      </div>

      <div class="benefits-section">
        <h2>Benefits</h2>
        <el-row :gutter="20">
          <el-col :span="6">
            <div class="benefit-item">
              <i class="el-icon-lock"></i>
              <h4>Secure</h4>
              <p>Blockchain-based protection for your intellectual property</p>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="benefit-item">
              <i class="el-icon-time"></i>
              <h4>Instant</h4>
              <p>Quick verification and NFT generation process</p>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="benefit-item">
              <i class="el-icon-connection"></i>
              <h4>Traceable</h4>
              <p>Permanent record on blockchain network</p>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="benefit-item">
              <i class="el-icon-share"></i>
              <h4>Shareable</h4>
              <p>Easy to share and verify ownership</p>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>

    <!-- Guide Dialog -->
    <el-dialog
      title="How It Works"
      :visible.sync="guideDialogVisible"
      width="60%">
      <div class="guide-content">
        <h3>Paper Upload and NFT Certification Process</h3>
        
        <div class="guide-section">
          <h4>1. Prepare Your Paper</h4>
          <ul>
            <li>Ensure your paper is already published in a recognized journal or conference</li>
            <li>Have the PDF version of the paper ready</li>
            <li>Gather all necessary publication details</li>
          </ul>
        </div>

        <div class="guide-section">
          <h4>2. Upload and Verification</h4>
          <ul>
            <li>Fill in the paper details including title, authors, and publication venue</li>
            <li>Upload the PDF file (max 50MB)</li>
            <li>Verify your email as the corresponding author</li>
          </ul>
        </div>

        <div class="guide-section">
          <h4>3. NFT Generation</h4>
          <ul>
            <li>Our system verifies the paper details</li>
            <li>Creates a unique NFT certificate</li>
            <li>Records the ownership on the blockchain</li>
          </ul>
        </div>

        <div class="guide-section">
          <h4>4. Management</h4>
          <ul>
            <li>View your papers and NFT certificates</li>
            <li>Download certificates and share with others</li>
            <li>Track paper verification status</li>
          </ul>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'PublishedPapers',
  data() {
    return {
      guideDialogVisible: false
    }
  },
  computed: {
    ...mapGetters(['isLoggedIn'])
  },
  methods: {
    showGuide() {
      this.guideDialogVisible = true
    }
  }
}
</script>

<style lang="less" scoped>
.published-papers {
  padding: 40px 0;
  background-color: #f5f7fa;
  min-height: calc(100vh - 60px);

  .container {
    width: 90%;
    max-width: 1400px;
    margin: 0 auto;

    .header {
      text-align: center;
      margin-bottom: 60px;

      h1 {
        color: #2c3e50;
        font-size: 36px;
        margin-bottom: 20px;
      }

      .description {
        color: #606266;
        font-size: 18px;
        max-width: 800px;
        margin: 0 auto;
      }
    }

    .features {
      margin-bottom: 80px;

      .feature-card {
        background: #fff;
        border-radius: 8px;
        padding: 30px;
        text-align: center;
        height: 100%;
        box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);
        transition: transform 0.3s ease;

        &:hover {
          transform: translateY(-5px);
        }

        .icon {
          font-size: 48px;
          color: #409EFF;
          margin-bottom: 20px;
        }

        h3 {
          color: #2c3e50;
          margin-bottom: 15px;
        }

        p {
          color: #606266;
          margin-bottom: 20px;
          min-height: 40px;
        }
      }
    }

    .process-section {
      background: #fff;
      padding: 40px;
      border-radius: 8px;
      margin-bottom: 60px;
      box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);

      h2 {
        color: #2c3e50;
        text-align: center;
        margin-bottom: 40px;
      }
    }

    .benefits-section {
      margin-bottom: 60px;

      h2 {
        color: #2c3e50;
        text-align: center;
        margin-bottom: 40px;
      }

      .benefit-item {
        background: #fff;
        padding: 30px;
        border-radius: 8px;
        text-align: center;
        height: 100%;
        box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);

        i {
          font-size: 36px;
          color: #409EFF;
          margin-bottom: 15px;
        }

        h4 {
          color: #2c3e50;
          margin-bottom: 10px;
        }

        p {
          color: #606266;
          font-size: 14px;
        }
      }
    }
  }
}

.guide-content {
  h3 {
    color: #2c3e50;
    margin-bottom: 30px;
    text-align: center;
  }

  .guide-section {
    margin-bottom: 30px;

    h4 {
      color: #409EFF;
      margin-bottom: 15px;
    }

    ul {
      list-style: none;
      padding-left: 20px;

      li {
        color: #606266;
        margin-bottom: 10px;
        position: relative;

        &:before {
          content: "•";
          color: #409EFF;
          position: absolute;
          left: -15px;
        }
      }
    }
  }
}
</style> 