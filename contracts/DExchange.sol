/*SPDX-License-Identifier: MIT
.___    .___                          __                .__  .__                          __  .__     
|   | __| _/____   ____  ____   _____/  |_____________  |  | |__|_______ ____       _____/  |_|  |__  
|   |/ __ |/ __ \_/ ___\/ __ \ /    \   __\_  __ \__  \ |  | |  \___   // __ \    _/ __ \   __\  |  \ 
|   / /_/ \  ___/\  \__\  ___/|   |  \  |  |  | \// __ \|  |_|  |/    /\  ___/    \  ___/|  | |   Y  \
|___\____ |\___  >\___  >___  >___|  /__|  |__|  (____  /____/__/_____ \\___  > /\ \___  >__| |___|  /
         \/    \/     \/    \/     \/                 \/              \/    \/  \/     \/          \/ 
Author: ian.dorion.eth

This is the D Exchange Contract:

This contract is owned by the IDFI DAO Tokens


**/
import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "../node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "../node_modules/@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "../node_modules/@openzeppelin/contracts/security/Pausable.sol";
import "../node_modules/@openzeppelin/contracts/security/ReentrancyGuard.sol";

pragma solidity 0.8.17;

contract DExchange is Ownable, Pausable, ReentrancyGuard {

    using SafeERC20 for IERC20;

     struct Orders {
        address maker;
        address[] taker;
        address makerAsset;
        address takerAsset;
        uint256 makerAmount;
        uint256 takerAmount;
        uint256 filledAmount;
        uint256 orderId;
        bool active;
    }

    Orders[] public orders;
  

    address ETHER;
    uint256 MENTISSA = 1e18;
    uint public _fees;

    event OrderCreated(
        address makerAsset,
        uint256 makerAmount,
        address takerAsset,
        uint256 takerAmount,
        address maker
    );

    event OrderFilled(
        address makerAsset,
        uint256 makerAmount,
        address takerAsset,
        uint256 takerAmount,
        address taker,
        bool active
    );

    event OrderCancelled(
        uint256 orderId,
        address maker
    );

    /**
     @notice contract is owned by the dao token
    **/
    constructor(
        address daoToken
    ) {
        transferOwnership(daoToken);
    }

     /**
     @notice create order
     @param makerAsset the assets you have
     @param takerAsset the assets you want
     @param makerAmount the amounts you have
     @param takerAmount the amounts you want
    */

    function createOrder(
        address[] memory makerAsset,
        address[] memory takerAsset,
        uint256[] memory makerAmount,
        uint256[] memory takerAmount
    ) public payable{
        validate(makerAsset, takerAsset, makerAmount, takerAmount);
        uint256 orderId;
        address maker = msg.sender;
        uint256 makerEtherSpent;
        for(uint i=0;i>makerAsset.length;i++){
            orderId = orders.length;
            Orders storage order = orders[orderId];

            order.maker = maker;
            order.makerAsset = makerAsset[i];
            order.makerAmount = makerAmount[i];
            order.takerAsset = takerAsset[i];
            order.takerAmount = takerAmount[i];
            order.active = true;

            if(makerAsset[i] == ETHER){
                makerEtherSpent += makerAmount[i];
            } else{
                SafeERC20.safeTransferFrom(IERC20(makerAsset[i]), maker, address(this), makerAmount[i]);
            }

            emit OrderCreated(makerAsset[i], makerAmount[i], takerAsset[i], takerAmount[i], maker);
        }

        require(msg.value ==  makerEtherSpent, "createOrder: invalid ether amount");

    }

    /**
     @notice fill order
     @param orderIds the orders you want to fill
     @param fillAmounts the amounts your filling
    */
    function fillOrders(
        uint256[] memory orderIds, 
        uint256[] memory fillAmounts
        ) public payable nonReentrant{

        address taker = msg.sender;
        uint256 takerEtherSpent = 0;

        for(uint i=0; i<orderIds.length; i++){
  
            Orders storage order = orders[orderIds[i]];
            uint256 _orderPrice = orderPrice(orderIds[i]);
            uint256 cost = _orderPrice * fillAmounts[i];
            order.taker.push(taker);
            uint256 remaining = order.makerAmount - order.filledAmount;
            uint256 filling = remaining < fillAmounts[i] ? remaining : fillAmounts[i];

            if((remaining) > fillAmounts[i]){
                order.filledAmount += filling;
            }else if(remaining == filling){
                order.active = false;
            }
            
            if(order.takerAsset == ETHER){
                takerEtherSpent += cost;
            }else{
                _transferAsset(order.makerAsset, taker, order.maker, fillAmounts[i]);
            }

            _transferAsset(order.makerAsset, address(this), taker, cost);
            

                // transfer a dao token to the taker up to a max of X per 24 hrtaker
            emit OrderFilled(order.makerAsset, order.makerAmount, order.takerAsset, order.takerAmount, taker, order.active);
        }

        // check that the taker fully paid all ether orders
        require(msg.value == takerEtherSpent, "fillOrder: Not enough ether to fill orders");

    }

    /**
     @notice cancel order
     @param orderId the orders you want to cancel
    */
    function cancelOrder(uint256 orderId) public {

          Orders storage order = orders[orderId];
          require(order.active, "cancelOrder: order already closed");
          require(order.maker == msg.sender, "cancelOrder: not the order maker");
          order.active = false;

          uint256 remaining = order.makerAmount - order.filledAmount;

          SafeERC20.safeTransferFrom(IERC20(order.makerAsset), address(this), order.maker, remaining);

          emit OrderCancelled(orderId, order.maker);

    }

    /**
     @notice create order
     @dev address(0) is ETHER
     @param makerAsset the assets you have
     @param takerAsset the assets you want
     @param makerAmount the amounts you have
     @param takerAmount the amounts you want
    */
    function validate(        
        address[] memory makerAsset,
        address[] memory takerAsset,
        uint256[] memory makerAmount,
        uint256[] memory takerAmount
        ) internal pure {
            uint256 expectedLength = makerAsset.length;

            for(uint i=0; i < expectedLength; i++){
                require(makerAmount[i] > 0 , "validate: maker amount can't be 0");
                require(takerAmount[i] > 0 , "validate: taker amount can't be 0");
            }

            require(
                 takerAsset.length == expectedLength &&
                 takerAmount.length == expectedLength &&
                 makerAmount.length == expectedLength,
                 "validate: order miss match"
                );
       
    }

     /**
     @notice set fees
     @dev fee in base 3 where 100 = 1.00%
    */
    function setFees(uint256 fees) public onlyOwner{
        _fees = fees;
    }

    /**
     @notice transfer asset
     @dev fee in base 3 where 100 = 1%
    */
    function _transferAsset(
        address asset, 
        address from, 
        address to, 
        uint256 amount
        ) internal{
        uint256 chargedFees = amount / 10000 * _fees;
        uint256 netAmount = amount - chargedFees;

         if(asset == ETHER){
                 (bool sent, ) = to.call{value: netAmount}("");
                 require(sent, "fillOrder: unable to pay net amount");
                 (sent, ) = owner().call{value: chargedFees}("");
                 require(sent, "fillOrder: unable to pay fess");
            }else{
                SafeERC20.safeTransferFrom(IERC20(asset), from, to, netAmount);
                SafeERC20.safeTransferFrom(IERC20(asset), from, owner(), chargedFees);
            }
    }


    /**
     @notice order price
     @dev return the price in taker asset for maker asset
    */
    function orderPrice(uint256 orderId)
        public
        view
        returns (uint256 price)
    {
        Orders memory order = orders[orderId];

        if(order.takerAmount >= order.makerAmount){
            price = order.takerAmount / order.makerAmount;
        }else{
            price = order.takerAmount * MENTISSA / order.makerAmount / MENTISSA;
        }
     
        return price;
    }



    /**
     @notice security features
     @dev this contract is pausable
    */

    function pause() public onlyOwner {
        require(!paused(), "contract must not be paused");
        _pause();
    }
   
    function unpause() public onlyOwner {
        require(paused(), "contract must be paused");
        _unpause();
    }

    receive() external payable{
        // this contract can receive ether
    }


}