const { ethers } = require("hardhat");

async function verifyStorageLayout(proxyAddress, expectedSlots = {}) {
  const slots = {};
  for (const [name, slot] of Object.entries(expectedSlots)) {
    // 使用 ethers.provider.getStorage 替换 ethers.provider.getStorageAt
    const value = await ethers.provider.getStorage(proxyAddress, slot);
    slots[name] = { slot, value };
    
    if (value === '0x' + '0'.repeat(64)) {
      console.warn(`⚠️ 存储槽 ${slot} (${name}) 为空`);
    }
  }
  return slots;
}

async function getImplementationAddress(proxyAddress) {
  const IMPLEMENTATION_SLOT = '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc';
  // 使用 ethers.provider.getStorage 替换 ethers.provider.getStorageAt
  const storageValue = await ethers.provider.getStorage(proxyAddress, IMPLEMENTATION_SLOT);
  // 使用 ethers.getAddress 和 ethers.dataSlice 替换旧的 ethers.utils.getAddress 和手动切片
  return ethers.getAddress(ethers.dataSlice(storageValue, 12));
}

module.exports = { verifyStorageLayout, getImplementationAddress };