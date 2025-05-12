const { ethers } = require("hardhat");
const { saveDeployment, getImplementationAddress } = require("./utils/deployment");
const { verifyStorageLayout } = require("./utils/storage");

async function main() {
  // 1. 获取部署者账户
  const [deployer] = await ethers.getSigners();
  const provider = ethers.provider; // 从 Hardhat 的 ethers 中获取 provider 实例
  console.log(`部署者地址: ${deployer.address}`);

  // 2. 部署逻辑合约
  console.log("正在部署MPER逻辑合约...");
  const MPER = await ethers.getContractFactory("MPER");
  const mperLogic = await MPER.deploy();
  const logicContractAddress = await mperLogic.getAddress(); // 获取逻辑合约地址
  console.log(`MPER逻辑合约已提交部署到: ${logicContractAddress}, 等待网络确认...`);

  // 等待逻辑合约部署完成并确认其代码已在链上可用
  // 通过尝试调用一个只读方法来确认
  try {
    const symbol = await mperLogic.symbol(); // 任何简单的只读调用都可以
    console.log(`MPER逻辑合约 (${symbol}) 已确认.`);
  } catch (e) {
    console.error("读取MPER逻辑合约状态失败。部署可能尚未完全传播到网络节点:", e);
    console.log("将等待几秒钟后重试...");
    await new Promise(resolve => setTimeout(resolve, 15000)); // 等待15秒
    try {
        const symbol = await mperLogic.symbol();
        console.log(`MPER逻辑合约 (${symbol}) 第二次尝试确认成功.`);
    } catch (e2) {
        console.error("第二次读取MPER逻辑合约状态仍然失败:", e2);
        throw new Error("MPER logic contract not ready after deployment and delay.");
    }
  }

  // 3. 配置Timelock参数
  const initialDelay = 86400; // 24小时延迟

  // 4. 部署代理合约
  console.log("正在部署代理合约...");
  const MPERProxy = await ethers.getContractFactory("MPERproxy");
  
  // 创建正确的初始化数据
  const initialSupply = 0; // 设置初始供应量 (例如 1 wei)
  // 如果想设置 1 MPER (假设18位小数): const initialSupply = ethers.parseUnits("1", 18);
  const initData = mperLogic.interface.encodeFunctionData("initialize", [initialSupply]);
  
  const proxy = await MPERProxy.deploy(
    logicContractAddress,
    deployer.address,
    initData,  // 修正: 使用生成的 initData 变量
    initialDelay
  ).catch(err => {
      console.error("Detailed deployment error:", {
          logicAddress: logicContractAddress,
          adminAddress: deployer.address,
          initialDelay: initialDelay,
          error: err,
          errorData: err.data,
          errorReason: err.reason
      });
      throw err;
  });
  // await proxy.deployed(); // 已移除

  const proxyContractAddress = await proxy.getAddress(); // 获取代理合约地址
  console.log(`MPER代理合约已提交部署到: ${proxyContractAddress}, 等待网络确认...`);

  // 等待代理合约部署完成并确认其代码已在链上可用
  // 通过尝试调用一个只读方法来确认 (e.g., timelock() itself)
  try {
    await proxy.timelock(); // 尝试调用一次以检查是否就绪
    console.log(`MPER代理合约 (timelock 可读) 已确认.`);
  } catch (e) {
    console.error("读取MPER代理合约状态失败 (例如 timelock)。部署可能尚未完全传播:", e);
    console.log("将等待几秒钟后重试...");
    await new Promise(resolve => setTimeout(resolve, 15000)); // 等待15秒
    try {
        await proxy.timelock(); // 第二次尝试
        console.log(`MPER代理合约 (timelock 可读) 第二次尝试确认成功.`);
    } catch (e2) {
        console.error("第二次读取MPER代理合约状态仍然失败:", e2);
        throw new Error("MPER proxy contract not ready after deployment and delay (timelock call failed).");
    }
  }

  // 5. 获取实现地址
  // 将 provider 实例传递给工具函数
  const implementationAddress = await getImplementationAddress(provider, proxyContractAddress);

  // 6. 验证存储布局
  console.log("验证存储布局...");
  await verifyStorageLayout(proxyContractAddress, {
    _hashToAddress: 0,
    _reviewToAddress: 1,
    __gap: 2
  });

  // 7. 保存部署记录
  const timelockAddress = await proxy.timelock(); // 此调用现在应该能成功
  const deploymentTx = proxy.deploymentTransaction(); // 获取部署交易对象
  const txHash = deploymentTx ? deploymentTx.hash : "N/A"; // 获取交易哈希

  await saveDeployment("local", {
    contract: "MPERProxy",
    address: proxyContractAddress,
    txHash: txHash,
    implementation: implementationAddress,
    timelock: timelockAddress,
    initialAdmin: deployer.address,
    upgradeDelay: initialDelay
  });

  // 获取通过代理调用的 totalSupply
  const mperViaProxy = MPER.attach(proxyContractAddress);

  console.log(`
✅ 部署完成
====================================
逻辑合约地址: ${logicContractAddress}
代理合约地址: ${proxyContractAddress}
时间锁合约地址: ${timelockAddress}
当前实现地址: ${implementationAddress}
临时管理员: ${deployer.address}
升级延迟: ${initialDelay}秒 (${initialDelay/3600}小时)
通过代理获取的总供应量: ${(await mperViaProxy.totalSupply()).toString()}
逻辑合约自身的总供应量: ${(await mperLogic.totalSupply()).toString()}
`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("部署失败:", error);
    process.exit(1);
  });