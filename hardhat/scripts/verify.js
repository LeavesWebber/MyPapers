const { ethers } = require("hardhat");
const { getDeployment } = require("./utils/deployment");
const { verifyStorageLayout } = require("./utils/storage");

async function main() {
  const [deployer] = await ethers.getSigners();
  const deployments = require('../deployments.json');
  
  const network = hre.network.name;
  const proxyInfo = deployments[network]?.MPSProxy;
  
  if (!proxyInfo) {
    throw new Error(`找不到${network}网络上的MPSProxy部署记录`);
  }

  const MPSProxy = await ethers.getContractFactory("MPSproxy");
  const proxy = MPSProxy.attach(proxyInfo.address);

  // 验证所有者
  const owner = await proxy.owner();
  console.log(`代理合约所有者: ${owner}`);
  
  // 验证实现地址
  const implementation = await proxy.implementation();
  console.log(`当前实现地址: ${implementation}`);
  
  // 验证时间锁
  const timelock = await proxy.timelock();
  console.log(`时间锁地址: ${timelock}`);
}

main().catch(console.error);