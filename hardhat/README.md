# MPER/MPS代币项目文档
需要用hardhat和要使用的管理员账户重新部署合约，然后根据生成的部署信息修改后端的变量，如mps合约地址。  
现在合约的注册赠送固定为500，对合约有不满意的可以修改后再部署或者使用合约升级功能，重要的功能都有我测试过的脚本可以使用。  
hardhat的使用、合同的更改、合约的升级、合约的调用都可以结合我生成的脚本和README问大模型。  
后端要调用合约还有一步比较关键的就是要使用geth的abigen，会根据合约的纯abi文件生成一个文件，具体也是需要结合大模型的回答。  

## 项目结构
MPER-project/  
├── contracts/ # 智能合约源文件  
│  ├── MPER.sol # 初始逻辑合约  
│  ├── MPERproxy.sol # 代理合约  
│  ├── MPS.sol # 新逻辑合约  
│  ├── MPSproxy.sol # 新代理合约  
│  ├── MyNFT.sol # NFT初始逻辑合约  
│  ├── MyNFTV2.sol # NFT升级版逻辑合约  
│  └── MyNFTproxy.sol # NFT代理合约  
├── deployments.json # 记录合约部署信息的JSON文件  
├── scripts/ # 部署和管理脚本  
│  ├── MPERdeploy.js # 初始部署脚本  
│  ├── MPSdeploy.js # 新合约部署脚本  
│  ├── upgradeMPER.js # 升级脚本  
│  ├── MyNFTdeploy.js # NFT初始部署脚本  
│  ├── MyNFTupgradeV2.js # NFT升级脚本  
│  ├── executeUpgradeV2.js # NFT升级执行脚本  
│  ├── fixPermissions.js # 权限修复脚本  
│  ├── checkStatus.js # 状态检查脚本  
│  ├── testNFT.js # NFT功能测试脚本  
│  ├── advanceTime.js # 时间推进脚本  
│  ├── verify.js # 部署后验证脚本  
│  └── utils/ # 辅助脚本目录  
│     ├── deployment.js # 部署相关的辅助函数  
│     └── storage.js # 存储布局验证相关的辅助函数  
├── test/ # 测试文件  
├── upgradeinfo/ # 升级相关信息  
├── MPERproxy_technical_docs.md # MPERproxy 合约详细技术说明文档  
├── hardhat.config.js # Hardhat配置  
├── package.json # 项目依赖  
├── package-lock.json # 依赖锁文件  
└── README.md # 项目文档  

## 文件说明

### 合约文件

#### `contracts/MPER.sol`
- 初始逻辑合约
- 实现功能：
  - 可升级的ERC-20代币
  - 批量铸币功能
  - 哈希存储系统
  - 预留存储空间用于未来升级

#### `contracts/MPERproxy.sol`
- 透明可升级代理合约，遵循ERC1967标准
- 特性：
  - 集成了OpenZeppelin的 `TimelockController`
  - 所有权转移给时间锁合约
  - 升级操作通过时间锁提案和执行流程进行管理

#### `contracts/MPS.sol`
- 新版本逻辑合约
- 新增功能：
  - 用户注册跟踪
  - 改进的存储结构
  - 兼容MPER的升级路径

#### `contracts/MPSproxy.sol`
- 新版代理合约
- 改进特性：
  - 简化升级流程
  - 增强安全性检查
  - 更好的错误处理

#### `contracts/MyNFT.sol`
- NFT初始逻辑合约
- 实现功能：
  - 可升级的ERC-721代币
  - 基础铸造功能
  - 元数据管理
  - 版税系统
  - 预留存储空间用于未来升级

#### `contracts/MyNFTV2.sol`
- NFT升级版逻辑合约
- 新增功能：
  - 白名单系统
  - 付费铸造
  - 增强的版税管理
  - 基础URI设置
  - 向后兼容MyNFT

#### `contracts/MyNFTproxy.sol`
- NFT透明可升级代理合约
- 特性：
  - 集成了OpenZeppelin的 `TimelockController`
  - 所有权转移给时间锁合约
  - 升级操作通过时间锁提案和执行流程进行管理

## 开发指南

