/*SPDX-License-Identifier: MIT

**/


import "../node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol";

pragma solidity 0.8.17;

contract DAOToken is ERC20 {

    constructor() ERC20("DAOToken", "DAOT"){

    }

}