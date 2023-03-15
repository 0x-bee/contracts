// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { ICurrency } from "./interfaces/ICheckout.sol";
import { BeeTokenChecker } from "./utils/BeeTokenChecker.sol";
import { BeePrivilegeChecker } from "./utils/BeePrivilegeChecker.sol";

contract BeeWharf is ICurrency {

    using SafeERC20 for IERC20;
    using BeePrivilegeChecker for mapping (address => bool);
    using BeeTokenChecker for bytes32[];

    address _administrator;
    mapping(address => bool) private capitalist;
    
    bytes32[] public whitelistedSymbols;
    mapping(bytes32 => address) private _whitelistedTokens;
    mapping (bytes32 => uint256) private _totalBalances;

    constructor () {
        _administrator = msg.sender;
    }

    function updateCapitalist(address ads, bool grant) external {
        require(msg.sender == _administrator, "This function is not public");
        capitalist[ads] = grant;
    }

    function whitelistToken(bytes32 symbol, address tokenAddress) external {
        require(msg.sender == _administrator, "This function is not public");
        whitelistedSymbols.push(symbol);
        _whitelistedTokens[symbol] = tokenAddress;
    }

    function getWhitelistedSymbols() external view returns (bytes32[] memory){
        return whitelistedSymbols;
    }

    function getWhitelistedTokenAddress(bytes32 symbol) public view returns(address) {
        return _whitelistedTokens[symbol];
    }

    receive() external payable {
        _totalBalances['Eth'] += msg.value;
    }
    /* override ETH */

    function payEth(string calldata orderId) payable external override {}
    function withdrawEth(uint256 amount, string calldata billId) external override {}
    function refundETH(address to, uint256 amount, string calldata billId) external override {}
    function balanceOfETH() external view override returns (uint256) {}

    /* override ERC20 */

    function balanceOfERC20(bytes32 symbol) external view override returns (uint256) {
        return _totalBalances[symbol];
    }

    // 给当前合约充值某种代币
    function topupToken(address from, uint256 amount, bytes32 symbol) external {
        require(whitelistedSymbols.containsValue(symbol), "Unsurpported token!");
        address tokenAddress = _whitelistedTokens[symbol];
        IERC20 token = IERC20(tokenAddress);
        require(token.balanceOf(from) >= amount, "Insufficient balance funds");

        // 用户侧需要调用 approve 先授权 (在 Dapp 中完成)
        token.safeTransferFrom(from, address(this), amount);
        _totalBalances[symbol] += amount;
    }

    function payERC20(
        uint256 amount, 
        string calldata orderId, 
        bytes32 symbol
        ) external override {
        this.topupToken(msg.sender, amount, symbol);
        address tokenAddress = _whitelistedTokens[symbol];
        emit PayERC20(msg.sender, amount, orderId, tokenAddress);
    }

    function withdrawERC20(
        uint256 amount, 
        string calldata billId,
        bytes32 symbol
        ) external override {

        require(capitalist.containsKey(msg.sender), 'You are not authorized to do that!');
        require(_totalBalances[symbol] >= amount, 'Insufficient funds');

        address tokenAddress = _whitelistedTokens[symbol];
        IERC20 token = IERC20(tokenAddress);
        token.safeTransfer(msg.sender, amount);
        _totalBalances[symbol] -= amount;
        emit WithdrawERC20(msg.sender, amount, billId, tokenAddress);
    }

    // 从当前合约退款
    function refundERC20ByContract(
        address to, 
        uint256 amount, 
        string calldata billId, 
        bytes32 symbol
        ) external override {
        require(capitalist.containsKey(msg.sender), 'You are not authorized to do that!');
        require(_totalBalances[symbol] >= amount, 'Insufficient funds');

        address tokenAddress = _whitelistedTokens[symbol];
        IERC20 token = IERC20(tokenAddress);
        token.safeTransfer(to, amount);
        emit RefundERC20(to, amount, billId, tokenAddress);
    }

    // 从调用者那里退款，途径合约，需要外部先授权
    function refundERC20 (
        address to, 
        uint256 amount, 
        string calldata billId, 
        bytes32 symbol
        ) external override {
        require(capitalist.containsKey(msg.sender), 'You are not authorized to do that!');
        this.topupToken(msg.sender, amount, symbol);
        this.refundERC20ByContract(to, amount, billId, symbol);
    }
}