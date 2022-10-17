require("dotenv").config();
const HDWalletProvider = require("truffle-hdwallet-provider");


const mnemonic = process.env.MNEMONIC;

module.exports = {
  plugins: ["truffle-security", "solidity-coverage"],
  networks: {
    // Quick dev using local ganache @port 7545
    development: {
      host: "127.0.0.1",
      port: 8545,
      network_id: "*",
      gas: 6700000,
    },
    hardhat: {
        host: "127.0.0.1",     
        port: 8545,            
        network_id: "31337",      
       },
    goerli: {
      provider: () =>
        new HDWalletProvider(
          [process.env.ALCHEMY_GOERLI_API_KEY],
          "https://eth-goerli.g.alchemy.com/v2/"+process.env.ALCHEMY_GOERLI_API_KEY
        ),
      network_id: "*",
      gas: 5242880,
    },
  },

  mocha: {
    useColors: true,
    reporter: "eth-gas-reporter",
    reporterOptions: { currency: "USD" },
  },
  // Configuring compiler
  compilers: {
    solc: {
      version: "0.8.17",
      docker: false,
      settings: {
        optimizer: {
          enabled: true,
          runs: 200,
        },
        evmVersion: "london"
      },
    },
  },
};