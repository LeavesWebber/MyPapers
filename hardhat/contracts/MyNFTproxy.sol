// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/access/Ownable2Step.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/utils/StorageSlot.sol";
import "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * @title MyNFTproxy
 * @dev 透明代理合约，使用TimelockController进行升级控制
 */
contract MyNFTproxy is ERC1967Proxy, Ownable2Step {
    // 事件定义
    event Upgraded(address indexed newImplementation, uint256 timestamp, string reason);
    event Initialized(address indexed implementation, address indexed admin);

    // 存储槽位
    bytes32 private constant _IMPLEMENTATION_INTERFACE_ID_SLOT = keccak256("mps.proxy.implementation.interface.id");
    bytes32 private constant _TIMELOCK_SLOT = keccak256("mps.proxy.timelock");

    // 状态变量
    TimelockController public timelock;

    // 添加receive函数以接收以太币
    receive() external payable {}

    /**
     * @dev 初始化代理合约
     * @param logic 逻辑合约地址
     * @param admin 初始管理员地址
     * @param data 初始化调用数据
     * @param initialDelay 初始升级延迟
     */
    constructor(
        address logic,
        address admin,
        bytes memory data,
        uint256 initialDelay
    ) ERC1967Proxy(logic, data) Ownable(admin) {
        require(address(logic).code.length > 0, "Logic address is not a contract");
        require(admin != address(0), "Admin address cannot be zero");

        // 存储逻辑合约的接口ID
        bytes4 interfaceId = type(ERC165).interfaceId;
        StorageSlot.getBytes32Slot(_IMPLEMENTATION_INTERFACE_ID_SLOT).value = bytes32(uint256(uint32(interfaceId)));

        // 创建TimelockController，设置合约自身为执行人
        address[] memory executors = new address[](1);
        executors[0] = address(this);
        address[] memory proposers = new address[](1);
        proposers[0] = admin;
        timelock = new TimelockController(initialDelay, proposers, executors, admin);
        StorageSlot.getAddressSlot(_TIMELOCK_SLOT).value = address(timelock);

        // 转移所有权
        _transferOwnership(address(timelock));

        emit Initialized(logic, admin);
    }

    /**
     * @dev 升级逻辑合约
     */
    function upgradeToAndCall(
        address newImplementation,
        bytes memory data,
        string memory reason
    ) external payable {
        require(msg.sender == address(timelock), "Unauthorized");
        ERC1967Utils.upgradeToAndCall(newImplementation, data);
        _verifyImplementation(newImplementation);
        emit Upgraded(newImplementation, block.timestamp, reason);
    }

    /**
     * @dev 验证升级数据
     */
    function _validateUpgradeData(address newImplementation, bytes memory data) internal view {
        // 检查是否为initialize调用
        bytes4 selector;
        assembly {
            selector := mload(add(data, 32))
        }
        
        if (selector == bytes4(keccak256("initialize(uint256)"))) {
            uint256 initialSupply;
            assembly {
                initialSupply := mload(add(add(data, 4), 32))
            }
            require(initialSupply > 0, "Initial supply must be positive");
        }
    }

    /**
     * @dev 验证实现合约接口
     */
    function _verifyImplementation(address newImplementation) internal view {
        require(address(newImplementation).code.length > 0, "New implementation is not a contract");
        bytes4 interfaceId;
        bytes32 slot = _IMPLEMENTATION_INTERFACE_ID_SLOT;
        assembly {
            interfaceId := sload(slot)
        }
        
        if (interfaceId != 0x00000000) {
            require(ERC165(newImplementation).supportsInterface(interfaceId), "Interface not supported");
        }
    }

    /**
     * @dev 获取当前实现地址
     */
    function implementation() external view returns (address) {
        return _implementation();
    }
}