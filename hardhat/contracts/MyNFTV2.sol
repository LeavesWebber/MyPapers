// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./MyNFT.sol";

contract MyNFTV2 is MyNFT {
    // 新状态变量 - 添加到原有变量的下方
    mapping(address => bool) public whitelist;
    uint256 public mintFee;
    
    // 添加新的初始化函数（可选）
    function initializeV2() public reinitializer(2) {
        mintFee = 0.01 ether;
    }
    
    // 添加白名单功能
    function addToWhitelist(address[] calldata users) public onlyOwner {
        for (uint256 i = 0; i < users.length; i++) {
            whitelist[users[i]] = true;
        }
    }
    
    // 收费铸币功能
    function mintWithFee(address to, string memory uri) public payable {
        require(msg.value >= mintFee, "Insufficient fee");
        require(whitelist[msg.sender], "Not in whitelist");
        safeMint(to, uri);
    }
    
    // 新增版税支持
    function royaltyInfo(uint256 tokenId, uint256 salePrice) 
        public 
        view 
        returns (address receiver, uint256 royaltyAmount) 
    {
        receiver = owner();
        royaltyAmount = (salePrice * royaltyPercentage) / 100;
    }
    
    // 预留存储空间
    uint256[48] private __gap;
}