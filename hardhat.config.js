require('dotenv').config()
require("@nomiclabs/hardhat-ethers")
require("@nomiclabs/hardhat-waffle");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  // defaultNetwork: "hardhat",
  // networks: {
  //   goerli: {
  //       url: "https://eth-goerli.g.alchemy.com/v2/"+process.env.ALCHEMY_GOERLI_API_KEY,
  //       chainId: 5,
  //       gasPrice: 5000000000,
  //       gas:2000000,
  //       gasLimit:3000000,
  //       accounts: [process.env.MNEMONIC]
    
  //   },
    
  //},
  paths: {
    sources: "./contracts",
    tests: "./test",
    artifacts: "./artifacts",
  
},

  solidity: "0.8.17",
};
