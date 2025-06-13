const { ethers, network } = require("hardhat");
const fs = require("fs");
const path = require("path");

async function getImplementationAddress(proxyAddress) {
    const implStorageSlot = '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc';
    const implHex = await ethers.provider.getStorage(proxyAddress, implStorageSlot);
    const implAddress = ethers.utils.getAddress(ethers.utils.hexDataSlice(implHex, 12));
    return implAddress;
}

async function main() {
  // 从环境变量获取状态文件路径
  const statePath = process.env.STATE_PATH;
  
  if (!statePath) {
    console.error("错误: 请设置 STATE_PATH 环境变量");
    console.log("示例: set STATE_PATH=\"./upgrades/...\" && npx hardhat run scripts/executeUpgrade.js --network localhost");
    process.exit(1);
  }

  const sanitizedPath = statePath.replace(/['"]/g, '');
  console.log(`从路径加载状态文件: ${sanitizedPath}`);
  
  let state;
  try {
    const fileContent = fs.readFileSync(sanitizedPath);
    state = JSON.parse(fileContent);
  } catch (e) {
    console.error("无法读取状态文件:", e.message);
    process.exit(1);
  }
  
  // 验证状态文件是否包含必要字段
  const requiredFields = [
    'proxyAddress', 
    'newImplementation', 
    'proxyContract', 
    'newLogicContract',
    'baseURI',
    'royaltyPercent'
  ];
  
  for (const field of requiredFields) {
    if (!state[field] && field !== 'royaltyPercent' && field !== 'baseURI') {
      console.error(`状态文件缺少必要字段: ${field}`);
      process.exit(1);
    }
  }
  
  console.log("====================================");
  console.log(`执行 NFT 升级`);
  console.log("====================================");
  console.log(`网络: ${state.network || network.name}`);
  console.log(`代理合约: ${state.proxyContract} @ ${state.proxyAddress}`);
  console.log(`新逻辑合约: ${state.newLogicContract} @ ${state.newImplementation}`);
  
  if (state.baseURI) {
    console.log(`基础URI: ${state.baseURI}`);
  }
  
  if (state.royaltyPercent !== undefined) {
    console.log(`版税比例: ${state.royaltyPercent}%`);
  }
  
  if (state.upgradeReason) {
    console.log(`升级原因: ${state.upgradeReason}`);
  }
  
  if (state.upgradeEta) {
    console.log(`预计升级执行时间: ${new Date(state.upgradeEta * 1000).toLocaleString()}`);
  }
  
  console.log(`当前时间: ${new Date().toLocaleString()}`);
  console.log("====================================");

  const [deployer] = await ethers.getSigners();
  console.log(`操作账户: ${deployer.address}`);

  // 获取代理合约实例 - 使用您的代理合约ABI
  const proxyABI = [
    "function timelock() view returns (address)",
    "function upgradeToAndCall(address newImplementation, bytes memory data, string memory reason) payable",
    "function implementation() view returns (address)",
    "function getUpgradeEta() public view returns (uint256)",
    "event Upgraded(address indexed newImplementation, uint256 timestamp, string reason)"
  ];
  
  const proxy = new ethers.Contract(state.proxyAddress, proxyABI, deployer);
  
  // 1. 检查升级计划时间
  try {
    const upgradeEta = await proxy.getUpgradeEta();
    console.log(`当前升级计划时间: ${upgradeEta > 0 ? new Date(Number(upgradeEta) * 1000).toLocaleString() : '无'}`);
    
    if (upgradeEta > 0) {
      // 检查是否到了执行时间
      const currentTime = Math.floor(Date.now() / 1000);
      if (currentTime < upgradeEta) {
        const remaining = upgradeEta - BigInt(currentTime);
        console.error(`错误: 升级延迟期还未结束，请等待 ${remaining} 秒`);
        console.log(`可执行时间: ${new Date(Number(upgradeEta) * 1000).toLocaleString()}`);
        console.log(`您可以使用脚本推进时间: npx hardhat run scripts/advanceTime.js --network localhost`);
        process.exit(1);
      }
    }
  } catch (e) {
    console.warn("无法获取升级计划时间:", e.message);
  }
  
  // 2. 准备初始化数据 (NFT特有)
  let initData;
  if (state.baseURI || state.royaltyPercent !== undefined) {
    const nftInterface = new ethers.Interface([
      "function initializeV2(string memory baseURI_)",
      "function initializeRoyalty(uint96 royaltyFraction_)"
    ]);
    
    if (state.baseURI) {
      initData = nftInterface.encodeFunctionData("initializeV2", [state.baseURI]);
      console.log(`初始化数据 - 基础URI: ${state.baseURI}`);
    } else if (state.royaltyPercent !== undefined) {
      const royaltyPoints = state.royaltyPercent * 100; // 转换为基点 (100 = 1%)
      initData = nftInterface.encodeFunctionData("initializeRoyalty", [royaltyPoints]);
      console.log(`初始化数据 - 版税比例: ${royaltyPoints} 基点 (${state.royaltyPercent}%)`);
    }
  } else {
    initData = "0x"; // 无初始化数据
    console.log("无额外初始化数据");
  }
  
  // 3. 获取时间锁合约
  let timelock;
  try {
    const timelockAddress = await proxy.timelock();
    const Timelock = await ethers.getContractFactory("TimelockController");
    timelock = Timelock.attach(timelockAddress);
    console.log(`时间锁合约地址: ${timelockAddress}`);
    
    // 检查调用者是否为时间锁管理员
    const isAdmin = await timelock.hasRole(await timelock.PROPOSER_ROLE(), deployer.address);
    if (!isAdmin) {
      console.error(`错误: 操作账户 ${deployer.address} 不是时间锁管理员`);
      console.log("请确保使用具有PROPOSER_ROLE的账户操作");
      process.exit(1);
    }
  } catch (e) {
    console.error("无法获取时间锁合约:", e.message);
    console.log("尝试直接执行升级...");
  }
  
  // 4. 执行升级
  try {
    let tx;
    if (timelock) {
      console.log("\n通过时间锁执行升级...");
      
      // 准备时间锁调用数据
      const dataForTimelock = proxy.interface.encodeFunctionData(
        "upgradeToAndCall",
        [state.newImplementation, initData, state.upgradeReason || "ERC721 Upgrade"]
      );
      
      // 生成唯一的salt
      const salt = ethers.id(Date.now().toString());
      
      tx = await timelock.execute(
        state.proxyAddress,
        0,
        dataForTimelock,
        ethers.ZeroHash,
        salt
      );
    } else {
      console.log("\n直接调用代理合约升级...");
      
      // 直接调用代理合约的升级方法
      tx = await proxy.upgradeToAndCall(
        state.newImplementation,
        initData,
        state.upgradeReason || "ERC721 Direct Upgrade"
      );
    }
    
    // 等待交易确认
    const receipt = await tx.wait();
    console.log(`✅ 升级交易成功! 交易哈希: ${tx.hash}`);
    
    // 解析事件
    const eventLogs = receipt.logs.filter(
      log => log.address.toLowerCase() === state.proxyAddress.toLowerCase()
    );
    
    if (eventLogs.length > 0) {
      const iface = new ethers.Interface(proxyABI);
      for (const log of eventLogs) {
        try {
          const parsedLog = iface.parseLog(log);
          if (parsedLog && parsedLog.name === "Upgraded") {
            console.log(`升级事件触发: 新实现地址 ${parsedLog.args.newImplementation}`);
            console.log(`原因: ${parsedLog.args.reason}`);
          }
        } catch (e) {
          // 无法解析的事件日志
        }
      }
    }
    
    // 5. 验证升级
    const currentImplementation = state.implementation;
    console.log(`当前实现地址: ${currentImplementation}`);
    console.log(`预期实现地址: ${state.newImplementation}`);
    
    if (currentImplementation.toLowerCase() === state.newImplementation.toLowerCase()) {
      console.log("✅ 升级验证成功: 实现地址已更新");
    } else {
      console.warn("⚠️ 实现地址不匹配，请检查升级是否成功");
    }
    
    // 更新状态文件
    state.upgradeExecutedAt = Math.floor(Date.now() / 1000);
    state.upgradeTxHash = tx.hash;
    state.currentImplementation = currentImplementation;
    fs.writeFileSync(sanitizedPath, JSON.stringify(state, null, 2));
    
    // 6. 测试新合约功能
    console.log("\n开始新合约功能测试...");
    await testNFTFunctionality(state);
    
    console.log("\n🎉 NFT 升级成功! 状态文件已更新");
  } catch (e) {
    console.error("\n❌ 升级失败:", e.message);
    
    // 分析常见错误原因
    if (e.message.includes("Unauthorized")) {
      console.log("可能原因: 操作账户无权限执行升级");
      console.log("解决方案: 请确保通过时间锁执行，且账户具有PROPOSER_ROLE角色");
    } else if (e.message.includes("invalid value")) {
      console.log("可能原因: 初始化数据格式不正确");
      console.log("解决方案: 检查初始化参数与合约构造函数/初始化函数是否匹配");
    } else if (e.message.includes("not a contract")) {
      console.log("可能原因: 新实现地址不是有效的合约");
      console.log("解决方案: 重新部署逻辑合约并更新状态文件");
    } else if (e.message.includes("timelock")) {
      console.log("可能原因: 时间锁合约调用失败");
      console.log("解决方案: 检查时间锁配置和延迟时间");
    }
    
    console.log("\n建议解决方案:");
    console.log("1. 检查代理合约和时间锁合约状态");
    console.log("2. 确保已部署的新逻辑合约正确");
    console.log("3. 重新部署整个系统: npx hardhat clean && npx hardhat compile && npx hardhat run scripts/deploy.js");
    
    process.exit(1);
  }
}

/**
 * 测试NFT新合约功能
 */
async function testNFTFunctionality(state) {
  const [deployer] = await ethers.getSigners();
  
  try {
    // 获取NFT合约实例
    const NFT = await ethers.getContractFactory(state.newLogicContract);
    const nft = NFT.attach(state.proxyAddress);
    
    // 1. 测试元数据
    if (typeof nft.baseURI === "function" && state.baseURI) {
      const actualBaseURI = await nft.baseURI();
      console.log(`- 基础URI: ${actualBaseURI}`);
      if (actualBaseURI !== state.baseURI) {
        console.warn(`  警告: 基础URI不匹配 (预期: ${state.baseURI})`);
      }
    } else {
      console.log("- 跳过基础URI测试 (未提供或方法不存在)");
    }
    
    // 2. 测试版税功能
    if (typeof nft.royaltyInfo === "function" && state.royaltyPercent !== undefined) {
      const testSaleAmount = ethers.parseEther("1.0");
      const [receiver, royaltyAmount] = await nft.royaltyInfo(0, testSaleAmount); // tokenId=0
      
      const expectedRoyalty = testSaleAmount * BigInt(state.royaltyPercent) / 100n;
      
      console.log(`- 版税测试:`);
      console.log(`  接收者: ${receiver}`);
      console.log(`  金额: ${ethers.formatEther(royaltyAmount)} ETH`);
      console.log(`  预期: ${ethers.formatEther(expectedRoyalty)} ETH`);
      
      if (royaltyAmount !== expectedRoyalty) {
        console.warn(`  警告: 版税金额不匹配`);
      }
    } else {
      console.log("- 跳过版税测试 (未提供或方法不存在)");
    }
    
    // 3. 测试铸造功能
    if (typeof nft.safeMint === "function") {
      console.log("- 尝试铸造测试NFT...");
      const mintTx = await nft.safeMint(deployer.address, "Test Token");
      await mintTx.wait();
      
      const tokenId = await nft.tokenOfOwnerByIndex(deployer.address, 0);
      console.log(`  铸造成功! Token ID: ${tokenId}`);
      
      // 测试转移
      if (typeof nft.transferFrom === "function") {
        console.log("  测试转移...");
        const transferTx = await nft.transferFrom(
          deployer.address, 
          "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", // Hardhat账户1
          tokenId
        );
        await transferTx.wait();
        console.log("  转移成功!");
      }
    } else {
      console.log("- 跳过铸造测试 (方法不存在)");
    }
    
    // 4. 测试版本号
    if (typeof nft.version === "function") {
      const version = await nft.version();
      console.log(`- 合约版本: ${version}`);
    } else {
      console.log("- 跳过版本测试 (方法不存在)");
    }
    
    console.log("✅ 功能测试完成");
  } catch (e) {
    console.error("⚠️ 功能测试失败:", e.message);
    console.log("请手动验证合约功能");
  }
}

main().catch((error) => {
  console.error("执行失败:", error);
  process.exit(1);
});