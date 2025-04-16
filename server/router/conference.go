package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type ConferenceRouter struct{}

// InitConferenceRouter 会议相关api
func (con *ConferenceRouter) InitConferenceRouter(Router *gin.RouterGroup) {
	conferenceRouter := Router.Group("conference")
	//conferenceRouterWithoutRecord := Router.Group("Conference").Use(middleware.OperationRecord())
	conferenceApi := api.ApiGroupApp.ConferenceApi
	{
		conferenceRouter.POST("create", conferenceApi.CreateConference)      // 创建会议
		conferenceRouter.DELETE("delete", conferenceApi.DeleteConference)    // 删除会议
		conferenceRouter.PUT("update", conferenceApi.UpdateConference)       // 修改信息
		conferenceRouter.GET("selfList", conferenceApi.GetConferencesByUser) // 查询自己所在的期刊列表
		//conferenceRouter.POST("/committee/join/", conferenceApi.JoinConference)              // 申请加入会议
		//conferenceRouter.POST("/committee/quit/", conferenceApi.QuitConference)              // 退出会议
		//conferenceRouter.GET("president/list", conferenceApi.ConferencePresidentList) // 查看会议主席(或副主席)列表  多出来了
		//conferenceRouter.POST("designateChief", conferenceApi.DesignateChief)   // 指定会议主席
		//conferenceRouter.POST("designateExpert", conferenceApi.DesignateExpert) // 指定会议专家
		conferenceRouter.POST("issue/create", conferenceApi.CreateConferenceIssue)   // 创建Issue
		conferenceRouter.DELETE("issue/delete", conferenceApi.DeleteConferenceIssue) // 删除Issue
		conferenceRouter.PUT("issue/update", conferenceApi.UpdateConferenceIssue)    // 修改Issue
		conferenceRouter.GET("level", conferenceApi.GetLevelInConference)            // 获取用户在期刊的level
	}
}
