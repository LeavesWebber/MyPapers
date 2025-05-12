const { ethers } = require("hardhat");
const { saveDeployment, getImplementationAddress } = require("./utils/deployment");
const { verifyStorageLayout } = require("./utils/storage");

async function main() {
  // 1. 获取部署者账户
  const [deployer] = await ethers.getSigners();
  const provider = ethers.provider;
  console.log(`部署者地址: ${deployer.address}`);

  // 2. 部署逻辑合约
  console.log("正在部署MPS逻辑合约...");
  const MPS = await ethers.getContractFactory("MPS");
  const mpsLogic = await MPS.deploy();
  const logicContractAddress = await mpsLogic.getAddress();
  console.log(`MPS逻辑合约已提交部署到: ${logicContractAddress}, 等待网络确认...`);

  // 等待逻辑合约部署完成
  try {
    const symbol = await mpsLogic.symbol();
    console.log(`MPS逻辑合约 (${symbol}) 已确认.`);
  } catch (e) {
    console.error("读取MPS逻辑合约状态失败:", e);
    console.log("将等待15秒后重试...");
    await new Promise(resolve => setTimeout(resolve, 15000));
    try {
      const symbol = await mpsLogic.symbol();
      console.log(`MPS逻辑合约 (${symbol}) 第二次尝试确认成功.`);
    } catch (e2) {
      console.error("第二次读取MPS逻辑合约状态仍然失败:", e2);
      throw new Error("MPS logic contract not ready after deployment and delay.");
    }
  }

  // 3. 配置Timelock参数
  const initialDelay = 86400; // 24小时延迟

  // 4. 部署代理合约
  console.log("正在部署代理合约...");
  const MPSProxy = await ethers.getContractFactory("MPSproxy");
  
  // 创建初始化数据
  const initialSupply = ethers.parseUnits("500", 18);
  const initData = mpsLogic.interface.encodeFunctionData("initialize", [initialSupply]);
  
  const proxy = await MPSProxy.deploy(
    logicContractAddress,
    deployer.address,
    initData,
    initialDelay
  ).catch(err => {
    console.error("部署错误:", err);
    throw err;
  });

  const proxyContractAddress = await proxy.getAddress();
  console.log(`MPS代理合约已提交部署到: ${proxyContractAddress}, 等待网络确认...`);

  // 等待代理合约部署完成
  try {
    await proxy.timelock();
    console.log(`MPS代理合约 (timelock 可读) 已确认.`);
  } catch (e) {
    console.error("读取MPS代理合约状态失败:", e);
    console.log("将等待15秒后重试...");
    await new Promise(resolve => setTimeout(resolve, 15000));
    try {
      await proxy.timelock();
      console.log(`MPS代理合约 (timelock 可读) 第二次尝试确认成功.`);
    } catch (e2) {
      console.error("第二次读取MPS代理合约状态仍然失败:", e2);
      throw new Error("MPS proxy contract not ready after deployment and delay.");
    }
  }

  // 5. 获取实现地址
  const implementationAddress = await getImplementationAddress(provider, proxyContractAddress);

  // 6. 验证存储布局
  console.log("验证存储布局...");
  await verifyStorageLayout(proxyContractAddress, {
    _owner: 0,           // OwnableUpgradeable的_owner
    _hashToAddress: 1,  // 哈希映射
    _reviewToAddress: 2, // 审稿映射
    _hasRegistered: 3,   // 用户注册状态
    __gap: 4            // 预留空间
  });

  // 7. 保存部署记录
  const timelockAddress = await proxy.timelock();
  const deploymentTx = proxy.deploymentTransaction();
  const txHash = deploymentTx ? deploymentTx.hash : "N/A";

  await saveDeployment("local", {
    contract: "MPSProxy",
    address: proxyContractAddress,
    txHash: txHash,
    implementation: implementationAddress,
    timelock: timelockAddress,
    initialAdmin: deployer.address,
    upgradeDelay: initialDelay
  });

  // 8. 初始化合约状态 - 将初始供应量转移到合约地址用于用户注册奖励
  const mpsViaProxy = MPS.attach(proxyContractAddress);
  await mpsViaProxy.transfer(proxyContractAddress, initialSupply);

  console.log(`
✅ 部署完成
====================================
逻辑合约地址: ${logicContractAddress}
代理合约地址: ${proxyContractAddress}
时间锁合约地址: ${timelockAddress}
当前实现地址: ${implementationAddress}
临时管理员: ${deployer.address}
升级延迟: ${initialDelay}秒 (${initialDelay/3600}小时)
通过代理调用的总供应量: ${(await mpsViaProxy.totalSupply()).toString()}
逻辑合约自身的总供应量: ${(await mpsLogic.totalSupply()).toString()}
`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("部署失败:", error);
    process.exit(1);
  });