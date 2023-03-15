// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ICurrency {

    // ETH Event
    event PayEth(address indexed from, uint256 amount, string orderId);
    event WithdrawEth(address indexed to, uint256 amount,string billId);
    event RefundEth(address indexed to, uint256 amount,string billId);

    // ERC20 Event
    event PayERC20(address indexed from, uint256 amount, string orderId, address indexed tokenAddress);
    event WithdrawERC20(address indexed to, uint256 amount, string billId, address indexed tokenAddress);
    event RefundERC20(address indexed to, uint256 amount, string billId, address indexed tokenAddress);
    
    // ETH function
    function payEth(string calldata orderId) payable external virtual;
    function withdrawEth(uint256 amount, string calldata billId) external virtual;
    function refundETH(address to, uint256 amount, string calldata billId) external virtual;
    function balanceOfETH() external view virtual returns (uint256);

    // ERC20 function
    function payERC20(uint256 amount, string calldata orderId, bytes32 symbol) external virtual;
    function withdrawERC20(uint256 amount, string calldata billId, bytes32 symbol) external virtual;
    function refundERC20(address to, uint256 amount, string calldata billId, bytes32 symbol) external virtual;
    function refundERC20ByContract(address to, uint256 amount, string calldata billId, bytes32 symbol) external virtual;
    function balanceOfERC20(bytes32 symbol) external view virtual returns (uint256);
}
