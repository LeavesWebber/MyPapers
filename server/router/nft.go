package router

import (
	v1 "server/api"

	"github.com/gin-gonic/gin"
)

type NFTRouter struct{}

func (n *NFTRouter) InitNFTRouter(Router *gin.RouterGroup) {
	nftRouter := Router.Group("nft")
	nftApi := v1.ApiGroupApp.NFTApi
	{
		// 基础NFT操作
		nftRouter.POST("mint", nftApi.MintNFT)                    // 铸造NFT
		nftRouter.GET("my-nfts", nftApi.GetMyNFTs)                // 获取我的NFT列表
		nftRouter.GET("token/:tokenId", nftApi.GetNFTByTokenId)   // 根据Token ID获取NFT信息
		nftRouter.GET("metadata/:tokenId", nftApi.GetNFTMetadata) // 获取NFT元数据
		nftRouter.PUT("set-price", nftApi.SetNFTPrice)            // 设置NFT价格

		// 市场相关
		nftRouter.GET("marketplace", nftApi.GetNFTMarketplace) // 获取NFT市场列表
		nftRouter.POST("buy", nftApi.BuyNFT)                   // 购买NFT
		nftRouter.POST("sell", nftApi.SellNFT)                 // 出售NFT

		// 统计信息
		nftRouter.GET("stats", nftApi.GetNFTStats) // 获取NFT统计信息
	}
}
