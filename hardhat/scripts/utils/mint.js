const { ethers, network } = require("hardhat");
const { getDeployment } = require("./deployment"); // 假设 deployment.js 在同一目录下

async function main() {
  // 1. 获取操作账户 (应为合约的Owner)
  const [deployer] = await ethers.getSigners();
  const networkName = network.name;

  console.log(`网络: ${networkName}`);
  console.log(`操作账户 (Owner): ${deployer.address}`);

  // --- 配置铸币参数 ---
  // 请将以下地址替换为您希望接收代币的实际地址
  const recipientAddresses = [
    "0x9aB8199C95b77FED1ee62ecE22a23DfcC72fE81C",
    // 可以添加更多地址
  ];

  // 为每个地址铸造的代币数量 (单位是 MPS, 假设18位小数)
  const amountToMintPerAddressString = "500"; // 例如，为每个地址铸造 100 个 MPS
  const amountToMintPerAddress = ethers.parseUnits(amountToMintPerAddressString, 18); // 转换为 wei

  if (recipientAddresses.length === 0) {
    console.error("错误: 请至少提供一个接收者地址。");
    return;
  }
  if (recipientAddresses.some(addr => !ethers.isAddress(addr))) {
    console.error("错误: recipientAddresses 数组中包含无效的以太坊地址。");
    return;
  }


  // --- 获取已部署的MPS合约实例 ---
  const mpsProxyDeployment = getDeployment(networkName, "MPSProxy");
  if (!mpsProxyDeployment) {
    console.error(`错误: 在 deployments.json 中未找到 ${networkName} 网络的 MPSProxy 部署信息。`);
    process.exit(1);
  }
  const mpsProxyAddress = mpsProxyDeployment.address;
  console.log(`MPSProxy 合约地址: ${mpsProxyAddress}`);

  // 获取 MPS 合约工厂
  const MPS = await ethers.getContractFactory("MPS");
  // 将合约工厂连接到已部署的代理合约地址
  const mps = MPS.attach(mpsProxyAddress);
  console.log(`已连接到 MPS 合约 (通过代理: ${mpsProxyAddress})`);

  // --- 调用 mint 函数 ---
  console.log(`\n准备为以下地址列表铸造代币:`);
  recipientAddresses.forEach(addr => console.log(`  - ${addr}`));
  console.log(`每个地址将接收: ${amountToMintPerAddressString} MPS`);

  try {
    console.log("\n正在发送铸币交易...");
    // 使用 deployer (owner) 账户调用 mint 函数
    const tx = await mps.connect(deployer).mint(recipientAddresses, amountToMintPerAddress);
    
    console.log(`铸币交易已发送: ${tx.hash}`);
    console.log("等待交易确认...");
    
    const receipt = await tx.wait();
    
    if (receipt.status === 1) {
        console.log("✅ 铸币成功！");
        console.log(`  交易区块号: ${receipt.blockNumber}`);
    } else {
        console.error("❌ 铸币交易失败。请检查交易详情。");
        process.exit(1);
    }

  } catch (error) {
    console.error("\n❌ 调用 mint 函数时发生错误:");
    console.error(error);
    process.exit(1);
  }

  // --- (可选) 验证余额 ---
  console.log("\n--- 验证新余额 ---");
  for (const address of recipientAddresses) {
    try {
      const balance = await mps.balanceOf(address);
      console.log(`地址 ${address} 的新余额: ${ethers.formatUnits(balance, 18)} MPS`);
    } catch (error) {
      console.error(`查询地址 ${address} 余额失败:`, error);
    }
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("脚本执行失败:", error);
    process.exit(1);
  });