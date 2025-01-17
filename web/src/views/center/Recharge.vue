<template>
  <div class="recharge-container">
    <el-card class="recharge-card">
      <div slot="header" class="card-header">
        <span>充值 MPS</span>
      </div>
      <div class="balance-info">
        <p>当前 MPS 余额：{{ mpsBalance }} MPS</p>
        <p>兑换比例：1 CNY = 1 MPS</p>
      </div>
      <el-form :model="rechargeForm" :rules="rules" ref="rechargeForm" label-width="100px" class="recharge-form">
        <el-form-item label="充值金额" prop="amount">
          <el-input-number 
            v-model="rechargeForm.amount" 
            :min="1"
            :max="10000"
            controls-position="right"
            @change="handleAmountChange">
          </el-input-number>
          <span class="unit">CNY</span>
        </el-form-item>
        <el-form-item label="获得MPS">
          <span class="mps-amount">{{ rechargeForm.amount }} MPS</span>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleRecharge" :loading="loading">微信支付充值</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { buyMPSWithFiat } from '@/api'
import { MPScontractInstance } from '@/constant'
import Web3 from 'web3'

export default {
  name: 'Recharge',
  data() {
    return {
      mpsBalance: '0',
      loading: false,
      rechargeForm: {
        amount: 100
      },
      rules: {
        amount: [
          { required: true, message: '请输入充值金额', trigger: 'blur' },
          { type: 'number', min: 1, max: 10000, message: '充值金额必须在1-10000之间', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.fetchMPSBalance()
    // 监听钱包账户变化
    if (window.ethereum) {
      window.ethereum.on('accountsChanged', (accounts) => {
        this.fetchMPSBalance()
      })
    }
  },
  methods: {
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
    handleAmountChange(value) {
      this.rechargeForm.amount = value
    },
    async handleRecharge() {
      try {
        this.$refs.rechargeForm.validate(async (valid) => {
          if (valid) {
            this.loading = true
            const res = await buyMPSWithFiat({
              amount: this.rechargeForm.amount
            })
            
            if (res.data.code === 1000) {
              // 假设后端返回微信支付所需的参数
              const payParams = res.data.data
              // 调用微信支付
              this.callWxPay(payParams)
            } else {
              this.$message.error(res.data.msg || '充值失败，请重试')
            }
          }
        })
      } catch (error) {
        console.error('充值失败:', error)
        this.$message.error('充值失败，请稍后重试')
      } finally {
        this.loading = false
      }
    },
    callWxPay(payParams) {
      // 这里需要集成微信支付SDK
      // 支付成功后刷新余额
      this.fetchMPSBalance()
    }
  }
}
</script>

<style scoped>
.recharge-container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.recharge-card {
  margin-bottom: 20px;
}

.card-header {
  font-size: 18px;
  font-weight: bold;
}

.balance-info {
  margin-bottom: 20px;
  color: #666;
}

.recharge-form {
  max-width: 500px;
  margin: 0 auto;
}

.unit {
  margin-left: 10px;
  color: #666;
}

.mps-amount {
  font-size: 16px;
  color: #409EFF;
  font-weight: bold;
}
</style> 