// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract MyNFT is 
    Initializable,
    ERC721Upgradeable, 
    ERC721EnumerableUpgradeable, 
    ERC721URIStorageUpgradeable, 
    OwnableUpgradeable 
{
    uint256 private _nextTokenId;
    string public contractMetadataURI; // 新的元数据功能
    uint256 public royaltyPercentage;  // 可升级的版税功能

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address initialOwner) initializer public {
        __ERC721_init("MyNFT", "MPER");
        __ERC721Enumerable_init();
        __ERC721URIStorage_init();
        __Ownable_init(initialOwner);
        
        _nextTokenId = 0;
        royaltyPercentage = 10; // 默认版税百分比
    }
    // 在合约中添加事件
    event Minted(address indexed to, uint256 tokenId);
    // 添加安全铸币功能
    function safeMint(address to, string memory uri) public onlyOwner {
        uint256 tokenId = _nextTokenId;
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);
        _nextTokenId++;
        emit Minted(to, tokenId); // 添加事件以便追踪
    }

    // 添加批量铸币功能
    function bulkMint(address[] memory recipients, string[] memory uris) public onlyOwner {
        require(recipients.length == uris.length, "Array length mismatch");
        for (uint256 i = 0; i < recipients.length; i++) {
            safeMint(recipients[i], uris[i]);
        }
    }
    
    // 新增功能：更新合约元数据
    function setContractMetadata(string memory newURI) public onlyOwner {
        contractMetadataURI = newURI;
    }
    
    // 新增功能：更新版税设置
    function updateRoyaltyPercentage(uint256 newPercentage) public onlyOwner {
        require(newPercentage <= 20, "Royalty too high");
        royaltyPercentage = newPercentage;
    }

    // 以下函数是必要的 Solidity 重写
    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
        returns (address)
    {
        return super._update(to, tokenId, auth);
    }

    function _increaseBalance(address account, uint128 value)
        internal
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
    {
        super._increaseBalance(account, value);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721Upgradeable, ERC721URIStorageUpgradeable)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable, ERC721URIStorageUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
    
    // 新增：安全升级预留空间
    uint256[50] private __gap;
}