### 环境搭建
1. 安装依赖：
```bash
npm install
```
2. 配置环境：
```bash
cp .env.example .env
# 编辑.env文件
```
3. 部署合约
```bash
npx hardhat run scripts/MPSdeploy.js --network paperschain
```
4. 验证部署
```bash
npx hardhat run scripts/verify.js --network paperschain
```
5. 运行测试
```bash
npx hardhat test
```

## MyNFT 完整升级流程

### 1. 初始部署

#### 启动本地节点
```bash
npx hardhat node
```

#### 部署MyNFT合约
```bash
npx hardhat run scripts/MyNFTdeploy.js --network localhost
```

**输出示例：**
```
部署者地址: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
正在部署MyNFT逻辑合约...
MyNFT逻辑合约已提交部署到: 0x7a2088a1bFc9d81c55368AE168C2C02570cB814F
正在部署代理合约...
MyNFT代理合约已提交部署到: 0x09635F643e140090A9A8Dcd712eD6285858ceBef
时间锁合约地址: 0xe73bc5BD4763A3307AB5F8F126634b7E12E3dA9b
```

### 2. 升级到MyNFTV2

#### 调度升级
```bash
npx hardhat run scripts/MyNFTupgradeV2.js --network localhost
```

**交互提示：**
- 输入基础元数据URI（例：`ipfs://QmXYZ/`）
- 输入版税百分比（0-100）

**输出示例：**
```
MyNFTV2逻辑合约已部署到: 0xc5a5C42992dECbae36851359345FE25997F5C42d
✅ 升级已调度! 交易哈希: 0x3770e320b869d0b0baa7bdb5b21f91d7f942f028bb3f54db3a64d517c5035119
预计可执行时间: 2025/6/24 10:15:36
```

### 3. 权限修复（如需要）

#### 检查权限状态
```bash
npx hardhat run scripts/fixPermissions.js --network localhost
```

**如果权限不足，进入控制台手动赋权：**
```bash
npx hardhat console --network localhost
```

**控制台命令：**
```javascript
const Timelock = await ethers.getContractFactory("TimelockController");
const timelock = Timelock.attach("0xe73bc5BD4763A3307AB5F8F126634b7E12E3dA9b");
const EXECUTOR_ROLE = await timelock.EXECUTOR_ROLE();
const PROPOSER_ROLE = await timelock.PROPOSER_ROLE();
await timelock.grantRole(EXECUTOR_ROLE, "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266");
await timelock.grantRole(PROPOSER_ROLE, "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266");
.exit
```

### 4. 执行升级

#### 检查升级状态
```bash
npx hardhat run scripts/checkStatus.js --network localhost
```

#### 执行升级
```bash
npx hardhat run scripts/executeUpgradeV2.js --network localhost
```

**成功输出示例：**
```
✅ 升级执行成功! 交易哈希: 0xe28cfc60f95c34d047be9e4a6254efd9ceeb66a86560bfc11ba9cc831a02441e
✅ 升级验证成功: 实现地址已更新
✅ 功能测试完成
🎉 MyNFTV2 升级成功完成!
```

### 5. 功能测试

#### 运行完整测试
```bash
npx hardhat run scripts/testNFT.js --network localhost
```

#### 交互式测试
```bash
npx hardhat console --network localhost
```

**控制台测试命令：**
```javascript
const NFT = await ethers.getContractFactory("MyNFTV2");
const nft = NFT.attach("0x09635F643e140090A9A8Dcd712eD6285858ceBef");
const [deployer] = await ethers.getSigners();

// 基本信息
const owner = await nft.owner();
const symbol = await nft.symbol();
console.log("Owner:", owner);
console.log("Symbol:", symbol);

// V2功能测试
const baseURI = await nft.baseURI();
const mintFee = await nft.mintFee();
console.log("BaseURI:", baseURI);
console.log("MintFee:", ethers.formatEther(mintFee));

// 铸造测试
const mintTx = await nft.safeMint(deployer.address, "Test Token");
await mintTx.wait();
console.log("铸造成功!");
```

### 6. 时间推进（如需要）

如果升级延迟期未到，可以推进时间：
```bash
npx hardhat run scripts/advanceTime.js --network localhost
```

## 升级说明
### 升级流程
1. **准备升级包**：
   - 确保新版本合约(MPS.sol)已通过测试
   - 准备升级脚本(upgradeMPER.js(！还未针对MPS适配，但只需简化))

2. **提交升级提案**：
```bash
npx hardhat run scripts/upgradeMPER.js --network paperschain
```

