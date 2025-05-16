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
  - [MPERproxy技术文档](MPERproxy_technical_docs.md)
  - [升级脚本](scripts/upgradeMPER.js)

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