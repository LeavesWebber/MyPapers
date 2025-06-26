const { ethers } = require("hardhat");
const hre = require("hardhat");
const fs = require("fs");
const path = require("path");

async function getTimelockAddress() {
  const networkName = hre.network.name;
  const deploymentsPath = path.resolve(__dirname, "../deployments.json");
  const contractName = process.env.CONTRACT_NAME;

  if (fs.existsSync(deploymentsPath)) {
    const deployments = JSON.parse(fs.readFileSync(deploymentsPath, "utf8"));
    if (deployments[networkName]) {
      // 1. 优先使用CONTRACT_NAME
      if (contractName && deployments[networkName][contractName]) {
        const timelock = deployments[networkName][contractName].timelock;
        if (timelock) {
          console.log(`[INFO] 使用指定合约 ${contractName} 的 timelock: ${timelock}`);
          return timelock;
        } else {
          console.log(`[WARN] 指定合约 ${contractName} 没有 timelock 字段`);
        }
      }
      // 2. fallback顺序查找MyNFTproxy、MPSProxy、MPERProxy
      const fallbackOrder = ["MyNFTproxy", "MPSProxy", "MPERProxy"];
      for (const name of fallbackOrder) {
        if (deployments[networkName][name] && deployments[networkName][name].timelock) {
          const timelock = deployments[networkName][name].timelock;
          console.log(`[INFO] 自动fallback使用合约 ${name} 的 timelock: ${timelock}`);
          return timelock;
        }
      }
      // 3. 遍历所有合约
      for (const contract in deployments[networkName]) {
        const timelock = deployments[networkName][contract].timelock;
        if (timelock) {
          console.log(`[INFO] 遍历所有合约，使用 ${contract} 的 timelock: ${timelock}`);
          return timelock;
        }
      }
      console.log(`[ERROR] ${networkName} 网络下没有任何合约包含 timelock 字段`);
    } else {
      console.log(`[ERROR] deployments.json 中没有 ${networkName} 网络的部署信息`);
    }
  } else {
    console.log(`[ERROR] 未找到 deployments.json 文件: ${deploymentsPath}`);
  }
  if (process.env.TIMELOCK_ADDRESS) {
    console.log(`[INFO] 使用环境变量 TIMELOCK_ADDRESS: ${process.env.TIMELOCK_ADDRESS}`);
    return process.env.TIMELOCK_ADDRESS;
  }
  throw new Error(
    `未找到 TimelockController 地址，请检查 deployments.json 的 ${networkName} 部分，或设置 CONTRACT_NAME/TIMELOCK_ADDRESS 环境变量`
  );
}

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log(`操作账户: ${deployer.address}`);

  // 获取 TimelockController 地址
  const timelockAddress = await getTimelockAddress();
  console.log(`\n=== TimelockController 地址: ${timelockAddress} ===`);

  const TimelockController = await ethers.getContractFactory("TimelockController");
  const timelock = TimelockController.attach(timelockAddress);

  // 获取角色
  const PROPOSER_ROLE = await timelock.PROPOSER_ROLE();
  const EXECUTOR_ROLE = await timelock.EXECUTOR_ROLE();
  const ADMIN_ROLE = await timelock.DEFAULT_ADMIN_ROLE();

  console.log(`PROPOSER_ROLE: ${PROPOSER_ROLE}`);
  console.log(`EXECUTOR_ROLE: ${EXECUTOR_ROLE}`);
  console.log(`ADMIN_ROLE: ${ADMIN_ROLE}`);

  // 检查当前权限
  console.log("\n=== 检查当前权限 ===");
  const hasProposerRole = await timelock.hasRole(PROPOSER_ROLE, deployer.address);
  const hasExecutorRole = await timelock.hasRole(EXECUTOR_ROLE, deployer.address);
  const hasAdminRole = await timelock.hasRole(ADMIN_ROLE, deployer.address);

  console.log(`当前账户是否有 PROPOSER_ROLE: ${hasProposerRole}`);
  console.log(`当前账户是否有 EXECUTOR_ROLE: ${hasExecutorRole}`);
  console.log(`当前账户是否有 ADMIN_ROLE: ${hasAdminRole}`);

  // 检查谁是管理员（兼容没有getRoleMember的情况）
  let adminRoleHolder = null;
  if (typeof timelock.getRoleMember === 'function') {
    try {
      adminRoleHolder = await timelock.getRoleMember(ADMIN_ROLE, 0);
      console.log(`ADMIN_ROLE 持有者: ${adminRoleHolder}`);
    } catch (e) {
      console.log("[WARN] getRoleMember 方法不可用，无法显示管理员地址");
    }
  } else {
    console.log("[WARN] getRoleMember 方法不存在，无法显示管理员地址");
  }

  // 如果当前账户是管理员，直接赋予权限
  if (hasAdminRole) {
    console.log("\n=== 当前账户是管理员，直接赋予权限 ===");

    if (!hasProposerRole) {
      console.log("赋予 PROPOSER_ROLE...");
      const tx1 = await timelock.grantRole(PROPOSER_ROLE, deployer.address);
      await tx1.wait();
      console.log("✅ PROPOSER_ROLE 已赋予");
    }

    if (!hasExecutorRole) {
      console.log("赋予 EXECUTOR_ROLE...");
      const tx2 = await timelock.grantRole(EXECUTOR_ROLE, deployer.address);
      await tx2.wait();
      console.log("✅ EXECUTOR_ROLE 已赋予");
    }
  } else {
    console.log("\n=== 当前账户不是管理员 ===");
    console.log("需要管理员账户来赋予权限");
    if (adminRoleHolder) {
      console.log(`管理员地址: ${adminRoleHolder}`);
    }

    // 尝试使用管理员账户
    const accounts = await ethers.getSigners();
    for (let i = 0; i < accounts.length; i++) {
      const account = accounts[i];
      const isAdmin = await timelock.hasRole(ADMIN_ROLE, account.address);
      if (isAdmin) {
        console.log(`找到管理员账户: ${account.address} (账户 ${i})`);

        // 使用管理员账户赋予权限
        const timelockWithAdmin = timelock.connect(account);

        if (!hasProposerRole) {
          console.log("赋予 PROPOSER_ROLE...");
          const tx1 = await timelockWithAdmin.grantRole(PROPOSER_ROLE, deployer.address);
          await tx1.wait();
          console.log("✅ PROPOSER_ROLE 已赋予");
        }

        if (!hasExecutorRole) {
          console.log("赋予 EXECUTOR_ROLE...");
          const tx2 = await timelockWithAdmin.grantRole(EXECUTOR_ROLE, deployer.address);
          await tx2.wait();
          console.log("✅ EXECUTOR_ROLE 已赋予");
        }

        break;
      }
    }
  }

  // 验证权限
  console.log("\n=== 验证权限 ===");
  const finalHasProposerRole = await timelock.hasRole(PROPOSER_ROLE, deployer.address);
  const finalHasExecutorRole = await timelock.hasRole(EXECUTOR_ROLE, deployer.address);

  console.log(`最终 - PROPOSER_ROLE: ${finalHasProposerRole}`);
  console.log(`最终 - EXECUTOR_ROLE: ${finalHasExecutorRole}`);

  if (finalHasProposerRole && finalHasExecutorRole) {
    console.log("\n✅ 权限修复完成！现在可以执行升级了");
  } else {
    console.log("\n❌ 权限修复失败，请手动检查");
  }
}

main().catch((error) => {
  console.error("权限修复失败:", error);
  process.exit(1);
}); 