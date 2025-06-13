const { ethers } = require("hardhat");
const { saveDeployment, getImplementationAddress } = require("./utils/deployment");
const { verifyStorageLayout } = require("./utils/storage"); // 添加存储布局验证

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

  // 等待逻辑合约部署确认
  try {
    const name = await nftLogic.name(); // ERC721标准函数
    console.log(`MyNFT逻辑合约 (${name}) 已确认.`);
  } catch (e) {
    console.error("读取MyNFT逻辑合约状态失败:", e);
    console.log("将等待几秒钟后重试...");
    await new Promise(resolve => setTimeout(resolve, 15000));
    try {
      const name = await nftLogic.name();
      console.log(`MyNFT逻辑合约 (${name}) 第二次尝试确认成功.`);
    } catch (e2) {
      console.error("第二次读取MyNFT逻辑合约状态仍然失败:", e2);
      throw new Error("MyNFT logic contract not ready after deployment and delay.");
    }
  }

  // 3. 设置代理合约参数
  const initialDelay = 1; // 24小时延迟

  // 4. 准备初始化数据
  const initData = MyNFT.interface.encodeFunctionData("initialize", [
    deployer.address // 初始所有者
  ]);

  // 5. 部署代理合约
  console.log("正在部署代理合约...");
  const NFTProxy = await ethers.getContractFactory("MyNFTproxy");
  
  const proxy = await NFTProxy.deploy(
    logicContractAddress,
    deployer.address,
    initData,
    initialDelay
  ).catch(err => {
    console.error("详细部署错误:", {
      logicAddress: logicContractAddress,
      adminAddress: deployer.address,
      initialDelay: initialDelay,
      error: err,
      errorData: err.data,
      errorReason: err.reason
    });
    throw err;
  });

  const proxyContractAddress = await proxy.getAddress();
  console.log(`MyNFT代理合约已提交部署到: ${proxyContractAddress}, 等待网络确认...`);

  // 等待代理合约部署确认
  try {
    await proxy.timelock();
    console.log(`MyNFT代理合约 (timelock 可读) 已确认.`);
  } catch (e) {
    console.error("读取MyNFT代理合约状态失败:", e);
    console.log("将等待几秒钟后重试...");
    await new Promise(resolve => setTimeout(resolve, 15000));
    try {
      await proxy.timelock();
      console.log(`MyNFT代理合约 (timelock 可读) 第二次尝试确认成功.`);
    } catch (e2) {
      console.error("第二次读取MyNFT代理合约状态仍然失败:", e2);
      throw new Error("MyNFT proxy contract not ready after deployment and delay.");
    }
  }

  // 6. 获取实现地址
  const implementationAddress = await getImplementationAddress(provider, proxyContractAddress);

  // 7. 验证存储布局（假设我们有相关工具函数）
  console.log("验证存储布局...");
  try {
    await verifyStorageLayout(proxyContractAddress, {
      // 根据实际存储布局定义槽位映射
      _owner: 0,
      _name: 1,
      _symbol: 2,
      // ... 其他存储槽
    });
  } catch (e) {
    console.warn("存储布局验证失败:", e);
  }

  // 8. 获取时间锁地址
  const timelockAddress = await proxy.timelock();

  // 9. 保存部署记录
  const deploymentTx = proxy.deploymentTransaction();
  const txHash = deploymentTx ? deploymentTx.hash : "N/A";
  

  await saveDeployment("localhost", {
    contract: "MyNFTproxy",
    address: proxyContractAddress,
    txHash: txHash,
    implementation: implementationAddress,
    timelock: timelockAddress,
    initialAdmin: deployer.address,
    upgradeDelay: initialDelay
  });

  // 10. 绑定代理实例到逻辑合约ABI
  const nft = MyNFT.attach(proxyContractAddress);

  // 11. 测试基本功能
  try {
    console.log("NFT名称:", await nft.name());
    console.log("NFT代号:", await nft.symbol());
    
    console.log("测试铸币...");
    const tx = await nft.safeMint(deployer.address, "ipfs://test1");
    const receipt = await tx.wait();
    
    // 事件解析
    const events = receipt.logs.map(log => {
        try {
            return nft.interface.parseLog(log);
        } catch (e) {
            return null;
        }
    }).filter(event => event !== null);
    
    const mintEvents = events.filter(e => e.name === "Minted" || e.name === "Transfer");
    
    if (mintEvents.length > 0) {
        const tokenId = mintEvents[0].args.tokenId || mintEvents[0].args[2];
        console.log("铸币ID:", tokenId.toString());
        console.log("代币所有者:", await nft.ownerOf(tokenId));
        console.log("代币URI:", await nft.tokenURI(tokenId));
    }
    
    console.log("总供应量:", (await nft.totalSupply()).toString());
  } catch (error) {
    console.error("初始化或测试失败:", error);
  }

  console.log(`
✅ 部署完成
====================================
逻辑合约地址: ${logicContractAddress}
代理合约地址: ${proxyContractAddress}
时间锁合约地址: ${timelockAddress}
当前实现地址: ${implementationAddress}
临时管理员: ${deployer.address}
升级延迟: ${initialDelay}秒 (${initialDelay/3600}小时)
`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("部署失败:", error);
    process.exit(1);
  });