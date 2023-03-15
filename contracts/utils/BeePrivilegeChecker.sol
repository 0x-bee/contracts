library BeePrivilegeChecker {
    function containsKey(mapping(address => bool) storage aMap, address aKey) internal view returns (bool) {
        return aMap[aKey];
    }
}