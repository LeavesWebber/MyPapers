// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/access/Ownable2Step.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/utils/StorageSlot.sol";
import "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * @title MPERproxy
 * @dev 支持分阶段治理的透明代理合约，专为MPER代币设计
 *
 * 架构说明：
 * 1. 代理合约本身不可升级，保持地址永久不变
 * 2. 逻辑合约可通过升级函数更换
 * 3. 管理员权限可转移给更复杂的治理合约
 */
contract MPERproxy is ERC1967Proxy, Ownable2Step {
    // 事件定义
    event Upgraded(address indexed newImplementation, uint256 timestamp, string reason);
    event Initialized(address indexed implementation, address indexed admin);
    event UpgradeCancelled(bytes32 indexed proposalId);
    event UpgradeDelayChanged(uint256 oldDelay, uint256 newDelay);

    // 存储槽位
    bytes32 private constant _IMPLEMENTATION_INTERFACE_ID_SLOT = keccak256("mper.proxy.implementation.interface.id");
    bytes32 private constant _TIMELOCK_SLOT = keccak256("mper.proxy.timelock");

    // 常量
    uint256 public constant MIN_UPGRADE_DELAY = 1 days;
    uint256 public constant MAX_REASON_LENGTH = 256;

    // 状态变量
    TimelockController public timelock;
    uint256 public upgradeDelay;

    // 提案跟踪
    mapping(bytes32 => bool) public isProposalActive;

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
        //require(data.length == 0 || data.length >= 4, "Invalid calldata");
        require(address(logic).code.length > 0, "Logic address is not a contract");
        require(admin != address(0), "Admin address cannot be zero");
        require(initialDelay >= MIN_UPGRADE_DELAY, "Initial delay too short");
    
        // The Ownable(admin) call above sets the initial owner.
        // The subsequent _transferOwnership(address(timelock)) will transfer ownership from 'admin' to 'timelock'.

        // 存储逻辑合约的接口ID
        bytes4 interfaceId = type(IERC165).interfaceId;
        StorageSlot.getBytes32Slot(_IMPLEMENTATION_INTERFACE_ID_SLOT).value = bytes32(uint256(uint32(interfaceId)));

        // 创建TimelockController，设置合约自身为执行人
        address[] memory executors = new address[](1);
        executors[0] = address(this);
        address[] memory proposers = new address[](1);  // 修改为1个提案者
        proposers[0] = admin;  // 将管理员设为初始提案者
        timelock = new TimelockController(initialDelay, proposers, executors, admin);  // 将admin作为初始管理员
        StorageSlot.getAddressSlot(_TIMELOCK_SLOT).value = address(timelock);

        // 转移所有权
        _transferOwnership(address(timelock));

        // 设置升级延迟
        upgradeDelay = initialDelay;

        emit Initialized(logic, admin);
    }

    /**
     * @dev 设置升级延迟
     * @param newDelay 新的升级延迟
     */
    function setUpgradeDelay(uint256 newDelay) external onlyOwner {
        require(newDelay >= MIN_UPGRADE_DELAY, "Delay too short");
        emit UpgradeDelayChanged(upgradeDelay, newDelay);
        upgradeDelay = newDelay;
    }

    /**
     * @dev 提案升级逻辑合约
     * @param newImplementation 新逻辑合约地址
     * @param data 初始化数据
     * @param reason 升级原因
     */
    function proposeUpgrade(
        address newImplementation,
        bytes memory data,
        string memory reason
    ) external onlyOwner returns (bytes32 proposalId) {
        require(address(newImplementation).code.length > 0, "New implementation is not a contract");
        require(bytes(reason).length > 0 && bytes(reason).length <= MAX_REASON_LENGTH, "Invalid reason length");

        _validateUpgradeData(newImplementation, data);

        uint256 eta = block.timestamp + upgradeDelay;
        bytes memory upgradeCall = abi.encodeWithSignature(
            "upgradeToAndCall(address,bytes,string)",
            newImplementation,
            data,
            reason
        );

        bytes32 predecessor = bytes32(0); // 新增 predecessor
        bytes32 proposalSalt = bytes32(keccak256(abi.encodePacked(block.timestamp)));
        proposalId = keccak256(abi.encode(address(this), 0, upgradeCall, proposalSalt, eta));
        timelock.schedule(address(this), 0, upgradeCall, predecessor, proposalSalt, eta); // 修正参数顺序和类型
        isProposalActive[proposalId] = true;
    }

    /**
     * @dev 取消升级提案
     * @param proposalId 提案ID
     */
    function cancelUpgrade(bytes32 proposalId) external onlyOwner {
        require(isProposalActive[proposalId], "Proposal is not active");
        timelock.cancel(proposalId);
        delete isProposalActive[proposalId];
        emit UpgradeCancelled(proposalId);
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
     * @dev 验证升级数据（针对MPER合约）
     */
    function _validateUpgradeData(address newImplementation, bytes memory data) internal view {
        // 基本验证
        require(data.length >= 4, "Invalid calldata");
        
        // 检查是否为initialize调用
        bytes4 selector;
        assembly {
            selector := mload(add(data, 32))
        }
        
        if (selector == bytes4(keccak256("initialize(uint256)"))) {
            // 解码initialize参数
            uint256 initialSupply;
            assembly {
                // 跳过前4字节的函数选择器
                initialSupply := mload(add(add(data, 4), 32))
            }
            require(initialSupply > 0, "Initial supply must be positive");
        }
        
        // 可以添加更多针对MPER合约特定函数的验证
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
