// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ICurrencyReceiver {
    event Pay(
        address currency,
        address indexed from,
        uint256 amount,
        string orderId
    );

    event Withdraw(
        address[] indexed currency,
        uint256[] amount,
        address indexed to,
        string billId
    );

    event Refund(
        address[] indexed currency,
        address[] indexed to,
        uint256[] amount,
        string[] orderIds
    );

    function pay(
        address currency,
        uint256 amount,
        string calldata orderId
    ) external virtual;

    function withdraw(
        address[] calldata currency,
        uint256[] calldata amount,
        string calldata billId
    ) external virtual;

    function refund(
        address[] calldata currency,
        uint256[] calldata amount,
        uint256[] calldata to,
        string[] calldata orderId
    ) external virtual;

    function balanceOf(
        address currency
    ) external view virtual returns (uint256);
}