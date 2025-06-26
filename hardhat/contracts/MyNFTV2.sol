// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./MyNFT.sol";

contract MyNFTV2 is MyNFT {
    mapping(address => bool) public whitelist;
    uint256 public mintFee;

    // 新的初始化函数，设置 baseURI
    function initializeV2(string memory baseURI_) public reinitializer(2) {
        _setBaseURI(baseURI_);
        mintFee = 0.01 ether;
    }

    // 设置版税百分比（10000 = 100%）
    function initializeRoyalty(uint96 royaltyFraction_) public onlyOwner {
        require(royaltyFraction_ <= 10000, "Too high");
        royaltyPercentage = royaltyFraction_;
    }

    function addToWhitelist(address[] calldata users) public onlyOwner {
        for (uint256 i = 0; i < users.length; i++) {
            whitelist[users[i]] = true;
        }
    }

    function mintWithFee(address to, string memory uri) public payable {
        require(msg.value >= mintFee, "Insufficient fee");
        require(whitelist[msg.sender], "Not in whitelist");
        safeMint(to, uri);
    }

    // OpenZeppelin IERC2981 兼容
    function royaltyInfo(uint256, uint256 salePrice)
        public
        view
        returns (address receiver, uint256 royaltyAmount)
    {
        receiver = owner();
        royaltyAmount = (salePrice * royaltyPercentage) / 10000;
    }

    // baseURI view 方法
    function baseURI() public view returns (string memory) {
        return _baseURI();
    }

    uint256[47] private __gap; // 预留存储空间
}