// scripts/deployTimelock.js
const accountsConfig = require("../accounts-config");

async function main() {
  const admin = new ethers.Wallet(accountsConfig.admin.privateKey, ethers.provider);
  
  const Timelock = await ethers.getContractFactory("TimelockController");
  const timelock = await Timelock.deploy(
    minDelay,
    [admin.address],
    [admin.address],
    admin.address
  );

  // 设置永久权限
  const TIMELOCK_ADMIN_ROLE = await timelock.TIMELOCK_ADMIN_ROLE();
  const PROPOSER_ROLE = await timelock.PROPOSER_ROLE();
  const EXECUTOR_ROLE = await timelock.EXECUTOR_ROLE();
  
  // 永久锁定权限
  await timelock.renounceRole(TIMELOCK_ADMIN_ROLE, admin.address); // 保留管理员
  await timelock.revokeRole(PROPOSER_ROLE, admin.address); // 不能取消
  await timelock.revokeRole(EXECUTOR_ROLE, admin.address); // 不能取消

  console.log("Admin permissions locked:", admin.address);
}