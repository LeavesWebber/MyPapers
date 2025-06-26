# NFT功能集成指南

## 概述

本项目已成功集成了NFT（非同质化代币）功能，包括铸造、交易、市场等完整功能。前后端已完全连接，支持用户通过Web界面进行NFT操作。

## 功能特性

### 1. NFT铸造
- 支持自定义Token ID和元数据URI
- 实时预览NFT信息
- 交易状态跟踪
- 区块浏览器链接

### 2. NFT市场
- 浏览所有可交易的NFT
- 搜索和筛选功能
- 价格排序
- 分类浏览
- 购买确认流程

### 3. NFT管理
- 查看个人NFT收藏
- 设置NFT价格
- 转移NFT所有权
- 查看交易历史

## 技术架构

### 前端技术栈
- Vue.js 2.x
- Element UI
- Axios HTTP客户端
- Vue Router

### 后端技术栈
- Go (Gin框架)
- Ethereum Go客户端
- 智能合约交互
- JWT认证

### 区块链技术
- Hardhat开发环境
- 可升级NFT合约
- 代理合约模式
- 时间锁控制器

## 配置说明

### 1. 环境变量配置

在服务器端创建 `.env` 文件，包含以下配置：

```bash
# 区块链配置
RPC_URL=http://localhost:8545
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
PROXY_CONTRACT=0x5FbDB2315678afecb367f032d93F642f64180aa3
NETWORK=localhost

# 其他配置...
```

### 2. 合约部署

确保已部署NFT合约：

```bash
# 部署NFT合约
npx hardhat run scripts/MyNFTdeploy.js --network localhost

# 升级到V2版本（可选）
npx hardhat run scripts/MyNFTupgradeV2.js --network localhost
npx hardhat run scripts/executeUpgradeV2.js --network localhost
```

### 3. 权限配置

确保时间锁控制器有正确的权限：

```javascript
// 在Hardhat控制台中执行
const timelock = await ethers.getContractAt("TimelockController", "TIMELOCK_ADDRESS");
const proposerRole = await timelock.PROPOSER_ROLE();
const executorRole = await timelock.EXECUTOR_ROLE();

// 授予权限
await timelock.grantRole(proposerRole, "YOUR_ADDRESS");
await timelock.grantRole(executorRole, "YOUR_ADDRESS");
```

## API接口

### 1. NFT铸造
```http
POST /mypapers/nft/mint
Content-Type: application/json

{
  "to": "0x...",
  "tokenId": 1,
  "uri": "ipfs://QmXYZ..."
}
```

### 2. 获取我的NFT
```http
GET /mypapers/nft/my-nfts
Authorization: Bearer <token>
```

### 3. 获取NFT详情
```http
GET /mypapers/nft/token/{tokenId}
```

### 4. 获取NFT市场
```http
GET /mypapers/nft/marketplace?page=1&limit=20&category=paper&sortBy=price
```

### 5. 购买NFT
```http
POST /mypapers/nft/buy
Content-Type: application/json

{
  "tokenId": 1
}
```

### 6. 出售NFT
```http
POST /mypapers/nft/sell
Content-Type: application/json

{
  "tokenId": 1,
  "price": 0.1
}
```

## 前端页面

### 1. NFT铸造页面
- 路径: `/center/nft-mint`
- 功能: 铸造新NFT，设置接收地址、Token ID和元数据URI

### 2. NFT市场页面
- 路径: `/center/nft-marketplace`
- 功能: 浏览和购买NFT，支持搜索、筛选和排序

### 3. 我的NFT页面
- 路径: `/center/myNFTs`
- 功能: 管理个人NFT收藏

## 使用流程

### 1. 铸造NFT
1. 登录系统
2. 进入NFT铸造页面
3. 填写接收地址、Token ID和元数据URI
4. 点击"铸造NFT"
5. 等待交易确认
6. 查看铸造结果

### 2. 购买NFT
1. 进入NFT市场页面
2. 浏览可用的NFT
3. 点击"购买"按钮
4. 确认购买信息
5. 完成交易

### 3. 出售NFT
1. 进入我的NFT页面
2. 选择要出售的NFT
3. 设置价格
4. 确认出售

## 开发说明

### 1. 前端开发
- NFT相关API在 `src/api/nft.js` 中定义
- 页面组件在 `src/views/center/` 目录下
- 路由配置在 `src/router/index.js` 中

### 2. 后端开发
- NFT API处理在 `server/api/nft.go` 中
- 路由配置在 `server/router/nft.go` 中
- 区块链配置在 `server/config/blockchain.go` 中

### 3. 智能合约
- NFT合约: `hardhat/contracts/MyNFT.sol`
- 代理合约: `hardhat/contracts/MyNFTproxy.sol`
- 升级合约: `hardhat/contracts/MyNFTV2.sol`

## 注意事项

1. **私钥安全**: 确保私钥安全存储，不要提交到版本控制系统
2. **网络配置**: 根据部署环境配置正确的RPC URL和网络类型
3. **Gas费用**: 确保账户有足够的ETH支付Gas费用
4. **合约地址**: 确保使用正确的代理合约地址
5. **权限管理**: 确保时间锁控制器有正确的权限配置

## 故障排除

### 1. 铸造失败
- 检查私钥是否正确
- 确认账户有足够ETH
- 验证合约地址是否正确
- 检查RPC连接是否正常

### 2. 升级失败
- 确认有PROPOSER_ROLE权限
- 检查时间锁是否已过期
- 验证升级合约是否兼容

### 3. API调用失败
- 检查JWT token是否有效
- 确认API路径是否正确
- 验证请求参数格式

## 扩展功能

### 1. 批量铸造
- 支持一次铸造多个NFT
- 批量设置元数据

### 2. NFT分类
- 支持自定义NFT分类
- 分类筛选和搜索

### 3. 版税系统
- 支持创作者版税
- 自动版税分配

### 4. 拍卖功能
- 支持NFT拍卖
- 竞价机制

### 5. 社交功能
- NFT评论系统
- 收藏夹功能
- 关注创作者

## 联系支持

如有问题或需要技术支持，请联系开发团队。 