package request

// MPSMintRequest MPS铸币请求
type MPSMintRequest struct {
	ToAddresses []string `json:"toAddresses" binding:"required"` // 接收地址列表
	Amount      string   `json:"amount" binding:"required"`      // 铸币数量
}

// MPSTransferRequest MPS转账请求
type MPSTransferRequest struct {
	To     string `json:"to" binding:"required"`     // 接收地址
	Amount string `json:"amount" binding:"required"` // 转账数量
}

// MPSStoreHashRequest 存储哈希请求
type MPSStoreHashRequest struct {
	Hash string `json:"hash" binding:"required"` // 哈希值
}

// MPSStoreReviewRequest 存储审稿内容请求
type MPSStoreReviewRequest struct {
	Content string `json:"content" binding:"required"` // 审稿内容
}

// MPSRegisterUserRequest 注册用户请求
type MPSRegisterUserRequest struct {
	UserAddress string `json:"userAddress" binding:"required"` // 用户地址
}

// MyNFTMintRequest NFT铸造请求
type MyNFTMintRequest struct {
	To  string `json:"to" binding:"required"`  // 接收地址
	URI string `json:"uri" binding:"required"` // 元数据URI
}

// MyNFTBulkMintRequest NFT批量铸造请求
type MyNFTBulkMintRequest struct {
	Recipients []string `json:"recipients" binding:"required"` // 接收地址列表
	URIs       []string `json:"uris" binding:"required"`       // 元数据URI列表
}

// MyNFTSetMetadataRequest 设置合约元数据请求
type MyNFTSetMetadataRequest struct {
	MetadataURI string `json:"metadataURI" binding:"required"` // 元数据URI
}

// MyNFTUpdateRoyaltyRequest 更新版税请求
type MyNFTUpdateRoyaltyRequest struct {
	RoyaltyPercentage uint64 `json:"royaltyPercentage" binding:"required"` // 版税百分比
}
