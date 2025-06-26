const { ethers, network } = require("hardhat");
const { getDeployment } = require("./utils/deployment");

async function main() {
  const [deployer] = await ethers.getSigners();
  const networkName = network.name;
  console.log(`测试账户: ${deployer.address}`);

  // 动态获取合约地址
  const proxyDeployment = getDeployment(networkName, "MyNFTproxy");
  if (!proxyDeployment) {
    console.error(`错误: 在 deployments.json 中未找到 ${networkName} 网络的 MyNFTproxy 部署信息。`);
    process.exit(1);
  }
  const proxyAddress = proxyDeployment.address;
  console.log(`\n=== 测试NFT合约功能 (地址: ${proxyAddress}) ===`);
  
  // 获取NFT合约实例
  const NFT = await ethers.getContractFactory("MyNFTV2");
  const nft = NFT.attach(proxyAddress);
  
  try {
    // 1. 测试基本信息
    console.log("\n1. 基本信息:");
    const owner = await nft.owner();
    const symbol = await nft.symbol();
    const name = await nft.name();
    console.log(`   Owner: ${owner}`);
    console.log(`   Symbol: ${symbol}`);
    console.log(`   Name: ${name}`);
    
    // 2. 测试V2特有功能
    console.log("\n2. V2特有功能:");
    
    try {
      const baseURI = await nft.baseURI();
      console.log(`   BaseURI: ${baseURI}`);
    } catch (e) {
      console.log(`   BaseURI: 方法不存在或调用失败 - ${e.message}`);
    }
    
    try {
      const mintFee = await nft.mintFee();
      console.log(`   MintFee: ${ethers.formatEther(mintFee)} ETH`);
    } catch (e) {
      console.log(`   MintFee: 方法不存在或调用失败 - ${e.message}`);
    }
    
    try {
      const isWhitelisted = await nft.whitelist(deployer.address);
      console.log(`   当前账户在白名单: ${isWhitelisted}`);
    } catch (e) {
      console.log(`   Whitelist: 方法不存在或调用失败 - ${e.message}`);
    }
    
    // 3. 测试版税功能
    console.log("\n3. 版税功能:");
    try {
      const testSaleAmount = ethers.parseEther("1.0");
      const [receiver, royaltyAmount] = await nft.royaltyInfo(0, testSaleAmount);
      console.log(`   版税接收者: ${receiver}`);
      console.log(`   版税金额: ${ethers.formatEther(royaltyAmount)} ETH`);
    } catch (e) {
      console.log(`   版税功能: 方法不存在或调用失败 - ${e.message}`);
    }
    
    // 4. 测试铸造功能
    console.log("\n4. 铸造功能:");
    try {
      console.log("   尝试铸造测试NFT...");
      const mintTx = await nft.safeMint(deployer.address, "Test Token URI");
      await mintTx.wait();
      console.log("   ✅ 铸造成功!");
      
      // 检查token数量
      try {
        const balance = await nft.balanceOf(deployer.address);
        console.log(`   当前账户NFT数量: ${balance}`);
      } catch (e) {
        console.log(`   无法获取余额: ${e.message}`);
      }
    } catch (e) {
      console.log(`   铸造失败: ${e.message}`);
    }
    
    // 5. 测试白名单功能
    console.log("\n5. 白名单功能:");
    try {
      const isWhitelisted = await nft.whitelist(deployer.address);
      if (!isWhitelisted) {
        console.log("   添加当前账户到白名单...");
        const addTx = await nft.addToWhitelist([deployer.address]);
        await addTx.wait();
        console.log("   ✅ 已添加到白名单");
      } else {
        console.log("   ✅ 当前账户已在白名单中");
      }
    } catch (e) {
      console.log(`   白名单操作失败: ${e.message}`);
    }
    
    // 6. 测试付费铸造
    console.log("\n6. 付费铸造功能:");
    try {
      const mintFee = await nft.mintFee();
      console.log(`   尝试付费铸造 (费用: ${ethers.formatEther(mintFee)} ETH)...`);
      
      const mintWithFeeTx = await nft.mintWithFee(
        deployer.address, 
        "Paid Test Token URI",
        { value: mintFee }
      );
      await mintWithFeeTx.wait();
      console.log("   ✅ 付费铸造成功!");
    } catch (e) {
      console.log(`   付费铸造失败: ${e.message}`);
    }
    
    console.log("\n✅ NFT功能测试完成!");
    
  } catch (error) {
    console.error("❌ 测试失败:", error.message);
  }
}

main().catch((error) => {
  console.error("测试脚本执行失败:", error);
  process.exit(1);
}); 