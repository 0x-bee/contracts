// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Administered is AccessControl, Ownable {

    bytes32 public constant MEMBER_ROLE = keccak256("MEMBER");

    constructor () {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setRoleAdmin(MEMBER_ROLE, DEFAULT_ADMIN_ROLE);
    }

    modifier onlyAdmin() {
        require(isAdmin(msg.sender), "Restricted to admins.");
        _;
    }

    modifier onlyMember() {
        require(isMemeber(msg.sender), "Restricted to members.");
        _;
    }

    function isAdmin(address account) public virtual view returns (bool) {
        return hasRole(DEFAULT_ADMIN_ROLE, account);
    }

    function isMemeber(address account) public virtual view returns (bool) {
        return hasRole(MEMBER_ROLE, account);
    }

    function addMember(address account) public virtual onlyAdmin {
        grantRole(MEMBER_ROLE, account);
    }

    function removeMember(address account) public virtual onlyAdmin {
        revokeRole(MEMBER_ROLE, account);
    }

    // override Ownable _transferOwnership 增加移交管理员角色逻辑
    
    function _transferOwnership(address newOwner) internal override {
        address oldOwner = owner();
        grantRole(DEFAULT_ADMIN_ROLE, newOwner);
        revokeRole(DEFAULT_ADMIN_ROLE, oldOwner);
        super._transferOwnership(newOwner);
    }
}