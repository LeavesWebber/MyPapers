package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type JournalRouter struct{}

// InitJournalRouter 期刊相关api
func (j *JournalRouter) InitJournalRouter(Router *gin.RouterGroup) {
	journalRouter := Router.Group("journal")
	//journalRouterWithoutRecord := Router.Group("Journal").Use(middleware.OperationRecord())
	journalApi := api.ApiGroupApp.JournalApi
	{
		journalRouter.POST("create", journalApi.CreateJournal)      // 创建会议
		journalRouter.DELETE("delete", journalApi.DeleteJournal)    // 删除会议
		journalRouter.PUT("update", journalApi.UpdateJournal)       // 修改信息
		journalRouter.GET("selfList", journalApi.GetJournalsByUser) // 查询自己所在的期刊列表
		//journalRouter.POST("/committee/join/", journalApi.JoinJournal)              // 申请加入会议
		//journalRouter.POST("/committee/quit/", journalApi.QuitJournal)              // 退出会议
		//journalRouter.GET("president/list", journalApi.JournalPresidentList) // 查看会议主席(或副主席)列表  多出来了
		//journalRouter.POST("designateChief", journalApi.DesignateChief)   // 指定会议主席
		//journalRouter.POST("designateExpert", journalApi.DesignateExpert) // 指定会议专家
		journalRouter.POST("issue/create", journalApi.CreateJournalIssue)   // 创建Issue
		journalRouter.DELETE("issue/delete", journalApi.DeleteJournalIssue) // 删除Issue
		journalRouter.PUT("issue/update", journalApi.UpdateJournalIssue)    // 修改Issue
		journalRouter.GET("level", journalApi.GetLevelInJournal)            // 获取用户在期刊的level
	}
}
