// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

// ERC20代币合约 - 可升级版本
contract MPS is Initializable, ERC20Upgradeable, OwnableUpgradeable { 
    // 哈希值到账户地址的映射
    mapping(string => address) private _hashToAddress;
    // 审稿内容到账户地址的映射
    mapping(string => address) private _reviewToAddress;
    // 新增：跟踪用户是否已注册
    mapping(address => bool) private _hasRegistered;

    // 事件：转账
    event TransferWithHash(address indexed from, address indexed to, uint256 value);
    // 事件：铸币
    event Mint(address indexed to, uint256 value);
    // 事件：存储哈希值
    event HashStored(string hash, address indexed sender);
    // 事件：扣减币
    event Burn(address indexed from, uint256 value);
    // 事件：存储审稿内容
    event ReviewStored(string content);
   
    function initialize(uint256 initialSupply) public initializer {
        __ERC20_init("MPS", "MPS");
        __Ownable_init(msg.sender); 

        if (initialSupply > 0) {
            _mint(msg.sender, initialSupply); // 初始代币分配给部署者 (也是初始 owner)
        }
    }

    // 获取哈希值对应的账户地址
    function getRecipientByHash(string memory hash) public view returns (address) {
        return _hashToAddress[hash];
    }

    // 用户注册函数 - 使用 onlyOwner，并从合约自身地址发放奖励
    function registerUser(address user) public onlyOwner { // 改为 onlyOwner
        require(!_hasRegistered[user], "MPS: User has already registered");
        // 检查合约自身是否有足够的代币
        uint256 requiredAmount = 500 * (10 ** decimals());
        require(balanceOf(address(this)) >= requiredAmount, "MPS: Insufficient tokens in the contract for reward");
        
        _hasRegistered[user] = true;
        // 从合约自身地址转移 500 个代币给用户
        _transfer(address(this), user, requiredAmount);
    }

    // 转账函数
    function transfer(address recipient, uint256 amount) public override returns (bool) {
        require(recipient != address(0), "ERC20: transfer to the zero address");
        super.transfer(recipient, amount);
        emit TransferWithHash(msg.sender, recipient, amount);
        return true;
    }
    
    // 铸币函数 - 使用 onlyOwner
    function mint(address[] memory toAddresses, uint256 amount) public onlyOwner { // 改为 onlyOwner
        for (uint256 i = 0; i < toAddresses.length; i++) {
            _mint(toAddresses[i], amount);
            emit Mint(toAddresses[i], amount);
        }
    }

    // 存储哈希值到区块链
    function storeHash(string memory hash) public {
        require(_hashToAddress[hash] == address(0) || _hashToAddress[hash] == msg.sender, "MPS: Hash already stored or not owner");
        _hashToAddress[hash] = msg.sender;
        emit HashStored(hash, msg.sender);
    }

    // 存储审稿内容
    function storeReview(string memory content) public {
        _reviewToAddress[content] = msg.sender;
        emit ReviewStored(content);
    }

    // 获取审稿内容的账户地址
    function getReviewByHash(string memory content) public view returns (address) {
        return _reviewToAddress[content];
    }
        // 为升级预留存储空间 (MPS自身的gap)
    uint256[50] private __gap; // 注意：OwnableUpgradeable 会占用一个存储槽给 _owner
}
