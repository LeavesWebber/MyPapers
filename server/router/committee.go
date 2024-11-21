package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type CommitteeRouter struct{}

// InitCommitteeRouter 委员会相关api
func (c *CommitteeRouter) InitCommitteeRouter(Router *gin.RouterGroup) {
	committeeRouter := Router.Group("committee")
	//committeeRouterWithoutRecord := Router.Group("Committee").Use(middleware.OperationRecord())
	committeeApi := api.ApiGroupApp.CommitteeApi
	{
		committeeRouter.POST("create", committeeApi.CreateCommittee)      // 创建委员会
		committeeRouter.DELETE("delete", committeeApi.DeleteCommittee)    // 删除委员会
		committeeRouter.PUT("update", committeeApi.UpdateCommittee)       // 修改委员会信息
		committeeRouter.GET("selfList", committeeApi.GetCommitteesByUser) // 查询自己所在的委员会列表
		////committeeRouter.POST("/committee/join/", committeeApi.JoinCommittee)              // 申请加入委员会
		////committeeRouter.POST("/committee/quit/", committeeApi.QuitCommittee)              // 退出委员会
		//committeeRouter.GET("president/list", committeeApi.CommitteePresidentList) // 查看委员会主席(或副主席)列表  多出来了
		////committeeRouter.POST("designateChief", committeeApi.DesignateChief)   // 指定委员会主席
		////committeeRouter.POST("designateExpert", committeeApi.DesignateExpert) // 指定委员会专家
	}
}
