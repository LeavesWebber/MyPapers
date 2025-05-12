# MPER代币项目文档

## 项目结构
MPER-project/
├── contracts/ # 智能合约源文件  
│  ├── MPER.sol # 主逻辑合约  
│  └── MPERproxy.sol # 代理合约  
├── deployments.json # 记录合约部署信息的JSON文件  
├── scripts/ # 部署和管理脚本  
│  ├── MPERdeploy.js # 主部署脚本  
│  ├── verify.js     # 部署后验证脚本  
│  └── utils/        # 辅助脚本目录  
│     ├── deployment.js # 部署相关的辅助函数 (例如保存部署信息、获取实现地址)  
│     └── storage.js    # 存储布局验证相关的辅助函数  
├── test/ # 测试文件  
├── MPERproxy_technical_docs.md # MPERproxy 合约详细技术说明文档  
├── hardhat.config.js # Hardhat配置  
├── package.json # 项目依赖  
├── package-lock.json # 依赖锁文件  
└── README.md # 项目文档  


## 文件说明

### 合约文件

#### `contracts/MPER.sol`
- 主逻辑合约 (目前暂未升级，具体功能待讨论)
- 实现功能：
  - 可升级的ERC-20代币
  - 批量铸币功能
  - 哈希存储系统
  - 预留存储空间用于未来升级

#### `contracts/MPERproxy.sol`
- 透明可升级代理合约，遵循ERC1967标准。
- 特性：
  - **治理机制**: 集成了OpenZeppelin的 `TimelockController`，所有权转移给时间锁合约，敏感操作（如升级）通过时间锁提案和执行流程进行管理。
  - **升级方式**: 逻辑合约的升级通过 `TimelockController` 调度 `MPERproxy` 的 `proposeUpgrade` 和 `upgradeToAndCall` 函数完成，包含时间延迟以确保安全。
  - **所有权**: 采用两步式所有权转移 (`Ownable2Step`)，并在部署后将所有权赋予内部创建的 `TimelockController` 实例。
  - 详细技术说明请参阅 [MPERproxy_technical_docs.md](./MPERproxy_technical_docs.md)。

### 脚本文件

#### `scripts/MPERdeploy.js`
- 主部署脚本。
- 功能：
  - 部署 `MPER.sol` 逻辑合约。
  - 部署 `MPERproxy.sol` 代理合约，并初始化（包括设置 `TimelockController`）。
  - 调用 `scripts/utils/deployment.js` 中的函数保存部署信息到 `deployments.json`。
  - （可选）执行初步的部署后验证。

#### `scripts/utils/deployment.js`
- 部署相关的辅助函数。
- 功能：
  - `saveDeployment()`: 将合约地址、交易哈希、实现地址等信息保存到 `deployments.json`。
  - `getDeployment()`: 从 `deployments.json` 读取已部署的合约信息。
  - `getImplementationAddress()`: 从代理合约的存储中读取当前逻辑合约（实现）的地址。

#### `scripts/utils/storage.js`
- 存储布局验证相关的辅助函数。
- 功能：
  - `verifyStorageLayout()`: 验证代理合约中关键状态变量的存储槽位是否符合预期，辅助确保升级兼容性。
  - `getImplementationAddress()`: (注意：此功能也存在于 `scripts/utils/deployment.js`，根据实际使用情况可能只保留一个或各有侧重)

#### `scripts/verify.js`
- 部署后验证脚本。
- 验证内容（示例）：
  - 确认合约已成功部署并获取地址。
  - 验证 `TimelockController` 是否正确设置为代理的拥有者。
  - 验证通过代理调用的某些只读函数是否返回预期值。

### 数据文件

#### `deployments.json`
- 存储各个网络上合约部署的详细信息，包括合约名称、地址、交易哈希、实现地址、部署时间戳等。
- 由 `scripts/utils/deployment.js` 中的 `saveDeployment` 函数生成和更新。

### 文档文件

#### `MPERproxy_technical_docs.md`
- `MPERproxy.sol` 合约的详细技术说明文档。
- 内容包括：合约架构、核心功能、状态变量、事件，特别是合约升级流程和 `TimelockController` 的使用方法及注意事项。

### 配置文件

#### `hardhat.config.js`
- Hardhat项目配置。
- 包含：
  - Solidity编译器版本设置。
  - 网络定义 (例如 `localhost`, `paperschain`, `mainnet` 等)。
  - Hardhat插件配置 (例如 `hardhat-ethers`, `@openzeppelin/hardhat-upgrades`)。
  - Gas报告器、Etherscan API密钥等。

#### `package.json`
- Node.js项目元数据和依赖管理。
- 关键依赖：
  - `@openzeppelin/contracts`: OpenZeppelin标准合约库。
  - `@openzeppelin/contracts-upgradeable`: OpenZeppelin可升级合约库。
  - `@openzeppelin/hardhat-upgrades`: Hardhat插件，用于管理可升级合约。
  - `hardhat`: 以太坊开发环境。
  - `ethers`: 以太坊JavaScript库。
  - 其他开发依赖如 `chai`, `dotenv` 等。

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
npx hardhat run scripts/MPERdeploy.js --network paperschain
```
4. 验证部署
```bash
npx hardhat run scripts/verify.js --network paperschain
```
5. 运行测试
```bash
npx hardhat test
```