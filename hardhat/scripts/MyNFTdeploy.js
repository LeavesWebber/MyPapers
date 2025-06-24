const { ethers } = require("hardhat");
const { saveDeployment, getImplementationAddress } = require("./utils/deployment");
const { verifyStorageLayout } = require("./utils/storage");
const network = hre.network.name;

async function main() {
  // 1. 获取部署者账户
  const [deployer] = await ethers.getSigners();
  const provider = ethers.provider;
  console.log(`部署者地址: ${deployer.address}`);

  // 2. 部署逻辑合约
  console.log("正在部署MyNFT逻辑合约...");
  const MyNFT = await ethers.getContractFactory("MyNFT");
  const nftLogic = await MyNFT.deploy();
  const logicContractAddress = await nftLogic.getAddress();
  console.log(`MyNFT逻辑合约已提交部署到: ${logicContractAddress}, 等待网络确认...`);

  // 等待逻辑合约部署完成
  try {
    const symbol = await nftLogic.symbol();
    console.log(`MyNFT逻辑合约 (${symbol}) 已确认.`);
  } catch (e) {
    console.error("读取MyNFT逻辑合约状态失败:", e);
    console.log("将等待15秒后重试...");
    await new Promise(resolve => setTimeout(resolve, 15000));
    try {
      const symbol = await nftLogic.symbol();
      console.log(`MyNFT逻辑合约 (${symbol}) 第二次尝试确认成功.`);
    } catch (e2) {
      console.error("第二次读取MyNFT逻辑合约状态仍然失败:", e2);
      throw new Error("MyNFT logic contract not ready after deployment and delay.");
    }
  }

  // 3. 配置Timelock参数
  const initialDelay = 1; // 24小时延迟

  // 4. 部署代理合约
  console.log("正在部署代理合约...");
  const NFTProxy = await ethers.getContractFactory("MyNFTproxy");

  // 创建初始化数据
  const initData = nftLogic.interface.encodeFunctionData("initialize", [deployer.address]);

  const proxy = await NFTProxy.deploy(
    logicContractAddress,
    deployer.address,
    initData,
    initialDelay
  ).catch(err => {
    console.error("部署错误:", err);
    throw err;
  });

  const proxyContractAddress = await proxy.getAddress();
  console.log(`MyNFT代理合约已提交部署到: ${proxyContractAddress}, 等待网络确认...`);

  // 等待代理合约部署完成
  try {
    await proxy.timelock();
    console.log(`MyNFT代理合约 (timelock 可读) 已确认.`);
  } catch (e) {
    console.error("读取MyNFT代理合约状态失败:", e);
    console.log("将等待15秒后重试...");
    await new Promise(resolve => setTimeout(resolve, 15000));
    try {
      await proxy.timelock();
      console.log(`MyNFT代理合约 (timelock 可读) 第二次尝试确认成功.`);
    } catch (e2) {
      console.error("第二次读取MyNFT代理合约状态仍然失败:", e2);
      throw new Error("MyNFT proxy contract not ready after deployment and delay.");
    }
  }

  // 5. 获取实现地址
  const implementationAddress = await getImplementationAddress(provider, proxyContractAddress);

  // 6. 验证存储布局
  console.log("验证存储布局...");
  await verifyStorageLayout(proxyContractAddress, {
    _nextTokenId: 0,
    contractMetadataURI: 1,
    royaltyPercentage: 2,
    _baseTokenURI: 3,
    __gap: 4
  });

  // 7. 保存部署记录
  const timelockAddress = await proxy.timelock();
  const deploymentTx = proxy.deploymentTransaction();
  const txHash = deploymentTx ? deploymentTx.hash : "N/A";

  await saveDeployment(network, {
    contract: "MyNFTproxy",
    address: proxyContractAddress,
    txHash: txHash,
    implementation: implementationAddress,
    timelock: timelockAddress,
    initialAdmin: deployer.address,
    upgradeDelay: initialDelay
  });

  // 8. 初始化合约状态（可选）
  // 这里不需要像 MPS 那样转移初始供应量

  console.log(`\n✅ 部署完成\n====================================\n逻辑合约地址: ${logicContractAddress}\n代理合约地址: ${proxyContractAddress}\n时间锁合约地址: ${timelockAddress}\n当前实现地址: ${implementationAddress}\n临时管理员: ${deployer.address}\n升级延迟: ${initialDelay}秒 (${initialDelay/3600}小时)\n`);

  // 9. 显示完整的升级流程引导
  console.log(`\n🚀 下一步：升级到 MyNFTV2\n====================================`);
  console.log(`📋 完整升级流程：`);
  console.log(`\n1️⃣ 调度升级 (部署 MyNFTV2 并准备升级)`);
  console.log(`   npx hardhat run scripts/MyNFTupgradeV2.js --network ${network}`);
  console.log(`   └─ 将部署新的 MyNFTV2 逻辑合约`);
  console.log(`   └─ 通过时间锁调度升级操作`);
  console.log(`   └─ 生成升级状态文件`);
  console.log(`\n2️⃣ 权限检查 (如需要)`);
  console.log(`   npx hardhat run scripts/fixPermissions.js --network ${network}`);
  console.log(`   └─ 检查并修复 TimelockController 权限`);
  console.log(`   └─ 确保账户具有 PROPOSER_ROLE 和 EXECUTOR_ROLE`);
  console.log(`\n3️⃣ 执行升级 (完成升级)`);
  console.log(`   npx hardhat run scripts/executeUpgradeV2.js --network ${network}`);
  console.log(`   └─ 通过时间锁执行升级操作`);
  console.log(`   └─ 验证升级是否成功`);
  console.log(`   └─ 自动测试新合约功能`);
  console.log(`\n4️⃣ 功能测试 (验证升级结果)`);
  console.log(`   npx hardhat run scripts/testNFT.js --network ${network}`);
  console.log(`   └─ 测试所有 NFT 功能`);
  console.log(`   └─ 验证白名单、付费铸造等新功能`);
  console.log(`\n5️⃣ 交互测试 (可选)`);
  console.log(`   npx hardhat console --network ${network}`);
  console.log(`   └─ 进入交互式控制台`);
  console.log(`   └─ 手动测试合约方法`);
  console.log(`\n⚠️  注意事项：`);
  console.log(`   • 如果升级延迟期未到，可使用 advanceTime.js 推进时间`);
  console.log(`   • 每次重启 hardhat node 后需要重新部署和赋权`);
  console.log(`   • 确保使用正确的网络和账户`);
  console.log(`\n📖 详细说明请参考 README.md 中的 "MyNFT 完整升级流程" 部分`);
  console.log(`====================================`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("部署失败:", error);
    process.exit(1);
  });