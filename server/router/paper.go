package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type PaperRouter struct{}

// InitPaperRouter 投稿相关api
func (p *PaperRouter) InitPaperRouter(Router *gin.RouterGroup) {
	paperRouter := Router.Group("paper")
	//PaperRouterWithoutRecord := Router.Group("Paper").Use(middleware.OperationRecord())
	paperApi := api.ApiGroupApp.PaperApi
	{
		paperRouter.POST("submit", paperApi.SubmitPaper)                        //投稿
		paperRouter.GET("detail", paperApi.GetPaper)                            // 查看投稿详情
		paperRouter.GET("selfList", paperApi.GetAllSelfPapers)                  // 查询自己的投稿列表
		paperRouter.GET("list", paperApi.GetAllPapers)                          // 查询投稿列表
		paperRouter.PUT("update", paperApi.UpdatePaper)                         // 修改投稿信息
		paperRouter.DELETE("delete", paperApi.DeletePaper)                      // 删除投稿
		paperRouter.GET("version", paperApi.GetPaperVersions)                   // 获取某个投稿的所有版本
		paperRouter.GET("honoraryCertificate", paperApi.GetHonoraryCertificate) // 获取荣誉证书
		paperRouter.POST("publish", paperApi.PublishPaper)                      // 发布投稿
		paperRouter.POST("addPaperViewer", paperApi.AddPaperViewer)             // 增加投稿可查看者
		paperRouter.GET("checkPaperViewer", paperApi.CheckPaperViewer)          // 查看用户是否有权限查看投稿
		paperRouter.GET("myNFTs", paperApi.GetMyNFTs)                           // 获取我的NFT
		paperRouter.PUT("updatePrice", paperApi.UpdatePrice)                    // 更新价格
		paperRouter.GET("getNFTInfo", paperApi.GetNFTInfoByTokenId)             // 根据tokenId获取NFT信息
		paperRouter.PUT("updatePaperUserId", paperApi.UpdatePaperUserId)        // 修改投稿对应的user_id
	}
}
