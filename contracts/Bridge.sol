/*SPDX-License-Identifier: MIT

author: ian.dorion.eth

This is the Operator Contract:

**/


import "../node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "../node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

pragma solidity 0.8.17;

contract Bridge is Ownable {

    event BridgeDeposit(
    address sender,
    uint256 destinationChainId,
    uint256 nonce,
    bytes32 resourceId,
    bytes data
    );

    constructor(){}

    function bridgeDeposit(uint256 chainId, bytes32 ressourceId, bytes memory data) public payable{


    }

}