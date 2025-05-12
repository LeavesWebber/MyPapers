require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");
require('dotenv').config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.24",
  networks: {
    paperschain: {
      url: "https://rpc.paperschain.io",
      chainId: 408,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY]
    }
  }
};
