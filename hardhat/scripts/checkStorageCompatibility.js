const { ethers, upgrades } = require("hardhat");

async function main() {
  // 根据您的 deployments.json 和 MPERdeploy.js，旧合约应为 MPER
  const oldContractName = "MPER"; 
  // 新合约，即您想要检查的 MPS.sol
  const newContractName = "MPS";  

  console.log(`正在验证从 ${oldContractName} 到 ${newContractName} 的升级兼容性...`);

  try {
    const OldContractFactory = await ethers.getContractFactory(oldContractName);
    const NewContractFactory = await ethers.getContractFactory(newContractName);

    // validateUpgrade 会检查存储布局、初始化函数等。
    // MPERproxy 是一个自定义的 ERC1967 代理，其升级由管理员（通过 Timelock）控制，
    // 行为上更接近透明代理，因此我们使用 kind: 'transparent'。
    // 如果 MPS.sol 引入了与 MPER.sol 不兼容的存储更改，此函数会抛出错误。
    await upgrades.validateUpgrade(OldContractFactory, NewContractFactory, { kind: 'transparent' });

    console.log(`✅ ${newContractName} 与 ${oldContractName} 的存储布局兼容，可以安全升级。`);
  } catch (error) {
    console.error(`❌ 验证失败: ${newContractName} 与 ${oldContractName} 不兼容。`);
    console.error("详细错误信息:");
    // OpenZeppelin Upgrades 插件的错误通常包含一个 errors 数组，其中有更具体的冲突信息
    if (error.errors && error.errors.length > 0) {
        error.errors.forEach(errDetail => {
            console.error(`- ${errDetail.message || errDetail}`);
            if (errDetail.src) {
                console.error(`  เกี่ยวข้องกับ: ${errDetail.src}`);
            }
        });
    } else {
        console.error(error.message);
    }
    // 你可以取消下面这行的注释来查看完整的错误对象
    // console.error(error); 
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("脚本执行时发生未捕获错误:", error);
    process.exit(1);
  });