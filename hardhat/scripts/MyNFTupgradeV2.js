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

  // 从deployments.json获取当前部署信息
  const proxyDeployment = getDeployment(networkName, "MyNFTproxy");
  if (!proxyDeployment) {
    console.error(`错误: 在 deployments.json 中未找到 ${networkName} 网络的 MyNFTproxy 部署信息。`);
    console.log(`\n💡 解决方案：`);
    console.log(`   1. 确保已运行 MyNFTdeploy.js 脚本`);
    console.log(`   2. 检查 deployments.json 文件是否存在`);
    console.log(`   3. 重新部署: npx hardhat run scripts/MyNFTdeploy.js --network ${networkName}`);
    readline.close();
    return;
  }

  const proxyAddress = proxyDeployment.address;
  const timelockAddress = proxyDeployment.timelock;
  const currentImplementation = proxyDeployment.implementation;

  console.log(`代理合约地址: ${proxyAddress}`);
  console.log(`时间锁地址: ${timelockAddress}`);
  console.log(`当前实现地址: ${currentImplementation}`);

  // 1. 部署新的MyNFTV2逻辑合约
  console.log("\n=== 部署新的MyNFTV2逻辑合约 ===");
  const MyNFTV2 = await ethers.getContractFactory("MyNFTV2");
  const nftV2Logic = await MyNFTV2.deploy();
  const newImplementationAddress = await nftV2Logic.getAddress();
  console.log(`MyNFTV2逻辑合约已部署到: ${newImplementationAddress}`);

  // 等待部署确认
  await nftV2Logic.waitForDeployment();
  console.log("MyNFTV2逻辑合约部署确认完成");

  // 2. 获取时间锁合约实例
  const TimelockController = await ethers.getContractFactory("TimelockController");
  const timelock = TimelockController.attach(timelockAddress);

  // 3. 准备初始化数据
  const baseURI = await prompt("输入新的基础元数据URI (例: 'ipfs://QmXYZ/'): ");
  const royaltyPercent = parseInt(await prompt("输入默认版税百分比 (0-100): ")) || 0;

  console.log(`\n配置参数:`);
  console.log(`- 基础URI: ${baseURI}`);
  console.log(`- 版税百分比: ${royaltyPercent}%`);

  // 准备初始化数据
  const nftInterface = new ethers.Interface([
    "function initializeV2(string memory baseURI_)",
    "function initializeRoyalty(uint96 royaltyFraction_)"
  ]);

  let initData;
  if (baseURI && royaltyPercent > 0) {
    // 如果有两个参数，使用multicall
    const calls = [
      nftInterface.encodeFunctionData("initializeV2", [baseURI]),
      nftInterface.encodeFunctionData("initializeRoyalty", [royaltyPercent * 100])
    ];
    const multicallInterface = new ethers.Interface([
      "function multicall(bytes[] calldata data)"
    ]);
    initData = multicallInterface.encodeFunctionData("multicall", [calls]);
  } else if (baseURI) {
    initData = nftInterface.encodeFunctionData("initializeV2", [baseURI]);
  } else if (royaltyPercent > 0) {
    initData = nftInterface.encodeFunctionData("initializeRoyalty", [royaltyPercent * 100]);
  } else {
    initData = "0x"; // 无初始化数据
  }

  // 4. 准备升级调用数据
  const NFTProxy = await ethers.getContractFactory("MyNFTproxy");
  const upgradeData = NFTProxy.interface.encodeFunctionData("upgradeToAndCall", [
    newImplementationAddress,
    initData,
    "Upgrading to MyNFTV2 with enhanced features"
  ]);

  // 5. 通过时间锁调度升级
  console.log("\n=== 调度升级 ===");
  const salt = ethers.id(`upgrade_MyNFTV2_${Date.now()}`);
  const predecessor = ethers.ZeroHash;
  const minDelay = await timelock.getMinDelay();

  console.log(`调度升级调用...`);
  console.log(`- 目标: ${proxyAddress}`);
  console.log(`- 新实现: ${newImplementationAddress}`);
  console.log(`- 延迟: ${minDelay} 秒`);

  const scheduleTx = await timelock.connect(deployer).schedule(
    proxyAddress,
    0,
    upgradeData,
    predecessor,
    salt,
    minDelay.toString()
  );

  await scheduleTx.wait();
  console.log(`✅ 升级已调度! 交易哈希: ${scheduleTx.hash}`);

  // 计算执行时间
  const currentBlock = await ethers.provider.getBlock('latest');
  const eta = currentBlock.timestamp + Number(minDelay);
  console.log(`预计可执行时间: ${new Date(eta * 1000).toLocaleString()}`);

  // 6. 保存升级状态
  const upgradeState = {
    network: networkName,
    proxyAddress: proxyAddress,
    newImplementation: newImplementationAddress,
    proxyContract: "MyNFTproxy",
    newLogicContract: "MyNFTV2",
    baseURI: baseURI,
    royaltyPercent: royaltyPercent,
    upgradeReason: "Upgrading to MyNFTV2 with enhanced features",
    upgradeData: upgradeData,
    salt: salt,
    eta: eta,
    scheduleTxHash: scheduleTx.hash
  };

  const fs = require('fs');
  const path = require('path');
  const statePath = path.join(__dirname, '../upgradeinfo/upgrade_state_v2.json');
  fs.writeFileSync(statePath, JSON.stringify(upgradeState, null, 2));

  console.log(`\n✅ 升级状态已保存到: ${statePath}`);
  
  // 7. 显示完整的后续流程引导
  console.log(`\n🚀 升级调度完成！下一步操作：`);
  console.log(`====================================`);
  console.log(`📋 剩余升级流程：`);
  console.log(`\n2️⃣ 权限检查 (如需要)`);
  console.log(`   npx hardhat run scripts/fixPermissions.js --network ${networkName}`);
  console.log(`   └─ 检查并修复 TimelockController 权限`);
  console.log(`   └─ 确保账户具有 PROPOSER_ROLE 和 EXECUTOR_ROLE`);
  console.log(`   └─ 如果权限正常，可以跳过此步骤`);
  console.log(`\n3️⃣ 执行升级 (完成升级)`);
  console.log(`   npx hardhat run scripts/executeUpgradeV2.js --network ${networkName}`);
  console.log(`   └─ 通过时间锁执行升级操作`);
  console.log(`   └─ 验证升级是否成功`);
  console.log(`   └─ 自动测试新合约功能`);
  console.log(`\n4️⃣ 功能测试 (验证升级结果)`);
  console.log(`   npx hardhat run scripts/testNFT.js --network ${networkName}`);
  console.log(`   └─ 测试所有 NFT 功能`);
  console.log(`   └─ 验证白名单、付费铸造等新功能`);
  console.log(`\n5️⃣ 交互测试 (可选)`);
  console.log(`   npx hardhat console --network ${networkName}`);
  console.log(`   └─ 进入交互式控制台`);
  console.log(`   └─ 手动测试合约方法`);
  console.log(`\n⏰ 时间管理：`);
  console.log(`   • 当前升级延迟: ${minDelay} 秒`);
  console.log(`   • 预计可执行时间: ${new Date(eta * 1000).toLocaleString()}`);
  if (minDelay > 60) {
    console.log(`   • 如需快速测试，可使用: npx hardhat run scripts/advanceTime.js --network ${networkName}`);
  }
  console.log(`\n⚠️  注意事项：`);
  console.log(`   • 确保 hardhat node 持续运行`);
  console.log(`   • 不要重启网络，否则需要重新部署`);
  console.log(`   • 如果遇到权限错误，先运行 fixPermissions.js`);
  console.log(`\n📖 详细说明请参考 README.md 中的 "MyNFT 完整升级流程" 部分`);
  console.log(`====================================`);
}

main()
  .then(() => {
    readline.close();
  })
  .catch((error) => {
    console.error("升级脚本执行失败:", error);
    console.log(`\n💡 常见问题解决方案：`);
    console.log(`   • 权限问题: 运行 fixPermissions.js`);
    console.log(`   • 网络问题: 确保 hardhat node 正在运行`);
    console.log(`   • 合约问题: 检查合约是否已正确部署`);
    console.log(`   • 部署问题: 确保已运行 MyNFTdeploy.js`);
    console.log(`\n🔧 具体解决步骤：`);
    console.log(`   1. 检查 hardhat node: npx hardhat node`);
    console.log(`   2. 重新部署: npx hardhat run scripts/MyNFTdeploy.js --network localhost`);
    console.log(`   3. 修复权限: npx hardhat run scripts/fixPermissions.js --network localhost`);
    console.log(`   4. 重新升级: npx hardhat run scripts/MyNFTupgradeV2.js --network localhost`);
    console.log(`\n📖 详细说明请参考 README.md 中的故障排除部分`);
    readline.close();
    process.exit(1);
  }); 