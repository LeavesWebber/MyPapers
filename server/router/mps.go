package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type MPSRouter struct{}

func InitMPSRouter(Router *gin.RouterGroup) {
	mpsApi := api.MPSApi{}
	Router.POST("/mps/mint", mpsApi.Mint)
	Router.POST("/mps/transfer", mpsApi.Transfer)
	Router.GET("/mps/balance-of/:address", mpsApi.GetBalanceOf)
	Router.POST("/mps/store-hash", mpsApi.StoreHash)
	Router.GET("/mps/recipient-by-hash/:hash", mpsApi.GetRecipientByHash)
	Router.POST("/mps/store-review", mpsApi.StoreReview)
	Router.GET("/mps/review-by-hash/:content", mpsApi.GetReviewByHash)
	Router.POST("/mps/register-user", mpsApi.RegisterUser)
}
