package router

import (
	"app/api"
	"github.com/gin-gonic/gin"
)

func InitBlockRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("block")
	{
		ConfigRouter.POST("/req", api.Select)
		ConfigRouter.POST("/allBlock", api.Select)
	}
}
