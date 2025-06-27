const axios = require('axios');
const { ethers } = require("ethers");
const wallet = new ethers.Wallet("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80");
console.log(wallet.address); // 应该和 userToken 代表的地址一致
// const { Web3 } = require('web3');
// const web3 = new Web3('http://127.0.0.1:8545');
// web3.eth.getChainId().then(console.log);
// 测试配置
const API_BASE_URL = process.env.API_BASE_URL || "http://localhost:8887/mypapers";
const TEST_DATA = {
  // MPS测试数据
  mps: {
    addresses: ["0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266", "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"],
    amount: "1000000000000000000", // 1 MPS (18位小数)
    hash: "QmTestHash123456789",
    reviewContent: "This is a test review content",
    transferAmount: "500000000000000000", // 0.5 MPS
    privateKey: "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
  },
  
  // MyNFT测试数据
  nft: {
    to: "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266",
    uri: "ipfs://QmTestNFT123456789",
    tokenId: 1,
    metadata: "https://example.com/metadata.json",
    royaltyPercentage: 15,
    interfaceId: "0x80ac58cd" // ERC721接口ID
  }
};

// 测试令牌（需要根据实际情况设置）
let adminToken = process.env.ADMIN_TOKEN || "your_admin_token";
let userToken = process.env.USER_TOKEN || "your_user_token";

// 测试结果统计
let testResults = {
  passed: 0,
  failed: 0,
  total: 0
};

// 测试工具函数
function logTest(name, passed, error = null) {
  testResults.total++;
  if (passed) {
    testResults.passed++;
    console.log(`✅ ${name} - 通过`);
  } else {
    testResults.failed++;
    console.log(`❌ ${name} - 失败`);
    if (error) {
      console.log(`   错误: ${error.message || error}`);
    }
  }
}

function logSection(title) {
  console.log(`\n${'='.repeat(50)}`);
  console.log(`📋 ${title}`);
  console.log(`${'='.repeat(50)}`);
}

function logSummary() {
  console.log(`\n${'='.repeat(50)}`);
  console.log(`📊 测试总结`);
  console.log(`${'='.repeat(50)}`);
  console.log(`总测试数: ${testResults.total}`);
  console.log(`通过: ${testResults.passed} ✅`);
  console.log(`失败: ${testResults.failed} ❌`);
  console.log(`成功率: ${((testResults.passed / testResults.total) * 100).toFixed(2)}%`);
}

// 异步测试函数
async function runTest(name, testFunction) {
  try {
    await testFunction();
    logTest(name, true);
  } catch (error) {
    logTest(name, false, error);
  }
}

