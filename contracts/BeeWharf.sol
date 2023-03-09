// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

interface WharfInterface {
    
    event PayEth(address indexed from, uint256 amount, string orderId);
    event PayERC20(address indexed from, uint256 amount, string orderId, address indexed tokenAddress);
    event WithdrawEth(address indexed to, uint256 amount, string billId);
    event WithdrawERC20(address indexed to, uint256 amount, string billId, address indexed tokenAddress);
    event RefundERC20(address indexed to, uint256 amount, string billId, address indexed tokenAddress);
    
    // orderId 中心化系统订单ID
    function payEth(string calldata orderId) payable external;

    // from 支出钱包地址
    // amount 支出金额
    // orderId 中心化系统订单ID 发出事件需要
    // tokenAdress erc20 token 合约地址
    function payERC20(uint256 amount, string calldata orderId, address tokenAddress) external;
    function payERC20From(address from, uint256 amount, string calldata orderId, address tokenAddress) external;

    // amount 提现金额
    // billId 中心化系统结算ID
    function withdrawEth(uint256 amount, string calldata billId) external;

    // amount 提现金额
    // currency erc20 合约地址
    // billId 中心化系统结算ID
    function withdrawERC20(uint256 amount, string calldata billId, address tokenAddress) external;

    // frome : 发起退款的账户
    // to : 退款地址
    // amount : 退款金额
    // billId : 账单 Id
    // tokenAdress erc20 token 合约地址
    function refundERC20(address to, uint256 amount, string calldata billId, address tokenAddress) external;
    function refundERC20From(address from, address to, uint256 amount, string calldata billId, address tokenAddress) external;
}

library BeeCheck {
    function containsKey(mapping(address => bool) storage aMap, address aKey) internal view returns (bool) {
        return aMap[aKey];
    }
}

contract BeeWharf is WharfInterface {

    using BeeCheck for mapping (address => bool);
    // token address => support ?
    mapping (address => bool) private tokenSupported;
    // token address --> totalBalance ?
    mapping (address => uint256) private totalBalances;

    address private owner;
    constructor () {
        owner = msg.sender;
        // Ethereum USDT
        tokenSupported[0xdAC17F958D2ee523a2206206994597C13D831ec7] = true;
        // Ethereum USDC
        tokenSupported[0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d] = true;
    }

    // 支持新的 token 支付
    function addNewSupportToken(address tokenAddress, bytes32 symbol) external {
        require(msg.sender == owner, 'This function is not public');
        if (tokenSupported.containsKey(tokenAddress) == false) {
            tokenSupported[tokenAddress] = true;
        }
    }

    function payEth(string calldata orderId) payable external {

    }

    // from 支出钱包地址
    // amount 支出金额
    // orderId 中心化系统订单ID 发出事件需要
    // currency erc20 合约地址
    function payERC20(
        uint256 amount, 
        string calldata orderId, 
        address tokenAddress
        ) external {
        this.payERC20From(msg.sender, amount, orderId, tokenAddress);
    }

    function payERC20From(
        address from, 
        uint256 amount, 
        string calldata orderId, 
        address tokenAddress
        ) external {
        
        if (tokenSupported.containsKey(tokenAddress) == false) {
            // unsupported token
        } else {
            // USDT 特殊处理
            if (tokenAddress == 0xdAC17F958D2ee523a2206206994597C13D831ec7) { // USDT

            } else {
                bool success = ERC20(tokenAddress).transferFrom(from, address(this), amount);
                if (success) {
                    totalBalances[tokenAddress] += amount;
                    emit PayERC20(from, amount, orderId, tokenAddress);
                } else {
                    // Error: 
                }
            }
        }
    }

    function withdrawEth(uint256 amount, string calldata billId) external {
        require(msg.sender == owner, 'This function is not public');
    }

    function withdrawERC20(
        uint256 amount, 
        string calldata billId,
        address tokenAddress
        ) external {
        require(msg.sender == owner, 'This function is not public');
        require(tokenSupported.containsKey(tokenAddress), "Not supported token");
        require(totalBalances[tokenAddress] >= amount, 'Insufficient funds');
        totalBalances[tokenAddress] -= amount;
        ERC20(tokenAddress).transfer(msg.sender, amount);
    }

    function refundERC20(
        address to, 
        uint256 amount, 
        string calldata billId, 
        address tokenAddress
        ) external {
        require(msg.sender == owner, 'This function is not public');
        this.refundERC20From(owner, to, amount, billId, tokenAddress);
    }

    function refundERC20From(
        address from, 
        address to, 
        uint256 amount, 
        string calldata billId, 
        address tokenAddress
        ) external {

        require(msg.sender == owner, 'This function is not public');
        if (tokenAddress == 0xdAC17F958D2ee523a2206206994597C13D831ec7) { // USDT

        } else {
            bool success = ERC20(tokenAddress).transferFrom(from, to, amount);
            if (success) {
                if (from == address(this)) {
                    totalBalances[tokenAddress] -= amount;
                }
                emit RefundERC20(to, amount, billId, tokenAddress);
            } else {
                // Error: 
            }
        }
    }

}