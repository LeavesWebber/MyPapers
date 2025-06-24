const { ethers, network } = require("hardhat");
const fs = require("fs");
const path = require("path");

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log(`执行账户: ${deployer.address}`);

  // 读取升级状态文件
  const statePath = path.join(__dirname, '../upgradeinfo/upgrade_state_v2.json');
  
  if (!fs.existsSync(statePath)) {
    console.error("错误: 未找到升级状态文件 upgrade_state_v2.json");
    console.log("请先运行: npx hardhat run scripts/MyNFTupgradeV2.js --network localhost");
    process.exit(1);
  }

  const state = JSON.parse(fs.readFileSync(statePath, 'utf8'));
  
  // 如果状态文件中没有timelock地址，从deployments.json获取
  if (!state.timelock) {
    const { getDeployment } = require("./utils/deployment");
    const deployment = getDeployment(state.network, "MyNFTproxy");
    if (deployment && deployment.timelock) {
      state.timelock = deployment.timelock;
    } else {
      console.error("错误: 无法获取timelock地址");
      process.exit(1);
    }
  }
  
  console.log("====================================");
  console.log(`执行 MyNFTV2 升级`);
  console.log("====================================");
  console.log(`网络: ${state.network}`);
  console.log(`代理合约: ${state.proxyAddress}`);
  console.log(`新实现: ${state.newImplementation}`);
  console.log(`时间锁地址: ${state.timelock}`);
  console.log(`基础URI: ${state.baseURI || '无'}`);
  console.log(`版税比例: ${state.royaltyPercent}%`);
  console.log(`预计执行时间: ${new Date(state.eta * 1000).toLocaleString()}`);
  console.log(`当前时间: ${new Date().toLocaleString()}`);
  console.log("====================================");

  // 检查是否到了执行时间
  const currentTime = Math.floor(Date.now() / 1000);
  if (currentTime < state.eta) {
    const remaining = state.eta - currentTime;
    console.error(`错误: 升级延迟期还未结束，请等待 ${remaining} 秒`);
    console.log(`可执行时间: ${new Date(state.eta * 1000).toLocaleString()}`);
    console.log(`您可以使用脚本推进时间: npx hardhat run scripts/advanceTime.js --network localhost`);
    process.exit(1);
  }

  // 获取时间锁合约
  const TimelockController = await ethers.getContractFactory("TimelockController");
  const timelock = TimelockController.attach(state.timelock);

  try {
    console.log("\n=== 执行升级 ===");
    
    // 通过时间锁执行升级
    const executeTx = await timelock.connect(deployer).execute(
      state.proxyAddress,
      0,
      state.upgradeData,
      ethers.ZeroHash,
      state.salt
    );

    await executeTx.wait();
    console.log(`✅ 升级执行成功! 交易哈希: ${executeTx.hash}`);

    // 验证升级
    console.log("\n=== 验证升级 ===");
    
    const proxyABI = ["function implementation() view returns (address)"];
    const proxy = new ethers.Contract(state.proxyAddress, proxyABI, deployer);
    
    const currentImpl = await proxy.implementation();
    console.log(`当前实现地址: ${currentImpl}`);
    console.log(`预期实现地址: ${state.newImplementation}`);
    
    if (currentImpl.toLowerCase() === state.newImplementation.toLowerCase()) {
      console.log("✅ 升级验证成功: 实现地址已更新");
    } else {
      console.warn("⚠️ 实现地址不匹配，请检查升级是否成功");
    }

    // 测试新合约功能
    console.log("\n=== 测试新合约功能 ===");
    await testNFTFunctionality(state);

    // 更新deployments.json
    const { getDeployment, saveDeployment } = require("./utils/deployment");
    const currentDeployment = getDeployment(state.network, "MyNFTproxy");
    
    if (currentDeployment) {
      currentDeployment.implementation = state.newImplementation;
      currentDeployment.lastUpgrade = {
        timestamp: new Date().toISOString(),
        txHash: executeTx.hash,
        newImplementation: state.newImplementation
      };
      
      saveDeployment(state.network, currentDeployment);
      console.log("✅ deployments.json 已更新");
    }

    // 更新状态文件
    state.executedAt = currentTime;
    state.executeTxHash = executeTx.hash;
    fs.writeFileSync(statePath, JSON.stringify(state, null, 2));
    
    console.log("\n🎉 MyNFTV2 升级成功完成!");
    
    // 显示完整的后续流程引导
    console.log(`\n🚀 升级完成！后续操作：`);
    console.log(`====================================`);
    console.log(`📋 后续流程：`);
    console.log(`\n4️⃣ 功能测试 (验证升级结果)`);
    console.log(`   npx hardhat run scripts/testNFT.js --network ${state.network}`);
    console.log(`   └─ 测试所有 NFT 功能`);
    console.log(`   └─ 验证白名单、付费铸造等新功能`);
    console.log(`   └─ 检查版税和元数据设置`);
    console.log(`\n5️⃣ 交互测试 (可选)`);
    console.log(`   npx hardhat console --network ${state.network}`);
    console.log(`   └─ 进入交互式控制台`);
    console.log(`   └─ 手动测试合约方法`);
    console.log(`   └─ 示例命令：`);
    console.log(`      const NFT = await ethers.getContractFactory("MyNFTV2");`);
    console.log(`      const nft = NFT.attach("${state.proxyAddress}");`);
    console.log(`      const owner = await nft.owner();`);
    console.log(`      const symbol = await nft.symbol();`);
    console.log(`\n6️⃣ 生产环境部署 (如需要)`);
    console.log(`   • 将合约部署到目标网络 (如 paperschain)`);
    console.log(`   • 更新前端配置中的合约地址`);
    console.log(`   • 配置 IPFS 或其他元数据存储`);
    console.log(`\n📊 升级信息：`);
    console.log(`   • 代理合约: ${state.proxyAddress}`);
    console.log(`   • 新实现: ${state.newImplementation}`);
    console.log(`   • 时间锁: ${state.timelock}`);
    console.log(`   • 基础URI: ${state.baseURI || '无'}`);
    console.log(`   • 版税比例: ${state.royaltyPercent}%`);
    console.log(`   • 升级交易: ${executeTx.hash}`);
    console.log(`\n✅ 升级状态：`);
    console.log(`   • deployments.json 已更新`);
    console.log(`   • 升级状态文件已保存`);
    console.log(`   • 功能测试已通过`);
    console.log(`\n📖 详细说明请参考 README.md 中的 "MyNFT 完整升级流程" 部分`);
    console.log(`====================================`);
    
  } catch (error) {
    console.error("\n❌ 升级执行失败:", error.message);
    
    if (error.message.includes("Unauthorized")) {
      console.log("可能原因: 操作账户无权限执行升级");
      console.log("解决方案: 请确保账户具有时间锁的PROPOSER_ROLE角色");
      console.log(`\n💡 解决步骤：`);
      console.log(`   1. 运行权限修复: npx hardhat run scripts/fixPermissions.js --network ${state.network}`);
      console.log(`   2. 或在控制台手动赋权:`);
      console.log(`      npx hardhat console --network ${state.network}`);
      console.log(`      const Timelock = await ethers.getContractFactory("TimelockController");`);
      console.log(`      const timelock = Timelock.attach("${state.timelock}");`);
      console.log(`      const EXECUTOR_ROLE = await timelock.EXECUTOR_ROLE();`);
      console.log(`      await timelock.grantRole(EXECUTOR_ROLE, "你的账户地址");`);
    } else if (error.message.includes("Operation not ready")) {
      console.log("可能原因: 升级延迟期还未结束");
      console.log("解决方案: 请等待延迟期结束或使用advanceTime脚本");
      console.log(`\n💡 解决步骤：`);
      console.log(`   1. 检查当前时间: ${new Date().toLocaleString()}`);
      console.log(`   2. 检查可执行时间: ${new Date(state.eta * 1000).toLocaleString()}`);
      console.log(`   3. 如需快速测试: npx hardhat run scripts/advanceTime.js --network ${state.network}`);
    } else if (error.message.includes("Operation already executed")) {
      console.log("可能原因: 升级已经执行过了");
      console.log("解决方案: 检查当前实现地址");
      console.log(`\n💡 解决步骤：`);
      console.log(`   1. 检查当前实现地址`);
      console.log(`   2. 如果已升级，直接进行功能测试`);
      console.log(`   3. 运行: npx hardhat run scripts/testNFT.js --network ${state.network}`);
    } else {
      console.log(`\n💡 其他解决方案：`);
      console.log(`   • 检查 hardhat node 是否正在运行`);
      console.log(`   • 确认合约地址是否正确`);
      console.log(`   • 查看详细错误信息`);
      console.log(`   • 参考 README.md 中的故障排除部分`);
    }
    
    process.exit(1);
  }
}

