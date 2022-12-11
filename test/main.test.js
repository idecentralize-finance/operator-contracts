const {
    loadFixture,
  } = require("@nomicfoundation/hardhat-network-helpers");
  
  //const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
  const { expect } = require("chai");
  
  const SHARES = "1020000000000000000000000" // this is equal to 51% of 2 000 000 shares with 18 decimals places
  const HIGH_PRICE = "1000000000000000000"
  const LOW_PRICE = "1000000000000000"
  const RIGHT_PRICE = "30000000000000000"
  const ONE_SHARE = "1000000000000000000"
  const ONE = 1
  const COLLECTIBLE_ONE_IN_ETH = ethers.utils.parseEther("1")
  const REBASE_URI = "https://ipfs.io/some-hash"
  const SHARE_MINT = "100"
  const SHARES_WITH_DECIMALS = ethers.utils.parseEther(SHARE_MINT)
  const COST_FOR_MINT = ethers.utils.parseEther("3")
  const NEG_COST_FOR_MINT = ethers.utils.parseEther("-3")
  
  describe("DEX DEPLOYMENT", function () {
    // We define a fixture to reuse the same setup in every test.
    // We use loadFixture to run this setup once, snapshot that state,
    // and reset Hardhat Network to that snapshot in every test.
    async function deployMaster() {
    
      // Contracts are deployed using the first signer/account by default
      const [owner, otherAccount] = await ethers.getSigners();
  
      // deploy DAO Token
      const DAOToken = await ethers.getContractFactory("DAOToken");
      const daoToken = await DAOToken.deploy();

      
      const DEX = await ethers.getContractFactory("DExchange");
      const dex = await DEX.deploy(daoToken.address);
  
      
      return { daoToken, dex };
    }
  
    describe("Deployment", function () {
      it("Should set the right owner", async function () {
        const { daoToken, dex } = await loadFixture(deployMaster);
        expect(await dex.owner()).to.equal(daoToken.address);
  
      });
  
    //   it("Should set the right master in CollectibleShares", async function () {
    //     const { master, collectibleShares} = await loadFixture(deployMaster);
    //     expect(await collectibleShares.QUPHMaster()).to.equal(master.address);
  
    //   });
  
    //   it("Should set the right owner in Master", async function () {
    //     const { master, owner } = await loadFixture(deployMaster);
  
    //     expect(await master.owner()).to.equal(owner.address);
    //   });
  
    //   it("Should initiate shares and mint 51% to owner", async function () {
    //     const { master, shares, owner } = await loadFixture(
    //       deployMaster
    //     );
  
    //     await master.initiateShares({from:owner.address})
  
    //     expect(await shares.balanceOf(owner.address)).to.equal(
    //       SHARES
    //     );
    //   });
  
    //   it("Should revert if the intitiator is called a second time", async function () {
    //     // We don't use the fixture here because we want a different deployment
    //     const { master, shares, owner } = await loadFixture(
    //       deployMaster
    //     );
    //     await master.initiateShares({from:owner.address})
    //     expect(await shares.balanceOf(owner.address)).to.equal(
    //       SHARES
    //     );
    //     await expect(master.initiateShares({from:owner.address})).to.be.revertedWith(
    //       "MASTER: Shares already initiated"
    //     );
    //   });
    });
  
  
  
//      describe("Minting", function () {
  
//        describe("Shares", function () {
//         it("Should revert with the right error if price is to high", async function () {
//           const { master, otherAccount } = await loadFixture(deployMaster);
  
//           await expect(master.connect(otherAccount).mintShares(ONE, otherAccount.address, { value: HIGH_PRICE})).to.be.revertedWith(
//             "MASTER: Price is not right"
//           );
//         });
//         it("Should revert with the right error if price is to low", async function () {
//           const { master, otherAccount } = await loadFixture(deployMaster);
  
//           await expect(master.connect(otherAccount).mintShares(ONE, otherAccount.address,{ value: LOW_PRICE})).to.be.revertedWith(
//             "MASTER: Price is not right"
//           );
//         });
//         it("Should not revert if price is right", async function () {
//           const { master, otherAccount } = await loadFixture(deployMaster);
  
//           await expect(master.connect(otherAccount).mintShares(ONE, otherAccount.address,{ value: RIGHT_PRICE})).not.to.be.revertedWith(
//             "MASTER: Price is not right"
//           );
//         });
//         it("Should mint the correct amount to the caller", async function () {
//           const { master, shares, otherAccount } = await loadFixture(deployMaster);
//           await master.connect(otherAccount).mintShares(ONE, otherAccount.address, { value: RIGHT_PRICE})
//           expect(await shares.balanceOf(otherAccount.address)).to.equal(
//             ONE_SHARE
//           );
//         });
  
//         it("Should not mint directly from the Shares contract", async function () {
//           const { shares, otherAccount } = await loadFixture(deployMaster);
//           await expect(shares.connect(otherAccount).mint(otherAccount.address, ONE)).to.be.revertedWith(
//             "SHARES: You are not the master"
//           );
//         });
//       });
  
//         describe("Colletible Shares", function () {
//         it("Should revert with the right error if no item have been created", async function () {
//           const { master, otherAccount } = await loadFixture(
//             deployMaster
//           );
  
//           // We use lock.connect() to send a transaction from another account
//           await expect(master.connect(otherAccount).mintCollectibleShares(ONE, otherAccount.address)).to.be.revertedWith(
//             "MASTER: Creation is not ready to mint"
//           );
//         });
  
//         it("Should release a new item", async function () {
//           const { master, owner } = await loadFixture(
//             deployMaster
//           );
//           await expect(master.connect(owner).releaseCreation(ONE, COLLECTIBLE_ONE_IN_ETH, REBASE_URI )).not.to.be.revertedWith(
//             "MASTER: Creation is not ready to mint"
//           );
//         });
  
//         it("Should revert when minting colletible at wrong price", async function () {
//           const { master, owner, otherAccount } = await loadFixture(
//             deployMaster
//           );
//           await expect(master.connect(owner).releaseCreation(ONE, COLLECTIBLE_ONE_IN_ETH, REBASE_URI )).not.to.be.revertedWith(
//             "MASTER: Creation is not ready to mint"
//           );
//           await expect(master.connect(otherAccount).mintCollectibleShares(ONE, otherAccount.address, {value:LOW_PRICE})).to.be.revertedWith(
//             "MASTER: Price is not right"
//           );
//         });
  
//         it("Should mint colletible when the price is right", async function () {
//           const { master, collectibleShares, owner, otherAccount } = await loadFixture(
//             deployMaster
//           );
//           await expect(master.connect(owner).releaseCreation(ONE, COLLECTIBLE_ONE_IN_ETH, REBASE_URI )).not.to.be.revertedWith(
//             "MASTER: Creation is not ready to mint"
//           );
//           await master.connect(otherAccount).mintCollectibleShares(ONE, otherAccount.address, {value:COLLECTIBLE_ONE_IN_ETH})
//           expect(await collectibleShares.ownerOf(ONE)).to.equal(otherAccount.address)
//         });
//         it("Should revert when trying to mint existing collectibles", async function () {
//           const { master, owner, otherAccount } = await loadFixture(
//             deployMaster
//           );
  
//           await expect(master.connect(owner).releaseCreation(ONE, COLLECTIBLE_ONE_IN_ETH, REBASE_URI )).not.to.be.revertedWith(
//             "MASTER: Creation is not ready to mint"
//           );
//           await master.connect(otherAccount).mintCollectibleShares(ONE, otherAccount.address, {value:COLLECTIBLE_ONE_IN_ETH})
  
//           await expect(master.connect(otherAccount).mintCollectibleShares(ONE, otherAccount.address, {value:COLLECTIBLE_ONE_IN_ETH})
//           ).to.be.revertedWith(
//             "ERC721: token already minted"
//           )
//         });
//        });
  
//        describe("Events", function () {
//         describe("Shares", function () {
//         it("Should emit an event on minting shares", async function () {
//           const { master, owner } = await loadFixture(
//             deployMaster
//           );
  
//           await expect(master.mintShares(SHARE_MINT, owner.address, {value: COST_FOR_MINT}))
//             .to.emit(master, "SharesMinted")
//             .withArgs(SHARES_WITH_DECIMALS, COST_FOR_MINT, owner.address); // We accept any value as `when` arg
//         });
//       });
  
  
//       describe("CollectibleShares", function () {
//         it("Should emit an event on minting CollectibleShares", async function () {
//           const { master, owner } = await loadFixture(
//             deployMaster
//           );
  
//           await expect(master.connect(owner).releaseCreation(ONE, COLLECTIBLE_ONE_IN_ETH, REBASE_URI )).not.to.be.revertedWith(
//             "MASTER: Creation is not ready to mint"
//           );
  
//           await expect(master. mintCollectibleShares(ONE, owner.address, {value: COLLECTIBLE_ONE_IN_ETH}))
//             .to.emit(master, "CollectibleShareMinted")
//             .withArgs(ONE, COLLECTIBLE_ONE_IN_ETH, owner.address); // We accept any value as `when` arg
//         });
      
        
  
//       });
  
//     });
  
//       describe("Transfers", function () {
//         it("Should transfer the funds to master and deduct from account", async function () {
//           const { master, owner } = await loadFixture(
//             deployMaster
//           );
  
//         await expect( await master.connect(owner).mintShares(SHARE_MINT, owner.address, {value: COST_FOR_MINT})).to.changeEtherBalances(
//             [master, owner],
//             [COST_FOR_MINT, NEG_COST_FOR_MINT]
//           );
         
//         });
//       });
//      });
  
    });
  