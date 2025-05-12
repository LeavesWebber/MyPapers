# MPER/MPS代币项目文档

## 项目结构
MPER-project/  
├── contracts/ # 智能合约源文件  
│  ├── MPER.sol # 初始逻辑合约  
│  ├── MPERproxy.sol # 代理合约  
│  ├── MPS.sol # 新逻辑合约  
│  └── MPSproxy.sol # 新代理合约  
├── deployments.json # 记录合约部署信息的JSON文件  
├── scripts/ # 部署和管理脚本  
│  ├── MPERdeploy.js # 初始部署脚本  
│  ├── MPSdeploy.js # 新合约部署脚本  
│  ├── upgradeMPER.js # 升级脚本  
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
  - `MPERproxy_technical_docs.md`
  - `upgradeMPER.js`
## 主要更新内容
### v2.0 (MPS版本)
1. 新增功能 ：
   
   - 用户注册跟踪系统
   - 改进的存储结构设计
   - 增强的权限管理
2. 合约改进 ：
   
   - 简化升级流程
   - 增强安全性检查
   - 优化错误处理机制
3. 工具链更新 ：
   
   - 新增MPS部署脚本
   - 改进部署验证流程
4. 文档完善 ：
   
   - 新增升级流程文档
   - 完善技术说明
   - 更新项目结构图