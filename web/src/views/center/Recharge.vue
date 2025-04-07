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
        <el-form-item label="支付方式">
          <el-radio-group v-model="rechargeForm.pay_type">
            <el-radio label="alipay">支付宝支付</el-radio>
            <el-radio label="wxpay">微信支付</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleRecharge" :loading="loading">立即充值</el-button>
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
        amount: 100,
        pay_type: 'alipay'
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
          gasPrice: "0",
        })
        
        console.log('代币发放成功:', result)
        return true
      } catch (error) {
        console.error('代币发放失败:', error)
        this.$message.error('代币发放失败: ' + (error.message || '未知错误'))
        return false
      }
    },
    async handleRecharge() {
      try {
        this.$refs.rechargeForm.validate(async (valid) => {
          if (valid) {
            this.loading = true
            const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' })
            if (!accounts || accounts.length === 0) {
              this.$message.error('请先连接 MetaMask 钱包')
              this.loading = false
              return
            }

            const res = await buyMPSWithFiat({
              amount: this.rechargeForm.amount,
              wallet_address: accounts[0],
              pay_type: this.rechargeForm.pay_type
            })
            
            if (res.code === 0) {
              if (this.rechargeForm.pay_type === 'alipay') {
                // 支付宝支付
                if (res.data && res.data.pay_url) {
                  // 保存订单号和钱包地址到本地存储
                  localStorage.setItem('current_order_no', res.data.order_no)
                  localStorage.setItem('current_wallet_address', accounts[0])
                  localStorage.setItem('current_mps_amount', this.rechargeForm.amount)
                  
                  // 直接使用后端返回的支付URL，不要修改
                  console.log('跳转到支付页面:', res.data.pay_url)
                  window.location.href = res.data.pay_url
                } else {
                  this.$message.error('获取支付链接失败')
                }
              } else {
                // 微信支付
                if (res.data && res.data.wx_pay_params) {
                  this.callWxPay(res.data.wx_pay_params)
                } else {
                  this.$message.error('获取微信支付参数失败')
                }
              }
            } else {
              this.$message.error(res.msg || '充值失败，请重试')
            }
          }
        })
      } catch (error) {
        console.error('充值失败:', error)
        this.$message.error('充值失败: ' + (error.message || '未知错误'))
      } finally {
        this.loading = false
      }
    },
    callWxPay(payParams) {
      if (typeof WeixinJSBridge == "undefined") {
        if (document.addEventListener) {
          document.addEventListener('WeixinJSBridgeReady', this.onBridgeReady(payParams), false);
        } else if (document.attachEvent) {
          document.attachEvent('WeixinJSBridgeReady', this.onBridgeReady(payParams));
          document.attachEvent('onWeixinJSBridgeReady', this.onBridgeReady(payParams));
        }
      } else {
        this.onBridgeReady(payParams);
      }
    },
    onBridgeReady(payParams) {
      WeixinJSBridge.invoke(
        'getBrandWCPayRequest',
        payParams,
        (res) => {
          if (res.err_msg == "get_brand_wcpay_request:ok") {
            // 支付成功
            this.$message.success('支付成功')
            // 轮询订单状态
            this.pollOrderStatus(payParams.orderNo)
          } else {
            // 支付失败
            this.$message.error('支付失败，请重试')
          }
        }
      );
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