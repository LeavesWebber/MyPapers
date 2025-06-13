// scripts/advanceTime.js
async function main() {
  const seconds = parseInt(process.env.SECONDS) || 170000;
  console.log(`推进区块链时间 ${seconds} 秒（约 ${Math.floor(seconds/86400)} 天）...`);
  
  // 连接到本地网络
  const provider = ethers.provider;
  
  try {
    // 推进时间
    await provider.send("evm_increaseTime", [seconds]);
    await provider.send("evm_mine");
    
    const latestBlock = await provider.getBlock("latest");
    console.log(`✅ 成功推进时间`);
    console.log(`新的区块时间: ${new Date(latestBlock.timestamp * 1000).toLocaleString()}`);
  } catch (error) {
    console.error("推进时间失败:", error.message);
  }
}

main().catch(console.error);