3. 等待时间锁延迟 ：
   
   - 提案执行前需等待预设的时间锁延迟期
4. 执行升级 ：

   - 时间锁到期后自动执行升级
   - 或手动调用 execute 函数完成升级
5. 验证升级 （未实现）：
```bash
npx hardhat run scripts/verify.js --network paperschain
```
### 注意事项
- 升级前务必备份当前合约状态
- 确保存储布局兼容性（重要）
- 建议在测试网完整测试后再部署到主网
- 详细技术说明请参考：
  - [MPERproxy技术文档](MPERproxy_technical_docs.md)
  - [升级脚本](scripts/upgradeMPER.js)

## MyNFT 升级注意事项

### 权限管理
- **TimelockController** 需要 PROPOSER_ROLE 和 EXECUTOR_ROLE 权限
- 升级前确保操作账户具有相应权限
- 权限问题可通过 `fixPermissions.js` 脚本解决

### 存储兼容性
- MyNFTV2 继承 MyNFT，保持存储布局兼容
- 使用 `__gap` 预留存储空间
- 升级前验证存储布局

### 网络重启
- **Hardhat 本地网络重启后，所有合约状态会丢失**
- 需要重新部署和赋权
- 生产环境使用持久化网络

### 常见问题解决

#### 权限错误
```
AccessControlUnauthorizedAccount
```
**解决方案：**
1. 运行 `fixPermissions.js` 脚本
2. 手动在控制台赋权
3. 确保使用正确的 timelock 地址

#### 升级失败
```
Operation not ready
```
**解决方案：**
1. 检查升级延迟期是否已到
2. 使用 `advanceTime.js` 推进时间
3. 重新执行升级

#### 合约调用失败
```
could not decode result data
```
**解决方案：**
1. 确保 hardhat node 正在运行
2. 检查合约地址是否正确
3. 验证合约是否已部署

## 主要更新内容
### v2.0 (MPS版本)

1. 合约架构升级
   - 从标准 ERC20 合约升级为​​可升级合约​​，使用了 OpenZeppelin 的可升级合约模式  
   - 新增了 Initializable、ERC20Upgradeable 和 OwnableUpgradeable 基础合约  
   - 保留了原有功能的同时，为未来升级预留了存储空间 (__gap)  
2. 权限管理增强
   - 引入了 Ownable 模式，关键函数增加了 onlyOwner 修饰符  
   - registerUser() 现在只能由合约所有者调用  
   - mint() 函数现在只能由合约所有者调用  
   - 移除了公开的 burnFrom() 函数（可能是安全考虑）  
3. 用户注册系统改进
   - 新增了 _hasRegistered 映射，跟踪用户是否已注册  
   - 注册奖励现在从合约自身地址发放（而非固定推广账户）  
   - 添加了防止重复注册的检查  
4. 代码优化与安全增强
   - 所有错误消息都添加了 "MPS:" 前缀以便识别  
   - 转账和铸币函数增加了更严格的参数检查  
   - 移除了硬编码的推广账户地址，改为使用合约自身地址  
5. 初始化流程变更
   - 使用 initialize() 替代构造函数（可升级合约的要求）  
   - 初始化时可选择是否铸造初始供应量  
   - 向后兼容性说明  
   - 升级后的合约保留了原有合约的所有核心功能：
     - 哈希值存储与查询
     - 审稿内容存储与查询
     - 代币转账与铸造
     - 用户注册奖励机制

### v2.0 (MyNFT版本)

1. NFT合约架构升级
   - 从标准 ERC721 合约升级为可升级合约
   - 使用 OpenZeppelin 的可升级合约模式
   - 集成 ERC721Enumerable 和 ERC721URIStorage
   - 预留存储空间用于未来升级

2. 权限管理增强
   - 引入 TimelockController 进行升级控制
   - 所有权转移给时间锁合约
   - 升级操作需要时间锁提案和执行

3. MyNFTV2 新功能
   - 白名单系统：控制铸造权限
   - 付费铸造：设置铸造费用
   - 增强版税管理：支持 ERC2981 标准
   - 基础URI设置：简化元数据管理
   - 向后兼容：完全兼容 MyNFT

4. 升级流程自动化
   - 完整的部署脚本链
   - 自动权限检查和修复
   - 升级状态跟踪和验证
   - 功能测试自动化