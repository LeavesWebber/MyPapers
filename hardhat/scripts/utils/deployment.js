const fs = require('fs');
const path = require('path');

// 定义部署记录文件的路径
const DEPLOYMENTS_PATH = path.join(__dirname, '../../deployments.json'); // 将部署文件放在项目根目录下的 deployments.json

// 确保 ethers 已被引入，例如在文件顶部:
// const { ethers } = require("hardhat");
// 或者如果只需要 ethers 中的 utils 并且 provider 是传递进来的:
// const { utils } = require("ethers");

// 假设 'ethers' 可用于 'ethers.utils'
const { ethers } = require("hardhat"); 

async function getImplementationAddress(provider, proxyAddress) { // 接受 provider 作为参数
  const implSlot = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc";
  // 使用传递进来的 provider 实例
  const implementationStorage = await provider.getStorage(proxyAddress, implSlot); 
  // 修正：使用 ethers v6 的方法
  return ethers.getAddress(ethers.dataSlice(implementationStorage, 12));
}

// 如果它是对象导出的一部分，请确保正确导出
// 例如:
// module.exports = {
//   // ... 其他函数
//   getImplementationAddress,
//   saveDeployment, // 确保其他导出项保持不变
//   verifyStorageLayout // 确保其他导出项保持不变
// };
// 如果你的这个文件中有其他函数如 saveDeployment 或 verifyStorageLayout，
// 请确保它们仍然被正确导出。上面的例子展示了如何做。
// 关键的改动仅仅是 getImplementationAddress 函数的签名和内部对 provider 的使用。
function saveDeployment(network, data) {
  let deployments = {};
  if (fs.existsSync(DEPLOYMENTS_PATH)) { // 现在 DEPLOYMENTS_PATH 已定义
    deployments = JSON.parse(fs.readFileSync(DEPLOYMENTS_PATH));
  }
  
  deployments[network] = deployments[network] || {};
  deployments[network][data.contract] = {
    address: data.address,
    txHash: data.txHash,
    timestamp: new Date().toISOString(),
    implementation: data.implementation,
    timelock: data.timelock, // 新增
    initialAdmin: data.initialAdmin, // 新增
    upgradeDelay: data.upgradeDelay // 新增
  };
  
  fs.writeFileSync(DEPLOYMENTS_PATH, JSON.stringify(deployments, null, 2)); // 现在 DEPLOYMENTS_PATH 已定义
}

function getDeployment(network, contractName) {
  if (!fs.existsSync(DEPLOYMENTS_PATH)) return null; // 现在 DEPLOYMENTS_PATH 已定义
  
  const deployments = JSON.parse(fs.readFileSync(DEPLOYMENTS_PATH)); // 现在 DEPLOYMENTS_PATH 已定义
  return deployments[network]?.[contractName] || null;
}

module.exports = { 
  saveDeployment,
  getDeployment,
  getImplementationAddress 
};