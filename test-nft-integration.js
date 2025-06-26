const axios = require('axios');

// 配置
const API_BASE_URL = 'http://localhost:8888/mypapers';
const TEST_USER = {
    username: 'testuser',
    password: 'testpass123'
};

// 测试数据
const TEST_NFT = {
    to: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    tokenId: 1,
    uri: 'ipfs://QmTest123456789'
};

// 工具函数
const log = (message, data = null) => {
    console.log(`[${new Date().toISOString()}] ${message}`);
    if (data) {
        console.log(JSON.stringify(data, null, 2));
    }
    console.log('---');
};

// 获取JWT Token
async function getAuthToken() {
    try {
        log('正在获取认证Token...');
        const response = await axios.post(`${API_BASE_URL}/user/login`, TEST_USER);
        
        if (response.data && response.data.token) {
            log('认证成功', { token: response.data.token.substring(0, 20) + '...' });
            return response.data.token;
        } else {
            throw new Error('登录响应格式错误');
        }
    } catch (error) {
        log('认证失败', error.response?.data || error.message);
        throw error;
    }
}

// 测试NFT铸造
async function testMintNFT(token) {
    try {
        log('测试NFT铸造...');
        const response = await axios.post(`${API_BASE_URL}/nft/mint`, TEST_NFT, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        });
        
        log('NFT铸造成功', response.data);
        return response.data;
    } catch (error) {
        log('NFT铸造失败', error.response?.data || error.message);
        throw error;
    }
}

// 测试获取我的NFT
async function testGetMyNFTs(token) {
    try {
        log('测试获取我的NFT...');
        const response = await axios.get(`${API_BASE_URL}/nft/my-nfts`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        log('获取我的NFT成功', response.data);
        return response.data;
    } catch (error) {
        log('获取我的NFT失败', error.response?.data || error.message);
        throw error;
    }
}

// 测试获取NFT详情
async function testGetNFTDetail(token, tokenId) {
    try {
        log(`测试获取NFT详情 (Token ID: ${tokenId})...`);
        const response = await axios.get(`${API_BASE_URL}/nft/token/${tokenId}`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        log('获取NFT详情成功', response.data);
        return response.data;
    } catch (error) {
        log('获取NFT详情失败', error.response?.data || error.message);
        throw error;
    }
}

// 测试获取NFT市场
async function testGetNFTMarketplace(token) {
    try {
        log('测试获取NFT市场...');
        const response = await axios.get(`${API_BASE_URL}/nft/marketplace?page=1&limit=10`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        log('获取NFT市场成功', response.data);
        return response.data;
    } catch (error) {
        log('获取NFT市场失败', error.response?.data || error.message);
        throw error;
    }
}

// 测试获取NFT统计
async function testGetNFTStats(token) {
    try {
        log('测试获取NFT统计...');
        const response = await axios.get(`${API_BASE_URL}/nft/stats`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        log('获取NFT统计成功', response.data);
        return response.data;
    } catch (error) {
        log('获取NFT统计失败', error.response?.data || error.message);
        throw error;
    }
}

// 测试设置NFT价格
async function testSetNFTPrice(token, tokenId) {
    try {
        log(`测试设置NFT价格 (Token ID: ${tokenId})...`);
        const priceData = {
            tokenId: tokenId,
            price: 0.1
        };
        
        const response = await axios.put(`${API_BASE_URL}/nft/set-price`, priceData, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        });
        
        log('设置NFT价格成功', response.data);
        return response.data;
    } catch (error) {
        log('设置NFT价格失败', error.response?.data || error.message);
        throw error;
    }
}

// 测试购买NFT
async function testBuyNFT(token, tokenId) {
    try {
        log(`测试购买NFT (Token ID: ${tokenId})...`);
        const buyData = {
            tokenId: tokenId
        };
        
        const response = await axios.post(`${API_BASE_URL}/nft/buy`, buyData, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        });
        
        log('购买NFT成功', response.data);
        return response.data;
    } catch (error) {
        log('购买NFT失败', error.response?.data || error.message);
        throw error;
    }
}

// 测试出售NFT
async function testSellNFT(token, tokenId) {
    try {
        log(`测试出售NFT (Token ID: ${tokenId})...`);
        const sellData = {
            tokenId: tokenId,
            price: 0.15
        };
        
        const response = await axios.post(`${API_BASE_URL}/nft/sell`, sellData, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        });
        
        log('出售NFT成功', response.data);
        return response.data;
    } catch (error) {
        log('出售NFT失败', error.response?.data || error.message);
        throw error;
    }
}

// 主测试函数
async function runNFTTests() {
    let token;
    
    try {
        log('开始NFT功能集成测试...');
        
        // 1. 获取认证Token
        token = await getAuthToken();
        
        // 2. 测试NFT铸造
        const mintResult = await testMintNFT(token);
        const tokenId = mintResult.tokenId || TEST_NFT.tokenId;
        
        // 3. 测试获取我的NFT
        await testGetMyNFTs(token);
        
        // 4. 测试获取NFT详情
        await testGetNFTDetail(token, tokenId);
        
        // 5. 测试获取NFT市场
        await testGetNFTMarketplace(token);
        
        // 6. 测试获取NFT统计
        await testGetNFTStats(token);
        
        // 7. 测试设置NFT价格
        await testSetNFTPrice(token, tokenId);
        
        // 8. 测试出售NFT
        await testSellNFT(token, tokenId);
        
        // 9. 测试购买NFT
        await testBuyNFT(token, tokenId);
        
        log('所有NFT功能测试完成！');
        
    } catch (error) {
        log('测试过程中出现错误', error.message);
        process.exit(1);
    }
}

// 健康检查
async function healthCheck() {
    try {
        log('检查服务器健康状态...');
        const response = await axios.get(`${API_BASE_URL}/health`);
        log('服务器健康状态正常', response.data);
        return true;
    } catch (error) {
        log('服务器健康检查失败', error.message);
        return false;
    }
}

// 运行测试
async function main() {
    try {
        // 首先进行健康检查
        const isHealthy = await healthCheck();
        if (!isHealthy) {
            log('服务器不健康，退出测试');
            process.exit(1);
        }
        
        // 运行NFT测试
        await runNFTTests();
        
    } catch (error) {
        log('测试执行失败', error.message);
        process.exit(1);
    }
}

// 如果直接运行此脚本
if (require.main === module) {
    main();
}

module.exports = {
    runNFTTests,
    healthCheck,
    getAuthToken,
    testMintNFT,
    testGetMyNFTs,
    testGetNFTDetail,
    testGetNFTMarketplace,
    testGetNFTStats,
    testSetNFTPrice,
    testBuyNFT,
    testSellNFT
}; 