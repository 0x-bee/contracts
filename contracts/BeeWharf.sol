// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { ICurrency } from "./interfaces/ICheckout.sol";
import { BeeTokenChecker } from "./utils/BeeTokenChecker.sol";

contract BeeWharf is ICurrency, Ownable {

    using SafeERC20 for IERC20;
    using BeeTokenChecker for mapping (address => bool);

    // token address => support ?
    mapping (address => bool) private tokenSupported;
    // token address --> totalBalance ?
    mapping (address => uint256) private totalBalances;

    constructor () {
        // Ethereum USDT
        tokenSupported[0xdAC17F958D2ee523a2206206994597C13D831ec7] = true;
        // Ethereum USDC
        tokenSupported[0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48] = true;
    }

    // 支持新的 token 支付
    function addNewSupportToken(address tokenAddress) external onlyOwner {
        if (tokenSupported.containsKey(tokenAddress) == false) {
            tokenSupported[tokenAddress] = true;
        }
    }

    /* override ETH */

    function payEth(string calldata orderId) payable external override {}
    function withdrawEth(uint256 amount, string calldata billId) external override onlyOwner {}
    function refundETH(address to, uint256 amount, string calldata billId) external override onlyOwner {}
    function balanceOfETH() external view override returns (uint256) {}

    /* override ERC20 */

    function balanceOfERC20(address tokenAddress) external view override returns (uint256) {
        require(tokenSupported.containsKey(tokenAddress), "Unsurpported token!");
        return totalBalances[tokenAddress];
    }

    function payERC20(
        uint256 amount, 
        string calldata orderId, 
        address tokenAddress
        ) external override {
        this.payERC20From(msg.sender, amount, orderId, tokenAddress);
    }

    function payERC20From(
        address from, 
        uint256 amount, 
        string calldata orderId, 
        address tokenAddress
        ) external override {
        require(tokenSupported.containsKey(tokenAddress), "Unsurpported token!");

        IERC20 token = IERC20(tokenAddress);
        require(token.balanceOf(from) >= amount, "Insufficient balance funds");

        token.safeTransferFrom(from, address(this), amount);
        totalBalances[tokenAddress] += amount;
        emit PayERC20(from, amount, orderId, tokenAddress);
    }

    function withdrawERC20(
        uint256 amount, 
        string calldata billId,
        address tokenAddress
        ) external override onlyOwner {

        require(tokenSupported.containsKey(tokenAddress), "Not supported token");
        require(totalBalances[tokenAddress] >= amount, 'Insufficient funds');

        IERC20 token = IERC20(tokenAddress);
        token.safeTransfer(msg.sender, amount);
        totalBalances[tokenAddress] -= amount;
        emit WithdrawERC20(msg.sender, amount, billId, tokenAddress);
    }

    function refundERC20(
        address to, 
        uint256 amount, 
        string calldata billId, 
        address tokenAddress
        ) external override onlyOwner {
        this.refundERC20From(address(this), to, amount, billId, tokenAddress);
    }

    function refundERC20From(
        address from, 
        address to, 
        uint256 amount, 
        string calldata billId, 
        address tokenAddress
        ) external override onlyOwner {
        require(tokenSupported.containsKey(tokenAddress), "Unsurpported token!");

        IERC20 token = IERC20(tokenAddress);
        require(token.balanceOf(from) >= amount, "Insufficient balance funds");
        token.safeTransferFrom(from, to, amount);
        if (from == address(this)) {
            totalBalances[tokenAddress] -= amount;
        }
        emit RefundERC20(to, amount, billId, tokenAddress);
    }

}