package initialize

import (
	"server/global"
	"server/middleware"
	"server/router"

	"github.com/gin-gonic/gin"
)

// Routers 路由总入口
func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp
	Router.Static("/public", "./public")
	Router.Static("/image", "./image")
	Router.Static("/assets", "./assets")
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors())        // 直接放行全部跨域请求
	Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.MPS_LOG.Info("use middleware cors")

	PublicGroup := Router.Group("mypapers")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 基础功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("mypapers")
	//PrivateGroup.Use(middleware.JWTAuthMiddleware()).Use(middleware.CasbinHandler())
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)       // 用户功能路由
		systemRouter.InitAuthorityRouter(PrivateGroup)  // 注册角色路由
		systemRouter.InitMenuRouter(PrivateGroup)       // 注册菜单路由
		systemRouter.InitCommitteeRouter(PrivateGroup)  // 委员会功能路由
		systemRouter.InitConferenceRouter(PrivateGroup) // 会议功能路由
		systemRouter.InitJournalRouter(PrivateGroup)    // 期刊功能路由
		systemRouter.InitPaperRouter(PrivateGroup)      // 投稿功能路由
		systemRouter.InitReviewRouter(PrivateGroup)     // 评审功能路由
		router.InitMPSRouter(PrivateGroup)              // MPS功能路由
		router.InitMyNFTRouter(PrivateGroup)            // MyNFT功能路由
	}
	global.MPS_LOG.Info("router register success")
	return Router
}
