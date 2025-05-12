const { ethers } = require("hardhat");
const { getDeployment } = require("./utils/deployment");
const { verifyStorageLayout } = require("./utils/storage");

async function main() {
  const network = "paperschain";
  const deployment = getDeployment(network, "MPERProxy");
  
  if (!deployment) {
    console.error("找不到部署记录");
    return;
  }

  console.log(`
🔍 验证部署
====================================
代理合约地址: ${deployment.address}
实现合约地址: ${deployment.implementation}
部署时间: ${deployment.timestamp}
`);

  // 验证存储布局
  console.log("正在验证存储布局...");
  const slots = await verifyStorageLayout(deployment.address, {
    _hashToAddress: 0,
    _reviewToAddress: 1,
    _hashCounter: 2
  });
  
  console.table(slots);
  
  // 验证代币供应量
  const proxy = await ethers.getContractFactory("MPER");
  
  console.log("验证初始化状态:");
  console.log("- 管理员:", await proxy.owner());
  console.log("- 总供应量:", await proxy.totalSupply());
  console.log("- 实现地址:", await upgrades.erc1967.getImplementationAddress(proxyAddress));
}

main().catch(console.error);