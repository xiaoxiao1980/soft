package router

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

func InitConfigRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("config")
	{
		ConfigRouter.POST("filepath", api.Select)
	}
}
