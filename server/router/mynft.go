package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

func InitMyNFTRouter(Router *gin.RouterGroup) {
	mynftApi := api.MyNFTApi{}

	// MyNFT 铸造相关接口
	Router.POST("/mynft/mint", mynftApi.MintNFT)                           // 铸造单个NFT
	Router.POST("/mynft/bulk-mint", mynftApi.BulkMintNFT)                  // 批量铸造NFT
	Router.POST("/mynft/set-metadata", mynftApi.SetContractMetadata)       // 设置合约元数据
	Router.POST("/mynft/update-royalty", mynftApi.UpdateRoyaltyPercentage) // 更新版税百分比

	// MyNFT 查询相关接口
	Router.GET("/mynft/token-uri", mynftApi.GetTokenURI)                // 查询token URI
	Router.GET("/mynft/supports-interface", mynftApi.SupportsInterface) // 查询接口支持
	Router.GET("/mynft/owner", mynftApi.GetOwner)                       // 查询合约owner
	Router.GET("/mynft/balance-of", mynftApi.GetBalanceOf)              // 查询地址余额
	Router.GET("/mynft/owner-of", mynftApi.GetOwnerOf)                  // 查询token所有者
	Router.GET("/mynft/total-supply", mynftApi.GetTotalSupply)          // 查询总供应量
}
