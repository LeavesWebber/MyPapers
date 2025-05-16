const fs = require('fs');
const path = require('path');
const { ethers, network } = require("hardhat");
const { getDeployment, saveDeployment } = require("./utils/deployment");
const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});

async function prompt(question) {
  return new Promise((resolve) => {
    readline.question(question, (answer) => {
      resolve(answer);
    });
  });
}

async function main() {
  const [deployer] = await ethers.getSigners();
  const networkName = network.name;

  console.log(`网络: ${networkName}`);
  console.log(`操作账户 (应为Timelock管理员/提议者): ${deployer.address}`);

  // --- 配置 ---
  const newLogicContractName = "MPS"; // 已修改为MPS
  const upgradeReason = "Upgrading from MPER to MPS contract with new features."; // 更新升级原因
  // 如果新逻辑合约的 initialize 函数需要参数，在这里准备 ABI 编码的数据
  // 例如: const mperV2Interface = new ethers.Interface(["function initializeV2(uint256 newValue)"]);
  // const initDataForV2 = mperV2Interface.encodeFunctionData("initializeV2", [123]);
  const initDataForV2 = "0x"; // 如果 V2 没有初始化函数或不需要参数，则为 "0x"

  // 1. 加载现有部署信息
  const proxyDeployment = getDeployment(networkName, "MPERProxy");
  if (!proxyDeployment) {
    console.error(`错误: 在 deployments.json 中未找到 ${networkName} 网络的 MPERProxy 部署信息。`);
    readline.close();
    return;
  }
  if (!proxyDeployment.timelock || !proxyDeployment.initialAdmin) {
    console.error(`错误: deployments.json 中的 MPERProxy 条目缺少 'timelock' 或 'initialAdmin' 地址。请确保 MPERdeploy.js 正确保存了这些信息。`);
    readline.close();
    return;
  }

  const proxyAddress = proxyDeployment.address;
  const timelockAddress = proxyDeployment.timelock;
  const currentImplementation = proxyDeployment.implementation;

  // 2. 先获取TimelockController合约实例
  const TimelockController = await ethers.getContractFactory("TimelockController");
  const timelock = TimelockController.attach(timelockAddress);

  // 3. 然后再获取时间延迟参数
  const mperProxyUpgradeDelay = BigInt(proxyDeployment.upgradeDelay);
  const minTimelockDelay = BigInt(await timelock.getMinDelay());

  console.log(`MPERProxy 地址: ${proxyAddress}`);
  console.log(`TimelockController 地址: ${timelockAddress}`);
  console.log(`当前实现地址: ${currentImplementation}`);
  console.log(`MPERProxy 升级延迟: ${mperProxyUpgradeDelay} 秒`);




  // 3. 部署新的逻辑合约 (MPERv2)
  console.log(`\n部署新的逻辑合约 ${newLogicContractName}...`);
  const MPERv2 = await ethers.getContractFactory(newLogicContractName);
  const mperV2Logic = await MPERv2.deploy();
  const newImplementationAddress = await mperV2Logic.getAddress();
  console.log(`${newLogicContractName} 逻辑合约已提交部署到: ${newImplementationAddress}, 等待网络确认...`);
  // 简单的等待确认
  try {
    await mperV2Logic.waitForDeployment(); // 等待部署完成
    // 尝试调用一个只读方法确认
    if (typeof mperV2Logic.symbol === "function") { // 假设有 symbol 方法
        const symbol = await mperV2Logic.symbol();
        console.log(`${newLogicContractName} (${symbol}) 已确认.`);
    } else {
        console.log(`${newLogicContractName} 已部署，但无 symbol 方法用于快速确认。`);
    }
  } catch (e) {
    console.error(`读取 ${newLogicContractName} 状态失败。部署可能尚未完全传播:`, e);
    console.log("将等待15秒后重试...");
    await new Promise(resolve => setTimeout(resolve, 15000));
    try {
        if (typeof mperV2Logic.symbol === "function") {
            const symbol = await mperV2Logic.symbol();
            console.log(`${newLogicContractName} (${symbol}) 第二次尝试确认成功.`);
        } else {
             console.log(`${newLogicContractName} 已部署 (第二次尝试)。`);
        }
    } catch (e2) {
        console.error(`第二次读取 ${newLogicContractName} 状态仍然失败:`, e2);
        readline.close();
        throw new Error(`${newLogicContractName} logic contract not ready after deployment and delay.`);
    }
  }
  console.log(`新的逻辑合约 (${newLogicContractName}) 地址: ${newImplementationAddress}`);

  // 4. 准备调用 MPERproxy.proposeUpgrade 的数据
  const MPERProxy = await ethers.getContractFactory("MPERproxy"); // 用于获取接口
  const mperProxyInterface = MPERProxy.interface;
  const dataForProposeUpgrade = mperProxyInterface.encodeFunctionData("proposeUpgrade", [
    newImplementationAddress,
    initDataForV2,
    upgradeReason
  ]);

  // 5. 通过 TimelockController 调度对 MPERproxy.proposeUpgrade 的调用
  // 这是升级流程的第一阶段时间锁
  const saltPropose = ethers.id(`propose_upgrade_${newLogicContractName}_${Date.now()}`); // 唯一盐值
  const predecessorPropose = ethers.ZeroHash;

  console.log(`\n--- 阶段 1: 调度 'MPERproxy.proposeUpgrade' ---`);
  console.log(`目标 (MPERproxy): ${proxyAddress}`);
  console.log(`调用数据 (proposeUpgrade): ${dataForProposeUpgrade.substring(0, 100)}...`);
  console.log(`TimelockController 最小延迟: ${minTimelockDelay} 秒`);

  const scheduleProposeTx = await timelock.connect(deployer).schedule(
    proxyAddress,         // target: MPERproxy
    0,                    // value
    dataForProposeUpgrade, // data
    predecessorPropose,   // predecessor
    saltPropose,          // salt
    minTimelockDelay.toString()      // delay (使用 TimelockController 的最小延迟)
  );
  await scheduleProposeTx.wait();
  const proposalIdForProposeUpgrade = await timelock.hashOperation(
    proxyAddress, 0, dataForProposeUpgrade, predecessorPropose, saltPropose
  );
  console.log(`'MPERproxy.proposeUpgrade' 调用已调度到 TimelockController。`);
  console.log(`  交易哈希: ${scheduleProposeTx.hash}`);
  console.log(`  提案 ID (用于执行 proposeUpgrade): ${proposalIdForProposeUpgrade}`);
  const etaPropose = BigInt((await ethers.provider.getBlock('latest')).timestamp) + minTimelockDelay;
  console.log(`  预计可执行时间 (ETA for proposeUpgrade): ${new Date(Number(etaPropose) * 1000).toLocaleString()} (大约 ${Number(minTimelockDelay) / 3600} 小时后)`);

  console.log(`\n--- 后续步骤 ---`);
  console.log(`1. 等待上述 ETA 到达 (${new Date(Number(etaPropose) * 1000).toLocaleString()}).`);
  console.log(`2. 执行 'MPERproxy.proposeUpgrade':`);
  console.log(`   通过调用 TimelockController.execute(`);
  console.log(`     target: "${proxyAddress}",`);
  console.log(`     value: 0,`);
  console.log(`     data: "${dataForProposeUpgrade}",`);
  console.log(`     predecessor: "${predecessorPropose}",`);
  console.log(`     salt: "${saltPropose}"`);
  console.log(`   )`);
  console.log(`   此操作将使 MPERproxy 内部调用 timelock.schedule 来调度真正的 upgradeToAndCall。`);
  console.log(`   MPERproxy.proposeUpgrade 函数会返回一个 *新的* 提案 ID，记下它！`);
  console.log(`   这个新的提案 ID 是用于最终执行 upgradeToAndCall 的。`);

  console.log(`\n3. 假设上一步成功执行，并且你获得了由 MPERproxy.proposeUpgrade 返回的 'actualUpgradeProposalId'。`);
  console.log(`   MPERproxy 将使用其自身的 'upgradeDelay' (${mperProxyUpgradeDelay} 秒) 来调度 'upgradeToAndCall'。`);
  console.log(`   等待 MPERproxy 的 'upgradeDelay' 结束。`);

  console.log(`\n4. 执行实际升级 'MPERproxy.upgradeToAndCall':`);
  console.log(`   通过调用 TimelockController.execute(`);
  console.log(`     target: "${proxyAddress}", // MPERproxy 地址`);
  console.log(`     value: 0,`);
  console.log(`     data: MPERproxyInterface.encodeFunctionData("upgradeToAndCall", ["${newImplementationAddress}", "${initDataForV2}", "${upgradeReason}"]),`);
  console.log(`     predecessor: ethers.ZeroHash, // MPERproxy 内部调度时通常用 0`);
  console.log(`     salt: <actualUpgradeProposalId 的 salt 部分> // 需要从 MPERproxy.proposeUpgrade 返回的 proposalId 或事件中解析`);
  console.log(`   )`);
  console.log(`   更简单的方式是，直接使用 MPERproxy.proposeUpgrade 返回的 'actualUpgradeProposalId' 作为 TimelockController.execute 的最后一个参数（如果 TimelockController 的 execute 接受 operationId 作为参数，或者你需要用它来构造所有参数）。`);
  console.log(`   查阅 MPERproxy_technical_docs.md 中关于 TimelockController.execute 的说明，以及 MPERproxy.proposeUpgrade 的返回值/事件。`);
  console.log(`   通常，MPERproxy.proposeUpgrade 返回的 proposalId 就是你直接用于 TimelockController.execute() 的 operationId。`);


  console.log(`\n5. 升级完成后，验证新的实现地址。`);
  console.log(`6. 更新 deployments.json 中的 'implementation' 字段为 '${newImplementationAddress}'。`);
  console.log(`   例如: saveDeployment("${networkName}", { contract: "MPERProxy", address: proxyAddress, implementation: newImplementationAddress, /* 其他字段保持或更新 */ });`);

const upgradeState = {
  network: networkName,
  proxyAddress,
  newImplementation: newImplementationAddress,
  proposeUpgradeData: dataForProposeUpgrade,
  salt: saltPropose,
  eta: Number(etaPropose)
};

const statePath = path.join(__dirname, '../upgradeinfo/upgrade_state.json');
fs.writeFileSync(statePath, JSON.stringify(upgradeState, null, 2));
  readline.close();
}

main()
  .then(() => {
    readline.close();
  })
  .catch((error) => {
    console.error("升级脚本执行失败:", error);
    readline.close();
    process.exit(1);
  });