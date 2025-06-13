const fs = require('fs');
const path = require('path');
const { ethers, network } = require("hardhat");
const { getDeployment, saveDeployment } = require("./utils/deployment");
const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});

async function prompt(question) {
  return new Promise((resolve) => {
    readline.question(question, (answer) => {
      resolve(answer);
    });
  });
}

async function main() {
  const [deployer] = await ethers.getSigners();
  const networkName = network.name;

  console.log(`网络: ${networkName}`);
  console.log(`操作账户: ${deployer.address}`);
  
  // ============ ERC721 专有配置 ============
  const currentContractName = "MyNFT"; // 当前NFT合约名称
  const newLogicContractName = "MyNFTV2"; // 新NFT逻辑合约名称
  const proxyContractName = "MyNFTproxy"; // ERC721代理合约名称
  const upgradeReason = "Upgrading to NFT v2 with royalties, whitelisting and enhanced metadata support";
  
  // ERC721 特定的初始化数据（使用避免函数重载的方法）
  const nftInterface = new ethers.Interface([
    // 明确的函数签名
    "function initializeV2(string memory baseURI_)",
    "function initializeRoyalty(uint96 royaltyFraction_)",
    "function multicall(bytes[] calldata data)"
  ]);
  
  // 提示用户输入NFT特定初始化参数
  const baseURI = await prompt("输入新的基础元数据URI (例: 'ipfs://QmXYZ/'): ");
  let royaltyPercent;
  
  try {
    royaltyPercent = parseInt(await prompt("输入默认版税百分比 (0-100): "));
    if (isNaN(royaltyPercent) || royaltyPercent < 0 || royaltyPercent > 100) {
      throw new Error("Invalid royalty percentage");
    }
  } catch (e) {
    console.log("使用默认版税值: 0%");
    royaltyPercent = 0;
  }

  // 准备初始化数据
  let initDataForV2;
  if (baseURI && royaltyPercent !== undefined) {
    // 创建多重初始化数组
    const calls = [
      nftInterface.encodeFunctionData("initializeV2", [baseURI]),
      nftInterface.encodeFunctionData("initializeRoyalty", [royaltyPercent * 100]) // 转换为基点 (100 = 1%)
    ];
    
    // 使用多重调用包装
    initDataForV2 = nftInterface.encodeFunctionData("multicall", [calls]);
  } else if (baseURI) {
    // 只初始化基础URI
    initDataForV2 = nftInterface.encodeFunctionData("initializeV2", [baseURI]);
  } else if (royaltyPercent !== undefined) {
    // 只初始化版税
    initDataForV2 = nftInterface.encodeFunctionData("initializeRoyalty", [royaltyPercent * 100]);
  } else {
    // 无初始化
    initDataForV2 = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512";
    console.log("警告: 未提供版税或基础URI初始化参数");
  }

  // 1. 加载NFT合约部署信息
  const proxyDeployment = getDeployment(networkName, proxyContractName);
  if (!proxyDeployment) {
    console.error(`错误: 未找到 ${proxyContractName} 部署信息`);
    console.log("尝试从最新部署中查找...");
    
    // 回退逻辑：从部署历史中查找最新代理
    const deployments = getDeploymentHistory(networkName);
    const latestNFTDeployment = deployments.find(d => d.contract.includes("NFT"));
    
    if (latestNFTDeployment) {
      console.log(`找到最近的NFT部署: ${latestNFTDeployment.contract} @ ${latestNFTDeployment.address}`);
      proxyDeployment = {
        ...latestNFTDeployment,
        timelock: latestNFTDeployment.timelock || "<unknown>",
        upgradeDelay: latestNFTDeployment.upgradeDelay || 0
      };
    } else {
      throw new Error("无法定位NFT代理合约");
    }
  }

  const proxyAddress = proxyDeployment.address;
  const timelockAddress = proxyDeployment.timelock;
  const currentImplementation = proxyDeployment.implementation;
  
  if (!timelockAddress) {
    const fallbackAdmin = proxyDeployment.initialAdmin || deployer.address;
    console.log(`使用回退管理员地址作为时间锁: ${fallbackAdmin}`);
    proxyDeployment.timelock = fallbackAdmin;
  }

  // 获取时间锁合约实例
  const TimelockController = await ethers.getContractFactory("TimelockController");
  const timelock = TimelockController.attach(proxyDeployment.timelock);
  
  // 获取时间锁参数
  const nftUpgradeDelay = BigInt(proxyDeployment.upgradeDelay || 86400);
  const minTimelockDelay = BigInt(await timelock.getMinDelay());

  console.log(`================================`);
  console.log(`ERC721 升级概览`);
  console.log(`================================`);
  console.log(`代理合约: ${proxyContractName} @ ${proxyAddress}`);
  console.log(`时间锁合约: ${proxyDeployment.timelock}`);
  console.log(`当前实现: ${currentImplementation}`);
  console.log(`新逻辑合约: ${newLogicContractName}`);
  console.log(`升级原因: ${upgradeReason}`);
  console.log(`基础元数据URI: ${baseURI || "无更新"}`);
  console.log(`默认版税: ${royaltyPercent}%`);
  console.log(`时间锁最小延迟: ${minTimelockDelay} 秒`);
  console.log(`代理升级延迟: ${nftUpgradeDelay} 秒`);
  console.log(`================================`);

  // 确认用户继续
  const confirm = await prompt("确认升级参数? (y/n): ");
  if (confirm.toLowerCase() !== 'y') {
    console.log("升级取消");
    readline.close();
    return;
  }

  // 2. 部署新的NFT逻辑合约
  console.log(`\n部署新的NFT逻辑合约 ${newLogicContractName}...`);
  const NFTV2 = await ethers.getContractFactory(newLogicContractName);
  const nftV2Logic = await NFTV2.deploy();
  const newImplementationAddress = await nftV2Logic.getAddress();
  
  console.log(`等待部署确认...`);
  await nftV2Logic.waitForDeployment();
  
  // NFT特定确认逻辑
  try {
    // 检查NFT元数据
    const nftName = await nftV2Logic.name();
    const nftSymbol = await nftV2Logic.symbol();
    console.log(`${newLogicContractName} 确认: ${nftName} (${nftSymbol})`);
    
    // 检查版本信息（如有）
    if (typeof nftV2Logic.version === "function") {
      const version = await nftV2Logic.version();
      console.log(`合约版本: ${version}`);
    }
  } catch (e) {
    console.warn("无法读取NFT合约元数据，但部署成功:", e);
  }
  
  console.log(`新的NFT逻辑合约地址: ${newImplementationAddress}`);

  // 3. 准备调用代理合约的proposeUpgrade数据
  const NFTProxy = await ethers.getContractFactory(proxyContractName);
  const proxyInterface = NFTProxy.interface;
  
  // ERC721升级特有：确保包含NFT初始化数据
  const dataForProposeUpgrade = proxyInterface.encodeFunctionData("upgradeToAndCall", [
    newImplementationAddress,
    initDataForV2,
    upgradeReason
  ]);

  // 4. 通过时间锁调度升级
  const saltPropose = ethers.id(`nft_upgrade_${newLogicContractName}_${Date.now()}`);
  const predecessorPropose = ethers.ZeroHash;

  console.log(`\n[阶段1] 调度NFT升级提案...`);
  const scheduleProposeTx = await timelock.connect(deployer).schedule(
    proxyAddress,         // target: NFT代理合约
    0,                    // value
    dataForProposeUpgrade, // data
    predecessorPropose,   // predecessor
    saltPropose,          // salt
    minTimelockDelay.toString() // delay
  );
  
  await scheduleProposeTx.wait();
  const proposalIdForProposeUpgrade = await timelock.hashOperation(
    proxyAddress, 0, dataForProposeUpgrade, predecessorPropose, saltPropose
  );
  
  const etaPropose = BigInt((await ethers.provider.getBlock('latest')).timestamp) + minTimelockDelay;
  
  console.log(`✅ NFT升级提案已调度`);
  console.log(`  交易哈希: ${scheduleProposeTx.hash}`);
  console.log(`  提案ID: ${proposalIdForProposeUpgrade}`);
  console.log(`  执行时间: ${new Date(Number(etaPropose) * 1000).toLocaleString()}`);
  
  // 创建升级状态文件（含NFT专有信息）
  const upgradeState = {
    network: networkName,
    contractType: "ERC721",
    proxyAddress,
    proxyContract: proxyContractName,
    newImplementation: newImplementationAddress,
    newLogicContract: newLogicContractName,
    baseURI,
    royaltyPercent,
    upgradeReason,
    timelockAddress,
    minTimelockDelay: minTimelockDelay.toString(),
    nftUpgradeDelay: nftUpgradeDelay.toString(),
    proposeUpgradeData: dataForProposeUpgrade,
    salt: saltPropose,
    proposeUpgradeEta: Number(etaPropose),
    upgradeEta: Number(etaPropose + nftUpgradeDelay)
  };
  
  // 创建ERC721专用的升级目录
  const upgradeDir = path.join(__dirname, `../upgrades/${networkName}/ERC721`);
  if (!fs.existsSync(upgradeDir)) {
    fs.mkdirSync(upgradeDir, { recursive: true });
  }
  
  const statePath = path.join(upgradeDir, `upgrade_state-${Date.now()}.json`);
  fs.writeFileSync(statePath, JSON.stringify(upgradeState, null, 2));
  console.log(`升级状态保存至: ${statePath}`);

  // ============ 生成后续操作指南 ============
  console.log(`\n🎨 ERC721 升级后续步骤指南 🎨`);
  console.log(`=====================================================`);
  console.log(`1. [提案准备] 等待时间锁延迟结束 (${minTimelockDelay} 秒)`);
  console.log(`   预计执行时间: ${new Date(Number(etaPropose) * 1000).toLocaleString()}`);
  console.log(`\n2. [执行提案] 执行NFT升级提案调用:`);
  console.log(`   npx hardhat run scripts/executeNFTupgrade.js --network ${networkName} --state ${statePath}`);
  console.log(`   或手动调用时间锁合约的 execute 函数:`);
  console.log(`   target: ${proxyAddress}`);
  console.log(`   data: ${dataForProposeUpgrade.substring(0, 100)}...`);
  console.log(`   salt: ${saltPropose}`);
  console.log(`\n3. [等待升级延迟] NFT特有的升级延迟期 (${nftUpgradeDelay} 秒)`);
  console.log(`   预计升级执行时间: ${new Date(Number(upgradeState.upgradeEta) * 1000).toLocaleString()}`);
  console.log(`\n4. [最终升级] 执行最终升级:`);
  console.log(`   npx hardhat run scripts/executeUpgrade.js --network ${networkName} --state ${statePath}`);
  console.log(`\n5. [验证] 升级后执行NFT功能测试:`);
  console.log(`   npx hardhat test test/NFTupgrade.test.js --network ${networkName}`);
  
  // NFT升级后验证步骤
  console.log(`\n✅ 升级后验证清单:`);
  console.log(`   - 测试元数据解析: 检查基础URI是否更新`);
  console.log(`   - 测试版税计算: 使用tokenId=0验证royaltyInfo`);
  console.log(`   - 测试新NFT铸造: 使用safeMintWithRoyalty功能`);
  console.log(`   - 检查所有权转移: 确保已有NFT所有权未受影响`);
  
  // 保存部署信息更新
  const updatedDeployment = {
    ...upgradeState,
    ...proxyDeployment,
    implementation: newImplementationAddress,
    upgradedAt: Date.now(),
    upgradedBy: deployer.address,
    version: "v2", // 更新版本号
    upgradeReason:upgradeReason
  };
  
fs.writeFileSync(statePath, JSON.stringify(updatedDeployment, null, 2));
  readline.close();
  
  console.log("\n🎉 NFT升级流程初始化完成! 请按上述步骤完成升级过程");
  readline.close();
}

// 辅助函数: 获取部署历史
function getDeploymentHistory(networkName) {
  try {
    const historyPath = path.join(__dirname, `../deployments/${networkName}/history.json`);
    if (fs.existsSync(historyPath)) {
      return JSON.parse(fs.readFileSync(historyPath));
    }
    return [];
  } catch (e) {
    console.error("部署历史读取失败:", e);
    return [];
  }
}

main().catch((error) => {
  console.error("NFT升级初始化失败:", error);
  readline.close();
  process.exit(1);
});