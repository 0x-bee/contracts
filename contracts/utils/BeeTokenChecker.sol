
// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

library BeeTokenChecker {
    function containsKey(mapping(address => bool) storage aMap, address aKey) internal view returns (bool) {
        return aMap[aKey];
    }
}