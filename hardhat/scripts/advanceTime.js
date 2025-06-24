// scripts/advanceTime.js
const { ethers } = require("hardhat");

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log(`操作账户: ${deployer.address}`);

  // 推进时间（24小时 = 86400秒）
  const timeToAdvance = 86400; // 24小时
  
  console.log(`推进时间 ${timeToAdvance} 秒 (${timeToAdvance/3600} 小时)...`);
  
  // 使用hardhat的evm_increaseTime
  await ethers.provider.send("evm_increaseTime", [timeToAdvance]);
  
  // 挖掘一个新区块来应用时间变化
  await ethers.provider.send("evm_mine");
  
  const currentBlock = await ethers.provider.getBlock('latest');
  console.log(`✅ 时间已推进到: ${new Date(currentBlock.timestamp * 1000).toLocaleString()}`);
  
  console.log("\n现在可以执行升级了!");
}

main().catch((error) => {
  console.error("时间推进失败:", error);
  process.exit(1);
});