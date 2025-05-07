<template>
  <div class="recharge-container">
    <el-card class="recharge-card">
      <div slot="header" class="card-header">
        <span class="header-title">充值 MPS</span>
        <span class="header-subtitle">快速充值，便捷支付</span>
      </div>
      <div class="balance-info">
        <div class="balance-card">
          <i class="el-icon-wallet"></i>
          <div class="balance-details">
            <span class="balance-label">当前 MPS 余额</span>
            <span class="balance-amount">{{ mpsBalance }} MPS</span>
          </div>
        </div>
        <div class="exchange-rate">
          <i class="el-icon-refresh"></i>
          <span>兑换比例：1 CNY = 1 MPS</span>
        </div>
      </div>
      <el-form :model="rechargeForm" :rules="rules" ref="rechargeForm" label-width="100px" class="recharge-form">
        <el-form-item label="充值金额" prop="amount">
          <el-input-number 
            v-model="rechargeForm.amount" 
            :min="1"
            :max="10000"
            controls-position="right"
            class="amount-input"
            @change="handleAmountChange">
          </el-input-number>
          <span class="unit">CNY</span>
        </el-form-item>
        <el-form-item label="获得MPS">
          <div class="mps-amount-card">
            <span class="mps-amount">{{ rechargeForm.amount }}</span>
            <span class="mps-unit">MPS</span>
          </div>
        </el-form-item>
        <el-form-item label="支付方式">
          <div class="payment-methods">
            <div 
              class="payment-method alipay" 
              :class="{ active: rechargeForm.pay_type === 'alipay' }"
              @click="rechargeForm.pay_type = 'alipay'"
            >
              <img src="@/assets/alipay-logo.svg" alt="支付宝" class="payment-logo">
            </div>
            <div 
              class="payment-method wxpay"
              :class="{ active: rechargeForm.pay_type === 'wxpay' }"
              @click="rechargeForm.pay_type = 'wxpay'"
            >
              <img src="@/assets/wechat-logo.svg" alt="微信" class="payment-logo">
              <span>微信支付</span>
            </div>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleRecharge" :loading="loading" class="recharge-button">
            立即充值
          </el-button>
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

        const balance = await MPScontractInstance.methods['balanceOf(address)'](accounts[0]).call()
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

            console.log('开始充值请求:', {
              amount: this.rechargeForm.amount,
              wallet_address: accounts[0],
              pay_type: this.rechargeForm.pay_type
            })

            const res = await buyMPSWithFiat({
              amount: this.rechargeForm.amount,
              wallet_address: accounts[0],
              pay_type: this.rechargeForm.pay_type
            })
            
            console.log('充值响应:', res)

            // 检查响应数据结构
            if (!res || !res.data) {
              console.error('响应数据格式错误:', res)
              this.$message.error('服务器响应格式错误')
              return
            }

            const { code, data, msg } = res.data

            if (code === 0) {
              if (this.rechargeForm.pay_type === 'alipay') {
                // 支付宝支付
                if (data && data.pay_url) {
                  // 保存订单号和钱包地址到本地存储
                  localStorage.setItem('current_order_no', data.order_no)
                  localStorage.setItem('current_wallet_address', accounts[0])
                  localStorage.setItem('current_mps_amount', this.rechargeForm.amount)
                  
                  // 直接使用后端返回的支付URL，不要修改
                  console.log('跳转到支付页面:', data.pay_url)
                  window.location.href = data.pay_url
                } else {
                  console.error('支付宝支付链接缺失:', data)
                  this.$message.error('获取支付链接失败')
                }
              } else {
                // 微信支付
                if (data && data.wx_pay_params) {
                  this.callWxPay(data.wx_pay_params)
                } else {
                  console.error('微信支付参数缺失:', data)
                  this.$message.error('获取微信支付参数失败')
                }
              }
            } else {
              console.error('充值失败:', msg || '未知错误')
              this.$message.error(msg || '充值失败，请重试')
            }
          }
        })
      } catch (error) {
        console.error('充值过程发生错误:', error)
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
  padding: 30px;
  max-width: 800px;
  margin: 0 auto;
}

.recharge-card {
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border: none;
}

.card-header {
  display: flex;
  /* flex-direction: column; */
  align-items: center;
  padding: 5px;
}

.header-title {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 5px;
}

.header-subtitle {
  font-size: 14px;
  color: #909399;
  margin-left: 10px;
}

.balance-info {
  margin: 5px 0;
}

.balance-card {
  display: flex;
  align-items: left;
  background: #f5f7fa;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 20px;
}

.balance-card i {
  font-size: 24px;
  color: #409EFF;
  margin-right: 15px;
}

.balance-details {
  display: flex;
  flex-direction: column;
}

.balance-label {
  font-size: 14px;
  color: #606266;
  margin-bottom: 5px;
}

.balance-amount {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
}

.exchange-rate {
  margin: 20px auto;
  display: flex;
  align-items: center;
  color: #909399;
  font-size: 14px;
}

.exchange-rate i {
  margin-right: 8px;
}

.recharge-form {
  max-width: 500px;
  /* margin: 0 auto; */
}

.amount-input {
  width: 200px;
}

.unit {
  margin-left: 10px;
  color: #909399;
}

.mps-amount-card {
  display: flex;
  align-items: baseline;
  background: #f5f7fa;
  padding: 10px 20px;
  border-radius: 8px;
}

.mps-amount {
  font-size: 24px;
  font-weight: bold;
  color: #409EFF;
}

.mps-unit {
  margin-left: 5px;
  color: #606266;
}

.payment-methods {
  display: flex;
  gap: 20px;
  width: 100%;
}

.payment-method {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 5px 20px;
  border: 1px solid #DCDFE6;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
  flex: 1;
  max-width: 200px;
}

.payment-logo {
  width: 80px;
  height: 28px;
  object-fit: contain;
}

.payment-method.wxpay .payment-logo {
  width: 40px;
  height: 40px;
  margin-right: 10px;
  filter: brightness(0) saturate(100%);
  transition: filter 0.3s;
  filter: brightness(0) saturate(100%) invert(56%) sepia(83%) saturate(543%) hue-rotate(93deg) brightness(96%) contrast(93%);
}


.payment-method.alipay:hover,
.payment-method.alipay.active {
  border-color: #1677FF;
  background: rgba(22, 119, 255, 0.05);
}

.payment-method.wxpay:hover,
.payment-method.wxpay.active {
  border-color: #07C160;
  background: rgba(7, 193, 96, 0.05);
}

.payment-method span {
  font-size: 18px;
  font-weight: 500;
  color: #606266;
}

.payment-method.active span {
  color: #07C160;
}

.recharge-button {
  width: 100%;
  height: 48px;
  font-size: 16px;
  border-radius: 8px;
  margin-top: 5px;
}
</style>