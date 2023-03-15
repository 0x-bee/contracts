
// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

library BeeTokenChecker {
    function containsValue(bytes32[] memory array, bytes32 aValue) internal pure returns (bool) {
        for (uint i = 0; i < array.length; i++) {
            if (array[i] == aValue) {
                return true;
            }
        }
        return false;
    }
}