// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ICurrency {
    
    mapping(address => uint256) balanceOf;

    event PayEth(address indexed from, uint256 amount,string orderId);
    event PayERC20(address indexed from, uint256 amount,string orderId, address indexed curency);
    event WithdrawEth(address indexed to, uint256 amount,string billId);
    event WithdrawERC20(address indexed to, uint256 amount, address indexed curency,string billId);
    
    function payEth(string calldata orderId) payable external virtual;
    function payERC20(address from, uint256 amount,string calldata orderId,address curency) payable external virtual;

    function withdrawEth(uint256 amount,string calldata billId) external virtual;
    function withdrawERC20(uint256 amount,address curency,string calldata billId) external virtual;

    function owner() public view virtual returns (address);
}