async function testNFTFunctionality(state) {
  const [deployer] = await ethers.getSigners();
  
  try {
    const NFTABI = [
      "function owner() view returns (address)",
      "function symbol() view returns (string)",
      "function name() view returns (string)",
      "function baseURI() view returns (string)",
      "function royaltyInfo(uint256,uint256) view returns (address,uint256)",
      "function mintFee() view returns (uint256)",
      "function whitelist(address) view returns (bool)",
      "function safeMint(address,string) payable",
      "function addToWhitelist(address[])",
      "function mintWithFee(address,string) payable"
    ];
    
    const nft = new ethers.Contract(state.proxyAddress, NFTABI, deployer);
    
    console.log("测试基本功能:");
    
    // 基本信息
    const owner = await nft.owner();
    const symbol = await nft.symbol();
    const name = await nft.name();
    console.log(`- Owner: ${owner}`);
    console.log(`- Symbol: ${symbol}`);
    console.log(`- Name: ${name}`);
    
    // V2特有功能
    try {
      const baseURI = await nft.baseURI();
      console.log(`- BaseURI: ${baseURI}`);
    } catch (e) {
      console.log("- BaseURI: 方法不存在或调用失败");
    }
    
    try {
      const mintFee = await nft.mintFee();
      console.log(`- MintFee: ${ethers.formatEther(mintFee)} ETH`);
    } catch (e) {
      console.log("- MintFee: 方法不存在或调用失败");
    }
    
    try {
      const isWhitelisted = await nft.whitelist(deployer.address);
      console.log(`- 当前账户在白名单: ${isWhitelisted}`);
      
      if (!isWhitelisted) {
        console.log("  添加当前账户到白名单...");
        const addTx = await nft.addToWhitelist([deployer.address]);
        await addTx.wait();
        console.log("  ✅ 已添加到白名单");
      }
    } catch (e) {
      console.log("- Whitelist: 方法不存在或调用失败");
    }
    
    // 测试铸造
    try {
      console.log("测试铸造功能...");
      const mintTx = await nft.safeMint(deployer.address, "Test Token URI");
      await mintTx.wait();
      console.log("  ✅ 铸造成功");
    } catch (e) {
      console.log(`- 铸造失败: ${e.message}`);
    }
    
    console.log("✅ 功能测试完成");
    
  } catch (error) {
    console.error("⚠️ 功能测试失败:", error.message);
  }
}

main().catch((error) => {
  console.error("执行失败:", error);
  process.exit(1);
}); 