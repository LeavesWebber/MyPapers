require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");
require('dotenv').config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.24",
  admin:{
    privateKey: "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", // Hardhat 默认第一个账户
    neverRevoke: true
  },
  networks: {
    paperschain: {
      url: "https://rpc.paperschain.io",
      chainId: 408,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY]
    }
  }
};
