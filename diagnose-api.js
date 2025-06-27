const axios = require('axios');

// 测试配置
const API_BASE_URL = process.env.API_BASE_URL || "http://localhost:8887/mypapers";

console.log('🔍 开始诊断区块链API连接问题...');
console.log(`📡 API地址: ${API_BASE_URL}`);
console.log('');

// 1. 测试基本连接
async function testBasicConnection() {
    console.log('1️⃣ 测试基本连接...');
    try {
        const response = await axios.get(`${API_BASE_URL}/health`, { timeout: 5000 });
        console.log('✅ 服务器连接成功');
        console.log(`   状态码: ${response.status}`);
        console.log(`   响应: ${JSON.stringify(response.data)}`);
    } catch (error) {
        console.log('❌ 服务器连接失败');
        if (error.code === 'ECONNREFUSED') {
            console.log('   原因: 服务器未启动或端口不正确');
        } else if (error.code === 'ENOTFOUND') {
            console.log('   原因: 无法解析主机名');
        } else if (error.code === 'ETIMEDOUT') {
            console.log('   原因: 连接超时');
        } else {
            console.log(`   错误: ${error.message}`);
        }
    }
    console.log('');
}

// 2. 测试MPS API端点
async function testMPSEndpoints() {
    console.log('2️⃣ 测试MPS API端点...');
    
    const endpoints = [
        { method: 'GET', path: '/mps/balance-of/0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266', name: '查询余额' },
        { method: 'GET', path: '/mps/recipient-by-hash/test', name: '查询哈希归属' },
        { method: 'GET', path: '/mps/review-by-hash/test', name: '查询审稿归属' }
    ];

    for (const endpoint of endpoints) {
        try {
            const response = await axios({
                method: endpoint.method,
                url: `${API_BASE_URL}${endpoint.path}`,
                timeout: 5000
            });
            console.log(`✅ ${endpoint.name} - 成功`);
            console.log(`   状态码: ${response.status}`);
        } catch (error) {
            console.log(`❌ ${endpoint.name} - 失败`);
            if (error.response) {
                console.log(`   状态码: ${error.response.status}`);
                console.log(`   响应: ${JSON.stringify(error.response.data)}`);
            } else {
                console.log(`   错误: ${error.message}`);
            }
        }
    }
    console.log('');
}

// 3. 测试MyNFT API端点
async function testMyNFTEndpoints() {
    console.log('3️⃣ 测试MyNFT API端点...');
    
    const endpoints = [
        { method: 'GET', path: '/mynft/owner', name: '查询owner' },
        { method: 'GET', path: '/mynft/total-supply', name: '查询总供应量' },
        { method: 'GET', path: '/mynft/balance-of/0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266', name: '查询余额' },
        { method: 'GET', path: '/mynft/token-uri/1', name: '查询tokenURI' }
    ];

    for (const endpoint of endpoints) {
        try {
            const response = await axios({
                method: endpoint.method,
                url: `${API_BASE_URL}${endpoint.path}`,
                timeout: 5000
            });
            console.log(`✅ ${endpoint.name} - 成功`);
            console.log(`   状态码: ${response.status}`);
        } catch (error) {
            console.log(`❌ ${endpoint.name} - 失败`);
            if (error.response) {
                console.log(`   状态码: ${error.response.status}`);
                console.log(`   响应: ${JSON.stringify(error.response.data)}`);
            } else {
                console.log(`   错误: ${error.message}`);
            }
        }
    }
    console.log('');
}

// 4. 测试POST端点（需要认证）
async function testPostEndpoints() {
    console.log('4️⃣ 测试POST端点（需要认证）...');
    
    const endpoints = [
        { 
            path: '/mps/mint', 
            name: 'MPS铸币',
            data: {
                toAddresses: ["0x270DE39CBB9d711f565AD74D56238689901aDC71"],
                amount: 1000000000000000000
            }
        },
        { 
            path: '/mynft/mint', 
            name: 'NFT铸造',
            data: {
                to: "0x270DE39CBB9d711f565AD74D56238689901aDC71",
                uri: "http://127.0.0.1:26000/ipfs/QmbF3PuGRZ6xgvzKuTLAtBydagXFLjowMb3JC9tWUzf4wM"
            }
        }
    ];

    for (const endpoint of endpoints) {
        try {
            const response = await axios({
                method: 'POST',
                url: `${API_BASE_URL}${endpoint.path}`,
                data: endpoint.data,
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjo1NzQ4NTY3NjcxMTcyMDE0MDgsIklEIjoyLCJVc2VybmFtZSI6IjEyMyIsIkZpcnN0TmFtZSI6IjEyMyIsIkxhc3ROYW1lIjoiMTIzIiwiQXV0aG9yaXR5SWQiOjEwMiwiZXhwIjoxNzUxNjEwMTQ3LCJpc3MiOiJibHVlYmVsbCJ9.WOkP4Z7eqHd28vXWlrkdPQAgFPO-CJ82gBy7dXk1gSI'
                },
                timeout: 5000
            });
            console.log(`✅ ${endpoint.name} - 成功`);
            console.log(`   状态码: ${response.status}`);
        } catch (error) {
            console.log(`❌ ${endpoint.name} - 失败`);
            if (error.response) {
                console.log(`   状态码: ${error.response.status}`);
                console.log(`   响应: ${JSON.stringify(error.response.data)}`);
            } else {
                console.log(`   错误: ${error.message}`);
            }
        }
    }
    console.log('');
}

// 5. 检查环境变量
function checkEnvironment() {
    console.log('5️⃣ 检查环境变量...');
    console.log(`   API_BASE_URL: ${process.env.API_BASE_URL || '未设置 (使用默认值)'}`);
    console.log(`   ADMIN_TOKEN: ${process.env.ADMIN_TOKEN ? '已设置' : '未设置'}`);
    console.log(`   USER_TOKEN: ${process.env.USER_TOKEN ? '已设置' : '未设置'}`);
    console.log('');
}

// 6. 提供建议
function provideRecommendations() {
    console.log('6️⃣ 建议和解决方案...');
    console.log('');
    console.log('如果所有测试都失败，请检查以下项目：');
    console.log('');
    console.log('1. 服务器状态:');
    console.log('   - 确保Go服务器正在运行');
    console.log('   - 检查端口8888是否被占用');
    console.log('   - 查看服务器日志是否有错误');
    console.log('');
    console.log('2. 配置文件:');
    console.log('   - 检查 config.yaml 中的区块链配置');
    console.log('   - 验证合约地址是否正确');
    console.log('   - 确认RPC URL是否可访问');
    console.log('');
    console.log('3. 网络连接:');
    console.log('   - 检查防火墙设置');
    console.log('   - 确认网络连接正常');
    console.log('');
    console.log('4. 依赖项:');
    console.log('   - 确保所有Go依赖项已安装');
    console.log('   - 检查区块链节点是否运行');
    console.log('');
}

// 主函数
async function main() {
    try {
        await testBasicConnection();
        await testMPSEndpoints();
        await testMyNFTEndpoints();
        await testPostEndpoints();
        checkEnvironment();
        provideRecommendations();
    } catch (error) {
        console.error('诊断过程中发生错误:', error.message);
    }
}

// 运行诊断
main(); 