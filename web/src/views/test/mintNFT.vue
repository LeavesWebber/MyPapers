<!-- NFTMinter.vue -->
<template>
  <div>
    <input v-model="recipient" placeholder="接收地址" />
    <input v-model.number="tokenId" placeholder="Token ID" type="number" />
    <button @click="mintToken">铸造NFT</button>
    <p v-if="txHash">交易哈希: {{ txHash }}</p>
    <p v-if="error" class="error">{{ error }}</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      recipient: '',
      tokenId: null,
      txHash: '',
      error: ''
    };
  },
  methods: {
    async mintToken() {
      try {
        const response = await axios.post('http://localhost:8080/mint', {
          to: this.recipient,
          tokenId: this.tokenId
        });
        
        this.txHash = response.data.txHash;
        this.error = '';
        
        // 可以添加交易状态轮询
        this.checkTransactionStatus(response.data.txHash);
      } catch (err) {
        this.error = `铸造失败: ${err.response?.data?.error || err.message}`;
        this.txHash = '';
      }
    },
    
    async checkTransactionStatus(txHash) {
      // 这里可以添加轮询逻辑查询交易状态
      console.log(`监控交易状态: ${txHash}`);
    }
  }
};
</script>