// MPS API 测试
async function testMPSAPIs() {
  logSection("MPS API 测试");
    
  // 1. 批量铸币测试
  await runTest("MPS批量铸币 - 成功", async () => {
    const mintData = {
      toAddresses: ["0x270DE39CBB9d711f565AD74D56238689901aDC71"],
      amount: "1000000000000000000"
    };
    
    const response = await axios.post(`${API_BASE_URL}/mps/mint`, mintData, {
      headers: {
        'Authorization':`Bearer ${adminToken}`,
        'Content-Type': 'application/json'
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
    
    if (!response.data.data.transactionHash) {
      throw new Error("缺少交易哈希");
    }
  });

  // 2. 查询余额测试
  await runTest("MPS查询余额 - 成功", async () => {
    const address = TEST_DATA.mps.addresses[0];
    
    const response = await axios.get(`${API_BASE_URL}/mps/balance-of/${address}`, {
      headers: {
        'Authorization': `Bearer ${userToken}`
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
    
    if (!response.data.data.balance) {
      throw new Error("缺少余额信息");
    }
  });

  // 3. 转账测试
  await runTest("MPS转账 - 成功", async () => {
    const transferData = {
      to: TEST_DATA.mps.addresses[1],
      amount: TEST_DATA.mps.transferAmount,
      privateKey: TEST_DATA.mps.privateKey
    };
    
    const response = await axios.post(`${API_BASE_URL}/mps/transfer`, transferData, {
      headers: {
        'Authorization': `Bearer ${userToken}`,
        'Content-Type': 'application/json'
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 4. 存储哈希测试
  await runTest("MPS存储哈希 - 成功", async () => {
    const hashData = {
      hash: TEST_DATA.mps.hash
    };
    
    const response = await axios.post(`${API_BASE_URL}/mps/store-hash`, hashData, {
      headers: {
        'Authorization': `Bearer ${userToken}`,
        'Content-Type': 'application/json'
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 5. 查询哈希归属测试
  await runTest("MPS查询哈希归属 - 成功", async () => {
    const hash = TEST_DATA.mps.hash;
    
    const response = await axios.get(`${API_BASE_URL}/mps/recipient-by-hash/${hash}`, {
      headers: {
        'Authorization': `Bearer ${userToken}`
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 6. 存储审稿内容测试
  await runTest("MPS存储审稿内容 - 成功", async () => {
    const reviewData = {
      content: TEST_DATA.mps.reviewContent
    };
    
    const response = await axios.post(`${API_BASE_URL}/mps/store-review`, reviewData, {
      headers: {
        'Authorization': `Bearer ${userToken}`,
        'Content-Type': 'application/json'
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 7. 查询审稿归属测试
  await runTest("MPS查询审稿归属 - 成功", async () => {
    const content = TEST_DATA.mps.reviewContent;
    
    const response = await axios.get(`${API_BASE_URL}/mps/review-by-hash/${encodeURIComponent(content)}`, {
      headers: {
        'Authorization': `Bearer ${userToken}`
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 8. 注册用户测试
  await runTest("MPS注册用户 - 成功", async () => {
    const registerData = {
      userAddress: TEST_DATA.mps.addresses[0]
    };
    
    const response = await axios.post(`${API_BASE_URL}/mps/register-user`, registerData, {
      headers: {
        'Authorization': `Bearer ${adminToken}`,
        'Content-Type': 'application/json'
      }
    });
    
    if (response.status !== 200) {
      throw new Error(`期望状态码200，实际${response.status}`);
    }
    if (typeof response.data.msg === 'string' && response.data.msg.includes('合约MPS余额不足')) {
      console.log('⚠️ 注册用户时合约MPS余额不足，无法完成注册奖励');
      return;
    }
    if (response.data.code !== 1000) {
      throw new Error(`期望代码1000，实际${response.data.code}，msg: ${response.data.msg}`);
    }
  });

  // 错误情况测试
  await runTest("MPS批量铸币 - 权限不足", async () => {
    const mintData = {
      userAddress: TEST_DATA.mps.addresses[0],
      amount: TEST_DATA.mps.amount
    };
    
    try {
      await axios.post(`${API_BASE_URL}/mps/mint`, mintData, {
        headers: {
          'Authorization': `Bearer ${userToken}`,
          'Content-Type': 'application/json'
        }
      });
      throw new Error("应该抛出权限不足错误");
    } catch (error) {
      if (error.response && error.response.status === 403) {
        // 预期的权限不足错误
        return;
      }
      throw error;
    }
  });

  await runTest("MPS查询余额 - 无效地址", async () => {
    const invalidAddress = "invalid_address";
    
    try {
      await axios.get(`${API_BASE_URL}/mps/balance-of/${invalidAddress}`, {
        headers: {
          'Authorization': `Bearer ${userToken}`
        }
      });
      throw new Error("应该抛出无效地址错误");
    } catch (error) {
      if (error.response && error.response.status === 400) {
        // 预期的无效地址错误
        return;
      }
      throw error;
    }
  });
}

// MyNFT API 测试
async function testMyNFTAPIs() {
  logSection("MyNFT API 测试");

  // 1. 铸造NFT测试
  await runTest("MyNFT铸造 - 成功", async () => {
    const mintData = {
      to: TEST_DATA.nft.to,
      uri: TEST_DATA.nft.uri
    };
    
    const response = await axios.post(`${API_BASE_URL}/mynft/mint`, mintData, {
      headers: {
        'Authorization': `Bearer ${adminToken}`,
        'Content-Type': 'application/json'
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
    
    if (!response.data.data.tokenId) {
      throw new Error("缺少tokenId");
    }
  });

  // 2. 批量铸造测试
  await runTest("MyNFT批量铸造 - 成功", async () => {
    const bulkMintData = {
      recipients: [TEST_DATA.mps.addresses[0], TEST_DATA.mps.addresses[1]],
      uris: [TEST_DATA.nft.uri, "ipfs://QmTestNFT2"]
    };
    
    const response = await axios.post(`${API_BASE_URL}/mynft/bulk-mint`, bulkMintData, {
      headers: {
        'Authorization': `Bearer ${adminToken}`,
        'Content-Type': 'application/json'
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 3. 设置合约元数据测试
  await runTest("MyNFT设置元数据 - 成功", async () => {
    const metadataUri = {
      metadataUri: TEST_DATA.nft.metadata
    };
    
    const response = await axios.post(`${API_BASE_URL}/mynft/set-metadata`, metadataUri, {
      headers: {
        'Authorization': `Bearer ${adminToken}`,
        'Content-Type': 'application/json'
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 4. 更新版税测试
  await runTest("MyNFT更新版税 - 成功", async () => {
    const royaltyData = {
      royaltyPercentage: TEST_DATA.nft.royaltyPercentage
    };
    
    const response = await axios.post(`${API_BASE_URL}/mynft/update-royalty`, royaltyData, {
      headers: {
        'Authorization': `Bearer ${adminToken}`,
        'Content-Type': 'application/json'
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 5. 查询tokenURI测试
  await runTest("MyNFT查询tokenURI - 成功", async () => {
    const tokenId = TEST_DATA.nft.tokenId;
    
    const response = await axios.get(`${API_BASE_URL}/mynft/token-uri?tokenId=${tokenId}`, {
      headers: {
        'Authorization': `Bearer ${userToken}`
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 6. 查询合约支持的接口测试
  await runTest("MyNFT查询接口支持 - 成功", async () => {
    const interfaceId = TEST_DATA.nft.interfaceId;
    
    const response = await axios.get(`${API_BASE_URL}/mynft/supports-interface?interfaceId=${interfaceId}`, {
      headers: {
        'Authorization': `Bearer ${userToken}`
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 7. 查询owner测试
  await runTest("MyNFT查询owner - 成功", async () => {
    const response = await axios.get(`${API_BASE_URL}/mynft/owner`, {
      headers: {
        'Authorization': `Bearer ${userToken}`
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 8. 查询余额测试
  await runTest("MyNFT查询余额 - 成功", async () => {
    const address = TEST_DATA.mps.addresses[0];
    
    const response = await axios.get(`${API_BASE_URL}/mynft/balance-of?address=${address}`, {
      headers: {
        'Authorization': `Bearer ${userToken}`
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 9. 查询ownerOf测试
  await runTest("MyNFT查询ownerOf - 成功", async () => {
    const tokenId = TEST_DATA.nft.tokenId;
    
    const response = await axios.get(`${API_BASE_URL}/mynft/owner-of?tokenId=${tokenId}`, {
      headers: {
        'Authorization': `Bearer ${userToken}`
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 10. 查询总供应量测试
  await runTest("MyNFT查询总供应量 - 成功", async () => {
    const response = await axios.get(`${API_BASE_URL}/mynft/total-supply`, {
      headers: {
        'Authorization': `Bearer ${userToken}`
      }
    });
    
    if (response.status !== 200 || response.data.code !== 1000) {
      throw new Error(`期望状态码200，实际${response.status}；期望代码1000，实际${response.data.code}`);
    }
  });

  // 错误情况测试
  await runTest("MyNFT铸造 - 权限不足", async () => {
    const mintData = {
      to: TEST_DATA.nft.to,
      uri: TEST_DATA.nft.uri
    };
    
    try {
      await axios.post(`${API_BASE_URL}/mynft/mint`, mintData, {
        headers: {
          'Authorization': `Bearer ${userToken}`,
          'Content-Type': 'application/json'
        }
      });
      throw new Error("应该抛出权限不足错误");
    } catch (error) {
      if (error.response && error.response.status === 403) {
        // 预期的权限不足错误
        return;
      }
      throw error;
    }
  });

  await runTest("MyNFT批量铸造 - 数组长度不匹配", async () => {
    const bulkMintData = {
      recipients: [TEST_DATA.mps.addresses[0]],
      uris: [TEST_DATA.nft.uri, "ipfs://QmTestNFT2"]
    };
    
    try {
      await axios.post(`${API_BASE_URL}/mynft/bulk-mint`, bulkMintData, {
        headers: {
          'Authorization': `Bearer ${adminToken}`,
          'Content-Type': 'application/json'
        }
      });
      throw new Error("应该抛出数组长度不匹配错误");
    } catch (error) {
      if (error.response && error.response.status === 400) {
        // 预期的数组长度不匹配错误
        return;
      }
      throw error;
    }
  });
}

// 集成测试
async function testIntegration() {
  logSection("集成测试");

  await runTest("完整MPS和NFT工作流程", async () => {
    // 1. 注册用户
    const registerResponse = await axios.post(`${API_BASE_URL}/mps/register-user`, {
      address: TEST_DATA.mps.addresses[0]
    }, {
      headers: { 'Authorization': `Bearer ${adminToken}` }
    });
    
    if (registerResponse.status !== 200) {
      throw new Error("用户注册失败");
    }
    
    // 2. 铸造MPS代币
    const mintResponse = await axios.post(`${API_BASE_URL}/mps/mint`, {
      addresses: [TEST_DATA.mps.addresses[0]],
      amount: TEST_DATA.mps.amount
    }, {
      headers: { 'Authorization': `Bearer ${adminToken}` }
    });
    
    if (mintResponse.status !== 200) {
      throw new Error("MPS铸造失败");
    }
    
    // 3. 查询MPS余额
    const balanceResponse = await axios.get(`${API_BASE_URL}/mps/balance-of/${TEST_DATA.mps.addresses[0]}`, {
      headers: { 'Authorization': `Bearer ${userToken}` }
    });
    
    if (balanceResponse.status !== 200) {
      throw new Error("MPS余额查询失败");
    }
    
    // 4. 铸造NFT
    const nftResponse = await axios.post(`${API_BASE_URL}/mynft/mint`, {
      to: TEST_DATA.mps.addresses[0],
      uri: TEST_DATA.nft.uri
    }, {
      headers: { 'Authorization': `Bearer ${adminToken}` }
    });
    
    if (nftResponse.status !== 200) {
      throw new Error("NFT铸造失败");
    }
    
    // 5. 查询NFT余额
    const nftBalanceResponse = await axios.get(`${API_BASE_URL}/mynft/balance-of/${TEST_DATA.mps.addresses[0]}`, {
      headers: { 'Authorization': `Bearer ${userToken}` }
    });
    
    if (nftBalanceResponse.status !== 200) {
      throw new Error("NFT余额查询失败");
    }
  });
}

// 性能测试
async function testPerformance() {
  logSection("性能测试");

  await runTest("并发余额查询", async () => {
    const address = TEST_DATA.mps.addresses[0];
    const concurrentRequests = 5;
    
    const promises = Array(concurrentRequests).fill().map(() => 
      axios.get(`${API_BASE_URL}/mps/balance-of/${address}`, {
        headers: { 'Authorization': `Bearer ${userToken}` }
      })
    );
    
    const responses = await Promise.all(promises);
    
    responses.forEach((response, index) => {
      if (response.status !== 200) {
        throw new Error(`并发请求${index + 1}失败`);
      }
    });
  });
}

// 主测试函数
async function runAllTests() {
  console.log("🚀 开始运行区块链API测试...");
  console.log(`📡 API地址: ${API_BASE_URL}`);
  console.log(`⏰ 开始时间: ${new Date().toLocaleString()}`);
  
  try {
    await testMPSAPIs();
    await testMyNFTAPIs();
    await testIntegration();
    await testPerformance();
  } catch (error) {
    console.error("❌ 测试过程中发生错误:", error.message);
  }
  
  logSummary();
  
  console.log(`⏰ 结束时间: ${new Date().toLocaleString()}`);
  
  // 如果有失败的测试，退出码为1
  if (testResults.failed > 0) {
    process.exit(1);
  }
}

// 如果直接运行此文件
if (require.main === module) {
  runAllTests().catch(console.error);
}

module.exports = {
  runAllTests,
  testMPSAPIs,
  testMyNFTAPIs,
  testIntegration,
  testPerformance
}; 