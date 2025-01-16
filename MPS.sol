// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

// ERC20代币合约
contract MPS is ERC20 {
    using Counters for Counters.Counter;
    Counters.Counter private _hashCounter;
    const PopularizeAccount = 0x7978c92C6DE8aAd1210FeBe413C7206917999B3c;
    // 哈希值到账户地址的映射
    mapping(string => address) private _hashToAddress;
    // 审稿内容到账户地址的映射
    mapping(string => address) private _reviewToAddress;



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

    // 初始化函数，分配总量给初始用户
    constructor(uint256 initialSupply) ERC20("MPS", "MPS") {
        _mint(msg.sender, initialSupply);
    }

    // 获取哈希值对应的账户地址
    function getRecipientByHash(string memory hash) public view returns (address) {
        return _hashToAddress[hash];
    }

     function registerUser(address user) public {
        //推广账户地址
        address tokenSource=0x7978c92C6DE8aAd1210FeBe413C7206917999B3c;
        // 检查发送者是否有足够的代币
        require(balanceOf(tokenSource) >= 500 * 10 ** decimals(), "Insufficient tokens in the source account");
        // 从代币来源账户转移 500 个代币给用户
        _transfer(tokenSource, user, 500 * 10 ** decimals());
    }

    // 转账函数
    function transfer(address recipient, uint256 amount) public override returns (bool) {
        require(recipient != address(0), "Invalid address");
        super.transfer(recipient, amount);
        emit TransferWithHash(msg.sender, recipient, amount);
        return true;
    }
    
    // 铸币函数
    // function mint(address to, uint256 amount) public returns (bool) {
    //     _mint(to, amount);
    //     emit Mint(to, amount);
    //     return true;
    // }
    function mint(address[] memory toAddresses, uint256 amount) public returns (bool) {
        for (uint256 i = 0; i < toAddresses.length; i++) {
            _mint(toAddresses[i], amount);
            emit Mint(toAddresses[i], amount);
        }
        return true;
    }


    // 存储哈希值到区块链
    function storeHash(string memory hash) public {
        // require(_hashToAddress[hash] == address(0), "Hash already stored");
        require(_hashToAddress[hash] == address(0) || _hashToAddress[hash] == msg.sender, "Hash already stored");
        _hashToAddress[hash] = msg.sender;
        emit HashStored(hash, msg.sender);
    }

    // 扣减币函数
    function burnFrom(address account, uint256 amount) public {
        _burn(account, amount);
        emit Burn(account, amount);
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
}
