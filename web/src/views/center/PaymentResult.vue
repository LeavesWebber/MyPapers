<template>
  <div class="payment-result-container">
    <el-card class="result-card">
      <div slot="header" class="card-header">
        <span>支付结果</span>
      </div>
      <div v-if="loading" class="loading">
        <i class="el-icon-loading"></i>
        <span>正在处理支付结果...</span>
      </div>
      <div v-else-if="success" class="success">
        <i class="el-icon-success"></i>
        <h2>支付成功</h2>
        <p>您已成功充值 {{ mpsAmount }} MPS</p>
        <p>当前 MPS 余额：{{ mpsBalance }} MPS</p>
        <el-button type="primary" @click="goBack">返回</el-button>
      </div>
      <div v-else class="error">
        <i class="el-icon-error"></i>
        <h2>支付失败</h2>
        <p>{{ errorMessage }}</p>
        <el-button type="primary" @click="goBack">返回</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
import { MPScontractInstance } from '@/constant'
import Web3 from 'web3'
import axios from 'axios'

export default {
  name: 'PaymentResult',
  data() {
    return {
      loading: true,
      success: false,
      errorMessage: '',
      mpsAmount: 0,
      mpsBalance: '0',
      orderNo: '',
      walletAddress: '',
      pollInterval: null,
      isMinting: false
    }
  },
  created() {
    // 从 URL 获取订单号
    this.orderNo = this.$route.query.order_no
    // 从本地存储获取钱包地址和充值金额
    this.walletAddress = localStorage.getItem('current_wallet_address')
    this.mpsAmount = localStorage.getItem('current_mps_amount')
    
    if (!this.orderNo || !this.walletAddress || !this.mpsAmount) {
      this.loading = false
      this.success = false
      this.errorMessage = '支付信息不完整，请重新充值'
      return
    }

    // 开始轮询订单状态
    this.pollOrderStatus()
  },
  beforeDestroy() {
    // 组件销毁前清除定时器
    if (this.pollInterval) {
      clearInterval(this.pollInterval)
    }
  },
  methods: {
    async pollOrderStatus() {
      // 每 5 秒检查一次订单状态
      this.pollInterval = setInterval(async () => {
        try {
          // 从 localStorage 获取 token
          const token = localStorage.getItem('token')
          if (!token) {
            this.loading = false
            this.success = false
            this.errorMessage = '未登录，请重新登录'
            clearInterval(this.pollInterval)
            return
          }

          // 调用后端接口检查订单状态
          const res = await axios.get(`/mypapers/mps/orderStatus?order_no=${this.orderNo}`, {
            headers: {
              'Authorization': `Bearer ${token}`
            }
          })
          
          if (res.data.code === 0) {
            if (res.data.data.status === 1) { // 1-支付成功
              this.loading = false
              this.success = true
              
              // 只有在未开始发放代币时才尝试发放
              if (!this.isMinting) {
                this.isMinting = true
                // 调用智能合约发放代币
                const success = await this.mintMPS(this.walletAddress, this.mpsAmount)
                if (success) {
                  // 更新余额显示
                  await this.fetchMPSBalance()
                } else {
                  this.success = false
                  this.errorMessage = '代币发放失败，请联系客服'
                }
                this.isMinting = false
              }
            } else if (res.data.data.status === 2) { // 2-支付失败
              this.loading = false
              this.success = false
              this.errorMessage = '支付失败，请重试'
            }
          } else {
            this.loading = false
            this.success = false
            this.errorMessage = res.data.msg || '检查订单状态失败'
          }
          
          // 如果支付成功或失败，清除定时器
          if (res.data.data.status === 1 || res.data.data.status === 2) {
            clearInterval(this.pollInterval)
          }
        } catch (error) {
          console.error('检查订单状态失败:', error)
          this.loading = false
          this.success = false
          this.errorMessage = '检查订单状态失败，请稍后重试'
          clearInterval(this.pollInterval)
        }
      }, 5000)
    },
    async fetchMPSBalance() {
      try {
        if (!window.ethereum) {
          this.$message.error('请安装 MetaMask 钱包')
          return
        }

        const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' })
        if (!accounts || accounts.length === 0) {
          this.$message.error('请先连接 MetaMask 钱包')
          return
        }

        const balance = await MPScontractInstance.methods.balanceOf(accounts[0]).call()
        // 将 wei 转换为 ether 单位
        this.mpsBalance = Web3.utils.fromWei(balance, 'ether')
      } catch (error) {
        console.error('获取MPS余额失败:', error)
        this.$message.error('获取余额失败，请检查钱包连接')
      }
    },
    // 调用智能合约发放代币
    async mintMPS(toAddress, amount) {
      try {
        // 获取代币精度
        const decimals = await MPScontractInstance.methods.decimals().call()
        // 将金额转换为代币的最小单位
        const tokenAmount = Web3.utils.toWei(amount.toString(), 'ether')
        
        // 调用智能合约的 mint 函数
        const result = await MPScontractInstance.methods.mint([toAddress], tokenAmount).send({
          from: window.ethereum.selectedAddress,
          // 移除 gasPrice 设置，让 MetaMask 自动计算
        })
        
        console.log('代币发放成功:', result)
        return true
      } catch (error) {
        console.error('代币发放失败:', error)
        this.$message.error('代币发放失败: ' + (error.message || '未知错误'))
        return false
      }
    },
    goBack() {
      // 清除本地存储的支付信息
      localStorage.removeItem('current_order_no')
      localStorage.removeItem('current_wallet_address')
      localStorage.removeItem('current_mps_amount')
      
      // 返回充值页面
      this.$router.push('/center/recharge')
    }
  }
}
</script>

<style scoped>
.payment-result-container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.result-card {
  margin-bottom: 20px;
}

.card-header {
  font-size: 18px;
  font-weight: bold;
}

.loading {
  text-align: center;
  padding: 40px 0;
}

.loading .el-icon-loading {
  font-size: 40px;
  margin-bottom: 20px;
}

.success, .error {
  text-align: center;
  padding: 40px 0;
}

.success .el-icon-success, .error .el-icon-error {
  font-size: 60px;
  margin-bottom: 20px;
}

.success .el-icon-success {
  color: #67C23A;
}

.error .el-icon-error {
  color: #F56C6C;
}

h2 {
  margin-bottom: 20px;
}

p {
  margin-bottom: 10px;
  color: #666;
}
</style> 