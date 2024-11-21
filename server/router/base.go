package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type BaseRouter struct{}

// InitBaseRouter 基本功能路由
func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) gin.IRoutes {
	baseRouter := Router.Group("user")
	committeeRouter := Router.Group("committee")
	conferenceRouter := Router.Group("conference")
	journalRouter := Router.Group("journal")
	paperRouter := Router.Group("paper")
	userApi := api.ApiGroupApp.UserApi
	committeeApi := api.ApiGroupApp.CommitteeApi
	conferenceApi := api.ApiGroupApp.ConferenceApi
	journalApi := api.ApiGroupApp.JournalApi
	PaperApi := api.ApiGroupApp.PaperApi
	{
		baseRouter.POST("register", userApi.Register) // 注册
		baseRouter.POST("login", userApi.Login)       // 登录
		//baseRouter.POST("captcha", userApi.Captcha)   // 生成验证码
		committeeRouter.GET("detail", committeeApi.GetCommittee)                                              // 查看委员会详情
		committeeRouter.GET("list", committeeApi.GetAllCommittees)                                            // 查询委员会列表
		conferenceRouter.GET("detail", conferenceApi.GetConference)                                           // 查看会议详情
		conferenceRouter.GET("list", conferenceApi.GetAllConferences)                                         // 查询会议列表
		conferenceRouter.GET("listByCommittee", conferenceApi.GetAllConferencesByCommittee)                   // 根据委员会查询会议列表
		journalRouter.GET("detail", journalApi.GetJournal)                                                    // 查看期刊详情
		journalRouter.GET("list", journalApi.GetAllJournals)                                                  // 查询期刊列表
		journalRouter.GET("listByCommittee", journalApi.GetAllJournalsByCommittee)                            // 根据委员会查询期刊列表
		paperRouter.GET("acceptPaperList", PaperApi.GetAllAcceptPapers)                                       // 查询已经审核通过的投稿列表
		paperRouter.GET("acceptPaperListByJournalAndTime", PaperApi.GetAllAcceptPapersByJournalAndTime)       // 按期刊和时间查询已经审核通过的投稿列表
		paperRouter.GET("acceptPaperListByConferenceAndTime", PaperApi.GetAllAcceptPapersByConferenceAndTime) // 按期刊和时间查询已经审核通过的投稿列表
		journalRouter.GET("issue/list", journalApi.GetAllJournalIssues)
		conferenceRouter.GET("issue/list", conferenceApi.GetAllConferenceIssues) // 查询会议Issue列表
	}
	return baseRouter
}
