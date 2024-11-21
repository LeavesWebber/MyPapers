package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type ReviewRouter struct{}

// InitReviewRouter 投稿相关api
func (r *ReviewRouter) InitReviewRouter(Router *gin.RouterGroup) {
	reviewRouter := Router.Group("review")
	//ReviewRouterWithoutRecord := Router.Group("Review").Use(middleware.OperationRecord())
	reviewApi := api.ApiGroupApp.ReviewApi
	{
		reviewRouter.GET("list", reviewApi.GetAllReviews)    // 查询审核列表
		reviewRouter.POST("allot", reviewApi.AllotReviewers) // 分配审核人
		reviewRouter.POST("submit", reviewApi.SubmitReview)  //审核
	}
}
