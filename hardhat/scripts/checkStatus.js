const { ethers } = require("hardhat");

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log(`检查账户: ${deployer.address}`);

  // 从deployments.json获取地址
  const proxyAddress = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512";
  const timelockAddress = "0xCafac3dD18aC6c6e92c921884f9E4176737C052c";

  console.log("\n=== 检查代理合约状态 ===");
  
  // 检查代理合约
  const proxyABI = [
    "function timelock() view returns (address)",
    "function implementation() view returns (address)",
    "function owner() view returns (address)"
  ];
  
  const proxy = new ethers.Contract(proxyAddress, proxyABI, deployer);
  
  try {
    const currentImpl = await proxy.implementation();
    console.log(`当前实现地址: ${currentImpl}`);
    
    const timelock = await proxy.timelock();
    console.log(`时间锁地址: ${timelock}`);
    
    const owner = await proxy.owner();
    console.log(`代理合约owner: ${owner}`);
  } catch (e) {
    console.error("读取代理合约失败:", e.message);
  }

  console.log("\n=== 检查时间锁合约状态 ===");
  
  // 检查时间锁合约
  const TimelockController = await ethers.getContractFactory("TimelockController");
  const timelock = TimelockController.attach(timelockAddress);
  
  try {
    const minDelay = await timelock.getMinDelay();
    console.log(`最小延迟: ${minDelay} 秒`);
    
    const isAdmin = await timelock.hasRole(await timelock.PROPOSER_ROLE(), deployer.address);
    console.log(`当前账户是否为管理员: ${isAdmin}`);
    
    const isExecutor = await timelock.hasRole(await timelock.EXECUTOR_ROLE(), deployer.address);
    console.log(`当前账户是否为执行者: ${isExecutor}`);
  } catch (e) {
    console.error("读取时间锁合约失败:", e.message);
  }

  console.log("\n=== 检查NFT合约功能 ===");
  
  // 检查NFT合约功能
  const NFTABI = [
    "function owner() view returns (address)",
    "function symbol() view returns (string)",
    "function name() view returns (string)",
    "function baseURI() view returns (string)",
    "function royaltyInfo(uint256,uint256) view returns (address,uint256)",
    "function mintFee() view returns (uint256)",
    "function whitelist(address) view returns (bool)"
  ];
  
  const nft = new ethers.Contract(proxyAddress, NFTABI, deployer);
  
  try {
    const owner = await nft.owner();
    console.log(`NFT合约owner: ${owner}`);
    
    const symbol = await nft.symbol();
    console.log(`NFT符号: ${symbol}`);
    
    const name = await nft.name();
    console.log(`NFT名称: ${name}`);
    
    try {
      const baseURI = await nft.baseURI();
      console.log(`基础URI: ${baseURI}`);
    } catch (e) {
      console.log("baseURI方法不存在或调用失败");
    }
    
    try {
      const mintFee = await nft.mintFee();
      console.log(`铸造费用: ${ethers.formatEther(mintFee)} ETH`);
    } catch (e) {
      console.log("mintFee方法不存在或调用失败");
    }
    
    try {
      const isWhitelisted = await nft.whitelist(deployer.address);
      console.log(`当前账户是否在白名单: ${isWhitelisted}`);
    } catch (e) {
      console.log("whitelist方法不存在或调用失败");
    }
  } catch (e) {
    console.error("读取NFT合约失败:", e.message);
  }
}

main().catch((error) => {
  console.error("检查失败:", error);
  process.exit(1);
}); 