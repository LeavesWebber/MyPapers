package response

// MPSMintResponse MPS铸币响应
type MPSMintResponse struct {
	TransactionHash string `json:"transactionHash"` // 交易哈希
	BlockNumber     uint64 `json:"blockNumber"`     // 区块号
	GasUsed         uint64 `json:"gasUsed"`         // Gas使用量
}

// MPSTransferResponse MPS转账响应
type MPSTransferResponse struct {
	TransactionHash string `json:"transactionHash"` // 交易哈希
	BlockNumber     uint64 `json:"blockNumber"`     // 区块号
	GasUsed         uint64 `json:"gasUsed"`         // Gas使用量
}

// MPSBalanceResponse MPS余额响应
type MPSBalanceResponse struct {
	Address string `json:"address"` // 地址
	Balance string `json:"balance"` // 余额
}

// MPSStoreHashResponse 存储哈希响应
type MPSStoreHashResponse struct {
	TransactionHash string `json:"transactionHash"` // 交易哈希
	BlockNumber     uint64 `json:"blockNumber"`     // 区块号
	GasUsed         uint64 `json:"gasUsed"`         // Gas使用量
}

// MPSStoreReviewResponse 存储审稿响应
type MPSStoreReviewResponse struct {
	TransactionHash string `json:"transactionHash"` // 交易哈希
	BlockNumber     uint64 `json:"blockNumber"`     // 区块号
	GasUsed         uint64 `json:"gasUsed"`         // Gas使用量
}

// MPSRegisterUserResponse 注册用户响应
type MPSRegisterUserResponse struct {
	TransactionHash string `json:"transactionHash"` // 交易哈希
	BlockNumber     uint64 `json:"blockNumber"`     // 区块号
	GasUsed         uint64 `json:"gasUsed"`         // Gas使用量
}

// MyNFTMintResponse NFT铸造响应
type MyNFTMintResponse struct {
	TokenId         uint64 `json:"tokenId"`         // Token ID
	To              string `json:"to"`              // 接收地址
	URI             string `json:"uri"`             // 元数据URI
	TransactionHash string `json:"transactionHash"` // 交易哈希
	BlockNumber     uint64 `json:"blockNumber"`     // 区块号
	GasUsed         uint64 `json:"gasUsed"`         // Gas使用量
}

// MyNFTMintResult NFT铸造结果
type MyNFTMintResult struct {
	TokenId         uint64 `json:"tokenId"`         // Token ID
	To              string `json:"to"`              // 接收地址
	URI             string `json:"uri"`             // 元数据URI
	TransactionHash string `json:"transactionHash"` // 交易哈希
	BlockNumber     uint64 `json:"blockNumber"`     // 区块号
	GasUsed         uint64 `json:"gasUsed"`         // Gas使用量
}

// MyNFTBulkMintResponse NFT批量铸造响应
type MyNFTBulkMintResponse struct {
	Results []MyNFTMintResult `json:"results"` // 铸造结果列表
	Count   int               `json:"count"`   // 铸造数量
}

// MyNFTSetMetadataResponse 设置合约元数据响应
type MyNFTSetMetadataResponse struct {
	MetadataURI     string `json:"metadataURI"`     // 元数据URI
	TransactionHash string `json:"transactionHash"` // 交易哈希
	BlockNumber     uint64 `json:"blockNumber"`     // 区块号
	GasUsed         uint64 `json:"gasUsed"`         // Gas使用量
}

// MyNFTUpdateRoyaltyResponse 更新版税响应
type MyNFTUpdateRoyaltyResponse struct {
	RoyaltyPercentage uint64 `json:"royaltyPercentage"` // 版税百分比
	TransactionHash   string `json:"transactionHash"`   // 交易哈希
	BlockNumber       uint64 `json:"blockNumber"`       // 区块号
	GasUsed           uint64 `json:"gasUsed"`           // Gas使用量
}

// MyNFTTokenURIResponse NFT Token URI响应
type MyNFTTokenURIResponse struct {
	TokenId uint64 `json:"tokenId"` // Token ID
	URI     string `json:"uri"`     // 元数据URI
}

// MyNFTSupportsInterfaceResponse NFT接口支持响应
type MyNFTSupportsInterfaceResponse struct {
	InterfaceId string `json:"interfaceId"` // 接口ID
	Supported   bool   `json:"supported"`   // 是否支持
}

// MyNFTOwnerResponse NFT Owner响应
type MyNFTOwnerResponse struct {
	Owner string `json:"owner"` // Owner地址
}

// MyNFTBalanceResponse NFT余额响应
type MyNFTBalanceResponse struct {
	Address string `json:"address"` // 地址
	Balance uint64 `json:"balance"` // 余额
}

// MyNFTOwnerOfResponse NFT Token所有者响应
type MyNFTOwnerOfResponse struct {
	TokenId uint64 `json:"tokenId"` // Token ID
	Owner   string `json:"owner"`   // 所有者地址
}

// MyNFTTotalSupplyResponse NFT总供应量响应
type MyNFTTotalSupplyResponse struct {
	TotalSupply uint64 `json:"totalSupply"` // 总供应量
}
