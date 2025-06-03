const { ethers } = require("hardhat");
const path = require("path");
const fs = require("fs");

// 路径标准化函数
function normalizePath(rawPath) {
  if (!rawPath) throw new Error("文件路径未提供");
  let normalized = rawPath.toString();
  
  // Windows 处理
  if (process.platform === 'win32') {
    normalized = normalized.replace(/\\/g, '/');
  }
  
  // 解决相对路径问题
  if (!path.isAbsolute(normalized)) {
    normalized = path.resolve(process.cwd(), normalized);
  }
  
  // 验证文件存在
  if (!fs.existsSync(normalized)) {
    throw new Error(`文件不存在: ${normalized}`);
  }
  
  return normalized;
}

async function main() {
  // 获取第一个参数作为状态文件
  const statePath = normalizePath(process.argv[2]);
  console.log(`处理升级文件: ${statePath}`);
  
  // 加载状态数据
  const state = require(statePath);
  
  // 执行升级逻辑
  const Timelock = await ethers.getContractFactory("TimelockController");
  const timelock = Timelock.attach(state.timelockAddress);
  
  const tx = await timelock.execute(
    state.proxyAddress,
    state.value || 0,
    state.dataForUpgrade,
    state.predecessor || ethers.ZeroHash,
    state.salt
  );
  
  const receipt = await tx.wait();
  console.log(`✅ 升级执行成功！交易哈希: ${tx.hash}`);
  
  process.exit(0);
}

main().catch(error => {
  console.error("执行失败:", error);
  process.exit(1);
});